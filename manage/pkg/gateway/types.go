package gateway

type Config struct {
	DeviceId int64  `json:"deviceId"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Type     string `json:"type"`
}

const (
	SERVER_OPC                 = "opc"
	SERVER_MODBUS_TCP          = "modbus_tcp"
	SERVER_MODBUS_RTU          = "modbus_rtu"
	SERVER_MODBUS_RTU_OVER_TCP = "modbus_rtu_over_tcp"
	SERVER_MQTT                = "mqtt"
)
