package json

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ToJson(obj interface{}) (string, error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func ToJsonIgnoreError(obj interface{}) string {
	jsonStr, _ := ToJson(obj)
	return jsonStr
}

func FromJson(jsonStr string, obj interface{}) error {
	return json.Unmarshal([]byte(jsonStr), obj)
}

func FromJsonIgnoreError(jsonStr string, obj interface{}) {
	_ = FromJson(jsonStr, obj)
}