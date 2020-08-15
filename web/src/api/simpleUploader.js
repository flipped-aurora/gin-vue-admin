

import service from '@/utils/request'

// @Tags SimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth

// @Produce  application/json
// @Param params md5 get "测试文件是否已经存在和判断已经上传过的切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /simpleUploader/checkFileMd5 [get]
export const checkFileMd5 = (params) => {
    return service({
        url: "/simpleUploader/checkFileMd5",
        method: 'get',
        params
    })
}


// @Tags SimpleUploader
// @Summary 合并文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param params md5 get "合并文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"合并成功"}"
// @Router /simpleUploader/mergeFileMd5 [get]
export const mergeFileMd5 = (params) => {
    return service({
        url: "/simpleUploader/mergeFileMd5",
        method: 'get',
        params
    })
}

