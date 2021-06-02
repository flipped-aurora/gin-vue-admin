import { store } from '@/store'
//  获取字典方法 使用示例 getDict('sex').then(res)  或者 async函数下 const res = await getDict('sex')
export const getDict = async(type) => {
  await store.dispatch('dictionary/getDictionary', type)
  return store.getters['dictionary/getDictionary'][type]
}
