package utils

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// GetListTest 复刻真实请求结构的形状:匿名内嵌 request.PageInfo,
// 分页字段(Page/PageSize)靠匿名内嵌提升上来,由 Verify 的匿名递归够到。
type GetListTest struct {
	request.PageInfo
	Name string
}

func TestVerify(t *testing.T) {
	PageInfoVerify := Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "Name": {NotEmpty()}}

	// 匿名内嵌字段的 0 值仍被捕捉(递归进 request.PageInfo)
	var testInfo GetListTest
	testInfo.Name = "test"
	testInfo.Page = 0
	testInfo.PageSize = 0
	if err := Verify(testInfo, PageInfoVerify); err == nil {
		t.Error("校验失败，未能捕捉匿名内嵌字段的0值")
	}

	// 顶层字段为空正常报错
	testInfo.Name = ""
	testInfo.Page = 1
	testInfo.PageSize = 10
	if err := Verify(testInfo, PageInfoVerify); err == nil {
		t.Error("校验失败，未能正常检测name为空")
	}

	// 全部合法时通过
	testInfo.Name = "test"
	testInfo.Page = 1
	testInfo.PageSize = 10
	if err := Verify(testInfo, PageInfoVerify); err != nil {
		t.Error("校验失败，未能正常通过检测")
	}
}

// InnerID 模拟内嵌 global.GVA_MODEL 提供的 ID 字段(导出,与 GVA_MODEL 一致,
// 否则 reflect 对未导出匿名字段取 Interface() 会 panic)
type InnerID struct {
	ID uint
}

// IDCarrier 匿名内嵌 InnerID,ID 被提升为自身字段(等价于业务模型内嵌 GVA_MODEL)
type IDCarrier struct {
	InnerID
}

// assocOwner 复刻"带具名关联对象的模型":匿名内嵌 IDCarrier 提供自身 ID,
// 同时有一个具名关联字段 Assoc(其内部也有 ID)。这正是 SysOperationRecord.User
// 这类结构的形状,曾导致 IdVerify 误报。
type assocOwner struct {
	IDCarrier
	Assoc IDCarrier // 具名关联对象,自身 ID 常为 0(未 Preload)
	Name  string
}

// TestVerifyIdWithNamedAssociation 回归:Verify 不应递归进具名关联字段,
// 否则关联对象内嵌的零值 ID 会让 IdVerify 误报。
func TestVerifyIdWithNamedAssociation(t *testing.T) {
	IdVerify := Rules{"ID": {NotEmpty()}}

	// 自身 ID 合法、关联对象 ID 为 0 → 应通过(修复前会误报"ID值不能为空")
	ok := assocOwner{Name: "x"}
	ok.ID = 56 // 提升自 innerID
	if err := Verify(ok, IdVerify); err != nil {
		t.Errorf("不应因具名关联对象的零值ID而报错, got: %v", err)
	}

	// 自身 ID 为 0 → 仍应报错(匿名内嵌链上的 ID 校验必须保留)
	bad := assocOwner{Name: "x"}
	bad.Assoc.ID = 999 // 关联对象有值也不能顶替自身 ID
	if err := Verify(bad, IdVerify); err == nil {
		t.Error("自身ID为0时应报错,匿名内嵌的ID校验丢失")
	}
}
