from opcua import Server
from opcua import ua
import datetime
import time
import random

# 创建服务器
server = Server()
server.set_endpoint("opc.tcp://0.0.0.0:4840/freeopcua/server/")

# 注册命名空间 1（原始命名空间）
uri1 = "http://example.org/opcua"
idx1 = server.register_namespace(uri1)

# 注册命名空间 2（新增命名空间）
uri2 = "http://example.org/opcua/extra"
idx2 = server.register_namespace(uri2)

# 获取 Objects 节点
objects = server.get_objects_node()

# 命名空间1下的 Devices 树
devices = objects.add_object(idx1, "Devices")

# 定义命名空间1下的设备
device_config = {
    "Device1": ["Temperature", "Humidity"],
    "Device2": ["Pressure", "Speed"],
    "Device3": ["Voltage", "Current", "Power"]
}

device_nodes = {}

for device_name, tags in device_config.items():
    device_node = devices.add_object(idx1, device_name)
    device_nodes[device_name] = {}
    for tag in tags:
        var = device_node.add_variable(idx1, tag, 0.0)
        var.set_writable()
        device_nodes[device_name][tag] = var

# 命名空间2下的另一个设备树
extra_devices = objects.add_object(idx2, "ExtraDevice")
extra_var1 = extra_devices.add_variable(idx2, "ExtraSensor1", 0.0)
extra_var2 = extra_devices.add_variable(idx2, "ExtraSensor2", 0.0)
extra_var1.set_writable()
extra_var2.set_writable()

# 启动服务器
server.start()
print("OPC UA Server started at opc.tcp://0.0.0.0:4840/freeopcua/server/")
print(f"Namespace 1 URI: {uri1} (index {idx1})")
print(f"Namespace 2 URI: {uri2} (index {idx2})")

try:
    while True:
        # 更新命名空间1的变量
        for device, tags in device_nodes.items():
            for tag_name, var in tags.items():
                new_value = round(random.uniform(10, 100), 2)
                var.set_value(new_value)
        # 更新命名空间2的变量
        extra_var1.set_value(round(random.uniform(1, 10), 2))
        extra_var2.set_value(round(random.uniform(100, 200), 2))
        print(f"ExtraDevice.ExtraSensor1 = {extra_var1.get_value()}")
        print(f"ExtraDevice.ExtraSensor2 = {extra_var2.get_value()}")

        time.sleep(4)

except KeyboardInterrupt:
    print("Shutting down server...")

finally:
    server.stop()