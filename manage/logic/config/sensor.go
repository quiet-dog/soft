package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/common"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/pkg/expr_template"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/pkg/hook"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"fmt"
	"time"

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

type sensorHook struct{}

func (s *sensorHook) AfterSelectHook(ctx context.Context, in *gdb.HookSelectInput, result *gdb.Result) (err error) {
	for _, item := range *result {
		if !item[dao.ManageSensor.Columns().Id].IsEmpty() {
			t, err := manage.ManageSensorCache().Get(ctx, item[dao.ManageSensor.Columns().Id].Int64())
			item["is_online"] = g.NewVar(true)
			if err != nil {
				item["is_online"] = g.NewVar(false)
			}
			template, err := manage.ManageSensorTemplateCache().Get(ctx, item[dao.ManageSensor.Columns().Id].Int64())
			if err == nil {
				v, _ := template.ToExprValueFloat64(t.Value)
				item["value"] = g.NewVar(v)
			}

		}

		if !item[dao.ManageSensor.Columns().DeviceId].IsEmpty() {
			item["device_name"], _ = dao.ManageDevice.Ctx(ctx).WherePri(item[dao.ManageSensor.Columns().DeviceId].Int64()).Value("name")
		}

		if !item[dao.ManageSensor.Columns().SensorTypeId].IsEmpty() {
			item["sensor_type_name"], _ = dao.ManageSensorType.Ctx(ctx).WherePri(item[dao.ManageSensor.Columns().SensorTypeId].Int64()).Value("name")
		}

	}
	return
}

func (s *sSensor) Model(ctx context.Context) *gdb.Model {
	v := &sensorHook{}
	return dao.ManageSensor.Ctx(ctx).Hook(hook.Bind(v)).Handler().Cache(orm.SetCacheOption(ctx)).OnConflict("id")
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

	if in.Type == gateway.SERVER_MODBUS_TCP && !in.Extend.IsNil() {
		if deviceId := in.Extend.Get("deviceId").Int64(); deviceId != 0 {
			out, err = manage.ManageModbus().TestDataByDeviceId(ctx, deviceId, in)
		}
	}

	return
}

func (s *sSensor) TranslateData(ctx context.Context, in *req.ManageSensorTranslate) (out common.Value, err error) {
	out = common.Value{}
	if in.Template == "" {
		return in.Env.Value, nil
	}
	template := expr_template.ExprTemplate(in.Template)
	out.Value, err = template.ToExprValueFloat64(in.Env.Value)

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
	out.Template = expr_template.ExprTemplate(r.Template)
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

func (s *sSensor) ReadEchart(ctx context.Context, re *model.PageListReq, in *req.ManageInfluxdbOneSensorSearch) (out *res.SensorEchart, err error) {
	out = &res.SensorEchart{}
	info, err := manage.ManageInfluxdb().SearchSensorEchart(ctx, re, in)
	if err != nil {
		return
	}

	out.SensorId = info.SensorId
	out.DeviceId = info.DeviceId
	out.SensorName = info.SensorName
	out.SensorTypeId = info.SensorTypeId
	out.SensorTypeName = info.SensorTypeName
	out.SensorUnit = info.SensorUnit
	for _, v := range info.Rows {
		out.CSeriesData = append(out.CSeriesData, v[fmt.Sprintf("c_%d", in.SensorId)])
		out.ESeriesData = append(out.ESeriesData, v[fmt.Sprintf("e_%d", in.SensorId)])
		// out.XData = append(out.XData, v["time"])
		if tt, ok := v["time"].(time.Time); ok {
			out.XData = append(out.XData, tt.UTC().Format("2006-01-02 15:04:05"))
		}
	}

	return
}
