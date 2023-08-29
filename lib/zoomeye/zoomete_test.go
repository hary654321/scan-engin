package zoomeye

import (
	"testing"
)

func TestZm(t *testing.T) {
	ZClinet := New()
	ZClinet.Search("www.zorelworld.com", "app,os,device")
}
