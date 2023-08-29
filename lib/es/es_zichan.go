package es

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type ZiChan struct {
	URL                     string `json:URL`
	Keyword                 string `json:Keyword`
	IP                      string `json:IP`
	IP_country              string `json:IP_country`
	IP_province             string `json:IP_province`
	IP_city                 string `json:IP_city`
	IP_isp                  string `json:IP_isp`
	Port                    string `json:Port`
	Service                 string `json:Service`
	Length                  string `json:Length`
	FingerPrint             string `json:FingerPrint`
	Addr                    string `json:Addr`
	Digest                  string `json:Digest`
	Info                    string `json:Info`
	Hostname                string `json:Hostname`
	OperatingSystem         string `json:OperatingSystem`
	DeviceType              string `json:DeviceType`
	ProductName             string `json:ProductName`
	Version                 string `json:Version`
	FoundDomain             string `json:FoundDomain` //body中正则匹配出的域名
	FoundIP                 string `json:FoundIP`
	ICP                     string `json:ICP` //正则匹配出来的icp 号
	ProbeName               string `json:ProbeName`
	Domain                  string `json:Domain`
	ICP_serviceLicence      string `json:ICP_serviceLicence`
	ICP_unitName            string `json:ICP_unitName`
	ICP_updateRecordTime    string `json:ICP_updateRecordTime`
	ICP_domain              string `json:ICP_domain`
	ICP_natureName          string `json:ICP_natureName`
	ICP_serviceId           string `json:ICP_serviceId`
	ICP_domainId            string `json:ICP_domainId`
	Cert_NotBefore          string `json:Cert_NotBefore`
	Cert_NotAfter           string `json:Cert_NotAfter`
	Cert_Subject            string `json:Cert_Subject`
	Cert_Issuer             string `json:Cert_Issuer`
	Cert_Version            string `json:Cert_Version`
	Cert_PublicKeyAlgorithm string `json:Cert_PublicKeyAlgorithm`
	Cert_SignatureAlgorithm string `json:Cert_SignatureAlgorithm`
	Cert_KeyUsage           string `json:Cert_KeyUsage`
	Asn                     string `json:Asn`
	Asn_start_ip            string `json:Asn_start_ip`
	Asn_end_ip              string `json:Asn_end_ip`
	Asn_info                string `json:Asn_info`
	Asn_type                string `json:Asn_type`
	Asn_domain              string `json:Asn_domain`
	MatchRegexString        string `json:MatchRegexString`
	CreateDate              string `json:CreateDate`
	CreateTime              string `json:CreateTime`
	Icon                    string `json:icon` //icon的资源base64的hash 可以用来icon的搜索
}

var index = "zichan"

// 数据创建
func ZiChanCreate(ziChanMap interface{}) {

	zichan := ZiChan{}

	err := mapstructure.Decode(ziChanMap, &zichan)
	if err != nil {
		fmt.Println(err.Error())
	}

	CreateData(index, zichan)
}
