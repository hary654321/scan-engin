package cdn

import (
	"testing"
)

func TestCdn(t *testing.T) {
	ip, err := Resolution("www.baidu.com")

	println(ip)
	println(err)

	FindWithDomain("www.baidu.com")
}

//获取ip地址

func TestWryt(*testing.T) {
	Init("qqwry.dat")

	res, _ := Find("171.113.110.235")
	println(res)
}
