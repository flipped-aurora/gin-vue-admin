// Package httpx ...
package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// RequestContext ...
type RequestContext struct {
	StartTime time.Time
	EndTime   time.Time
	Cost      time.Duration
	Url       string
	ReqBody   string
	Method    string
	//ReqBody       string
	ResBodyString string
	resetBody     string
	Code          int
}

// ResetLogBody ...
func (r *RequestContext) ResetLogBody(resetLogBody string) *RequestContext {
	r.resetBody = resetLogBody
	return r
}

// String
func (r *RequestContext) String(messages ...string) string {
	if r != nil {
		body := r.ResBodyString
		if r.resetBody != "" {
			body = r.resetBody
		}
		msg := ""
		if len(messages) > 0 {
			msg = messages[0]
		}
		return fmt.Sprintf("msg:%s request method:%v cost:%v url:%v body:%v "+
			",res_code: %v res_body: %v", msg, r.Method, r.Cost, r.Url, r.ReqBody, r.Code, body)
	}
	if len(messages) > 0 {
		return fmt.Sprintf("%s nil", messages[0])
	}
	return "nil"
}

// OK ...
func (r *RequestContext) OK() bool {
	return r.Code == 200
}

func request(
	method string,
	url string, body *string, resBindBond interface{}, headers ...map[string]string) (*RequestContext, error) {
	bdReq := ""
	start := time.Now()
	var bd io.Reader
	if body != nil {
		bd = strings.NewReader(*body)
		bdReq = *body
	}
	request, err := http.NewRequest(method, url, bd)
	if err != nil {
		return nil, err
	}
	if len(headers) > 0 {
		for k, v := range headers[0] {
			request.Header.Set(k, v)
		}
	}
	request.Header.Set("content-type", "application/json")
	client := &http.Client{
		Timeout: time.Second * 60,
	}
	//request.
	resp, err := client.Do(request)
	if err != nil {
		return &RequestContext{
			Url:           url,
			Method:        method,
			StartTime:     start,
			EndTime:       time.Now(),
			Cost:          time.Now().Sub(start),
			ReqBody:       bdReq,
			Code:          0,
			ResBodyString: ""}, err
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RequestContext{
			Url:           url,
			Method:        method,
			StartTime:     start,
			EndTime:       time.Now(),
			Cost:          time.Now().Sub(start),
			ReqBody:       bdReq,
			Code:          0,
			ResBodyString: ""}, err
	}
	reqCtx := RequestContext{
		Url:           url,
		Method:        method,
		StartTime:     start,
		EndTime:       time.Now(),
		Cost:          time.Now().Sub(start),
		ReqBody:       bdReq,
		Code:          resp.StatusCode,
		ResBodyString: string(all)}
	if resBindBond != nil {
		err = json.Unmarshal(all, resBindBond)
		if err != nil {
			return &reqCtx, err
		}
	}
	return &reqCtx, nil
}

// Get ...
func Get(url string, resBindBond interface{}, headers ...map[string]string) (*RequestContext, error) {
	return request(http.MethodGet, url, nil, resBindBond, headers...)
}

// GetString ...
func GetString(url string, headers ...map[string]string) (*RequestContext, error) {
	return request(http.MethodGet, url, nil, nil, headers...)
}

// PostString ...
func PostString(url string, body interface{}, headers ...map[string]string) (*RequestContext, error) {
	return Post(url, body, nil, headers...)
}

// Post ...
func Post(
	url string, body interface{}, resBindBond interface{}, headers ...map[string]string) (*RequestContext, error) {
	var bd string
	if body == nil {
		bd = "{}"
	} else {
		i, ok := body.([]byte)
		if ok {
			bd = string(i)
		} else {
			typeOf := reflect.TypeOf(body)
			kd := typeOf.Kind()
			switch kd {
			case reflect.Slice, reflect.Array, reflect.Struct, reflect.Map:
				marshal, err := json.Marshal(body)
				if err != nil {
					return nil, err
				}
				bd = string(marshal)
			case reflect.String:
				bd = body.(string)
			default:

			}
		}
	}

	return request(http.MethodPost, url, &bd, resBindBond, headers...)
}
