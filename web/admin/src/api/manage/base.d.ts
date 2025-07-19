
export interface Response<T> {
    code: number;
    data?: T;
    message: string;
    path: string;
    requestId: string;
    success: boolean;
    takeUpTime: number;
}


export interface ResponsePage<T> {
    pageInfo: PageInfo;
    items: T[];
}

export interface BasePageReq {
    pageSize?: number;
    page?: number;
}

export interface PageInfo {
    pageSize: number;
    currentPage: number;
    totalPage: number;
    total: number;
}


export interface BaseRow {
    created_by: number;
    created_at: string;
    updated_by: number;
    updated_at: string;
    id: number;
}

export interface BaseId {
    id: number;
}

export interface BaseIds {
    ids: number[];
}


export interface TreeLeaf {
    label: string;
    value: number;
    children?: TreeLeaf[];
    disabled?: boolean;
    isLeaf?: boolean;
}



export interface Page<T> extends ResponsePage<T> {

}