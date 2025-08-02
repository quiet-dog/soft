import { number } from "echarts"
import { BasePageReq, BaseRow } from "../base";

export interface SearchEvent extends BasePageReq {
    alarmIds: number[];
    eventIds: number[];
    sensorIds: number[];
}


export interface EventRow extends BaseRow {
    sensorId: number;
    value: number;
    level: string;
    color: string;
    description: string;
}