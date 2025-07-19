import { BaseId, BasePageReq } from "../base";
import sensorType from '@/api/manage/sensorType';

export interface SensorAdd {
    name: string;
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


export interface TemplateEnv {
    value: any;
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