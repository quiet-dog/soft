import { BaseId, BasePageReq, BaseRow } from '../base';
export interface AreaAdd {
    name: string;
    remark: string;
    parentId: number;
    sort: number;
}

export interface AreaEdit extends AreaAdd, BaseId {

}

export interface AreaRow extends BaseRow, AreaEdit { }

export interface AreaSearch extends AreaAdd, BasePageReq { }

export interface AreaTree {
    label: string;
    value: number;
    children?: AreaTree[];
    isLeaf?: boolean;
}
