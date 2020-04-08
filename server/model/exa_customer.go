package model

import (
	"github.com/jinzhu/gorm"
)

type ExaCustomer struct {
	gorm.Model
	CustomerName       string  `json:"customerName"`
	CustomerPhoneData  string  `json:"customerPhoneData"`
	SysUserID          uint    `json:"sysUserId"`
	SysUserAuthorityID string  `json:"sysUserAuthorityID"`
	SysUser            SysUser `json:"sysUser"`
}