package zorm

import (
	"fmt"
	"github.com/gzorm/commons/sharding"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"gorm.io/plugin/prometheus"
	"time"
)

func NewClient(config *DBConfig) *gorm.DB {
	dbMain := getConn(config.MainSource.SchemaConfig, config.Mode)

	for _, shading := range config.Sharding {
		var args []interface{}
		for _, table := range shading.Tables {
			args = append(args, table)
		}
		middleware := sharding.Register(sharding.Config{
			ShardingKey:    shading.ShardingKey,
			NumberOfShards: uint(shading.NumberOfShards),
			Name:           shading.Name,
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

func getConn(config SchemaConfig, mode string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=20s",
		config.Username, config.Password, config.Host,
		config.Port, config.DBName,
	)
	Db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		Logger: NewGormLogger(mode),
	})
	if err != nil {
		panic(err)
	}
	db, err := Db.DB()

	db.SetMaxIdleConns(config.MaxIdelConn)
	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetConnMaxLifetime(5 * time.Minute)        // 设置连接的最大存活时间
	db.SetConnMaxIdleTime(5 * time.Minute)        // 设置连接的最大空闲时间
	db.Exec("SET innodb_lock_wait_timeout = 50;") // 设置锁等待超时
	if err != nil {
		panic(err)
	}
	return Db
}
