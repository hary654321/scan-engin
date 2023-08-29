package otherScanner

import (
	"context"
	"fmt"
	"github.com/Ullaakut/nmap"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"log"
	"os/exec"
	"testing"
	"time"
	"zrWorker/core/slog"
	"zrWorker/pkg/utils"
)

func TestSave(t *testing.T) {
	SaveSubDomain1("law086.com")

}

func TestRun(t *testing.T) {
	cmd := exec.Command("ls", "-l", "/var/log/")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

// ./Starmap -d zorelworld.com -b -rW -oJ -o subDomain.json
func SaveSubDomain1(domain string) string {

	cmd := exec.Command("/zrtx/bin/cyberspace/Starmap", "-d", domain, "-o", "/zrtx/log/cyberspace/2023-01-06/subDomain.json", "-oJ")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	return string(out)
}

func TestNmapScan(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5 minute timeout.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("baidu.com", "qq.com"),
		nmap.WithPorts("80,443,843"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	fmt.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
}

func TestOp(t *testing.T) {
	slog.Println(slog.DEBUG, GetOpInfo1("192.168.79.1"))
}

func GetOpInfo1(ip string) string {
	slog.Println(slog.DEBUG, "开始调用nmap 操作系统：", ip)
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "nmap", "-O", ip)

	out, err := cmd.CombinedOutput()
	if err != nil {
		//slog.Println(slog.DEBUG,  string(out))
		slog.Println(slog.DEBUG, err)
	}
	//slog.Println(slog.INFO,  string(out))

	res := string(out)
	strArr := utils.Explode(res, "\n")

	for _, line := range strArr {
		//slog.Println(slog.DEBUG,  line)
		if utils.GetStrACount("OS details", line) > 0 {
			return utils.SubStrAfter(line, ":")
		}
	}

	elapsed := time.Since(start)
	slog.Println(slog.DEBUG, "该函数执行完成耗时：", elapsed)

	return ""
}

func TestWhois(t *testing.T) {
	GetWhois("a")
}

func TestWhoisa(t *testing.T) {
	result, err := whois.Whois("zorelworld.com")
	if err != nil {
		slog.Println(slog.DEBUG, "GetWhoisInfo：", result)
	}

	info, err := whoisparser.Parse(result)
	if err != nil {
		slog.Println(slog.DEBUG, "GetWhoisInfo：", result)
	}
	utils.WriteJson("c.json", info)

}

func TestSubDo(t *testing.T) {
	res := SaveSubDomain("zorelworld.com", "51yjsteel.com")

	slog.Println(slog.DEBUG, res)
}
