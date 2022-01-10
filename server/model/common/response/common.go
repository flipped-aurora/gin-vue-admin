package response

type PageResult struct {
	List     interface{} `json:"list" comment:"列表"`
	Total    int64       `json:"total" comment:"总数"`
	Page     int         `json:"page" comment:"页码"`
	PageSize int         `json:"pageSize" comment:"每页数量"`
}
