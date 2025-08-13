<script lang='ts' setup>
import device from '@/api/manage/device';
import deviceControl from '@/api/manage/device-control';
import { DeviceRead } from '@/api/manage/device/types';
import server from '@/api/manage/server';
import { ServerRow } from '@/api/manage/server/types';
import { ref } from 'vue';
import ModbusControl from "@/package/control/modbus/index.vue";
import OpcControl from "@/package/control/opc/index.vue";

const visible = ref(false)
const table = ref([])
const num = ref(0)
const columns = [
    {
        title: '名称',
        dataIndex: 'name',
        slotName: 'name',
    },
    {
        title: '',
        dataIndex: 'opeara',
        slotName: 'opeara',
    }
]
const deviceInfo = ref<DeviceRead>()
const serverInfo = ref<ServerRow>()
const cursorIndex = ref(0)

function open(deviceId: number) {
    visible.value = true
    device.read(deviceId).then(res => {
        deviceInfo.value = res.data
        server.read(deviceInfo.value?.serverId!).then(s => {
            serverInfo.value = s.data
        })
    })
    deviceControl.list({
        deviceIds: [deviceId],
        page: 1,
        pageSize: 100,
        name: "",
        extend: "",
        deviceId: 0
    }).then(res => {
        num.value = res.data?.pageInfo.total!
        // @ts-expect-error
        table.value = res.data?.items
    })
}

function changeNum(val: number) {
    // 假设 writeValues.value 是 number[] 数组
    if (table.value.length < val) {
        // 长度小于 val，往数组尾部追加 0x01
        while (table.value.length < val) {
            // @ts-expect-error
            table.value.push({
                name: "",
                deviceId: deviceInfo.value?.id,
                extend: {}
            });
        }
    } else if (table.value.length > val) {
        // 长度大于 val，删除多余元素
        while (table.value.length > val) {
            table.value.pop();
        }
    }
}

const modbusControlRef = ref<InstanceType<typeof ModbusControl>>()
const opcControlRef = ref<InstanceType<typeof OpcControl>>()


async function configCommand(index) {
    cursorIndex.value = index
    if (serverInfo.value?.type == "opc") {
        // @ts-expect-error
        opcControlRef.value?.open(table.value[cursorIndex.value].extend)
    } else if (serverInfo.value?.type.includes('modbus')) {
        // @ts-expect-error
        modbusControlRef.value?.open(table.value[cursorIndex.value].extend)
    }

}

function changeExtend(extend) {
    // @ts-expect-error
    table.value[cursorIndex.value].extend = extend
}


function handleOk() {
    deviceControl.addControl(table.value).then(res => {

    }).catch(err => {

    })
}

defineExpose({
    open
})
</script>

<template>
    <AModal @ok="handleOk" v-model:visible="visible">
        <AInputNumber @change="changeNum" v-model="num" mode="button" />
        <ATable :columns="columns" :data="table">
            <template #name="{ rowIndex }">
                <!-- @vue-expect-error -->
                <AInput v-model="table[rowIndex].name" />
            </template>
            <template #opeara="{ rowIndex }">
                <AButton @click="configCommand(rowIndex)">配置</AButton>
            </template>
        </ATable>
        <!-- <component v-if="controlComponent" ref="controlComponentRef" :is="controlComponent"
            @changeExtend="changeExtend" /> -->
        <ModbusControl ref="modbusControlRef" />
        <OpcControl ref="opcControlRef" />
    </AModal>
</template>

<style scoped></style>
