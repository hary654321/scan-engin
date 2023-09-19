package run

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
	"zrWorker/app"
	"zrWorker/core/cdn"
	"zrWorker/core/hydra"
	"zrWorker/core/scanner"
	"zrWorker/core/slog"
	"zrWorker/global"
	"zrWorker/lib/appfinger"

	"github.com/hary654321/gonmap"
	"github.com/lcvvvv/stdio/chinese"

	//"zrWorker/lib/chrome"
	"zrWorker/lib/color"
	domain2 "zrWorker/lib/domain"
	ip2 "zrWorker/lib/ip"
	"zrWorker/lib/ip/xdb"
	"zrWorker/lib/misc"
	"zrWorker/lib/simplehttp"
	"zrWorker/lib/udp"
	"zrWorker/lib/uri"
	"zrWorker/pkg/utils"
)

var EngineArr map[string]*Engine //引擎数组
var TaskLooP map[string]string   //循环控制

func init() {
	EngineArr = make(map[string]*Engine)
	TaskLooP = make(map[string]string)
}

type Engine struct {
	DomainScanner *scanner.DomainClient
	IPScanner     *scanner.IPClient
	PortScanner   *scanner.PortClient
	URLScanner    *scanner.URLClient
	HydraScanner  *scanner.HydraClient
	RunTaskId     string
	Total         int
	Done          int
	StartTime     int64
}

func NewEngine(runTaskID string) (*Engine, bool) {
	slog.Println(slog.DEBUG, "new NewEngine")
	if EngineArr[runTaskID] != nil {
		slog.Println(slog.WARN, "已有任务：", runTaskID)
		return EngineArr[runTaskID], false
	}

	//下发扫描任务
	var wg = &sync.WaitGroup{}
	wg.Add(5)
	var e = &Engine{
		DomainScanner: generateDomainScanner(runTaskID, wg),
		IPScanner:     generateIPScanner(runTaskID, wg),
		PortScanner:   generatePortScanner(runTaskID, wg),
		URLScanner:    generateURLScanner(runTaskID, wg),
		HydraScanner:  generateHydraScanner(runTaskID, wg),
		RunTaskId:     runTaskID,
		StartTime:     time.Now().Unix(),
	}
	e.start()
	//启用看门狗函数定时输出负载情况
	go e.watchDog()
	//wg.Wait()
	EngineArr[runTaskID] = e

	slog.Println(slog.DEBUG, "done NewEngine")
	return e, true
}

func (e *Engine) PushTarget(expr string) {
	slog.Println(slog.DEBUG, "PushTarget：", expr)

	if uri.IsIPv4(expr) {
		//slog.Println(slog.DEBUG, "IsIPv4：", expr)
		//遍历的ip
		e.IPScanner.Push(net.ParseIP(expr))

		return
	}
	if uri.IsIPv6(expr) {
		slog.Println(slog.WARN, "暂时不支持IPv6的扫描对象：", expr)
		return
	}
	if uri.IsCIDR(expr) {

		slog.Println(slog.DEBUG, "IsCIDR：", expr)

		for _, ip := range uri.CIDRToIP(expr) {
			e.PushTarget(ip.String())
		}
		return
	}
	if uri.IsIPRanger(expr) {
		for _, ip := range uri.RangerToIP(expr) {
			e.PushTarget(ip.String())
		}
		return
	}
	if uri.IsDomain(expr) {
		//在这里控制发散
		if !utils.In_array(expr, global.AppSetting.Target) {
			ip := utils.GetDomainIp(expr)
			slog.Println(slog.DEBUG, "IsDomain：")
			ipLocation := ip2.SearchIpAddr(ip)
			if len(global.AppSetting.AdrrArr) > 0 && !utils.In_array(ipLocation.Country, global.AppSetting.AdrrArr) {
				slog.Println(slog.DEBUG, "不发散扫描：", expr)
				return
			}
		}
		//slog.Println(slog.WARN, "IsDomain：", expr)

		e.DomainScanner.Push(expr)
		e.pushURLTarget(uri.URLParse("http://"+expr), nil)
		e.pushURLTarget(uri.URLParse("https://"+expr), nil)
		return
	}
	if uri.IsHostPath(expr) {
		e.pushURLTarget(uri.URLParse("http://"+expr), nil)
		e.pushURLTarget(uri.URLParse("https://"+expr), nil)

		e.PushTarget(uri.GetNetlocWithHostPath(expr))

		return
	}
	if uri.IsNetlocPort(expr) {
		netloc, port := uri.SplitWithNetlocPort(expr)
		if uri.IsIPv4(netloc) {
			e.PortScanner.Push(net.ParseIP(netloc), port)
		}
		if uri.IsDomain(netloc) {
			e.pushURLTarget(uri.URLParse("http://"+expr), nil)
			e.pushURLTarget(uri.URLParse("https://"+expr), nil)
		}

		e.PushTarget(netloc)

		return
	}
	if uri.IsURL(expr) {

		//slog.Println(slog.DEBUG, "IsURL：", expr)
		//重复扫描
		//pushURLTarget(uri.URLParse(expr), nil)
		e.PushTarget(uri.GetNetlocWithURL(expr))

		return
	}
	slog.Println(slog.WARN, "无法识别的Target字符串:", expr)
}

