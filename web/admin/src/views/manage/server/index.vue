<script lang='ts' setup>
import useServerHook from '.';

const { crud, columns, crudRef } = useServerHook()
</script>

<template>
  <div class="ma-content-block lg:flex justify-between p-4">
    <ma-crud :options="crud" :columns="columns" ref="crudRef">
      <template #isOnline="{ record }">
        <a-tag v-if="record.isOnline" :color="'#00b42a'">
          在线
        </a-tag>
        <a-tag :color="'#f53f3f'" v-else>
          离线
        </a-tag>
      </template>
    </ma-crud>
  </div>
</template>

<style scoped></style>
