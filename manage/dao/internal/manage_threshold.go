// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManageThresholdDao is the data access object for the table manage_threshold.
type ManageThresholdDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  ManageThresholdColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// ManageThresholdColumns defines and stores column names for the table manage_threshold.
type ManageThresholdColumns struct {
	SensorId     string // 传感器id
	AlarmLabelId string // 报警标签id
	Template     string // expr表达式
	Sort         string // 优先级
}

// manageThresholdColumns holds the columns for the table manage_threshold.
var manageThresholdColumns = ManageThresholdColumns{
	SensorId:     "sensor_id",
	AlarmLabelId: "alarm_label_id",
	Template:     "template",
	Sort:         "sort",
}

// NewManageThresholdDao creates and returns a new DAO object for table data access.
func NewManageThresholdDao(handlers ...gdb.ModelHandler) *ManageThresholdDao {
	return &ManageThresholdDao{
		group:    "default",
		table:    "manage_threshold",
		columns:  manageThresholdColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ManageThresholdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ManageThresholdDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ManageThresholdDao) Columns() ManageThresholdColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ManageThresholdDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ManageThresholdDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ManageThresholdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
