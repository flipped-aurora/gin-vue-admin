import { useRoute } from 'vue-router'
import { reactive } from 'vue'
export const useBtnAuth = () => {
  const route = useRoute()
  return route.meta.btns || reactive({})
}
