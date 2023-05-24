package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// CommonError 微信返回的通用错误json
type CommonError struct {
	apiName string
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (c *CommonError) Error() string {
	return fmt.Sprintf("%s Error , errcode=%d , errmsg=%s", c.apiName, c.ErrCode, c.ErrMsg)
}

// NewCommonError 新建CommonError错误，对于无errcode和errmsg的返回也可以返回该通用错误
func NewCommonError(apiName string, code int64, msg string) *CommonError {
	return &CommonError{
		apiName: apiName,
		ErrCode: code,
		ErrMsg:  msg,
	}
}

// DecodeWithCommonError 将返回值按照CommonError解析
func DecodeWithCommonError(response []byte, apiName string) (err error) {
	var commError CommonError
	err = json.Unmarshal(response, &commError)
	if err != nil {
		return
	}
	commError.apiName = apiName
	if commError.ErrCode != 0 {
		return &commError
	}
	return nil
}

// DecodeWithError 将返回值按照解析
func DecodeWithError(response []byte, obj interface{}, apiName string) error {
	err := json.Unmarshal(response, obj)
	if err != nil {
		return fmt.Errorf("json Unmarshal Error, err=%v", err)
	}
	responseObj := reflect.ValueOf(obj)
	if !responseObj.IsValid() {
		return fmt.Errorf("obj is invalid")
	}
	commonError := responseObj.Elem().FieldByName("CommonError")
	if !commonError.IsValid() || commonError.Kind() != reflect.Struct {
		return fmt.Errorf("commonError is invalid or not struct")
	}
	errCode := commonError.FieldByName("ErrCode")
	errMsg := commonError.FieldByName("ErrMsg")
	if !errCode.IsValid() || !errMsg.IsValid() {
		return fmt.Errorf("errcode or errmsg is invalid")
	}
	if errCode.Int() != 0 {
		return &CommonError{
			apiName: apiName,
			ErrCode: errCode.Int(),
			ErrMsg:  errMsg.String(),
		}
	}
	return nil
}
