<template>
    <div>
        <AInputNumber @change="changeNum" mode="button" v-model="num" />
        <ATable :columns="columns" :data="thresholds">
            <template #alarmLabelId="{ rowIndex }">
                <ASelect @change="changeAlarmLabelId($event, rowIndex)" v-model="thresholds![rowIndex].alarmLabelId">
                    <AOption :value="item.id" v-for="item in alarmLabels" :key="item.id">
                        <!-- @vue-expect-error -->
                        {{ item.level }}</AOption>
                </ASelect>
            </template>
            <template #template="{ rowIndex }">
                <AInput v-model="thresholds![rowIndex].template" />
            </template>
            <template #color="{ rowIndex }">
                <AColorPicker v-model="thresholds![rowIndex].color" disabled />
            </template>
        </ATable>
    </div>
</template>

<script lang="ts" setup>
import { AlarmLabelRow } from '@/api/manage/alarmLabel/types';
import { ThresholdRow } from '@/api/manage/threshold/types';
import { computed, onMounted, reactive, ref, watch } from 'vue';

const columns = reactive([
    { title: "报警标签", dataIndex: "alarmLabelId", slotName: "alarmLabelId", width: 150 },
    { title: "颜色", dataIndex: "color", slotName: "color", width: 100 },
    { title: "报警表达式", dataIndex: "template", slotName: "template", width: 100 },
])
const num = ref(0)

const {  sensorId = 0, alarmLabels = [] } = defineProps<{
    sensorId?: number
    alarmLabels?: AlarmLabelRow[]
}>()

const thresholds = defineModel<ThresholdRow[]>("thresholds")

function changeAlarmLabelId(val: number, rowIndex: number) {
    thresholds!.value![rowIndex].color = alarmLabels?.find(item => item.id == val)?.color || ""
}

function changeNum(val: number) {
    if(!thresholds.value){
        thresholds.value = []
    }

    if(val > thresholds.value?.length){
        let diff = val - thresholds.value?.length
        for(let i = 0; i < diff; i++){
            thresholds.value?.push({
                sensorId: sensorId,
                alarmLabelId: 0,
                template: "",
                sort: 0,
            })
        }
    }
    if(val < thresholds.value?.length){
        thresholds.value?.splice(val)
    }
}


// 如何知道组建打开了
onMounted(() => {
    num.value = thresholds.value?.length || 0
})


</script>
<style scoped></style>