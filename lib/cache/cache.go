package cache

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"time"
	"zrWorker/core/slog"
)

var cli *bigcache.BigCache

func NewCacheClient(LifeWindow time.Duration) {

	cli, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(LifeWindow*time.Minute))

}

func Set(k string, v []byte) {
	err := cli.Set(k, v)

	if err != nil {
		slog.Println(slog.DEBUG, "Set err", err.Error())
	}
}

func Get(k string) []byte {
	res, err := cli.Get(k)
	if err != nil {
		//slog.Println(slog.DEBUG, "Get err:", err.Error())
	}

	//slog.Println(slog.INFO, "cache Get info:", k,string(res))

	return res
}

func SetTaskLog(k string, taskLog TaskLog) {
	res, err := json.Marshal(taskLog)
	if err != nil {
		slog.Println(slog.DEBUG, "SetTaskLog err", err.Error())
	}

	Set(k, res)
}

func GetTaskLog(k string) TaskLog {

	res := Get(k)
	var taskLog TaskLog

	err := json.Unmarshal(res, &taskLog)
	if err != nil {
		slog.Printf(slog.DEBUG, "错误%s", err.Error())
	}

	return taskLog

}
