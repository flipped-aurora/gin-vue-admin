package stringsx

import (
	"fmt"
	"strings"
	"testing"
)

func TestParserHtmlTagContent(t *testing.T) {
	content := `
	<a>a标签1</a>test
	<b>b标签</b>你好
	<c>c标签</c>哈哈
	<a>a标签2</a>test11
<a>a标签3</a>test
`
	tagContent := ParserHtmlTagContent(content, "c")
	fmt.Println(strings.Join(tagContent, "\n"))
}
func TestParserHtmlTagContent2(t *testing.T) {
	content := `
	<Logger>{"a_command":"hi","a_soft":"json","a_soft_classify":"/beiluo/json","a_tenant":"beiluo","level":"INFO","msg":"hello info log","stack":"D:/code/github.com/apphub/server/soft_cmd/beiluo/json/main.go:24","ts":1725606536}</Logger>
	<Response>{"status_code":200,"msg":"","content_type":"text","has_file":false,"path":"","delete_file":false,"data":"hello info res"}</Response>
	<UserCost>0s</UserCost>
`
	tagContent := ParserHtmlTagContent(content, "UserCost")
	fmt.Println(strings.Join(tagContent, "\n"))
}
func BenchmarkParserHtmlTagContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		content := `
	<a>a标签1</a>test
	<b>b标签</b>你好
	<c>c标签</c>哈哈
	<a>a标签2</a>test11
<a>a标签3</a>test
`
		_ = ParserHtmlTagContent(content, "c")
	}
}

func BenchmarkParserHtmlTagContent2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		content := `
	<Logger>{"a_command":"hi","a_soft":"json","a_soft_classify":"/beiluo/json","a_tenant":"beiluo","level":"INFO","msg":"hello info log","stack":"D:/code/github.com/apphub/server/soft_cmd/beiluo/json/main.go:24","ts":1725606536}</Logger>
	<Response>{"status_code":200,"msg":"","content_type":"text","has_file":false,"path":"","delete_file":false,"data":"hello info res"}</Response>
	<UserCost>0s</UserCost>
`
		_ = ParserHtmlTagContent(content, "Logger")

	}
}
