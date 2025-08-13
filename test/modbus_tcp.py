import asyncio
import logging
import random
from pymodbus.server import StartAsyncTcpServer
from pymodbus.datastore import ModbusSequentialDataBlock, ModbusServerContext, ModbusDeviceContext

# Configure logging
logging.basicConfig()
log = logging.getLogger()
log.setLevel(logging.INFO)

class CallbackDataBlock(ModbusSequentialDataBlock):
    """ 扩展 DataBlock，在写入时触发回调 """
    def setValues(self, address, value):
        super().setValues(address, value)
        if address < 10:
            log.info(f"寄存器写入: 地址={address}, 数据={value}")

async def simulate_random_data(store, interval=1):
    """ 模拟随机数据 """
    while True:
        for device_id, device in store.items():
            for i in range(100):
                if i < 10 :
                    continue
                device.setValues(3, i, [random.randint(0, 100)])  # Holding registers
                device.setValues(4, i, [random.randint(0, 100)])  # Input registers
                device.setValues(2, i, [random.randint(0, 1)])    # Discrete inputs
                device.setValues(1, i, [random.randint(0, 1)])    # Coils
        await asyncio.sleep(interval)

async def run_modbus_server():
    # 每个设备的寄存器块
    def block(): return CallbackDataBlock(0, [0] * 100)

    store = {}
    for device_id in range(10):
        store[device_id] = ModbusDeviceContext(
            di=block(),
            co=block(),
            hr=block(),
            ir=block(),
        )

    context = ModbusServerContext(devices=store, single=False)

    # 启动随机数据模拟
    asyncio.create_task(simulate_random_data(store, interval=2))

    # 启动 TCP 服务器
    address = ("0.0.0.0", 9999)
    log.info(f"启动 Modbus TCP 服务器: {address[0]}:{address[1]}")
    await StartAsyncTcpServer(context=context, address=address)

if __name__ == "__main__":
    asyncio.run(run_modbus_server())
