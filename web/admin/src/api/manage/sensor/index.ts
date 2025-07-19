import { http } from "@/utils/http";
import { BaseIds, Page, TreeLeaf } from "../base";
import { ReadData, SensorEdit, SensorRow, SensorSearch, TemplateEnv, Translate } from "./types";

export default {
    list: (params: SensorSearch) => {
        return http<Page<SensorRow>>({
            url: '/manage/sensor/index',
            method: 'get',
            params
        });
    },
    save: (params: SensorEdit) => {
        return http({
            url: '/manage/sensor/save',
            method: 'post',
            data: params
        });
    },
    deletes: (params: BaseIds) => {
        return http({
            url: '/manage/sensor/delete',
            method: 'delete',
            data: params
        });
    },
    tree: (params: SensorSearch) => {
        return http<TreeLeaf[]>({
            url: '/manage/sensor/tree',
            method: 'get',
            params
        });
    },
    readData: (data: ReadData) => {
        return http<TemplateEnv>({
            url: "/manage/sensor/readData",
            method: "post",
            data,
        })
    },
    translate: (data: Translate) => {
        return http({
            url: "/manage/sensor/translate",
            method: "post",
            data
        })
    }
}