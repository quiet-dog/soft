from opcua import Server
import random
import time

# 1. 创建 OPC UA 服务器
server = Server()
server.set_endpoint("opc.tcp://0.0.0.0:4843/freeopcua/server/")

# 2. 设置命名空间
uri = "http://example.com/sensors"
idx = server.register_namespace(uri)

# 3. 创建对象节点
root = server.get_objects_node()
sensors_node = root.add_object(idx, "Sensors")

# 4. 创建传感器节点
temperature_nodes = []
humidity_nodes = []
pressure_nodes = []

for i in range(1, 6):
    temperature_nodes.append(sensors_node.add_variable(idx, f"Temperature_{i}", 20.0))
    humidity_nodes.append(sensors_node.add_variable(idx, f"Humidity_{i}", 50.0))
    pressure_nodes.append(sensors_node.add_variable(idx, f"Pressure_{i}", 1013.0))

# 5. 设置变量可写
for node in temperature_nodes + humidity_nodes + pressure_nodes:
    node.set_writable()

# 6. 启动服务器
server.start()
print("OPC UA Server started at opc.tcp://0.0.0.0:4840/freeopcua/server/")

try:
    while True:
        # 模拟实时数据更新
        for node in temperature_nodes:
            node.set_value(round(random.uniform(20, 30), 2))
        for node in humidity_nodes:
            node.set_value(round(random.uniform(40, 60), 2))
        for node in pressure_nodes:
            node.set_value(round(random.uniform(1000, 1020), 2))
        
        time.sleep(1)  # 每秒更新一次
finally:
    server.stop()
    print("Server stopped")
