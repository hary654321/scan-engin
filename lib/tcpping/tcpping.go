package tcpping

import (
	"fmt"
	"runtime"
	"strings"
	"time"
	"zrWorker/lib/client"
)

var CommonPorts = []int{
	80, 443, 7547, 22, 5060, 8080, 8443, 161, 2083, 2096, 8000,
	21, 2087, 8888, 53, 8089, 2082, 2095, 30005, 2086, 554,
}

func Ping(ip string, port int, timeout time.Duration) error {
	return connect("tcp", ip, port, timeout)
}

func PingPorts(ip string, timeout time.Duration) (err error) {
	for _, port := range CommonPorts {
		if err = Ping(ip, port, timeout); err == nil {
			return nil
		}
	}
	return err
}

func connect(protocol string, ip string, port int, duration time.Duration) error {
	address := fmt.Sprintf("%s:%d", ip, port)

	//在这里设置代理
	conn, err := client.GetConn(protocol, address, duration)
	if err != nil {
		return err
	}
	defer conn.Close()
	if (runtime.GOOS == "windows" && (port == 110 || port == 25)) == false {
		return nil
	}
	err = conn.SetDeadline(time.Now())
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte("\r\n"))
	if err != nil {
		if strings.Contains(err.Error(), "forcibly closed by the remote host") {
			return err
		}
		if strings.Contains(err.Error(), "timeout") {
			return err
		}
	}
	return nil
}
