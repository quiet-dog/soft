<script lang='ts' setup>
import { ref } from 'vue';



const visible = defineModel({
    type: Boolean,
    default: true,
})


const emit = defineEmits<{
    (e: 'changeExtend', value: { start: string;quantity:number }): void;
}>();

const extend = ref({
    start: "",
    quantity:0
})


const handleOk = () => {
    emit("changeExtend", extend.value);
    visible.value = false;
}
const handleCancel = () => {
    visible.value = false;
}

function onClose(){
    visible.value = false;
}


function changeSlave(val) {
    extend.value.start = "0x" + val
}

</script>

<template>
   <AModal @close="onClose" v-model:visible="visible" title="modbus节点配置" @ok="handleOk" @cancel="handleCancel">
        <AForm :model="extend">
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

<style scoped>

</style>
