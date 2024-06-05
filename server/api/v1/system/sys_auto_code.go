package system

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	sysReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeApi struct{}

// PreviewTemp
// @Tags      AutoCode
// @Summary   预览创建后的代码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.AutoCodeStruct                                      true  "预览创建代码"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "预览创建后的代码"
// @Router    /autoCode/preview [post]
func (autoApi *AutoCodeApi) PreviewTemp(c *gin.Context) {
	var a system.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Pretreatment() // 处理go关键字
	a.PackageT = utils.FirstUpper(a.Package)
	autoCode, err := autoCodeService.PreviewTemp(a)
	if err != nil {
		global.GVA_LOG.Error("预览失败!", zap.Error(err))
		response.FailWithMessage("预览失败", c)
	} else {
		response.OkWithDetailed(gin.H{"autoCode": autoCode}, "预览成功", c)
	}
}

// CreateTemp
// @Tags      AutoCode
// @Summary   自动代码模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.AutoCodeStruct  true  "创建自动代码"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/createTemp [post]
func (autoApi *AutoCodeApi) CreateTemp(c *gin.Context) {
	var a system.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Pretreatment()
	var apiIds []uint
	var menuId uint
	if a.AutoCreateApiToSql {
		if ids, err := autoCodeService.AutoCreateApi(&a); err != nil {
			global.GVA_LOG.Error("自动化创建失败!请自行清空垃圾数据!", zap.Error(err))
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
			return
		} else {
			apiIds = ids
		}
	}
	if a.AutoCreateMenuToSql {
		if id, err := autoCodeService.AutoCreateMenu(&a); err != nil {
			global.GVA_LOG.Error("自动化创建失败!请自行清空垃圾数据!", zap.Error(err))
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
		} else {
			menuId = id
		}
	}
	a.PackageT = utils.FirstUpper(a.Package)
	err := autoCodeService.CreateTemp(a, menuId, apiIds...)
	if err != nil {
		c.Writer.Header().Add("success", "false")
		c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
		return
	}
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Header().Add("success", "true")
	if a.AutoKeepCode {
		var records []system.RecordsDeleteCode
		result := global.GVA_DB.Table("records").Select("path", "file", "update_time").Find(&records)
		// 从 records 表中获取被删除的代码文件的路径和文件名
		if result.Error != nil {
			global.GVA_LOG.Error("无任何删除记录!", zap.Error(result.Error))
			response.FailWithMessage("无任何删除记录", c)
			return
		}

		for _, record := range records {
			srcFile := record.Path
			file := record.File

			// destFile 是新创建的文件的路径
			destFile := filepath.Join(global.GVA_CONFIG.AutoCode.Root, file)
			// 检查新文件的路径是否存在
			if _, err := os.Stat(destFile); err == nil {
				if err := extractAndAppendCodeBlocks(srcFile, destFile); err != nil {
					global.GVA_LOG.Error("提取代码块失败!", zap.Error(err))
					response.FailWithMessage("提取代码块失败", c)
					return
				}
				// 删除数据库中的记录
				result := global.GVA_DB.Where("path = ? AND file = ?", srcFile, file).Delete(&system.RecordsDeleteCode{})
				if result.Error != nil {
					global.GVA_LOG.Error("删除记录失败!", zap.Error(result.Error))
					response.FailWithMessage("删除记录失败", c)
					return
				}
			}
		}
	}
}

