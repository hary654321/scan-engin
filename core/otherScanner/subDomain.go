package otherScanner

// 使用  https://github.com/ZhuriLab/Starmap
import (
	"context"
	"encoding/json"
	"github.com/hary654321/Starmap/pkg/runner"
	"os/exec"
	"time"
	"zrWorker/core/slog"
	"zrWorker/global"
	"zrWorker/pkg/utils"
)

type SubDomainInfo struct {
	Host string `json:"host"`
	Ip   string `json:"ip"`
}

// ./SubDomain -d zorelworld.com -o subDomain.json
func SaveSubDomainCli(runTaskID, domain string) []SubDomainInfo {
	slog.Println(slog.DEBUG, "开始调用Starmap 查询子域名：", domain)
	start := time.Now()
	fileName := utils.GetLogPath(runTaskID, "subDomain")
	cmd := exec.Command(global.ServerSetting.ScanPath+"/SubDomain", "-d", domain, "-o", fileName, "-oJ")

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, string(out))
		slog.Println(slog.DEBUG, err)
	}
	slog.Println(slog.DEBUG, string(out))

	var OutArr []SubDomainInfo

	res := string(out)
	strArr := utils.Explode(res, "\n")

	for _, subDomain := range strArr {
		if utils.GetStrACount("{", subDomain) > 0 && utils.GetStrACount("}", subDomain) > 0 {
			var SubDomainInfo SubDomainInfo
			if err = json.Unmarshal([]byte(subDomain), &SubDomainInfo); err != nil {
				slog.Printf(slog.WARN, "错误%s", err.Error())
				continue
			}
			OutArr = append(OutArr, SubDomainInfo)
		}
	}
	elapsed := time.Since(start)
	slog.Println(slog.DEBUG, "该函数执行完成耗时：", elapsed)
	return OutArr
}

// 采用包引入的方式
func SaveSubDomain(runTaskID, root string) (domainArr []string) {
	if root == "" {
		return
	}
	options := runner.DefaultOptions()

	newRunner, err := runner.NewRunner(options)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	err, res, _ := newRunner.EnumerateSingleDomain(context.Background(), root, nil)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	for subDomain, HostEntry := range *res {
		domainArr = append(domainArr, subDomain)
		//utils.WriteJson(root+".json", HostEntry)
		utils.WriteJson(utils.GetLogPath(runTaskID, "subDomain"), HostEntry)
	}

	return
}
