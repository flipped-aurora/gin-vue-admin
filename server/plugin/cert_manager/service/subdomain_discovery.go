package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

type subdomainDiscovery struct{}

type CTLogResponse struct {
	ID        int    `json:"id"`
	NameValue string `json:"name_value"`
}

func (s *subdomainDiscovery) DiscoverSubdomainsViaCTLog(rootDomain string) ([]string, error) {
	apiURL := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", rootDomain)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("CT Log API 请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CT Log API 返回状态码: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var ctResponses []CTLogResponse
	if err := json.Unmarshal(body, &ctResponses); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败: %v", err)
	}

	subdomainSet := make(map[string]bool)
	for _, item := range ctResponses {
		names := strings.Split(item.NameValue, "\n")
		for _, name := range names {
			name = strings.TrimSpace(name)
			if name != "" && strings.HasSuffix(name, rootDomain) {
				subdomainSet[name] = true
			}
		}
	}

	subdomains := make([]string, 0, len(subdomainSet))
	for subdomain := range subdomainSet {
		subdomains = append(subdomains, subdomain)
	}

	return subdomains, nil
}

// CommonSubdomains 常用子域名字典
var CommonSubdomains = []string{
	"www", "mail", "remote", "blog", "webmail", "server", "ns1", "ns2", "smtp", "vpn",
	"m", "shop", "ftp", "mail2", "test", "portal", "host", "support", "dev", "web",
	"api", "data", "client", "manage", "admin", "crm", "erp", "cloud", "static", "img",
}

func (s *subdomainDiscovery) BruteForceSubdomains(rootDomain string) []string {
	var subdomains []string
	var mu sync.Mutex
	var wg sync.WaitGroup
	sem := make(chan struct{}, 20) // 控制并发数

	for _, sub := range CommonSubdomains {
		wg.Add(1)
		go func(sub string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			domain := fmt.Sprintf("%s.%s", sub, rootDomain)
			_, err := net.LookupHost(domain)
			if err == nil {
				mu.Lock()
				subdomains = append(subdomains, domain)
				mu.Unlock()
			}
		}(sub)
	}
	wg.Wait()
	return subdomains
}

func (s *subdomainDiscovery) ResolveIP(domain string) string {
	ips, err := net.LookupHost(domain)
	if err != nil || len(ips) == 0 {
		return ""
	}
	return ips[0]
}

func (s *subdomainDiscovery) DiscoverAndStoreSubdomains(rootDomain string, category string, deepScan bool) error {
	// 1. 获取所有子域名
	ctSubdomains, _ := s.DiscoverSubdomainsViaCTLog(rootDomain)
	subdomainSet := make(map[string]string)
	for _, d := range ctSubdomains {
		subdomainSet[d] = "CT Log"
	}

	if deepScan {
		bruteSubdomains := s.BruteForceSubdomains(rootDomain)
		for _, d := range bruteSubdomains {
			if _, ok := subdomainSet[d]; !ok {
				subdomainSet[d] = "Brute Force"
			}
		}
	}
	subdomainSet[rootDomain] = "Input"

	// 2. 将域名按长度排序（由短到长），确保父域名先处理
	allDomains := make([]string, 0, len(subdomainSet))
	for d := range subdomainSet {
		allDomains = append(allDomains, d)
	}
	sort.Slice(allDomains, func(i, j int) bool {
		return len(allDomains[i]) < len(allDomains[j])
	})

	// 3. 并发解析 IP
	type result struct {
		domain string
		ip     string
	}
	resChan := make(chan result, len(allDomains))
	var wg sync.WaitGroup
	sem := make(chan struct{}, 50)

	for _, d := range allDomains {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			resChan <- result{domain: domain, ip: s.ResolveIP(domain)}
		}(d)
	}
	wg.Wait()
	close(resChan)

	ipMap := make(map[string]string)
	for r := range resChan {
		ipMap[r.domain] = r.ip
	}

	// 4. 按排序后的顺序入库
	for _, subdomain := range allDomains {
		source := subdomainSet[subdomain]
		ip := ipMap[subdomain]

		var existingAsset model.DomainAsset
		err := global.GVA_DB.Where("domain = ?", subdomain).First(&existingAsset).Error

		if err != nil {
			// 新增记录
			newAsset := model.DomainAsset{
				Domain:       subdomain,
				RootDomain:   rootDomain,
				Category:     category,
				DomainStatus: 1,
				CertStatus:   1,
				IP:           ip,
				Source:       source,
			}
			if ip == "" {
				newAsset.DomainStatus = 3
			}

			// 关联父域名
			if subdomain != rootDomain {
				parentDomain := extractParentDomain(subdomain, rootDomain)
				if parentDomain != "" {
					var parentAsset model.DomainAsset
					global.GVA_DB.Where("domain = ?", parentDomain).First(&parentAsset)
					if parentAsset.ID > 0 {
						pID := parentAsset.ID
						newAsset.ParentID = &pID
					}
				}
			}
			global.GVA_DB.Create(&newAsset)
		} else {
			// 更新现有记录
			updates := make(map[string]interface{})
			updates["root_domain"] = rootDomain
			updates["category"] = category
			if existingAsset.IP != ip {
				updates["ip"] = ip
				if ip == "" {
					updates["domain_status"] = 3
				} else {
					updates["domain_status"] = 1
				}
			}
			if existingAsset.Source == "" {
				updates["source"] = source
			}
			// 尝试修复 ParentID (如果之前没关联上)
			if existingAsset.ParentID == nil && subdomain != rootDomain {
				parentDomain := extractParentDomain(subdomain, rootDomain)
				if parentDomain != "" {
					var parentAsset model.DomainAsset
					global.GVA_DB.Where("domain = ?", parentDomain).First(&parentAsset)
					if parentAsset.ID > 0 {
						pID := parentAsset.ID
						updates["parent_id"] = &pID
					}
				}
			}
			if len(updates) > 0 {
				global.GVA_DB.Model(&existingAsset).Updates(updates)
			}
		}
	}

	return nil
}

