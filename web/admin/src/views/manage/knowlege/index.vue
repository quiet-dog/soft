<template>
    <div class="ma-content-block lg:flex justify-between p-4">
        <!-- CRUD 组件 -->
        <ma-crud :options="options" :columns="columns" ref="crudRef">

            <template #sort="{ record }">
                <a-input-number :default-value="record.sort" mode="button" :min="0" :max="1000" />
            </template>

            <template #status="{ record }">
                <a-switch :checked-value="1" unchecked-value="2" @change="switchStatus($event, record)"
                    :default-checked="record.status == 1" />
            </template>


            <!-- 操作前置扩展 -->
            <template #operationBeforeExtend="{ record }">
                <a-link @click="openPreviewModal(record)"><icon-eye /> 预览 </a-link>
            </template>

        </ma-crud>
        <Preview ref="previewRef" />
    </div>
</template>
<script setup lang="ts">
import { ref, reactive } from 'vue'
import manageKnowledge from '@/api/manage/knowledge'
import { Message } from '@arco-design/web-vue'
import tool from '@/utils/tool'
import * as common from '@/utils/common'
import Preview from '@/components/preview/index.vue'
const crudRef = ref()
const previewRef = ref<InstanceType<typeof Preview>>()

const numberOperation = (newValue, id, numberName) => {
    manageKnowledge.numberOperation({ id, numberName, numberValue: newValue }).then(res => {
        res.success && Message.success(res.message)
    }).catch(e => { console.log(e) })
}

const switchStatus = (status, record) => {
    manageKnowledge.changeStatus({ id: record.id, status }).then(res => {
        res.success && Message.success(res.message)
    }).catch(e => { console.log(e) })
}

const options = reactive({
    add: {
        api: manageKnowledge.save,
        auth: [
            ":manageKnowledge:save"
        ],
        show: true
    },
    api: manageKnowledge.getPageList,
    delete: {
        api: manageKnowledge.deletes,
        auth: [
            ":manageKnowledge:delete"
        ],
        realApi: manageKnowledge.realDeletes,
        realAuth: [
            ":manageKnowledge:realDelete"
        ],
        show: true
    },
    edit: {
        api: manageKnowledge.update,
        auth: [
            ":manageKnowledge:update"
        ],
        show: true
    },
    export: {
        auth: [
            ":manageKnowledge:export"
        ],
        show: true,
        url: 'manage/manageKnowledge/export'
    },
    formOption: {
        viewType: 'modal',
        width: 600
    },
    id: 'manage_knowledge',
    import: {
        auth: [
            ":manageKnowledge:import"
        ],
        show: true,
        templateUrl: 'manage/manageKnowledge/downloadTemplate',
        url: 'manage/manageKnowledge/import'
    },
    operationColumn: true,
    operationColumnWidth: 160,
    pk: 'id',
    recovery: {
        api: manageKnowledge.recoverys,
        auth: [
            ":manageKnowledge:recovery"
        ],
        show: true
    },
    recycleApi: manageKnowledge.getPageRecycleList,
    rowSelection: {
        showCheckedAll: true
    }
})
const columns = reactive([
    {
        addDisplay: false,
        commonRules: {
            message: "请输入主键",
            required: false
        },
        dataIndex: "id",
        editDisplay: false,
        formType: "input",
        hide: true,
        search: false,
        sortable: {},
        title: "主键"
    },
    {
        addDisplay: true,
        commonRules: {
            message: "请输入文件名称",
            required: true
        },
        dataIndex: "name",
        editDisplay: true,
        formType: "input",
        hide: false,
        search: true,
        sortable: {},
        title: "文件名称"
    },
    {
        addDisplay: true,
        commonRules: {
            message: "请输入文件编号",
            required: true
        },
        dataIndex: "code",
        editDisplay: true,
        formType: "input",
        hide: false,
        search: true,
        sortable: {},
        title: "文件编号"
    },
    {
        addDisplay: true,
        commonRules: {
            message: "请输入知识库类型",
            required: true
        },
        dataIndex: "knowledge_type",
        dict: {
            name: "knowledge",
            props: {
                label: "title",
                value: "key"
            },
            translation: true
        },
        editDisplay: true,
        formType: "select",
        hide: false,
        search: true,
        sortable: {},
        title: "知识库类型"
    },
    {
        addDisplay: false,
        commonRules: {
            message: "请输入创建者",
            required: false
        },
        dataIndex: "created_by",
        editDisplay: false,
        formType: "input",
        hide: true,
        search: false,
        sortable: {},
        title: "创建者"
    },
    {
        addDisplay: false,
        commonRules: {
            message: "请输入更新者",
            required: false
        },
        dataIndex: "updated_by",
        editDisplay: false,
        formType: "input",
        hide: true,
        search: false,
        sortable: {},
        title: "更新者"
    },
    {
        addDisplay: false,
        commonRules: {
            message: "请输入创建时间",
            required: false
        },
        dataIndex: "created_at",
        editDisplay: false,
        formType: "input",
        hide: true,
        search: false,
        sortable: {},
        title: "创建时间"
    },
    {
        addDisplay: false,
        commonRules: {
            message: "请输入更新时间",
            required: false
        },
        dataIndex: "updated_at",
        editDisplay: false,
        formType: "input",
        hide: true,
        search: false,
        sortable: {},
        title: "更新时间"
    },
    {
        addDisplay: false,
        commonRules: {
            message: "请输入删除时间",
            required: false
        },
        dataIndex: "deleted_at",
        editDisplay: false,
        formType: "input",
        hide: true,
        search: false,
        sortable: {},
        title: "删除时间"
    },
    {
        addDisplay: true,
        commonRules: {
            message: "请输入备注",
            required: false
        },
        dataIndex: "remark",
        editDisplay: true,
        formType: "input",
        hide: false,
        search: false,
        sortable: {},
        title: "备注"
    },
    {
		addDisplay: true,
		chunk: true,
		commonRules: {
			message: "请输入上传文件",
			required: true
		},
		dataIndex: "path",
		editDisplay: true,
		formType: "upload",
		hide: true,
		multiple: false,
		onlyData: true,
		returnType: "url",
		search: false,
		sortable: {},
		title: "上传文件",
        type: "file",
        width: 200
	}
])

function openPreviewModal(record) {
    previewRef.value?.open(record.path);
}
</script>
