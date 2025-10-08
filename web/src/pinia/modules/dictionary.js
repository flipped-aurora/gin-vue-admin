import { findSysDictionary } from '@/api/sysDictionary'

import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useDictionaryStore = defineStore('dictionary', () => {
  const dictionaryMap = ref({})

  const setDictionaryMap = (dictionaryRes) => {
    dictionaryMap.value = { ...dictionaryMap.value, ...dictionaryRes }
  }

  const getDictionary = async (type) => {
    if (dictionaryMap.value[type] && dictionaryMap.value[type].length) {
      return dictionaryMap.value[type]
    } else {
      const res = await findSysDictionary({ type })
      if (res.code === 0) {
        const dictionaryRes = {}
        const dict = []
        res.data.resysDictionary.sysDictionaryDetails &&
          res.data.resysDictionary.sysDictionaryDetails.forEach((item) => {
            dict.push({
              label: item.label,
              value: item.value,
              extend: item.extend
            })
          })
        dictionaryRes[res.data.resysDictionary.type] = dict
        setDictionaryMap(dictionaryRes)
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
