package zoomeye

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var logger = Logger(log.New(os.Stdout, "[zoomeye]", log.Ldate|log.Ltime))

type Logger interface {
	Println(...interface{})
	Printf(string, ...interface{})
}

func SetLogger(log Logger) {
	logger = log
}

type Client struct {
	baseUrl, searchPath, apiKey string
	fieldList                   []string
}

type ResponseJson struct {
	Code      int           `json:"code"`
	Total     int           `json:"total"`
	Available int           `json:"available"`
	Matches   []interface{} `json:"matches"`
	Facets    interface{}   `json:"facets"`
}

const (
	baseURL    = "https://api.zoomeye.org"
	searchPath = "/host/search"
	//loginPath  = "/api/v1/info/my"
)

func New() *Client {
	f := &Client{
		baseUrl:    baseURL,
		searchPath: searchPath,
		apiKey:     "7F3623c6-f3b9-bc48e-f74c-43a9122a79c",
		fieldList: []string{
			"host",
			"title",
			"banner",
			"header",
			"ip", "domain", "port", "country", "province",
			"city", "country_name",
			"server",
			"protocol",
			"cert", "isp", "as_organization",
		},
	}
	return f
}

func (f *Client) Search(keyword, facets string) (int, ResponseJson) {
	url := f.baseUrl + f.searchPath
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("query", keyword)
	q.Add("facets", facets)

	req.Header.Add("API-KEY", f.apiKey)

	req.URL.RawQuery = q.Encode()
	var responseJson ResponseJson
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Println(err)
		return 0, responseJson
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Println(err)
		return 0, responseJson
	}

	if err = json.Unmarshal(body, &responseJson); err != nil {
		logger.Println(body, err)
		return 0, responseJson
	}
	logger.Println(responseJson)
	//r := f.makeResult(responseJson)
	return responseJson.Available, responseJson
}

//func (f *Client) makeResult(responseJson ResponseJson) (results []Result) {
//	for _, row := range responseJson.Matches {
//		var result Result
//		m := reflect.ValueOf(&result).Elem()
//		for index, f := range f.fieldList {
//			//首字母大写
//			f = strings.ToUpper(f[:1]) + f[1:]
//			m.FieldByName(f).SetString(row[index])
//		}
//		results = append(results, result)
//	}
//	return results
//}
