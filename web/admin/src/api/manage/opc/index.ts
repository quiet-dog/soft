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
    }
}