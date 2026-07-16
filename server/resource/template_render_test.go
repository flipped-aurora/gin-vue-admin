package resource

import (
	"go/format"
	"strings"
	"testing"
	"text/template"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/autocode"
)

// 渲染 api/service 模板，断言：
//  1. 模板能正常执行（无模板语法错误）
//  2. 产物是合法 Go 源码（go/format 可解析）
//  3. 产物不再含旧的 global.GVA_LOG / go.uber.org/zap 形态
//
// 这是对「代码生成模板」的静态门禁——go build 照不到 .tpl，必须显式渲染验证。
// 参考 service/system/auto_code_template.go:228 的渲染方式（text/template + GetTemplateFuncMap）。

// baseData 覆盖 api/service 模板用到的常见字段（取 auto_code 模板用到的 key）。
func baseData() map[string]any {
	primary := request.AutoCodeField{
		FieldName:  "ID",
		FieldJson:  "id",
		ColumnName: "id",
		FieldType:  "int",
	}
	primaryPtr := primary // 取地址用于 []*AutoCodeField
	return map[string]any{
		"Package":            "biz",
		"Module":             "github.com/flipped-aurora/gin-vue-admin/server",
		"StructName":         "SysFoo",
		"Abbreviation":       "sysFoo",
		"Description":        "示例",
		"BusinessDB":         "",
		"OnlyTemplate":       false,
		"IsTree":             false,
		"AutoCreateResource": false,
		"HasDataSource":      false,
		"GvaModel":           true,
		"NeedSort":           false,
		"Fields":             []*request.AutoCodeField{&primaryPtr},
		"PrimaryField":       primary,
		"FuncName":           "DoSomething",
		"FuncDesc":           "示例函数",
		"Router":             "doSomething",
		"Method":             "get",
	}
}

func renderTB(t *testing.T, name, path string, data map[string]any) string {
	t.Helper()
	tpl, err := template.New(name).Funcs(autocode.GetTemplateFuncMap()).ParseFiles(path)
	if err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	var sb strings.Builder
	if err := tpl.Execute(&sb, data); err != nil {
		t.Fatalf("execute %s: %v", path, err)
	}
	return sb.String()
}

// assertGoSource 用 go/format 校验产物是合法 Go 源码（仅语法，不编译）。
func assertGoSource(t *testing.T, src, label string) {
	t.Helper()
	if _, err := format.Source([]byte(src)); err != nil {
		t.Errorf("%s 产物不是合法 Go 源码: %v\n--- 产物 ---\n%s", label, err, src)
	}
}

func TestPackageApiTemplate(t *testing.T) {
	out := renderTB(t, "api.go.tpl", "package/server/api/api.go.tpl", baseData())
	assertGoSource(t, out, "package/api")
	if strings.Contains(out, "global.GVA_LOG") {
		t.Error("package/api 模板产物仍含 global.GVA_LOG")
	}
	if strings.Contains(out, `"go.uber.org/zap"`) {
		t.Error("package/api 模板产物仍含 zap import")
	}
	if !strings.Contains(out, "utils/logger") {
		t.Error("package/api 模板产物缺 logger import")
	}
}

func TestPackageServiceTemplate(t *testing.T) {
	// service 模板默认渲染 IsAdd 分支（追加片段），需强制走 else（完整定义）
	data := baseData()
	data["IsAdd"] = false
	out := renderTB(t, "service.go.tpl", "package/server/service/service.go.tpl", data)
	assertGoSource(t, out, "package/service")
	// service 方法 DB 调用必须带 WithContext(ctx)
	if !strings.Contains(out, ".WithContext(ctx)") {
		t.Error("package/service 模板产物缺 .WithContext(ctx)")
	}
}

func TestPluginApiTemplate(t *testing.T) {
	out := renderTB(t, "api.go.tpl", "plugin/server/api/api.go.tpl", baseData())
	assertGoSource(t, out, "plugin/api")
	if strings.Contains(out, "global.GVA_LOG") {
		t.Error("plugin/api 模板产物仍含 global.GVA_LOG")
	}
	if !strings.Contains(out, "utils/logger") {
		t.Error("plugin/api 模板产物缺 logger import")
	}
}

func TestPluginServiceTemplate(t *testing.T) {
	data := baseData()
	data["IsAdd"] = false
	out := renderTB(t, "service.go.tpl", "plugin/server/service/service.go.tpl", data)
	assertGoSource(t, out, "plugin/service")
	if !strings.Contains(out, ".WithContext(ctx)") {
		t.Error("plugin/service 模板产物缺 .WithContext(ctx)")
	}
}

func TestFunctionApiTemplate(t *testing.T) {
	// function 模板生成单函数，包裹进合法 Go 文件再校验
	// 非 plugin 分支（else）
	data := baseData()
	out := renderTB(t, "api.go.tpl", "function/api.go.tpl", data)
	wrapped := "package p\nimport (\n\"context\"\n\"github.com/gin-gonic/gin\"\n\"github.com/flipped-aurora/gin-vue-admin/server/utils/logger\"\n)\nvar _ = context.Background\nvar _ = (*gin.Context)(nil)\nvar _ = logger.WithCtx\n" + out
	assertGoSource(t, wrapped, "function/api")
	if strings.Contains(out, "global.GVA_LOG") {
		t.Error("function/api 模板产物仍含 global.GVA_LOG")
	}
}
