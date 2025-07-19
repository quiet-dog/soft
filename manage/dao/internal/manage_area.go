// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManageAreaDao is the data access object for the table manage_area.
type ManageAreaDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ManageAreaColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ManageAreaColumns defines and stores column names for the table manage_area.
type ManageAreaColumns struct {
	Id        string // 主键
	ParentId  string // 父ID
	Name      string // 区域名称
	Sort      string // 排序
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	Remark    string // 备注
}

// manageAreaColumns holds the columns for the table manage_area.
var manageAreaColumns = ManageAreaColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Name:      "name",
	Sort:      "sort",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Remark:    "remark",
}

// NewManageAreaDao creates and returns a new DAO object for table data access.
func NewManageAreaDao(handlers ...gdb.ModelHandler) *ManageAreaDao {
	return &ManageAreaDao{
		group:    "default",
		table:    "manage_area",
		columns:  manageAreaColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ManageAreaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ManageAreaDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ManageAreaDao) Columns() ManageAreaColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ManageAreaDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ManageAreaDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ManageAreaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
