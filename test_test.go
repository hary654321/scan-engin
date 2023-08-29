package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/robfig/cron/v3"
	"log"
	"math"
	"reflect"
	"testing"
	"time"
	"zrWorker/core/slog"
	"zrWorker/pkg/utils"
	"zrWorker/run"
)

func TestMd5(t *testing.T) {
	pwd := utils.Md5("http://172.17.0.1:9200")
	print(pwd)
}

func TestStruct(t *testing.T) {

	type hello struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	h := hello{
		Name: "小王",
		Age:  15,
	}
	retH := reflect.TypeOf(h)
	//获取结构体里的名称
	for i := 0; i < retH.NumField(); i++ {
		field := retH.Field(i)
		fmt.Println("结构体里的字段名", field.Name)
		fmt.Println("结构体里的字段属性:", field.Type)
		fmt.Println("结构体里面的字段的tag标签", field.Tag)
	}
}

func TestCron(t *testing.T) {
	//创建任务调度器实例
	c := cron.New()
	//注册任务到调度器，注册的任务都是异步执行的。
	c.AddFunc("1 * * * * *", func() {
		fmt.Println("1  r run...")
	})
	//注册任务到调度器，注册的任务都是异步执行的。
	c.AddFunc("0 30 * * * *", func() {
		fmt.Println("every hour on the half hour run...")
	})
	c.AddFunc("@hourly", func() {
		fmt.Println("every hour run...")
	})
	//启动计划任务
	c.Start()
}

// 定时任务
func jobTask() {
	fmt.Printf("任务启动: %s \n", time.Now().Format("2006-01-02 15:04:05"))
}
func TestCronS(t *testing.T) {
	// 创建一个cron对象
	c := cron.New()

	// 任务调度
	enterId, err := c.AddFunc("@every 3s", jobTask)
	if err != nil {
		panic(err)
	}
	fmt.Printf("任务id是 %d \n", enterId)

	// 同步执行任务会阻塞当前执行顺序  一般使用Start()
	//c.Run()
	//fmt.Println("当前执行顺序.......")

	// goroutine 协程启动定时任务(看到后面Start函数和run()函数，就会明白启动这一步也可以写在任务调度之前执行)
	c.Start()
	// Start()内部有一个running 布尔值 限制只有一个Cron对象启动 所以这一步多个 c.Start() 也只会有一个运行
	c.Start()
	c.Start()

	// 用于阻塞 后面可以使用 select {} 阻塞
	time.Sleep(time.Second * 1000)

	// 关闭定时任务(其实不关闭也可以，主进程直接结束了, 内部的goroutine协程也会自动结束)
	c.Stop()

}

func InitFinger() {

}

/*
*
遍历域名
*/
func Check4dight() {

	CharStr := "0123456789abcdefghijklmnopqrstuvwxyz"
	for {
		i := utils.RandInt(35)
		n := utils.RandInt(35)
		m := utils.RandInt(35)
		j := utils.RandInt(35)
		domain := fmt.Sprintf("%c%c%c%c.com", CharStr[i], CharStr[n], CharStr[m], CharStr[j])
		slog.Printf(slog.INFO, "域名%s", domain)
		t := time.NewTicker(time.Millisecond * 200)
		<-t.C
		run.PushTarget(domain)

	}
}

func TestJt(t *testing.T) {
	// 禁用chrome headless
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoDefaultBrowserCheck, //不检查默认浏览器
		chromedp.Flag("headless", true),
		chromedp.Flag("blink-settings", "imagesEnabled=true"), //开启图像界面,重点是开启这个
		chromedp.Flag("ignore-certificate-errors", true),      //忽略错误
		chromedp.Flag("disable-web-security", true),           //禁用网络安全标志
		chromedp.Flag("disable-extensions", true),             //开启插件支持
		chromedp.Flag("disable-default-apps", true),
		chromedp.WindowSize(1920, 1080),    // 设置浏览器分辨率（窗口大小）
		chromedp.Flag("disable-gpu", true), //开启gpu渲染
		chromedp.Flag("hide-scrollbars", true),
		chromedp.Flag("mute-audio", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.NoFirstRun, //设置网站不是首次运行
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"), //设置UserAgent
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// 创建上下文实例
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// 创建超时上下文
	ctx, cancel = context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	//导航到目标页面，等待一个元素，捕捉元素的截图
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(`https://baidu.com/`, 100, &buf)); err != nil {
		log.Fatal(err)
	}
	utils.WritePng("a", buf)

}

// 获取整个浏览器窗口的截图（全屏）
// 这将模拟浏览器操作设置。
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		//chromedp.WaitVisible("style"),
		chromedp.Sleep(10 * time.Second),
		//chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 得到布局页面
			_, _, _, _, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// 浏览器视窗设置模拟
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// 捕捉屏幕截图
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func TestZipFile(t *testing.T) {
	res := utils.ZipFile("/zrtx/log/cyberspace/2023-02-07", "/zrtx/log/cyberspace/2023-02-07/all.zip", "*.json")
	println(res)
}
