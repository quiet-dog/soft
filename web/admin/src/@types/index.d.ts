export interface Column {
    title: string;
    dataIndex: string;
    width?: number;
    hide?: boolean;
    search?: boolean;
    formType?: 'upload' | 'tree-select' | 'select' | 'radio' | 'textarea' | 'range' | 'password' | 'input-number';
    returnType?: 'hash';
    type?: 'image' | 'password';
    rounded?: boolean;
    labelWidth?: string;
    addDisabled?: boolean;
    editDisabled?: boolean;
    commonRules?: Array<{
        required?: boolean;
        message?: string;
        match?: RegExp;
        type?: 'email';
    }>;
    addDefaultValue?: string | number;
    editDefaultValue?: string | ((record: any) => Promise<string[] | string>);
    multiple?: boolean;
    treeCheckable?: boolean;
    treeCheckStrictly?: boolean;
    dict?: {
        url?: string;
        name?: string;
        props?: {
            label: string;
            value: string;
        };
        translation?: boolean;
    };
    addDisplay?: boolean;
    editDisplay?: boolean;
    autocomplete?: 'off';
    addRules?: Array<{
        required?: boolean;
        message?: string;
    }>;
}

export interface Crud {
    api: (params?: any) => Promise<any>;
    recycleApi?: (params?: any) => Promise<any>;
    searchColNumber?: number;
    showIndex?: boolean;
    pageLayout?: 'fixed' | string;
    rowSelection?: { showCheckedAll?: boolean };
    operationColumn?: boolean;
    operationColumnWidth?: number;
    add?: {
        show: boolean;
        api: (data: any) => Promise<any>;
        auth: string[];
    };
    edit?: {
        show: boolean;
        api: (data: any) => Promise<any>;
        auth: string[];
    };
    delete?: {
        show: boolean;
        api: (ids: number[]) => Promise<any>;
        auth: string[];
        realApi?: (ids: number[]) => Promise<any>;
        realAuth?: string[];
    };
    recovery?: {
        show: boolean;
        api: (ids: number[]) => Promise<any>;
        auth: string[];
    };
    import?: {
        show: boolean;
        url: string;
        templateUrl: string;
        auth: string[];
    };
    export?: {
        show: boolean;
        url: string;
        auth: string[];
    };
    formOption?: {
        width: number;
        layout: Array<{
            formType: 'grid';
            cols: Array<{
                span: number;
                formList: Array<{ dataIndex: string }>;
            }>;
        }>;
    };
    isDbClickEdit?: boolean;
    beforeOpenEdit?: (record: any) => boolean;
    beforeDelete?: (ids: number[]) => boolean;
}