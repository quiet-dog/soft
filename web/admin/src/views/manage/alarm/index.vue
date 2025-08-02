<script lang='ts' setup>
import { ref } from 'vue';
import { useAlarmHook } from '.';
import AlarmEventDetail from "@/components/alarm-event-detail/index.vue";

const { crud, columns, crudRef, loadMore } = useAlarmHook()
const alarmEventDetailRef = ref<InstanceType<typeof AlarmEventDetail>>()


function selectOperation(value: "detail", id: number) {
    if (value == "detail") {
        alarmEventDetailRef.value?.handleOpen([id])
    }
}




</script>

<template>
    <div class="ma-content-block lg:flex justify-between p-4">
        <ma-crud :options="crud" :columns="columns" ref="crudRef">


            <!-- 操作之后的 -->
            <template #operationAfterExtend="{ record }">
                <a-dropdown @select="selectOperation($event, record.id)" trigger="hover">
                    <a-link><icon-double-right /> 更多</a-link>
                    <template #content>
                        <a-doption value="detail" v-auth="['manage:event:list']">查看详情</a-doption>
                    </template>
                </a-dropdown>
            </template>
        </ma-crud>

        <AlarmEventDetail ref="alarmEventDetailRef" />
    </div>
</template>

<style scoped></style>