// 提取rm_file(删除文件存放)代码文件中的标记代码段,添加到目标文件末尾,如果目标文件不存在则自动创建
func extractAndAppendCodeBlocks(srcFile, destFile string) error {

	source, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer source.Close()

	// 检查目标文件是否存在，如果不存在则创建
	dest, err := os.OpenFile(destFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer dest.Close()

	scanner := bufio.NewScanner(source)
	writer := bufio.NewWriter(dest)
	defer writer.Flush()

	keepWriting := false
	nestCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		if isStartTag(trimmedLine) {
			nestCount++
			if nestCount == 1 {
				keepWriting = true
			}
		}

		if keepWriting {
			_, err = writer.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}

		if isEndTag(trimmedLine) {
			if nestCount > 0 {
				nestCount--
			}
			if nestCount == 0 {
				keepWriting = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if nestCount != 0 {
		return errors.New("发现未匹配的标签")
	}

	return nil
}

// isStartTag 检查一行是否包含开始标签。
func isStartTag(line string) bool {
	match, _ := regexp.MatchString(`^\s*//\s*@gvastartkeep\s*$`, line)
	return match
}

// isEndTag 检查一行是否包含结束标签。
func isEndTag(line string) bool {
	match, _ := regexp.MatchString(`^\s*//\s*@gvaendkeep\s*$`, line)
	return match
}

// GetDB
// @Tags      AutoCode
// @Summary   获取当前所有数据库
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前所有数据库"
// @Router    /autoCode/getDatabase [get]
func (autoApi *AutoCodeApi) GetDB(c *gin.Context) {
	businessDB := c.Query("businessDB")
	dbs, err := autoCodeService.Database(businessDB).GetDB(businessDB)
	var dbList []map[string]interface{}
	for _, db := range global.GVA_CONFIG.DBList {
		var item = make(map[string]interface{})
		item["aliasName"] = db.AliasName
		item["dbName"] = db.Dbname
		item["disable"] = db.Disable
		item["dbtype"] = db.Type
		dbList = append(dbList, item)
	}
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs, "dbList": dbList}, "获取成功", c)
	}
}

// GetTables
// @Tags      AutoCode
// @Summary   获取当前数据库所有表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前数据库所有表"
// @Router    /autoCode/getTables [get]
func (autoApi *AutoCodeApi) GetTables(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	businessDB := c.Query("businessDB")
	tables, err := autoCodeService.Database(businessDB).GetTables(businessDB, dbName)
	if err != nil {
		global.GVA_LOG.Error("查询table失败!", zap.Error(err))
		response.FailWithMessage("查询table失败", c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, "获取成功", c)
	}
}

// GetColumn
// @Tags      AutoCode
// @Summary   获取当前表所有字段
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前表所有字段"
// @Router    /autoCode/getColumn [get]
func (autoApi *AutoCodeApi) GetColumn(c *gin.Context) {
	businessDB := c.Query("businessDB")
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	tableName := c.Query("tableName")
	columns, err := autoCodeService.Database(businessDB).GetColumn(businessDB, tableName, dbName)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, "获取成功", c)
	}
}

// CreatePackage
// @Tags      AutoCode
// @Summary   创建package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/createPackage [post]
func (autoApi *AutoCodeApi) CreatePackage(c *gin.Context) {
	var a system.SysAutoCode
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.AutoPackageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// PackageName可能导致路径穿越的问题 / 和 \ 都要防止
	if strings.Contains(a.PackageName, "\\") || strings.Contains(a.PackageName, "/") || strings.Contains(a.PackageName, "..") {
		response.FailWithMessage("包名不合法", c)
		return
	}

	err := autoCodeService.CreateAutoCode(&a)
	if err != nil {

		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdatePackage
// @Tags      AutoCode
// @Summary   更新package展示名字/描述
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "更新package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "更新package成功"
// @Router    /autoCode/updatePackageDetail [post]
func (autoApi *AutoCodeApi) UpdatePackageDetail(c *gin.Context) {
	var a sysReq.SysAutoCode
	var autoCode system.SysAutoCode
	_ = c.ShouldBindJSON(&a)
	if err := global.GVA_DB.Where("id = ?", a.ID).First(&autoCode).Error; err != nil {
		response.FailWithMessage("获取package失败", c)
		return
	}
	// 更新展示名字/描述
	autoCode.Label = a.Label
	autoCode.Desc = a.Desc
	if err := global.GVA_DB.Save(&autoCode).Error; err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetPackage
// @Tags      AutoCode
// @Summary   获取package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/getPackage [post]
func (autoApi *AutoCodeApi) GetPackage(c *gin.Context) {
	pkgs, err := autoCodeService.GetPackage()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"pkgs": pkgs}, "获取成功", c)
	}
}

