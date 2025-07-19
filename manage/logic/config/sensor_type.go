package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/pkg/hook"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSensorType struct {
	base.BaseService
}

func init() {
	manage.RegisterManageSensorType(NewManageSensorType())
}

func NewManageSensorType() *sSensorType {
	return &sSensorType{}
}

func (s *sSensorType) Model(ctx context.Context) *gdb.Model {
	return dao.ManageSensorType.Ctx(ctx).Hook(hook.Bind()).Handler().Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSensorType) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageSensorTypeSearch) (res []*res.SensorTypeTableRow, total int, err error) {
	m := s.handleSensorTypeSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sSensorType) Save(ctx context.Context, in *req.ManageSensorTypeSave) (id int64, err error) {
	var sensorType *do.ManageSensorType
	if err = gconv.Struct(in, &sensorType); err != nil {
		return
	}

	rs, err := s.Model(ctx).Data(sensorType).Insert()
	if utils.IsError(err) {
		return 0, err
	}
	id, err = rs.LastInsertId()
	return
}

func (s *sSensorType) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	return
}

func (s *sSensorType) Tree(ctx context.Context, req *model.PageListReq, in *req.ManageSensorTypeSearch) (out []*res.AreaTree, err error) {
	out = []*res.AreaTree{}
	items, _, err := s.GetPageListForSearch(ctx, req, in)
	for _, item := range items {
		out = append(out, &res.AreaTree{
			Label:    item.Name,
			Value:    item.Id,
			Children: nil,
			IsLeaf:   true, // Assuming sensorTypes are leaf nodes
		})
	}
	return
}

func (s *sSensorType) handleSensorTypeSearch(ctx context.Context, in *req.ManageSensorTypeSearch) (query *gdb.Model) {
	query = s.Model(ctx)
	if in == nil {
		return query
	}

	if !g.IsEmpty(in.Name) {
		query = query.WhereLike("name", "%"+in.Name+"%")
	}

	return
}
