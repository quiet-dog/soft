<template>
    <AModal @close="close" v-model:visible="visible" width="800px">

        <iframe :src="filePath" height="800px" width="100%"></iframe>
    </AModal>

</template>

<script lang="ts" setup>
import { ref } from 'vue';

const visible = ref(false);
const filePath = ref('');
function base64Encode(str: string): string {
    const bytes = new TextEncoder().encode(str);
    let binary = '';
    bytes.forEach(b => binary += String.fromCharCode(b));
    return btoa(binary);
}


function open(url: string) {
    visible.value = true;
    url = location.origin + import.meta.env.VITE_APP_PROXY_PREFIX + url;
    filePath.value = location.origin + import.meta.env.VITE_APP_PROXY_PREFIX + "/kkfileview/onlinePreview?url=" + encodeURIComponent(base64Encode(url));
}

function close() {
    filePath.value = "";
}

defineExpose({
    open
})

</script>

<style scoped></style>