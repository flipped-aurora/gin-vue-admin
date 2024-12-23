import { useDictionaryStore } from '@/pinia/modules/dictionary'
import { getSysParam } from '@/api/sysParams'
//  获取字典方法 使用示例 getDict('sex').then(res)  或者 async函数下 const res = await getDict('sex')
export const getDict = async (type) => {
  const dictionaryStore = useDictionaryStore()
  await dictionaryStore.getDictionary(type)
  return dictionaryStore.dictionaryMap[type]
}

//  字典文字展示方法
export const showDictLabel = (
  dict,
  code,
  keyCode = 'value',
  valueCode = 'label'
) => {
  if (!dict) {
    return ''
  }
  const dictMap = {}
  dict.forEach((item) => {
    if (Reflect.has(item, keyCode) && Reflect.has(item, valueCode)) {
      dictMap[item[keyCode]] = item[valueCode]
    }
  })
  return Reflect.has(dictMap, code) ? dictMap[code] : ''
}

export const getParams = async (key) => {
  const res = await getSysParam({ key })
  if (res.code === 0) {
    return res.data.value
  }
}
