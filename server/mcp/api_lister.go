package mcpTool

import (
	"context"

	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/gin-gonic/gin"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&ApiLister{})
}

type ApiInfo struct {
	ID          uint   `json:"id,omitempty"`
	Path        string `json:"path"`
	Description string `json:"description,omitempty"`
	ApiGroup    string `json:"apiGroup,omitempty"`
	Method      string `json:"method"`
	Source      string `json:"source"`
}

type ApiListResponse struct {
	Success      bool      `json:"success"`
	Message      string    `json:"message"`
	DatabaseApis []ApiInfo `json:"databaseApis"`
	GinApis      []ApiInfo `json:"ginApis"`
	TotalCount   int       `json:"totalCount"`
}

type mcpRoutesResponse struct {
	Routes gin.RoutesInfo `json:"routes"`
}

type ApiLister struct{}

func (a *ApiLister) New() mcp.Tool {
	return mcp.NewTool("list_all_apis",
		mcp.WithDescription(`获取系统中所有的API接口，分为两组：

**功能说明：**
- 返回数据库中已注册的API列表
- 返回gin框架中实际注册的路由API列表
- 帮助前端判断是使用现有API还是需要创建新的API,如果api在前端未使用且需要前端调用的时候，请到api文件夹下对应模块的js中添加方法并暴露给当前业务调用

**返回数据结构：**
- databaseApis: 数据库中的API记录（包含ID、描述、分组等完整信息）
- ginApis: gin路由中的API（仅包含路径和方法），需要AI根据路径自行揣摩路径的业务含义，例如：/api/user/:id 表示根据用户ID获取用户信息`),
		mcp.WithString("_placeholder",
			mcp.Description("占位符，防止json schema校验失败"),
		),
	)
}

func (a *ApiLister) Handle(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiResp, err := postUpstream[systemRes.SysAPIListResponse](ctx, "/api/getAllApis", map[string]any{})
	if err != nil {
		return nil, err
	}

	routeResp, err := postUpstream[mcpRoutesResponse](ctx, "/autoCode/mcpRoutes", map[string]any{})
	if err != nil {
		return nil, err
	}

	databaseApis := make([]ApiInfo, 0, len(apiResp.Data.Apis))
	for _, api := range apiResp.Data.Apis {
		databaseApis = append(databaseApis, ApiInfo{
			ID:          api.ID,
			Path:        api.Path,
			Description: api.Description,
			ApiGroup:    api.ApiGroup,
			Method:      api.Method,
			Source:      "database",
		})
	}

	ginApis := make([]ApiInfo, 0, len(routeResp.Data.Routes))
	for _, route := range routeResp.Data.Routes {
		ginApis = append(ginApis, ApiInfo{
			Path:   route.Path,
			Method: route.Method,
			Source: "gin",
		})
	}

	return textResultWithJSON("", ApiListResponse{
		Success:      true,
		Message:      "获取API列表成功",
		DatabaseApis: databaseApis,
		GinApis:      ginApis,
		TotalCount:   len(databaseApis) + len(ginApis),
	})
}
