<script lang='ts' setup>
import event from '@/api/manage/event';
import { EventRow } from '@/api/manage/event/types';
import { reactive, ref } from 'vue';
import { PaginationProps } from "@arco-design/web-vue"

const visible = ref(false)
const columns = reactive([
    { title: "编号", dataIndex: "id" },
    { title: "等级", dataIndex: "level", slotName: "level" },
    { title: "值", dataIndex: "value", slotName: "value" },
    { title: "创建时间", dataIndex: "createdAt", },
])
const data = ref<EventRow[]>([])
const searchParams = ref({
    page: 1,
    pageSize: 10,
    eventIds: [],
    sensorIds: [],
    alarmIds: []
})
const pagination = ref({
    // current: 1,
    // pageSize: 10,
    total: 0
})

function getEvent() {
    event.list({
        // page:pagination.value.current,
        // pageSize:pagination.value.pageSize,
        ...searchParams.value
    }).then(res => {
        data.value = res.data?.items!
        pagination.value.total = res.data?.pageInfo.total!
    }).finally(() => {
        visible.value = true
    })
}

function handleOpen(alarmIds: number[]) {
    // pagination.value.current = 1
    // pagination.value.pageSize = 10
    searchParams.value.page = 1
    searchParams.value.pageSize = 10
    // @ts-expect-error
    searchParams.value.alarmIds = alarmIds
    getEvent()

}

function pageChange(value) {
    searchParams.value.page = value
    getEvent()
}

defineExpose({
    handleOpen
})

</script>

<template>
    <AModal v-model:visible="visible">
        <ATabs>
            <ATabPane key="1" title="报警历史数据">
                <ATable page-position="top" :columns="columns" @page-change="pageChange" :pagination="pagination"
                    :data="data">
                    <template #level="{ record }">
                        <ATag :color="record.color">{{ record.level }}</ATag>
                    </template>
                    <template #value="{ record }">
                        <span :style="{ color: record.color }">{{ record.value }}</span>
                    </template>
                </ATable>
            </ATabPane>
        </ATabs>
    </AModal>
</template>

<style scoped></style>
