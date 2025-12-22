<template>
    <div class="grid-stack" style="height: 100%;">
        <div class="grid-stack-item" v-for="(item, index) in data">
            <div class="grid-stack-item-content">
                <a-card :bordered="true">
                    <template #title>
                        <span>{{ item?.name }}</span>
                    </template>
                    <template #extra>
                        <span>{{ item?.sensors?.length || 0 }}</span>
                    </template>

                    <a-descriptions :column="1" size="small">
                        <a-descriptions-item v-for="sensor in item?.sensors" :key="sensor.id" label="">
                            <span>{{ sensor.name }}:</span>
                            <span :style="{ color: sensor.isAlarm ? 'red' : '' }">
                                {{ sensor.value }} {{ sensor.unit }}
                            </span>
                        </a-descriptions-item>
                    </a-descriptions>
                </a-card>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { GridStack } from 'gridstack';
import 'gridstack/dist/gridstack.min.css';
import { onMounted, nextTick, watch } from 'vue';

import { DragData } from './types';

let grid: GridStack | null = null; // 保持这样，不要用 ref

const { data = [] } = defineProps<{
    data: DragData[]
}>()

function isAlarm(value) {
    if (value == undefined || value == null) {
        return false
    }
    return value
}

onMounted(() => {
    // 只在这里初始化一次 grid！
    grid = GridStack.init({
        column: 12,           // 建议显式设置列数，默认为 12
        cellHeight: 'auto',       // 推荐固定高度，避免自动计算导致混乱
        margin: 6,
        float: false,          // 允许自动向下排列
        resizable: { handles: 'all' },
        acceptWidgets: true,
        auto: true,       // 自动布局容
    });

    // 监听布局变化（拖拽、缩放后触发）
    grid.on('change', (event, items) => {
        console.log('布局变化：', items.map(item => ({
            id: item.id,
            x: item.x,
            y: item.y,
            w: 0,
            h: 0
        })));
    });

    // 页面首次加载时，初始化已有的格子
    nextTick(() => {
        initAllWidgets();
    });
});

function computedWidth(item: DragData) {
    const sensorCount = item?.sensors?.length || 0;
    if (sensorCount <= 3) return 3;      // 少量传感器，窄卡片
    if (sensorCount <= 8) return 6;      // 中等，半宽
    if (sensorCount <= 15) return 9;     // 多，3/4宽
    return 12;                           // 超多，全宽（或你的 column 数）
}

// 关键：监听 data 长度变化，只注册新格子
watch(
    () => data.length,
    async () => {
        await nextTick(); // 等待新 DOM 渲染完成

        // 只把还没注册过的 .grid-stack-item 变成 widget
        const items = document.querySelectorAll('.grid-stack-item');
        items.forEach((el: Element) => {
            const htmlEl = el as HTMLElement;
            // gridstack 初始化后会给元素加 data-gs-instance 属性
            if (!htmlEl.dataset.gsInstance) {
                try {
                    // 设置高度
                    // 获取属性length
                    grid!.makeWidget(htmlEl, {
                        h: Number(htmlEl.getAttribute('data-length')),
                    });

                    // 设置h

                    console.log('新格子已注册：', grid?.getGridItems());
                } catch (e) {
                    // 重复调用会报错，catch 忽略即可
                }
            }
        });
    }
);

// 可选：手动初始化所有当前格子（首次加载用）
function initAllWidgets() {
    const items = document.querySelectorAll('.grid-stack-item');
    items.forEach((el: Element) => {
        try {
            grid!.makeWidget(el as HTMLElement);
        } catch (e) {
            // 已注册的忽略
        }
    });
}
</script>
<style scoped>
.drag-handle {
    cursor: move;
}

/* 内容容器必须有明确的高度传递机制 */
.grid-stack-item-content {
    padding: 0;
    /* Arco Card 自带 padding，无需额外 */
    overflow: visible;
    /* 允许内容决定高度 */
    height: auto;
    /* 关键 */
    min-height: 100px;
    /* 可选：最小高度，避免太矮 */
}

/* 确保 Arco Card 填满容器 */
:deep(.arco-card) {
    height: 100%;
    display: flex;
    flex-direction: column;
}

:deep(.arco-card-body) {
    flex: 1;
    overflow: auto;
    /* 内容太多时只在 body 内滚动，不影响外层高度 */
}
</style>
