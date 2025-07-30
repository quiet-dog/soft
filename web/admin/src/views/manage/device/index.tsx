import { BasicColumn, BasicCrud } from "@/components/ma-crud/types"
import { defineAsyncComponent, reactive, ref } from "vue"
import device from "@/api/manage/device"
import { getTreeAreaChildren } from "@/views/manage/area/index"
import { TreeLeaf } from "@/api/manage/base"
import server from "@/api/manage/server"
import { ServerTreeLeaf } from "@/api/manage/server/types"


export async function getAsyncServerConfigComponent(serverId: number) {
    const res = await server.read(serverId)
    if (res.data?.type === 'opc') {
        return defineAsyncComponent(() => {
            return import("@/package/device/opc/index.vue")
        })
    }

    if (res.data?.type.includes('modbus')) {
        return defineAsyncComponent(() => {
            return import("@/package/device/modbus/index.vue")
        })
    }

}

export function useDeviceHook() {
    const crud = reactive<BasicCrud>({
        api: device.list,
        showIndex: false,
        rowSelection: { showCheckedAll: true },
        operationColumn: true,
        add: { show: true, api: device.save, auth: ['manage:device:save'] },
        beforeRequest(params) {
            // 可以在这里处理请求参数
        },
        delete: {
            show: true,
            api: device.deletes, auth: ['manage:device:delete'],
        },
        beforeOpenAdd: (params) => {
            console.log('beforeAdd', params)
            loadServerOptions()
            return true
        }
    })


    const columns = reactive<BasicColumn[]>([
        { title: 'ID', dataIndex: 'id', addDisplay: false, editDisplay: false, width: 50, hide: true },
        {
            title: '设备名称', dataIndex: 'name', formType: 'input', search: true, width: 200, commonRules: [
                { required: true, message: '请输入设备名称' },
            ]
        },
        {
            title: '所属区域', hide: true, dataIndex: 'areaId', formType: 'tree-select', search: false, width: 200, addDisplay: true, editDisplay: true, commonRules: [
                {
                    required: true,
                    message: '请选择所属区域',
                },
                {
                    required: true,
                    validator: (value, callback) => {
                        if (Number(value) > 0) {
                            callback()
                        } else {
                            callback('请选择所属区域')
                        }
                    }
                }
            ],
            dict: {
                async data() {
                    const list = await getTreeAreaChildren(0)
                    return list
                }
            },
            loadMore(option) {
                return new Promise(async (resolve) => {
                    const result = await getTreeAreaChildren(option.value)
                    option.children = result
                    resolve(result)
                })
            },
        },
        {
            title: '所属服务', hide: true, dataIndex: 'serverId', formType: 'select', search: false, width: 200, addDisplay: true, editDisplay: true, commonRules: [
                {
                    required: true,
                    message: '请选择所属服务',
                }
            ],
        },
        {
            title: '服务名称', dataIndex: 'serverName', formType: 'input', width: 200, addDisplay: false, editDisplay: false,
        },
        {
            title: '区域名称', dataIndex: 'areaName', formType: 'input', width: 200, addDisplay: false, editDisplay: false,
        },
        {
            title: '厂商', dataIndex: 'manufacturer', formType: 'input', width: 150, addDisplay: true, editDisplay: true,
            commonRules: [
                { required: true, message: '请输入厂商名称' },
            ]
        },
        {
            title: '型号', dataIndex: 'model', formType: 'input', width: 150, addDisplay: true, editDisplay: true,
            commonRules: [
                { required: true, message: '请输入设备型号' },
            ]
        },
        {
            title: '安装位置', dataIndex: 'installationLocation', formType: 'input', width: 200, addDisplay: true, editDisplay: true,
            commonRules: [
                { required: true, message: '请输入安装位置' },
            ]
        },
        {
            title: '备注', dataIndex: 'remark', formType: 'textarea', width: 200, addDisplay: true, editDisplay: true,
        },
    ])

    const serverOptions = ref<ServerTreeLeaf>({
        items: [],
        pageInfo: {
            name: '',
            ip: '',
            port: '',
            type: '',
            remark: '',
            username: '',
            password: '',
            page: 0,
            pageSize: 10
        }
    })


    const asyncComponent = ref()
    const asyncServerId = ref(0)
    const changeServer = async (value: number) => {
        console.log('changeServer', value)
        asyncComponent.value = await getAsyncServerConfigComponent(value)
        asyncServerId.value = value
        crudRef.value.getFormData().serverId = value
    }

    const loadServerOptions = () => {
        serverOptions.value.pageInfo.page!++
        server.tree(serverOptions.value.pageInfo).then(res => {
            serverOptions.value.items.push(...res.data!);
            if (res.data!.length === 0) {
                serverOptions.value.pageInfo.page!--
            }
        })
    }

    const crudRef = ref()


    const changeExtend = (extend: any) => {
        console.log('changeExtend', extend)
        crudRef.value.getFormData().extend = JSON.stringify(extend)
    }

    return {
        crud,
        columns,
        crudRef,
        serverOptions,
        loadServerOptions,
        changeServer,
        asyncComponent,
        asyncServerId,
        changeExtend
    }
}