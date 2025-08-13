// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageDevice is the golang structure for table manage_device.
type ManageDevice struct {
	Id                   int64       `json:"id"                   orm:"id"                    description:"主键"`     // 主键
	Name                 string      `json:"name"                 orm:"name"                  description:"设备名称"`   // 设备名称
	Manufacturer         string      `json:"manufacturer"         orm:"manufacturer"          description:"厂商"`     // 厂商
	Model                string      `json:"model"                orm:"model"                 description:"型号"`     // 型号
	InstallationLocation string      `json:"installationLocation" orm:"installation_location" description:"安装位置"`   // 安装位置
	AreaId               int64       `json:"areaId"               orm:"area_id"               description:"所属区域"`   // 所属区域
	ServerId             int64       `json:"serverId"             orm:"server_id"             description:"设备服务id"` // 设备服务id
	CreatedBy            int64       `json:"createdBy"            orm:"created_by"            description:"创建者"`    // 创建者
	UpdatedBy            int64       `json:"updatedBy"            orm:"updated_by"            description:"更新者"`    // 更新者
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"            description:"创建时间"`   // 创建时间
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"            description:"更新时间"`   // 更新时间
	DeletedAt            *gtime.Time `json:"deletedAt"            orm:"deleted_at"            description:"删除时间"`   // 删除时间
	Remark               string      `json:"remark"               orm:"remark"                description:"备注"`     // 备注
	Extend               *gjson.Json `json:"extend"               orm:"extend"                description:"扩展信息"`   // 扩展信息
	ModelPath            string      `json:"modelPath"            orm:"model_path"            description:"模型路径"`   // 模型路径
}