// auto_code_api.go

// GetPackageByID
// @Tags      AutoCode
// @Summary   根据ID获取package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  body      int                                                true  "package ID"
// @Success   200  {object}  response.Response{data=system.SysAutoCode,msg=string}  "根据ID获取package成功"
// @Router    /autoCode/getPackageByID/ POST
func (autoApi *AutoCodeApi) GetPackageById(c *gin.Context) {
	var pkgId request.GetById
	err := c.ShouldBindJSON(&pkgId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var autoCode system.SysAutoCode
	fmt.Printf("pkgId%+v", pkgId)
	if err := global.GVA_DB.First(&autoCode, pkgId).Error; err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"pkg": autoCode}, "获取成功", c)
	}
}

// DelPackage
// @Tags      AutoCode
// @Summary   删除package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "删除package成功"
// @Router    /autoCode/delPackage [post]
func (autoApi *AutoCodeApi) DelPackage(c *gin.Context) {
	var a system.SysAutoCode
	_ = c.ShouldBindJSON(&a)
	err := autoCodeService.DelPackage(a)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// AutoPlug
// @Tags      AutoCode
// @Summary   创建插件模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建插件模板"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "创建插件模板成功"
// @Router    /autoCode/createPlug [post]
func (autoApi *AutoCodeApi) AutoPlug(c *gin.Context) {
	var a system.AutoPlugReq
	err := c.ShouldBindJSON(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if strings.Contains(a.PlugName, "\\") || strings.Contains(a.PlugName, "/") || strings.Contains(a.PlugName, "..") {
		response.FailWithMessage("插件名称不合法", c)
		return
	}

	a.Snake = strings.ToLower(a.PlugName)
	a.NeedModel = a.HasRequest || a.HasResponse
	err = autoCodeService.CreatePlug(a)
	if err != nil {
		global.GVA_LOG.Error("预览失败!", zap.Error(err))
		response.FailWithMessage("预览失败", c)
		return
	}
	response.Ok(c)
}

// InstallPlugin
// @Tags      AutoCode
// @Summary   安装插件
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     plug  formData  file                                              true  "this is a test file"
// @Success   200   {object}  response.Response{data=[]interface{},msg=string}  "安装插件成功"
// @Router    /autoCode/installPlugin [post]
func (autoApi *AutoCodeApi) InstallPlugin(c *gin.Context) {
	header, err := c.FormFile("plug")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	web, server, err := autoCodeService.InstallPlugin(header)
	webStr := "web插件安装成功"
	serverStr := "server插件安装成功"
	if web == -1 {
		webStr = "web端插件未成功安装，请按照文档自行解压安装，如果为纯后端插件请忽略此条提示"
	}
	if server == -1 {
		serverStr = "server端插件未成功安装，请按照文档自行解压安装，如果为纯前端插件请忽略此条提示"
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData([]interface{}{
		gin.H{
			"code": web,
			"msg":  webStr,
		},
		gin.H{
			"code": server,
			"msg":  serverStr,
		}}, c)
}

// PubPlug
// @Tags      AutoCode
// @Summary   打包插件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "打包插件"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "打包插件成功"
// @Router    /autoCode/pubPlug [get]
func (autoApi *AutoCodeApi) PubPlug(c *gin.Context) {
	plugName := c.Query("plugName")
	zipPath, err := autoCodeService.PubPlug(plugName)
	if err != nil {
		global.GVA_LOG.Error("打包失败!", zap.Error(err))
		response.FailWithMessage("打包失败"+err.Error(), c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("打包成功,文件路径为:%s", zipPath), c)
}
