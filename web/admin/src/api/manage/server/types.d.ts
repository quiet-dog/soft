import { BaseId, BasePageReq } from "../base";

export interface ServerAdd {
    name: string;
    ip: string;
    port: string;
    username: string;
    password: string;
    remark: string;
    type: string;
}


export interface ServerEdit extends ServerAdd, BaseId {
}


export interface ServerRow extends BaseRow, ServerEdit { }

export interface ServerSearch extends ServerAdd, BasePageReq {
}


export interface ServerTreeLeaf {
    pageInfo: ServerSearch;
    items: TreeLeaf[];
}
