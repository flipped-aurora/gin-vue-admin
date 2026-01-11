import service from '@/utils/request'

export const getShopPluginList = (params) => {
  return service({
    baseURL: "plugin",
    url: '/shopPlugin/getShopPluginList',
    method: 'get',
    params
  })
}