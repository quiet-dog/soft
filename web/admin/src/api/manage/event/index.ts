import { http } from "@/utils/http"
import { Page } from "../base"
import { EventRow, SearchEvent } from "./types"
import qs from "qs"

export default {
    list(params: SearchEvent) {
        return http<Page<EventRow>>({
            url: "/manage/event/index",
            method: "get",
            params,
            paramsSerializer: (params) => {
                return qs.stringify(params, { arrayFormat: 'repeat' })
            }
        })
    },
}