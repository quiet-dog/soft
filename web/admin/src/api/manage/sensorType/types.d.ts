import { BaseId, BasePageReq } from "../base";

export interface SensorTypeAdd {
    name: string;
    unit: string;
}


export interface SensorTypeEdit extends SensorTypeAdd, BaseId {
}


export interface SensorTypeRow extends BaseRow, SensorTypeEdit {

}

export interface SensorTypeSearch extends SensorTypeAdd, BasePageReq {
}

export interface SensorTypeTreeLeaf {
    pageInfo: SensorTypeSearch;
    items: TreeLeaf[];
}
