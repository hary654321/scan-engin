package scanner

import (
	"github.com/hary654321/gonmap"
	"net"
	"time"
	"zrWorker/lib/udp"
	"zrWorker/pkg/utils"
)

type foo1 struct {
	addr net.IP
	num  int
}

type PortClient struct {
	*client

	HandlerClosed     func(addr net.IP, port int)
	OutputUdpResponse func(addr net.IP, port int, response *udp.Result)
	HandlerOpen       func(addr net.IP, port int)
	HandlerNotMatched func(addr net.IP, port int, response string)
	HandlerMatched    func(addr net.IP, port int, response *gonmap.Response)
	HandlerError      func(addr net.IP, port int, err error)
}

func NewPortScanner(config *Config) *PortClient {
	var client = &PortClient{
		client:            newConfig(config, config.Threads),
		HandlerClosed:     func(addr net.IP, port int) {},
		HandlerOpen:       func(addr net.IP, port int) {},
		HandlerNotMatched: func(addr net.IP, port int, response string) {},
		HandlerMatched:    func(addr net.IP, port int, response *gonmap.Response) {},
		HandlerError:      func(addr net.IP, port int, err error) {},
	}
	client.pool.Interval = config.Interval
	client.pool.Function = func(in interface{}) {
		//println(1)
		nmap := gonmap.New()
		nmap.SetTimeout(time.Second * 10)
		//if config.DeepInspection == true {
		//	nmap.OpenDeepIdentify()
		//}
		value := in.(foo1)

		if utils.In_array(value.num, udp.UdpPort) {
			//slog.Println(slog.WARN, "udp 检测", addr.String(), udpPort)
			res, err := udp.UdpInfo(value.addr.String(), value.num)
			if err == nil {
				client.OutputUdpResponse(value.addr, value.num, res)
			}
		} else {
			//具体进行端口扫描
			status, response := nmap.ScanTimeout(value.addr.String(), value.num, time.Second*2)
			//slog.Println(slog.DEBUG, "端口状态：", value.addr.String(), ":", value.num, status.String(), response)
			switch status {
			case gonmap.Closed:
				client.HandlerClosed(value.addr, value.num)
			case gonmap.Open:
				client.HandlerOpen(value.addr, value.num)
			case gonmap.NotMatched:
				client.HandlerNotMatched(value.addr, value.num, response.Raw)
			case gonmap.Matched:
				client.HandlerMatched(value.addr, value.num, response)
			}
		}
	}
	return client
}

func (c *PortClient) Push(ip net.IP, num int) {
	c.pool.Push(foo1{ip, num})
}
