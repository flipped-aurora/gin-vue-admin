import { store } from '@/store/index'

export const getDict = async (type) => {
    await store.dispatch("dictionary/getDictionary", type)
    return store.getters["dictionary/getDictionary"][type]
}