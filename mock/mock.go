package mock

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/gzorm/commons/utils"
	"github.com/gzorm/commons/zorm/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"testing"

	zr "github.com/gzorm/commons/core/stores/redis"
)

func GetPlatLogicService[T any, Ctx any](t *testing.T, gc any, constructor any) (*T, sqlmock.Sqlmock, *miniredis.Miniredis, utils.Snowflake) {
	svc := new(Ctx)
	s := getRedis(t)
	rc := zr.MustNewRedis(zr.RedisConf{
		Host: s.Addr(),
		Type: "node",
	})
	gdb, mock := getMockDB()
	dao.SetDefault(gdb)
	svcRV := reflect.ValueOf(svc)
	svcRV.Elem().FieldByName("Redis").Set(reflect.ValueOf(rc))
	snow := utils.NewSnowflake(rc)
	svcRV.Elem().FieldByName("Snowflake").Set(reflect.ValueOf(snow))
	svcRV.Elem().FieldByName("GameCallbackRpc").Set(reflect.ValueOf(gc))
	consRV := reflect.ValueOf(constructor)
	retList := consRV.Call([]reflect.Value{
		reflect.ValueOf(context.Background()),
		reflect.ValueOf(svc),
	})
	return retList[0].Interface().(*T), mock, s, snow
}

func GetGameCallBackLogicService[T any, Ctx any](t *testing.T, constructor any) (*T, sqlmock.Sqlmock, *miniredis.Miniredis, utils.Snowflake) {
	svc := new(Ctx)
	s := getRedis(t)
	rc := zr.MustNewRedis(zr.RedisConf{
		Host: s.Addr(),
		Type: "node",
	})
	gdb, mock := getMockDB()
	dao.SetDefault(gdb)
	svcRV := reflect.ValueOf(svc)
	svcRV.Elem().FieldByName("Redis").Set(reflect.ValueOf(rc))
	snow := utils.NewSnowflake(rc)
	svcRV.Elem().FieldByName("Snowflake").Set(reflect.ValueOf(snow))
	consRV := reflect.ValueOf(constructor)
	retList := consRV.Call([]reflect.Value{
		reflect.ValueOf(context.Background()),
		reflect.ValueOf(svc),
	})
	return retList[0].Interface().(*T), mock, s, snow
}

func getRedis(t *testing.T) *miniredis.Miniredis {
	s := miniredis.RunT(t)
	return s
}

func getMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	var mock sqlmock.Sqlmock
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if conn == nil {
		panic(err)
	}

	return conn, mock
}
