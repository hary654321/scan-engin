package client

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
	"zrWorker/core/slog"
	"zrWorker/global"
	"zrWorker/pkg/utils"

	"golang.org/x/net/proxy"
)

type ResponseJson struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}
type CountJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data int    `json:"data"`
}

type HostInfoJson struct {
	Cpuinfos  map[string]interface{} `json:"cpuinfos"`
	Hostinfos map[string]interface{} `json:"hostinfos"`
	Meminfos  map[string]interface{} `json:"meminfos"`
	Netinfos  []interface{}          `json:"netinfos"`
	Netspeed  []interface{}          `json:"netspeed"`
	Parts     []interface{}          `json:"parts"`
}

var ProxyMap = []string{"104.129.182.86:8880", "174.137.55.184:8880", "104.243.23.33:8880", "104.129.180.98:8880"}

func getUrl(ip string, port int, path string) string {
	url := "https://" + ip + ":" + utils.GetInterfaceToString(port) + path

	//log.Info("url :" + url)

	return url
}

func GetCli(timeout time.Duration) (*http.Client, string) {

	// setup a http client
	httpTransport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient := &http.Client{Transport: httpTransport, Timeout: timeout}

	// set our socks5 as the dialer
	// create a socks5 dialer
	addr := GetAddr()
	// slog.Println(slog.WARN, "addr", addr)
	if global.AppSetting.Proxy && addr != "0" {
		dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
		if err != nil {
			fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
			os.Exit(1)
		}
		httpTransport.Dial = dialer.Dial
	}

	return httpClient, addr
}

func GetAddr() string {
	num := utils.RanNum(len(ProxyMap))
	addr := ProxyMap[num]

	return addr
}

func GetConn(protocol, address string, timeout time.Duration) (net.Conn, error) {
	addr := GetAddr()
	if global.AppSetting.Proxy && addr != "0" {
		dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)

		if err != nil {
			slog.Println(slog.DEBUG, "can't connect to the proxy:", err, "addr", addr)
		}
		return dialer.Dial(protocol, address)

	}

	return net.DialTimeout(protocol, address, timeout)
}
