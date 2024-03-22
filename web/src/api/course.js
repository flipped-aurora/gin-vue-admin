import service from '@/utils/request'

// @Tags Course
// @Summary 创建Course
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Course true "创建Course"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /course/createCourse [post]
export const createCourse = (data) => {
  return service({
    url: '/course/createCourse',
    method: 'post',
    data
  })
}

// @Tags Course
// @Summary 删除Course
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Course true "删除Course"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /course/deleteCourse [delete]
export const deleteCourse = (data) => {
  return service({
    url: '/course/deleteCourse',
    method: 'delete',
    data
  })
}

// @Tags Course
// @Summary 删除Course
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Course"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /course/deleteCourse [delete]
export const deleteCourseByIds = (data) => {
  return service({
    url: '/course/deleteCourseByIds',
    method: 'delete',
    data
  })
}

// @Tags Course
// @Summary 更新Course
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Course true "更新Course"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /course/updateCourse [put]
export const updateCourse = (data) => {
  return service({
    url: '/course/updateCourse',
    method: 'put',
    data
  })
}

// @Tags Course
// @Summary 用id查询Course
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Course true "用id查询Course"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /course/findCourse [get]
export const findCourse = (params) => {
  return service({
    url: '/course/findCourse',
    method: 'get',
    params
  })
}

// @Tags Course
// @Summary 分页获取Course列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Course列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /course/getCourseList [get]
export const getCourseList = (params) => {
  return service({
    url: '/course/getCourseList',
    method: 'get',
    params
  })
}


export const getCourseIdName = () => {
  return service({
    url: '/course/getCourseIdName',
    method: 'get',
  })
}
