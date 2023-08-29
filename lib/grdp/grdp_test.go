package grdp

import (
	"fmt"
	"testing"
	"zrWorker/lib/grdp/glog"
)

func testrdp(target, username, password string) {
	domain := ""
	//username := "administrator"
	//password := "zaq1@WSX"
	//target = "180.102.17.30:3389"
	var err error
	g := NewClient(target, glog.NONE)
	//SSL协议登录测试
	err = g.loginForSSL(domain, username, password)
	if err == nil {
		fmt.Println("Login Success")
		return
	}
	if err.Error() != "PROTOCOL_RDP" {
		fmt.Println("Login Error:", err)
		return
	}
	//RDP协议登录测试
	err = g.loginForRDP(domain, username, password)
	if err == nil {
		fmt.Println("Login Success")
		return
	} else {
		fmt.Println("Login Error:", err)
		return
	}
}

func TestName(t *testing.T) {
	targetArr := []string{
		//"50.57.49.172:3389",
		//"20.49.22.250:3389",
		"172.16.130.138:22",
	}
	for _, target := range targetArr {
		fmt.Println(target)
		testrdp(target, "root", "depth@2013")
	}
}
