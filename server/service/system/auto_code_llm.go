package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	commonResp "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/request"
	"github.com/goccy/go-json"
	"io"
	"strings"
)

// LLMAuto 调用大模型服务，返回生成结果数据
// 入参为通用 JSONMap，需包含 mode（例如 ai/butler/eye/painter 等）以及业务 prompt/payload
func (s *AutoCodeService) LLMAuto(ctx context.Context, llm common.JSONMap) (interface{}, error) {
	if global.GVA_CONFIG.AutoCode.AiPath == "" {
		return nil, errors.New("请先前往插件市场个人中心获取AiPath并填入config.yaml中")
	}

	// 构建调用路径：{AiPath} 中的 {FUNC} 由 mode 替换
	mode := fmt.Sprintf("%v", llm["mode"]) // 统一转字符串，避免 nil 造成路径异常
	path := strings.ReplaceAll(global.GVA_CONFIG.AutoCode.AiPath, "{FUNC}", fmt.Sprintf("api/chat/%s", mode))

	res, err := request.HttpRequest(
		path,
		"POST",
		nil,
		nil,
		llm,
	)
	if err != nil {
		return nil, fmt.Errorf("大模型生成失败: %w", err)
	}
	defer res.Body.Close()

	var resStruct commonResp.Response
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("读取大模型响应失败: %w", err)
	}
	if err = json.Unmarshal(b, &resStruct); err != nil {
		return nil, fmt.Errorf("解析大模型响应失败: %w", err)
	}
	if resStruct.Code == 7 { // 业务约定：7 表示模型生成失败
		return nil, fmt.Errorf("大模型生成失败: %s", resStruct.Msg)
	}
	return resStruct.Data, nil
}
