// 自动生成模板Test
package autocode

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type Test struct {
      gorm.Model 
      TestComponent  string `json:"testComponent"`    
      TestBigComponent  int `json:"testBigComponent"`    
}

// 创建Test
func (t *Test)CreateTest()(err error){
        err = qmsql.DEFAULTDB.Create(t).Error
        return err
}

// 删除Test
func (t *Test)DeleteTest()(err error){
        err = qmsql.DEFAULTDB.Delete(t).Error
        return err
}

// 更新Test
func (t *Test)UpdateTest()(err error, ret Test){
        err = qmsql.DEFAULTDB.Save(t).Error
        return err, *t
}

// 根据ID查看单条Test
func (t *Test)FindById()(err error,ret Test){
    err = qmsql.DEFAULTDB.Where("id = ?",t.ID).First(&ret).Error
    return err,ret
}

// 分页获取Test
func (t *Test)GetInfoList(info modelInterface.PageInfo)(err error, list interface{}, total int){
    	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
    	err, db, total := servers.PagingServer(t, info)
    	if err != nil {
    		return
    	} else {
    		var reTestList []Test
    		err = db.Find(&reTestList).Error
    		return err, reTestList, total
    	}
}