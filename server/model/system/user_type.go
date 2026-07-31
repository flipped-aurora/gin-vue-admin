package system

// UserType identifies the login subject domain carried by a JWT.
type UserType string

const (
	UserTypeAdmin  UserType = "admin"
	UserTypeClient UserType = "client"
)
