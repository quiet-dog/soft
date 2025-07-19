<script lang='ts' setup>
import { computed, onMounted, ref, shallowRef, watch } from "vue"
import "./index"
import { VueMonacoEditor } from '@guolao/vue-monaco-editor'
import sensor from "@/api/manage/sensor";
import { TemplateEnv } from "@/api/manage/sensor/types";
import { Message } from '@arco-design/web-vue';


const { templateInfo = {
    type: "",
    extend: ""
} } = defineProps<{
    templateInfo: {
        type: string;
        extend: string
    }
}>()

const emit = defineEmits<{
    (e: 'changeTemplate', value: string): void;
}>();


const visible = defineModel({
    type: Boolean,
    default: false,
})


const MONACO_EDITOR_OPTIONS = {
    automaticLayout: true,
    formatOnType: true,
    formatOnPaste: true,
}

const code = ref('// some code...')
const editor = shallowRef()
const handleMount = editorInstance => (editor.value = editorInstance)
const sensorData = ref<TemplateEnv>()


// your action
function formatCode() {
    editor.value?.getAction('editor.action.formatDocument').run()
}

watch(code, (value) => {
    console.log(value)
})

function beforeOpen() {
    sensor.readData({
        type: templateInfo.type,
        extend: templateInfo.extend
    }).then(res => {
        sensorData.value = res.data
    })
}

const data = computed(() => {
    return [
        {
            label: "值",
            value: sensorData.value?.value
        }, {
            label: "值类型",
            value: sensorData.value?.type
        }, {
            label: "产生时间",
            value: sensorData.value?.createTime
        }
    ]
})

function handleOk() {
    emit("changeTemplate", code.value)
    visible.value = false
}

function handleClose() {
    templateInfo.type = ""
    templateInfo.extend = ""
}

function translateData(){
    sensor.translate({
        env:sensorData.value!,
        template:code.value
    }).then(res=>{
        Message.success(String(res.data))
    })
}

// onMounted(() => {

// })
</script>

<template>
    <AModal @close="handleClose" @ok="handleOk" v-model:visible="visible" @before-open="beforeOpen" width="1200px">
        <ADescriptions title="传感器数据" bordered :label-style="{
            color: 'black'
        }" align="center" :data="data" />
        <AButton type="primary" @click="translateData">转换</AButton>
        <div style="height: 500px;width:1000px;margin: auto;">
            <VueMonacoEditor v-model:value="code" language="javascript" theme="vs-dark" :options="MONACO_EDITOR_OPTIONS"
                @mount="handleMount" />
        </div>
    </AModal>
</template>

<style scoped></style>
