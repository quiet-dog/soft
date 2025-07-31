import { http } from "@/utils/http"
import { BaseId, BaseIds, Page, } from "../base"
import { AreaEdit, AreaRow, AreaSearch, AreaTree } from "./types"

export default {
    list(params: AreaSearch) {
        return http<Page<AreaRow>>({
            url: '/manage/area/index',
            method: 'get',
            params
        })
    },
    save(params: AreaEdit) {
        return http({
            url: '/manage/area/save',
            method: 'post',
            data: params
        })
    },
    deletes(params: BaseIds) {
        return http({
            url: '/manage/area/delete',
            method: 'delete',
            data: params
        })
    },
    tree(params: AreaSearch) {
        return http<AreaTree[]>({
            url: '/manage/area/tree',
            method: 'get',
            params
        })
    },
    update(id: number, data: AreaEdit) {
        return http({
            url: '/manage/area/update/' + id,
            method: 'put',
            data: data
        })
    },
    treeTop(id: number) {
        return http({
            url: "/manage/area/treeTop",
            method: 'get',
            params: {
                parentId: id
            }
        })
    }
}