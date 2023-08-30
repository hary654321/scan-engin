package open

import (
	"net/http"
	"strings"
	"zrWorker/core/slog"
	"zrWorker/lib/cache"
	"zrWorker/pkg/e"
	"zrWorker/pkg/utils"
	"zrWorker/run"

	"github.com/gin-gonic/gin"
)

func RecTask(c *gin.Context) {

	f, _ := c.FormFile("file")

	runTaskID := c.PostForm("runTaskId")
	taskId := c.PostForm("taskId")

	slog.Println(slog.DEBUG, "runTaskID:", runTaskID, taskId)

	c.SaveUploadedFile(f, "./"+f.Filename)

	mul := utils.Read("./" + f.Filename)
	// slog.Println(slog.DEBUG, "mul:", mul, runTaskID, taskId)
	engine, newT := run.NewEngine(runTaskID)
	if !newT {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  "重复任务",
		})
		return
	}

	if mul != "" {
		strArrayNew := strings.Split(mul, ",")

		engine.Total = len(strArrayNew)

		for _, v := range strArrayNew {
			slog.Println(slog.DEBUG, v)
			go engine.PushTarget(v)
		}
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
