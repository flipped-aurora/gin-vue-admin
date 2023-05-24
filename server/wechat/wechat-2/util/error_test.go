package util

import "testing"

var okErrData string = `{"errcode": 0}`
var errData string = `{"errcode": 43101, "errmsg": "user refuse to accept the msg"}`
var expectError string = "Send Error , errcode=43101 , errmsg=user refuse to accept the msg"

func TestDecodeWithCommonErrorNoError(t *testing.T) {
	err := DecodeWithCommonError([]byte(okErrData), "Send")
	if err != nil {
		t.Error("DecodeWithCommonError should not return error")
		return
	}
}

func TestDecodeWithCommonError(t *testing.T) {
	err := DecodeWithCommonError([]byte(errData), "Send")
	if err == nil {
		t.Error("DecodeWithCommonError should return error")
		return
	}

	cErr, ok := err.(*CommonError)
	if !ok {
		t.Errorf("DecodeWithCommonError should return *CommonError but %T", err)
		return
	}
	if !(cErr.ErrCode == 43101 && cErr.ErrMsg == "user refuse to accept the msg" && cErr.Error() == expectError) {
		t.Error("DecodeWithCommonError return bad *CommonError")
		return
	}
}

func TestDecodeWithError(t *testing.T) {
	type DE struct {
		CommonError
	}
	var obj DE
	err := DecodeWithError([]byte(errData), &obj, "Send")
	if err == nil {
		t.Error("DecodeWithError should return error")
		return
	}

	cErr, ok := err.(*CommonError)
	if !ok {
		t.Errorf("DecodeWithError should return *CommonError but %T", err)
		return
	}
	if !(cErr.ErrCode == 43101 && cErr.ErrMsg == "user refuse to accept the msg" && cErr.Error() == expectError) {
		t.Error("DecodeWithError return bad *CommonError")
		return
	}
}
