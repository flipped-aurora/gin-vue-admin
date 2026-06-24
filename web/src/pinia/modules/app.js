import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAppStore = defineStore('app', () => {
  const device = ref('')
  const drawerSize = ref('')
  const operateMinWith = ref('240')
  // 侧栏折叠态（运行时态，不持久化），供 vertical 布局的固定侧栏与外壳共享
  const sideCollapse = ref(false)

  const toggleDevice = (e) => {
    if (e === 'mobile') {
      drawerSize.value = '100%'
      operateMinWith.value = '80'
    } else {
      drawerSize.value = '800'
      operateMinWith.value = '240'
    }
    device.value = e
  }

  const toggleSideCollapse = (v) => {
    sideCollapse.value = typeof v === 'boolean' ? v : !sideCollapse.value
  }

  return {
    device,
    drawerSize,
    operateMinWith,
    sideCollapse,
    toggleSideCollapse,
    toggleDevice
  }
})
