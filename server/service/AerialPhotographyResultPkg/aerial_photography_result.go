package AerialPhotographyResultPkg

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/AerialPhotographyResultPkg"
	AerialPhotographyResultPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/AerialPhotographyResultPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/NestInfo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func init() {
	go AutoCompressAerialPhotographyFile()
}

type AerialPhotographyResultService struct {
}

// CreateAerialPhotographyResult 创建AerialPhotographyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (ALPhotographyResultService *AerialPhotographyResultService) CreateAerialPhotographyResult(ALPhotographyResult *AerialPhotographyResultPkg.AerialPhotographyResult) (err error) {
	err = global.GVA_DB.Create(ALPhotographyResult).Error
	return err
}

// DeleteAerialPhotographyResult 删除AerialPhotographyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (ALPhotographyResultService *AerialPhotographyResultService) DeleteAerialPhotographyResult(ALPhotographyResult AerialPhotographyResultPkg.AerialPhotographyResult) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&AerialPhotographyResultPkg.AerialPhotographyResult{}).Where("id = ?", ALPhotographyResult.ID).Update("deleted_by", ALPhotographyResult.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&ALPhotographyResult).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAerialPhotographyResultByIds 批量删除AerialPhotographyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (ALPhotographyResultService *AerialPhotographyResultService) DeleteAerialPhotographyResultByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&AerialPhotographyResultPkg.AerialPhotographyResult{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&AerialPhotographyResultPkg.AerialPhotographyResult{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAerialPhotographyResult 更新AerialPhotographyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (ALPhotographyResultService *AerialPhotographyResultService) UpdateAerialPhotographyResult(ALPhotographyResult AerialPhotographyResultPkg.AerialPhotographyResult) (err error) {
	err = global.GVA_DB.Save(&ALPhotographyResult).Error
	return err
}

// GetAerialPhotographyResult 根据id获取AerialPhotographyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (ALPhotographyResultService *AerialPhotographyResultService) GetAerialPhotographyResult(id uint) (ALPhotographyResult AerialPhotographyResultPkg.AerialPhotographyResult, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ALPhotographyResult).Error
	return
}

// GetAerialPhotographyResultInfoList 分页获取AerialPhotographyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (ALPhotographyResultService *AerialPhotographyResultService) GetAerialPhotographyResultInfoList(info AerialPhotographyResultPkgReq.AerialPhotographyResultSearch, c *gin.Context) (list []AerialPhotographyResultPkg.AerialPhotographyResult, total int64, err error) {
	nestInfoService := new(NestInfo.NestInfoService)
	nestIDList, err := nestInfoService.GetNestIDListByUser(c)
	if err != nil {
		return
	}

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&AerialPhotographyResultPkg.AerialPhotographyResult{})
	if len(nestIDList) > 0 {
		for i, str := range nestIDList {
			if i == 0 {
				db.Where("nest_ids like ?", str)
			} else {
				db.Or("nest_ids like ?", str)
			}
		}
	}
	var ALPhotographyResults []AerialPhotographyResultPkg.AerialPhotographyResult
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Status != nil {
		if *info.Status >= 0 {
			db = db.Where("status = ?", info.Status)
		}
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&ALPhotographyResults).Error
	return ALPhotographyResults, total, err
}

//QueryAerialPhotographyResult 查询并处理航摄成果数据
func (ALPhotographyResultService *AerialPhotographyResultService) QueryAerialPhotographyResult(c *gin.Context) ([]AerialPhotographyResultPkg.AerialPhotographyResult, []AerialPhotographyResultPkg.AerialPhotographyResult, error) {
	var dataList []AerialPhotographyResultPkg.AerialPhotographyResult
	var modelList []AerialPhotographyResultPkg.AerialPhotographyResult
	var orthoList []AerialPhotographyResultPkg.AerialPhotographyResult
	nestInfoService := new(NestInfo.NestInfoService)
	nestIDList, err := nestInfoService.GetNestIDListByUser(c)
	if err != nil {
		return modelList, orthoList, err
	}
	querySql := "select id,name,photography_createtime,type,nest_ids, REPLACE(JSON_EXTRACT(aerial_photography_file, '$[0].url'),'\"','') aerial_photography_file, position, load_or_not from aerial_photography_result where 1 = 1 and status = 2 "
	db := global.GVA_DB.Model(&AerialPhotographyResultPkg.AerialPhotographyResult{})
	if len(nestIDList) > 0 {
		for i, str := range nestIDList {
			if i == 0 {
				querySql += " and nest_ids like '%" + str + "%'"
			} else {
				querySql += " or nest_ids like '%" + str + "%'"
			}
		}
	}
	querySql += " order by photography_createtime desc "
	queryErr := db.Raw(querySql).Find(&dataList)
	if queryErr.Error != nil {
	} else {
		for _, item := range dataList {
			if item.AerialPhotographyFile.String() != "" {
				url := item.AerialPhotographyFile.String()[0:strings.LastIndex(item.AerialPhotographyFile.String(), ".")]

				if *item.Type == 0 {
					//ortho
					if url != "" && item.Position != "" {
						posMap := make(map[string]int)
						parseErr := json.Unmarshal([]byte(item.Position), &posMap)
						if parseErr == nil {
							var info string
							info = filepath.Join(url, strconv.Itoa(posMap["x"]), strconv.Itoa(posMap["y"]), strconv.Itoa(posMap["z"])+".png")
							item.FileUrl = &info
							item.AerialPhotographyFile = nil //此处赋空值为避免最后返回结果时解析第三方包json报错
							orthoList = append(orthoList, item)

						} else {
							global.GVA_LOG.Error(parseErr.Error())
						}
					}
				} else if *item.Type == 1 {
					//model
					var info string
					info = filepath.Join(url, "tileset.json")
					item.FileUrl = &info
					item.AerialPhotographyFile = nil
					modelList = append(modelList, item)
				}
			}
		}
	}
	return modelList, orthoList, queryErr.Error

}

