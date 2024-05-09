package utils

import (
	"fmt"
	"testing"
)

func TestGetJSONKeys(t *testing.T) {
	var jsonStr = `
	{
		"Name": "test",
		"TableName": "test",
		"TemplateID": "test",
		"TemplateInfo": "test",
		"Limit": 0
}`
	keys, err := GetJSONKeys(jsonStr)
	if err != nil {
		t.Errorf("GetJSONKeys failed" + err.Error())
		return
	}
	if len(keys) != 5 {
		t.Errorf("GetJSONKeys failed" + err.Error())
		return
	}
	if keys[0] != "Name" {
		t.Errorf("GetJSONKeys failed" + err.Error())

		return
	}
	if keys[1] != "TableName" {
		t.Errorf("GetJSONKeys failed" + err.Error())

		return
	}
	if keys[2] != "TemplateID" {
		t.Errorf("GetJSONKeys failed" + err.Error())

		return
	}
	if keys[3] != "TemplateInfo" {
		t.Errorf("GetJSONKeys failed" + err.Error())

		return
	}
	if keys[4] != "Limit" {
		t.Errorf("GetJSONKeys failed" + err.Error())

		return
	}

	fmt.Println(keys)
}
