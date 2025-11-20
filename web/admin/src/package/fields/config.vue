<script lang='ts' setup>
import opc from '@/api/manage/opc';
import { Message } from '@arco-design/web-vue';

import { ref } from 'vue';

const { type, info, serverId } = defineProps<{
    type?: string;
    info?: any;
    serverId?: number
}>()

const extend = ref()
const opcConf = ref("")

function handleConfirm(done) {
    if (type == "opc") {
        opc.nodeIdIsExit(serverId!, opcConf.value).then(res => {
            if (res.data! > 0) {
                extend.value.id = res.data
                extend.value.nodeId = opcConf.value
                done(true)
            } else {
                Message.warning("该节点无法使用")
            }
        })
    }
}

function ok() {
    emit("success", extend.value)
}

function popupVisibleChange(val) {
    if (val) {
        extend.value = info
        if (type === 'opc') {
            opcConf.value = extend.value.nodeId
        }
    }
}

const emit = defineEmits(["success"])

</script>

<template>
    <APopconfirm @ok="ok" @popup-visible-change="popupVisibleChange" :onBeforeOk="handleConfirm">
        <template #content>
            <AInput v-model="opcConf" placeholder="填写点位信息" v-if="type === 'opc'" />
        </template>
        <AButton type="text">已绑定</AButton>
    </APopconfirm>
</template>

<style scoped></style>
