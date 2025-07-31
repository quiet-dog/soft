<script lang='ts' setup>
import { TreeLeaf } from '@/api/manage/base';
import opc from '@/api/manage/opc';
import { ref } from 'vue';
import { ExtendType } from ".";

const { sExtend } = defineProps<{
    sExtend: ExtendType
}>()

const visible = defineModel({
    type: Boolean,
    default: true,
})

const extend = ref({
    id: 0,
})

const emit = defineEmits<{
    (e: 'changeExtend', value: ExtendType): void;
}>();


const treeData = ref<TreeLeaf[]>([]);
const { serverId = 0 } = defineProps<{
    serverId: number
}>()


const loadMore = (nodeData) => {
    return new Promise((resolve) => {
        console.log("loadMore", nodeData);
        opc.treeLazy(serverId, nodeData.value).then(res => {
            if (res.data) {
                nodeData.children = res.data;
            }
            resolve(null);
        });
    });
};

const select = (selectedKeys, e) => {
    extend.value.id = selectedKeys[0] || 0;
    visible.value = false;
}

function onClose() {
    visible.value = false;
    emit("changeExtend", extend.value);
}

function handleOpen() {
    opc.treeLazy(serverId, 0).then(res => {
        console.log("aaa", res.data);
        treeData.value = res.data!;
    })
    if (sExtend != undefined && sExtend != null) {
        extend.value = JSON.parse(JSON.stringify(sExtend))
    } else {
        extend.value.id = 0
    }

}


</script>

<template>
    <AModal @open="handleOpen" @close="onClose" v-model:visible="visible" title="OPC节点配置">
        <a-tree :checked-strategy="'child'" @select="select" :data="treeData" :load-more="loadMore" :field-names="{
            key: 'value',
            title: 'label',
            children: 'children'
        }" />
    </AModal>
</template>

<style scoped></style>
