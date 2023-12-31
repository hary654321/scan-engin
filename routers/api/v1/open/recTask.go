package open

import (
	"net/http"
	"strings"
	"zrWorker/core/client"
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
	mul := c.PostForm("mul")

	slog.Println(slog.DEBUG, "runTaskID:", runTaskID, "taskId", taskId, "mul", mul)

	// slog.Println(slog.DEBUG, "mul:", mul, runTaskID, taskId)
	engine, newT := run.NewEngine(runTaskID)
	if !newT {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  "重复任务",
		})
		return
	}
	var tarArr []string

	if mul != "" {
		tarArr = strings.Split(mul, ",")
	} else {
		c.SaveUploadedFile(f, "./"+f.Filename)

		fileinfo := utils.Read("./" + f.Filename)
		tarArr = strings.Split(fileinfo, "\n")
	}

	if len(tarArr) > 0 {

		engine.Total = len(tarArr)

		addrCount := len(client.ProxyMap)

		bc := engine.Total / addrCount
		for i := 0; i < addrCount; i++ {
			s := bc * i
			e := bc * (i + 1)
			go client.RunTask(taskId, runTaskID, tarArr[s:e], i)
		}

		// for _, v := range tarArr {
		// 	slog.Println(slog.DEBUG, v)

		// 	if global.AppSetting.Engin {
		// 		slog.Println(slog.DEBUG, v, "走vps")
		// 		go client.RunTask(taskId, runTaskID, v)
		// 		continue
		// 	}

		// 	go engine.PushTarget(v)
		// }
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
