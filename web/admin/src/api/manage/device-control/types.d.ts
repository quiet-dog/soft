import { BaseId, BasePageReq } from "../base";
import { ServerRow } from "../server/types";

export interface DeviceControlAdd {
    deviceId: number;
    name: string;
    extend: any;
}


export interface DeviceControlEdit extends DeviceControlAdd, BaseId {
}


export interface DeviceControlRow extends BaseRow, DeviceControlEdit {

}

export interface DeviceControlSearch extends DeviceControlAdd, BasePageReq {
    deviceIds?: number[];
}


export interface DeviceControlRead extends DeviceControlRow {
}



export interface DeviceControlTreeLeaf {
    pageInfo: DeviceControlSearch;
    items: TreeLeaf[];
}




export interface DeviceControlTest {
    serverId: number;
    extend: string;
}