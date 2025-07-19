import { http } from "@/utils/http";
import { BaseIds, Page, TreeLeaf } from "../base";
import { SensorTypeEdit, SensorTypeRow, SensorTypeSearch } from "./types";

export default {
    list: (params: SensorTypeSearch) => {
        return http<Page<SensorTypeRow>>({
            url: '/manage/sensorType/index',
            method: 'get',
            params
        });
    },
    save: (params: SensorTypeEdit) => {
        return http({
            url: '/manage/sensorType/save',
            method: 'post',
            data: params
        });
    },
    deletes: (params: BaseIds) => {
        return http({
            url: '/manage/sensorType/delete',
            method: 'delete',
            data: params
        });
    },
    tree: (params: SensorTypeSearch) => {
        return http<TreeLeaf[]>({
            url: '/manage/sensorType/tree',
            method: 'get',
            params
        });
    }
}