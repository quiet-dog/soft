from pymodbus.client import ModbusTcpClient

# 连接到 Modbus 服务器（设备）
client = ModbusTcpClient('0.0.0.0', port=9999)

# 建立连接
if not client.connect():
    print("连接失败")
    exit(1)

# 读取保持寄存器，起始地址0，读10个寄存器
result = client.read_holding_registers(address=0, count=10, device_id=0)

if not result.isError():
    # result.registers 是一个列表，包含读取到的寄存器值
    print("读取到的寄存器数据:", result.registers)
else:
    print("读取错误:", result)

# 关闭连接
client.close()
