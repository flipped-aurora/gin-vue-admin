import service from '@/utils/request'

// @Summary 获取指定月份存在日志的日期
// @Router /logViewer/dates [get]
export const getLogDates = (params) => {
  return service({
    url: '/logViewer/dates',
    method: 'get',
    params,
    donNotShowLoading: true
  })
}

// @Summary 获取指定日期的日志文件
// @Router /logViewer/files [get]
export const getLogFiles = (params) => {
  return service({
    url: '/logViewer/files',
    method: 'get',
    params,
    donNotShowLoading: true
  })
}

// @Summary 分块读取日志文件内容
// @Router /logViewer/content [get]
export const getLogContent = (params) => {
  return service({
    url: '/logViewer/content',
    method: 'get',
    params,
    donNotShowLoading: true
  })
}
