package otherScanner

import (
	"context"
	"github.com/Ullaakut/nmap"
	"log"
	"os/exec"
	"strconv"
	"time"
	"zrWorker/core/slog"
	"zrWorker/pkg/utils"
)

func NmapScan(ip string, port string) (resip, resport, resprotocol string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts(port),
		nmap.WithContext(ctx),
		nmap.WithSkipHostDiscovery(), // s.args = append(s.args, "-Pn") 加上 -Pn 就不去ping主机，因为有的主机防止ping,增加准确度
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
		return
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("Unable to run nmap scan: %v", err)
		return
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			if port.State.State == "open" {
				if port.Service.Name == "microsoft-ds" {
					port.Service.Name = "SMB"
				}

				b := strconv.Itoa(int(port.ID))
				c := string(b)
				return host.Addresses[0].String(), c, port.Service.Name
			}
			return
		}
		return
	}
	return

}

// 获取操作系统信息
func GetOpInfo(ip string) string {
	slog.Println(slog.DEBUG, "开始调用nmap 操作系统：", ip)
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
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

//获取whois 信息
/*
res
|    Domain Name: ZORELWORLD.COM
|    Registry Domain ID: 1857671320_DOMAIN_COM-VRSN
|    Registrar WHOIS Server: whois.ename.com
|    Registrar URL: http://www.ename.net
|    Updated Date: 2021-10-27T07:27:32Z
|    Creation Date: 2014-05-07T09:03:50Z
|    Registry Expiry Date: 2025-05-07T09:03:50Z
|    Registrar: eName Technology Co., Ltd.
|    Registrar IANA ID: 1331
|    Registrar Abuse Contact Email: abuse@ename.com
|    Registrar Abuse Contact Phone: 86.4000044400
|    Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited
|    Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited
|    Name Server: F1G1NS1.DNSPOD.NET
|    Name Server: F1G1NS2.DNSPOD.NET
|    DNSSEC: unsigned
|    URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/

*/

func getWhoisKey() []string {
	whoisKey := []string{
		"Domain Name",
		"Registry Domain ID",
		"Registrar WHOIS Server",
		"Registrar URL",
		"Updated Date",
		"Creation Date",
		"Registry Expiry Date",
		"Registrar",
		"Registrar IANA ID",
		"Registrar Abuse Contact Email",
		"Registrar Abuse Contact Phone",
		"Domain Status",
		"Name Server",
		"DNSSEC",
		"URL of the ICANN Whois Inaccuracy Complaint Form",
	}
	return whoisKey
}

func GetWhois(ip string) map[string]string {
	slog.Println(slog.DEBUG, "开始调用nmap 操作系统：", ip)
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
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

	var whoisInfo map[string]string

	getWhoisKey := getWhoisKey()

	for _, line := range strArr {
		for _, key := range getWhoisKey {
			//slog.Println(slog.DEBUG,  line)
			if utils.GetStrACount(key, line) > 0 {
				whoisInfo[key] = utils.SubStrAfter(line, ":")
			}
		}
	}

	elapsed := time.Since(start)
	slog.Println(slog.DEBUG, "该函数执行完成耗时：", elapsed)

	return whoisInfo
}
