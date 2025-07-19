<script lang='ts' setup>
import device from '@/api/manage/device';
import opc from '@/api/manage/opc';
import { OpcExtend } from '@/api/manage/sensor/types';
import { reactive, defineProps, onMounted, ref } from 'vue';

const { deviceId = 0 } = defineProps<{
    deviceId: number
}>()

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

const form = reactive({})

const handleOk = () => {
    emit("changeExtend", extend.value);
    visible.value = false;
}
const handleCancel = () => {
    visible.value = false;
}

function onClose(){
    visible.value = false;
}

const selectOptions = ref()

onMounted(()=>{
    device.read(deviceId).then(res=>{
        if(res.data?.extend){
            let extend = res.data.extend as OpcExtend
            opc.treeLazy(res.data.serverId,extend.id).then(r=>{
                // console.log("rrrrrrrrrrr=",r)
                selectOptions.value = r.data
            })
        }
    })
})
</script>

<template>
    <AModal @close="onClose" v-model:visible="visible" title="OPC节点配置" @ok="handleOk" @cancel="handleCancel">
        <AForm :model="extend">
            <ARow>
                <ACol>
                    <AFormItem label="监测点位">
                        <ACascader v-model="extend.id" :options="selectOptions"  />
                    </AFormItem>
                </ACol>
            </ARow>
        </AForm>
    </AModal>
</template>

<style scoped></style>
