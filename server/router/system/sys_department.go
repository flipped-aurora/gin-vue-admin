package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysDepartmentRouter struct{}

func (s *SysDepartmentRouter) InitSysDepartmentRouter(Router *gin.RouterGroup) {
	deptRouter := Router.Group("department").Use(middleware.OperationRecord())
	deptRouterWithoutRecord := Router.Group("department")
	{
		deptRouter.POST("createDepartment", departmentApi.CreateSysDepartment)   // 创建部门
		deptRouter.PUT("updateDepartment", departmentApi.UpdateSysDepartment)    // 更新部门
		deptRouter.DELETE("deleteDepartment", departmentApi.DeleteSysDepartment) // 删除部门
		deptRouter.POST("setDepartmentUsers", departmentApi.SetDepartmentUsers)  // 设置部门成员(反向分配)
	}
	{
		deptRouterWithoutRecord.POST("getDepartmentList", departmentApi.GetSysDepartmentList) // 获取部门树
		deptRouterWithoutRecord.GET("findDepartment", departmentApi.FindSysDepartment)        // 根据ID获取部门
		deptRouterWithoutRecord.GET("getDepartmentUsers", departmentApi.GetDepartmentUsers)   // 获取部门成员ID列表
	}
}
