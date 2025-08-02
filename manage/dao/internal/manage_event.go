// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManageEventDao is the data access object for the table manage_event.
type ManageEventDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ManageEventColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ManageEventColumns defines and stores column names for the table manage_event.
type ManageEventColumns struct {
	Id          string // 主键
	SensorId    string // 传感器Id
	Value       string // 报警时候的数值
	Description string // 报警描述
	CreatedBy   string // 创建者
	UpdatedBy   string // 更新者
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
	Level       string // 报警等级
	Color       string // 报警颜色
	AlarmId     string //
}

// manageEventColumns holds the columns for the table manage_event.
var manageEventColumns = ManageEventColumns{
	Id:          "id",
	SensorId:    "sensor_id",
	Value:       "value",
	Description: "description",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	Level:       "level",
	Color:       "color",
	AlarmId:     "alarm_id",
}

// NewManageEventDao creates and returns a new DAO object for table data access.
func NewManageEventDao(handlers ...gdb.ModelHandler) *ManageEventDao {
	return &ManageEventDao{
		group:    "default",
		table:    "manage_event",
		columns:  manageEventColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ManageEventDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ManageEventDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ManageEventDao) Columns() ManageEventColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ManageEventDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ManageEventDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ManageEventDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
