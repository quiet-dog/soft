// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManageSensorTypeDao is the data access object for the table manage_sensor_type.
type ManageSensorTypeDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  ManageSensorTypeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// ManageSensorTypeColumns defines and stores column names for the table manage_sensor_type.
type ManageSensorTypeColumns struct {
	Id        string // 主键
	Name      string // 传感器类型名称
	Unit      string // 单位
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	Remark    string // 备注
}

// manageSensorTypeColumns holds the columns for the table manage_sensor_type.
var manageSensorTypeColumns = ManageSensorTypeColumns{
	Id:        "id",
	Name:      "name",
	Unit:      "unit",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Remark:    "remark",
}

// NewManageSensorTypeDao creates and returns a new DAO object for table data access.
func NewManageSensorTypeDao(handlers ...gdb.ModelHandler) *ManageSensorTypeDao {
	return &ManageSensorTypeDao{
		group:    "default",
		table:    "manage_sensor_type",
		columns:  manageSensorTypeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ManageSensorTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ManageSensorTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ManageSensorTypeDao) Columns() ManageSensorTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ManageSensorTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ManageSensorTypeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ManageSensorTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
