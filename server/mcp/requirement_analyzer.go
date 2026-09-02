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
		mcp.WithDescription(`** 智能需求分析与模块设计工具 - 首选入口工具（最高优先级）**

** 重要提示：这是所有MCP工具的首选入口，请优先使用！**

** 核心能力：**
作为资深系统架构师，智能分析用户需求并自动设计完整的模块架构

** 核心功能：**
1. **智能需求解构**：深度分析用户需求，识别核心业务实体、业务流程、数据关系
2. **自动模块设计**：基于需求分析，智能确定需要多少个模块及各模块功能
3. **字段智能推导**：为每个模块自动设计详细字段，包含数据类型、关联关系、字典需求
4. **架构优化建议**：提供模块拆分、关联设计、扩展性等专业建议

** 输出内容：**
- 模块数量和架构设计
- 每个模块的详细字段清单
- 数据类型和关联关系设计
- 字典需求和类型定义
- 模块间关系图和扩展建议

** 适用场景：**
- 用户需求描述不完整，需要智能补全
- 复杂业务系统的模块架构设计
- 需要专业的数据库设计建议
- 想要快速搭建生产级业务系统

** 推荐工作流：**
 requirement_analyzer → gva_analyze → gva_execute → 其他辅助工具
 
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

// generateAIPrompt 生成AI提示词 - 智能分析需求并确定模块结构
func (t *RequirementAnalyzer) generateAIPrompt(userRequirement string) string {
	prompt := fmt.Sprintf(`# 智能需求分析与模块设计任务

## 用户原始需求
%s

## 核心任务
你需要作为一个资深的系统架构师，深度分析用户需求，智能设计出完整的模块架构。

## 分析步骤

### 第一步：需求解构分析
请仔细分析用户需求，识别出：
1. **核心业务实体**（如：用户、商品、订单、疫苗、宠物等）
2. **业务流程**（如：注册、购买、记录、管理等）
3. **数据关系**（实体间的关联关系）
4. **功能模块**（需要哪些独立的管理模块）

### 第二步：模块架构设计
基于需求分析，设计出模块架构，格式如下：

**模块1：[模块名称]**
- 功能描述：[该模块的核心功能]
- 主要字段：[列出关键字段，注明数据类型]
- 关联关系：[与其他模块的关系，明确一对一/一对多]
- 字典需求：[需要哪些字典类型]

**模块2：[模块名称]**
- 功能描述：[该模块的核心功能]
- 主要字段：[列出关键字段，注明数据类型]
- 关联关系：[与其他模块的关系]
- 字典需求：[需要哪些字典类型]

**...**

### 第三步：字段详细设计
为每个模块详细设计字段：

#### 模块1字段清单：
- 字段名1 (数据类型) - 字段描述 [是否必填] [关联信息/字典类型]
- 字段名2 (数据类型) - 字段描述 [是否必填] [关联信息/字典类型]
- ...

#### 模块2字段清单：
- 字段名1 (数据类型) - 字段描述 [是否必填] [关联信息/字典类型]
- ...

## 智能分析指导原则

### 模块拆分原则
1. **单一职责**：每个模块只负责一个核心业务实体
2. **数据完整性**：相关数据应该在同一模块中
3. **业务独立性**：模块应该能够独立完成特定业务功能
4. **扩展性考虑**：为未来功能扩展预留空间

### 字段设计原则
1. **必要性**：只包含业务必需的字段
2. **规范性**：遵循数据库设计规范
3. **关联性**：正确识别实体间关系
4. **字典化**：状态、类型等枚举值使用字典

### 关联关系识别
- **一对一**：一个实体只能关联另一个实体的一个记录
- **一对多**：一个实体可以关联另一个实体的多个记录
- **多对多**：通过中间表实现复杂关联

## 特殊场景处理

### 复杂实体识别
当用户提到某个概念时，要判断它是否需要独立模块：
- **字典处理**：简单的常见的状态、类型（如：开关、性别、完成状态等）
- **独立模块**：复杂实体（如：疫苗管理、宠物档案、注射记录）

## 输出要求

### 必须包含的信息
1. **模块数量**：明确需要几个模块
2. **模块关系图**：用文字描述模块间关系
3. **核心字段**：每个模块的关键字段（至少5-10个）
4. **数据类型**：string、int、bool、time.Time、float64等
5. **关联设计**：明确哪些字段是关联字段
6. **字典需求**：列出需要创建的字典类型

### 严格遵循用户输入
- 如果用户提供了具体字段，**必须使用**用户提供的字段
- 如果用户提供了SQL文件，**严格按照**SQL结构设计
- **不要**随意发散，**不要**添加用户未提及的功能
---

**现在请开始深度分析用户需求："%s"**

请按照上述框架进行系统性分析，确保输出的模块设计既满足当前需求，又具备良好的扩展性。`, userRequirement, userRequirement)

	return prompt
}
