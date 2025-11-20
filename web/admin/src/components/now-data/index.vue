<script lang="ts" setup>
import sensor from '@/api/manage/sensor';
import { SensorRow } from '@/api/manage/sensor/types';
import { ref } from 'vue';

const visible = ref(false)
const tabs = ref([
    {
        label: '总览',
        key: 0,
    },
])
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
    }
])

function open(deviceId: number) {
    sensor.list({
        deviceId: deviceId,
        page: 1,
        pageSize: 100,
        name: '',
        sensorTypeId: 0,
    }).then(res => {
        console.log(res.data);
        if (res.data?.items?.length! > 0) {
            res.data?.items?.forEach(item => {
                tabs.value.push({
                    label: item.name,
                    key: item.id,
                })
            })
            data.value = res.data?.items!
        }
        visible.value = true
    })
    visible.value = true
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
        <ATabs>
            <ATabPane v-for="tab in tabs" :key="tab.key" :title="tab.label">
                <ATable v-if="tab.key == 0" :data="data" :columns="columns"></ATable>
            </ATabPane>
        </ATabs>
    </AModal>
</template>
<style scoped></style>