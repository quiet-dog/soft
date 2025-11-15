<script lang='ts' setup>
import device from '@/api/manage/device';
import { DeviceInfo } from '@/api/manage/device/types';
import { SensorRow } from '@/api/manage/sensor/types';
import { reactive, ref } from 'vue';
import config from './config.vue';
import { Message } from '@arco-design/web-vue';
const visible = ref(false)
const deviceInfo = ref<DeviceInfo>()
const columns = reactive([{
    title: "名称",
    dataIndex: 'name',
    slotName: "name"
}, {
    title: "单位",
    dataIndex: "unit",
    slotName: "unit"
}, {
    title: "转换格式",
    dataIndex: "template",
    slotName: "template"
}, {
    title: "",
    dataIndex: "",
    slotName: 'extend'
}])
const data = ref<SensorRow[]>([])

function open(deviceId: number) {
    device.readSensorInfo(deviceId).then(res => {
        console.log("deviceInfo", deviceInfo.value)
        deviceInfo.value = res.data
        data.value = res.data?.sensors!
        visible.value = true
    }).catch(err => {

    })
}

function changeExtend(value, index: number) {
    data.value[index].extend = value
}

// 提交信息
function handleConfirm(done) {
    device.saveSensorInfo(data.value).then(res => {
        Message.success("配置成功")
        done(true)
    }).catch(err => {
        Message.error("配置失败")
    })
}


defineExpose({
    open
})
</script>

<template>
    <AModal :on-before-ok="handleConfirm" v-model:visible="visible" title="节点配置">
        <ATable :columns="columns" :data="data">
            <template #name="{ rowIndex }">
                <AInput v-model="data[rowIndex].name" />
            </template>
            <template #unit="{ rowIndex }">
                <AInput v-model="data[rowIndex].unit" />
            </template>
            <template #template="{ rowIndex }">
                <AInput v-model="data[rowIndex].template" />
            </template>
            <template #extend="{ record, rowIndex }">
                <config @success="changeExtend($event, rowIndex)" :type="deviceInfo?.server?.type"
                    :server-id="deviceInfo?.serverId" :info="record.extend" />
            </template>
        </ATable>
    </AModal>
</template>

<style scoped></style>
