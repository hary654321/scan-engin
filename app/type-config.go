package app

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strings"
	"time"
	"zrWorker/core/hydra"
	"zrWorker/core/slog"
	"zrWorker/global"
	"zrWorker/pkg/utils"
)

type Config struct {
	//Target                       chan string
	TaskPercent                  map[string]int
	Port                         []int
	Output                       *os.File
	Proxy, Encoding              string
	Path, Host                   []string
	OutputJson                   *JSONWriter
	OutputCSV                    *CSVWriter
	Threads                      int
	Timeout                      time.Duration
	ClosePing, Check, CloseColor bool
	ScanVersion                  bool
	Spy                          string
	//hydra
	Hydra, HydraUpdate             bool
	HydraPass, HydraUser, HydraMod []string
	//fofa
	Fofa           []string
	FofaFixKeyword string
	FofaSize       int
	Scan           bool
	//CDN检测模块
	DownloadQQwry bool
	CloseCDN      bool
	//输出修饰
	Match    string
	NotMatch string
}

type JSONWriter struct {
	f *os.File
}

func (jw *JSONWriter) Push(m map[string]string) {

	m["CreateDate"] = utils.GetDate()
	m["CreateTime"] = utils.GetTime()
	if global.ServerSetting.RunMode == "debug" {
		m["Body"] = ""
		m["FingerPrint"] = ""
		m["Response"] = ""
		m["FoundDomain"] = ""
	}
	stat, err := jw.f.Stat()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	jw.f.Seek(stat.Size()-1, 0)
	jsonBuf, _ := json.Marshal(m)
	//jsonBuf = append(jsonBuf, []byte("]")...)
	//if stat.Size() != 2 {
	//	jsonBuf = append([]byte(","), jsonBuf...)
	//}
	jsonBuf = append([]byte("\r\n"), jsonBuf...)
	jw.f.Write(jsonBuf)
}

type CSVWriter struct {
	f     *csv.Writer
	title []string
}

func (cw *CSVWriter) inTitle(title string) bool {
	for _, value := range cw.title {
		if value == title {
			return true
		}
	}
	return false
}

func (cw *CSVWriter) Push(m map[string]string) {
	var cells []string
	for _, key := range cw.title {
		if value, ok := m[key]; ok {
			cells = append(cells, value)
			delete(m, key)
		} else {
			cells = append(cells, "")
		}
	}
	for key, value := range m {
		cells = append(cells, value)
		cw.title = append(cw.title, key)
	}
	cw.f.Write(cells)
	cw.f.Flush()
}

var Setting = New()

func LoadOutputJSON(path string) *JSONWriter {
	if path == "" {
		return nil
	}
	//if _, err := os.Stat(path); err == nil || os.IsExist(err) {
	//	slog.Println(slog.DEBUG, "检测到JSON输出文件已存在，将自动删除该文件：", path)
	//	if err := os.Remove(path); err != nil {
	//		slog.Println(slog.DEBUG, "删除文件失败，请检查：", err)
	//	}
	//}
	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR+os.O_APPEND, 0764)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	jw := &JSONWriter{f}
	jw.f.Seek(0, 0)
	//_, err = jw.f.WriteString(`[]`)
	//if err != nil {
	//	slog.Println(slog.DEBUG, err)
	//}
	return &JSONWriter{f}
}

func loadOutputCSV(path string) *CSVWriter {
	if path == "" {
		return nil
	}

	if _, err := os.Stat(path); err == nil || os.IsExist(err) {
		slog.Println(slog.DEBUG, "检测到CSV输出文件已存在，将自动删除该文件：", path)
		if err := os.Remove(path); err != nil {
			slog.Println(slog.DEBUG, "删除文件失败，请检查：", err)
		}
	}

	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR, 0764)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	writer := &CSVWriter{w, []string{
		"URL", "Keyword", "IP", "Port", "Service", "Length",
		"FingerPrint", "Addr",
		"Digest", "Info", "Hostname", "OperatingSystem",
		"DeviceType", "ProductName", "Version",
		"FoundDomain", "FoundIP", "ICP",
		"ProbeName", "MatchRegexString",
		//"Header", "Cert", "Response", "Body",
	}}
	writer.f.Write(writer.title)
	writer.f.Flush()
	return writer
}

