package zorm

import (
	"fmt"
	"github.com/gzorm/common/sharding"
	"github.com/gzorm/common/zorm/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"gorm.io/plugin/prometheus"
	"sync"
	"testing"
)

type Order struct {
	Id     int64 `gorm:"primarykey"`
	UserID int64 `gorm:"column:xb_uid"`
	Stake  float64
}

func (o *Order) TableName() string {
	return "win_betslips"
}

type Log struct {
	Id     int64 `gorm:"primarykey"`
	UserID int64 `gorm:"column:uid"`
	Stake  float64
}

func (o *Log) TableName() string {
	return "win_coin_log"
}

type User struct {
	Id       int64   `gorm:"primarykey"`
	Username string  `gorm:"column:username"`
	LevelId  float64 `gorm:"column:level_id"`
}

func (o *User) TableName() string {
	return "win_user"
}

func TestZorm(t *testing.T) {
	dsn := "admin:Mr61JkEz5KDEMRHAzsu4@tcp(filbet-dev-benchmark.cluster-cxloozl5ivvh.ap-east-1.rds.amazonaws.com:3306)/filbet?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	//redisClient := redis.NewClient(&redis.Options{
	//	Addr:     "16.162.112.251:6379",
	//	Password: "fasdaSASDAAao012312312ob1aredlk90",
	//	DB:       2,
	//})
	//
	//cache, err := cache.NewGorm2Cache(&config.CacheConfig{
	//	CacheLevel:           config.CacheLevelAll,
	//	CacheStorage:         storage.NewRedis(&storage.RedisStoreConfig{Client: redisClient}),
	//	InvalidateWhenUpdate: true,  // when you create/update/delete objects, invalidate cache
	//	CacheTTL:             50000, // 5000 ms
	//	CacheMaxItemCnt:      50,    // if length of objects retrieved one single time
	//	// exceeds this number, then don't cache
	//	Tables: []string{"test"},
	//})
	//
	//if err != nil {
	//	panic(err)
	//}

	//err = db.Use(cache)
	//if err != nil {
	//	panic(err)
	//}

	dsn2 := "admin:Mr61JkEz5KDEMRHAzsu4@tcp(filbet-dev-benchmark.cluster-cxloozl5ivvh.ap-east-1.rds.amazonaws.com:3306)/filbet_sharding?charset=utf8mb4&parseTime=True&loc=Local"
	db2, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn2}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	middleware := sharding.Register(sharding.Config{
		ShardingKey:    "xb_uid",
		NumberOfShards: 1024,
		Name:           "gorm:sharding",
	}, "win_betslips", "win_betslips_details")

	err = db.Use(middleware)
	if err != nil {
		panic(err)
	}

	middleware2 := sharding.Register(sharding.Config{
		ShardingKey:    "uid",
		NumberOfShards: 1024,
		Name:           "gorm:sharding2",
	}, "win_coin_log")

	//m2 := &Middleware2{
	//	middleware2,
	//}
	//
	err = db.Use(middleware2)
	if err != nil {
		panic(err)
	}

	dsnRead := "admin:Mr61JkEz5KDEMRHAzsu4@tcp(filbet-dev-benchmark.cluster-ro-cxloozl5ivvh.ap-east-1.rds.amazonaws.com:3306)/filbet?charset=utf8mb4&parseTime=True&loc=Local"
	dbRead, err := gorm.Open(mysql.New(mysql.Config{DSN: dsnRead}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	dsnRead2 := "admin:Mr61JkEz5KDEMRHAzsu4@tcp(filbet-dev-benchmark.cluster-ro-cxloozl5ivvh.ap-east-1.rds.amazonaws.com:3306)/filbet_sharding?charset=utf8mb4&parseTime=True&loc=Local"
	dbRead2, err := gorm.Open(mysql.New(mysql.Config{DSN: dsnRead2}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	err = db.Use(dbresolver.Register(
		dbresolver.Config{
			Replicas: []gorm.Dialector{
				dbRead.Dialector, dbRead.Dialector, dbRead.Dialector, dbRead.Dialector, dbRead.Dialector, dbRead.Dialector,
			},
			Policy: dbresolver.RandomPolicy{},
		},
	).Register(
		dbresolver.Config{
			Sources: []gorm.Dialector{
				db2.Dialector,
			},
			Replicas: []gorm.Dialector{
				dbRead2.Dialector, dbRead2.Dialector, dbRead2.Dialector, dbRead2.Dialector, dbRead2.Dialector, dbRead2.Dialector,
			},
			Policy: dbresolver.RandomPolicy{},
		}, "win_betslips", "win_coin_log",
	))
	if err != nil {
		panic(err)
	}

	db.Use(prometheus.New(prometheus.Config{
		DBName:          "db1", // 使用 `DBName` 作为指标 label
		RefreshInterval: 5,     // 指标刷新频率（默认为 15 秒）
		//	PushAddr:        "prometheus pusher address", // 如果配置了 `PushAddr`，则推送指标
		StartServer:    true, // 启用一个 http 服务来暴露指标
		HTTPServerPort: 8084, // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		}, // 用户自定义指标
	}))

	//middleware := sharding.Register(sharding.Config{
	//	ShardingKey:    "xb_uid",
	//	NumberOfShards: 1024,
	//}, "win_betslips")
	//err = db.Use(middleware)
	//if err != nil {
	//	panic(err)
	//}

	db = db.Debug()

	//cache := go_gorm_cache.NewDBCache(store.NewRedis(redisClient), go_gorm_cache.Conf{})
	//err = db.Use(cache)
	//if err != nil {
	//	panic("db.Use connect database")
	//}
	//db = db.Debug()
	//cache.AttachToDB(db)
	// More options

	// this will redirect query to orders_02
	//var orders []Order
	for {
		wg := sync.WaitGroup{}

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var order Order
				//var count int64
				tx := db.Model(&Order{}).Where("xb_uid", int64(23551)).First(&order)
				if tx.Error != nil {
					panic(tx.Error)
				}
				t.Log(tx.Dialector.Name())
				fmt.Printf("%#v\n", order)
				var log = &Log{}

				tx = db.Model(&Log{}).Where("uid", int64(23551)).First(&log)
				if tx.Error != nil {
					panic(tx.Error)
				}
				fmt.Printf("%#v\n", log)
			}()
		}
		wg.Wait()

		//var user User
		//err = db.Model(&User{}).Where("id", int64(10003)).First(&user).Error
		//if err != nil {
		//	panic(err)
		//}
		//t.Log(user)
		//tx := db.Model(&Order{}).Where("xb_uid=? and id =?", int64(23551), int64(4368990922220210)).
		//	Update("stake", rand.Int31n(33333))
		//if tx.Error != nil {
		//	panic(tx.Error)
		//}
		//time.Sleep(time.Second * 3)
	}
}

func getDBTest() *gorm.DB {

	dsn := "admin:Mr61JkEz5KDEMRHAzsu4@tcp(filbet-dev-benchmark.cluster-cxloozl5ivvh.ap-east-1.rds.amazonaws.com:3306)/filbet?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	//redisClient := redis.NewClient(&redis.Options{
	//	Addr:     "16.162.112.251:6379",
	//	Password: "fasdaSASDAAao012312312ob1aredlk90",
	//	DB:       2,
	//})
	//
	//cache, err := cache.NewGorm2Cache(&config.CacheConfig{
	//	CacheLevel:           config.CacheLevelAll,
	//	CacheStorage:         storage.NewRedis(&storage.RedisStoreConfig{Client: redisClient}),
	//	InvalidateWhenUpdate: true,  // when you create/update/delete objects, invalidate cache
	//	CacheTTL:             50000, // 5000 ms
	//	CacheMaxItemCnt:      50,    // if length of objects retrieved one single time
	//	// exceeds this number, then don't cache
	//	Tables: []string{"test"},
	//})
	//
	//if err != nil {
	//	panic(err)
	//}

	//err = db.Use(cache)
	//if err != nil {
	//	panic(err)
	//}

	dsn2 := "admin:Mr61JkEz5KDEMRHAzsu4@tcp(filbet-dev-benchmark.cluster-cxloozl5ivvh.ap-east-1.rds.amazonaws.com:3306)/filbet_sharding?charset=utf8mb4&parseTime=True&loc=Local"
	db2, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn2}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	middleware := sharding.Register(sharding.Config{
		ShardingKey:    "xb_uid",
		NumberOfShards: 1024,
		Name:           "gorm:sharding",
	}, "win_betslips", "win_betslips_details")

	err = db.Use(middleware)
	if err != nil {
		panic(err)
	}

	middleware2 := sharding.Register(sharding.Config{
		ShardingKey:    "uid",
		NumberOfShards: 1024,
		Name:           "gorm:sharding2",
	}, "win_coin_log")

	//m2 := &Middleware2{
	//	middleware2,
	//}
	//
	err = db.Use(middleware2)
	if err != nil {
		panic(err)
	}

	dsnRead := "admin:Mr61JkEz5KDEMRHAzsu4@tcp(filbet-dev-benchmark.cluster-ro-cxloozl5ivvh.ap-east-1.rds.amazonaws.com:3306)/filbet?charset=utf8mb4&parseTime=True&loc=Local"
	dbRead, err := gorm.Open(mysql.New(mysql.Config{DSN: dsnRead}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	dsnRead2 := "admin:Mr61JkEz5KDEMRHAzsu4@tcp(filbet-dev-benchmark.cluster-ro-cxloozl5ivvh.ap-east-1.rds.amazonaws.com:3306)/filbet_sharding?charset=utf8mb4&parseTime=True&loc=Local"
	dbRead2, err := gorm.Open(mysql.New(mysql.Config{DSN: dsnRead2}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	err = db.Use(dbresolver.Register(
		dbresolver.Config{
			Replicas: []gorm.Dialector{
				dbRead.Dialector, dbRead.Dialector, dbRead.Dialector, dbRead.Dialector, dbRead.Dialector, dbRead.Dialector,
			},
			Policy: dbresolver.RandomPolicy{},
		},
	).Register(
		dbresolver.Config{
			Sources: []gorm.Dialector{
				db2.Dialector,
			},
			Replicas: []gorm.Dialector{
				dbRead2.Dialector, dbRead2.Dialector, dbRead2.Dialector, dbRead2.Dialector, dbRead2.Dialector, dbRead2.Dialector,
			},
			Policy: dbresolver.RandomPolicy{},
		}, "win_betslips", "win_coin_log",
	))
	if err != nil {
		panic(err)
	}

	db.Use(prometheus.New(prometheus.Config{
		DBName:          "db1", // 使用 `DBName` 作为指标 label
		RefreshInterval: 5,     // 指标刷新频率（默认为 15 秒）
		//	PushAddr:        "prometheus pusher address", // 如果配置了 `PushAddr`，则推送指标
		StartServer:    true, // 启用一个 http 服务来暴露指标
		HTTPServerPort: 8084, // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		}, // 用户自定义指标
	}))

	//middleware := sharding.Register(sharding.Config{
	//	ShardingKey:    "xb_uid",
	//	NumberOfShards: 1024,
	//}, "win_betslips")
	//err = db.Use(middleware)
	//if err != nil {
	//	panic(err)
	//}

	db = db.Debug()

	//cache := go_gorm_cache.NewDBCache(store.NewRedis(redisClient), go_gorm_cache.Conf{})
	//err = db.Use(cache)
	//if err != nil {
	//	panic("db.Use connect database")
	//}
	//db = db.Debug()
	//cache.AttachToDB(db)
	// More options

	// this will redirect query to orders_02
	//var orders []Order
	return db
}

func TestNewClient(t *testing.T) {
	c := &DBConfig{
		MainSource: Source{
			SchemaConfig: SchemaConfig{
				Host:        "filbet-dev-benchmark.cluster-cxloozl5ivvh.ap-east-1.rds.amazonaws.com",
				Port:        3306,
				Password:    "Mr61JkEz5KDEMRHAzsu4",
				Username:    "admin",
				DBName:      "filbet",
				MaxOpenConn: 200,
			},
			Replicas: []SchemaConfig{
				{
					Host:        "filbet-dev-benchmark.cluster-ro-cxloozl5ivvh.ap-east-1.rds.amazonaws.com",
					Port:        3306,
					Password:    "Mr61JkEz5KDEMRHAzsu4",
					Username:    "admin",
					DBName:      "filbet",
					MaxOpenConn: 200,
				},
			},
		},
		OtherSource: []Source{
			{
				SchemaConfig: SchemaConfig{
					Host:        "filbet-dev-benchmark.cluster-cxloozl5ivvh.ap-east-1.rds.amazonaws.com",
					Port:        3306,
					Password:    "Mr61JkEz5KDEMRHAzsu4",
					Username:    "admin",
					DBName:      "filbet_sharding",
					MaxOpenConn: 200,
				},
				Replicas: []SchemaConfig{
					{
						Host:        "filbet-dev-benchmark.cluster-ro-cxloozl5ivvh.ap-east-1.rds.amazonaws.com",
						Port:        3306,
						Password:    "Mr61JkEz5KDEMRHAzsu4",
						Username:    "admin",
						DBName:      "filbet_sharding",
						MaxOpenConn: 200,
					},
				},
				Tables: []string{
					"win_betslips", "win_coin_log",
				},
			},
		},
		Prometheus: DBPrometheus{
			Name:            "common",
			HTTPServerPort:  8084,
			RefreshInterval: 5,
		},
		Sharding: []DBSharding{
			{
				Name:           "gorm:sharding",
				ShardingKey:    "xb_uid,uid",
				NumberOfShards: 1024,
				Tables: []string{
					"win_betslips", "win_betslips_details", "win_coin_log",
				},
			},
		},
	}
	db := NewClient(c)

	//db = getDBTest()
	dao.SetDefault(db)
	con, err := dao.WinConfig.Where(dao.WinConfig.ID.Eq(1)).First()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(con)
	con, err = dao.WinConfig.Where(dao.WinConfig.ID.Eq(1)).First()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(con)
	con, err = dao.WinConfig.Where(dao.WinConfig.ID.Eq(1)).First()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(con)
}
