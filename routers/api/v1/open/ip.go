package open

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zrWorker/core/slog"
	ip2 "zrWorker/lib/ip"
	"zrWorker/pkg/e"
)

func Ip(c *gin.Context) {

	ip := c.Query("ip")

	slog.Println(slog.DEBUG, "ip")
	res := ip2.SearchIpAddr(ip)

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "e.GetMsg(code)",
		"data": res,
	})
}
