package masscan

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os/exec"
	"time"
)

type Address struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

type Ports []struct {
	Protocol string  `xml:"protocol,attr"`
	Portid   string  `xml:"portid,attr"`
	State    State   `xml:"state"`
	Service  Service `xml:"service"`
}

type State struct {
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL string `xml:"reason_ttl,attr"`
}

type Nmaprun struct {
	XMLName    xml.Name `xml:"nmaprun"`
	StartTime  string   `xml:"start,attr"`
	Scanner    string   `xml:"scanner,attr"`
	Version    string   `xml:"version,attr"`
	XmlVersion string   `xml:"xmloutputversion,attr"`
}

type Host struct {
	StartTime       string
	Endtime         string  `xml:"endtime,attr"`
	Address         Address `xml:"address"`
	Ports           Ports   `xml:"ports>port"`
	LastScanTime    int
	LastScanEndTime int
}

type Service struct {
	Name   string `xml:"name,attr"`
	Banner string `xml:"banner,attr"`
}
type Masscan struct {
	SystemPath      string
	Args            []string
	Ports           string
	Ranges          string
	Rate            string
	Exclude         string
	Result          []byte
	LastScanTime    int
	LastScanEndTime int
}

func New() *Masscan {
	return &Masscan{}
}

func (m *Masscan) SetSystemPath(systemPath string) {
	if systemPath != "" {
		m.SystemPath = systemPath
	}
}
func (m *Masscan) SetArgs(arg ...string) {
	m.Args = arg
}

func (m *Masscan) SetRate(rate string) {
	m.Rate = rate
}

//
func (m *Masscan) Run() error {
	var cmd *exec.Cmd
	var outb, errs bytes.Buffer
	if m.Rate != "" {
		m.Args = append(m.Args, "--rate")
		m.Args = append(m.Args, m.Rate)
	}
	if m.Ports != "" {
		m.Args = append(m.Args, "-p")
		m.Args = append(m.Args, m.Ports)
	}
	//m.Args = append(m.Args, "--excludefile")
	//m.Args = append(m.Args, m.Exclude)

	m.Args = append(m.Args, "-oX")
	m.Args = append(m.Args, "-")

	cmd = exec.Command(m.SystemPath, m.Args...)
	fmt.Println("Masscan => ", cmd.Args)
	fmt.Println("Masscan:", cmd)
	cmd.Stdout = &outb
	cmd.Stderr = &errs
	err := cmd.Run()
	if err != nil {
		if errs.Len() > 0 {
			fmt.Printf("masscan run err", err)
			return nil
		}
		return err
	}
	m.Result = outb.Bytes()
	return nil
}

func (m *Masscan) Parse() ([]Host, error) {
	var tmp string
	var hosts []Host
	decoder := xml.NewDecoder(bytes.NewReader(m.Result))
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if t == nil {
			break
		}
		switch res := t.(type) {
		case xml.StartElement:
			time.Sleep(3)
			//if res.Name.Local == "nmaprun" {
			//	for _, v := range res.Attr {
			//		if v.Name.Local == "start" {
			//			tmp = res.Attr[1].Value
			//			break
			//		}
			//	}
			//}
			if res.Name.Local == "host" {
				var host Host
				err := decoder.DecodeElement(&host, &res)
				if err == io.EOF {
					break
				}
				if err != nil {
					return nil, err
				}
				host.StartTime = tmp + "000"
				host.Endtime = host.Endtime + "000"
				hosts = append(hosts, host)
			}
		default:
		}
	}
	return hosts, nil
}
