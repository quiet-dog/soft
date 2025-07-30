<script lang='ts' setup>
import { onMounted, provide, ref, watch } from 'vue';
import { use } from "echarts/core";
import VChart, { THEME_KEY } from "vue-echarts";
import { CanvasRenderer } from "echarts/renderers";
import { LineChart } from "echarts/charts";
import dayjs from 'dayjs';
import sensor from '@/api/manage/sensor';
import * as echarts from "echarts";
import {
    TitleComponent,
    TooltipComponent,
    LegendComponent
} from "echarts/components";

use([
    CanvasRenderer,
    LineChart,
    TitleComponent,
    TooltipComponent,
    LegendComponent
])
provide(THEME_KEY, "dark");



const options = ref({
    color: ['#80FFA5', '#00DDFF', '#37A2FF', '#FF0087', '#FFBF00'],
    title: {
        text: ''
    },
    tooltip: {
        trigger: 'axis',
        axisPointer: {
            type: 'cross',
            label: {
                backgroundColor: '#6a7985'
            }
        }
    },
    legend: {
        data: []
    },
    toolbox: {
        feature: {
            saveAsImage: {}
        }
    },
    grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true
    },
    xAxis: [
        {
            type: 'category',
            boundaryGap: false,
            data: []
        }
    ],
    yAxis: [
        {
            type: 'value'
        }
    ],
    series: [
        {
            name: '',
            type: 'line',
            stack: 'Total',
            smooth: true,
            lineStyle: {
                width: 0
            },
            showSymbol: false,
            areaStyle: {
                opacity: 0.8,
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                    {
                        offset: 0,
                        color: 'rgb(128, 255, 165)'
                    },
                    {
                        offset: 1,
                        color: 'rgb(1, 191, 236)'
                    }
                ])
            },
            emphasis: {
                focus: 'series'
            },
            data: []
        },
        {
            name: '',
            type: 'line',
            stack: 'Total',
            smooth: true,
            lineStyle: {
                width: 0
            },
            showSymbol: false,
            areaStyle: {
                opacity: 0.8,
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                    {
                        offset: 0,
                        color: 'rgb(0, 221, 255)'
                    },
                    {
                        offset: 1,
                        color: 'rgb(77, 119, 255)'
                    }
                ])
            },
            emphasis: {
                focus: 'series'
            },
            data: []
        }
    ]
})

const form = ref({
    beginTime: 0,
    endTime: 0,
})

function changeDate(value, date, dateString) {
    if (Array.isArray(value)) {
        form.value.beginTime = Number(value[0])
        form.value.endTime = Number(value[1])
    } else {
        form.value.beginTime = dayjs().subtract(1, "hour").millisecond()
        form.value.endTime = dayjs().millisecond()
    }
    renderEchart()
}


const { sensorId = 0 } = defineProps<{
    sensorId: number
}>()

const vChartRef = ref()
function renderEchart() {
    sensor.readEchart({
        sensorId: sensorId == 0 ? 48 : sensorId,
        beginTime: Number(form.value.beginTime),
        endTime: Number(form.value.endTime),
    }).then(res => {

        // @ts-ignore
        options.value.series[0].name = "真实_" + res.data?.sensorTypeName
        // @ts-ignore
        options.value.series[1].name = "转换_" + res.data?.sensorTypeName
        // @ts-ignore
        options.value.legend.data = [options.value.series[0].name, options.value.series[1].name]
        // @ts-ignore
        options.value.series[0].data = res.data?.cSeiresData
        // @ts-ignore
        options.value.series[1].data = res.data?.eSeiresData
        // @ts-ignore
        options.value.xAxis[0].data = res.data?.xData
        vChartRef.value.resize()
    })
}


function render() {
    renderEchart()
}


defineExpose({
    render
})


onMounted(() => {
    renderEchart()
})

</script>

<template>
    <div>
        <a-form :model="form">
            <a-row>
                <a-col :span="12">
                    <a-form-item label="精确度">
                        <a-input-number :precision="1" :step="1" />
                    </a-form-item>
                </a-col>
                <a-col :span="12">
                    <a-form-item label="时间选择">
                        <a-range-picker value-format="x" show-time @change="changeDate"
                            :time-picker-props="{ defaultValue: ['00:00:00', '23:59:59'] }" />
                    </a-form-item>
                </a-col>
            </a-row>
        </a-form>
        <div style="display: flex; justify-content: center; align-items: center; margin-top: 10px;">
            <a-space :size="16" :align="'center'">
                <a-button type="primary">搜索</a-button>
                <a-button type="dashed">重置</a-button>
            </a-space>
        </div>
        <VChart ref="vChartRef" class="echart" :option="options" />
    </div>
</template>

<style scoped>
.echart {
    height: 400px;
    width: 100%;
    margin-top: 10px;
}
</style>
