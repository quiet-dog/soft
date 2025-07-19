import sensorType from "@/api/manage/sensorType"
import { BasicColumn, BasicCrud } from "@/components/ma-crud/types"
import { reactive, ref } from "vue"

export function useSensorTypeHook() {
    const crudRef = ref()
    const crud = reactive<BasicCrud>({
        api: sensorType.list,
        showIndex: false,
        rowSelection: { showCheckedAll: true },
        operationColumn: true,
        add: { show: true, api: sensorType.save, auth: ['manage:sensorType:save'] },
        beforeRequest(params) {
            // 可以在这里处理请求参数
        },
        delete: {
            show: true,
            api: sensorType.deletes, auth: ['manage:sensorType:delete'],
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
            title: '单位', dataIndex: 'unit', formType: 'input', search: true, width: 150, addDisplay: true, editDisplay: true, commonRules: [
                { required: true, message: '请输入单位' },
            ]
        },
        {
            title: '备注', dataIndex: 'remark', formType: 'textarea', width: 200, addDisplay: true, editDisplay: true
        }
    ])
    return {
        crud,
        columns,
        crudRef
    }
}