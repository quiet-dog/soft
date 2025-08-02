// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManageAlarmDao is the data access object for the table manage_alarm.
type ManageAlarmDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ManageAlarmColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ManageAlarmColumns defines and stores column names for the table manage_alarm.
type ManageAlarmColumns struct {
	Id        string // 主键
	IsLift    string // 是否解除报警
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	Level     string //
	SensorId  string // 传感器Id
	Color     string // 颜色
}

// manageAlarmColumns holds the columns for the table manage_alarm.
var manageAlarmColumns = ManageAlarmColumns{
	Id:        "id",
	IsLift:    "is_lift",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Level:     "level",
	SensorId:  "sensor_id",
	Color:     "color",
}

// NewManageAlarmDao creates and returns a new DAO object for table data access.
func NewManageAlarmDao(handlers ...gdb.ModelHandler) *ManageAlarmDao {
	return &ManageAlarmDao{
		group:    "default",
		table:    "manage_alarm",
		columns:  manageAlarmColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ManageAlarmDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ManageAlarmDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ManageAlarmDao) Columns() ManageAlarmColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ManageAlarmDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ManageAlarmDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ManageAlarmDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
