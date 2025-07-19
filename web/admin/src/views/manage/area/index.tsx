
import MaTreeSlider from '@cps/ma-treeSlider/index.vue'
import { onMounted, reactive, ref, watch } from 'vue';
import area from '@/api/manage/area/index'
import { BasicColumn, BasicCrud } from '@/components/ma-crud/types';
import { AreaTree } from '@/api/manage/area/types';
import { TreeNodeData } from "@arco-design/web-vue/es/tree/interface";


export async function getTreeAreaChildren(id): Promise<AreaTree[]> {
    try {
        const res = await area.tree({
            parentId: id,
            name: '',
            sort: 0,
            remark: '',
        })
        if (res.data && res.data.length > 0) {
            return res.data.map(item => {
                if (item.isLeaf) {
                    return { label: item.label, value: item.value }
                }
                return { label: item.label, value: item.value, isLeaf: true, children: [] }
            })
        } else {
            return []
        }
    } catch (err) {
        return []
    }
}
export function useAreaHook() {


    const areaList = ref<TreeNodeData[]>([{ title: '全部', key: 0, children: [], }])
    const defaultKey = ref([0])
    const crudRef = ref()

    const changeAreaId = async () => {
        try {
            const res = await area.tree({
                parentId: defaultKey.value[0],
                name: '',
                sort: 0,
                remark: '',
            })
            if (res.data && res.data.length > 0) {
                areaList.value[0].children = res.data.map(item => {
                    if (item.isLeaf) {
                        return { title: item.label, key: item.value }
                    }
                    return { title: item.label, key: item.value, isLeaf: true }
                })
            } else {
            }
        } catch (error) {

        }
    }



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



    const curd = reactive<BasicCrud>({
        api: area.list,
        // searchColNumber: 3,
        showIndex: false,
        // pageLayout: 'fixed',
        rowSelection: { showCheckedAll: true },
        operationColumn: true,
        // operationColumnWidth: 200,
        add: { show: true, api: area.save, auth: ['manage:area:save'] },
        beforeRequest(params) {
            params.parentId = defaultKey.value[0]
        },
        delete: {
            show: true,
            api: area.deletes, auth: ['system:area:delete'],
        },
    })

    const selectAreaList = ref<AreaTree[]>([
        { label: '全部', value: 0, isLeaf: true }
    ])

    const columns = reactive<BasicColumn[]>([
        { title: 'ID', dataIndex: 'id', addDisplay: false, editDisplay: false, width: 50, hide: true },
        {
            title: '区域名称', dataIndex: 'name', addDisplay: true, editDisplay: true, width: 150, search: true, commonRules: [
                { required: true, message: '请输入区域名称' },
            ]
        },
        {
            title: '父级区域', dataIndex: 'parentId', hide: true, addDisplay: true, editDisplay: true, width: 150, formType: 'tree-select',
            dict: {
                async data() {
                    const list = await getTreeAreaChildren(0)
                    return list
                }
            },
            loadMore(option) {
                return new Promise(async (resolve) => {
                    const result = await getTreeAreaChildren(option.value)
                    option.children = result
                    resolve(result)
                })
            },
        },
        { title: '排序', dataIndex: 'sort', addDisplay: true, editDisplay: true, width: 150, formType: "input-number" },
        { title: '备注', dataIndex: 'remark', addDisplay: true, editDisplay: true, width: 150, formType: 'textarea', },
    ])



    function selectArea(val) {
        console.log('selectArea', val);
        defaultKey.value = val
        // crudRef.value?.requestParams.parentId = val[0]
        crudRef.value?.refresh()
    }


    onMounted(() => {
        changeAreaId()
    })

    return {
        areaList,
        defaultKey,
        crudRef,
        curd,
        columns,
        selectAreaList,
        loadAreaTree,
        selectArea,
    }
}