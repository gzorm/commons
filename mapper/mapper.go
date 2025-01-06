package mapper

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gzorm/commons/core/logx"
	"github.com/shopspring/decimal"
)

func Map[T any](from any) *T {
	t := new(T)
	tv := reflect.ValueOf(t)
	tt := reflect.TypeOf(t)
	rv := reflect.ValueOf(from)
	rt := reflect.TypeOf(from)

	for i := 0; i < rv.Elem().NumField(); i++ {
		if js, ok := rt.Elem().Field(i).Tag.Lookup("json"); ok {
			index, ok := findIndex(tt, js)
			if ok {

				// tv.Elem().Field(index).SetString(rv.Elem().Field(i).String())
				switch rv.Elem().Field(i).Kind() {
				case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
					if tv.Elem().Field(index).Kind() == reflect.String {
						tv.Elem().Field(index).SetString(fmt.Sprint(rv.Elem().Field(i).Interface()))
					} else if tv.Elem().Field(index).Kind() == reflect.Int64 {
						tv.Elem().Field(index).SetInt(int64(rv.Elem().Field(i).Int()))
					}
				case reflect.Uint, reflect.Uint16, reflect.Uint8, reflect.Uint32, reflect.Uint64:
					if tv.Elem().Field(index).Kind() == reflect.String {
						tv.Elem().Field(index).SetString(fmt.Sprint(rv.Elem().Field(i).Interface()))
					} else if tv.Elem().Field(index).Kind() == reflect.Int64 {
						tv.Elem().Field(index).SetInt(int64(rv.Elem().Field(i).Uint()))
					}

				case reflect.String:
					if tv.Elem().Field(index).Kind() == reflect.String {
						tv.Elem().Field(index).SetString(rv.Elem().Field(i).String())
					} else {
						v, err := strconv.Atoi(rv.Elem().Field(i).String())
						if err != nil {
							logx.Error(err)
						} else {
							tv.Elem().Field(index).Set(reflect.ValueOf(v))
						}

					}
				case reflect.Float32, reflect.Float64:
					d, err := decimal.NewFromString(fmt.Sprint(rv.Elem().Field(i).Interface()))
					if err != nil {
						logx.Error(err)
					}
					tv.Elem().Field(index).Set(reflect.ValueOf(d))
				}
			}
		}
		if js, ok := rt.Elem().Field(i).Tag.Lookup("gorm"); ok {
			find := false
			sp := strings.Split(js, ";")
			if len(sp) > 0 {
				sp1 := strings.Split(sp[0], ":")
				if len(sp1) > 1 {
					find = true
					js = sp1[1]
				}

			}

			if !find {
				continue
			}

			index, ok := findIndex(tt, js)
			if ok {

				// tv.Elem().Field(index).SetString(rv.Elem().Field(i).String())
				switch rv.Elem().Field(i).Kind() {
				case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
					if tv.Elem().Field(index).Kind() == reflect.String {
						tv.Elem().Field(index).SetString(fmt.Sprint(rv.Elem().Field(i).Interface()))
					} else if tv.Elem().Field(index).Kind() == reflect.Int64 || tv.Elem().Field(index).Kind() == reflect.Int32 || tv.Elem().Field(index).Kind() == reflect.Int8 {
						tv.Elem().Field(index).SetInt(rv.Elem().Field(i).Int())
					} else if tv.Elem().Field(index).Kind() == reflect.Uint64 || tv.Elem().Field(index).Kind() == reflect.Uint32 || tv.Elem().Field(index).Kind() == reflect.Uint8 {
						tv.Elem().Field(index).SetUint(uint64(rv.Elem().Field(i).Int()))
					}
				case reflect.Uint, reflect.Uint16, reflect.Uint8, reflect.Uint32, reflect.Uint64:
					if tv.Elem().Field(index).Kind() == reflect.String {
						tv.Elem().Field(index).SetString(fmt.Sprint(rv.Elem().Field(i).Interface()))
					} else if tv.Elem().Field(index).Kind() == reflect.Int64 {
						tv.Elem().Field(index).SetInt(int64(rv.Elem().Field(i).Uint()))
					}

				case reflect.String:
					if tv.Elem().Field(index).Kind() == reflect.String {
						tv.Elem().Field(index).SetString(rv.Elem().Field(i).String())
					} else {
						v, err := strconv.Atoi(rv.Elem().Field(i).String())
						if err != nil {
							logx.Error(err)
						} else {
							tv.Elem().Field(index).Set(reflect.ValueOf(v))
						}

					}
				case reflect.Float32, reflect.Float64:
					d, err := decimal.NewFromString(fmt.Sprint(rv.Elem().Field(i).Interface()))
					if err != nil {
						logx.Error(err)
					}
					tv.Elem().Field(index).Set(reflect.ValueOf(d))

				case reflect.Struct:
					if tv.Elem().Field(index).Kind() == reflect.Float64 || tv.Elem().Field(index).Kind() == reflect.Float32 {
						tv.Elem().Field(index).SetFloat(rv.Elem().Field(i).Interface().(decimal.Decimal).InexactFloat64())
					}
				}
			}
		}
	}
	return t
}

func findIndex(rt reflect.Type, key string) (int, bool) {
	for i := 0; i < rt.Elem().NumField(); i++ {
		if js, ok := rt.Elem().Field(i).Tag.Lookup("json"); ok {
			if strings.Contains(js, key) {
				return i, true
			}
		}
		if js, ok := rt.Elem().Field(i).Tag.Lookup("gorm"); ok {
			sp := strings.Split(js, ";")
			if len(sp) > 0 {
				sp1 := strings.Split(sp[0], ":")
				if len(sp1) > 1 {
					if sp1[1] == key {
						return i, true
					}
				}

			}
		}
	}
	return 0, false
}
