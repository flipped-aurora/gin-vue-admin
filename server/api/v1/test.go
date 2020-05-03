package v1

import (
	"fmt"
    "gin-vue-admin/global/response"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    resp "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
)


// @Tags TestTest
// @Summary 创建TestTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTest true "创建TestTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /t/createTestTest [post]
func CreateTestTest(c *gin.Context) {
	var t model.TestTest
	_ = c.ShouldBindJSON(&t)
	err := service.CreateTestTest(t)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}


// @Tags TestTest
// @Summary 删除TestTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTest true "删除TestTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /t/deleteTestTest [delete]
func DeleteTestTest(c *gin.Context) {
	var t model.TestTest
	_ = c.ShouldBindJSON(&t)
	err := service.DeleteTestTest(t)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


// @Tags TestTest
// @Summary 更新TestTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTest true "更新TestTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /t/updateTestTest [put]
func UpdateTestTest(c *gin.Context) {
	var t model.TestTest
	_ = c.ShouldBindJSON(&t)
	err := service.UpdateTestTest(&t)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


// @Tags TestTest
// @Summary 用id查询TestTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTest true "用id查询TestTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /t/findTestTest [get]
func FindTestTest(c *gin.Context) {
	var t model.TestTest
	_ = c.ShouldBindQuery(&t)
	err,ret := service.GetTestTest(t.ID)
	if err != nil {
	response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData( gin.H{"ret":ret,}, c)
	}
}


// @Tags TestTest
// @Summary 分页获取TestTest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TestTest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /t/getTestTestList [get]
func GetTestTestList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetTestTestInfoList(pageInfo)
	if err != nil {
	    response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
        response.OkWithData(resp.PageResult{
                    List:     list,
                    Total:    total,
                    Page:     pageInfo.Page,
                    PageSize: pageInfo.PageSize,
                }, c)
	}
}