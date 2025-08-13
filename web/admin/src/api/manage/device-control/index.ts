import { http } from "@/utils/http";
import { BaseIds, Page, TreeLeaf } from "../base";
import { DeviceControlEdit, DeviceControlRead, DeviceControlRow, DeviceControlSearch, DeviceControlTest } from "./types";
import qs from "qs";

export default {
    list: (params: DeviceControlSearch) => {
        return http<Page<DeviceControlRow>>({
            url: '/manage/deviceControl/index',
            method: 'get',
            params,
            paramsSerializer: params => {
                return qs.stringify(params, { arrayFormat: 'repeat' })
            }
        });
    },
    save: (params: DeviceControlEdit) => {
        return http({
            url: '/manage/deviceControl/save',
            method: 'post',
            data: params
        });
    },
    deletes: (params: BaseIds) => {
        return http({
            url: '/manage/deviceControl/delete',
            method: 'delete',
            data: params
        });
    },
    tree: (params: DeviceControlSearch) => {
        return http<TreeLeaf[]>({
            url: '/manage/deviceControl/tree',
            method: 'get',
            params
        });
    },
    read: (id: number) => {
        return http<DeviceControlRead>({
            url: '/manage/deviceControl/read/' + id,
            method: 'get',
        });
    },
    addControl: (data: DeviceControlRow[]) => {
        return http({
            url: '/manage/deviceControl/addControl',
            method: 'post',
            data: {
                command: data
            }
        });
    },
    control(data: number) {
        return http({
            url: "/manage/deviceControl/control",
            method: "get",
            params: {
                id: data
            }
        })
    }
}