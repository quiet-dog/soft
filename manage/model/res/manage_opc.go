package res

import "devinggo/manage/model/base"

type OpcTableRow struct {
	base.BaseTable
	NodeId         string `json:"nodeId" description:"节点ID"`    // 节点ID
	ServerId       int64  `json:"serverId" description:"服务器ID"` // 服务器ID
	NodeClass      string `json:"nodeClass" description:"节点类型"` // 节点类型
	Type           string `json:"type" description:"节点类型"`      // 节点类型
	ParentId       int64  `json:"parentId" description:"父节点ID"` // 父节点ID
	NamespaceIndex string `json:"namespace" description:"命名空间"` // 命名空间
}

type OpcTree struct {
	AreaTree
}

type OpcInfo struct {
	OpcTableRow
}
