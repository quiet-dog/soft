import { http } from "@/utils/http";
import { BaseIds, Page } from "../base";
import { AlarmLabelRow, AlarmLabelSearch } from "./types";

export default {
    list: (params: AlarmLabelSearch) => {
        return http<Page<AlarmLabelRow>>({
            url: '/manage/alarmLabel/index',
            method: 'get',
            params
        })
    },
    save: (params: AlarmLabelRow) => {
        return http({
            url: '/manage/alarmLabel/save',
            method: 'post',
            data: params
        });
    },
    deletes: (params: BaseIds) => {
        return http({
            url: '/manage/alarmLabel/delete',
            method: 'delete',
            data: params
        });
    }
}