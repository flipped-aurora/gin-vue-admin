import service from '@/utils/request'

export const createArticle = (data) => service({ url: '/article/createArticle', method: 'post', data })
export const deleteArticle = (params) => service({ url: '/article/deleteArticle', method: 'delete', params })
export const deleteArticleByIds = (params) => service({ url: '/article/deleteArticleByIds', method: 'delete', params })
export const updateArticle = (data) => service({ url: '/article/updateArticle', method: 'put', data })
export const findArticle = (params) => service({ url: '/article/findArticle', method: 'get', params })
export const readArticle = (params) => service({ url: '/article/readArticle', method: 'get', params })
export const getArticleList = (params) => service({ url: '/article/getArticleList', method: 'get', params })

export const createCategory = (data) => service({ url: '/article/createCategory', method: 'post', data })
export const deleteCategory = (params) => service({ url: '/article/deleteCategory', method: 'delete', params })
export const updateCategory = (data) => service({ url: '/article/updateCategory', method: 'put', data })
export const getCategoryList = (params) => service({ url: '/article/getCategoryList', method: 'get', params })
export const getAllCategories = () => service({ url: '/article/getAllCategories', method: 'get' })
