<template>
    <ALayout class="h-full">
        <ALayoutSider>
            <a-tree :data="areaList" :load-more="loadAreaTree" @select="selectArea" />
        </ALayoutSider>
        <ALayoutContent>
            <ASpin :loading="loading" class="w-full h-full">
                <ASpace class="card-row">
                    <ACard v-for="value in deviceList" :key="value.id">
                        <template #title>
                            <span>{{ value.name }}</span>
                        </template>
                        <template #extra>
                            <span>{{ value.sensors.length }}</span>
                        </template>

                        <ADescriptions :column="1" size="small">
                            <ADescriptionsItem v-for="sensor in value.sensors" :key="sensor.id">
                                <span>{{ sensor.name }}</span>:
                                <span :style="{ color: isAlarm(value.isAlarm) ? 'red' : '' }">{{ sensor.value }}</span>
                                <span :style="{ color: isAlarm(value.isAlarm) ? 'red' : '' }">{{ sensor.unit }}</span>
                            </ADescriptionsItem>
                        </ADescriptions>
                    </ACard>
                </ASpace>
            </ASpin>

            <!-- <MyGridLayout :layout="gridLayout" cols="12" row-height="50" /> -->
            <!-- <VueDrag :data="deviceList" /> -->
        </ALayoutContent>
    </ALayout>
</template>

<script lang="ts" setup>
import { markRaw } from 'vue';
import { useAllHook } from '.';
import VueDrag from "@/package/vue-drag/index.vue";
// import MyGridLayout from '@/package/react/drag/drag';

const { areaList, loadAreaTree, searchParams, getDeviceList, deviceList, loading, selectArea } = useAllHook()

function isAlarm(value) {
    if (value == undefined || value == null) {
        return false
    }
    return value
}



</script>

<style scoped>
.card-row {
    display: flex;
    /* 强制 flex 布局 */
    flex-wrap: wrap;
    /* 允许换行 */
    align-items: stretch;
    /* 关键：同一行卡片拉伸到最高高度 */
    gap: 16px;
    /* 可选：卡片间距，更美观（替代 ASpace 的 split） */
}

.card-row:deep(.arco-card) {
    display: flex;
    /* ACard 内部也用 flex */
    flex-direction: column;
    /* 垂直排列 */
    height: 100%;
    /* 填满父容器高度 */
    /* width: calc(33.333% - 16px); */
    /* 可选：如果你想固定每行 3 张，可以这样控制宽度 */
}

/* ACard body 部分伸展填充剩余空间（可选，让内容区更灵活） */
:deep(.arco-card-body) {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    /* 可选：让底部内容对齐 */
}
</style>