// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManageDeviceDao is the data access object for the table manage_device.
type ManageDeviceDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  ManageDeviceColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// ManageDeviceColumns defines and stores column names for the table manage_device.
type ManageDeviceColumns struct {
	Id                   string // 主键
	Name                 string // 设备名称
	Manufacturer         string // 厂商
	Model                string // 型号
	InstallationLocation string // 安装位置
	AreaId               string // 所属区域
	ServerId             string // 设备服务id
	CreatedBy            string // 创建者
	UpdatedBy            string // 更新者
	CreatedAt            string // 创建时间
	UpdatedAt            string // 更新时间
	DeletedAt            string // 删除时间
	Remark               string // 备注
	Extend               string // 扩展信息
}

// manageDeviceColumns holds the columns for the table manage_device.
var manageDeviceColumns = ManageDeviceColumns{
	Id:                   "id",
	Name:                 "name",
	Manufacturer:         "manufacturer",
	Model:                "model",
	InstallationLocation: "installation_location",
	AreaId:               "area_id",
	ServerId:             "server_id",
	CreatedBy:            "created_by",
	UpdatedBy:            "updated_by",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
	DeletedAt:            "deleted_at",
	Remark:               "remark",
	Extend:               "extend",
}

// NewManageDeviceDao creates and returns a new DAO object for table data access.
func NewManageDeviceDao(handlers ...gdb.ModelHandler) *ManageDeviceDao {
	return &ManageDeviceDao{
		group:    "default",
		table:    "manage_device",
		columns:  manageDeviceColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ManageDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ManageDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ManageDeviceDao) Columns() ManageDeviceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ManageDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ManageDeviceDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ManageDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
