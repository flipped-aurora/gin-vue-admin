package mcpTool

// pageResultData 解码上游 response.PageResult 的泛型分页载体
type pageResultData[T any] struct {
	List     T     `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}
