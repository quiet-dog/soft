<script lang='ts' setup>
import sensor from '@/api/manage/sensor';
import { InfluxdbData } from '@/api/manage/sensor/types';
import { onMounted, ref } from 'vue';


const visible = defineModel({
    type: Boolean,
    default: false,
})

const { sensorId = 0 } = defineProps<{
    sensorId: number
}>()

const params = ref<InfluxdbData>(
    {
        sensorIds: [],
        beginTime: '',
        endTime: '',
        deviceId: 0,
        page: 1,
        pageSize: 10
    })

const columns = ref([])
const data = ref([])
const pageInfo = ref({
    total: 0,
    page: 1,
    pageSize: 10
})


function beforeOpen() {
    sensor.read(sensorId).then(res => {
        params.value.sensorIds = [res.data!.id];
        params.value.deviceId = res.data!.deviceId;
        sensor.data(params.value).then(ress => {
            console.log(ress.data);
            columns.value = [{
                title: "真实_" + ress.data!.sensorTypeName,
                dataIndex: "c_" + sensorId
            }, {
                title: "转换_" + ress.data!.sensorTypeName,
                dataIndex: "e_" + sensorId
            }, {
                title: "时间",
                dataIndex: "time"
            }]
            data.value = ress.data.rows
            pageInfo.value.total = ress.data.total
        }).catch(err => {
            console.error(err);
        })
    })
}

function changePage() {
    sensor.data({
        ...params.value,
        page: pageInfo.value.page,
        pageSize: pageInfo.value.pageSize
    }).then(res => {
        data.value = res.data.rows
        pageInfo.value.total = res.data.total
    })
}

onMounted(() => {
    sensor.data(params.value).then(res => {
        console.log(res.data);
    }).catch(err => {
        console.error(err);
    })
})

</script>

<template>
    <AModal v-model:visible="visible" @before-open="beforeOpen" title="查看数据" width="60%">
        <div>
            <ATabs>
                <ATabPane key="1" title="数据详情">
                    <ATable :pagination="false" :columns="columns" :data="data"></ATable>
                    <APagination @change="changePage" v-model:current="pageInfo.page"
                        v-model:page-size="pageInfo.pageSize" :total="pageInfo.total" />
                </ATabPane>
                <ATabPane key="2" title="数据图表">

                </ATabPane>
            </ATabs>
        </div>
    </AModal>
</template>

<style scoped></style>
