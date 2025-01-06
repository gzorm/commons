package zorm

import (
	"github.com/gzorm/common/zorm/dao"
	"testing"
)

func TestNewRangeClient(t *testing.T) {
	c := &DBConfig{
		MainSource: Source{
			SchemaConfig: SchemaConfig{
				Host:        "16.162.51.186",
				Port:        3306,
				Password:    "scott123456",
				Username:    "root",
				DBName:      "filbet_dev_main",
				MaxOpenConn: 200,
			},
			Replicas: []SchemaConfig{
				{
					Host:        "16.162.51.186",
					Port:        3306,
					Password:    "scott123456",
					Username:    "root",
					DBName:      "filbet_dev_main",
					MaxOpenConn: 200,
				},
			},
		},
		OtherSource: []Source{
			{
				SchemaConfig: SchemaConfig{
					Host:        "16.162.51.186",
					Port:        3306,
					Password:    "scott123456",
					Username:    "root",
					DBName:      "dev_range_sharding",
					MaxOpenConn: 200,
				},
				Replicas: []SchemaConfig{
					{
						Host:        "16.162.51.186",
						Port:        3306,
						Password:    "scott123456",
						Username:    "root",
						DBName:      "dev_range_sharding",
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
				ShardingKey:    "dt_completed",
				NumberOfShards: 1024,
				Tables: []string{
					"win_betslips", "win_betslips_details", "win_coin_log",
				},
			},
		},
	}
	db := NewRangeShardingClient(c)

	//db = getDBTest()
	dao.SetDefault(db)
	con, err := dao.WinBetslip.Where(dao.WinBetslip.DtCompleted.Gt(1700451566)).Find()
	if err != nil {
		t.Fatal(err)
	}
	for _, betslip := range con {
		t.Log(*betslip)
	}

}
