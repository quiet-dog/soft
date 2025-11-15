// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManageOpcDao is the data access object for the table manage_opc.
type ManageOpcDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ManageOpcColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ManageOpcColumns defines and stores column names for the table manage_opc.
type ManageOpcColumns struct {
	Id              string // 主键
	NodeId          string // 标签名称
	ServerId        string // 服务id
	NodeClass       string // 变量名称
	Type            string // 变量类型
	CreatedBy       string // 创建者
	UpdatedBy       string // 更新者
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
	DeletedAt       string // 删除时间
	Remark          string // 备注
	ParentId        string //
	NamespacesIndex string //
	BrowseName      string //
	DisplayName     string // 显示名称
}

// manageOpcColumns holds the columns for the table manage_opc.
var manageOpcColumns = ManageOpcColumns{
	Id:              "id",
	NodeId:          "node_id",
	ServerId:        "server_id",
	NodeClass:       "node_class",
	Type:            "type",
	CreatedBy:       "created_by",
	UpdatedBy:       "updated_by",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	Remark:          "remark",
	ParentId:        "parent_id",
	NamespacesIndex: "namespaces_index",
	BrowseName:      "browse_name",
	DisplayName:     "display_name",
}

// NewManageOpcDao creates and returns a new DAO object for table data access.
func NewManageOpcDao(handlers ...gdb.ModelHandler) *ManageOpcDao {
	return &ManageOpcDao{
		group:    "default",
		table:    "manage_opc",
		columns:  manageOpcColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ManageOpcDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ManageOpcDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ManageOpcDao) Columns() ManageOpcColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ManageOpcDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ManageOpcDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ManageOpcDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
