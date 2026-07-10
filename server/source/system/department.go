package system

import (
	"context"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// 部门/岗位建表与种子; 与菜单/授权无 context 依赖, 排在授权之后即可
const initOrderDepartment = initOrderMenuAuthority + 1

type initDepartment struct{}

// auto run
func init() {
	system.RegisterInit(initOrderDepartment, &initDepartment{})
}

func (i *initDepartment) InitializerName() string {
	return sysModel.SysDepartment{}.TableName()
}

func (i *initDepartment) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&sysModel.SysDepartment{},
		&sysModel.SysPosition{},
		&sysModel.SysDataAccessLog{},
		&sysModel.SysAuthorityDepartment{},
	)
}

func (i *initDepartment) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&sysModel.SysDepartment{}) && m.HasTable(&sysModel.SysPosition{})
}

func (i *initDepartment) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	enable := true
	departments := []sysModel.SysDepartment{
		{Name: "总公司", ParentId: 0, Ancestors: "0", Sort: 0, Status: &enable},
	}
	if err := db.Create(&departments).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysDepartment{}.TableName()+"表数据初始化失败!")
	}
	positions := []sysModel.SysPosition{
		{Name: "总经理", Code: "CEO", Sort: 1, Status: &enable},
		{Name: "普通员工", Code: "STAFF", Sort: 2, Status: &enable},
	}
	if err := db.Create(&positions).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysPosition{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), departments)
	return next, nil
}

func (i *initDepartment) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "总公司").First(&sysModel.SysDepartment{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
