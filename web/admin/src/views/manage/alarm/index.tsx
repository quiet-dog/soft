import alarm from "@/api/manage/alarm";
import event from "@/api/manage/event";
import { BasicColumn, BasicCrud } from "@/components/ma-crud/types";
import { reactive, ref } from "vue";

export function useAlarmHook() {

    const crud = reactive<BasicCrud>({
        api: alarm.list,
        operationColumn: true,
    })

    const columns = reactive<BasicColumn[]>([
        { title: "报警编号", dataIndex: 'id', },
        { title: "报警等级", dataIndex: 'level' },
        { title: "传感器名称", dataIndex: 'sensorId' },
        { title: "创建时间", dataIndex: "createdAt", },
    ])

    const crudRef = ref()

    const loadMore = (record, done) => {
        console.log("=====", record)
        // @ts-ignore
        event.list({
            page: 1,
            pageSize: 100,
            alarmIds: [record.id]
        }).then(res => {
            res.data?.items.forEach(item => {
                // @ts-ignore
                item.isLeaf = true
            })
            done(res.data?.items)
        })
    }

    return {
        crud,
        columns,
        crudRef,
        loadMore
    }
}