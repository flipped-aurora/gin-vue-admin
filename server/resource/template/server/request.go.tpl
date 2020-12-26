package request

import "gin-vue-admin/model"

type {{.StructName}}Search struct {
    model.{{.StructName}}Base
    PageInfo
}

type {{.StructName}}Update struct {
	model.{{.StructName}}Base
	ID uint `gorm:"primarykey"`
}

