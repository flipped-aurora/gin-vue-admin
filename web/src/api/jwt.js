import { createService } from '@/utils/request'
// @Tags jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
export const jsonInBlacklist = () => {
  const service = createService()
  return service({
    url: '/jwt/jsonInBlacklist',
    method: 'post'
  })
}
