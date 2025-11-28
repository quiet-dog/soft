<script lang='ts' setup>
import { ref } from 'vue';
import { useAlarmHook } from '.';
import AlarmEventDetail from "@/components/alarm-event-detail/index.vue";
import alarm from '@/api/manage/alarm';
import { Message } from '@arco-design/web-vue';

const { crud, columns, crudRef, loadMore } = useAlarmHook()
const alarmEventDetailRef = ref<InstanceType<typeof AlarmEventDetail>>()




function selectOperation(value: "detail", id: number) {
    if (value == "detail") {
        alarmEventDetailRef.value?.handleOpen([id])
    } else if (value == "alarmRelease") {
        alarm.lift(id).then(res => {
            Message.success("报警解除成功")
        })
    }
}




</script>

<template>
    <div class="ma-content-block lg:flex justify-between p-4">
        <ma-crud :options="crud" :columns="columns" ref="crudRef">


            <!-- 报警等级 -->
            <template #level="{ record }">
                <a-tag :color="record.color">
                    {{ record.level }}
                </a-tag>
            </template>


            <!-- 报警状态 -->
            <template #isLift="{ record }">
                <a-tag :color="record.isLift ? 'green' : 'red'">
                    {{ record.isLift ? '报警解除' : '报警触发' }}
                </a-tag>
            </template>
            <!-- 操作之后的 -->
            <template #operationAfterExtend="{ record }">
                <a-dropdown @select="selectOperation($event, record.id)" trigger="hover">
                    <a-link><icon-double-right /> 更多</a-link>
                    <template #content>
                        <a-doption value="detail" v-auth="['manage:event:list']">查看详情</a-doption>
                        <!-- 报警解除 -->
                        <a-doption value="alarmRelease" v-auth="['manage:alarm:release']">报警解除</a-doption>
                    </template>
                </a-dropdown>
            </template>
        </ma-crud>

        <AlarmEventDetail ref="alarmEventDetailRef" />
    </div>
</template>

<style scoped></style>
