package customerModel

import (
	"errors"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/tools"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Customers struct {
	gorm.Model
	UUID     uuid.UUID       `json:"uuid"`
	Username string          `json:"username"`
	Password string          `json:"password"`
	Nickname string          `json:"nickname"`
	Email    string          `json:"email"`
	Phone    string          `json:"phone"`
	Image    string          `json:"image"`
	Address  []Address       `json:"address" gorm:"ForeignKey:UserId;"`
	Order    []CustomerOrder `json:"order" gorm:"ForeginKey:UserId"`
}

func (c *Customers) Register() (err error, u *Customers) {
	var user Customers
	//判断用户名是否注册
	notResigt := qmsql.DEFAULTDB.Where("username = ?", c.Username).First(&user).RecordNotFound()
	//notResigt为false表明读取到了 不能注册
	if !notResigt {
		return errors.New("用户名已注册"), nil
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		c.Password = tools.MD5V([]byte(c.Password))
		c.UUID = uuid.NewV4()
		err = qmsql.DEFAULTDB.Create(c).Error
	}
	return err, c
}

//修改用户密码
func (c *Customers) ChangePassword(newPassword string) (err error, u *Customers) {
	var user Customers
	//后期修改jwt+password模式
	c.Password = tools.MD5V([]byte(c.Password))
	err = qmsql.DEFAULTDB.Where("username = ? AND password = ?", c.Username, c.Password).First(&user).Update("password", tools.MD5V([]byte(newPassword))).Error
	return err, c
}

func (c *Customers) Login() (err error, u *Customers) {
	var user Customers
	c.Password = tools.MD5V([]byte(c.Password))
	err = qmsql.DEFAULTDB.Where("username = ? AND password = ?", c.Username, c.Password).First(&user).Error
	if err != nil {
		return err, &user
	}
	return err, &user
}

// 用户头像上传更新地址
func (c *Customers) UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *Customers) {
	var user Customers
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).First(&user).Update("image", filePath).First(&user).Error
	return err, &user
}

// 分页获取数据  需要分页实现这个接口即可
func (c *Customers) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(c, info)
	if err != nil {
		return
	} else {
		var userList []Customers
		err = db.Find(&userList).Error
		return err, userList, total
	}
}

func (c *Customers) AddCustomer() (err error) {
	findOne := qmsql.DEFAULTDB.Where("username = ?", c.Username).Find(&Customers{}).Error
	if findOne != nil {
		c.UUID = uuid.NewV4()
		err = qmsql.DEFAULTDB.Create(c).Error
	} else {
		err = errors.New("存在重复名称，请修改名称")
	}
	return
}

func (c *Customers) UpdateCustomer() (err error) {
	err = qmsql.DEFAULTDB.Where("username = ?", c.Username).Find(&Customers{}).Error
	if err != nil {
		err = errors.New("存在重复名称,请重新输入")
		return
	}
	updataMap := make(map[string]interface{})
	updataMap["username"] = c.Username
	updataMap["password"] = c.Password
	updataMap["nickname"] = c.Nickname
	updataMap["image"] = c.Image
	updataMap["email"] = c.Email
	updataMap["phone"] = c.Phone
	err = qmsql.DEFAULTDB.Where("uuid = ?", c.UUID).Find(&Customers{}).Update(updataMap).Error
	return
}

func (c *Customers) DeleteCustomer(uuid uuid.UUID) (err error) {
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).Delete(&Customers{}).Error
	return
}

func (c *Customers) GetCustomerById(uuid uuid.UUID) (err error) {
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).Find(&c).Error
	return
}
