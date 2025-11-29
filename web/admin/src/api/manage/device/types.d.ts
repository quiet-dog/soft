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
    extend?: any;
    modelPath?: string;
}

export interface DeviceRowsHaveSensors extends DeviceRow {
    sensors: SensorRow[];
}

export interface GetPageListForSearchHaveSensors {
    rows: DeviceRowsHaveSensors[];
    total: number;
}

export interface DeviceSearch extends DeviceAdd, BasePageReq {
    areaIds?: number[];
}


export interface DeviceRead extends DeviceRow {
    server?: ServerRow;
}



export interface DeviceTreeLeaf {
    pageInfo: DeviceSearch;
    items: TreeLeaf[];
}

export interface DeviceInfo extends DeviceRow {
    sensors?: SensorRow[];
    server?: ServerRow;
}

export interface DeviceTest {
    serverId: number;
    extend: string;
}