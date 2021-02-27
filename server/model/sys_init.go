package model

type InitDB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password"`
	DBName   string `json:"db_name" binding:"required"`
}