func (e *Engine) pushURLTarget(URL *url.URL, response *gonmap.Response) {
	var cli *http.Client
	//判断是否初始化client
	if app.Setting.Proxy != "" || app.Setting.Timeout != 3*time.Second {
		cli = simplehttp.NewClient()
	}
	//判断是否需要设置代理
	if app.Setting.Proxy != "" {
		simplehttp.SetProxy(cli, app.Setting.Proxy)
	}
	//判断是否需要设置超时参数
	if app.Setting.Timeout != 3*time.Second {
		simplehttp.SetTimeout(cli, app.Setting.Timeout)
	}

	//判断是否存在请求修饰性参数

	e.URLScanner.Push(URL, response, nil, cli)

	//如果存在，则逐一建立请求下发队列
	//var reqs []*http.Request
	//for _, host := range app.Setting.Host {
	//	req, _ := simplehttp.NewRequest(http.MethodGet, URL.String(), nil)
	//	req.Host = host
	//	reqs = append(reqs, req)
	//}
	//for _, path := range app.Setting.Path {
	//	req, _ := simplehttp.NewRequest(http.MethodGet, URL.String()+path, nil)
	//	reqs = append(reqs, req)
	//}
	//for _, req := range reqs {
	//	URLScanner.Push(req.URL, response, req, cli)
	//}
}

func (e *Engine) start() {
	go e.DomainScanner.Start()
	go e.IPScanner.Start()
	go e.PortScanner.Start()
	go e.URLScanner.Start()
	go e.HydraScanner.Start()
	time.Sleep(time.Second * 1)
	slog.Println(slog.INFO, "任务"+e.RunTaskId+"Domain、IP、Port、URL、Hydra引擎已准备就绪")
}

func (e *Engine) stop() {

	if e.DomainScanner.RunningThreads() == 0 && e.DomainScanner.IsDone() == false {
		e.DomainScanner.Stop()
		slog.Println(slog.DEBUG, "任务"+e.RunTaskId+"检测到所有Domian检测任务已完成，Domain扫描引擎已停止")
	}
	if e.IPScanner.RunningThreads() == 0 && e.IPScanner.IsDone() == false {
		e.IPScanner.Stop()
		slog.Println(slog.DEBUG, "任务"+e.RunTaskId+"检测到所有IP检测任务已完成，IP扫描引擎已停止")
	}

	if e.PortScanner.RunningThreads() == 0 && e.PortScanner.IsDone() == false {
		e.PortScanner.Stop()
		slog.Println(slog.DEBUG, "任务"+e.RunTaskId+"检测到所有Port检测任务已完成，Port扫描引擎已停止")
	}
	if e.URLScanner.RunningThreads() == 0 && e.URLScanner.IsDone() == false {
		e.URLScanner.Stop()
		slog.Println(slog.DEBUG, "任务"+e.RunTaskId+"检测到所有URL检测任务已完成，URL扫描引擎已停止")
	}
	if e.HydraScanner.RunningThreads() == 0 && e.HydraScanner.IsDone() == false {
		e.HydraScanner.Stop()
		slog.Println(slog.DEBUG, "任务"+e.RunTaskId+"检测到所有暴力破解任务已完成，暴力破解引擎已停止")
	}
}

