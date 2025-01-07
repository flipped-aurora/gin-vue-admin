import service from '@/utils/request'
// 分类列表
export const getCategoryList = () => {
    return service({
        url: '/attachmentCategory/getCategoryList',
        method: 'get',
    })
}

// 添加/编辑分类
export const addCategory = (data) => {
    return service({
        url: '/attachmentCategory/addCategory',
        method: 'post',
        data
    })
}

// 删除分类
export const deleteCategory = (data) => {
    return service({
        url: '/attachmentCategory/deleteCategory',
        method: 'post',
        data
    })
}
