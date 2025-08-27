<script lang='ts' setup>
import { ref } from 'vue';


const visible = ref(false)
const emit = defineEmits(['changeExtend'])
const extend = ref({
    url: "",
    speed: 2,
    dataBits: 8,
    stopBits: 1,
    parity: 0
})

function open(data) {
    if (data != undefined && data != null) {
        extend.value = data
    }
    visible.value = true
}

function handleOk() {
    emit('changeExtend', extend.value)
    visible.value = false
}

defineExpose({
    open
})


</script>

<template>
    <!-- rtu的配置  var speed uint = 2
	var dataBits uint = 8
	var stopBits uint = 1
	var parity uint = modbus.PARITY_NONE-->
    <AModal @ok="handleOk" v-model:visible="visible">
        <AForm>
            <ARow>
                <ACol>
                    <AFormItem label="地址" name="url">
                        <AInput v-model="extend.url" />
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow>
                <ACol>
                    <AFormItem label="波特率" name="speed">
                        <AInputNumber v-model="extend.speed" />
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow>
                <ACol>
                    <AFormItem label="数据位" name="dataBits">
                        <AInputNumber v-model="extend.dataBits" />
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow>
                <ACol>
                    <AFormItem label="停止位" name="stopBits">
                        <AInputNumber v-model="extend.stopBits" />
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow>
                <ACol>
                    <AFormItem label="校验位" name="parity">
                        <ASelect v-model="extend.parity">
                            <AOption value="0">无</AOption>
                            <AOption value="1">奇校验</AOption>
                            <AOption value="2">偶校验</AOption>
                        </ASelect>
                    </AFormItem>
                </ACol>
            </ARow>


        </AForm>
    </AModal>
</template>

<style scoped></style>
