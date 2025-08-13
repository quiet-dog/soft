import { BasicCrud } from "@/components/ma-crud/types"
import { reactive, ref } from "vue"
import server from "@/api/manage/server"

export default function useServerHook() {
    const formType = ref('')
    const crud = reactive<BasicCrud>({
        api: server.list,
        showIndex: false,
        // pageLayout: 'fixed',
        rowSelection: { showCheckedAll: true },
        operationColumn: true,
        // operationColumnWidth: 200,
        add: { show: true, api: server.save, auth: ['manage:server:save'] },
        beforeRequest(params) {
        },
        delete: {
            show: true,
            api: server.deletes, auth: ['system:server:delete'],
        },
        edit: { show: true, api: server.update, auth: ['manage:server:update'] },
        beforeEdit: (params) => {
            return true
        }
    })
    const columns = reactive([
        { title: 'ID', dataIndex: 'id', addDisplay: false, editDisplay: false, width: 50, hide: true },
        {
            title: '服务器名称', dataIndex: 'name', formType: 'input', search: true, width: 200, commonRules: [
                { required: true, message: '请输入标签名称' },
            ]
        },
        {
            title: 'ip', dataIndex: 'ip', formType: 'input', search: true, width: 200, commonRules: [
                { required: true, message: '请输入主机地址' },
            ]
        },
        {
            title: '端口', dataIndex: 'port', formType: 'input', search: true, width: 150, addDisplay: true, editDisplay: true, commonRules: [
                { required: true, message: '请输入端口' },
            ]
        },
        {
            title: '类型', dataIndex: 'type', formType: 'select', search: true, width: 150, addDisplay: true, editDisplay: true, commonRules: [
                { required: true, message: '请选择服务器类型' },
            ],
            dict: {
                url: '/manage/server/types',
            }
        },
        {
            title: '采集间隔', dataIndex: 'interval', formType: 'input-number', search: false, width: 150, addDisplay: true, editDisplay: true, commonRules: [
                { required: true, message: '请设置采集间隔' },
            ],
        },
        {
            title: '用户名', dataIndex: 'username', formType: 'input', width: 150, addDisplay: true, editDisplay: true, commonRules: [
            ]
        },
        {
            title: '密码', dataIndex: 'password', formType: 'input-password', width: 150, addDisplay: true, editDisplay: true, commonRules: [
            ]
        },
        {
            title: '在线状态', dataIndex: 'isOnline', addDisplay: false, editDisplay: false
        },
        {
            title: '备注', dataIndex: 'remark', formType: 'textarea', width: 200, addDisplay: true, editDisplay: true
        },
    ])
    const crudRef = ref()
    return {
        crud,
        columns,
        crudRef,
        formType
    }
}