func extractParentDomain(subdomain, rootDomain string) string {
	if subdomain == rootDomain {
		return ""
	}

	if !strings.HasSuffix(subdomain, rootDomain) {
		return ""
	}

	prefix := strings.TrimSuffix(subdomain, "."+rootDomain)
	parts := strings.Split(prefix, ".")
	if len(parts) > 1 {
		return strings.Join(parts[1:], ".") + "." + rootDomain
	}

	return rootDomain
}

func (s *subdomainDiscovery) ExportSubdomainReport(rootDomain string) (string, string, error) {
	var assets []model.DomainAsset
	err := global.GVA_DB.Where("root_domain = ?", rootDomain).Order("id asc").Find(&assets).Error
	if err != nil {
		return "", "", err
	}

	f := excelize.NewFile()
	sheetName := "域名资产报告"
	index, _ := f.NewSheet(sheetName)
	f.DeleteSheet("Sheet1")

	// 设置表头
	headers := []string{"ID", "域名", "根域名", "父域名ID", "项目分类", "解析IP", "发现来源", "域名状态", "证书状态", "最后探测时间"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, h)
	}

	// 填充数据
	for i, asset := range assets {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), asset.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), asset.Domain)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), asset.RootDomain)
		if asset.ParentID != nil {
			f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), *asset.ParentID)
		}
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), asset.Category)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), asset.IP)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), asset.Source)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), map[int]string{1: "正常", 3: "异常"}[asset.DomainStatus])
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), map[int]string{1: "正常", 3: "异常"}[asset.CertStatus])
		if asset.LastProbedAt != nil {
			f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), asset.LastProbedAt.Format("2006-01-02 15:04:05"))
		}
	}

	f.SetActiveSheet(index)
	fileName := fmt.Sprintf("DomainReport_%s_%s.xlsx", rootDomain, time.Now().Format("20060102150405"))
	filePath := "./" + fileName // 临时存放在当前目录，API 层读取后删除

	if err := f.SaveAs(filePath); err != nil {
		return "", "", err
	}

	return filePath, fileName, nil
}

func (s *subdomainDiscovery) BatchDiscoverSubdomains(rootDomains []string, category string, deepScan bool) error {
	for _, rootDomain := range rootDomains {
		if err := s.DiscoverAndStoreSubdomains(rootDomain, category, deepScan); err != nil {
			global.GVA_LOG.Error("批量发现子域名失败", zap.String("rootDomain", rootDomain), zap.Error(err))
		}
	}
	return nil
}

func (s *subdomainDiscovery) GetDomainTree(rootDomain string) ([]*model.DomainAsset, error) {
	var assets []model.DomainAsset
	err := global.GVA_DB.Where("root_domain = ?", rootDomain).Find(&assets).Error
	if err != nil {
		return nil, err
	}

	treeMap := make(map[uint]*model.DomainAsset)
	var rootNodes []*model.DomainAsset

	// 第一遍：建立 ID 到指针的映射
	for i := range assets {
		treeMap[assets[i].ID] = &assets[i]
		// 初始化子节点切片，防止前端解析 nil
		assets[i].Children = []*model.DomainAsset{}
	}

	// 第二遍：构建树结构
	for i := range assets {
		if assets[i].ParentID == nil || *assets[i].ParentID == 0 {
			rootNodes = append(rootNodes, &assets[i])
		} else if parent, exists := treeMap[*assets[i].ParentID]; exists {
			parent.Children = append(parent.Children, &assets[i])
			// 状态汇总
			if assets[i].DomainStatus == 3 {
				parent.DomainStatus = 3
			}
			if assets[i].CertStatus == 3 {
				parent.CertStatus = 3
			}
		} else {
			// 如果 ParentID 不在当前 rootDomain 的结果集中，也将其作为根节点展示
			rootNodes = append(rootNodes, &assets[i])
		}
	}

	return rootNodes, nil
}
