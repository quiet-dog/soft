// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManageAlarmLabelDao is the data access object for the table manage_alarm_label.
type ManageAlarmLabelDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  ManageAlarmLabelColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// ManageAlarmLabelColumns defines and stores column names for the table manage_alarm_label.
type ManageAlarmLabelColumns struct {
	Id        string // 主键
	Name      string // 标签名称
	Color     string // 标签颜色
	Level     string // 标签等级
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	Remark    string // 备注
	NoticeId  string // 通知ID
}

// manageAlarmLabelColumns holds the columns for the table manage_alarm_label.
var manageAlarmLabelColumns = ManageAlarmLabelColumns{
	Id:        "id",
	Name:      "name",
	Color:     "color",
	Level:     "level",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Remark:    "remark",
	NoticeId:  "notice_id",
}

// NewManageAlarmLabelDao creates and returns a new DAO object for table data access.
func NewManageAlarmLabelDao(handlers ...gdb.ModelHandler) *ManageAlarmLabelDao {
	return &ManageAlarmLabelDao{
		group:    "default",
		table:    "manage_alarm_label",
		columns:  manageAlarmLabelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ManageAlarmLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ManageAlarmLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ManageAlarmLabelDao) Columns() ManageAlarmLabelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ManageAlarmLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ManageAlarmLabelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ManageAlarmLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
