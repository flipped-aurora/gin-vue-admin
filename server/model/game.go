package model

type ConnectionToken struct {
	Token string `json:"token" header:"C-Token" binding:"required"`
}

type DestroyResult struct {
	Total   int `json:"total"`
	Success int `json:"success"`
	Failed  int `json:"failed"`
}
