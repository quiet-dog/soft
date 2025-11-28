<script lang='ts' setup>
import event from '@/api/manage/event';
import { EventRow } from '@/api/manage/event/types';
import { reactive, ref } from 'vue';
import { PaginationProps } from "@arco-design/web-vue"
import sensor from '@/api/manage/sensor';
import alarm from '@/api/manage/alarm';
import threshold from '@/api/manage/threshold';
import { InfluxdbRow } from '@/api/manage/sensor/types';
import { AlarmRow } from '@/api/manage/alarm/types';
import dayjs from 'dayjs';

const visible = ref(false)
const columns = ref([

])
const alarmInfoDetail = ref<AlarmRow>()
const data = ref<InfluxdbRow[]>([])
const searchParams = ref({
    page: 1,
    pageSize: 10,
    eventIds: [],
    sensorIds: [],
    alarmIds: [],
    sensorId: 0,
    beginTime: 0,
    endTime: 0,
})


const pagination = ref({
    total: 0,
    page: 1,
    pageSize: 10,
    loading: false
})

function getSenorData() {
    pagination.value.loading = true
    sensor.readHistoryData({ ...searchParams.value, page: pagination.value.page, pageSize: pagination.value.pageSize }).then(res => {
        data.value = res.data?.rows!
        pagination.value.total = res.data?.total!
    }).finally(() => {
        pagination.value.loading = false
    })
}

function handleOpen(alarmId: number) {
    visible.value = true
    pagination.value.loading = true
    alarm.read(alarmId).then(alarmInfo => {
        alarmInfoDetail.value = alarmInfo.data
        searchParams.value.sensorId = alarmInfo.data?.sensorId!
        searchParams.value.beginTime = alarmInfo.data?.sendTime!
        if (alarmInfo.data?.isLift) {
            searchParams.value.endTime = alarmInfo.data?.endTime!
        }
        sensor.readHistoryData({ ...searchParams.value, page: pagination.value.page, pageSize: pagination.value.pageSize }).then(res => {
            // @ts-ignore
            columns.value = [{
                title: res.data?.sensorName + "/" + res.data?.sensorUnit,
                dataIndex: "c_" + alarmInfo.data?.sensorId,
                slotName: "value",
            },
            // @ts-ignore
            {
                title: "报警等级",
                dataIndex: "level",
                slotName: "level",
            },
            // @ts-ignore
            {
                title: "时间",
                dataIndex: "time",
                slotName: "time",
            }]
            data.value = res.data?.rows!
            pagination.value.total = res.data?.total!
        }).finally(() => {
            pagination.value.loading = false
        })
    })
}



function pageChange(value) {
    pagination.value.page = value
    getSenorData()
}

defineExpose({
    handleOpen
})

</script>

<template>
    <AModal v-model:visible="visible">
        <ATable :loading="pagination.loading" page-position="top" :columns="columns" @page-change="pageChange"
            :pagination="pagination" :data="data">
            <template #level="{ record }">
                <ATag :color="record.color">{{ record.level }}</ATag>
            </template>
            <template #value="{ rowIndex, record }">
                <span :style="{ color: record.color }">{{ data[rowIndex]["c_" + alarmInfoDetail?.sensorId] }}</span>
            </template>
            <template #time="{ record }">
                <span>{{ dayjs(record.time).format("YYYY-MM-DD HH:mm:ss") }}</span>
            </template>
        </ATable>
    </AModal>
</template>

<style scoped></style>
