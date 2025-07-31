<script lang='ts' setup>
import { useSensorHook } from '.';
import Template from "@/package/template/index.vue"
import InSensor from "@/components/in-sensor/index.vue"
import { ref } from 'vue';

const { crud, columns,
    crudRef, changeSensorType, loadSensorOptions,
    sensorTypeOptions, deviceOptions,
    loadDeviceOptions, changeDevice, asyncComponent,
    visible, asyncDeviceId, changeExtend, templateVisable,
    templateInfo, changeTemplate,sExtend
} = useSensorHook()

const sensorDataVisible = ref(false)
const sensorDataId = ref(0)
function changeSensorData(id: number) {
    sensorDataId.value = id
    sensorDataVisible.value = true
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
                <a-select  @dropdown-reach-bottom="loadDeviceOptions" placeholder="请选择设备">
                    <a-option @click="changeDevice(item.value)" v-for="item in deviceOptions.items" :key="item.value" :value="item.value">
                        {{ item.label }}
                    </a-option>
                </a-select>
            </template>

            <template #view="{ record }">
                <AButton type="primary" @click="changeSensorData(record.id)">查看</AButton>
            </template>
        </ma-crud>
        <component :is="asyncComponent" :deviceId="asyncDeviceId" @changeExtend="changeExtend" :sExtend="sExtend" />
        <Template v-model:model-value="templateVisable" :templateInfo="templateInfo"
            @change-template="changeTemplate" />
        <InSensor v-model:model-value="sensorDataVisible" :sensor-id="sensorDataId" />
    </div>
</template>

<style scoped></style>
