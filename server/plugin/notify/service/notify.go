package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notify/global"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type NotifyService struct {
}

//@author: [Espoir](https://github.com/nightsimon)
//@function: SendTextMessage
//@description: 发送钉钉文字信息
//@params content string发送的文字内容
//@params atMobiles []string 艾特的手机号
//@params isAtAll bool 是否艾特全体
//@return: err error

func (e *NotifyService) SendTextMessage(content string, atMobiles []string, isAtAll bool) (err error) {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": content,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   isAtAll,
		},
	}
	return SendMessage(msg)
}

//@author: [Espoir](https://github.com/nightsimon)
//@function: SendLinkMessage
//@description: 发送钉钉图文链接信息
//@params content string 发送的文字内容
//@params title string 发送的标题
//@params picUrl string 艾特的手机号
//@params messageUrl string 是否艾特全体
//@return: err error

func (e *NotifyService) SendLinkMessage(content, title, picUrl, messageUrl string) (err error) {
	msg := map[string]interface{}{
		"msgtype": "link",
		"link": map[string]string{
			"text":       content,
			"title":      title,
			"picUrl":     picUrl,
			"messageUrl": messageUrl,
		},
	}
	return SendMessage(msg)
}

//@author: [Espoir](https://github.com/nightsimon)
//@function: SendMarkdownMessage
//@description: 发送钉钉Markdown信息
//@params content 发送的文字内容
//@params title 发送的标题
//@params atMobiles []string 艾特的手机号
//@params isAtAll bool 是否艾特全体
//@return: err error

func (e *NotifyService) SendMarkdownMessage(content, title string, atMobiles []string, isAtAll bool) (err error) {
	msg := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"text":  content,
			"title": title,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   isAtAll,
		},
	}
	return SendMessage(msg)
}

func SendMessage(msg interface{}) error {
	body := bytes.NewBuffer(nil)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		return fmt.Errorf("msg json failed, msg: %v, err: %v", msg, err.Error())
	}

	value := url.Values{}
	value.Set("access_token", global.GlobalConfig_.Token)
	if global.GlobalConfig_.Secret != "" {
		t := time.Now().UnixNano() / 1e6
		value.Set("timestamp", fmt.Sprintf("%d", t))
		value.Set("sign", sign(t, global.GlobalConfig_.Secret))
	}

	request, err := http.NewRequest(http.MethodPost, global.GlobalConfig_.Url, body)
	if err != nil {
		return fmt.Errorf("error request: %v", err.Error())
	}
	request.URL.RawQuery = value.Encode()
	request.Header.Add("Content-Type", "application/json")
	res, err := (&http.Client{}).Do(request)
	if err != nil {
		return fmt.Errorf("send dingTalk message failed, error: %v", err.Error())
	}
	defer func() { _ = res.Body.Close() }()
	result, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return fmt.Errorf("send dingTalk message failed, %s", httpError(request, res, result, "http code is not 200"))
	}
	if err != nil {
		return fmt.Errorf("send dingTalk message failed, %s", httpError(request, res, result, err.Error()))
	}

	type response struct {
		ErrCode int `json:"errcode"`
	}
	var ret response

	if err := json.Unmarshal(result, &ret); err != nil {
		return fmt.Errorf("send dingTalk message failed, %s", httpError(request, res, result, err.Error()))
	}

	if ret.ErrCode != 0 {
		return fmt.Errorf("send dingTalk message failed, %s", httpError(request, res, result, "errcode is not 0"))
	}

	return nil
}

func httpError(request *http.Request, response *http.Response, body []byte, error string) string {
	return fmt.Sprintf(
		"http request failure, error: %s, status code: %d, %s %s, body:\n%s",
		error,
		response.StatusCode,
		request.Method,
		request.URL.String(),
		string(body),
	)
}
func sign(t int64, secret string) string {
	strToHash := fmt.Sprintf("%d\n%s", t, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}

//	其余方法请参考 https://developers.dingtalk.com/document/robots/custom-robot-access?spm=ding_open_doc.document.0.0.7f8710afbfzduV#topic-2026027
