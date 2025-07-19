import { BaseId, BasePageReq } from "../base";

export interface AlarmLabelAdd {
    name: string;
    remark: string;
    color: string;
    lebel: string;
}


export interface AlarmLabelEdit extends AlarmLabelAdd, BaseId {
}


export interface AlarmLabelRow extends BaseRow, AlarmLabelEdit { }

export interface AlarmLabelSearch extends AlarmLabelAdd, BasePageReq {
}
