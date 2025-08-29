// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GfUsersDao is the data access object for the table gf_users.
type GfUsersDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  GfUsersColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// GfUsersColumns defines and stores column names for the table gf_users.
type GfUsersColumns struct {
	Id            string //
	Username      string // 账号
	Nickname      string // 昵称
	Password      string // 密码
	Status        string // 状态:0=正常,1=禁用
	RegistTime    string // 注册时间
	LastLoginTime string // 上次登录时间
	OpenId        string // 第三方openId
}

// gfUsersColumns holds the columns for the table gf_users.
var gfUsersColumns = GfUsersColumns{
	Id:            "id",
	Username:      "username",
	Nickname:      "nickname",
	Password:      "password",
	Status:        "status",
	RegistTime:    "registTime",
	LastLoginTime: "lastLoginTime",
	OpenId:        "openId",
}

// NewGfUsersDao creates and returns a new DAO object for table data access.
func NewGfUsersDao(handlers ...gdb.ModelHandler) *GfUsersDao {
	return &GfUsersDao{
		group:    "default",
		table:    "gf_users",
		columns:  gfUsersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GfUsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GfUsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GfUsersDao) Columns() GfUsersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GfUsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GfUsersDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *GfUsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
