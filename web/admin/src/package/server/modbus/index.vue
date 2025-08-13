<script lang='ts' setup>
import { ref } from 'vue';

const visible = ref(true)
</script>

<template>
    <AModal v-model:visible="visible">
    </AModal>
</template>

<style scoped></style>
