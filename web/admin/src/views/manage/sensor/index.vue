<script lang='ts' setup>
import { useSensorHook } from '.';
import Template from "@/package/template/index.vue"
import InSensor from "@/components/in-sensor/index.vue"
import ThreshodEditAdd from "@/components/threshold-edit-add/index.vue"
import { ref } from 'vue';

const { crud, columns,
    crudRef, changeSensorType, loadSensorOptions,
    sensorTypeOptions, deviceOptions,
    loadDeviceOptions, changeDevice, asyncComponent,
    visible, asyncDeviceId, changeExtend, templateVisable,
    templateInfo, changeTemplate, sExtend
} = useSensorHook()

const sensorDataVisible = ref(false)
const sensorDataId = ref(0)
function changeSensorData(id: number) {
    sensorDataId.value = id
    sensorDataVisible.value = true
}
const thresholdEddAddRef = ref<InstanceType<typeof ThreshodEditAdd>>()

function selectOperation(value: string, id: number) {
    if (value == "thresholdSet") {
        thresholdEddAddRef.value?.handleOpen(id);
    }
}

</script>

<template>
    <div class="ma-content-block lg:flex justify-between p-4">
        <ma-crud :options="crud" :columns="columns" ref="crudRef">
            <template #form-sensorTypeId>
                <a-select @change="changeSensorType" @dropdown-reach-bottom="loadSensorOptions" placeholder="请选择服务器">
                    <a-option v-for="item in sensorTypeOptions.items" :key="item.value" :value="item.value">
                        {{ item.label }}
                    </a-option>
                </a-select>
            </template>

            <template #form-deviceId>
                <a-select @dropdown-reach-bottom="loadDeviceOptions" placeholder="请选择设备">
                    <a-option @click="changeDevice(item.value)" v-for="item in deviceOptions.items" :key="item.value"
                        :value="item.value">
                        {{ item.label }}
                    </a-option>
                </a-select>
            </template>

            <!-- 设备在线状态  -->
            <template #isOnline="{ record }">
                <a-tag v-if="record.isOnline" :color="'#00b42a'">
                    在线
                </a-tag>
                <a-tag :color="'#f53f3f'" v-else>
                    离线
                </a-tag>
            </template>

            <template #view="{ record }">
                <AButton type="primary" @click="changeSensorData(record.id)">查看</AButton>
            </template>

            <!-- 操作之后的 -->
            <template #operationAfterExtend="{ record }">
                <a-dropdown @select="selectOperation($event, record.id)" trigger="hover">
                    <a-link><icon-double-right /> 更多</a-link>
                    <template #content>
                        <a-doption value="thresholdSet" v-auth="['manage:threshold:save']">阈值设置</a-doption>
                    </template>
                </a-dropdown>
            </template>
        </ma-crud>
        <component :is="asyncComponent" :deviceId="asyncDeviceId" @changeExtend="changeExtend" :sExtend="sExtend" />
        <Template v-model:model-value="templateVisable" :templateInfo="templateInfo"
            @change-template="changeTemplate" />
        <InSensor v-model:model-value="sensorDataVisible" :sensor-id="sensorDataId" />
        <ThreshodEditAdd ref="thresholdEddAddRef" />
    </div>
</template>

<style scoped></style>
