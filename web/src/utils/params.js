import { useParamsStore } from '@/pinia/modules/params'
/*
 * 获取参数方法 使用示例 getParams('key').then(res)  或者 async函数下 const res = await getParams('key')
 *   const res = ref('')
 *   const fun = async () => {
 *       res.value = await getParams('test')
 *   }
 *   fun()
 */
export const getParams = async(key) => {
    const paramsStore = useParamsStore()
    await paramsStore.getParams(key)
    return paramsStore.paramsMap[key]
}
