import { BaseRow, BasePageReq } from '../base';

export interface AlarmRow extends BaseRow {
    sensorId: number;
    level: string;
    endTime?: number;
    sendTime?: number;
    isLift?: boolean;
}


export interface SearchAlarm extends BasePageReq, AlarmRow {

}