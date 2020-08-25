| 路由path                                          | 路由描述                | 测试情况 |
| ------------------------------------------------- | ----------------------- | :------: |
| /base/login                                       | 用户登录                |    √     |
| /base/register                                    | 用户注册                |          |
| /api/createApi                                    | 创建api                 |          |
| /api/getApiList                                   | 获取api详细信息         |          |
| /api/deleteApi                                    | 删除Api                 |          |
| /api/updateApi                                    | 更新Api                 |          |
| /api/getAllApis                                   | 获取所有api             |          |
| /authority/createAuthority                        | 创建角色                |    √     |
| /authority/createAuthority                        | 创建子角色              |    √     |
| /authority/deleteAuthority                        | 删除角色                |    √     |
| /authority/getAuthorityList                       | 获取角色列表            |    √     |
| /menu/getMenu                                     | 获取菜单树              |          |
| /menu/getMenuList                                 | 分页获取基础menu列表    |          |
| /menu/addBaseMenu                                 | 新增菜单                |          |
| /menu/getBaseMenuTree                             | 获取用户动态路由        |          |
| /menu/addMenuAuthority                            | 增加menu和角色关联关系  |          |
| /menu/getMenuAuthority                            | 获取指定角色menu        |          |
| /menu/deleteBaseMenu                              | 删除菜单                |          |
| /menu/updateBaseMenu                              | 更新菜单                |          |
| /menu/getBaseMenuById                             | 根据id获取菜单          |          |
| /user/changePassword                              | 修改密码                |          |
| /user/uploadHeaderImg                             | 上传头像                |          |
| /user/getInfoList                                 | 分页获取用户列表        |          |
| /user/getUserList                                 | 获取用户列表            |          |
| /user/setUserAuthority                            | 修改用户角色            |          |
| /fileUploadAndDownload/upload                     | 文件上传示例            |          |
| /fileUploadAndDownload/getFileList                | 获取上传文件列表        |          |
| /casbin/updateCasbin                              | 更改角色api权限         |          |
| /casbin/getPolicyPathByAuthorityId                | 获取权限列表            |          |
| /fileUploadAndDownload/deleteFile                 | 删除文件                |          |
| /jwt/jsonInBlacklist                              | jwt加入黑名单           |          |
| /authority/setDataAuthority                       | 设置角色资源权限        |          |
| /system/getSystemConfig                           | 获取配置文件内容        |          |
| /system/setSystemConfig                           | 设置配置文件内容        |          |
| /customer/customer                                | 创建客户                |          |
| /customer/customer                                | 更新客户                |          |
| /customer/customer                                | 删除客户                |          |
| /customer/customer                                | 获取单一客户            |          |
| /customer/customerList                            | 获取客户列表            |          |
| /casbin/casbinTest/:pathParam                     | RESTFUL模式测试         |          |
| /autoCode/createTemp                              | 自动化代码              |          |
| /authority/updateAuthority                        | 更新角色信息            |          |
| /authority/copyAuthority                          | 拷贝角色                |          |
| /user/deleteUser                                  | 删除用户                |          |
| /sysDictionaryDetail/createSysDictionaryDetail    | 新增字典内容            |          |
| /sysDictionaryDetail/deleteSysDictionaryDetail    | 新增字典内容            |          |
| /sysDictionaryDetail/updateSysDictionaryDetail    | 新增字典内容            |          |
| /sysDictionaryDetail/findSysDictionaryDetail      | 根据ID获取字典内容      |          |
| /sysDictionaryDetail/getSysDictionaryDetailList   | 获取字典内容列表        |          |
| /sysDictionary/createSysDictionary                | 新增字典                |          |
| /sysDictionary/deleteSysDictionary                | 删除字典                |          |
| /sysDictionary/updateSysDictionary                | 更新字典                |          |
| /sysDictionary/findSysDictionary                  | 根据ID获取字典          |          |
| /sysDictionary/getSysDictionaryList               | 获取字典列表            |          |
| /sysOperationRecord/createSysOperationRecord      | 新增操作记录            |          |
| /sysOperationRecord/deleteSysOperationRecord      | 删除操作记录            |          |
| /sysOperationRecord/updateSysOperationRecord      | 更新操作记录            |          |
| /sysOperationRecord/findSysOperationRecord        | 根据ID获取操作记录      |          |
| /sysOperationRecord/getSysOperationRecordList     | 获取操作记录列表        |          |
| /autoCode/getTables                               | 获取数据库表            |    √     |
| /autoCode/getDB                                   | 获取所有数据库          |    √     |
| /autoCode/getColume                               | 获取所选table的所有字段 |    √     |
| /sysOperationRecord/deleteSysOperationRecordByIds | 批量删除操作历史        |          |
| /simpleUploader/upload                            | 插件版分片上传          |          |
| /simpleUploader/checkFileMd5                      | 文件完整度验证          |          |
| /simpleUploader/mergeFileMd5                      | 上传完成合并文件        |          |
