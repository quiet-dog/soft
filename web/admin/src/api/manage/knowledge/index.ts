// js api
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

import { request } from '@/utils/request.js'

export default {

    /**
    * 获取manageKnowledge分页列表
    * @returns
    */
    getPageList(params = {}) {
        return request({
            url: 'manage/manageKnowledge/index',
            method: 'get',
            params
        })
    },

    getList(params = {}) {
        return request({
            url: 'manage/manageKnowledge/list',
            method: 'get',
            params
        })
    },

    /**
    * 从回收站获取manageKnowledge数据列表
    * @returns
    */
    getPageRecycleList(params = {}) {
        return request({
            url: 'manage/manageKnowledge/recycle',
            method: 'get',
            params
        })
    },

    getRecycleList(params = {}) {
        return request({
            url: 'manage/manageKnowledge/recycleList',
            method: 'get',
            params
        })
    },

    /**
    * 恢复manageKnowledge数据
    * @returns
    */
    recoverys(data) {
        return request({
            url: 'manage/manageKnowledge/recovery',
            method: 'put',
            data
        })
    },

    /**
    * 真实删除manageKnowledge
    * @returns
    */
    realDeletes(data) {
        return request({
            url: 'manage/manageKnowledge/realDelete',
            method: 'delete',
            data
        })
    },

    /**
    * 添加manageKnowledge
    * @returns
    */
    save(data = {}) {
        return request({
            url: 'manage/manageKnowledge/save',
            method: 'post',
            data
        })
    },

    /**
    * 更新manageKnowledge数据
    * @returns
    */
    update(id, data = {}) {
        return request({
            url: 'manage/manageKnowledge/update/' + id,
            method: 'put',
            data
        })
    },

    /**
    * 读取manageKnowledge
    * @returns
    */
    read(id) {
        return request({
            url: 'manage/manageKnowledge/read/' + id,
            method: 'get'
        })
    },

    /**
    * 将manageKnowledge删除，有软删除则移动到回收站
    * @returns
    */
    deletes(data) {
        return request({
            url: 'manage/manageKnowledge/delete',
            method: 'delete',
            data
        })
    },

    /**
    * 更改manageKnowledge数据
    * @returns
    */
    changeStatus(data = {}) {
        return request({
            url: 'manage/manageKnowledge/changeStatus',
            method: 'put',
            data
        })
    },

    /**
    * 修改manageKnowledge数值数据，自增自减
    * @returns
    */
    numberOperation(data = {}) {
        return request({
            url: 'manage/manageKnowledge/numberOperation',
            method: 'put',
            data
        })
    },

    /**
    * manageKnowledge导入
    * @returns
    */
    importExcel(data = {}) {
        return request({
            url: 'manage/manageKnowledge/import',
            method: 'post',
            data
        })
    },

    /**
    * manageKnowledge下载模板
    * @returns
    */
    downloadTemplate() {
        return request({
            url: 'manage/manageKnowledge/downloadTemplate',
            method: 'post',
            responseType: 'blob'
        })
    },

    /**
    * manageKnowledge导出
    * @returns
    */
    exportExcel(params = {}) {
        return request({
            url: 'manage/manageKnowledge/export',
            method: 'post',
            responseType: 'blob',
            params
        })
    },

}