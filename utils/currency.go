package utils

import (
	"strings"

	"github.com/gzorm/commons/zorm/enum"
	"github.com/spf13/cast"
)

func WalletCurrencyCode(currency string) int64 {
	switch strings.ToUpper(strings.TrimSpace(currency)) {
	case "USDT", "2":
		return 2
	case "EGP", "1", "0", "":
		return int64(enum.WalletCurrencyPhp)
	default:
		if n := cast.ToInt64(currency); n > 0 {
			return n
		}
		return int64(enum.WalletCurrencyPhp)
	}
}
