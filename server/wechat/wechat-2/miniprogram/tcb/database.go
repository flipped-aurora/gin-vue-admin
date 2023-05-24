package tcb

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 数据库导入
	databaseMigrateImportURL = "https://api.weixin.qq.com/tcb/databasemigrateimport"
	// 数据库导出
	databaseMigrateExportURL = "https://api.weixin.qq.com/tcb/databasemigrateexport"
	// 数据库迁移状态查询
	databaseMigrateQueryInfoURL = "https://api.weixin.qq.com/tcb/databasemigratequeryinfo"
	// 变更数据库索引
	updateIndexURL = "https://api.weixin.qq.com/tcb/updateindex"
	// 新增集合
	databaseCollectionAddURL = "https://api.weixin.qq.com/tcb/databasecollectionadd"
	// 删除集合
	databaseCollectionDeleteURL = "https://api.weixin.qq.com/tcb/databasecollectiondelete"
	// 获取特定云环境下集合信息
	databaseCollectionGetURL = "https://api.weixin.qq.com/tcb/databasecollectionget"
	// 数据库插入记录
	databaseAddURL = "https://api.weixin.qq.com/tcb/databaseadd"
	// 数据库删除记录
	databaseDeleteURL = "https://api.weixin.qq.com/tcb/databasedelete"
	// 数据库更新记录
	databaseUpdateURL = "https://api.weixin.qq.com/tcb/databaseupdate"
	// 数据库查询记录
	databaseQueryURL = "https://api.weixin.qq.com/tcb/databasequery"
	// 统计集合记录数或统计查询语句对应的结果记录数
	databaseCountURL = "https://api.weixin.qq.com/tcb/databasecount"

	// ConflictModeInster 冲突处理模式 插入
	ConflictModeInster ConflictMode = 1
	// ConflictModeUpsert 冲突处理模式 更新
	ConflictModeUpsert ConflictMode = 2

	// FileTypeJSON 的合法值 json
	FileTypeJSON FileType = 1
	// FileTypeCsv 的合法值 csv
	FileTypeCsv FileType = 2
)

// ConflictMode 冲突处理模式
type ConflictMode int

// FileType 文件上传和导出的允许文件类型
type FileType int

// ValidDirections 合法的direction值
var ValidDirections = []string{"1", "-1", "2dsphere"}

// DatabaseMigrateExportReq 数据库出 请求参数
type DatabaseMigrateExportReq struct {
	Env      string   `json:"env,omitempty"`       // 云环境ID
	FilePath string   `json:"file_path,omitempty"` // 导出文件路径(导入文件需先上传到同环境的存储中，可使用开发者工具或 HTTP API的上传文件 API上传）
	FileType FileType `json:"file_type,omitempty"` // 导出文件类型，文件格式参考数据库导入指引中的文件格式部分  1:json 2:csv
	Query    string   `json:"query,omitempty"`     // 导出条件
}

// DatabaseMigrateExportRes 数据库导出 返回结果
type DatabaseMigrateExportRes struct {
	util.CommonError
	JobID int64 `json:"job_id"` // 导出任务ID，可使用数据库迁移进度查询 API 查询导入进度及结果
}

// DatabaseMigrateImportReq 数据库导入 请求参数
type DatabaseMigrateImportReq struct {
	Env            string       `json:"env,omitempty"`             // 云环境ID
	CollectionName string       `json:"collection_name,omitempty"` // 集合名称
	FilePath       string       `json:"file_path,omitempty"`       // 导出文件路径（文件会导出到同环境的云存储中，可使用获取下载链接 API 获取下载链接）
	FileType       FileType     `json:"file_type,omitempty"`       // 导入文件类型，文件格式参考数据库导入指引中的文件格式部分  1:json 2:csv
	StopOnError    bool         `json:"stop_on_error,omitempty"`   // 是否在遇到错误时停止导入
	ConflictMode   ConflictMode `json:"conflict_mode,omitempty"`   // 冲突处理模式  1:inster 2:UPSERT
}

// DatabaseMigrateImportRes 数据库导入 返回结果
type DatabaseMigrateImportRes struct {
	util.CommonError
	JobID int64 `json:"job_id"` // 导入任务ID，可使用数据库迁移进度查询 API 查询导入进度及结果
}

// DatabaseMigrateQueryInfoRes 数据库迁移状态查询
type DatabaseMigrateQueryInfoRes struct {
	util.CommonError
	Status        string `json:"status"`         // 导出状态
	RecordSuccess int64  `json:"record_success"` // 导出成功记录数
	RecordFail    int64  `json:"record_fail"`    // 导出失败记录数
	ErrMsg        string `json:"err_msg"`        // 导出错误信息
	FileURL       string `json:"file_url"`       // 导出文件下载地址
}

