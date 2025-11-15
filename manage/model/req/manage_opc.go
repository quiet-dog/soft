package req

type OpcTreeReq struct {
	ServerId int64 `json:"serverId" description:"OPC服务器ID"` // OPC服务器ID
	ParentId int64 `json:"parentId" description:"父节点ID"`    // 父节点ID
}

type OpcReadByServer struct {
	ServerId int64  `json:"serverId" description:"服务Id"`
	NodeId   string `json:"nodeId" description:"节点Id"`
}
