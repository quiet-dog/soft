import { BaseId, BasePageReq } from "../base";
import sensorType from '@/api/manage/sensorType';
import { ThresholdRow } from "../threshold/types";

export interface SensorAdd {
    name: string;
    deviceId: number;
    extend?: any;
    remark?: string;
    sensorTypeId: number;
    unit?: string,
    template?: string
}


export interface SensorEdit extends SensorAdd, BaseId {
}


export interface SensorRow extends BaseRow, SensorEdit {

}

export interface SensorSearch extends SensorAdd, BasePageReq {
}


export interface SensorAlarmRow extends SensorRow {
    thresholds: ThresholdRow[]
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



export interface ReadSensorEchart extends BasePageReq {
    sensorId: number;
    // deviceId: number;
    beginTime: string | number;
    endTime: string | number;
}


export interface SensorEchart {
    sensorId: number;
    deviceId: number;
    sensorName: string;
    sensorTypeName: string;
    sensorTypeId: string;
    xData: string[];
    eSeiresData: number[];
    cSeiresData: number[];
}


export interface ReadHistory extends BasePageReq {
    sensorId?: number;
    beginTime?: number;
    endTime?: number;
}


export interface InfluxdbRow {
    [key: string]: any;
    sensor: string;
    time: string
}
export interface SensorData extends SensorRow {
    rows: InfluxdbRow[];
    total: number;
    sensorUnit?: string;
    sensorName?: string;
}