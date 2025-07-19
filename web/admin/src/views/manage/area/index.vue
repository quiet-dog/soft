<script lang='ts' setup>
import {useAreaHook} from '.';


const { areaList, loadAreaTree, selectArea, curd, columns,crudRef } = useAreaHook()

</script>

<template>
    <div class="ma-content-block lg:flex justify-between p-4">
        <div class="">
            <a-tree :data="areaList" :load-more="loadAreaTree" @select="selectArea" />
        </div>
        <div class="lg:w-10/12 w-full lg:ml-4 mt-5 lg:mt-0">
            <ma-crud :options="curd" :columns="columns" ref="crudRef">
            </ma-crud>
        </div>
    </div>
</template>

<style scoped></style>
