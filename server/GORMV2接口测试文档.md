| 路由path                                          | 路由描述                | 测试情况 | 备注       |
| ------------------------------------------------- | ----------------------- | :------: | ---------- |
| /base/login                                       | 用户登录                |    √     | SliverHorn |
| /base/register                                    | 用户注册                |    √     | SliverHorn |
| /api/createApi                                    | 创建api                 |    √     | SliverHorn |
| /api/getApiList                                   | 获取api详细信息         |    √     | SliverHorn |
| /api/deleteApi                                    | 删除Api                 |    √     | SliverHorn |
| /api/updateApi                                    | 更新Api                 |    √     | SliverHorn |
| /api/getAllApis                                   | 获取所有api             |    √     | SliverHorn |
| /authority/createAuthority                        | 创建角色                |    √     | SliverHorn |
| /authority/createAuthority                        | 创建子角色              |    √     | SliverHorn |
| /authority/deleteAuthority                        | 删除角色                |    √     | SliverHorn |
| /authority/getAuthorityList                       | 获取角色列表            |    √     | SliverHorn |
| /menu/getMenu                                     | 获取菜单树              |    √     | SliverHorn |
| /menu/getMenuList                                 | 分页获取基础menu列表    |    √     | SliverHorn |
| /menu/addBaseMenu                                 | 新增菜单                |    √     | SliverHorn |
| /menu/getBaseMenuTree                             | 获取用户动态路由        |    √     | SliverHorn |
| /menu/addMenuAuthority                            | 增加menu和角色关联关系  |    √     | SliverHorn |
| /menu/getMenuAuthority                            | 获取指定角色menu        |    √     | SliverHorn |
| /menu/deleteBaseMenu                              | 删除菜单                |    √     | SliverHorn |
| /menu/updateBaseMenu                              | 更新菜单                |    √     | SliverHorn |
| /menu/getBaseMenuById                             | 根据id获取菜单          |    √     | SliverHorn |
| /user/changePassword                              | 修改密码                |    √     | SliverHorn |
| /user/uploadHeaderImg                             | 上传头像                |    √     | SliverHorn |
| /user/getUserList                                 | 获取用户列表            |    √     | SliverHorn |
| /user/setUserAuthority                            | 修改用户角色            |    √     | SliverHorn |
| /fileUploadAndDownload/upload                     | 文件上传示例            |    √     | SliverHorn |
| /fileUploadAndDownload/getFileList                | 获取上传文件列表        |    √     | SliverHorn |
| /casbin/updateCasbin                              | 更改角色api权限         |    √     | SliverHorn |
| /casbin/getPolicyPathByAuthorityId                | 获取权限列表            |    √     | SliverHorn |
| /fileUploadAndDownload/deleteFile                 | 删除文件                |    √     | SliverHorn |
| /jwt/jsonInBlacklist                              | jwt加入黑名单           |    √     | SliverHorn |
| /authority/setDataAuthority                       | 设置角色资源权限        |    √     | SliverHorn |
| /system/getSystemConfig                           | 获取配置文件内容        |    √     | SliverHorn |
| /system/setSystemConfig                           | 设置配置文件内容        |    √     | SliverHorn |
| /customer/customer                                | 创建客户                |    √     | SliverHorn |
| /customer/customer                                | 更新客户                |    √     | SliverHorn |
| /customer/customer                                | 删除客户                |    √     | SliverHorn |
| /customer/customer                                | 获取单一客户            |    √     | SliverHorn |
| /customer/customerList                            | 获取客户列表            |    √     | SliverHorn |
| /autoCode/createTemp                              | 自动化代码              |    √     | SliverHorn |
| /authority/updateAuthority                        | 更新角色信息            |    √     | SliverHorn |
| /authority/copyAuthority                          | 拷贝角色                |    √     | SliverHorn |
| /user/deleteUser                                  | 删除用户                |    √     | SliverHorn |
| /sysDictionaryDetail/createSysDictionaryDetail    | 新增字典内容            |    √     | SliverHorn |
| /sysDictionaryDetail/deleteSysDictionaryDetail    | 删除字典内容            |    √     | SliverHorn |
| /sysDictionaryDetail/updateSysDictionaryDetail    | 更新字典内容            |    √     | SliverHorn |
| /sysDictionaryDetail/findSysDictionaryDetail      | 根据ID获取字典内容      |    √     | SliverHorn |
| /sysDictionaryDetail/getSysDictionaryDetailList   | 获取字典内容列表        |    √     | SliverHorn |
| /sysDictionary/createSysDictionary                | 新增字典                |    √     | SliverHorn |
| /sysDictionary/deleteSysDictionary                | 删除字典                |    √     | SliverHorn |
| /sysDictionary/updateSysDictionary                | 更新字典                |    √     | SliverHorn |
| /sysDictionary/findSysDictionary                  | 根据ID获取字典          |    √     | SliverHorn |
| /sysDictionary/getSysDictionaryList               | 获取字典列表            |    √     | SliverHorn |
| /sysOperationRecord/createSysOperationRecord      | 新增操作记录            |    √     | SliverHorn |
| /sysOperationRecord/deleteSysOperationRecord      | 删除操作记录            |    √     | SliverHorn |
| /sysOperationRecord/findSysOperationRecord        | 根据ID获取操作记录      |    √     | SliverHorn |
| /sysOperationRecord/getSysOperationRecordList     | 获取操作记录列表        |    √     | SliverHorn |
| /autoCode/getTables                               | 获取数据库表            |    √     | SliverHorn |
| /autoCode/getDB                                   | 获取所有数据库          |    √     | SliverHorn |
| /autoCode/getColume                               | 获取所选table的所有字段 |    √     | SliverHorn |
| /sysOperationRecord/deleteSysOperationRecordByIds | 批量删除操作历史        |    √     | SliverHorn |
| /simpleUploader/upload                            | 插件版分片上传          |    √     | SliverHorn |
| /simpleUploader/checkFileMd5                      | 文件完整度验证          |    √     | SliverHorn |
| /simpleUploader/mergeFileMd5                      | 上传完成合并文件        |    √     | SliverHorn |

### Gorm V2 Vs Gorm V1
- in查询,v2不需要括号了
```go
// V2文档
// IN
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');

// V1文档
// IN
db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
//// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');
```

- v2将RecordNotFound链式调用,取而代之是errors.Is(error, gorm.ErrRecordNotFound)

```go
// V2版本代码(https://github.com/flipped-aurora/gin-vue-admin/blob/gva_gormv2_dev/server/service/sys_user.go)
func Register(u model.SysUser) (err error, userInter model.SysUser) {
    var user model.SysUser
    if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound){ // 判断用户名是否注册
        return errors.New("用户名已注册"), userInter
    }
    // 否则 附加uuid 密码md5简单加密 注册
    u.Password = utils.MD5V([]byte(u.Password))
    u.UUID = uuid.NewV4()
    err = global.GVA_DB.Create(&u).Error
	return err, u
}

// V1版本Register(https://github.com/flipped-aurora/gin-vue-admin/blob/master/server/service/sys_user.go)
func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	// 判断用户名是否注册
	notRegister := global.GVA_DB.Where("username = ?", u.Username).First(&user).RecordNotFound()
	// notRegister为false表明读取到了 不能注册
	if !notRegister {
		return errors.New("用户名已注册"), userInter
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = utils.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = global.GVA_DB.Create(&u).Error
	}
	return err, u
}
```

