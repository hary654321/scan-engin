package open

import (
	"net/http"
	"path"
	"zrWorker/app"
	"zrWorker/core/slog"
	"zrWorker/global"
	"zrWorker/lib/cache"
	"zrWorker/pkg/e"
	"zrWorker/pkg/utils"
	"zrWorker/run"

	"github.com/gin-gonic/gin"
)

func GetZip(c *gin.Context) {

	//任务信息记录下来
	//startTime := utils.GetTime()
	//runTaskID := c.PostForm("runTaskId")
	//taskId := c.PostForm("taskId")

	slog.Printf(slog.DEBUG, "GetZip")
	day := c.Query("day")
	if day == "" {
		day = utils.GetDate()
	}

	target := global.ServerSetting.LogPath + "/" + day + ".zip"

	//if _, err := os.Stat(target); err == nil || os.IsExist(err) {
	//	slog.Printf(slog.DEBUG, "已有文件"+target)
	//	slog.Printf(slog.DEBUG, "已有文件"+"*"+day+".json")
	//} else {
	utils.ZipFile(global.ServerSetting.LogPath, target, "*"+day+".json")
	//}

	//获取文件的名称
	fileName := path.Base(target)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(target)

	return
}

func GetTaskRes(c *gin.Context) {
	runTaskId := c.Query("runTaskId")
	taskInfo := cache.GetTaskLog(runTaskId)

	TaskEngine := run.EngineArr[runTaskId]

	//被删除了  说明执行完成了
	if TaskEngine == nil {
		taskInfo.Progress = 100
		taskInfo.Res = utils.Read(utils.GetLogPath(runTaskId, "ipInfo"))
		delete(app.Setting.TaskPercent, runTaskId)
	} else {
		taskInfo.Progress = TaskEngine.GetPercent()
	}
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": taskInfo,
	})
}

func Image(c *gin.Context) {

	url := c.Query("url")

	filePath := utils.GetScreenPath() + utils.Md5(url) + ".png"

	slog.Printf(slog.DEBUG, filePath)
	//获取文件的名称
	fileName := path.Base(filePath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)

	return
}

func TaskCount(c *gin.Context) {

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
