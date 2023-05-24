package addresslist

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// DepartmentSimpleListURL 获取子部门ID列表
	DepartmentSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/department/simplelist?access_token=%s&id=%d"
)

type (
	// DepartmentSimpleListResponse 获取子部门ID列表响应
	DepartmentSimpleListResponse struct {
		util.CommonError
		DepartmentID []*DepartmentID `json:"department_id"`
	}
	// DepartmentID 子部门ID
	DepartmentID struct {
		ID       int `json:"id"`
		ParentID int `json:"parentid"`
		Order    int `json:"order"`
	}
)

// DepartmentSimpleList 获取子部门ID列表
// see https://developer.work.weixin.qq.com/document/path/95350
func (r *Client) DepartmentSimpleList(departmentID int) ([]*DepartmentID, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.HTTPGet(fmt.Sprintf(DepartmentSimpleListURL, accessToken, departmentID)); err != nil {
		return nil, err
	}
	result := &DepartmentSimpleListResponse{}
	if err = util.DecodeWithError(response, result, "DepartmentSimpleList"); err != nil {
		return nil, err
	}
	return result.DepartmentID, nil
}
