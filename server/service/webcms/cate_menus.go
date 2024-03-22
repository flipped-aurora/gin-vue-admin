package webcms

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
)

type CateMenusService struct {
}

// CreateCateMenus 创建CateMenus记录
// Author [piexlmax](https://github.com/piexlmax)
func (cateMenusService *CateMenusService) CreateCateMenus(cateMenus []webcms.CateMenus) (err error) {
	err = global.GVA_DB.Create(&cateMenus).Error
	return err
}

// DeleteCateMenus 删除CateMenus记录
// Author [piexlmax](https://github.com/piexlmax)
func (cateMenusService *CateMenusService) DeleteCateMenus(cateMenus webcms.CateMenus) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&cateMenus).Error
	return err
}

// DeleteCateMenusByIds 批量删除CateMenus记录
// Author [piexlmax](https://github.com/piexlmax)
func (cateMenusService *CateMenusService) DeleteCateMenusByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]webcms.CateMenus{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCateMenus 更新CateMenus记录
// Author [piexlmax](https://github.com/piexlmax)
func (cateMenusService *CateMenusService) UpdateCateMenus(cateMenus webcms.CateMenus) (err error) {
	err = global.GVA_DB.Save(&cateMenus).Error
	return err
}

// GetCateMenus 根据id获取CateMenus记录
// Author [piexlmax](https://github.com/piexlmax)
func (cateMenusService *CateMenusService) GetCateMenus(id uint) (cateMenus webcms.CateMenus, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cateMenus).Error
	switch cateMenus.CateType {
	case 1:
		cateMenus.Url = fmt.Sprint("/category/", cateMenus.ID)
	case 2:
		cateMenus.Url = fmt.Sprint("/list/", cateMenus.ID)
	default:
		cateMenus.Url = cateMenus.Islink
	}
	return
}

// GetCateMenusInfoList 分页获取CateMenus记录
// Author [piexlmax](https://github.com/piexlmax)
func (cateMenusService *CateMenusService) GetCateMenusInfoList(info webcmsReq.CateMenusSearch, siteinfo map[string]any) (list []webcms.CateMenus, total int64, err error) {
	var cateMenuss []webcms.CateMenus
	treeMap, err := cateMenusService.getBaseMenuTreeMap(siteinfo)
	cateMenuss = treeMap[0]
	for i := 0; i < len(cateMenuss); i++ {
		err = cateMenusService.getBaseChildrenList(&cateMenuss[i], treeMap, 0)
	}
	return cateMenuss, total, err
}

func (cateMenusService *CateMenusService) getBaseChildrenList(menu *webcms.CateMenus, treeMap map[uint][]webcms.CateMenus, level int) (err error) {
	menu.Children = treeMap[menu.ID]
	menu.MenuLevel = level
	for i := 0; i < len(menu.Children); i++ {
		err = cateMenusService.getBaseChildrenList(&menu.Children[i], treeMap, level+1)
	}
	return err
}

func (cateMenusService *CateMenusService) getBaseMenuTreeMap(siteinfo map[string]any) (treeMap map[uint][]webcms.CateMenus, err error) {
	var allMenus []webcms.CateMenus
	treeMap = make(map[uint][]webcms.CateMenus)
	err = global.GVA_DB.Where("siteid", siteinfo["id"]).Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		switch v.CateType {
		case 1:
			v.Url = fmt.Sprint(siteinfo["siteUrl"], "category/", v.ID)
		case 2:
			v.Url = fmt.Sprint(siteinfo["siteUrl"], "list/", v.ID)
		default:
			v.Url = v.Islink
		}
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// 通过 parent_id 获取栏目信息
func (cateMenusService *CateMenusService) GetCateMenusListByPid(pid uint) (list []webcms.CateMenus, err error) {
	err = global.GVA_DB.Where("parent_id = ?", pid).Where("hidden = ?", 0).Order("sort").Find(&list).Error
	for k, v := range list {
		switch v.CateType {
		case 1:
			list[k].Url = fmt.Sprint("/category/", v.ID)
		case 2:
			list[k].Url = fmt.Sprint("/list/", v.ID)
		default:
			list[k].Url = v.Islink
		}
	}
	return
}

// 获取全部栏目
func (cateMenusService *CateMenusService) GetCateMenusInfoList2(siteid int) (list []webcms.CateMenus, err error) {
	var cateMenuss []webcms.CateMenus
	treeMap, err := cateMenusService.getBaseMenuTreeMap2(siteid)
	cateMenuss = treeMap[0]
	for i := 0; i < len(cateMenuss); i++ {
		err = cateMenusService.getBaseChildrenList(&cateMenuss[i], treeMap, i)
	}
	return cateMenuss, err
}

func (cateMenusService *CateMenusService) getBaseMenuTreeMap2(siteid int) (treeMap map[uint][]webcms.CateMenus, err error) {
	var allMenus []webcms.CateMenus
	treeMap = make(map[uint][]webcms.CateMenus)
	err = global.GVA_DB.Where("hidden = ?", 0).Where("siteid", siteid).Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		switch v.CateType {
		case 1:
			v.Url = fmt.Sprint("/category/", v.ID)
		case 2:
			v.Url = fmt.Sprint("/list/", v.ID)
		default:
			v.Url = v.Islink
		}
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// GetTemplateList 获取模板名称列表
func (cateMenusService *CateMenusService) GetTemplateList() (templateList []string, err error) {
	removeslice := []string{"header.html", "index.html", "footer.html", "404.html", "500.html"}
	filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileExt := filepath.Ext(info.Name())
			if fileExt == ".html" {
				filename := info.Name()
				if !StrSlice(removeslice, filename) {
					templateList = append(templateList, filename)
				}
			}
		}
		return nil
	})
	return
}

// GetModelsList  获取模型列表
func (cateMenusService *CateMenusService) GetModelsList() (modellist []system.SysAutoCodeHistory, err error) {
	err = global.GVA_DB.Find(&modellist).Error
	return
}
