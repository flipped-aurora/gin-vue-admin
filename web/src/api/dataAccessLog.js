import service from '@/utils/request'

// 分页获取数据权限审计日志 @Router /dataAccessLog/getDataAccessLogList [post]
export const getDataAccessLogList = (data) => {
  return service({
    url: '/dataAccessLog/getDataAccessLogList',
    method: 'post',
    data
  })
}

// 批量删除数据权限审计日志 @Router /dataAccessLog/deleteDataAccessLogByIds [delete]
export const deleteDataAccessLogByIds = (data) => {
  return service({
    url: '/dataAccessLog/deleteDataAccessLogByIds',
    method: 'delete',
    data
  })
}
