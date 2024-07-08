package system

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"testing"
)

func TestSysAutoCode_PackageTemplate(t *testing.T) {
	info := request.AutoCode{
		Package:         "blender",
		PackageName:     "user",
		Abbreviation:    "user",
		HumpPackageName: "user",
	}
	code, enter, err := AutoCodeTemplate.PackageTemplate(context.Background(), info)
	if err != nil {
		t.Error(err)
		return
	}
	a, _ := json.Marshal(code)
	b, _ := json.Marshal(enter)
	t.Log(string(a))
	t.Log(string(b))
}
