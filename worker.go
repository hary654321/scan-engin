package main

import (
	"embed"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
	"zrWorker/core/cdn"
	"zrWorker/core/hydra"
	"zrWorker/core/scanner"
	"zrWorker/core/slog"
	"zrWorker/global"
	"zrWorker/lib/appfinger"
	"zrWorker/lib/cache"
	"zrWorker/lib/gonmap"
	"zrWorker/lib/misc"
	"zrWorker/pkg/setting"
	"zrWorker/pkg/utils"
	"zrWorker/routers"

	"github.com/fvbock/endless"
)

func init() {
	err := setupSetting()
	if err != nil {
		slog.Printf(slog.WARN, "init.setupSetting err: %v", err)
	}
}
func main() {

	endless.DefaultReadTimeOut = global.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = global.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	endPoint := fmt.Sprintf(":%d", global.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		slog.Printf(slog.INFO, "Actual pid is %d", syscall.Getpid())
	}
	InitKscan()

	//http server
	err := server.ListenAndServeTLS("/zrtx/config/cyberspace/.cert.pem", "/zrtx/config/cyberspace/.key.pem")
	if err != nil {
		slog.Printf(slog.DEBUG, "Server err: %v", err)
	}
}

//go:embed configs/fingerprint.txt
var fingerprintEmbed embed.FS

const (
	qqwryPath       = "qqwry.dat"
	fingerprintPath = "configs/fingerprint.txt"
)

func InitKscan() {
	//HTTP指纹库初始化
	fs, _ := fingerprintEmbed.Open(fingerprintPath)
	if n, err := appfinger.InitDatabaseFS(fs); err != nil {
		slog.Println(slog.WARN, "指纹库加载失败，请检查【fingerprint.txt】文件", err)
	} else {
		slog.Printf(slog.INFO, "成功加载HTTP指纹:[%d]条", n)
	}
	//超时及日志配置
	gonmap.SetLogger(slog.Debug())
	slog.Printf(slog.INFO, "成功加载NMAP探针:[%d]个,指纹[%d]条", gonmap.UsedProbesCount, gonmap.UsedMatchCount)
	//CDN检测初始化

	if _, err := os.Lstat(qqwryPath); os.IsNotExist(err) == true {
		slog.Printf(slog.WARN, "未检测到qqwry.dat,将关闭CDN检测功能，如需开启，请执行kscan --download-qqwry下载该文件")
		slog.Println(slog.INFO, "现在开始下载最新qqwry，请耐心等待！")
		err := cdn.DownloadQQWry()
		if err != nil {
			slog.Println(slog.WARN, "纯真IP库下载失败，请手动下载解压后保存到kscan同一目录")
			slog.Println(slog.WARN, "下载链接： https://qqwry.mirror.noc.one/qqwry.rar")
			slog.Println(slog.WARN, err)
		}
	}
	slog.Println(slog.INFO, "qqwry.dat下载成功！")
	slog.Println(slog.INFO, "检测到qqwry.dat,将自动启动CDN检测功能，可使用-Dn参数关闭该功能")
	scanner.CDNCheck = true
	cdn.Init(qqwryPath)

	slog.Println(slog.INFO, "hydra模块已开启，开始监听暴力破解任务")
	slog.Println(slog.INFO, "当前已开启的hydra模块为：", misc.Intersection(hydra.ProtocolList, hydra.ProtocolList))
	//加载Hydra模块自定义字典
	userDict, uErr := utils.ReadLineData("/zrtx/config/cyberspace/user.txt")
	passDict, pErr := utils.ReadLineData("/zrtx/config/cyberspace/pass.txt")

	if uErr == nil && pErr == nil {
		//slog.Println(slog.DEBUG,userDict,passDict)
		hydra.InitCustomAuthMap(userDict, passDict)
	}

	var restart = "/zrtx/log/cyberspace/restart" + utils.GetDate() + ".json"
	utils.WriteAppend(restart, utils.GetTime())

	//cmd.CleanLog()
	utils.Write("/zrtx/log/cyberspace/worker.log", "")
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split("/zrtx/config/cyberspace", ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Masscan", &global.MasscanSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	cache.NewCacheClient(time.Duration(global.ServerSetting.LifeWindow))

	return nil
}
