package model

type ConnectionToken struct {
	Token string `json:"token" header:"C-Token" binding:"required"`
}
