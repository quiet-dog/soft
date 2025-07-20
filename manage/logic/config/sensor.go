package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/common"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/pkg/hook"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"fmt"

	"github.com/expr-lang/expr"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSensor struct {
	base.BaseService
}

func init() {
	manage.RegisterManageSensor(NewManageSensor())
}

func NewManageSensor() *sSensor {
	return &sSensor{}
}

func (s *sSensor) Model(ctx context.Context) *gdb.Model {
	return dao.ManageSensor.Ctx(ctx).Hook(hook.Bind()).Handler().Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSensor) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageSensorSearch) (res []*res.SensorTableRow, total int, err error) {
	m := s.handleSensorSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sSensor) Save(ctx context.Context, in *req.ManageSensorSave) (id int64, err error) {
	var device *do.ManageSensor
	if err = gconv.Struct(in, &device); err != nil {
		return
	}

	rs, err := s.Model(ctx).Data(device).Insert()
	if utils.IsError(err) {
		return 0, err
	}
	id, err = rs.LastInsertId()
	return
}

func (s *sSensor) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	return
}

func (s *sSensor) Tree(ctx context.Context, req *model.PageListReq, in *req.ManageSensorSearch) (out []*res.AreaTree, err error) {
	out = []*res.AreaTree{}
	items, _, err := s.GetPageListForSearch(ctx, req, in)
	for _, item := range items {
		out = append(out, &res.AreaTree{
			Label:    item.Name,
			Value:    item.Id,
			Children: nil,
			IsLeaf:   true, // Assuming devices are leaf nodes
		})
	}
	return
}

func (s *sSensor) handleSensorSearch(ctx context.Context, in *req.ManageSensorSearch) (query *gdb.Model) {
	query = s.Model(ctx)
	if in == nil {
		return query
	}

	if !g.IsEmpty(in.Name) {
		query = query.WhereLike("name", "%"+in.Name+"%")
	}

	if len(in.DeviceIds) > 0 {
		query = query.WhereIn("device_id", in.DeviceIds)
	}

	return
}

func (s *sSensor) ReadData(ctx context.Context, in *req.ManageSensorReadData) (out *common.TemplateEnv, err error) {

	// opc 数据读取
	if in.Type == gateway.SERVER_OPC && !in.Extend.IsNil() {
		if opcId := in.Extend.Get("id").Int64(); opcId != 0 {
			out, err = NewManageOpc().ReadData(ctx, opcId)
		}
	}

	return
}

func (s *sSensor) TranslateData(ctx context.Context, in *req.ManageSensorTranslate) (out common.Value, err error) {
	out = common.Value{}
	if in.Template == "" {
		return in.Env.Value, nil
	}
	program, err := expr.Compile(in.Template, expr.Env(in.Env.PrepareExprEnv()))
	if err != nil {
		return
	}
	result, err := expr.Run(program, in.Env.PrepareExprEnv())
	// s.ReadInfluxdbFormat(ctx, 48)
	out.Value = result
	return
}

func (s *sSensor) ReadInfluxdbFormat(ctx context.Context, sensorId int64) (out *common.SensorToInfluxdb, err error) {
	fmt.Println("===========")
	out = &common.SensorToInfluxdb{}
	r, err := s.Read(ctx, sensorId)
	if err != nil {
		return
	}
	out.DeviceId = r.DeviceId
	out.SensorId = r.Id
	out.SensorTypeId = r.SensorTypeId
	out.Template = r.Template
	fmt.Println(out.Template, sensorId)
	// os.Exit(0)
	return
}

func (s *sSensor) Read(ctx context.Context, sensorId int64) (sensorInfo *res.SensorInfo, err error) {
	sensorInfo = &res.SensorInfo{}
	err = s.Model(ctx).Where(dao.ManageSensor.Columns().Id, sensorId).Scan(&sensorInfo)
	if utils.IsError(err) {
		return
	}
	return
}
