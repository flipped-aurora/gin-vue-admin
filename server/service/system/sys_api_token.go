package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	sysReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type ApiTokenService struct{}

func (apiVersion *ApiTokenService) CreateApiToken(apiToken system.SysApiToken, days int) (string, error) {
	var user system.SysUser
	if err := global.GVA_DB.Where("id = ?", apiToken.UserID).First(&user).Error; err != nil {
		return "", errors.New("用户不存在")
	}

	hasAuth := false
	for _, auth := range user.Authorities {
		if auth.AuthorityId == apiToken.AuthorityID {
			hasAuth = true
			break
		}
	}
	if !hasAuth && user.AuthorityId != apiToken.AuthorityID {
		return "", errors.New("用户不具备该角色权限")
	}

	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一不同的部分是过期时间

	expireTime := time.Duration(days) * 24 * time.Hour
	if days == -1 {
		expireTime = 100 * 365 * 24 * time.Hour
	}

	bf, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)

	claims := sysReq.CustomClaims{
		BaseClaims: sysReq.BaseClaims{
			UUID:        user.UUID,
			ID:          user.ID,
			Username:    user.Username,
			NickName:    user.NickName,
			AuthorityId: apiToken.AuthorityID,
		},
		BufferTime: int64(bf / time.Second), // 缓冲时间
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
			Issuer:    global.GVA_CONFIG.JWT.Issuer,
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return "", err
	}

	apiToken.Token = token
	apiToken.Status = true
	apiToken.ExpiresAt = time.Now().Add(expireTime)
	err = global.GVA_DB.Create(&apiToken).Error
	return token, err
}

func (apiVersion *ApiTokenService) GetApiTokenList(info sysReq.SysApiTokenSearch) (list []system.SysApiToken, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysApiToken{})

	db = db.Preload("User")

	if info.UserID != 0 {
		db = db.Where("user_id = ?", info.UserID)
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&list).Error
	return list, total, err
}

func (apiVersion *ApiTokenService) DeleteApiToken(id uint) error {
	var apiToken system.SysApiToken
	err := global.GVA_DB.First(&apiToken, id).Error
	if err != nil {
		return err
	}

	jwtService := JwtService{}
	err = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: apiToken.Token})
	if err != nil {
		return err
	}

	return global.GVA_DB.Model(&apiToken).Update("status", false).Error
}
