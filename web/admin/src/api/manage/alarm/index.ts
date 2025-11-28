import { http } from "@/utils/http"
import { Page } from "../base"
import { AlarmRow, SearchAlarm } from "./types"

export default {
    list(params: SearchAlarm) {
        return http<Page<AlarmRow>>({
            url: "/manage/alarm/index",
            method: "get",
            params
        })
    },
    // 报警解除
    lift(id: number) {
        return http({
            url: "/manage/alarm/lift",
            method: "post",
            data: {
                id: id
            }
        })
    },
    read(id: number) {
        return http<AlarmRow>({
            url: "/manage/alarm/read/" + id,
            method: "get"
        })
    }

}