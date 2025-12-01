<template>
    <ALayout class="h-full">
        <ALayoutSider>
            <a-tree :data="areaList" :load-more="loadAreaTree" @select="selectArea" />
        </ALayoutSider>
        <ALayoutContent>
            <ASpin :loading="loading" class="w-full h-full">
                <ASpace>
                    <ACard v-for="value in deviceList" :key="value.id">
                        <template #title>
                            <span>{{ value.name }}</span>
                        </template>
                        <template #extra>
                            <span>{{ value.sensors.length }}</span>
                        </template>

                        <ADescriptions>
                            <ADescriptionsItem v-for="sensor in value.sensors" :key="sensor.id">
                                <span>{{ sensor.name }}</span>:
                                <span :style="{ color: isAlarm(value.isAlarm) ? 'red' : '' }">{{ sensor.value }}</span>
                                <span :style="{ color: isAlarm(value.isAlarm) ? 'red' : '' }">{{ sensor.unit }}</span>
                            </ADescriptionsItem>
                        </ADescriptions>
                    </ACard>
                </ASpace>
            </ASpin>
        </ALayoutContent>
    </ALayout>
</template>

<script lang="ts" setup>
import { useAllHook } from '.';

const { areaList, loadAreaTree, searchParams, getDeviceList, deviceList, loading, selectArea } = useAllHook()

function isAlarm(value) {
    if (value == undefined || value == null) {
        return false
    }
    return value
}

</script>

<style scoped></style>