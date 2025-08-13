// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageDeviceControl is the golang structure of table manage_device_control for DAO operations like Where/Data.
type ManageDeviceControl struct {
	g.Meta    `orm:"table:manage_device_control, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 命令作用描述
	DeviceId  interface{} // 设备Id
	Extend    *gjson.Json //
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Remark    interface{} // 备注
}
