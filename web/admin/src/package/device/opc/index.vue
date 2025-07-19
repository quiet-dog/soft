<script lang='ts' setup>
import { TreeLeaf } from '@/api/manage/base';
import opc from '@/api/manage/opc';
import { onMounted, reactive, ref } from 'vue';

const visible = defineModel({
    type: Boolean,
    default: true,
})

const extend = ref({
    id: 0,
})

const emit = defineEmits<{
    (e: 'changeExtend', value: { id: number }): void;
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

function onClose(){
    visible.value = false;
    emit("changeExtend", extend.value);
}

onMounted(() => {
    opc.treeLazy(serverId, 0).then(res => {
        console.log("aaa",res.data);
        treeData.value = res.data!;
    })
})

</script>

<template>
    <AModal @close="onClose" v-model:visible="visible" title="OPC节点配置">
        <a-tree :checked-strategy="'child'" @select="select" :data="treeData" :load-more="loadMore" :field-names="{
            key:'value',
            title: 'label',
            children: 'children'
        }" />
    </AModal>
</template>

<style scoped></style>
