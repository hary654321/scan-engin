package domain

import (
	"regexp"
	"zrWorker/core/slog"
	"zrWorker/pkg/utils"
)

var roots string
var pattern string

func init() {

	source, _ := utils.ReadLineData("/zrtx/config/cyberspace/public-suffix-list.txt")
	for _, line := range source {

		if utils.GetStrACount("//", line) < 1 && len(line) > 0 {
			//slog.Println(slog.INFO,line)
			a := line + "|"
			roots += a
		}
	}
	pattern = "(\\w*\\.?){1}\\.(" + roots + ")$"
}

func GetRoot(domain string) string {
	topPattern := "(\\w*\\.?){1}\\.(com.cn|net.cn|gov.cn|.nz|org.cn|com|net|org|gov|cc|biz|info|cn|co|uk|jp|kr|)$"
	_, err := regexp.MatchString(topPattern, domain)
	if err == nil {
		reg := regexp.MustCompile(topPattern)
		data := reg.Find([]byte(domain))
		res := string(data)

		//slog.Println(slog.DEBUG, "topPattern", res)

		return res
	}
	reg := regexp.MustCompile(pattern)
	data := reg.Find([]byte(domain))
	res := string(data)

	slog.Println(slog.DEBUG, res)

	return res
}