// UpdateIndexReq 变更数据库索引 请求参数
type UpdateIndexReq struct {
	Env            string        `json:"env,omitempty"`             // 云环境ID
	CollectionName string        `json:"collection_name,omitempty"` // 集合名称
	CreateIndexes  []CreateIndex `json:"create_indexes,omitempty"`  // 新增索引
	DropIndexes    []DropIndex   `json:"drop_indexes,omitempty"`    // 删除索引
}

// CreateIndex 新增索引
type CreateIndex struct {
	Name   string           `json:"name,omitempty"`   // 索引名
	Unique bool             `json:"unique,omitempty"` // 是否唯一
	Keys   []CreateIndexKey `json:"keys,omitempty"`   // 索引字段
}

// CreateIndexKey create index key
type CreateIndexKey struct {
	Name      string `json:"name,omitempty"`      // 字段名
	Direction string `json:"direction,omitempty"` // 字段排序
}

// DropIndex 删除索引
type DropIndex struct {
	Name string `json:"name,omitempty"`
}

// DatabaseCollectionReq 新增/删除集合请求参数
type DatabaseCollectionReq struct {
	Env            string `json:"env,omitempty"`             // 云环境ID
	CollectionName string `json:"collection_name,omitempty"` // 集合名称
}

// DatabaseCollectionGetReq 获取特定云环境下集合信息请求
type DatabaseCollectionGetReq struct {
	Env    string `json:"env,omitempty"`    // 云环境ID
	Limit  int64  `json:"limit,omitempty"`  // 获取数量限制
	Offset int64  `json:"offset,omitempty"` // 偏移量
}

// DatabaseCollectionGetRes 获取特定云环境下集合信息结果
type DatabaseCollectionGetRes struct {
	util.CommonError
	Pager struct {
		Limit  int64 `json:"limit"`  // 单次查询限制
		Offset int64 `json:"offset"` // 偏移量
		Total  int64 `json:"total"`  // 符合查询条件的记录总数
	} `json:"pager"`
	Collections []struct {
		Name       string `json:"name"`        // 集合名
		Count      int64  `json:"count"`       // 表中文档数量
		Size       int64  `json:"size"`        // 表的大小（即表中文档总大小），单位：字节
		IndexCount int64  `json:"index_count"` // 索引数量
		IndexSize  int64  `json:"index_size"`  // 索引占用大小，单位：字节
	} `json:"collections"`
}

// DatabaseReq 数据库插入/删除/更新/查询/统计记录请求参数
type DatabaseReq struct {
	Env   string `json:"env,omitempty"`   // 云环境ID
	Query string `json:"query,omitempty"` // 数据库操作语句
}

// DatabaseAddRes 数据库插入记录返回结果
type DatabaseAddRes struct {
	util.CommonError
	IDList []string `json:"id_list"` // 插入成功的数据集合主键_id。
}

// DatabaseDeleteRes 数据库删除记录返回结果
type DatabaseDeleteRes struct {
	util.CommonError
	Deleted int64 `json:"deleted"` // 删除记录数量
}

// DatabaseUpdateRes 数据库更新记录返回结果
type DatabaseUpdateRes struct {
	util.CommonError
	Matched  int64  `json:"matched"`  // 更新条件匹配到的结果数
	Modified int64  `json:"modified"` // 修改的记录数，注意：使用set操作新插入的数据不计入修改数目
	ID       string `json:"id"`
}

// DatabaseQueryRes 数据库查询记录 返回结果
type DatabaseQueryRes struct {
	util.CommonError
	Pager struct {
		Limit  int64 `json:"limit"`  // 单次查询限制
		Offset int64 `json:"offset"` // 偏移量
		Total  int64 `json:"total"`  // 符合查询条件的记录总数
	} `json:"pager"`
	Data []string `json:"data"`
}

// DatabaseCountRes 统计集合记录数或统计查询语句对应的结果记录数 返回结果
type DatabaseCountRes struct {
	util.CommonError
	Count int64 `json:"count"` // 记录数量
}

// DatabaseMigrateImport 数据库导入
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseMigrateImport.html
func (tcb *Tcb) DatabaseMigrateImport(req *DatabaseMigrateImportReq) (*DatabaseMigrateImportRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseMigrateImportURL, accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}
	databaseMigrateImportRes := &DatabaseMigrateImportRes{}
	err = util.DecodeWithError(response, databaseMigrateImportRes, "DatabaseMigrateImport")
	return databaseMigrateImportRes, err
}

// DatabaseMigrateExport 数据库导出
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseMigrateExport.html
func (tcb *Tcb) DatabaseMigrateExport(req *DatabaseMigrateExportReq) (*DatabaseMigrateExportRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseMigrateExportURL, accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}
	databaseMigrateExportRes := &DatabaseMigrateExportRes{}
	err = util.DecodeWithError(response, databaseMigrateExportRes, "DatabaseMigrateExport")
	return databaseMigrateExportRes, err
}

