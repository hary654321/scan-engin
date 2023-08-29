package ip2domian

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"zrWorker/core/slog"
)

var BaseUrl = "http://api.webscan.cc/?action=query&ip="

type ResponseJsonRt []map[string]string

func GETIpDomains(ip string) ResponseJsonRt {

	httpClient := http.Client{
		Timeout: 20 * time.Second,
	}

	req, _ := http.NewRequest(http.MethodGet, BaseUrl+ip, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	//req.Header.Add("Content-Length", "26")
	//req.Header.Add("Host", "tools.ipip.net")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	//req.Header.Add("Origin", "https://tools.ipip.net")
	//req.Header.Add("Referer", "https://tools.ipip.net/as.php")
	//req.Header.Add("X-Requested-With", "XMLHttpRequest")
	//req.Header.Add("sec-ch-ua", "Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108")
	//req.Header.Add("sec-ch-ua-platform", "Windows")
	//req.Header.Add("Cookie", " __root_domain_v=.ipip.net; _qddaz=QD.810970811957720; lastSE=baidu; _ga=GA1.2.35085502.1672133495; _gid=GA1.2.1971748185.1672133495; https_waf_cookie=6797258e-51e5-4ac291087444bdc46389597bef154a261e17; LOVEAPP_SESSID=a7930364022e2f9c77c2e87ac7d4724d4c245d5f; Hm_lvt_6b4a9140aed51e46402f36e099e37baf=1670811959,1672133495,1672189861; _gat=1; Hm_lpvt_6b4a9140aed51e46402f36e099e37baf=1672193761; _qdda=3-1.1; _qddab=3-5qg2et.lc70zndh")

	var ResponseJsonRt ResponseJsonRt

	resp, err := httpClient.Do(req)
	if err != nil {
		slog.Println(slog.DEBUG, "%s", err.Error())
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	//[{"domain":"www.zorelworld.com","title":"中睿天下_威胁监测 攻击溯源"},{"domain":"zorelworld.com","title":"中睿天下_威胁监测 攻击溯源"}]
	slog.Println(slog.DEBUG, string(body))
	if err != nil {
		slog.Println(slog.DEBUG, "错误%s", err.Error())
		return nil
	}
	if err = json.Unmarshal(body, &ResponseJsonRt); err != nil {
		slog.Println(slog.DEBUG, "错误%s", err.Error())
		return nil
	}

	if len(ResponseJsonRt) < 1 {
		slog.Println(slog.DEBUG, ip+"GETIpDomains 空数据")
		return nil
	}

	slog.Println(slog.DEBUG, "ResponseJsonRt", ResponseJsonRt)

	return ResponseJsonRt
}