func (e *Engine) GetPercent() int {

	slog.Println(slog.DEBUG, e.DomainScanner.DoneCount()+e.IPScanner.DoneCount(), e.Total)
	if e.Total == 0 {
		return 0
	}

	done := e.DomainScanner.DoneCount() + e.IPScanner.DoneCount()
	return done / e.Total * 100
}

func generateDomainScanner(runTaskID string, wg *sync.WaitGroup) *scanner.DomainClient {
	DomainConfig := scanner.DefaultConfig()
	DomainConfig.Threads = global.AppSetting.Threads
	client := scanner.NewDomainScanner(DomainConfig)
	client.HandlerRealIP = func(domain string, ip net.IP) {
		slog.Println(slog.DEBUG, "HandlerRealIP", ip)

		root := domain2.GetRoot(domain)
		client.SaveDomainRootInfo(runTaskID, root)

		if utils.GetStrACount("d", TaskLooP[runTaskID]) <= 0 {
			TaskLooP[runTaskID] += "d"
			subDomains := client.SaveSubDomain(runTaskID, root)
			if len(subDomains) > 1 {
				for _, subDomain := range subDomains {
					slog.Println(slog.WARN, "开始扫描子域名 ：", subDomain)
					//app.Setting.Target <- subDomain
					EngineArr[runTaskID].PushTarget(subDomain)
				}
			}
		}

		EngineArr[runTaskID].IPScanner.Push(ip)
	}
	client.HandlerIsCDN = func(domain, CDNInfo string) {
		outputCDNRecord(runTaskID, domain, CDNInfo)
	}
	client.HandlerError = func(domain string, err error) {
		slog.Println(slog.DEBUG, "DomainScanner Error: ", domain, err)
	}
	client.Defer(func() {
		wg.Done()
	})
	return client
}

func generateIPScanner(runTaskID string, wg *sync.WaitGroup) *scanner.IPClient {
	IPConfig := scanner.DefaultConfig()
	IPConfig.Threads = global.AppSetting.Threads * 2
	IPConfig.Timeout = 8 * time.Second
	client := scanner.NewIPScanner(IPConfig)
	client.HandlerDie = func(addr net.IP) {
		//slog.Println(slog.DEBUG, addr.String(), " is die")
	}
	client.HandlerAlive = func(addr net.IP) {
		slog.Println(slog.DEBUG, addr.String(), "HandlerAlive")
		EngineArr[runTaskID].pushURLTarget(uri.URLParse("http://"+addr.String()), nil)
		EngineArr[runTaskID].pushURLTarget(uri.URLParse("https://"+addr.String()), nil)

		cha := utils.GetStrACount("i", TaskLooP[runTaskID]) <= 0
		ipDomain := client.SaveIpInfo(runTaskID, addr.String(), cha)
		if cha {
			TaskLooP[runTaskID] += "i"
			if len(ipDomain) > 1 {
				for _, d := range ipDomain {
					slog.Println(slog.WARN, "开始扫描ipDomain ：", d)
					go EngineArr[runTaskID].PushTarget(d)
				}
			}
		}

		for _, port := range app.TOP_1000[:100] {
			//slog.Println(slog.DEBUG, "扫描端口：", addr.String(), ":", port)
			go EngineArr[runTaskID].PortScanner.Push(addr, port)
		}

		for _, udpPort := range udp.UdpPort {
			go EngineArr[runTaskID].PortScanner.Push(addr, udpPort)
		}
	}
	client.HandlerError = func(addr net.IP, err error) {
		slog.Println(slog.DEBUG, "IPScanner Error: ", addr.String(), err)
	}
	client.Defer(func() {
		wg.Done()
	})
	return client
}

