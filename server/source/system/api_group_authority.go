package system

import (
	"context"
	sysModel "github.com/edufriendchen/hertz-vue-admin/server/model/system"
	"github.com/edufriendchen/hertz-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderApiGroupAuthority = initOrderApiGroup + initOrderAuthority

type initApiGroupAuthority struct{}

// auto run
func init() {
	system.RegisterInit(initOrderApiGroupAuthority, &initApiGroupAuthority{})
}

func (i *initApiGroupAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysRolesApiGroup{})
}

func (i *initApiGroupAuthority) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysRolesApiGroup{})
}

func (i initApiGroupAuthority) InitializerName() string {
	return (&sysModel.SysRolesApiGroup{}).TableName()
}

func (i *initApiGroupAuthority) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	// 权限定义：0=有权限，1=无权限
	// 对应 common.PermissionType 的索引
	fullAccess := []byte{0, 0, 0, 0, 0} // 完全访问权限
	readOnly := []byte{0, 1, 1, 1, 1}   // 只读权限
	noAccess := []byte{1, 1, 1, 1, 1}   // 无权限

	entities := []sysModel.SysRolesApiGroup{
		// 超级管理员(888) - 完全权限
		{RoleId: 888, ApiGroupId: 1, Permission: fullAccess}, // base
		{RoleId: 888, ApiGroupId: 2, Permission: fullAccess}, // jwt
		{RoleId: 888, ApiGroupId: 3, Permission: fullAccess}, // 系统用户
		{RoleId: 888, ApiGroupId: 4, Permission: fullAccess}, // controller
		{RoleId: 888, ApiGroupId: 5, Permission: fullAccess}, // role
		{RoleId: 888, ApiGroupId: 6, Permission: fullAccess}, // casbin
		{RoleId: 888, ApiGroupId: 7, Permission: fullAccess}, // 菜单
		{RoleId: 888, ApiGroupId: 8, Permission: fullAccess}, // 分片上传
		{RoleId: 888, ApiGroupId: 9, Permission: fullAccess}, // 文件上传与下载
		{RoleId: 888, ApiGroupId: 10, Permission: fullAccess}, // 系统服务
		{RoleId: 888, ApiGroupId: 11, Permission: fullAccess}, // 客户
		{RoleId: 888, ApiGroupId: 12, Permission: fullAccess}, // 代码生成器
		{RoleId: 888, ApiGroupId: 13, Permission: fullAccess}, // 包（pkg）生成器
		{RoleId: 888, ApiGroupId: 14, Permission: fullAccess}, // 代码生成器历史
		{RoleId: 888, ApiGroupId: 15, Permission: fullAccess}, // 系统字典详情
		{RoleId: 888, ApiGroupId: 16, Permission: fullAccess}, // 系统字典
		{RoleId: 888, ApiGroupId: 17, Permission: fullAccess}, // 操作记录
		{RoleId: 888, ApiGroupId: 18, Permission: fullAccess}, // 断点续传(插件版)
		{RoleId: 888, ApiGroupId: 19, Permission: fullAccess}, // email
		{RoleId: 888, ApiGroupId: 20, Permission: fullAccess}, // excel
		{RoleId: 888, ApiGroupId: 21, Permission: fullAccess}, // 按钮权限

		// 测试角色(9528) - 只读权限
		{RoleId: 9528, ApiGroupId: 1, Permission: readOnly}, // base
		{RoleId: 9528, ApiGroupId: 2, Permission: readOnly}, // jwt
		{RoleId: 9528, ApiGroupId: 3, Permission: readOnly}, // 系统用户
		{RoleId: 9528, ApiGroupId: 4, Permission: readOnly}, // controller
		{RoleId: 9528, ApiGroupId: 5, Permission: readOnly}, // role
		{RoleId: 9528, ApiGroupId: 6, Permission: readOnly}, // casbin
		{RoleId: 9528, ApiGroupId: 7, Permission: readOnly}, // 菜单
		{RoleId: 9528, ApiGroupId: 8, Permission: readOnly}, // 分片上传
		{RoleId: 9528, ApiGroupId: 9, Permission: readOnly}, // 文件上传与下载
		{RoleId: 9528, ApiGroupId: 10, Permission: readOnly}, // 系统服务
		{RoleId: 9528, ApiGroupId: 11, Permission: readOnly}, // 客户
		{RoleId: 9528, ApiGroupId: 12, Permission: readOnly}, // 代码生成器
		{RoleId: 9528, ApiGroupId: 13, Permission: readOnly}, // 包（pkg）生成器
		{RoleId: 9528, ApiGroupId: 14, Permission: readOnly}, // 代码生成器历史
		{RoleId: 9528, ApiGroupId: 15, Permission: readOnly}, // 系统字典详情
		{RoleId: 9528, ApiGroupId: 16, Permission: readOnly}, // 系统字典
		{RoleId: 9528, ApiGroupId: 17, Permission: readOnly}, // 操作记录
		{RoleId: 9528, ApiGroupId: 18, Permission: readOnly}, // 断点续传(插件版)
		{RoleId: 9528, ApiGroupId: 19, Permission: readOnly}, // email
		{RoleId: 9528, ApiGroupId: 20, Permission: readOnly}, // excel
		{RoleId: 9528, ApiGroupId: 21, Permission: readOnly}, // 按钮权限

		// 普通用户子角色(8881) - 无权限
		{RoleId: 8881, ApiGroupId: 1, Permission: noAccess}, // base
		{RoleId: 8881, ApiGroupId: 2, Permission: noAccess}, // jwt
		{RoleId: 8881, ApiGroupId: 3, Permission: noAccess}, // 系统用户
		{RoleId: 8881, ApiGroupId: 4, Permission: noAccess}, // controller
		{RoleId: 8881, ApiGroupId: 5, Permission: noAccess}, // role
		{RoleId: 8881, ApiGroupId: 6, Permission: noAccess}, // casbin
		{RoleId: 8881, ApiGroupId: 7, Permission: noAccess}, // 菜单
		{RoleId: 8881, ApiGroupId: 8, Permission: noAccess}, // 分片上传
		{RoleId: 8881, ApiGroupId: 9, Permission: noAccess}, // 文件上传与下载
		{RoleId: 8881, ApiGroupId: 10, Permission: noAccess}, // 系统服务
		{RoleId: 8881, ApiGroupId: 11, Permission: noAccess}, // 客户
		{RoleId: 8881, ApiGroupId: 12, Permission: noAccess}, // 代码生成器
		{RoleId: 8881, ApiGroupId: 13, Permission: noAccess}, // 包（pkg）生成器
		{RoleId: 8881, ApiGroupId: 14, Permission: noAccess}, // 代码生成器历史
		{RoleId: 8881, ApiGroupId: 15, Permission: noAccess}, // 系统字典详情
		{RoleId: 8881, ApiGroupId: 16, Permission: noAccess}, // 系统字典
		{RoleId: 8881, ApiGroupId: 17, Permission: noAccess}, // 操作记录
		{RoleId: 8881, ApiGroupId: 18, Permission: noAccess}, // 断点续传(插件版)
		{RoleId: 8881, ApiGroupId: 19, Permission: noAccess}, // email
		{RoleId: 8881, ApiGroupId: 20, Permission: noAccess}, // excel
		{RoleId: 8881, ApiGroupId: 21, Permission: noAccess}, // 按钮权限
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, (&sysModel.SysRolesApiGroup{}).TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApiGroupAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("role_id = ? AND api_group_id = ?", 888, 5).
		First(&sysModel.SysRolesApiGroup{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
