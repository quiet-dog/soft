<script lang='ts' setup>
import { ref } from 'vue';
import tool from '@/utils/tool';
import device from '@/api/manage/device';
import { Message } from '@arco-design/web-vue';

const env = import.meta.env;
const headers = ref({
    "authorization": tool.local.get(env.VITE_APP_TOKEN_PREFIX),
})


function success(file) {
    if (file.response.code == 0) {
        device.importModel(deviceId.value, file.response.data.url)
        Message.success("导入成功")
    } else {
        Message.error("导入失败")
    }
}

const visible = ref(false)
const deviceId = ref(0)
function open(id: number) {
    visible.value = true;
    deviceId.value = id;
}

defineExpose({
    open
})

</script>

<template>
    <AModal v-model:visible="visible">
        <AUpload @success="success" :headers="headers" with-credentials action="/dev/system/uploadFile" />
    </AModal>
</template>

<style scoped></style>
