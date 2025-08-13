import { http } from "@/utils/http";
import { BaseIds, Page, TreeLeaf } from "../base";
import { ServerEdit, ServerRow, ServerSearch } from "./types";

export default {
    list: (params: ServerSearch) => {
        return http<Page<ServerRow>>({
            url: '/manage/server/index',
            method: 'get',
            params
        });
    },
    save: (params: ServerEdit) => {
        return http({
            url: '/manage/server/save',
            method: 'post',
            data: params
        });
    },
    deletes: (params: BaseIds) => {
        return http({
            url: '/manage/server/delete',
            method: 'delete',
            data: params
        });
    },
    tree: (params: ServerSearch) => {
        return http<TreeLeaf[]>({
            url: '/manage/server/tree',
            method: 'get',
            params
        });
    },
    read: (id: number) => {
        return http<ServerRow>({
            url: '/manage/server/read/' + id,
            method: 'get',
        });
    },
    update(id: number, data: ServerEdit) {
        return http({
            url: "/manage/server/update/" + id,
            method: 'put',
            data
        })
    },
    types() {
        return http({
            url: "/manage/server/types",
            method: 'get',
        })
    }
}