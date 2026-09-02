package mcpTool

type pageResultData[T any] struct {
	List     T     `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}
