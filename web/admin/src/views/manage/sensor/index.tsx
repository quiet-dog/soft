import device from "@/api/manage/device"
import { DeviceTreeLeaf } from "@/api/manage/device/types"
import sensor from "@/api/manage/sensor"
import { SensorTreeLeaf } from "@/api/manage/sensor/types"
import sensorType from "@/api/manage/sensorType"
import { SensorTypeTreeLeaf } from "@/api/manage/sensorType/types"
import { BasicColumn, BasicCrud } from "@/components/ma-crud/types"
import { defineAsyncComponent, reactive, ref } from "vue"


export async function getAsyncDeviceConfigComponent(deviceId: number) {
    const res = await device.read(deviceId)
    if (res.data?.server?.type === 'opc') {
        console.log('加载OPC设备配置组件')
        return defineAsyncComponent(() => {
            return import("@/package/sensor/opc/index.vue")
        })
    }

}
export function useSensorHook() {
    const crudRef = ref()
    const crud = reactive<BasicCrud>({
        api: sensor.list,
        showIndex: false,
        rowSelection: { showCheckedAll: true },
        operationColumn: true,
        add: { show: true, api: sensor.save, auth: ['manage:sensor:save'] },
        beforeRequest(params) {
            // 可以在这里处理请求参数
        },
        delete: {
            show: true,
            api: sensor.deletes, auth: ['manage:sensor:delete'],
        },
        beforeOpenAdd: (params) => {
            loadSensorOptions()
            loadDeviceOptions()
            return true
        },
    })

    const columns = reactive<BasicColumn[]>([
        { title: 'ID', dataIndex: 'id', addDisplay: false, editDisplay: false, width: 50, hide: true },
        {
            title: '传感器名称', dataIndex: 'name', formType: 'input', search: true, width: 200, commonRules: [
                { required: true, message: '请输入传感器名称' },
            ]
        },
        {
            title: '所属设备', dataIndex: 'deviceId', formType: 'select', search: true, width: 200, addDisplay: true, editDisplay: true, commonRules: [
                { required: true, message: '请选择所属设备' },
            ]
        },
        {
            title: '类型', dataIndex: 'sensorTypeId', formType: 'select', search: false, width: 150, addDisplay: true, editDisplay: true, commonRules: [
                { required: true, message: '请选择传感器类型' },
            ]
        },
        {
            title: '标记', dataIndex: 'extend', search: false, width: 150, addDisplay: false, editDisplay: false,
        },
        {
            title: '备注', dataIndex: 'remark', formType: 'textarea', width: 200, addDisplay: true, editDisplay: true
        },
        {
            title:"数据列",dataIndex:"view",addDisplay: false, editDisplay: false,width:100
        }
    ])

    const sensorTypeOptions = ref<SensorTypeTreeLeaf>({
        pageInfo: {
            pageSize: 10,
            page: 0,
            name: '',
            unit: '',
        },
        items: [],
    })

    const changeSensorType = (value: number) => {
        crudRef.value.getFormData().sensorTypeId = value
    }

    const loadSensorOptions = () => {
        sensorTypeOptions.value.pageInfo.page!++
        sensorType.tree(sensorTypeOptions.value.pageInfo).then(res => {
            sensorTypeOptions.value.items.push(...res.data!);
            if (res.data!.length === 0) {
                sensorTypeOptions.value.pageInfo.page!--
            }
        })
    }


    const deviceOptions = ref<DeviceTreeLeaf>({
        pageInfo: {
            pageSize: 10,
            page: 0,
            name: '',
            manufacturer: '',
            model: '',
            installationLocation: '',
            areaId: 0,
            serverId: 0,
        },
        items: [],
    })

    const loadDeviceOptions = () => {
        deviceOptions.value.pageInfo.page!++
        device.tree(deviceOptions.value.pageInfo).then(res => {
            deviceOptions.value.items.push(...res.data!);
            if (res.data!.length === 0) {
                deviceOptions.value.pageInfo.page!--
            }
        })
    }

    const asyncComponent = ref()
    const asyncDeviceId = ref(0)
    const changeDevice = async (value: number) => {
        asyncComponent.value = await getAsyncDeviceConfigComponent(value)
        asyncDeviceId.value = value
        crudRef.value.getFormData().deviceId = value
    }

    const visible = ref(false)


    // 动态更改扩展信息
    const templateVisable = ref(false)
    const templateInfo = ref({
        type: "",
        extend: ""
    })
    const changeExtend = (extend: any) => {
        crudRef.value.getFormData().extend = JSON.stringify(extend)
        device.read(asyncDeviceId.value).then(res => {
            templateInfo.value.type = res.data?.server?.type!
            templateInfo.value.extend = JSON.stringify(extend)
            templateVisable.value = true
        })
    }

    const changeTemplate = (extend: any) => {
        crudRef.value.getFormData().template = extend
    }


    return {
        crud,
        columns,
        crudRef,
        sensorTypeOptions,
        changeSensorType,
        loadSensorOptions,
        deviceOptions,
        loadDeviceOptions,
        changeDevice,
        asyncComponent,
        visible,
        asyncDeviceId,
        changeExtend,
        templateVisable,
        templateInfo,
        changeTemplate
    }
}