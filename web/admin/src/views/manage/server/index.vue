<script lang='ts' setup>
import { onMounted, ref } from 'vue';
import useServerHook from '.';
import server from '@/api/manage/server';
import OpcConfig from "@/package/server/opc/index.vue";

const { crud, columns, crudRef,formType } = useServerHook()
const serverTypes = ref([])
const opcConfigRef = ref<InstanceType<typeof OpcConfig>>()
function changeTypes(value) {
  formType.value = value
  crudRef.value.getFormData().type = value;
  if (value == "opc") {
    opcConfigRef.value?.open(crudRef.value.getFormData().extend)
  }
}

function changeExtend(value) {
  crudRef.value.getFormData().extend = value
}

onMounted(() => {
  server.types().then(res => {
    serverTypes.value = res.data
  })
})
</script>

<template>
  <div class="ma-content-block lg:flex justify-between p-4">
    <ma-crud :options="crud" :columns="columns" ref="crudRef">
      <template #form-type>
        <a-select v-model="formType">
          <!-- @vue-expect-error -->
          <a-option @click="changeTypes(item.value)" v-for="item in serverTypes" :value="item.value">
            <!-- @vue-expect-error -->
            {{ item.label }}
          </a-option>
        </a-select>
      </template>

      <template #isOnline="{ record }">
        <a-tag v-if="record.isOnline" :color="'#00b42a'">
          在线
        </a-tag>
        <a-tag :color="'#f53f3f'" v-else>
          离线
        </a-tag>
      </template>
    </ma-crud>
    <OpcConfig ref="opcConfigRef" @changeExtend="changeExtend" />
  </div>
</template>

<style scoped></style>
