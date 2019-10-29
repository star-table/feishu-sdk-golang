package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/polaris-team/feishu-sdk-golang/core/util/json"
	"github.com/polaris-team/feishu-sdk-golang/core/util/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const defaultContentType = "application/json"

var httpClient = &http.Client{}

type HeaderOption struct {
	Name  string
	Value string
}

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{
		Transport: tr,
		Timeout:   time.Duration(30) * time.Second,
	}
}

func Post(url string, params map[string]interface{}, body string, headerOptions ...HeaderOption) (string, error) {
	fullUrl := url + ConvertToQueryParams(params)
	req, err := http.NewRequest("POST", fullUrl, strings.NewReader(body))
	if err != nil{
		return "", err
	}
	req.Header.Set("Content-Type", defaultContentType)
	for _, headerOption := range headerOptions {
		req.Header.Set(headerOption.Name, headerOption.Value)
	}
	resp, err := httpClient.Do(req)
	defer func(){
		if e := resp.Body.Close(); e != nil{
			fmt.Println(e)
		}
	}()
	return responseHandle(resp, err)
}

func Get(url string, params map[string]interface{}, headerOptions ...HeaderOption) (string, error) {
	fullUrl := url + ConvertToQueryParams(params)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil{
		return "", err
	}
	for _, headerOption := range headerOptions {
		req.Header.Set(headerOption.Name, headerOption.Value)
	}
	resp, err := httpClient.Do(req)
	defer func(){
		if e := resp.Body.Close(); e != nil{
			fmt.Println(e)
		}
	}()
	return responseHandle(resp, err)
}

func responseHandle(resp *http.Response, err error) (string, error) {
	if err != nil {
		log.Error(err)
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return string(b), nil
}

func ConvertToQueryParams(params map[string]interface{}) string {
	paramsJson := json.ToJsonIgnoreError(params)
	params = map[string]interface{}{}
	_ = json.FromJson(paramsJson, &params)

	if &params == nil || len(params) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString("?")
	for k, v := range params {
		if v == nil {
			continue
		}
		buffer.WriteString(fmt.Sprintf("%s=%v&", k, v))
	}
	buffer.Truncate(buffer.Len() - 1)
	return buffer.String()
}