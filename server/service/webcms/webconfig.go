package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
)

type WebconfigService struct {
}

// 创建记录
func (WebconfigService *WebconfigService) CreateWebconfig(webconfig webcms.Webconfig) (err error) {
	err = global.GVA_DB.Create(&webconfig).Error
	return err
}

// 获取记录
func (WebconfigService *WebconfigService) GetWebconfig() (webconfig []webcms.Webconfig, err error) {
	err = global.GVA_DB.Find(&webconfig).Error
	if err != nil {
		return
	}
	// for _, val := range webconfig {
	// 	err = severedis(val)
	// 	if err != nil {
	// 		return
	// 	}
	// }
	return
}

// 删除记录
func (WebconfigService *WebconfigService) DeleteWebconfig(id uint) (err error) {
	err = global.GVA_DB.Delete(&[]webcms.Webconfig{}, "id = ?", id).Error
	return err
}

// func severedis(webconfig webcms.Webconfig) (err error) {
// 	// 将数据保存到redis
// 	t := reflect.TypeOf(webconfig)
// 	v := reflect.ValueOf(webconfig)
// 	keyvalue := make(map[string]interface{}, t.NumField())
// 	for i := 0; i < t.NumField(); i++ {
// 		if t.Field(i).Tag.Get("json") == "" {
// 			continue
// 		}
// 		keyvalue[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
// 	}
// 	// fmt.Println(keyvalue)
// 	keyvaluestr, _ := json.Marshal(keyvalue)
// 	err = global.GVA_BigCache.Set(fmt.Sprint("webconfig:", webconfig.ID), keyvaluestr)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// UpdatesWebconfig 更新记录
func (WebconfigService *WebconfigService) UpdatesWebconfig(webconfig webcms.Webconfig) (err error) {
	err = global.GVA_DB.Updates(&webconfig).Error
	if err != nil {
		return
	}
	// err = severedis(webconfig)
	// if err != nil {
	// 	return
	// }
	return
}

// 通过field 获取值
func (WebconfigService *WebconfigService) GetValueByField(webconfig webcms.Webconfig, filed string) (value string, err error) {
	err = global.GVA_DB.Model(webcms.Webconfig{}).Select(filed).Find(&value).Error
	// webconfigbyte, err := global.GVA_BigCache.Get(fmt.Sprint("webconfig:", webconfig.ID))
	// if err != nil {
	// 	return
	// }
	// var webconfigsstruct map[string]string
	// err = json.Unmarshal(webconfigbyte, &webconfigsstruct)
	// if err != nil {
	// 	return
	// }
	// if value, ok := webconfigsstruct[filed]; !ok {
	// 	return value, nil
	// }
	return
}
