package mapper

import (
	"testing"

	"github.com/shopspring/decimal"
)

func Test_Mapper(t *testing.T) {
	a1 := A1{
		Name:    "stu1",
		Address: "333",
		Age:     14,
		Coin:    1.111,
		Test:    "dsfasdfsdf",
	}
	d := Map[D1](&a1)
	t.Log(d)
	t.Log(d.Coin.InexactFloat64())
}

func TestMapper2(t *testing.T) {
	d1 := D1{
		UserName: "stu1",
		Address:  222,
		Age:      14,
		Coin:     decimal.NewFromFloat(1.11),
		Test2:    "dsfasdfsdf",
	}
	d := Map[A1](&d1)
	t.Log(d)
	t.Log(d.Test)
}

func TestDecimal(t *testing.T) {
	t.Log(decimal.NewFromFloat(180))
}

type A1 struct {
	Name    string  `json:"name"`
	Address string  `json:"add"`
	Age     uint8   `json:"age"`
	Coin    float32 ` json:"coin"`
	Test    string  `json:"test"`
}

type D1 struct {
	UserName string          `gorm:"column:name"`
	Address  int             `gorm:"column:add"`
	Age      int64           `gorm:"column:age"`
	Coin     decimal.Decimal `gorm:"column:coin"`
	Test2    string          `gorm:"column:test"`
}
