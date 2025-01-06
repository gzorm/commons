package zorm

import (
	"github.com/gzorm/commons/sharding"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"gorm.io/plugin/prometheus"
	"time"
)

func NewRangeShardingClient(config *DBConfig) *gorm.DB {
	dbMain := getConn(config.MainSource.SchemaConfig, config.Mode)

	for _, shading := range config.Sharding {
		var args []interface{}
		for _, table := range shading.Tables {
			args = append(args, table)
		}
		middleware := sharding.Register(sharding.Config{
			ShardingKey: shading.ShardingKey,
			//NumberOfShards: uint(shading.NumberOfShards),
			ShardingAlgorithm: func(columnValue interface{}) (suffix string, err error) {
				var tm int64
				if v, ok := columnValue.(int64); ok {
					tm = v
				} else if v, ok := columnValue.(int); ok {
					tm = int64(v)
				} else if v, ok := columnValue.(uint64); ok {
					tm = int64(v)
				} else if v, ok := columnValue.(uint); ok {
					tm = int64(v)
				}
				t := time.Unix(int64(tm), 0)
				return "_" + t.Format("2006_01"), err
			},
			Name: shading.Name,
		}, args...)
		err := dbMain.Use(middleware)
		if err != nil {
			panic(err)
		}
	}
	var dbMainRes []gorm.Dialector
	for _, replica := range config.MainSource.Replicas {
		dbMainRes = append(dbMainRes, getConn(replica, config.Mode).Dialector)
	}

	mainSource := dbresolver.Register(
		dbresolver.Config{
			Replicas: dbMainRes,
			Policy:   dbresolver.RandomPolicy{},
		},
	)
	for _, source := range config.OtherSource {

		var res []gorm.Dialector
		for _, replica := range source.Replicas {
			db := getConn(replica, config.Mode)
			res = append(res, db.Dialector)
		}
		var s []gorm.Dialector
		db := getConn(source.SchemaConfig, config.Mode)
		s = append(s, db.Dialector)
		var args []interface{}
		for _, table := range source.Tables {
			args = append(args, table)
		}
		mainSource.Register(
			dbresolver.Config{
				Sources:  s,
				Replicas: res,
				Policy:   dbresolver.RandomPolicy{},
			}, args...)
	}

	err := dbMain.Use(mainSource)

	if err != nil {
		panic(err)
	}

	err = dbMain.Use(prometheus.New(prometheus.Config{
		DBName:          config.Prometheus.Name,                    // 使用 `DBName` 作为指标 label
		RefreshInterval: uint32(config.Prometheus.RefreshInterval), // 指标刷新频率（默认为 15 秒）
		//	PushAddr:        "prometheus pusher address", // 如果配置了 `PushAddr`，则推送指标
		StartServer:    true,                                     // 启用一个 http 服务来暴露指标
		HTTPServerPort: uint32(config.Prometheus.HTTPServerPort), // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		}, // 用户自定义指标
	}))

	if err != nil {
		panic(err)
	}

	return dbMain
}
