<script lang="ts" setup>
import sensor from '@/api/manage/sensor';
import { SensorRow } from '@/api/manage/sensor/types';
import SensorTable from './sensor-table.vue';
import { ref } from 'vue';
import device from '@/api/manage/device';
import dayjs from 'dayjs';

const visible = ref(false)
const tabs = ref<{ label: string, key: number, unit?: string, name?: string, sensorId?: number }[]>([
    {
        label: '总览',
        key: 0,
    },
])
const activeTab = ref(0)
const data = ref<SensorRow[]>([])
const columns = ref([
    {
        title: '名称',
        dataIndex: 'name',
    },
    {
        title: "当前值",
        dataIndex: 'value',
        slotName: 'value',
    },
    {
        title: "单位",
        dataIndex: 'unit',
        slotName: 'unit',
    },
    {
        title: "时间",
        dataIndex: 'dataTime',
        slotName: 'dataTime',
    }
])

function changeTab(key: number) {
    sensor.readHistoryData({
        sensorId: key,
    })
}

function open(deviceId: number) {
    device.getSensorNow(deviceId).then(res=>{
        if (res.data?.length! > 0) {
            res.data?.forEach(item => {
                tabs.value.push({
                    label: item.name,
                    key: item.id,
                    unit: item.unit,
                    name: item.name,
                    sensorId: item.id,
                })
            })
            data.value = res.data!
        }
        visible.value = true
    })

    // sensor.list({
    //     deviceId: deviceId,
    //     page: 1,
    //     pageSize: 100,
    //     name: '',
    //     sensorTypeId: 0,
    // }).then(res => {
    //     console.log(res.data);
    //     if (res.data?.items?.length! > 0) {
    //         res.data?.items?.forEach(item => {
    //             tabs.value.push({
    //                 label: item.name,
    //                 key: item.id,
    //                 unit: item.unit,
    //                 name: item.name,
    //                 sensorId: item.id,
    //             })
    //         })
    //         data.value = res.data?.items!
    //     }
    //     visible.value = true
    // })
    // visible.value = true
}

function close() {
    // 除第一个之外的都删除
    tabs.value.splice(1)
}

defineExpose({
    open
})

</script>
<template>
    <AModal @close="close" v-model:visible="visible" title="当前数据">
        <ATabs v-model:active-key="activeTab">
            <ATabPane @change="changeTab" v-for="tab in tabs" :key="tab.key" :title="tab.label">
                <ATable v-if="tab.key == 0" :data="data" :columns="columns">
                    <template #dataTime="{ record }">
                        <span>{{ dayjs(record.dataTime).format("YYYY-MM-DD HH:mm:ss") }}</span>
                    </template>

                </ATable>
                <SensorTable :current-tab="tab.key" :active-tab="activeTab" :name="tab.name" :unit="tab.unit" :sensorId="tab.sensorId" v-else />
            </ATabPane>
        </ATabs>
    </AModal>
</template>
<style scoped></style>