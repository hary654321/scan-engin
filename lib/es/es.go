package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"reflect"
	"zrWorker/pkg/utils"
)

// Elasticsearch demo

type Person struct {
	Name    string `json:name`
	Age     int    `json:age`
	Married bool   `json:married`
}

var client *elastic.Client

// 初始化es驱动
func Init(host string) {
	errorlog := log.New(os.Stdout, "app", log.LstdFlags)
	//to, err := time.ParseDuration(Timeout)
	//if err != nil {
	//	fmt.Printf("time duration: %s", err.Error())
	//}
	//httpCli := &http.Client{
	//	Timeout: to,
	//	Transport: &http.Transport{
	//		TLSClientConfig: &tls.Config{
	//			InsecureSkipVerify: true,
	//		},
	//		TLSHandshakeTimeout: 10 * time.Second,
	//		MaxIdleConns:        100,
	//		MaxIdleConnsPerHost: 100,
	//	},
	//}
	var err error
	client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host), elastic.SetSniff(false))
	if err != nil {
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
	}
	fmt.Printf("Es return with code %d and version %s \n", code, info.Version.Number)
	esversionCode, err := client.ElasticsearchVersion(host)
	if err != nil {
	}
	fmt.Printf("es version %s\n", esversionCode)
}

/*
*====================================================索引=============================================/
/*创建索引
*/
func CreateIndex(index, mappingInfo string) {
	createIndex, err := client.CreateIndex(index).BodyString(mappingInfo).Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if !createIndex.Acknowledged {
			fmt.Println("创建失败")
		} else {
			fmt.Println("创建成功")
		}
	}
}

// DelIndex /**删除索引*/
func DelIndex() {
	client.DeleteIndex("enIndex_20211128").Do(context.Background())
} /*
判断索引是否存在*/
func ExistsIndex() {
	exist, err := client.IndexExists("enIndex_20211128").Do(context.Background())
	if err != nil {
	}
	if !exist {
		fmt.Println("不存在")
	} else {
		fmt.Println("存在")
	}
} /*
索引添加别名*/
func AddAliasses() {
	client.Alias().Add("enIndex_20211129", "enIndex").Do(context.Background())
} /*
关闭索引*/
func CloseIndex() {
	client.CloseIndex("enIndex_20211129").Do(context.Background())
} /*
打开索引*/
func OpenIndex() {
	client.OpenIndex("enIndex_20211129").Do(context.Background())
}

/*=============================================================数据=====================================================================*/
type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	About     string `json:"about"`
}

// 数据创建
func CreateData(index string, data interface{}) {
	//1.使用结构体方式存入到es里面
	//e1 := Employee{"jane2", "Smith", 20, "I like music"}
	put, err := client.Index().Index(index).BodyJson(data).Do(context.Background())
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf(put.Id, put.Index, put.Type)
}

func GetDataById(index string, id string) json.RawMessage {
	get, err := client.Get().Index(index).Id(id).Do(context.Background())
	if err != nil {
		fmt.Printf(err.Error())
	}

	println(get)
	return get.Source
}

func Query(index string, boolQ *elastic.BoolQuery, pageInfo utils.PageInfo) *elastic.SearchResult {

	size := pageInfo.PageSize
	from := (pageInfo.Page - 1) * size
	var res *elastic.SearchResult
	var err error
	res, err = client.Search(index).Query(boolQ).From(from).Size(size).Do(context.Background())
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}

	return res
}

func printEmployee(res *elastic.SearchResult) {
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}

func QueryES(method, path, contentType string, body interface{}) (json.RawMessage, error) {
	reqopts := elastic.PerformRequestOptions{
		Method:      method,
		Path:        path, // build url
		Body:        body,
		ContentType: contentType,
	}
	//timeout := "200"

	//to, err := time.ParseDuration(timeout)
	//if err != nil {
	//	fmt.Printf("time duration: %s", err.Error())
	//	return ZiChan{}, err
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), to)
	//defer cancel()

	esRawResp, err := client.PerformRequest(context.Background(), reqopts)
	if err != nil {
		fmt.Printf("time duration: %s", err.Error())
		return nil, err
	}

	if esRawResp.StatusCode != 200 {
		fmt.Printf("[%s] err: es response code %d", esRawResp.Body, esRawResp.StatusCode)
	}

	return esRawResp.Body, nil
}

type CountJson struct {
	Count  int            `json:"count"`
	Shards map[string]int `json:"_shards"`
}

func GetCount(index string) CountJson {
	res, err := QueryES("GET", "/"+index+"/_count", "application/json;charset=UTF-8", "")

	//println(string(res))
	var esResp CountJson
	if err = json.Unmarshal(res, &esResp); err != nil {
		fmt.Printf("[%v] err: %s", esResp, err.Error())
		return esResp
	}
	return esResp
}

type GroupJson struct {
	Aggregations map[string]interface{} `json:"aggregations"`
}

func GetDateGroup(index, body string) interface{} {
	//body := `{"from":0,"size":0,"aggs":{"Date":{"terms":{"field":"CreateDate"}}}}`
	res, err := QueryES("POST", "/"+index+"/_search", "application/json;charset=UTF-8", body)

	if err != nil {
		println(string(res), err.Error())
	}
	var esResp GroupJson
	if err = json.Unmarshal(res, &esResp); err != nil {
		fmt.Printf("[%v] err: %s", esResp, err.Error())
		return false
	}

	buckets := esResp.Aggregations["Date"]
	bb := buckets.(map[string]interface{})
	list := bb["buckets"]

	return list
}
