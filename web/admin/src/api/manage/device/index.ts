import { http } from "@/utils/http";
import { BaseIds, Page, TreeLeaf } from "../base";
import { DeviceEdit, DeviceInfo, DeviceRead, DeviceRow, DeviceSearch, DeviceTest } from "./types";
import { SensorRow } from "../sensor/types";

export default {
    list: (params: DeviceSearch) => {
        return http<Page<DeviceRow>>({
            url: '/manage/device/index',
            method: 'get',
            params
        });
    },
    save: (params: DeviceEdit) => {
        return http({
            url: '/manage/device/save',
            method: 'post',
            data: params
        });
    },
    deletes: (params: BaseIds) => {
        return http({
            url: '/manage/device/delete',
            method: 'delete',
            data: params
        });
    },
    tree: (params: DeviceSearch) => {
        return http<TreeLeaf[]>({
            url: '/manage/device/tree',
            method: 'get',
            params
        });
    },
    read: (id: number) => {
        return http<DeviceRead>({
            url: '/manage/device/read/' + id,
            method: 'get',
        });
    },
    readSensorInfo: (id: number) => {
        return http<DeviceInfo>({
            url: '/manage/device/readSensorInfo/' + id,
            method: 'get',
        });
    },
    test: (data: DeviceTest) => {
        return http({
            url: '/manage/device/test',
            method: 'post',
            data
        })
    },
    importModel(deviceId: number, path: string) {
        return http({
            url: '/manage/device/importModel',
            method: 'post',
            data: {
                deviceId,
                path
            }
        })
    },
    saveSensorInfo(data: SensorRow[]) {
        return http({
            url: "/manage/device/saveSensorInfo",
            method: "post",
            data
        })
    }
}