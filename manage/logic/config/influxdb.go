package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/utils"
	"fmt"
	"strings"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/v2/influxdb3"
)

type sInfluxdb struct {
	base.BaseService
}

func init() {
	manage.RegisterManageInfluxdb(NewManageInfluxdb())

}

func NewManageInfluxdb() *sInfluxdb {
	return &sInfluxdb{}
}

func (s *sInfluxdb) Model(ctx context.Context) (*influxdb3.Client, error) {
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host: "http://localhost:8181",
		// Token:    "apiv3_m5pZL1Z_fuVx4oEKwkwSiL5qyIYu3CQrih5394FoDuURdYxPqwtWO3IYiZG06-0AXysYINo_f46Pi5-xDQa-pw",
		Token:    "apiv3_KDfgrll4Hg3VKFpOt5wLOtStjWSmZNeIcW-obG1SYJGc5W2OAZRrH-pXq_5Q-_E7LT0bhKwcMOglg-Ml2J3EJg",
		Database: "DATABASE_NAME",
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *sInfluxdb) SearchSensorDataList(ctx context.Context, req *model.PageListReq, in *req.ManageInfluxdbSearch) (out *res.SensorDataList, err error) {
	out = &res.SensorDataList{}
	fmt.Println(len(in.SensorIds))
	dao.ManageSensor.Ctx(ctx).As("s").Fields("s.device_id", "s.sensor_type_id",
		"s.id as sensor_id", "st.name as sensor_type_name",
		"s.name as sensor_name", "st.unit as sensor_unit").
		LeftJoin("manage_sensor_type st", "st.id = s.sensor_type_id").
		LeftJoin("manage_device d", "d.id = s.device_id").
		WhereIn("s.id", in.SensorIds).Scan(&out)
	out.Rows, out.Total, err = s.SearchTable(ctx, req, in)
	if utils.IsError(err) {
		return nil, err
	}
	return
}

func (s *sInfluxdb) SearchSensorEchart(ctx context.Context, re *model.PageListReq, in *req.ManageInfluxdbOneSensorSearch) (out *res.SensorDataList, err error) {
	out = &res.SensorDataList{}
	dao.ManageSensor.Ctx(ctx).As("s").Fields("s.device_id", "s.sensor_type_id",
		"s.id as sensor_id", "st.name as sensor_type_name",
		"s.name as sensor_name", "st.unit as sensor_unit").
		LeftJoin("manage_sensor_type st", "st.id = s.sensor_type_id").
		LeftJoin("manage_device d", "d.id = s.device_id").
		Where("s.id", in.SensorId).Scan(&out)
	i := &req.ManageInfluxdbSearch{
		DeviceId:  out.DeviceId,
		SensorIds: []int64{in.SensorId},
		BeginTime: in.BeginTime,
		EndTime:   in.EndTime,
		Precision: in.Precision,
	}

	out.Rows, out.Total, err = s.SearchTable(ctx, re, i)
	// if err != nil {
	// 	return
	// }

	// for _, v := range list {
	// 	out.CSeriesData = append(out.CSeriesData, v[fmt.Sprintf("c_%d", in.SensorId)])
	// 	out.ESeriesData = append(out.ESeriesData, v[fmt.Sprintf("e_%d", in.SensorId)])
	// 	out.XData = append(out.XData, v["time"])
	// }

	// if utils.IsError(err) {
	// 	return nil, err
	// }
	return
}

func (s *sInfluxdb) SearchTable(ctx context.Context, req *model.PageListReq, in *req.ManageInfluxdbSearch) (list []map[string]interface{}, total int64, err error) {
	line, totalLine := s.handleInfluxdbSearch(ctx, req, in)
	c, err := s.Model(context.Background())
	if err != nil {
		fmt.Println("Failed to create InfluxDB client:", err)
		return
	}
	res, err := c.Query(ctx, line)
	if err != nil {
		return
	}

	for res.Next() {
		list = append(list, res.Value())
	}
	res, err = c.Query(ctx, totalLine)
	if err != nil {
		return
	}
	for res.Next() {
		total = res.Value()["count(*)"].(int64)
	}

	return

}

