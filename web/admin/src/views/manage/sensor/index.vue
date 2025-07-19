<script lang='ts' setup>
import { useSensorHook } from '.';
import Template from "@/package/template/index.vue"

const { crud, columns,
    crudRef, changeSensorType, loadSensorOptions,
    sensorTypeOptions, deviceOptions,
    loadDeviceOptions, changeDevice, asyncComponent,
    visible, asyncDeviceId, changeExtend,templateVisable,
    templateInfo,changeTemplate
} = useSensorHook()


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
                <a-select @change="changeDevice" @dropdown-reach-bottom="loadDeviceOptions" placeholder="请选择设备">
                    <a-option v-for="item in deviceOptions.items" :key="item.value" :value="item.value">
                        {{ item.label }}
                    </a-option>
                </a-select>
            </template>
        </ma-crud>
        <component :is="asyncComponent" :deviceId="asyncDeviceId" @changeExtend="changeExtend" />
        <Template v-model:model-value="templateVisable" :templateInfo="templateInfo" @change-template="changeTemplate" />
    </div>
</template>

<style scoped></style>
