package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/common"
	"devinggo/manage/model/entity"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/model/res/device"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

type sOpc struct {
	base.BaseService
}

func init() {
	manage.RegisterManageOpc(NewManageOpc())
}

func NewManageOpc() *sOpc {
	return &sOpc{}
}

func (s *sOpc) Model(ctx context.Context) *gdb.Model {
	return dao.ManageOpc.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func browseNamespaceTree(c *opcua.Client, nodeID *ua.NodeID, visited map[string]bool, targetNS uint16, treeName string) *device.OpcTree {
	key := nodeID.String()
	if visited[key] {
		return nil
	}
	visited[key] = true

	root := &device.OpcTree{
		NodeClass:    0,
		VariableName: treeName,
		Namespace:    fmt.Sprintf("%d", targetNS),
		NodeId:       nodeID.String(),
		Type:         nodeID.Type().String(),
		Children:     []*device.OpcTree{},
	}

	req := &ua.BrowseRequest{
		View: &ua.ViewDescription{
			ViewID: ua.NewTwoByteNodeID(0),
		},
		NodesToBrowse: []*ua.BrowseDescription{
			{
				NodeID:          nodeID,
				BrowseDirection: ua.BrowseDirectionForward,
				ReferenceTypeID: ua.NewNumericNodeID(0, 0),
				IncludeSubtypes: true,
				NodeClassMask:   0,
				ResultMask:      uint32(ua.BrowseResultMaskAll),
			},
		},
	}

	resp, err := c.Browse(context.Background(), req)
	if err != nil || len(resp.Results) == 0 || resp.Results[0].StatusCode != ua.StatusOK {
		log.Printf("Browse error at node %s: %v", nodeID, err)
		return root // 返回空的 root 也可以
	}

	for _, ref := range resp.Results[0].References {
		childNode := ref.NodeID.NodeID
		childKey := childNode.String()
		if visited[childKey] {
			continue
		}

		ns := childNode.Namespace()
		if ns == targetNS {
			child := &device.OpcTree{
				NodeClass:    ref.NodeClass,
				VariableName: ref.BrowseName.Name,
				Namespace:    fmt.Sprintf("%d", ns),
				NodeId:       childNode.String(),
				Children:     []*device.OpcTree{},
			}
			if childNode != nil {
				child.Type = childNode.Type().String()
			}
			child.Children = browseNamespaceTree(c, childNode, visited, targetNS, ref.DisplayName.Text).Children
			root.Children = append(root.Children, child)
		}
	}

	return root
}

func (s *sOpc) saveTreeRecursive(ctx context.Context, tree *device.OpcTree, serverId, parentId int64, saveFunc func(*device.OpcTree, int64, int64) (int64, error)) error {
	newId, err := saveFunc(tree, serverId, parentId)
	if err != nil {
		return fmt.Errorf("保存节点失败: %w", err)
	}

	for _, child := range tree.Children {
		if err := s.saveTreeRecursive(ctx, child, serverId, newId, saveFunc); err != nil {
			return err
		}
	}
	return nil
}

func (s *sOpc) InitOpc(ctx context.Context, serverId int64) (result []*device.OpcTree, err error) {
	var server *res.ServerTableRow
	if err = dao.ManageServer.Ctx(ctx).WherePri(serverId).Scan(&server); err != nil {
		return
	}

	endpoint := fmt.Sprintf("opc.tcp://%s:%s", server.Ip, server.Port)
	c, err := opcua.NewClient(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create OPC UA client: %w", err)
	}

	if err = c.Connect(ctx); err != nil {
		return nil, fmt.Errorf("连接opc服务器错误: %w", err)
	}

	defer c.Close(context.Background())

	// 获取全部节点

	// 先打印 NamespaceArray，确认索引和URI
	nsNode := c.Node(ua.NewNumericNodeID(0, 2255))
	val, err := nsNode.Value(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	namespaces, ok := val.Value().([]string)
	if !ok {
		return nil, fmt.Errorf("expected []string but got %T", val)
	}
	fmt.Println("Namespaces:")

	for i, ns := range namespaces {
		fmt.Printf("  %d: %s\n", i, ns)
		objectsNode := ua.NewNumericNodeID(0, 85)
		visited := make(map[string]bool)
		nodeTree := browseNamespaceTree(c, objectsNode, visited, uint16(i), objectsNode.String())
		result = append(result, nodeTree)
		err = s.saveTreeRecursive(ctx, nodeTree, serverId, 0, func(ot *device.OpcTree, i1, i2 int64) (int64, error) {
			id, err := s.Save(ctx, ot, i1, i2)
			return id, err
		})
		if err != nil {
			return result, fmt.Errorf("保存节点树失败: %w", err)
		}
	}

	return
}

func (s *sOpc) Save(ctx context.Context, in *device.OpcTree, serverId int64, parentId int64) (id int64, err error) {
	var device entity.ManageOpc
	if err = gconv.Struct(in, &device); err != nil {
		return
	}

	// 查询是否存在同名的节点
	s.Model(ctx).Where(dao.ManageOpc.Columns().ServerId, serverId).
		Where(dao.ManageOpc.Columns().NodeId, in.NodeId).Scan(&device)

	device.NodeClass = in.NodeClass.String()
	device.ServerId = serverId
	device.NodeId = in.NodeId
	device.Type = in.Type
	device.ParentId = parentId
	device.NamespacesIndex = int(gconv.Int(in.Namespace))
	device.BrowseName = in.VariableName

	rs, err := s.Model(ctx).Data(device).Save()
	if utils.IsError(err) {
		return 0, err
	}

	if device.Id > 0 {
		return device.Id, nil
	}

	return rs.LastInsertId()
}

func (s *sOpc) Read(ctx context.Context, opcId int64) (opcInfo *res.OpcInfo, err error) {
	opcInfo = &res.OpcInfo{}
	err = s.Model(ctx).Where(dao.ManageOpc.Columns().Id, opcId).Scan(&opcInfo)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sOpc) Tree(ctx context.Context, in *req.OpcTreeReq) (rs []*res.OpcTree, err error) {
	data := []*res.OpcTableRow{}
	query := s.Model(ctx).Where(dao.ManageOpc.Columns().ServerId, in.ServerId)
	if in.ParentId > 0 {
		query = query.Where(dao.ManageOpc.Columns().ParentId, in.ParentId)
	} else {
		query = query.Where(dao.ManageOpc.Columns().ParentId, 0)
	}
	err = query.Scan(&data)
	if utils.IsError(err) {
		return nil, err
	}

	for _, item := range data {

		rs = append(rs, &res.OpcTree{
			AreaTree: res.AreaTree{
				Label:    item.NodeId,
				Value:    item.Id,
				Children: nil,
				IsLeaf:   s.isLeafNode(item.Id), // Assuming these are not leaf nodes
			},
		})
	}
	return
}

// 判断是否是叶子节点
func (s *sOpc) isLeafNode(id int64) bool {
	// 这里可以根据实际情况判断是否是叶子节点
	// 例如，如果没有子节点，则认为是叶子节点
	count, err := s.Model(context.Background()).Where(dao.ManageOpc.Columns().ParentId, id).Count()
	if err != nil {
		log.Printf("Error checking leaf node status for ID %d: %v", id, err)
		return false
	}
	return count == 0
}

// 根据opc_id 读取数据的类型
func (s *sOpc) ReadData(ctx context.Context, opcId int64) (rs *common.TemplateEnv, err error) {
	r, err := s.Read(context.Background(), opcId)
	if err != nil {
		return
	}

	server, err := NewManageServer().Read(context.Background(), r.ServerId)
	if err != nil {
		return
	}

	c, err := opcua.NewClient(fmt.Sprintf("opc.tcp://%s:%s", server.Ip, server.Port), opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err != nil {
		return
	}

	if err = c.Connect(ctx); err != nil {
		return
	}
	defer c.Close(ctx)

	id, err := ua.ParseNodeID(r.NodeId)
	if err != nil {
		return
	}

	req := &ua.ReadRequest{
		MaxAge: 2000,
		NodesToRead: []*ua.ReadValueID{
			{NodeID: id},
		},
		TimestampsToReturn: ua.TimestampsToReturnBoth,
	}

	var resp *ua.ReadResponse
	for {
		resp, err = c.Read(ctx, req)
		if err == nil {
			break
		}

		// Following switch contains known errors that can be retried by the user.
		// Best practice is to do it on read operations.
		switch {
		case err == io.EOF && c.State() != opcua.Closed:
			// has to be retried unless user closed the connection
			time.After(1 * time.Second)
			continue

		case errors.Is(err, ua.StatusBadSessionIDInvalid):
			// Session is not activated has to be retried. Session will be recreated internally.
			time.After(1 * time.Second)
			continue

		case errors.Is(err, ua.StatusBadSessionNotActivated):
			// Session is invalid has to be retried. Session will be recreated internally.
			time.After(1 * time.Second)
			continue

		case errors.Is(err, ua.StatusBadSecureChannelIDInvalid):
			// secure channel will be recreated internally.
			time.After(1 * time.Second)
			continue

		default:
			log.Fatalf("Read failed: %s", err)
		}
	}

	if resp != nil && resp.Results[0].Status != ua.StatusOK {
		log.Fatalf("Status not OK: %v", resp.Results[0].Status)
	}
	rs = &common.TemplateEnv{}
	rs.Value = common.Value{
		Value: resp.Results[0].Value.Value(),
	}
	rs.Type = resp.Results[0].Value.Type().String()
	rs.CreateTime = resp.Results[0].SourceTimestamp

	return
}
