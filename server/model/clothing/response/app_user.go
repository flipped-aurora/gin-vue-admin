package response

type LoginResponse struct {
	User      UserInfo `json:"user"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}

type UserInfo struct {
	ID       uint          `json:"ID"`
	Roles    []interface{} `json:"roles"`
	PhoneNum string        `json:"phoneNum"`
	Username string        `json:"username"`
}

type CheckExistResponse struct {
	IsExist bool   `json:"isExist"` // 是否存在
	Msg     string `json:"msg"`     // 错误信息
}