func generatePortScanner(runTaskID string, wg *sync.WaitGroup) *scanner.PortClient {
	PortConfig := scanner.DefaultConfig()
	PortConfig.Threads = global.AppSetting.Threads * 10
	PortConfig.Timeout = 10 * time.Second // getTimeout(len(app.Setting.Port))
	if app.Setting.ScanVersion == true {
		PortConfig.DeepInspection = true
	}
	client := scanner.NewPortScanner(PortConfig)
	client.HandlerClosed = func(addr net.IP, port int) {

	}
	client.HandlerOpen = func(addr net.IP, port int) {
		outputOpenResponse(runTaskID, addr, port)
	}
	client.HandlerNotMatched = func(addr net.IP, port int, response string) {
		outputUnknownResponse(runTaskID, addr, port, response)
	}
	client.HandlerMatched = func(addr net.IP, port int, response *gonmap.Response) {
		//slog.Println(slog.DEBUG, "HandlerMatched：", response.FingerPrint.Service, addr.String(), port)
		URLRaw := fmt.Sprintf("%s://%s:%d", response.FingerPrint.Service, addr.String(), port)
		URL, _ := url.Parse(URLRaw)
		if appfinger.SupportCheck(URL.Scheme) == true {
			EngineArr[runTaskID].pushURLTarget(URL, response)
			//port的扫描 http 也存到port
			//return
		}
		outputNmapFinger(runTaskID, URL, response)
		if app.Setting.Hydra == true {
			if protocol := response.FingerPrint.Service; hydra.Ok(protocol) {
				EngineArr[runTaskID].HydraScanner.Push(addr, port, protocol)
			}
		}
	}
	client.HandlerError = func(addr net.IP, port int, err error) {
		slog.Println(slog.DEBUG, "PortScanner Error: ", fmt.Sprintf("%s:%d", addr.String(), port), err)
	}

	client.OutputUdpResponse = func(addr net.IP, port int, res *udp.Result) {
		//输出结果
		protocol := gonmap.GuessProtocol(port)
		target := fmt.Sprintf("%s://%s:%d", protocol, addr.String(), port)
		URL, _ := url.Parse(target)

		m := map[string]interface{}{
			"IP":        URL.Hostname(),
			"Port":      strconv.Itoa(port),
			"Keyword":   res.Service.Name,
			"ProbeName": "UDP",
			"UdpInfo":   res,
		}
		utils.WriteJsonAny(utils.GetLogPath(runTaskID, "ipInfo"), m)
		utils.WriteJsonAny(utils.GetLogPath(runTaskID, "portInfo"), m)
	}

	client.Defer(func() {
		wg.Done()
	})
	return client
}

func generateURLScanner(runTaskID string, wg *sync.WaitGroup) *scanner.URLClient {
	URLConfig := scanner.DefaultConfig()
	URLConfig.Threads = global.AppSetting.Threads * 2

	client := scanner.NewURLScanner(URLConfig)
	client.HandlerMatched = func(URL *url.URL, banner *appfinger.Banner, finger *appfinger.FingerPrint) {
		slog.Println(slog.DEBUG, "generateURLScanner: HandlerMatched")
		//在这里把截图存下来
		//go chrome.Screenshot(URL.String())
		EngineArr[runTaskID].outputAppFinger(runTaskID, URL, banner, finger)
	}
	client.HandlerError = func(url *url.URL, err error) {
		//slog.Println(slog.DEBUG, "URLScanner Error: ", url.String(), err)
	}
	client.Defer(func() {
		wg.Done()
	})
	return client
}

func generateHydraScanner(runTaskID string, wg *sync.WaitGroup) *scanner.HydraClient {
	HydraConfig := scanner.DefaultConfig()
	HydraConfig.Threads = global.AppSetting.Threads / 5

	client := scanner.NewHydraScanner(HydraConfig)
	client.HandlerSuccess = func(addr net.IP, port int, protocol string, auth *hydra.Auth) {
		outputHydraSuccess(runTaskID, addr, port, protocol, auth)
	}
	client.HandlerError = func(addr net.IP, port int, protocol string, err error) {
		slog.Println(slog.DEBUG, fmt.Sprintf("%s://%s:%d", protocol, addr.String(), port), err)
	}
	client.Defer(func() {
		wg.Done()
	})
	return client
}

