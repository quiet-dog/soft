from pymodbus.server import StartAsyncTcpServer
from pymodbus.datastore import ModbusSequentialDataBlock, ModbusSlaveContext, ModbusServerContext
import random
import asyncio
import time
import logging

# 设置日志用于调试
logging.basicConfig()
log = logging.getLogger()
log.setLevel(logging.INFO)

# 定义数据存储
def create_datastore():
    # 创建数据块：保持寄存器，地址 1-2
    # 初始化温度=25.0°C (250), 湿度=60.0% (600)
    block = ModbusSequentialDataBlock(1, [250, 600])  # 地址从 1 开始，uint16 值
    store = ModbusSlaveContext(
        hr=block,  # 保持寄存器
        zero_mode=False  # Modbus 地址从 1 开始
    )
    return store

# 更新传感器数据（模拟动态变化）
async def update_sensor_data(context):
    while True:
        # 获取从站上下文（从站 ID=1）
        slave_context = context[1]
        # 模拟温度 (20.0°C 到 30.0°C，放大 10 倍存为 uint16)
        temperature = int(random.uniform(20.0, 30.0) * 10)
        # 模拟湿度 (50.0% 到 80.0%，放大 10 倍存为 uint16)
        humidity = int(random.uniform(50.0, 80.0) * 10)
        
        # 写入寄存器：地址 1=温度，地址 2=湿度
        slave_context.setValues(3, 1, [temperature, humidity])
        
        # 记录当前值
        log.info(f"更新数据: 温度={temperature/10:.1f}°C, 湿度={humidity/10:.1f}%")
        
        # 每 2 秒更新一次
        await asyncio.sleep(2)

# 主函数运行服务器
async def run_server():
    # 创建从站上下文（从站 ID=1）
    store = create_datastore()
    context = ModbusServerContext(slaves={1: store}, single=False)
    
    # 启动数据更新任务
    asyncio.create_task(update_sensor_data(context))
    
    # 启动 Modbus TCP 服务器
    address = ("127.0.0.1", 5020)  # 监听 localhost:5020（避免权限问题）
    log.info(f"启动 Modbus TCP 服务器于 {address[0]}:{address[1]}, 从站 ID=1")
    await StartAsyncTcpServer(context=context, address=address)

# 运行服务器
if __name__ == "__main__":
    try:
        asyncio.run(run_server())
    except KeyboardInterrupt:
        log.info("服务器已停止")