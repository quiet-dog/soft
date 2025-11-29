<script lang='ts' setup>
import opc from '@/api/manage/opc';
import { Message } from '@arco-design/web-vue';

import { ref } from 'vue';

const { type, info, serverId } = defineProps<{
    type?: string;
    info?: any;
    serverId?: number
}>()

const extend = ref()
const opcConf = ref("")
const modbusConf = ref({
    slaveId: 0,
    startAddress: 0,
    quantity: 0,
    readType: 1
})
function handleConfirm(done) {
    if (type == "opc") {
        opc.nodeIdIsExit(serverId!, opcConf.value).then(res => {
            if (res.data! > 0) {
                extend.value.id = res.data
                extend.value.nodeId = opcConf.value
                done(true)
            } else {
                Message.warning("该节点无法使用")
            }
        })
    }

    // 判断type是否包含modbus
    if (type!.includes("modbus")) {
        extend.value.slaveId = modbusConf.value.slaveId
        extend.value.startAddress = modbusConf.value.startAddress
        extend.value.quantity = modbusConf.value.quantity
        extend.value.readType = modbusConf.value.readType
        done(true)
    }
}

function ok() {
    emit("success", extend.value)
}

function popupVisibleChange(val) {
    if (val) {
        extend.value = info
        if (type === 'opc') {
            opcConf.value = extend.value.nodeId
        }
        if (type!.includes("modbus")) {
            if (extend.value == undefined) {
                extend.value = {
                    slaveId: 0,
                    startAddress: 0,
                    quantity: 0,
                    readType: 0
                }
            }
            modbusConf.value.slaveId = extend.value?.slaveId != undefined ? extend.value.slaveId : 0
            modbusConf.value.startAddress = extend.value?.startAddress != undefined ? extend.value.startAddress : 0
            modbusConf.value.quantity = extend.value?.quantity != undefined ? extend.value.quantity : 0
            modbusConf.value.readType = extend.value?.readType != undefined ? extend.value.readType : 1
        }
    }
}

const emit = defineEmits(["success"])

</script>

<template>
    <APopconfirm @ok="ok" @popup-visible-change="popupVisibleChange" :onBeforeOk="handleConfirm">
        <template #content>
            <AInput v-model="opcConf" placeholder="填写点位信息" v-if="type === 'opc'" />

            <AForm v-else>
                <AFormItem label="功能类型">
                    <ASelect v-model="modbusConf.readType" placeholder="功能类型">
                        <AOption :value="1">读寄存器</AOption>
                        <AOption :value="2">读写寄存器</AOption>
                    </ASelect>
                </AFormItem>
                <AFormItem label="从站">
                    <AInputNumber v-model="modbusConf.slaveId" placeholder="从站" />
                </AFormItem>
                <AFormItem label="开始">
                    <AInputNumber v-model="modbusConf.startAddress" placeholder="开始地址" />
                </AFormItem>
                <AFormItem label="数量">
                    <AInputNumber v-model="modbusConf.quantity" placeholder="读取数量" />
                </AFormItem>
            </AForm>

        </template>
        <AButton type="text">已绑定</AButton>
    </APopconfirm>
</template>

<style scoped></style>
