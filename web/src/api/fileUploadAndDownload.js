import service from '@/utils/request';
import { Message } from 'element-ui';

const handleFileError = (res, fileName) => {
    if (typeof(res.data) !== "undefined") {
        if (res.data.type == "application/json") {
            const reader = new FileReader();
            reader.onload = function() {
                let message = JSON.parse(reader.result).msg;
                Message({
                    showClose: true,
                    message: message,
                    type: 'error'
                })
            };
            reader.readAsText(new Blob([res.data]));
        } 
    } else {
        var downloadUrl = window.URL.createObjectURL(new Blob([res]));
        var a = document.createElement('a');
        a.style.display = 'none';
        a.href = downloadUrl;
        a.download = fileName;
        var event = new MouseEvent("click");
        a.dispatchEvent(event);
    }        
} 

// @Tags FileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取文件户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
export const getFileList = (data) => {
    return service({
        url: "/fileUploadAndDownload/getFileList",
        method: "post",
        data
    })
}

// @Tags FileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body dbModel.FileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
export const deleteFile = (data) => {
    return service({
        url: "/fileUploadAndDownload/deleteFile",
        method: "post",
        data
    })
}

// @Tags ExaFileUploadAndDownload
// @Summary 导出Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body request.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /fileUploadAndDownload/exportExcel [post]
export const exportExcel = (tableData, fileName) => {
    service({
        url: "/fileUploadAndDownload/exportExcel",
        method: 'post',
        data: {
            fileName: fileName,
            infoList: tableData
        },
        responseType: 'blob'
    }).then((res)=>{
        handleFileError(res, fileName)
    })
}

// @Tags ExaFileUploadAndDownload
// @Summary 导入Excel文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "导入Excel文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
// @Router /fileUploadAndDownload/importExcel [post]
export const loadExcelData = () => {
    return service({
        url: "/fileUploadAndDownload/loadExcel",
        method: 'get'
    })
}

// @Tags ExaFileUploadAndDownload
// @Summary 下载模板
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param fileName query fileName true "模板名称"
// @Success 200
// @Router /fileUploadAndDownload/downloadTemplate [get]
export const downloadTemplate = (fileName) => {   
    return service({
        url: "/fileUploadAndDownload/downloadTemplate",
        method: 'get',
        params:{
            fileName: fileName
        },
        responseType: 'blob'
    }).then((res)=>{
        handleFileError(res, fileName)        
    })
}