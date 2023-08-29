package routers

import (
	"github.com/gin-gonic/gin"
	"zrWorker/middleware/jwt"
	"zrWorker/routers/api/v1/open"
)

func loadApiV1Open(r *gin.Engine) {

	openApi := r.Group("/api/v1")
	//开放接口
	openApi.Use(jwt.Open())
	{
		//心跳
		openApi.GET("/heartbeat", open.HeartBeat)
		//获取任务结果
		openApi.GET("/getTaskRes", open.GetTaskRes)
		//接收任务
		openApi.POST("/recTask", open.RecTask)

		//本机信息
		openApi.GET("/info", open.InfoGet)

		//zip压缩包结果
		openApi.GET("/zip", open.GetZip)

		//图片
		openApi.GET("/image", open.Image)
		//结果数据
		openApi.GET("/taskCount", open.TaskCount)

		openApi.GET("/serviceLog", open.ServiceLog)

		openApi.GET("/resLog", open.ResLog)

		openApi.GET("/ip", open.Ip)

		openApi.GET("/test", open.Test)

	}
}