// AutoCompressAerialPhotographyFile 自动解压航摄文件并更新记录状态
// Author [piexlmax](https://github.com/piexlmax)
func AutoCompressAerialPhotographyFile() {
	dataHandlerTicker := time.NewTicker(time.Second * 60)
	wg := sync.WaitGroup{}
	wg.Add(1)
	quit := make(chan int)
	go func() {
		for {
			select {
			case <-dataHandlerTicker.C:
				alRes := new(AerialPhotographyResultPkg.AerialPhotographyResult)
				db := global.GVA_DB.Model(&AerialPhotographyResultPkg.AerialPhotographyResult{})
				db.Where("status = 0").Where("aerial_photography_file <> ''").Order("photography_createtime desc").Limit(1).Scan(alRes)
				if alRes.ID > 0 {
					global.GVA_LOG.Info("航摄成果自动解压开始")
					db.Exec("update aerial_photography_result set status = 1 where id = ?", alRes.ID)
					fileJsonStr := alRes.AerialPhotographyFile.String()
					if fileJsonStr != "" {
						fileDataArr := make([]map[string]interface{}, 0, 0)
						fmt.Println(fileJsonStr)
						err := json.Unmarshal([]byte(fileJsonStr), &fileDataArr)
						if err == nil {
							var url string
							//var name string
							if _, exist := fileDataArr[0]["url"]; exist {
								url = fileDataArr[0]["url"].(string)
							}
							//if _,exist := fileDataArr[0]["name"]; exist {
							//	name = fileDataArr[0]["name"].(string)
							//}
							var flag bool
							_, unzipErr := AerialPhotographyFileUnzip(url, url[0:strings.LastIndex(url, ".")], *alRes, flag)
							if unzipErr == nil {
								//解压成功,更新航摄成果状态
								db.Exec("update aerial_photography_result set status = 2 where id = ?", alRes.ID)
								global.GVA_LOG.Info("航摄成果自动解压结束")
							} else {
								db.Exec("update aerial_photography_result set status = 3 where id = ?", alRes.ID)
								global.GVA_LOG.Error("航摄成果自动解压异常:" + unzipErr.Error())
							}

						}
					}
				}
			case <-quit:
				dataHandlerTicker.Stop()
				return
			}
		}
	}()
	wg.Wait()
	dataHandlerTicker.Stop()
}

//解压航摄成果压缩包并更新坐标
func AerialPhotographyFileUnzip(zipFile string, destDir string, alRes AerialPhotographyResultPkg.AerialPhotographyResult, flag bool) ([]string, error) {
	zipReader, err := zip.OpenReader(zipFile)
	var paths []string
	if err != nil {
		return []string{}, err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		if strings.Index(f.Name, "..") > -1 {
			return []string{}, fmt.Errorf("%s 文件名不合法", f.Name)
		}
		fpath := filepath.Join(destDir, f.Name)
		paths = append(paths, fpath)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return []string{}, err
			}

			inFile, err := f.Open()
			if err != nil {
				return []string{}, err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return []string{}, err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return []string{}, err
			}

			if *alRes.Type == 0 {
				readFile, err := ioutil.ReadFile(outFile.Name())
				if err != nil {
					global.GVA_LOG.Info("ioutil.ReadFile error:" + err.Error())
				}
				if len(readFile) > 102400 && !flag {
					posMap := make(map[string]int64)
					name := outFile.Name()
					name = strings.ReplaceAll(name, "/", "\\")
					name = name[0:strings.LastIndex(name, ".")]
					splitResult := strings.Split(name, "\\")
					//OrthoPhotoLayer.Pos="{x:"+PosArr[2]+",y:"+PosArr[3]+",z:"+PosArr[1]+"}"
					posMap["y"], _ = strconv.ParseInt(splitResult[len(splitResult)-2], 10, 64)
					posMap["x"], _ = strconv.ParseInt(splitResult[len(splitResult)-3], 10, 64)
					posMap["z"], _ = strconv.ParseInt(splitResult[len(splitResult)-1], 10, 64)
					//map转json
					posJson, err := json.Marshal(posMap)
					if err != nil {
						global.GVA_LOG.Info("json.Marshal error:" + err.Error())
					}
					//json转str
					posStr := string(posJson)
					db := global.GVA_DB.Model(&AerialPhotographyResultPkg.AerialPhotographyResult{})
					db.Exec("update aerial_photography_result set Position = ? where id = ?", posStr, alRes.ID)
					flag = true
					global.GVA_LOG.Info("高清正射坐标更新成功")
				}
			}
		}
	}
	return paths, nil
}
