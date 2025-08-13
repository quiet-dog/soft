// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageDevice is the golang structure of table manage_device for DAO operations like Where/Data.
type ManageDevice struct {
	g.Meta               `orm:"table:manage_device, do:true"`
	Id                   interface{} // 主键
	Name                 interface{} // 设备名称
	Manufacturer         interface{} // 厂商
	Model                interface{} // 型号
	InstallationLocation interface{} // 安装位置
	AreaId               interface{} // 所属区域
	ServerId             interface{} // 设备服务id
	CreatedBy            interface{} // 创建者
	UpdatedBy            interface{} // 更新者
	CreatedAt            *gtime.Time // 创建时间
	UpdatedAt            *gtime.Time // 更新时间
	DeletedAt            *gtime.Time // 删除时间
	Remark               interface{} // 备注
	Extend               *gjson.Json // 扩展信息
	ModelPath            interface{} // 模型路径
}
