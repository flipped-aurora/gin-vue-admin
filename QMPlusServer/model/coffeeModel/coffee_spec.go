package coffeeModel

import (
	"errors"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CoffeeSpec struct {
	gorm.Model
	CoffeeSpecDetail []CoffeeSpecDetail `json:"coffee_spec_detail" gorm:"ForeignKey:SpecId;AssociationForeignKey:SpecId"`
	SpecId           uuid.UUID          `json:"spec_id"`
	Name             string             `json:"name"`
}

func (c *CoffeeSpec) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(c, info)
	if err != nil {
		return
	} else {
		var coffeeSpecList []CoffeeSpec
		err = db.Preload("CoffeeSpecDetail").Find(&coffeeSpecList).Error
		return err, coffeeSpecList, total
	}
}

func (c *CoffeeSpec) AddCoffeeSpec(name string, detail []CoffeeSpecDetail) (err error) {
	findOne := qmsql.DEFAULTDB.Where("name = ?", name).Find(&CoffeeSpec{}).Error
	if findOne != nil {
		c.SpecId = uuid.NewV4()
		c.Name = name
		err = qmsql.DEFAULTDB.Create(&c).Error
		if err != nil {
			return
		}
		c.CoffeeSpecDetail = detail
		for i := 0; i < len(c.CoffeeSpecDetail); i++ {
			c.CoffeeSpecDetail[i].SpecId = c.SpecId
			err = qmsql.DEFAULTDB.Create(&c.CoffeeSpecDetail[i]).Error
		}
	} else {
		err = errors.New("存在重复名称，请修改名称")
	}

	return
}

func (c *CoffeeSpec) UpdateCoffeeSpec() (err error) {
	return nil
}

func (c *CoffeeSpec) DeleteCoffeeSpec(id int64) (err error) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).Delete(&c).Error
	return
}

func (c *CoffeeSpec) GetCoffeeById(id int64) (err error) {
	return nil
}

func (c *CoffeeSpec) GetCoffeeSpecByCoffeeId(coffeeId uuid.UUID) (list []CoffeeSpec, err error) {
	var detail []CoffeeSpecDetail
	var spec CoffeeSpec
	err = qmsql.DEFAULTDB.Where("coffee_id = ?", coffeeId).Find(&detail).Error
	if err != nil {
		return nil, err
	}
	var temp = RemoveRepByLoop(detail)
	for i := 0; i < len(temp); i++ {
		err = qmsql.DEFAULTDB.Where("coffee_id = ?", coffeeId).Where("spec_id = ?", temp[i].SpecId).Find(&detail).Error
		if err != nil {
			return nil, err
		}
		spec.CoffeeSpecDetail = detail
		err = qmsql.DEFAULTDB.Where("spec_id = ?", temp[i].SpecId).Find(&spec).Error
		if err != nil {
			return nil, err
		}
		list = append(list, spec)
	}
	return list, nil
}
func RemoveRepByLoop(slc []CoffeeSpecDetail) []CoffeeSpecDetail {
	result := []CoffeeSpecDetail{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i].SpecId == result[j].SpecId {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}
