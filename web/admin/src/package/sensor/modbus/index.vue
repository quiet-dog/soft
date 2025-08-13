<script lang='ts' setup>
import { ref } from 'vue';
import { ExtendType } from "."

const visible = defineModel({
    type: Boolean,
    default: true,
})

const { sExtend } = defineProps<{
    sExtend: ExtendType
}>()

const emit = defineEmits<{
    (e: 'changeExtend', value: ExtendType): void;
}>();

const extend = ref({
    start: "",
    quantity: 0,
    readType:0
})


const handleOk = () => {
    emit("changeExtend", extend.value);
    visible.value = false;
}

const handleCancel = () => {
    visible.value = false;
}

function onClose() {
    visible.value = false;
}

function changeSlave(val) {
    extend.value.start = "0x" + val
}

function handleOpen() {
    if (sExtend != undefined && sExtend != null) {
        extend.value = JSON.parse(JSON.stringify(sExtend))
    } else {
        extend.value.quantity = 0
        extend.value.start = ""
        extend.value.readType = 1
    }

}

function changeSelect(value){
    extend.value.readType = Number(value)
}

</script>

<template>
    <AModal @open="handleOpen" @close="onClose" v-model:visible="visible" title="modbus节点配置" @ok="handleOk"
        @cancel="handleCancel">
        <AForm :model="extend">
         <ARow>
                <ACol>
                    <AFormItem label="功能类型">
                        <ASelect default-value="1" @change="changeSelect">
                            <AOption value="1">读寄存器</AOption>
                            <AOption value="2">读写寄存器</AOption>
                        </ASelect>
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow>
                <ACol>
                    <AFormItem label="起始地址">
                        <AInput @change="changeSlave" placeholder="请输入十六进制地址">
                            <template #prepend>
                                0x
                            </template>
                        </AInput>
                    </AFormItem>
                </ACol>
            </ARow>

            <ARow>
                <ACol>
                    <AFormItem label="数量">
                        <AInputNumber v-model="extend.quantity" />
                    </AFormItem>
                </ACol>
            </ARow>
        </AForm>
    </AModal>
</template>

<style scoped></style>
