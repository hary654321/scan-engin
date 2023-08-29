package icp

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"zrWorker/core/slog"
	"zrWorker/pkg/utils"
)

type ResponseJson struct {
	Code   int                    `json:"code"`
	Msg    string                 `json:"msg"`
	Params map[string]interface{} `json:"params"`
}

const (
	baseURL = "https://hlwicpfwc.miit.gov.cn/icpproject_query/api/"
)

type authParam struct {
	AuthKey   string `json:"authKey"`
	TimeStamp string `json:"timeStamp"`
}

type searchParam struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	UnitName string `json:"unitName"`
}

func getToken() ResponseJson {
	authkey := utils.Md5("testtest" + utils.GetTime())
	authParam := authParam{AuthKey: authkey, TimeStamp: utils.GetTime()}
	res := SendAuth("auth", "application/x-www-form-urlencoded;charset=UTF-8", "0", authParam)

	return res
}

/*
*

	{
	    "contentTypeName": "",
	    "domain": "zorelworld.com",
	    "domainId": 1.10000415475e+11,
	    "leaderName": "",
	    "limitAccess": "否",
	    "mainId": 1.10000161934e+11,
	    "mainLicence": "京ICP备14031481号",
	    "natureName": "企业",
	    "serviceId": 1.10000284809e+11,
	    "serviceLicence": "京ICP备14031481号-1",
	    "unitName": "北京中睿天下信息技术有限公司",
	    "updateRecordTime": "2021-12-22 13:09:26"
	}
*/
func GetIcpInfo(unitName string) (map[string]interface{}, bool) {
	authRes := getToken()
	token, _ := authRes.Params["bussiness"].(string)
	res := SendS("icpAbbreviateInfo/queryByCondition", "application/json;charset=UTF-8", token, searchParam{PageNum: 1, PageSize: 10, UnitName: unitName})
	outObj := make(map[string]interface{})
	if res.Code == 401 {
		return outObj, false
	}
	//fmt.Printf("map: %#v", res)
	ArrI := res.Params["list"]

	if ArrI == nil {
		return outObj, false
	}
	Arr := ArrI.([]interface{})
	if len(Arr) > 0 {
		ObjI := Arr[0]
		outObj = ObjI.(map[string]interface{})

		return outObj, true
	}
	return outObj, false
}

func SendAuth(path, Content, token string, data authParam) ResponseJson {

	formdata := make(url.Values)
	formdata["authKey"] = []string{data.AuthKey}
	formdata["timeStamp"] = []string{data.TimeStamp}

	req, _ := http.NewRequest(http.MethodPost, baseURL+path, strings.NewReader(formdata.Encode()))

	ip := "101." + getRandNum() + "." + getRandNum() + "." + getRandNum()

	//slog.Println(slog.DEBUG, ip)
	req.Header.Add("Content-Type", Content)
	req.Header.Add("Origin", "https://beian.miit.gov.cn/")
	req.Header.Add("Referer", "https://beian.miit.gov.cn/")
	req.Header.Add("token", token)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36")
	req.Header.Add("CLIENT-IP", ip)
	req.Header.Add("X-FORWARDED-FOR", ip)

	var responseJson ResponseJson
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Println(slog.DEBUG, err)
		return responseJson
	}
	body, err := ioutil.ReadAll(resp.Body)
	//slog.Println(slog.INFO, "icp获取token", string(body))
	if err != nil {
		slog.Println(slog.DEBUG, err)
		return responseJson
	}

	if err = json.Unmarshal(body, &responseJson); err != nil {
		slog.Println(slog.DEBUG, body, err)
		return responseJson
	}

	return responseJson
}

func SendS(path, Content, token string, data searchParam) ResponseJson {

	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, baseURL+path, bytes.NewBuffer(bytesData))

	ip := "101." + getRandNum() + "." + getRandNum() + "." + getRandNum()

	req.Header.Add("Content-Type", Content)
	req.Header.Add("Origin", "https://beian.miit.gov.cn/")
	req.Header.Add("Referer", "https://beian.miit.gov.cn/")
	req.Header.Add("token", token)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36")
	req.Header.Add("CLIENT-IP", ip)
	req.Header.Add("X-FORWARDED-FOR", ip)

	var responseJson ResponseJson
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Println(slog.DEBUG, err)
		return responseJson
	}
	body, err := ioutil.ReadAll(resp.Body)
	//slog.Println(slog.DEBUG, "获取结果返回：", data, string(body))
	if err != nil {
		slog.Println(slog.DEBUG, err)
		return responseJson
	}

	if err = json.Unmarshal(body, &responseJson); err != nil {
		slog.Println(slog.DEBUG, err)
		return responseJson
	}

	return responseJson
}

func getRandNum() string {
	result, _ := rand.Int(rand.Reader, big.NewInt(255))

	return result.String()
}
