package spy

import (
	"testing"
	"zrWorker/core/slog"
	"zrWorker/lib/uri"
)

func TestIp(t *testing.T) {
	var gatewayArr []string
	slog.Println(slog.INFO, "现在开始枚举常见网段192.168.0.0")
	gatewayArr = uri.GetGatewayList("192.168.0.1", "b")
	slog.Println(slog.INFO, "ip数组", gatewayArr)
}
