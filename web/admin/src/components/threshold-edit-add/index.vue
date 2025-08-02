<script lang='ts' setup>
import alarmLabel from '@/api/manage/alarmLabel';
import threshold from '@/api/manage/threshold';
import { ThresholdRow } from '@/api/manage/threshold/types';
import { reactive, ref } from 'vue';
import { Message } from '@arco-design/web-vue';


const visible = ref(false)

const sensorId = ref(0)

const thresholdsV = ref<ThresholdRow[]>([])

const num = ref(1)

const selectOptions = ref([])
const selectTotal = ref(0)
const selectParams = ref({
    page: 1,
    pageSize: 10
})

function beforeOpen() {
    threshold.info(sensorId.value).then(res => {
        if (Array.isArray(res.data) && res.data.length > 0) {
            thresholdsV.value = res.data!
            num.value = thresholdsV.value.length
        } else {
            num.value = 1
            thresholdsV.value.push({
                sensorId: sensorId.value,
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
        selectTotal.value = res.data?.pageInfo.total!
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

// 懒加载标签
function scrollEvent() {
    if (selectTotal.value == 0) {
        return
    }

    if ((selectParams.value.pageSize * selectParams.value.page) >= selectTotal.value) {
        return
    }

    selectParams.value.page++
    // @ts-ignore
    alarmLabel.list({
        page: selectParams.value.page,
        pageSize: selectParams.value.pageSize
    }).then(res => {
        // @ts-ignore
        selectOptions.value.push(...res.data?.items)
        selectTotal.value = res.data?.pageInfo.total!
    })
}

function changeNum(val) {
    if (thresholdsV.value.length < val) {
        let diff = val - thresholdsV.value.length
        for (let i = 0; i < diff; i++) {
            thresholdsV.value.push({
                alarmLabelId: 0,
                template: "",
                sort: 0,
                sensorId: sensorId.value
            })
        }
    }

    if (thresholdsV.value.length > val) {
        thresholdsV.value.splice(val)
    }
}

function handleTable(val) {
    thresholdsV.value = val
}

function handleOk() {
    threshold.save(sensorId.value, thresholdsV.value).then(res => {
        if (res.code == 0) {
            Message.success(res.message)
            visible.value = false
        } else {
            Message.error(res.message)
        }
    })
}

defineExpose({
    handleOpen, handleClose
})


</script>

<template>
    <AModal @ok="handleOk" @before-open="beforeOpen" v-model:visible="visible">
        <AInputNumber @change="changeNum" mode="button" :default-value="num" />
        <ATable @change="handleTable" :draggable="{ type: 'handle', width: 40 }" :columns="columns" :data="thresholdsV">
            <template #alarmLabelId="{ rowIndex }">
                <ASelect v-model="thresholdsV[rowIndex].alarmLabelId" @dropdown-reach-bottom="scrollEvent">
                    <!-- @vue-skip -->
                    <AOption :value="item.id" v-for="item in selectOptions">{{ item.name }}</AOption>
                </ASelect>
            </template>
            <template #template="{ rowIndex }">
                <AInput placeholder="请输入expr表达式" v-model="thresholdsV[rowIndex].template" />
            </template>
        </ATable>
    </AModal>
</template>

<style scoped></style>
