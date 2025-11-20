import { http } from "@/utils/http";
import { TreeLeaf } from "../base";

export default {
    treeLazy: (serverId: number, parentId: number) => {
        return http<TreeLeaf[]>({
            url: '/manage/opc/tree',
            method: 'get',
            params: {
                serverId,
                parentId
            }
        });
    },
    nodeIdIsExit(serverId: number, nodeId: string) {
        return http<number>({
            url: "/manage/opc/isExit",
            method: 'post',
            data: {
                serverId: serverId,
                nodeId: nodeId
            }
        })
    },
}