package request

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func HttpRequest(
	urlStr string,
	method string,
	headers map[string]string,
	params map[string]string,
	data any) (*http.Response, error) {
	// 创建URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	query := u.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	u.RawQuery = query.Encode()

	// 将数据编码为JSON
	buf := new(bytes.Buffer)
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	// 创建请求
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// 返回响应，让调用者处理
	return resp, nil
}
