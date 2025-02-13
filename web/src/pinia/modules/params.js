import { getSysParam } from '@/api/sysParams'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useParamsStore = defineStore('params', () => {
    const paramsMap = ref({})

    const setParamsMap = (paramsRes) => {
        paramsMap.value = { ...paramsMap.value, ...paramsRes }
    }

    const getParams = async(key) => {
        if (paramsMap.value[key] && paramsMap.value[key].length) {
            return paramsMap.value[key]
        } else {
            const res = await getSysParam({ key })
            if (res.code === 0) {
                const paramsRes = {}
                paramsRes[key] = res.data.value
                setParamsMap(paramsRes)
                return paramsMap.value[key]
            }
        }
    }

    return {
        paramsMap,
        setParamsMap,
        getParams
    }
})
