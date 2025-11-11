package stacktrace

import (
    "regexp"
    "strconv"
    "strings"
)

// Frame 表示一次栈帧解析结果
type Frame struct {
    File string
    Line int
    Func string
}

var fileLineRe = regexp.MustCompile(`\s*(.+\.go):(\d+)\s*$`)

// FindFinalCaller 从 zap 的 entry.Stack 文本中，解析“最终业务调用方”的文件与行号
// 策略：自顶向下解析，优先选择第一条项目代码帧，过滤第三方库/标准库/框架中间件
func FindFinalCaller(stack string) (Frame, bool) {
    if stack == "" {
        return Frame{}, false
    }
    lines := strings.Split(stack, "\n")
    var currFunc string
    for i := 0; i < len(lines); i++ {
        line := strings.TrimSpace(lines[i])
        if line == "" {
            continue
        }
        if m := fileLineRe.FindStringSubmatch(line); m != nil {
            file := m[1]
            ln, _ := strconv.Atoi(m[2])
            if shouldSkip(file) {
                // 跳过此帧，同时重置函数名以避免错误配对
                currFunc = ""
                continue
            }
            return Frame{File: file, Line: ln, Func: currFunc}, true
        }
        // 记录函数名行，下一行通常是文件:行
        currFunc = line
    }
    return Frame{}, false
}

func shouldSkip(file string) bool {
    // 第三方库与 Go 模块缓存
    if strings.Contains(file, "/go/pkg/mod/") {
        return true
    }
    if strings.Contains(file, "/go.uber.org/") {
        return true
    }
    if strings.Contains(file, "/gorm.io/") {
        return true
    }
    // 标准库
    if strings.Contains(file, "/go/go") && strings.Contains(file, "/src/") { // e.g. /Users/name/go/go1.24.2/src/net/http/server.go
        return true
    }
    // 框架内不需要作为最终调用方的路径
    if strings.Contains(file, "/server/core/zap.go") {
        return true
    }
    if strings.Contains(file, "/server/core/") {
        return true
    }
    if strings.Contains(file, "/server/utils/errorhook/") {
        return true
    }
    if strings.Contains(file, "/server/middleware/") {
        return true
    }
    if strings.Contains(file, "/server/router/") {
        return true
    }
    return false
}