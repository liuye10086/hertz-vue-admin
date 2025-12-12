package system

import (
	"context"

	sysModel "github.com/edufriendchen/hertz-vue-admin/server/model/system"
	"github.com/edufriendchen/hertz-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initApiGroup struct{}

const initOrderApiGroup = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderApiGroup, &initApiGroup{})
}

func (i initApiGroup) InitializerName() string {
	return (&sysModel.SysApiGroup{}).TableName()
}

func (i *initApiGroup) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysApiGroup{})
}

func (i *initApiGroup) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysApiGroup{})
}

func (i *initApiGroup) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysApiGroup{
		{Name: "base", Path: "/base"},
		{Name: "jwt", Path: "/jwt"},
		{Name: "系统用户", Path: "/user"},
		{Name: "controller", Path: "/controller"},
		{Name: "role", Path: "/authority"},
		{Name: "casbin", Path: "/casbin"},
		{Name: "菜单", Path: "/menu"},
		{Name: "分片上传", Path: "/fileUploadAndDownload"},
		{Name: "文件上传与下载", Path: "/fileUploadAndDownload"},
		{Name: "系统服务", Path: "/system"},
		{Name: "客户", Path: "/customer"},
		{Name: "代码生成器", Path: "/autoCode"},
		{Name: "包（pkg）生成器", Path: "/autoCode"},
		{Name: "代码生成器历史", Path: "/autoCode"},
		{Name: "系统字典详情", Path: "/sysDictionaryDetail"},
		{Name: "系统字典", Path: "/sysDictionary"},
		{Name: "操作记录", Path: "/sysOperationRecord"},
		{Name: "断点续传(插件版)", Path: "/simpleUploader"},
		{Name: "email", Path: "/email"},
		{Name: "excel", Path: "/excel"},
		{Name: "按钮权限", Path: "/authorityBtn"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, (&sysModel.SysApiGroup{}).TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApiGroup) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "role").
		First(&sysModel.SysApiGroup{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
