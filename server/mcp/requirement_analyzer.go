package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&RequirementAnalyzer{})
}

type RequirementAnalyzer struct{}

// RequirementAnalysisRequest 需求分析请求
type RequirementAnalysisRequest struct {
	UserRequirement string `json:"userRequirement"`
}



// RequirementAnalysisResponse 需求分析响应
type RequirementAnalysisResponse struct {
	AIPrompt string `json:"aiPrompt"` // 给AI的提示词
}

// New 返回工具注册信息
func (t *RequirementAnalyzer) New() mcp.Tool {
	return mcp.NewTool("requirement_analyzer",
		mcp.WithDescription(`**🚀 需求分析工具 - 首选入口工具（最高优先级）**

**⭐ 重要提示：这是所有MCP工具的首选入口，请优先使用！**

**🎯 核心职责：**
将用户的自然语言需求转换为AI可理解的结构化提示词

**📋 工作流程：**
1. 接收用户自然语言需求描述
2. 生成专业的AI提示词，要求AI将需求梳理为清晰的逻辑步骤：
   - **1. 第一步功能描述**
   - **2. 第二步功能描述** 
   - **3. 第三步功能描述**
   - **...**
3. 指导后续使用 gva_auto_generate 工具进行代码生成

**✅ 适用场景：**
- 用户有新的业务需求需要开发
- 需要创建新的功能模块
- 想要快速搭建业务系统
- 需求描述比较模糊，需要AI帮助梳理

**❌ 不负责的事情：**
- 不生成具体的包名和模块名（交给 gva_auto_generate）
- 不进行代码生成（交给 gva_auto_generate）
- 不创建数据库表结构（交给 gva_auto_generate）

**🔄 推荐工作流：**
requirement_analyzer → gva_auto_generate → 其他辅助工具

`),
		mcp.WithString("userRequirement",
			mcp.Required(),
			mcp.Description("用户的需求描述，支持自然语言，如：'我要做一个猫舍管理系统，用来录入猫的信息，并且记录每只猫每天的活动信息'"),
		),
	)
}

// Handle 处理工具调用
func (t *RequirementAnalyzer) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	userRequirement, ok := request.GetArguments()["userRequirement"].(string)
	if !ok || userRequirement == "" {
		return nil, errors.New("参数错误：userRequirement 必须是非空字符串")
	}

	// 分析用户需求
	analysisResponse, err := t.analyzeRequirement(userRequirement)
	if err != nil {
		return nil, fmt.Errorf("需求分析失败: %v", err)
	}

	// 序列化响应
	responseData, err := json.Marshal(analysisResponse)
	if err != nil {
		return nil, fmt.Errorf("序列化响应失败: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(responseData)),
		},
	}, nil
}

// analyzeRequirement 分析用户需求 - 专注于AI需求传递
func (t *RequirementAnalyzer) analyzeRequirement(userRequirement string) (*RequirementAnalysisResponse, error) {
	// 生成AI提示词 - 这是唯一功能
	aiPrompt := t.generateAIPrompt(userRequirement)

	return &RequirementAnalysisResponse{
		AIPrompt: aiPrompt,
	}, nil
}

// generateAIPrompt 生成AI提示词 - 要求AI梳理逻辑为1xxx2xxx格式
func (t *RequirementAnalyzer) generateAIPrompt(userRequirement string) string {
	prompt := fmt.Sprintf(`# 🤖 AI需求逻辑梳理任务

## 📝 用户原始需求
%s

## 🎯 AI任务要求
请将上述用户需求梳理成清晰的逻辑步骤，格式要求：

**1. 第一步功能描述**
**2. 第二步功能描述**  
**3. 第三步功能描述**
**...**

## 📋 梳理要求
- 将需求拆解为具体的功能步骤
- 每个步骤用数字编号（1、2、3...）
- 步骤描述要清晰、具体、可执行
- 按照业务逻辑顺序排列
- 考虑数据流和用户操作流程

## 🔄 后续流程
梳理完成后，请使用 gva_auto_generate 工具进行代码生成：
- gva_auto_generate 会根据梳理的逻辑步骤自动生成包名、模块名
- gva_auto_generate 会设计数据表结构和API接口
- gva_auto_generate 会生成完整的前后端代码


现在请开始梳理用户需求："%s"`, userRequirement, userRequirement)

	return prompt
}
