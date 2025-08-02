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

}