func (c *Config) loadPort() {
	if len(Args.Port) > 0 {
		c.Port = append(c.Port, Args.Port...)
	}
	if Args.Top != 400 {
		c.Port = append(c.Port, TOP_1000[:Args.Top]...)
	}
	if len(c.Port) == 0 {
		c.Port = PORT_SCAN
	}
}

func (c *Config) loadOutput() {
	expr := Args.Output
	if expr == "" {
		return
	}
	f, err := os.OpenFile(expr, os.O_CREATE+os.O_RDWR, 0764)
	if err != nil {
		slog.Println(slog.DEBUG, err.Error())
	} else {
		c.Output = f
	}
}

func (c *Config) loadScanPing() {
	if len(c.Port) < 10 {
		c.ClosePing = true
		slog.Println(slog.INFO, "由于扫描端口数量小于10，已自动关闭主机存活性检测功能")
	} else {
		c.ClosePing = Args.ClosePing
	}
}

func (c *Config) loadHydraMod(splice []string) {
	if len(splice) == 0 {
		c.HydraMod = hydra.ProtocolList
		return
	}
	if splice[0] == "all" {
		c.HydraMod = hydra.ProtocolList
		return
	}
	c.HydraMod = splice
}

func (c *Config) loadFofaField(expr string) []string {
	//判断对象是否为多个
	if strArr := strings.ReplaceAll(expr, "\\,", "[DouHao]"); strings.Count(strArr, ",") > 0 {
		var passArr []string
		for _, str := range strings.Split(strArr, ",") {
			passArr = append(passArr, strings.ReplaceAll(str, "[DouHao]", ","))
		}
		return passArr
	}
	//对象为单个且不为空时直接返回
	if expr != "" {
		return []string{expr}
	}
	return []string{}
}

func New() Config {
	return Config{
		TaskPercent: make(map[string]int),
		Path:        []string{"/"},
		Port:        []int{},
		Output:      nil,
		Proxy:       "",
		Host:        []string{},
		Threads:     800,
		Timeout:     0,
		Encoding:    "utf-8",
	}
}

