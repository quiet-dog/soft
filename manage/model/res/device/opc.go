package device

import "github.com/gopcua/opcua/ua"

type OpcTree struct {
	NodeClass    ua.NodeClass `json:"nodeClass" description:"节点类型"`    // 节点类型
	VariableName string       `json:"variableName" description:"变量名称"` // 变量名称
	Namespace    string       `json:"namespaceIndex" description:"命名空间索引"`
	NodeId       string       `json:"nodeId" description:"节点ID"`      // 节点ID
	Children     []*OpcTree   `json:"children" description:"子节点"`     // 子节点
	Type         string       `json:"type" description:"节点类型"`        // 节点类型
	DisplayName  string       `json:"displayName" description:"显示名称"` // 显示名称
	BrowseName   string       `json:"browseName" description:"浏览名称"`  // 浏览名称
}
