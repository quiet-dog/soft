import asyncio
import logging
import random
from pymodbus.server import StartAsyncTcpServer
from pymodbus.datastore import ModbusSequentialDataBlock, ModbusServerContext, ModbusDeviceContext

# Configure logging
logging.basicConfig()
log = logging.getLogger()
log.setLevel(logging.INFO)

async def simulate_random_data(store, interval=1):
    """ Simulate random data for each device in the store """
    while True:
        for device_id, device in store.items():
            # Simulate random values for each register and coil
            # Random values for holding registers (hr), input registers (ir), discrete inputs (di), and coils (co)
            for i in range(100):  # Assuming 100 registers or coils
                device.setValues(3, i, [random.randint(0, 100)])  # Holding registers (hr) range 0-65535
                device.setValues(4, i, [random.randint(0, 100)])  # Input registers (ir) range 0-65535
                device.setValues(2, i, [random.randint(0, 1)])  # Discrete inputs (di) range 0-1
                device.setValues(1, i, [random.randint(0, 1)])  # Coils (co) range 0-1

        await asyncio.sleep(interval)  # Update every 'interval' seconds

async def run_modbus_server():
    # Create a simple data store for a range of devices
    block = lambda: ModbusSequentialDataBlock(0, [0] * 100)  # 100 registers/coils
    store = {}
    for device_id in range(10):  # 10 devices
        store[device_id] = ModbusDeviceContext(
            di=block(),
            co=block(),
            hr=block(),
            ir=block(),
        )

    # Create server context
    context = ModbusServerContext(devices=store, single=False)

    # Start random data simulation task
    asyncio.create_task(simulate_random_data(store, interval=2))  # Update every 2 seconds

    # Start TCP server
    address = ("0.0.0.0", 9999)  # Listen on all interfaces, port 9999
    log.info(f"Starting Modbus TCP server on {address[0]}:{address[1]}")
    await StartAsyncTcpServer(context=context, address=address)

if __name__ == "__main__":
    asyncio.run(run_modbus_server())