// 输出爆破信息
func outputHydraSuccess(runTaskID string, addr net.IP, port int, protocol string, auth *hydra.Auth) {
	var target = fmt.Sprintf("%s://%s:%d", protocol, addr.String(), port)
	var m = auth.Map()
	URL, _ := url.Parse(target)

	m["HydraSuccess"] = "1"
	outputHandler(runTaskID, URL, color.Important("CrackSuccess"), m, "portInfo")
}

func outputNmapFinger(runTaskID string, URL *url.URL, resp *gonmap.Response) {
	if responseFilter(resp.Raw) == true {
		return
	}
	finger := resp.FingerPrint
	m := misc.ToMap(finger)
	m["Response"] = resp.Raw
	m["IP"] = URL.Hostname()
	m["Port"] = URL.Port()
	//补充归属地信息
	//if app.Setting.CloseCDN == false {
	//	result, _ := cdn.Find(URL.Hostname())
	//	println("addr:::::", result)
	//	m["Addr"] = result
	//}
	outputHandler(runTaskID, URL, finger.Service, m, "ipInfo")

	utils.WriteJsonString(utils.GetLogPath(runTaskID, "portInfo"), m)
}

func (e *Engine) outputAppFinger(runTaskID string, URL *url.URL, banner *appfinger.Banner, finger *appfinger.FingerPrint) {
	if responseFilter(banner.Response, banner.Cert) == true {
		return
	}
	m := misc.ToMap(finger)
	//补充归属地信息
	if app.Setting.CloseCDN == false {
		result, _ := cdn.Find(URL.Hostname())
		//slog.Println(slog.DEBUG,"outputAppFinger  addr:", result)
		m["Addr"] = result
	}
	m["Service"] = URL.Scheme
	m["FoundDomain"] = banner.FoundDomain
	m["FoundIP"] = banner.FoundIP
	m["Response"] = banner.Response
	m["Cert"] = banner.Cert
	m["Header"] = banner.Header
	m["Body"] = banner.Body
	m["ICP"] = banner.ICP
	m["Icon"] = banner.Icon
	m["FingerPrint"] = m["ProductName"]
	delete(m, "ProductName")
	//增加IP、Domain、Port字段
	m["Port"] = uri.GetURLPort(URL)
	if m["Port"] == "" {
		slog.Println(slog.WARN, "无法获取端口号：", URL)
	}
	if hostname := URL.Hostname(); uri.IsIPv4(hostname) {
		m["IP"] = hostname
	} else {
		m["Domain"] = hostname
		m["IP"] = utils.GetDomainIp(hostname)
	}

	outputHandler(runTaskID, URL, banner.Title, m, "ipInfo")
}

// 可以在这里查询任务的进度
func (e *Engine) watchDog() {
	time.Sleep(time.Second * 3)

	for {
		var (
			nDomain = e.DomainScanner.RunningThreads()
			nIP     = e.IPScanner.RunningThreads()
			nPort   = e.PortScanner.RunningThreads()
			nURL    = e.URLScanner.RunningThreads()
			nHydra  = e.HydraScanner.RunningThreads()
		)
		warn := fmt.Sprintf("任务"+e.RunTaskId+"当前存活协程数：Domain：%d 个，IP：%d 个，Port：%d 个，URL：%d 个，Hydra：%d 个  ,进度：%d", nDomain, nIP, nPort, nURL, nHydra, e.GetPercent())
		total := nDomain + nIP + nPort + nURL + nHydra
		if total == 0 && time.Now().Unix()-e.StartTime > 60 {
			e.stop()
			delete(EngineArr, e.RunTaskId)
			delete(TaskLooP, e.RunTaskId)
			break
		}
		slog.Println(slog.WARN, warn)
		time.Sleep(time.Second * 3)
	}
}

// 可以在这里查询任务的进度
func (e *Engine) GetRunningThreads() (res map[string]int, total int) {
	res = make(map[string]int)

	res["nDomain"] = e.DomainScanner.RunningThreads()
	res["nIP"] = e.IPScanner.RunningThreads()
	res["nPort"] = e.PortScanner.RunningThreads()
	res["nURL"] = e.URLScanner.RunningThreads()
	res["nHydra"] = e.HydraScanner.RunningThreads()

	for _, v := range res {
		total += v
	}

	return
}

var overload bool

