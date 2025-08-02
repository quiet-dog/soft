import { BaseRow, BasePageReq } from '../base';

export interface AlarmRow extends BaseRow {
    sensorId: number;
    level: string;
}


export interface SearchAlarm extends BasePageReq, AlarmRow {

}