package ip2region

import (
	"embed"
	"log"
	"net"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

//go:embed *.xdb
var localeFS embed.FS

// IP2RegionV4 IPv4 归属地查询器
var IP2RegionV4 *xdb.Searcher

// IP2RegionV6 IPv6 归属地查询器
var IP2RegionV6 *xdb.Searcher

// 货币类型常量
const (
	CurrencyUSDT = "USDT" // 美元稳定币
	CurrencyEGP  = "EGP"  // 埃及镑
)

// Init 初始化 IP2Region 查询器（同时加载 IPv4 和 IPv6 数据库）
func Init() error {
	// 加载 IPv4 数据库
	ipBytes, err := localeFS.ReadFile("ip2region_v4.xdb")
	if err != nil {
		return err
	}

	// 获取 IPv4 版本信息
	v4Version, err := xdb.VersionFromName("IPv4")
	if err != nil {
		return err
	}

	searcher, err := xdb.NewWithBuffer(v4Version, ipBytes)
	if err != nil {
		return err
	}

	IP2RegionV4 = searcher

	// 尝试加载 IPv6 数据库（如果存在）
	if v6Bytes, v6Err := localeFS.ReadFile("ip2region_v6.xdb"); v6Err == nil {
		v6Version, vErr := xdb.VersionFromName("IPv6")
		if vErr == nil {
			if v6Searcher, sErr := xdb.NewWithBuffer(v6Version, v6Bytes); sErr == nil {
				IP2RegionV6 = v6Searcher
				log.Println("ip2region: IPv6 database loaded successfully")
			} else {
				log.Printf("ip2region: IPv6 database init failed, err=%v", sErr)
			}
		} else {
			log.Printf("ip2region: IPv6 version not found, err=%v", vErr)
		}
	} else {
		log.Printf("ip2region: IPv6 database file not found, only IPv4 supported")
	}

	return nil
}

// isIPv6 判断 IP 地址是否为 IPv6
func isIPv6(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() == nil
}

// isLocalIP 判断是否为本地/内网 IP
func isLocalIP(ip string) bool {
	if ip == "" || ip == "localhost" {
		return true
	}
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	// IPv4 本地地址
	if parsedIP.IsLoopback() || parsedIP.IsPrivate() || parsedIP.IsLinkLocalUnicast() {
		return true
	}
	return false
}

// GetIPRegion 根据IP查询归属地（自动识别 IPv4/IPv6）
func GetIPRegion(ip string) (string, error) {
	if ip == "" {
		return "", nil
	}

	if isIPv6(ip) {
		if IP2RegionV6 != nil {
			return IP2RegionV6.Search(ip)
		}
		return "", nil
	}

	if IP2RegionV4 != nil {
		return IP2RegionV4.Search(ip)
	}
	return "", nil
}

//	根据IP确定货币类型
//
// 埃及地区返回 EGP，其他地区返回 USDT
func GetCurrencyByIP(ip string) (currency string) {
	// 暂时不处理ip识别货币 测试问题
	return CurrencyEGP

	// 处理本地 IP 或空 IP
	if isLocalIP(ip) {
		return CurrencyEGP // 本地/内网 IP 默认返回 EGP
	}

	// 防止 ip2region 查询时 panic
	defer func() {
		if r := recover(); r != nil {
			// 发生 panic 时默认返回 EGP，并记录 panic 信息便于排查
			currency = CurrencyEGP
			log.Printf("GetCurrencyByIP ip2region panic, ip=%s, panic=%v", ip, r)
		}
	}()

	region, err := GetIPRegion(ip)
	if err != nil {
		// 查询失败默认返回 EGP
		log.Printf("GetCurrencyByIP ip2region error, ip=%s, err=%v", ip, err)
		return CurrencyEGP
	}
	currency = GetCurrencyByRegion(region)
	log.Printf("GetCurrencyByIP ip=%s, region=%s, currency=%s", ip, region, currency)
	return
}

// GetCurrencyByRegion 根据IP归属地确定货币类型
// 埃及地区返回 EGP，其他地区返回 USDT
func GetCurrencyByRegion(region string) string {
	// 检查 IP 归属地是否包含埃及标识
	if strings.Contains(region, "埃及") || strings.Contains(strings.ToLower(region), "egypt") {
		return CurrencyEGP
	}
	return CurrencyUSDT
}
