package test

import (
	"github.com/gzorm/commons/zorm-cache/cache"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

func testPluck(cache cache.Cache, db *gorm.DB) {
	var value9 []string
	result := db.Model(&TestModel{}).Pluck("value9", &value9)
	So(result.Error, ShouldBeNil)
	So(len(value9), ShouldEqual, testSize)
	So(value9[0], ShouldEqual, "1")
}
