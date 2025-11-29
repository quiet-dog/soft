import area from "@/api/manage/area";
import device from "@/api/manage/device";
import { DeviceRowsHaveSensors } from "@/api/manage/device/types";
import { TreeNodeData } from "@arco-design/web-vue/es/tree/interface";
import { ref } from "vue";

import Msg from '@/ws-serve/msg'

export function useAllHook() {

    const Wsm = new Msg()

    const areaList = ref<TreeNodeData[]>([{ title: '全部', key: 0, children: [], }])
    const deviceList = ref<DeviceRowsHaveSensors[]>([])
    const total = ref(0)
    const loading = ref(false)
    const searchParams = ref({
        page: 1,
        pageSize: 10,
        areaIds: []
    })

    function loadAreaTree(nodeData) {
        return new Promise(async (resolve) => {
            console.log('loadAreaTree', nodeData);
            const res = await area.tree({
                parentId: Number(nodeData.key),
                name: '',
                sort: 0,
                remark: '',
            })
            if (res.data && res.data.length > 0) {
                const list = res.data.map(item => {
                    if (item.isLeaf) {
                        return { title: item.label, key: item.value }
                    }
                    return { title: item.label, key: item.value, isLeaf: true, children: [] }
                })
                nodeData.children = list
                resolve(list)
            } else {
                resolve([])
            }
        })
    }

    function getDeviceList() {
        loading.value = true
        // @ts-expect-error
        device.getPageListForSearchHaveSensors(searchParams.value).then(res => {
            console.log('getDeviceList', res);
            deviceList.value = res.data?.rows!
            total.value = res.data?.total!
        }).finally(() => {
            loading.value = false
        })
    }
    const lastSubscribe = ref("")

    function selectArea(val) {
        searchParams.value.areaIds = val
        getDeviceList()
        const toSubscribe = "area_data_" + val[0]
        // 取消订阅上一个订阅
        if (lastSubscribe.value) {
            // @ts-ignore
            Wsm.ws.unsubscribe(lastSubscribe.value)
        }
        // @ts-ignore
        Wsm.ws.subscribe(toSubscribe, (data, ws) => {
            alert("123123")
        })
        lastSubscribe.value = toSubscribe
    }




    return {
        loadAreaTree,
        areaList,
        searchParams,
        getDeviceList,
        deviceList,
        loading,
        selectArea
    }
}