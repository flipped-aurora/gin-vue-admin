package api

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type certAdvanced struct{}

var CertAdvancedApi = new(certAdvanced)

type DiscoverSubdomainsRequest struct {
	RootDomain string `json:"rootDomain" binding:"required"`
	Category   string `json:"category"`
	DeepScan   bool   `json:"deepScan"`
}

type BatchDiscoverRequest struct {
	RootDomains []string `json:"rootDomains" binding:"required"`
	Category    string   `json:"category"`
	DeepScan    bool     `json:"deepScan"`
}

type ReProbeRequest struct {
	RootDomain string `json:"rootDomain" binding:"required"`
}

type IgnoreDomainRequest struct {
	DomainID uint `json:"domainId" binding:"required"`
	Ignore   bool `json:"ignore"`
}

func (a *certAdvanced) DiscoverSubdomains(c *gin.Context) {
	var req DiscoverSubdomainsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		return
	}

	if err := service.ServiceGroupApp.SubdomainDiscovery.DiscoverAndStoreSubdomains(req.RootDomain, req.Category, req.DeepScan); err != nil {
		global.GVA_LOG.Error("子域名发现失败", zap.Error(err))
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "子域名发现成功"})
}

func (a *certAdvanced) BatchDiscoverSubdomains(c *gin.Context) {
	var req BatchDiscoverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		return
	}

	if err := service.ServiceGroupApp.SubdomainDiscovery.BatchDiscoverSubdomains(req.RootDomains, req.Category, req.DeepScan); err != nil {
		global.GVA_LOG.Error("批量子域名发现失败", zap.Error(err))
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "批量子域名发现成功"})
}

func (a *certAdvanced) GetDomainTree(c *gin.Context) {
	rootDomain := c.Query("rootDomain")
	if rootDomain == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "rootDomain 参数必填"})
		return
	}

	tree, err := service.ServiceGroupApp.SubdomainDiscovery.GetDomainTree(rootDomain)
	if err != nil {
		global.GVA_LOG.Error("获取域名树失败", zap.Error(err))
		c.JSON(500, gin.H{"code": 500, "msg": "获取域名树失败"})
		return
	}

	c.JSON(200, gin.H{"code": 0, "data": tree})
}

func (a *certAdvanced) ExportSubdomainReport(c *gin.Context) {
	rootDomain := c.Query("rootDomain")
	if rootDomain == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "rootDomain 参数必填"})
		return
	}

	filePath, fileName, err := service.ServiceGroupApp.SubdomainDiscovery.ExportSubdomainReport(rootDomain)
	if err != nil {
		global.GVA_LOG.Error("导出报告失败", zap.Error(err))
		c.JSON(500, gin.H{"code": 500, "msg": "导出报告失败"})
		return
	}
	defer os.Remove(filePath)

	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
}

func (a *certAdvanced) ReProbeDomainTree(c *gin.Context) {
	var req ReProbeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		return
	}

	var assets []model.DomainAsset
	if err := global.GVA_DB.Where("root_domain = ?", req.RootDomain).Find(&assets).Error; err != nil {
		global.GVA_LOG.Error("查询域名资产失败", zap.Error(err))
		c.JSON(500, gin.H{"code": 500, "msg": "查询域名资产失败"})
		return
	}

	for _, asset := range assets {
		if asset.IsIgnored {
			continue
		}
		if err := service.ServiceGroupApp.CertCertificate.ProbeAndUpdateCertificate(asset.Domain); err != nil {
			global.GVA_LOG.Error("探测域名失败", zap.String("domain", asset.Domain), zap.Error(err))
		}
	}

	c.JSON(200, gin.H{"code": 0, "msg": "重新探测成功"})
}

func (a *certAdvanced) IgnoreDomain(c *gin.Context) {
	var req IgnoreDomainRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		return
	}

	if err := global.GVA_DB.Model(&model.DomainAsset{}).Where("id = ?", req.DomainID).Update("is_ignored", req.Ignore).Error; err != nil {
		global.GVA_LOG.Error("更新忽略状态失败", zap.Error(err))
		c.JSON(500, gin.H{"code": 500, "msg": "更新忽略状态失败"})
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "更新成功"})
}

func (a *certAdvanced) GetDomainAssetList(c *gin.Context) {
	var pageInfo request.DomainAssetSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		return
	}

	list, total, err := service.ServiceGroupApp.DomainAssetService.GetDomainAssetList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取域名资产列表失败", zap.Error(err))
		c.JSON(500, gin.H{"code": 500, "msg": "获取域名资产列表失败"})
		return
	}

	c.JSON(200, gin.H{"code": 0, "data": gin.H{"list": list, "total": total}})
}

type BatchIdsRequest struct {
	IDs []uint `json:"ids" binding:"required"`
}

func (a *certAdvanced) BatchReprobe(c *gin.Context) {
	var req BatchIdsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		c.JSON(400, gin.H{"code": 400, "msg": "参数格式错误"})
		return
	}

	var assets []model.DomainAsset
	if err := global.GVA_DB.Where("id IN ?", req.IDs).Find(&assets).Error; err != nil {
		global.GVA_LOG.Error("查询域名记录失败", zap.Error(err))
		c.JSON(500, gin.H{"code": 500, "msg": "查询失败"})
		return
	}

	if len(assets) == 0 {
		c.JSON(200, gin.H{"code": 0, "msg": "无匹配记录"})
		return
	}

	var wg sync.WaitGroup
	errChan := make(chan string, len(assets))

	for _, asset := range assets {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			if err := service.ServiceGroupApp.CertCertificate.ProbeAndUpdateCertificate(domain); err != nil {
				global.GVA_LOG.Error("重新探测失败", zap.String("domain", domain), zap.Error(err))
				errChan <- fmt.Sprintf("%s: %v", domain, err)
			}
		}(asset.Domain)
	}

	wg.Wait()
	close(errChan)

	var errMsgs []string
	for msg := range errChan {
		errMsgs = append(errMsgs, msg)
	}

	if len(errMsgs) > 0 {
		c.JSON(200, gin.H{
			"code": 7, // 业务逻辑层面的部分失败
			"msg":  fmt.Sprintf("部分探测失败: %s", strings.Join(errMsgs, "; ")),
		})
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "批量探测完成"})
}

func (a *certAdvanced) BatchIgnore(c *gin.Context) {
	var req BatchIdsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		return
	}

	// 这里的忽略逻辑如果是针对 DomainAsset 的，需要查询 DomainAsset
	// 如果是针对 CertCertificate 的，可能需要一个新的字段。
	// 目前我们假设是针对 DomainAsset 的
	if err := global.GVA_DB.Model(&model.DomainAsset{}).Where("id IN ?", req.IDs).Update("is_ignored", true).Error; err != nil {
		global.GVA_LOG.Error("批量忽略失败", zap.Error(err))
		c.JSON(500, gin.H{"code": 500, "msg": "操作失败"})
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "批量忽略成功"})
}
