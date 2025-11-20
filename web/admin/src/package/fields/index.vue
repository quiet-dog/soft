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
const dataLength = ref(0)

function open(deviceId: number) {
    device.readSensorInfo(deviceId).then(res => {
        console.log("deviceInfo", deviceInfo.value)
        deviceInfo.value = res.data
        data.value = res.data?.sensors!
        if (Array.isArray(res.data?.sensors)) {
            dataLength.value = res.data?.sensors.length
        }
        visible.value = true
    }).catch(err => {

    })
}

function changeExtend(value, index: number) {
    data.value[index].extend = value
}

// 提交信息
function handleConfirm(done) {
    device.saveSensorInfo({
        sensors: data.value,
        deviceId: deviceInfo.value?.id
    }).then(res => {
        Message.success("配置成功")
        done(true)
    }).catch(err => {
        Message.error("配置失败")
    })
}

function changeDataLength(val) {
    console.log("data.value", data.value)
    if (data.value == null || data.value == undefined) {
        data.value = []
    }
    if (val > data.value.length) {
        const need = val - data.value.length;
        for (let i = 0; i < need; i++) {
            let extend;
            if (deviceInfo.value?.server?.type == "opc") {
                extend = {
                    id: 0,
                    nodeId: ""
                }
            }
            data.value.push({
                id: 0,
                deviceId: deviceInfo.value?.id!,
                name: "",
                unit: "",
                template: "",
                remark: "",
                sensorTypeId: 0,
                extend: extend
            });
        }
    } else if (val < data.value.length) {
        // 删掉多余的
        data.value.splice(val);
    }
}


defineExpose({
    open
})
</script>

<template>
    <AModal :on-before-ok="handleConfirm" v-model:visible="visible" title="节点配置">
        <AInputNumber @change="changeDataLength" v-model="dataLength" :default-value="dataLength" mode="button" />
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
