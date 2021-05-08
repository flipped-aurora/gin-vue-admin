package model

type ConnectionToken struct {
	Token string `json:"token" form:"token" binding:"required"`
}
