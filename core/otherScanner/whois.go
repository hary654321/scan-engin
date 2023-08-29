package otherScanner

import (
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"zrWorker/core/slog"
)

func GetWhoisInfo(t string) (info whoisparser.WhoisInfo, err error) {
	result, err := whois.Whois(t)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	info, err = whoisparser.Parse(result)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	slog.Println(slog.DEBUG, t, "有数据")
	return
}
