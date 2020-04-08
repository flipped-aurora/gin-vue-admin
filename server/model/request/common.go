package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id"`
}
