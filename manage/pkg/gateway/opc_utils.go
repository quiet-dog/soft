package gateway

import (
	"context"
	"fmt"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

func variantFromValueByNodeID(client *opcua.Client, nodeID *ua.NodeID, value interface{}) (*ua.Variant, error) {
	// 1. 读取节点的数据类型 DataType 属性
	attr, err := client.Node(nodeID).Attribute(context.Background(), ua.AttributeIDDataType)
	if err != nil {
		return nil, fmt.Errorf("读取 DataType 失败: %w", err)
	}

	dtNodeID, ok := attr.Value().(*ua.NodeID)
	if !ok || dtNodeID == nil {
		return nil, fmt.Errorf("DataType 属性不是 NodeID 类型")
	}

	// 2. 根据 DataType NodeID 的数值（IntID）判断类型
	switch dtNodeID.Type() {
	case ua.NodeIDTypeNumeric:
		typeID := dtNodeID.IntID()
		switch typeID {
		case 1: // Boolean
			b, ok := value.(bool)
			if !ok {
				return nil, fmt.Errorf("期望 bool 类型, 实际 %T", value)
			}
			val, err := ua.NewVariant(b)
			if err != nil {
				return nil, fmt.Errorf("数据转换失败")
			}
			return val, nil
		case 6: // Int32
			v, ok := toInt64(value)
			if !ok {
				return nil, fmt.Errorf("期望数字类型可转 int32")
			}
			val, err := ua.NewVariant(int32(v))
			if err != nil {
				return nil, fmt.Errorf("数据转换失败")
			}
			return val, nil
		case 12: // String
			s, ok := value.(string)
			if !ok {
				return nil, fmt.Errorf("期望字符串类型")
			}
			val, err := ua.NewVariant(s)
			if err != nil {
				return nil, fmt.Errorf("数据转换失败")
			}
			return val, nil
		// 这里可以继续扩展其他类型...
		default:
			return nil, fmt.Errorf("不支持的数据类型ID: %d", typeID)
		}
	default:
		return nil, fmt.Errorf("暂不支持 DataType NodeIDType: %v", dtNodeID.Type())
	}
}

// 简单转换辅助（只示意，实际可更完善）
func toInt64(v interface{}) (int64, bool) {
	switch val := v.(type) {
	case int:
		return int64(val), true
	case int32:
		return int64(val), true
	case int64:
		return val, true
	case float32:
		return int64(val), true
	case float64:
		return int64(val), true
	default:
		return 0, false
	}
}