// DatabaseMigrateQueryInfo 数据库迁移状态查询
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseMigrateQueryInfo.html
func (tcb *Tcb) DatabaseMigrateQueryInfo(env string, jobID int64) (*DatabaseMigrateQueryInfoRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseMigrateQueryInfoURL, accessToken)
	response, err := util.PostJSON(uri, map[string]interface{}{
		"env":    env,
		"job_id": jobID,
	})
	if err != nil {
		return nil, err
	}
	databaseMigrateQueryInfoRes := &DatabaseMigrateQueryInfoRes{}
	err = util.DecodeWithError(response, databaseMigrateQueryInfoRes, "DatabaseMigrateQueryInfo")
	return databaseMigrateQueryInfoRes, err
}

// UpdateIndex 变更数据库索引
// https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/updateIndex.html
func (tcb *Tcb) UpdateIndex(req *UpdateIndexReq) error {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("%s?access_token=%s", updateIndexURL, accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "UpdateIndex")
}

// DatabaseCollectionAdd 新增集合
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseCollectionAdd.html
func (tcb *Tcb) DatabaseCollectionAdd(env, collectionName string) error {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseCollectionAddURL, accessToken)
	response, err := util.PostJSON(uri, &DatabaseCollectionReq{
		Env:            env,
		CollectionName: collectionName,
	})
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "DatabaseCollectionAdd")
}

// DatabaseCollectionDelete 删除集合
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseCollectionDelete.html
func (tcb *Tcb) DatabaseCollectionDelete(env, collectionName string) error {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseCollectionDeleteURL, accessToken)
	response, err := util.PostJSON(uri, &DatabaseCollectionReq{
		Env:            env,
		CollectionName: collectionName,
	})
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "DatabaseCollectionDelete")
}

// DatabaseCollectionGet 获取特定云环境下集合信息
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseCollectionGet.html
func (tcb *Tcb) DatabaseCollectionGet(env string, limit, offset int64) (*DatabaseCollectionGetRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseCollectionGetURL, accessToken)
	response, err := util.PostJSON(uri, &DatabaseCollectionGetReq{
		Env:    env,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	databaseCollectionGetRes := &DatabaseCollectionGetRes{}
	err = util.DecodeWithError(response, databaseCollectionGetRes, "DatabaseCollectionGet")
	return databaseCollectionGetRes, err
}

// DatabaseAdd 数据库插入记录
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseAdd.html
func (tcb *Tcb) DatabaseAdd(env, query string) (*DatabaseAddRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseAddURL, accessToken)
	response, err := util.PostJSON(uri, &DatabaseReq{
		Env:   env,
		Query: query,
	})
	if err != nil {
		return nil, err
	}
	databaseAddRes := &DatabaseAddRes{}
	err = util.DecodeWithError(response, databaseAddRes, "DatabaseAdd")
	return databaseAddRes, err
}

// DatabaseDelete 数据库插入记录
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseDelete.html
func (tcb *Tcb) DatabaseDelete(env, query string) (*DatabaseDeleteRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseDeleteURL, accessToken)
	response, err := util.PostJSON(uri, &DatabaseReq{
		Env:   env,
		Query: query,
	})
	if err != nil {
		return nil, err
	}
	databaseDeleteRes := &DatabaseDeleteRes{}
	err = util.DecodeWithError(response, databaseDeleteRes, "DatabaseDelete")
	return databaseDeleteRes, err
}

// DatabaseUpdate 数据库插入记录
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseUpdate.html
func (tcb *Tcb) DatabaseUpdate(env, query string) (*DatabaseUpdateRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseUpdateURL, accessToken)
	response, err := util.PostJSON(uri, &DatabaseReq{
		Env:   env,
		Query: query,
	})
	if err != nil {
		return nil, err
	}
	databaseUpdateRes := &DatabaseUpdateRes{}
	err = util.DecodeWithError(response, databaseUpdateRes, "DatabaseUpdate")
	return databaseUpdateRes, err
}

// DatabaseQuery 数据库查询记录
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseQuery.html
func (tcb *Tcb) DatabaseQuery(env, query string) (*DatabaseQueryRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseQueryURL, accessToken)
	response, err := util.PostJSON(uri, &DatabaseReq{
		Env:   env,
		Query: query,
	})
	if err != nil {
		return nil, err
	}
	databaseQueryRes := &DatabaseQueryRes{}
	err = util.DecodeWithError(response, databaseQueryRes, "DatabaseQuery")
	return databaseQueryRes, err
}

// DatabaseCount 统计集合记录数或统计查询语句对应的结果记录数
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/database/databaseCount.html
func (tcb *Tcb) DatabaseCount(env, query string) (*DatabaseCountRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", databaseCountURL, accessToken)
	response, err := util.PostJSON(uri, &DatabaseReq{
		Env:   env,
		Query: query,
	})
	if err != nil {
		return nil, err
	}
	databaseCountRes := &DatabaseCountRes{}
	err = util.DecodeWithError(response, databaseCountRes, "DatabaseCount")
	return databaseCountRes, err
}
