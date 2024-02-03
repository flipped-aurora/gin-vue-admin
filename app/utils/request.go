package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var client = &http.Client{}

// 发起请求极速版
func Request(url string, method string, data map[string]interface{}) (string, error) {
	datt, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json序列化失败")
		return "", err
	}
	req, err := http.NewRequest(method, url, strings.NewReader(string(datt)))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
