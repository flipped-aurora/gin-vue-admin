import service from '@/utils/request'
// @Summary Ldap用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /ldap/login [post]
export const ldapLogin = (data) => {
  return service({
    url: '/ldap/login',
    method: 'post',
    data: data
  })
}
