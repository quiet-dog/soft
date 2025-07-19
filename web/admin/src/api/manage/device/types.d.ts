import { BaseId, BasePageReq } from "../base";
import { ServerRow } from "../server/types";

export interface DeviceAdd {
    name: string;
    manufacturer: string;
    model: string;
    installationLocation: string;
    areaId: number;
    serverId: number;
}


export interface DeviceEdit extends DeviceAdd, BaseId {
}


export interface DeviceRow extends BaseRow, DeviceEdit {
    areaName: string;
    serverName: string;
    extend?: any
}

export interface DeviceSearch extends DeviceAdd, BasePageReq {
}


export interface DeviceRead extends DeviceRow {
    server?: ServerRow;
}



export interface DeviceTreeLeaf {
    pageInfo: DeviceSearch;
    items: TreeLeaf[];
}
