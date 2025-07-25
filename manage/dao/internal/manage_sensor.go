// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManageSensorDao is the data access object for the table manage_sensor.
type ManageSensorDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  ManageSensorColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// ManageSensorColumns defines and stores column names for the table manage_sensor.
type ManageSensorColumns struct {
	Id           string // 主键
	Name         string // 传感器名称
	SensorTypeId string // 传感器类型Id
	DeviceId     string // 设备Id
	Extend       string //
	CreatedBy    string // 创建者
	UpdatedBy    string // 更新者
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
	Remark       string // 备注
	Template     string //
}

// manageSensorColumns holds the columns for the table manage_sensor.
var manageSensorColumns = ManageSensorColumns{
	Id:           "id",
	Name:         "name",
	SensorTypeId: "sensor_type_id",
	DeviceId:     "device_id",
	Extend:       "extend",
	CreatedBy:    "created_by",
	UpdatedBy:    "updated_by",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	Remark:       "remark",
	Template:     "template",
}

// NewManageSensorDao creates and returns a new DAO object for table data access.
func NewManageSensorDao(handlers ...gdb.ModelHandler) *ManageSensorDao {
	return &ManageSensorDao{
		group:    "default",
		table:    "manage_sensor",
		columns:  manageSensorColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ManageSensorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ManageSensorDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ManageSensorDao) Columns() ManageSensorColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ManageSensorDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ManageSensorDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ManageSensorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
