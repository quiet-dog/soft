<script lang='ts' setup>
import device from '@/api/manage/device';
import { ref } from 'vue';
import { Message } from '@arco-design/web-vue';
import { ExtendType } from '.';


const visible = defineModel({
    type: Boolean,
    default: true,
})

const { serverId = 0,sExtend } = defineProps<{
    serverId: number,
    sExtend:ExtendType
}>()

const extend = ref({
    slave: "",
    duration: 5
})

const testBtnLoading = ref(false)

function changeSlave(val) {
    extend.value.slave = "0x" + val
}

const emit = defineEmits<{
    (e: 'changeExtend', value: ExtendType): void;
}>();


function onClose() {
    visible.value = false;
    emit("changeExtend", extend.value);
}

function testConnect() {
    testBtnLoading.value = true
    device.test({
        serverId: serverId,
        extend: JSON.stringify(extend.value)
    }).then(res=>{
        if(res.code != 0){
            Message.error(res.message)
        }else{
            Message.success("测试通过")
        }
    }).finally(()=>{
        testBtnLoading.value = false
    })
}


function handleOpen(){
    if(sExtend != undefined && sExtend !=null){
        extend.value = JSON.parse(JSON.stringify(sExtend))
    } else {
        extend.value.duration = 5
        extend.value.slave = ""
    }
}
</script>

<template>
    <AModal @open="handleOpen" @close="onClose" v-model:visible="visible" title="modbus设备配置">
        <AForm :model="extend">
            <AFormItem label="从站地址">
                <AInput @change="changeSlave" placeholder="请输入十六进制地址">
                    <template #prepend>
                        0x
                    </template>
                </AInput>
            </AFormItem>
            <AFormItem label="采集间隔">
                <AInputNumber :default-value="5" v-model="extend.duration" placeholder="请输入采集间隔">
                    <template #suffix>
                        <span>秒</span>
                    </template>
                </AInputNumber>
            </AFormItem>
        </AForm>
        <div style="margin: 0 auto;display: flex;justify-content: center;">
            <AButton type="primary" :loading="testBtnLoading" @click="testConnect">
                测试
            </AButton>
        </div>
    </AModal>
</template>

<style scoped></style>
