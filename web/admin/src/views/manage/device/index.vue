<script lang='ts' setup>
import { useDeviceHook } from '.';

const {
  crud, columns, crudRef,
  serverOptions, loadServerOptions, changeServer,
  asyncServerId, asyncComponent, changeExtend
} = useDeviceHook()


</script>

<template>
  <div class="ma-content-block lg:flex justify-between p-4">
    <ma-crud :options="crud" :columns="columns" ref="crudRef">
      <template #form-serverId>
        <a-select @change="changeServer" @dropdown-reach-bottom="loadServerOptions" placeholder="请选择服务器">
          <a-option v-for="item in serverOptions.items" :key="item.value" :value="item.value">
            {{ item.label }}
          </a-option>
        </a-select>
      </template>
    </ma-crud>
    <component :is="asyncComponent" :serverId="asyncServerId" @changeExtend="changeExtend" />
  </div>
</template>

<style scoped></style>
