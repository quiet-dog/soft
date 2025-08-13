<script lang='ts' setup>
import { ref } from 'vue';

const visible = ref(false)
const emit = defineEmits(['changeExtend'])

const extend = ref({
    type: "value",
    value: "",
    objectId: "",
    nodeId: "",
    methondId: ""
})


function changeTextArea(val) {
    try {
        const parsed = JSON.parse(val);
        if (Array.isArray(parsed)) {
            // @ts-expect-error
            extend.value.value = parsed
            console.log('这是数组');
        } else if (parsed === null) {
            console.log('这是 null');
        } else {
            // const type = typeof parsed;
            extend.value.value = parsed
        }
    } catch (e) {
        console.log('JSON 解析错误');
    }
}

function open(value) {
    visible.value = true
    extend.value = value
}

function handleOk() {
    emit("changeExtend", extend.value)
}

defineExpose({
    open
})

</script>

<template>
    <AModal @ok="handleOk" v-model:visible="visible">
        <AForm>
            <ARow>
                <ACol>
                    <AFormItem label="类型">
                        <ASelect v-model="extend.type">
                            <AOption value="value">
                                节点
                            </AOption>
                            <AOption value="method">
                                方法
                            </AOption>
                        </ASelect>
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow v-if="extend.type == 'value'">
                <ACol>
                    <AFormItem label="节点ID">
                        <AInput v-model="extend.nodeId" />
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow v-if="extend.type == 'method'">
                <ACol>
                    <AFormItem label="方法ID">
                        <AInput v-model="extend.methondId" />
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow v-if="extend.type == 'method'">
                <ACol>
                    <AFormItem label="方法对象ID">
                        <AInput v-model="extend.objectId" />
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow>
                <ACol>
                    <AFormItem label="值">
                        <ATextarea @change="changeTextArea" />
                    </AFormItem>
                </ACol>
            </ARow>
        </AForm>
    </AModal>
</template>

<style scoped></style>
