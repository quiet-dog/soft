<template>
    <ATable :loading="pagination.loading" :data="data" :columns="columns" :pagination="pagination" page-position="top" @page-change="pageChange">
        <template #time="{ record }">
            <span>{{ dayjs(record.time).format("YYYY-MM-DD HH:mm:ss") }}</span>
        </template>
    </ATable>
</template>

<script lang="ts" setup>
import { InfluxdbRow, SensorRow } from '@/api/manage/sensor/types';
import { onActivated, onMounted, ref, watch } from 'vue';
import { TableColumnData } from '@arco-design/web-vue';
import sensor from '@/api/manage/sensor';
import dayjs from 'dayjs';

const { unit = "", name = "", sensorId = 0,activeTab = 0,currentTab=0 } = defineProps<{
    unit?: string,
    name?: string,
    sensorId?: number,
    activeTab?: number,
    currentTab?:number
}>()

watch(()=>activeTab,(newVal,oldVal)=>{
    
    if(newVal == 0){
        return
    }

    if(activeTab == currentTab){
        loadSensorData()
    }
})

const pagination = ref({
    total: 0,
    page: 1,
    pageSize: 10,
    loading: false
})

const data = ref<InfluxdbRow[]>([])
const columns = ref([])

function loadSensorData() {
    pagination.value.loading = true
    sensor.readHistoryData({
        ...pagination.value,
        sensorId: sensorId
    }).then(res => {
        data.value = res.data?.rows!
        pagination.value.total = res.data?.total!
        pagination.value.loading = false
    })
}

function pageChange(page: number) {
    pagination.value.page = page
    loadSensorData()
}

onMounted(() => {
    columns.value = [
        // @ts-expect-error 
        {
            title: name,
            dataIndex: 'c_' + sensorId,
            slotName: 'value',
        },
        // @ts-expect-error
        {
            title: '时间',
            dataIndex: 'time',
            slotName: 'time',
        }]
})

</script>

<style scoped></style>