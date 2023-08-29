package es

import (
	"testing"
)

var mapping = `{"settings":{"index":{"max_result_window":"1000000000"}},"mappings":{"properties":{"Service":{"type":"keyword"},"IP":{"type":"text"},"IP_country":{"type":"text"},"IP_province":{"type":"text"},"IP_city":{"type":"text"},"IP_isp":{"type":"text"},"Domain":{"type":"text"},"Keyword":{"type":"text"},"URL":{"type":"text"},"Addr":{"type":"text"},"FingerPrint":{"type":"text"},"Hostname":{"type":"text"},"OperatingSystem":{"type":"text"},"ProductName":{"type":"text"},"Length":{"type":"text"},"Digest":{"type":"text"},"DeviceType":{"type":"text"},"FoundIP":{"type":"text"},"MatchRegexString":{"type":"text"},"ICP_serviceLicence":{"type":"text"},"ICP_unitName":{"type":"text"},"ICP_updateRecordTime":{"type":"text"},"ICP_domain":{"type":"text"},"ICP_natureName":{"type":"text"},"ICP_serviceId":{"type":"text"},"ICP_domainId":{"type":"text"},"Cert_NotBefore":{"type":"text"},"Cert_NotAfter":{"type":"text"},"Cert_Subject":{"type":"text"},"Cert_Issuer":{"type":"text"},"Cert_Version":{"type":"text"},"Cert_PublicKeyAlgorithm":{"type":"text"},"Cert_SignatureAlgorithm":{"type":"text"},"Cert_KeyUsage":{"type":"text"},"ProbeName":{"type":"text"},"Version":{"type":"text"},"Info":{"type":"text"},"FoundDomain":{"type":"text"},"Asn":{"type":"text"},"Asn_start_ip":{"type":"text"},"Asn_end_ip":{"type":"text"},"Asn_info":{"type":"text"},"Asn_type":{"type":"text"},"Asn_domain":{"type":"text"},"Port":{"type":"integer"},"CreateTime":{"type":"integer"},"CreateDate":{"type":"date","format":"yyyy-MM-dd"}}}}`

func TestCreatMp(t *testing.T) {
	//mapping := `{"settings":{"index":{"max_result_window":"1000000000"}},"mappings":{"properties":{"FirstName":{"type":"text"},"LastName":{"type":"text"},"Age":{"type":"integer"},"About":{"type":"text"}}}}`

	//"URL", "", "", "Port", "", "Length",
	//	"FingerPrint", "Addr",
	//	"Digest", "Info", "Hostname", "OperatingSystem",
	//	"DeviceType", "ProductName", "Version",
	//	"", "FoundIP", "ICP",
	//	"ProbeName", "MatchRegexString",
	//"Header", "Cert", "Response", "Body",
	Init("http://192.168.56.132:9200") //http://172.16.130.138:9200
	CreateIndex("zichan", mapping)
	//GetDataById("info","1")
	//Query("info")
}

func TestEsCount(t *testing.T) {
	Init("http://172.16.130.138:9200")
	GetCount("zichan")
}

func TestEsGroup(t *testing.T) {
	Init("http://172.16.130.138:9200")
	body := `{"from":0,"size":0,"aggs":{"Date":{"terms":{"field":"Service"}}}}`
	GetDateGroup("zichan", body)
}
