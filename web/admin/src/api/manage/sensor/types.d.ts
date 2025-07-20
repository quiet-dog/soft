import { BaseId, BasePageReq } from "../base";
import sensorType from '@/api/manage/sensorType';

export interface SensorAdd {
    name: string;
    deviceId: number;
    extend?: any;
    remark?: string;
    sensorTypeId: number;
}


export interface SensorEdit extends SensorAdd, BaseId {
}


export interface SensorRow extends BaseRow, SensorEdit {

}

export interface SensorSearch extends SensorAdd, BasePageReq {
}


export interface ReadData {
    extend: any;
    type: string;
}

export interface Value {
    value: any;
}

export interface TemplateEnv {
    value: Value;
    type: string;
    createTime: string
}


export interface Translate {
    env: TemplateEnv;
    template: string;
}

export interface SensorTreeLeaf {
    pageInfo: SensorSearch;
    items: TreeLeaf[];
}


export interface OpcExtend {
    id: number
}

export interface InfluxdbData extends BasePageReq {
    sensorIds: number[];
    deviceId: number;
    beginTime: string;
    endTime: string;
}