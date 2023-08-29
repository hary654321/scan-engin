package glog

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	SetLevel(INFO)
	buf := new(bytes.Buffer)
	logger := log.New(buf, "hary", 0)
	SetLogger(logger)

	Debug("aaaaa", "bbbbbbbbbb")

	fmt.Print(buf)
}
