// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageSensorType is the golang structure of table manage_sensor_type for DAO operations like Where/Data.
type ManageSensorType struct {
	g.Meta    `orm:"table:manage_sensor_type, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 传感器类型名称
	Unit      interface{} // 单位
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Remark    interface{} // 备注
}
