import service from '@/utils/request'
// @Router /authority/getAuthorityList [post]
export const getAuthorityList = (data) => {
  return service({
    url: '/authority/getAuthorityList',
    method: 'post',
    data
  })
}

// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {authorityId uint} true "删除角色"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/deleteAuthority [post]
export const deleteAuthority = (data) => {
  return service({
    url: '/authority/deleteAuthority',
    method: 'post',
    data
  })
}

// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "创建角色"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/createAuthority [post]
export const createAuthority = (data) => {
  return service({
    url: '/authority/createAuthority',
    method: 'post',
    data
  })
}

// @Tags authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "拷贝角色"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]
export const copyAuthority = (data) => {
  return service({
    url: '/authority/copyAuthority',
    method: 'post',
    data
  })
}

// @Summary 设置角色数据范围(数据权限)
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {authorityId:uint,dataScope:int} true "设置角色数据范围"
// @Router /authority/setDataScope [post]
export const setDataScope = (data) => {
  return service({
    url: '/authority/setDataScope',
    method: 'post',
    data
  })
}

// @Summary 获取角色自定义部门集(数据权限第5档)
// @Router /authority/getDataScopeDepts [get]
export const getDataScopeDepts = (authorityId) => {
  return service({
    url: '/authority/getDataScopeDepts',
    method: 'get',
    params: { authorityId }
  })
}

// @Summary 修改角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "修改角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/updateAuthority [put]
export const updateAuthority = (data) => {
  return service({
    url: '/authority/updateAuthority',
    method: 'put',
    data
  })
}

/**
 * 获取拥有指定角色的用户ID列表
 * @param {number} authorityId 角色ID
 * @returns {Promise<number[]>} 用户ID数组
 */
export const getUsersByAuthorityId = (authorityId) => {
  return service({
    url: '/authority/getUsersByAuthority',
    method: 'get',
    params: { authorityId }
  })
}

/**
 * 全量覆盖某角色关联的用户列表
 * @param {Object} data
 * @param {number} data.authorityId 角色ID
 * @param {number[]} data.userIds 用户ID列表
 * @returns {Promise}
 */
export const setRoleUsers = (data) => {
  return service({
    url: '/authority/setRoleUsers',
    method: 'post',
    data
  })
}
