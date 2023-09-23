package client

import (
	"crypto/tls"
	"net/http"
	"time"
	"zrWorker/pkg/utils"
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

func getUrl(ip string, port int, path string) string {
	url := "https://" + ip + ":" + utils.GetInterfaceToString(port) + path

	//log.Info("url :" + url)

	return url
}

func GetCli(timeout time.Duration) *http.Client {

	// setup a http client
	httpTransport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient := &http.Client{Transport: httpTransport, Timeout: timeout}

	// set our socks5 as the dialer
	// create a socks5 dialer
	// num := utils.RanNum(len(define.ProxyMap))
	// addr := define.ProxyMap[num]
	// // slog.Println(slog.WARN, "addr", addr)
	// if addr != "0" && config.CoreConf.Env != "dev" {
	// 	dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
	// 		os.Exit(1)
	// 	}
	// 	httpTransport.Dial = dialer.Dial
	// }

	return httpClient
}
