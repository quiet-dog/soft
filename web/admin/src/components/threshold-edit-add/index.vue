<script lang='ts' setup>
import alarmLabel from '@/api/manage/alarmLabel';
import threshold from '@/api/manage/threshold';
import { ThresholdRow } from '@/api/manage/threshold/types';
import { reactive, ref } from 'vue';

const visible = ref(false)

const sensorId = ref(0)

const thresholdsV = ref<ThresholdRow[]>([])

const num = ref(1)

const selectOptions = ref([])
const selectParams = ref({
    page: 1,
    pageSize: 10
})

function beforeOpen() {
    threshold.info(sensorId.value).then(res => {
        if (Array.isArray(res.data) && res.data.length > 0) {
            thresholdsV.value = res.data!
        } else {
            num.value = 1
            thresholdsV.value.push({
                sensorId: 0,
                sort: 0,
                alarmLabelId: 0,
                template: ""
            })
        }
    })

    // @ts-ignore
    alarmLabel.list({
        page: selectParams.value.page,
        pageSize: selectParams.value.pageSize
    }).then(res => {
        // @ts-ignore
        selectOptions.value = res.data?.items
    })
}

const columns = reactive([
    { title: "报警标签", dataIndex: "alarmLabelId", slotName: "alarmLabelId", width: 150 },
    { title: "报警表达式", dataIndex: "template", slotName: "template", width: 100 },
])

const handleOpen = (id: number) => {
    sensorId.value = id
    visible.value = true
}

const handleClose = () => {
    visible.value = false
}

defineExpose({
    handleOpen, handleClose
})


</script>

<template>
    <AModal @before-open="beforeOpen" v-model:visible="visible">
        <AInputNumber mode="button" v-model="num" />
        <ATable :columns="columns" :data="thresholdsV">
            <template #alarmLabelId="{ rowIndex }">
                <ASelect>
                    <!-- @vue-skip -->
                    <AOption v-for="item in selectOptions">{{ item.name }}</AOption>
                </ASelect>
            </template>
            <template #template="{ rowIndex }">
                <ATextarea placeholder="请输入expr表达式" allow-clear v-model="thresholdsV[rowIndex].template" />
            </template>
        </ATable>
    </AModal>
</template>

<style scoped></style>
