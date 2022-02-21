import { useRoute } from 'vue-router'
export const useBtnAuth = () => {
  const route = useRoute()
  return route.meta.btns
}
