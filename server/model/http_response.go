package model

import uuid "github.com/satori/go.uuid"

/****************************** common start ****************************************************/
// 分页公用入参结构体
type PageInfo struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

//根据id查询结构体
type GetById struct {
	Id float64 `json:"id"`
}

/****************************** common end ****************************************************/

/****************************** api start ****************************************************/
//api分页条件查询及排序结构体
type SearchApiParams struct {
	SysApi
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}

/****************************** api end ****************************************************/

/****************************** Authority start ****************************************************/

// 添加角色和menu关系
type AddMenuAuthorityInfo struct {
	Menus       []SysBaseMenu
	AuthorityId string
}

//	根据角色id获取角色
type AuthorityIdInfo struct {
	AuthorityId string
}

/****************************** Authority end ****************************************************/

/****************************** user start ****************************************************/

// 修改密码结构体
type ChangePasswordStutrc struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// 设置用户权限
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}

/****************************** user end ****************************************************/
