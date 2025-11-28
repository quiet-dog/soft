<template>
    <AModal :unmount-on-close="true"    @ok="handleOk" title="报警配置" v-model:visible="visible">
        <ATabs>
            <ATabPane v-for="item in sensorAlarmList" :key="item.id" :title="item.name">
                <ThresholdConfig v-model:thresholds="item.thresholds" :alarm-labels="alarmLabels" :sensor-id="item.id" />
            </ATabPane>
        </ATabs>

    </AModal>
</template>

<script lang="ts" setup>
import device from '@/api/manage/device';
import { SensorAlarmRow } from '@/api/manage/sensor/types';
import { ref } from 'vue';
import ThresholdConfig from './threshold-config.vue';
import { AlarmLabelRow } from '@/api/manage/alarmLabel/types';
import alarmLabel from '@/api/manage/alarmLabel';




const visible = ref(false)
const sensorAlarmList = ref<SensorAlarmRow[]>([])
const alarmLabels = ref<AlarmLabelRow[]>([])
const deviceId = ref(0)

async function open(id: number) {
    deviceId.value = id
    device.getSensorAlarmList(id).then(res => {
        sensorAlarmList.value = res.data!
    })

    alarmLabel.list({
        page: 1,
        pageSize: 1000,
        name: "",
        remark: "",
        color: "",
        label: "",
    }).then(res => {
        alarmLabels.value = res.data!.items
    })

    visible.value = true
}

function handleOk(){
    console.log(sensorAlarmList.value)
    device.saveSensorAlarmList({
        sensors: sensorAlarmList.value,
        deviceId: deviceId.value,
    }).then(res => {
        // console.log(res)
        visible.value = false
    })
}

defineExpose({
    open
})
</script>

<style scoped></style>