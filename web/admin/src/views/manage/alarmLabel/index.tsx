import { BasicColumn, BasicCrud } from "@/components/ma-crud/types";
import { h, reactive, ref } from "vue";
import alarmLabel from "@/api/manage/alarmLabel";
import { Tag } from "@arco-design/web-vue";

export default function useAlarmLabelHook() {
    const crud = reactive<BasicCrud>({
        api: alarmLabel.list,
        showIndex: false,
        // pageLayout: 'fixed',
        rowSelection: { showCheckedAll: true },
        operationColumn: true,
        // operationColumnWidth: 200,
        add: { show: true, api: alarmLabel.save, auth: ['manage:alarmLabel:save'] },
        beforeRequest(params) {
        },
        delete: {
            show: true,
            api: alarmLabel.deletes, auth: ['system:alarmLabel:delete'],
        },
    })
    const columns = reactive<BasicColumn[]>([
        { title: 'ID', dataIndex: 'id', addDisplay: false, editDisplay: false, width: 50, hide: true },
        {
            title: '标签名称', dataIndex: 'name', formType: 'input', search: true, width: 200, commonRules: [
                { required: true, message: '请输入标签名称' },
            ]
        },
        {
            title: '标签等级', dataIndex: 'level', search: true, width: 150, addDisplay: true, editDisplay: true, commonRules: [
                { required: true, message: '请输入标签等级' },
            ], customRender: ({ record }) => {
                return h(Tag, {
                    color: record.color == null ? 'green' : record.color,
                }, record.level)
            }
        },
        {
            title: '通知类型', dataIndex: 'type', search: true, width: 180, editDisabled: true,
            commonRules: [{ required: true, message: '公告类型必选' }], formType: 'radio',
            dict: { name: 'backend_notice_type', props: { label: 'title', value: 'key' }, translation: true },
            addDefaultValue: 1
        },
        {
            title: '标签颜色', dataIndex: 'color', formType: "color-picker", width: 150, addDisplay: true, editDisplay: true, commonRules: [
                { required: true, message: '请选择标签颜色' },
            ]
        },
        { title: '备注', dataIndex: 'remark', formType: 'textarea', width: 200, addDisplay: true, editDisplay: true },
    ])
    const crudRef = ref()
    return {
        crud,
        columns,
        crudRef
    }
}