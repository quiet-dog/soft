<script lang='ts' setup>
import { ref } from 'vue';
import { useDeviceHook } from '.';
import DeviceControl from "@/components/device-control/index.vue";
import deviceControl from '@/api/manage/device-control';
import { DeviceControlRow } from '@/api/manage/device-control/types';
import { Message } from '@arco-design/web-vue';
import ImportModel from "@/components/import-model/index.vue";
import ModelPreview from "@/components/model-preview/index.vue";
import Fields from "@/package/fields/index.vue";
import NowData from "@/components/now-data/index.vue"

const {
  crud, columns, crudRef,
  serverOptions, loadServerOptions, changeServer,
  asyncServerId, asyncComponent, changeExtend,
  sExtend
} = useDeviceHook()

const importModelRef = ref<InstanceType<typeof ImportModel>>()
const deviceControlRef = ref<InstanceType<typeof DeviceControl>>()
const modelPreviewRef = ref<InstanceType<typeof ModelPreview>>()
const fieldsRef = ref<InstanceType<typeof Fields>>()
const nowDataRef = ref<InstanceType<typeof NowData>>()
function selectOperation(val: string, id: number) {
  if (val == "deviceControl") {
    deviceControlRef.value?.open(id)
  } else if (val == "importModel") {
    importModelRef.value?.open(id)
  } else if (val == "previewModel") {
    modelPreviewRef.value?.open(id)
  } else if (val == "sensorControl") {
    fieldsRef.value?.open(id)
  } else if (val == "nowDataModel") {
    nowDataRef.value?.open(id)
  }
}

const controlCommands = ref<DeviceControlRow[]>([])
function getControl(val: boolean, id: number) {
  if (val) {
    deviceControl.list({
      page: 1,
      pageSize: 100,
      name: "",
      extend: "",
      deviceIds: [id],
      deviceId: 0
    }).then(res => {
      // @ts-expect-error
      controlCommands.value = res.data?.items
    })
  }
}

function selectControl(val: number) {
  deviceControl.control(val).then(res => {
    if (res.code == 0) {
      Message.success("发送成功")
    } else {
      Message.error(res.message)
    }
  })
}

</script>

<template>
  <div class="ma-content-block lg:flex justify-between p-4">
    <ma-crud :options="crud" :columns="columns" ref="crudRef">
      <template #form-serverId>
        <a-select @dropdown-reach-bottom="loadServerOptions" placeholder="请选择服务器">
          <a-option @click="changeServer(item.value)" v-for="item in serverOptions.items" :key="item.value"
            :value="item.value">
            {{ item.label }}
          </a-option>
        </a-select>
      </template>
      <!-- 操作之前的 -->
      <template #operationBeforeExtend="{ record }">
        <a-dropdown @select="selectControl" @popup-visible-change="getControl($event, record.id)" trigger="hover">
          <a-link><icon-double-left /> 控制</a-link>
          <template #content>
            <a-doption v-for="item in controlCommands" :value="item.id">
              {{ item.name }}
            </a-doption>
          </template>
        </a-dropdown>
      </template>

      <!-- 操作之后的 -->
      <template #operationAfterExtend="{ record }">
        <a-dropdown @select="selectOperation($event, record.id)" trigger="hover">
          <a-link><icon-double-right /> 更多</a-link>
          <template #content>
            <a-doption value="sensorControl" v-auth="['manage:sensorControl:save']">信息配置</a-doption>
            <a-doption value="deviceControl" v-auth="['manage:deviceControl:save']">设备控制</a-doption>
            <a-doption value="importModel" v-auth="['manage:device:importModel']">导入模型</a-doption>
            <a-doption value="previewModel" v-auth="['manage:device:previewModel']">查看模型</a-doption>
            <a-doption value="nowDataModel" v-auth="['manage:device:nowData']">当前数据</a-doption>
          </template>
        </a-dropdown>
      </template>
    </ma-crud>
    <component :is="asyncComponent" :serverId="asyncServerId" @changeExtend="changeExtend" :sExtend="sExtend" />
    <DeviceControl ref="deviceControlRef" />
    <ImportModel ref="importModelRef" />
    <ModelPreview ref="modelPreviewRef" />
    <Fields ref="fieldsRef" />
    <NowData ref="nowDataRef" />
  </div>
</template>

<style scoped></style>
