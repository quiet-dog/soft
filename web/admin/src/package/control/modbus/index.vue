<script lang='ts' setup>
import { ref } from 'vue';

const visible = ref(false)
const emit = defineEmits(['changeExtend'])

const num = ref(1)
const writeValues = ref<string[]>(['0'])
const copyWriteValues = ref<number[]>([0])
const extend = ref({
    startAddr: "0x0001",
    value: []
})

function changeNum(val: number) {
    // 假设 writeValues.value 是 number[] 数组
    if (writeValues.value.length < val) {
        // 长度小于 val，往数组尾部追加 0x01
        while (writeValues.value.length < val) {
            writeValues.value.push('0');
            copyWriteValues.value.push(0)
        }
    } else if (writeValues.value.length > val) {
        // 长度大于 val，删除多余元素
        while (writeValues.value.length > val) {
            writeValues.value.pop();
            copyWriteValues.value.pop();
        }
    }
}

function changeWrite(v: string, index: number) {
    // writeValues.value[index] = Number.parseInt(v, 16);
    copyWriteValues.value[index] = Number.parseInt(v, 16);
    console.log("writeValues", writeValues.value, index)
}

function handleOk() {
    // @ts-expect-error
    extend.value.value = copyWriteValues.value
    emit("changeExtend", extend.value)
    visible.value = false
}

function open(value) {
    visible.value = true
    console.log("value", value)
    if (value && value != null && value != undefined) {
        extend.value = value;
        if (Array.isArray(extend.value.value)) {
            writeValues.value = extend.value.value.map((item: number) => {
                return item.toString(16).toUpperCase()
            })
            copyWriteValues.value = extend.value.value.map(item => item)
            num.value = writeValues.value.length;
        }
    } else {
        writeValues.value = ['0']
        copyWriteValues.value = [0]
        extend.value = {
            startAddr: "0x0001",
            // @ts-expect-error
            value: [0],
        }
    }
}
defineExpose({
    open
})
</script>

<template>
    <AModal @ok="handleOk" v-model:visible="visible">
        <AForm :model="extend">
            <ARow>
                <ACol>
                    <AFormItem prop="startAddr" label="起始地址">
                        <AInput v-model="extend.startAddr" placeholder="请输入起始地址">
                            <template #prepend>
                                0x
                            </template>
                        </AInput>
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow>
                <ACol>
                    <AFormItem prop="num" label="写入数量">
                        <AInputNumber :default-value="num" @change="changeNum" :step="1" mode="button" />
                    </AFormItem>
                </ACol>
            </ARow>

            <ARow v-for="(item, index) in writeValues" :key="index">
                <ACol>
                    <AFormItem>
                        <AInput v-model="writeValues[index]" @input="(v) => changeWrite(v, index)">
                            <template #prepend>
                                0x
                            </template>
                        </AInput>
                    </AFormItem>
                </ACol>
            </ARow>
        </AForm>
    </AModal>
</template>

<style scoped></style>
