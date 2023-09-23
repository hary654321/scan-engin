package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
	"zrWorker/core/slog"
	"zrWorker/global"
	"zrWorker/pkg/utils"
)

var ProxyMap = []string{"104.129.182.86", "174.137.55.184", "104.243.23.33", "104.129.180.98"}

func GetAddr() string {
	num := utils.RanNum(len(ProxyMap))
	addr := ProxyMap[num]

	return addr
}

func RunTask(taskId, runTaskId string, ip string) (map[string]interface{}, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.WriteField("taskId", taskId)
	bodyWriter.WriteField("runTaskId", runTaskId)
	bodyWriter.WriteField("mul", ip)

	//设置文件入参

	// utils.Write(runTaskId+".csv", strings.Join(ipArr, ","))

	// fileWriter, err0 := bodyWriter.CreateFormFile("file", runTaskId+".csv")
	// // 把文件流写入到缓冲区里去
	// file, err := os.Open(runTaskId + ".csv")
	// if err == nil {
	// 	defer file.Close()
	// }

	// _, err1 := io.Copy(fileWriter, file)

	// if err1 != nil {
	// 	slog.Println(slog.DEBUG, "上传文件失败", zap.Error(err0))
	// }
	// bodyWriter.Close()
	req, _ := http.NewRequest(http.MethodPost, getUrl(GetAddr(), 18000, "/api/v1/recTask"), bodyBuf)

	contentType := bodyWriter.FormDataContentType()
	slog.Println(slog.DEBUG, "contentType", contentType)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", global.ServerSetting.BasicAuth)

	var responseJson ResponseJson
	var output map[string]interface{}

	cli := GetCli(20 * time.Second)
	resp, err := cli.Do(req)
	if err != nil {
		slog.Println(slog.DEBUG, "发送任务失败===", err, "target============")
		return output, err
	}
	slog.Println(slog.DEBUG, "成功代理", "目标ip", "\n")
	body, err := ioutil.ReadAll(resp.Body)
	//{"code":200,"data":{"runningTasks":"3412341234","time":"1672733119","version":"1.1.1"},"msg":""}
	if err != nil {
		slog.Println(slog.DEBUG, "读取任务response失败===", err)
		return output, err
	}
	if err = json.Unmarshal(body, &responseJson); err != nil {
		slog.Println(slog.DEBUG, "json读取失败==", err)
		return output, err
	}

	if responseJson.Code != 200 {
		slog.Println(slog.DEBUG, "===任务返回错误:"+responseJson.Msg)
		return output, err
	}
	return responseJson.Data, nil
}