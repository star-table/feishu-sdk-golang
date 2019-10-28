package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var httpClient = &http.Client{}

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{Transport: tr}
}

func Post(url string, params map[string]interface{}, body string) (string, error) {
	resp, err := httpClient.Post(url+ConvertToQueryParams(params), "application/json", strings.NewReader(body))
	defer func(){
		if err := resp.Body.Close(); err != nil{
			fmt.Println(err)
		}
	}()
	return ResponseHandle(resp, err)
}

func Get(url string, params map[string]interface{}) (string, error) {
	resp, err := httpClient.Get(url + ConvertToQueryParams(params))
	defer func(){
		if err := resp.Body.Close(); err != nil{
			fmt.Println(err)
		}
	}()
	return ResponseHandle(resp, err)
}

func ResponseHandle(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ConvertToQueryParams(params map[string]interface{}) string {
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