var TOP_1000 = []int{
	80, 443, 7547, 22, 5060, 8080, 8443, 161, 2083, 2096, 8000,
	21, 2087, 8888, 53, 8089, 2082, 2095, 30005, 2086, 554,
	888, 8081, 4567, 1701, 58000, 8008, 3389, 8085, 25, 49152,
	51005, 1723, 123, 8088, 2000, 500, 1024, 23, 3306, 5985,
	7170, 9000, 81, 5000, 49154, 47001, 37777, 49153, 50805, 111,
	110, 50001, 49155, 139, 445, 587, 8082, 50995, 30010, 143,
	993, 8001, 1194, 995, 135, 14440, 49665, 8291, 1025, 6881,
	465, 14430, 9020, 50996, 51001, 5001, 50999, 9080, 51000, 50998,
	50997, 51003, 49156, 58603, 51002, 88, 1717, 51004, 4433, 7000,
	8090, 20002, 3000, 8084, 9090, 5678, 49157, 8002, 20201, 37443,
	520, 3128, 82, 7848, 9876, 52869, 10250, 51007, 2107, 2103,
	60000, 49666, 2105, 2080, 9200, 2077, 8999, 2222, 1026, 49667,
	2078, 49158, 7777, 8083, 49664, 9001, 5432, 8880, 9010, 137,
	10000, 8181, 10443, 8086, 9999, 6466, 6467, 8020, 10001, 50777,
	49668, 7001, 55555, 8015, 10101, 2053, 40000, 60002, 1080, 4443,
	8010, 2079, 1433, 9100, 85, 9305, 9091, 4444, 6443, 2052,
	4343, 3001, 5523, 18080, 20000, 5006, 90, 32400, 9306, 12345,
	5555, 50580, 2525, 444, 7443, 6363, 9303, 4430, 49669, 1027,
	9304, 30006, 9307, 1883, 50000, 2121, 6000, 52200, 10002, 8728,
	8099, 9443, 9002, 2601, 5900, 49502, 20202, 5353, 8899, 83,
	6699, 2323, 8009, 7080, 646, 8800, 17000, 51200, 10022, 8889,
	3479, 49501, 1234, 7548, 10010, 9998, 8200, 3307, 541, 2049,
	6264, 9003, 8444, 49159, 5061, 5357, 800, 3333, 42235, 1900,
	65004, 3005, 4000, 1000, 10011, 8159, 873, 8087, 9009, 5431,
	6001, 515, 5986, 8100, 631, 8058, 179, 8282, 26, 8096,
	8069, 6379, 1028, 38520, 9500, 7005, 9012, 50011, 10080, 27017,
	9101, 8003, 5222, 22222, 602, 18017, 9092, 8006, 52230, 8580,
	5683, 1688, 4190, 8445, 9800, 999, 119, 52931, 990, 9004,
	10020, 3050, 5005, 3400, 6666, 16001, 4431, 30003, 8022, 7676,
	8989, 8887, 8092, 8383, 8061, 8011, 8091, 30000, 84, 49161,
	8101, 7070, 9191, 8123, 89, 9600, 5090, 6036, 8043, 9102,
	21242, 6789, 8005, 8554, 40005, 1029, 9013, 9444, 42443, 1344,
	8991, 7002, 6060, 1443, 4434, 3080, 43080, 44444, 808, 50002,
	2091, 20001, 4040, 6002, 8012, 9089, 4500, 50050, 44158, 8004,
	19080, 3030, 4080, 10243, 8070, 4911, 4369, 548, 3120, 2200,
	3690, 2012, 9088, 30001, 28080, 10003, 3391, 86, 7999, 9005,
	8050, 8013, 4505, 60001, 3003, 5800, 9093, 5672, 9094, 3100,
	3006, 20080, 4321, 5007, 9109, 3031, 9997, 5002, 8787, 3155,
	3002, 10004, 30004, 3008, 10005, 7100, 8866, 5080, 3268, 2090,
	1967, 8016, 8094, 8021, 9095, 3103, 4445, 10008, 6400, 843,
	3109, 3010, 21300, 1935, 8333, 7500, 6003, 4899, 6100, 8093,
	8105, 8881, 4848, 3443, 3790, 2443, 3570, 3013, 7010, 8040,
	9988, 20443, 8060, 3390, 8110, 2022, 5050, 3261, 6005, 5500,
	9801, 3299, 4022, 8014, 4488, 5443, 8007, 3004, 30007, 15672,
	18443, 7004, 3160, 8883, 1720, 9099, 2020, 8172, 4533, 3580,
	3015, 92, 7081, 50022, 9098, 3097, 8686, 4045, 8180, 4001,
	40029, 12380, 9105, 9062, 9103, 8077, 18081, 32768, 3900, 7788,
	3269, 3157, 3337, 91, 10009, 9081, 4064, 7778, 9021, 5701,
	853, 9030, 8990, 7003, 8111, 50003, 8098, 8023, 8182, 8500,
	4063, 8222, 442, 7071, 9300, 8025, 8031, 8187, 3156, 8018,
	8095, 9991, 2196, 50004, 3131, 10006, 555, 30083, 8585, 5901,
	3170, 99, 20121, 3133, 18083, 3165, 2001, 21305, 9900, 4222,
	9797, 1521, 5600, 9096, 8073, 8765, 10013, 8042, 9050, 3183,
	50100, 6080, 8066, 9211, 6050, 5009, 4002, 8530, 18001, 3388,
	8442, 9051, 2379, 9097, 5704, 7681, 9008, 3195, 6011, 9501,
	3138, 6004, 9210, 3456, 3130, 10023, 5223, 8885, 1988, 7011,
	8848, 7090, 8026, 4555, 3136, 3147, 5601, 3139, 100, 3180,
	9082, 3493, 3144, 3308, 6405, 7007, 28658, 3143, 9111, 9205,
	3351, 52951, 9014, 8197, 8890, 9107, 3141, 6006, 2048, 4999,
	7700, 9212, 2101, 2600, 9026, 8241, 3632, 3179, 9295, 8401,
	96, 4050, 3190, 9993, 5569, 4800, 3148, 3145, 8027, 6008,
	6580, 25565, 10028, 8017, 50005, 3260, 22345, 1043, 30021, 8055,
	9301, 98, 8315, 18018, 6565, 3161, 9108, 9213, 18888, 3883,
	9023, 16000, 3780, 10033, 1042, 6661, 9888, 18000, 7373, 3171,
	6007, 8183, 9106, 8446, 7800, 8112, 5709, 3134, 7780, 3123,
	3176, 8789, 1337, 3174, 3182, 9663, 8019, 4730, 7444, 18088,
	3173, 6433, 20084, 9110, 3188, 3185, 3187, 8205, 3197, 11001,
	3168, 18085, 8108, 6543, 3158, 3132, 3149, 3192, 8028, 10021,
	3146, 3159, 9207, 4880, 24442, 4712, 30011, 19999, 3193, 62016,
	10099, 10015, 3178, 9308, 8935, 7801, 9550, 15006, 4786, 8097,
	3169, 3164, 6348, 7057, 3186, 8051, 9085, 3142, 93, 8045,
	9201, 8161, 8038, 3167, 6380, 8118, 8243, 9943, 9553, 18089,
	3181, 10030, 9006, 8701, 3703, 8103, 902, 3379, 12443, 1031,
	8822, 3162, 2010, 3184, 3166, 50500, 9992, 3986, 30009, 18010,
	9083, 3175, 3137, 7063, 3191, 10026, 3749, 3135, 4840, 3194,
	2050, 7065, 264, 801, 5201, 55055, 8030, 9007, 12578, 8024,
	3163, 18007, 7006, 10014, 6561, 8545, 30050, 10024, 3198, 1554,
	9151, 6009, 5433, 8033, 3172, 10086, 3483, 3177, 8192, 5706,
	18084, 9197, 3022, 8079, 3140, 10084, 8029, 700, 8900, 5702,
	3352, 30017, 50443, 1050, 9310, 7088, 8886, 3892, 10025, 10082,
	8402, 9104, 8035, 10554, 94, 30025, 18003, 12000, 18005, 6590,
	636, 5566, 12321, 3189, 9042, 4035, 3288, 3196, 50010, 8912,
	32080, 48888, 9302, 3382, 10017, 10031, 7171, 9040, 12291, 3150,
	7776, 3273, 8076, 12300, 12292, 3460, 5003, 18002, 8844, 4003,
	8114, 18101, 880, 9215, 18100, 9505, 9086, 9033, 9447, 6601,
	50006, 18090, 45667, 1500, 8855, 69, 12016, 12588, 13200, 3541,
	3127, 60443, 30013, 10089, 5705, 9994, 18013, 50012, 18182, 12358,
	12028, 18004, 4949, 14400, 5557, 12501, 8400, 12392, 9803, 1201,
	50007, 5708, 9704, 4070, 8849, 5703, 12401, 16010, 3542, 50009,
	9700, 12509, 18098, 12431, 10032, 12350, 9208, 15000, 8812, 3052,
	9160, 8857, 7272, 17777, 12500, 18016, 7980, 9087, 12480, 11027,
	18067, 9209, 8460, 18014, 12334, 7082, 8882, 8834, 16443, 5556,
	9019, 10035, 5707, 6510, 32800, 8859, 10943, 9060, 7779, 12448,
	199, 12572, 15001, 45111, 18086, 10255, 9765, 30080, 8916, 447,
	19091, 8480, 6020, 12590, 12445, 95, 8884, 45555, 18073, 6600,
	1199, 1180, 3263, 30019, 9206, 9044, 12506, 12303, 9933, 50008,
	8302, 8581, 12520, 12310, 16080, 12332, 12333, 900, 3303, 7050,
	12290, 11084, 8075, 9658, 8901, 12349, 8041, 12322, 12400, 20005,
	47080, 15555, 12357, 47002, 5100, 10225, 12318, 12478, 16403, 12309,
	8126, 8902, 12523, 12320, 2181, 8160, 12370, 12369, 9966}

var TOP_TEST = []int{8888}

var PORT_SCAN = TOP_1000[:50]
