package open

import (
	"net/http"
	"strings"
	"zrWorker/app"
	"zrWorker/core/slog"
	"zrWorker/core/spy"
	"zrWorker/global"
	"zrWorker/lib/cache"
	"zrWorker/pkg/e"
	"zrWorker/pkg/utils"
	"zrWorker/run"

	"github.com/gin-gonic/gin"
)

func RecTask(c *gin.Context) {

	t := c.PostForm("t")
	spyParam := c.PostForm("spy")
	mul := c.PostForm("mul")
	hydra := c.PostForm("hydra")
	addr := c.PostForm("addr")
	runTaskID := c.PostForm("runTaskId")
	taskId := c.PostForm("taskId")

	engine, newT := run.NewEngine(runTaskID)
	if !newT {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  "重复任务",
		})
		return
	}
	if t != "" {
		go engine.PushTarget(t)
	}
	if spyParam != "" {
		go spy.Start(engine, spyParam)
	}
	if mul != "" {
		strArrayNew := strings.Split(mul, ",")

		engine.Total = len(strArrayNew)

		for _, v := range strArrayNew {
			slog.Println(slog.DEBUG, v)
			go engine.PushTarget(v)
		}
	}
	if hydra == "1" {
		app.Setting.Hydra = true
	}
	if addr != "" {
		global.AppSetting.AdrrArr = append(global.AppSetting.AdrrArr, addr)
		go engine.ScanAddr(addr)
	}
	//任务信息记录下来
	startTime := utils.GetTime()

	logData := cache.TaskLog{
		TaskID:    taskId,
		RunTaskID: runTaskID,
		StartTime: startTime,
		Progress:  0,
	}

	cache.SetTaskLog(runTaskID, logData)

	data := make(map[string]interface{})
	code := e.SUCCESS
	data["taskId"] = taskId
	data["runTaskId"] = runTaskID
	data["startTime"] = startTime

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
