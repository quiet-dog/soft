<script lang='ts' setup>
import useServerHook from '.';

const { crud, columns, crudRef } = useServerHook()
</script>

<template>
  <div class="ma-content-block lg:flex justify-between p-4">
        <ma-crud :options="crud" :columns="columns" ref="crudRef"></ma-crud>
    </div>
</template>

<style scoped>

</style>
