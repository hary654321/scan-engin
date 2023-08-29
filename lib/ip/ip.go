package ip

import (
	"strings"
	"zrWorker/core/slog"
	"zrWorker/lib/ip/xdb"
	"zrWorker/pkg/utils"
)

type Region struct {
	Country  string
	Province string
	City     string
	Isp      string
}

var searcher *xdb.Searcher

func init() {
	var dbFile = "/zrtx/config/cyberspace/ip2region.xdb"
	searcher, _ = xdb.NewSearcher(dbFile)
}
func SearchIpAddr(ipStr string) *Region {
	var region = new(Region)
	ip, err := xdb.CheckIP(ipStr)
	if err != nil {
		slog.Println(slog.DEBUG, ipStr)

	}

	regionstr, ioCount, err := searcher.Search(ip)
	if err != nil {
		slog.Println(slog.DEBUG, err.Error(), ioCount)
	} else {
		slog.Println(slog.DEBUG, ipStr, "-------", regionstr, ioCount)
	}

	strArrayNew := strings.Split(regionstr, "|")
	length := len(strArrayNew)
	region.Country = strArrayNew[0]
	if length > 2 {
		region.Province = strArrayNew[2]
	}
	if length > 3 {
		region.City = strArrayNew[3]
	}
	if length > 4 {
		region.Isp = strArrayNew[4]
	}

	//slog.Println(slog.DEBUG, region)
	return region

}

// 获取一个城市/省/国家 的ip
func GetAddrIp(addr string) (res []string) {
	userDict, _ := utils.ReadLineData("/zrtx/config/cyberspace/ip.merge.txt")

	for _, v := range userDict {

		if utils.GetStrACount(addr, v) > 0 {
			arr := strings.Split(v, "|")

			//slog.Println(slog.DEBUG,arr)

			startIp := xdb.GetIpUint32(arr[0])
			endIp := xdb.GetIpUint32(arr[1])

			for i := startIp; i < endIp; i++ {
				//println(xdb.Long2IP(i))
				res = append(res, xdb.Long2IP(i))
			}
		}
	}

	//println(len(res))
	//utils.WriteJson("sz.json",res)
	return
}
