import { findSysDictionary } from '@/api/sysDictionary'

import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useDictionaryStore = defineStore('dictionary', () => {
  const dictionaryMap = ref({})

  const setDictionaryMap = (dictionaryMap) => {
    dictionaryMap.value = { ...dictionaryMap.value, ...dictionaryMap }
  }

  const getDictionary = async(type) => {
    if (dictionaryMap.value[type] && dictionaryMap.value[type].length) {
      return dictionaryMap.value[type]
    } else {
      const res = await findSysDictionary({ type })
      if (res.code === 0) {
        const dictionaryMap = {}
        const dict = []
        res.data.resysDictionary.sysDictionaryDetails && res.data.resysDictionary.sysDictionaryDetails.forEach(item => {
          dict.push({
            label: item.label,
            value: item.value
          })
        })
        dictionaryMap[res.data.resysDictionary.type] = dict
        setDictionaryMap(dictionaryMap)
        return dictionaryMap.value[type]
      }
    }
  }

  return {
    dictionaryMap,
    setDictionaryMap,
    getDictionary
  }
})
