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
type GfGoodsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  GfGoodsColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// GfUsersColumns defines and stores column names for the table gf_users.
type GfGoodsColumns struct {
	Id            string //
	GoodsName     string // 商品名称
	GoodsDesc     string // 商品描述
	GoodsPrice    string // 商品价格
	GoodsStock    string // 商品库存
	GoodsType     string // 商品类型
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
}

// gfGoodsColumns holds the columns for the table gf_goods.
var gfGoodsColumns = GfGoodsColumns{
	Id:            "id",
	GoodsName:     "goodsName", // 商品名称
	GoodsDesc:     "goodsDesc", // 商品描述
	GoodsPrice:    "goodsPrice", // 商品价格
	GoodsStock:    "goodsStock", // 商品库存
	GoodsType:     "goodsType", // 商品类型
	CreateTime:    "createTime", // 创建时间
	UpdateTime:    "updateTime", // 更新时间
}

// NewGfGoodsDao creates and returns a new DAO object for table data access.
func NewGfGoodsDao(handlers ...gdb.ModelHandler) *GfGoodsDao {
	return &GfGoodsDao{
		group:    "default",
		table:    "gf_goods",
		columns:  gfGoodsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GfGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GfGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GfGoodsDao) Columns() GfGoodsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GfGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GfGoodsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *GfGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