func (s *sInfluxdb) Store(ctx context.Context, data gateway.Value, sensorId int64) (err error) {
	c, err := s.Model(context.Background())
	if err != nil {
		fmt.Println("Failed to create InfluxDB client:", err)

		return
	}
	defer c.Close()

	fmt.Println("Writing data to InfluxDB...", sensorId)
	influxdbData, err := manage.ManageSensor().ReadInfluxdbFormat(ctx, sensorId)
	if err != nil {
		fmt.Println("Failed to read InfluxDB format for sensor ID:", sensorId, "Error:", err)

		return
	}
	fmt.Println("====================", influxdbData.Template)

	current := influxdbData.Template.ToValueInfluxdbFloat64(data.Value)

	// line := "1,sensor=2 value=23.5,current=45i"
	line := fmt.Sprintf("t_%d,sensor=s_%d c_%d=%s,e_%d=%s %d",
		influxdbData.DeviceId,
		influxdbData.SensorId,
		influxdbData.SensorId,
		current,
		influxdbData.SensorId,
		influxdbData.Template.ToExprValueInfluxdbFloat64(data.Value),
		data.CreateTime.UnixNano(),
	)
	fmt.Println(line)
	err = c.Write(ctx, []byte(line))
	return
}

func (s *sInfluxdb) StoreDataChannel(ctx context.Context, msg gateway.Msg) (err error) {
	// 存储redis
	manage.ManageSensorDataCache().Store(ctx, msg.Value.ID, msg.Value)

	// 是否报警
	manage.ManageEvent().CheckEvent(ctx, msg.Value.ID, msg.Value)

	{
		// 第三方的接入
	}

	// 存储到influxdb
	err = s.Store(ctx, msg.Value, msg.Value.ID)

	return
}

func (s *sInfluxdb) handleInfluxdbSearch(ctx context.Context, req *model.PageListReq, in *req.ManageInfluxdbSearch) (line string, total string) {
	line = "SELECT * FROM "
	total = "SELECT COUNT(*) FROM "

	// 拼接 measurement
	if in.DeviceId != 0 {
		line += fmt.Sprintf("t_%d", in.DeviceId)
		total += fmt.Sprintf("t_%d", in.DeviceId)
	}
	// where 条件起始
	conditions := ""

	// 传感器 ID
	if len(in.SensorIds) > 0 {
		sensors := make([]string, len(in.SensorIds))
		exprSensors := make([]string, len(in.SensorIds))
		currentSensors := make([]string, len(in.SensorIds))
		for i, sensorId := range in.SensorIds {
			sensors[i] = fmt.Sprintf("'s_%d'", sensorId)
			exprSensors[i] = fmt.Sprintf("'e_%d'", sensorId)
			currentSensors[i] = fmt.Sprintf("'c_%d'", sensorId)
		}
		conditions += fmt.Sprintf("sensor IN (%s)", strings.Join(sensors, ","))
		for _, v := range exprSensors {
			conditions += fmt.Sprintf(" AND %s IS NOT NULL", v)
		}
		for _, v := range currentSensors {
			conditions += fmt.Sprintf(" AND %s IS NOT NULL", v)
		}
	}

	// 时间范围
	if in.BeginTime != 0 && in.EndTime != 0 {
		timeCond := fmt.Sprintf("time >= '%s' and time <= '%s'", time.UnixMilli(in.BeginTime).Format(time.RFC3339), time.UnixMilli(in.EndTime).Format(time.RFC3339))
		if conditions != "" {
			conditions += " AND " + timeCond
		} else {
			conditions = timeCond
		}
	}

	if conditions != "" {
		line += " WHERE " + conditions
		total += " WHERE " + conditions
	}

	if in.Precision != 0 {
		line += fmt.Sprintf(" GROUP BY time(%ds) ", in.Precision)
		line = strings.Replace(line, "SELECT *", `SELECT MEAN("*")`, 1)
		total = strings.Replace(line, "SELECT COUNT(*)", `SELECT COUNT(MEAN("*"))`, 1)
	}

	// 排序
	line += " ORDER BY time DESC"

	// 分页（limit 和 offset）
	if req.PageSize > 0 {
		line += fmt.Sprintf(" LIMIT %d", req.PageSize)
		if req.Page > 1 {
			line += fmt.Sprintf(" OFFSET %d", (req.Page-1)*req.PageSize)
		}
	}
	fmt.Println("InfluxDB Query Line:==========================")
	fmt.Println("InfluxDB Query Line:", line)
	fmt.Println("InfluxDB Total Query Line:", total)
	// os.Exit(0)
	return
}
