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
