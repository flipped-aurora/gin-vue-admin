package stringsx

import "strings"

//todo 注意这个html标签不支持携带参数
//<a>你好</a> 这个是ok的
//<a herf="http://127.0.0.1">你好</a> 这个是不行的

// ParserHtmlTagContent
// tagName=a可以取出output：a标签1
// content：<a>a标签1</a>test<b>b标签</b>你好
func ParserHtmlTagContent(content string, tagName string) []string {
	leftTag := "<" + tagName + ">"
	rightTag := "</" + tagName + ">"

	dataList := make([]string, 0, 3)
	split := strings.Split(content, leftTag) //<a>你好呢</a>adkadsa
	for _, line := range split {
		index := strings.Index(line, rightTag)
		if index == -1 {
			continue
		}
		s := strings.TrimPrefix(line[:index], leftTag)
		dataList = append(dataList, s)
	}
	return dataList
}
