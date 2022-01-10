package response

type PageResult struct {
	List     interface{} `json:"list" comment:"列表"`
	Total    int64       `json:"total" comment:"总数" example:"10"`
	Page     int         `json:"page" comment:"页码" example:"1"`
	PageSize int         `json:"pageSize" comment:"每页数量" example:"10"`
}
