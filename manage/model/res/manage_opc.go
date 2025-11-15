package res

import "devinggo/manage/model/base"

type OpcTableRow struct {
	base.BaseTable
	NodeId         string `json:"nodeId" description:"节点ID"`      // 节点ID
	ServerId       int64  `json:"serverId" description:"服务器ID"`   // 服务器ID
	NodeClass      string `json:"nodeClass" description:"节点类型"`   // 节点类型
	Type           string `json:"type" description:"节点类型"`        // 节点类型
	ParentId       int64  `json:"parentId" description:"父节点ID"`   // 父节点ID
	NamespaceIndex string `json:"namespace" description:"命名空间"`   // 命名空间
	DisplayName    string `json:"displayName" description:"显示名称"` // 显示名称
	BrowseName     string `json:"browseName" description:"浏览名称"`  // 浏览名称
}

type OpcTree struct {
	Key         string    `json:"key"`
	Label       string    `json:"label" description:"名称"`         // 名称
	Value       int64     `json:"value" description:"ID"`         // ID
	Children    []OpcTree `json:"children" description:"子节点"`     // 子节点
	IsLeaf      bool      `json:"isLeaf" description:"是否叶子节点"`    // 是否叶子节点
	DisplayName string    `json:"displayName" description:"显示名称"` // 显示名称
	BrowseName  string    `json:"browseName" description:"浏览名称"`  // 浏览名称
	NodeId      string    `json:"nodeId" description:"节点ID"`      // 节点ID
}

type OpcInfo struct {
	OpcTableRow
}