func (e *Engine) ScanAddr(addr string) {
	// 获取一个城市/省/国家 的ip

	var ipLastPath = "/zrtx/log/cyberspace/ipLast.txt"
	userDict, _ := utils.ReadLineData("/zrtx/config/cyberspace/ip.merge.txt")

	ipLast := strings.Trim(utils.Read(ipLastPath), "\n")

	for _, v := range userDict {
		if utils.GetStrACount(addr, v) > 0 {
			arr := strings.Split(v, "|")
			startIp := xdb.GetIpUint32(arr[0])
			endIp := xdb.GetIpUint32(arr[1])
			for i := startIp; i < endIp; i++ {

				if i < xdb.GetIpUint32(ipLast) {
					continue
				}
				if i%global.AppSetting.ScanSpeed == 0 {
					var (
						nDomain = e.DomainScanner.RunningThreads()
						nIP     = e.IPScanner.RunningThreads()
						nPort   = e.PortScanner.RunningThreads()
						nURL    = e.URLScanner.RunningThreads()
						nHydra  = e.HydraScanner.RunningThreads()
					)
					warn := fmt.Sprintf("watchDog：Domain：%d 个，IP：%d 个，Port：%d 个，URL：%d 个，Hydra：%d 个", nDomain, nIP, nPort, nURL, nHydra)
					slog.Println(slog.WARN, warn)
					slog.Println(slog.WARN, "速度一分钟：", global.AppSetting.ScanSpeed)
					if nIP > int(global.AppSetting.ScanSpeed/10) {
						slog.Println(slog.WARN, "首次超载")
						//time.Sleep(1 * 60 * time.Second)
						overload = true
					}
					time.Sleep(60 * time.Second)
				}
				if overload {
					//超载时候用的port
					nPort := e.PortScanner.RunningThreads()
					nIP := e.IPScanner.RunningThreads()
					max := utils.Max(nPort, nIP)
					if max > int(global.AppSetting.ScanSpeed/10) {
						slog.Println(slog.WARN, "任然超载停止1分钟")
						time.Sleep(1 * 60 * time.Second)
					} else {
						overload = false
					}
				}
				slog.Println(slog.WARN, xdb.Long2IP(i))
				utils.Write(ipLastPath, xdb.Long2IP(i))
				e.PushTarget(xdb.Long2IP(i))
			}
		}
	}
}

func outputCDNRecord(runTaskID, domain, info string) {
	if responseFilter(info) == true {
		return
	}
	//输出结果
	target := fmt.Sprintf("cdn://%s", domain)
	URL, _ := url.Parse(target)

	outputHandler(runTaskID, URL, "CDN资产", map[string]string{
		"CDNInfo": info,
		"Domain":  domain,
	}, "/ipInfo.json")
}

func outputUnknownResponse(runTaskID string, addr net.IP, port int, response string) {
	if responseFilter(response) == true {
		return
	}
	//输出结果
	target := fmt.Sprintf("unknown://%s:%d", addr.String(), port)
	URL, _ := url.Parse(target)

	m := map[string]string{
		"Response": response,
		"IP":       URL.Hostname(),
		"Port":     strconv.Itoa(port),
		"Keyword":  "无法识别该协议",
	}
	outputHandler(runTaskID, URL, "无法识别该协议", m, "ipInfo")

	utils.WriteJsonString(utils.GetLogPath(runTaskID, "portInfo"), m)
}

func outputOpenResponse(runTaskID string, addr net.IP, port int) {
	//输出结果
	protocol := gonmap.GuessProtocol(port)
	target := fmt.Sprintf("%s://%s:%d", protocol, addr.String(), port)
	URL, _ := url.Parse(target)

	m := map[string]string{
		"IP":      URL.Hostname(),
		"Port":    strconv.Itoa(port),
		"Keyword": "response is empty",
	}
	outputHandler(runTaskID, URL, "response is empty", m, "ipInfo")

	utils.WriteJsonString(utils.GetLogPath(runTaskID, "portInfo"), m)
}

