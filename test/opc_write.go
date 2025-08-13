import asyncio
from asyncua import Client, ua

async def main():
    url = "opc.tcp://localhost:4840/freeopcua/server/"
    try:
        async with Client(url=url) as client:
            # 检查服务器连接
            print(f"尝试连接服务器: {url}")
            await client.connect()
            print("服务器连接成功")
            
            # 等待服务器初始化
            await asyncio.sleep(1)  # 延迟 1 秒确保地址空间加载
            
            # 获取命名空间索引
            uri = "http://examples.freeopcua.github.io"
            idx = await client.get_namespace_index(uri)
            print(f"命名空间索引: {idx}")

            # 获取 MyObject 节点
            myobj_node = client.get_node(f"ns={idx};s=MyObject")
            try:
                obj_name = await myobj_node.read_display_name()
                print(f"找到 MyObject 节点: {obj_name.Text}")
                
                # 列出 MyObject 的子节点
                children = await myobj_node.get_children()
                for child in children:
                    child_name = await child.read_browse_name()
                    print(f"MyObject 下子节点: {child_name}")
            except ua.UaStatusCodeError as e:
                print(f"无法访问 MyObject: {e}")
                return

            # 获取 MyVariable 节点
            node_id = f"ns={idx};s=MyObject.MyVariable"
            node = client.get_node(node_id)
            print(f"尝试访问节点: {node_id}")

            # 检查节点是否存在
            try:
                display_name = await node.read_display_name()
                print(f"找到节点: {display_name.Text}")
            except ua.UaStatusCodeError as e:
                print(f"节点不存在: {e}")
                # 列出 Objects 下的节点
                objects_node = client.nodes.objects
                children = await objects_node.get_children()
                for child in children:
                    child_name = await child.read_browse_name()
                    print(f"Objects 下节点: {child_name}")
                return

            # 检查节点权限
            access_level = await node.read_attribute(ua.AttributeIds.UserAccessLevel)
            print(f"节点权限: {access_level}")

            # 写入数据
            new_value = 42.0
            await node.write_value(new_value)
            print(f"写入值 {new_value} 到 MyVariable")

            # 读取验证
            value = await node.read_value()
            print(f"从 MyVariable 读取值: {value}")

    except ua.UaStatusCodeError as e:
        print(f"操作失败: {e}")
    except Exception as e:
        print(f"意外错误: {e}")

if __name__ == "__main__":
    asyncio.run(main())