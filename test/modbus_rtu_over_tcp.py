import asyncio
import threading, random, time
from pymodbus.server.async_io import StartTcpServer
from pymodbus.datastore import ModbusSparseDataBlock, ModbusServerContext
from pymodbus.transaction import ModbusRtuFramer

# 自定义数据块
class CustomDataBlock(ModbusSparseDataBlock):
    pass

# 初始化寄存器: 0=温度, 1=湿度
block = CustomDataBlock({0: 25, 1: 60})
context = ModbusServerContext({1: block})

# 定时更新寄存器线程
def update_values():
    while True:
        temp = random.randint(20, 30)
        hum = random.randint(40, 70)
        block.setValues(3, {0: temp, 1: hum})
        print(f"更新温湿度: 温度={temp}°C, 湿度={hum}%")
        time.sleep(5)

threading.Thread(target=update_values, daemon=True).start()

# AsyncIO 启动 RTU over TCP 服务器
async def run_server():
    await StartTcpServer(context, framer=ModbusRtuFramer, address=("0.0.0.0", 5021))

if __name__ == "__main__":
    print("启动 RTU over TCP 服务器，端口 5020")
    asyncio.run(run_server())