func responseFilter(strArgs ...string) bool {
	var match = app.Setting.Match
	var notMatch = app.Setting.NotMatch

	if match != "" {
		for _, str := range strArgs {
			//主要结果中包含关键则，则会显示
			if strings.Contains(str, app.Setting.Match) == true {
				return false
			}
		}
	}

	if notMatch != "" {
		for _, str := range strArgs {
			//主要结果中包含关键则，则会显示
			if strings.Contains(str, app.Setting.NotMatch) == true {
				return true
			}
		}
	}
	return false
}

var (
	disableKey       = []string{"MatchRegexString", "Service", "ProbeName", "Response", "Cert", "Header", "Body", "IP"}
	importantKey     = []string{"ProductName", "DeviceType"}
	varyImportantKey = []string{"Hostname", "FingerPrint", "ICP"}
)

func getHTTPDigest(s string) string {
	var length = 24
	var digestBuf []rune
	_, body := simplehttp.SplitHeaderAndBody(s)
	body = chinese.ToUTF8(body)
	for _, r := range []rune(body) {
		buf := []byte(string(r))
		if len(digestBuf) == length {
			return string(digestBuf)
		}
		if len(buf) > 1 {
			digestBuf = append(digestBuf, r)
		}
	}
	return string(digestBuf) + misc.StrRandomCut(body, length-len(digestBuf))
}

func getRawDigest(s string) string {
	var length = 24
	if len(s) < length {
		return s
	}
	var digestBuf []rune
	for _, r := range []rune(s) {
		if len(digestBuf) == length {
			return string(digestBuf)
		}
		if 0x20 <= r && r <= 0x7E {
			digestBuf = append(digestBuf, r)
		}
	}
	return string(digestBuf) + misc.StrRandomCut(s, length-len(digestBuf))
}

func outputHandler(runTaskID string, URL *url.URL, keyword string, m map[string]string, fileName string) {
	m = misc.FixMap(m)
	if respRaw := m["Response"]; respRaw != "" {
		if m["Service"] == "http" || m["Service"] == "https" {
			m["Digest"] = strconv.Quote(getHTTPDigest(respRaw))
		} else {
			m["Digest"] = strconv.Quote(getRawDigest(respRaw))
		}
	}
	m["Length"] = strconv.Itoa(len(m["Response"]))
	sourceMap := misc.CloneMap(m)
	for _, keyword := range disableKey {
		delete(m, keyword)
	}
	for key, value := range m {
		if key == "FingerPrint" {
			continue
		}
		m[key] = misc.StrRandomCut(value, 24)
	}
	fingerPrint := color.StrMapRandomColor(m, true, importantKey, varyImportantKey)
	fingerPrint = misc.FixLine(fingerPrint)
	//format := "%-30v %-" + strconv.Itoa(misc.AutoWidth(color.Clear(keyword), 26+color.Count(keyword))) + "v %s"
	//printStr := fmt.Sprintf(format, URL.String(), keyword, fingerPrint)

	if m["IP"] == "" {
		//slog.Println(slog.DEBUG, URL.Host)
		m["IP"] = URL.Host
	}
	if uri.IsDomain(m["IP"]) {
		m["IP"] = utils.GetDomainIp(m["IP"])
	} else {
		if utils.GetStrACount(":", m["IP"]) > 0 {
			m["Port"] = utils.SubStrAfter(m["IP"], ":")
			m["IP"] = utils.SubStrBefore(m["IP"], ":")
		}
	}

	if m["Service"] == "" {
		m["Service"] = URL.Scheme
	}
	//es存储
	//es.ZiChanCreate(m)
	//slog.Println(slog.DATA, printStr)

	merge := utils.MergeMap(m, sourceMap)
	utils.WriteJsonString(utils.GetLogPath(runTaskID, fileName), merge)
	if cw := app.Setting.OutputCSV; cw != nil {
		sourceMap["URL"] = URL.String()
		sourceMap["Keyword"] = keyword
		delete(sourceMap, "Header")
		delete(sourceMap, "Cert")
		delete(sourceMap, "Response")
		delete(sourceMap, "Body")
		sourceMap["Digest"] = strconv.Quote(sourceMap["Digest"])
		for key, value := range sourceMap {
			sourceMap[key] = chinese.ToUTF8(value)
		}
		cw.Push(sourceMap)
	}
}
