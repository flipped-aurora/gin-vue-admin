import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useAppStore = defineStore('app', () => {
  const device = ref('')
  const drawerSize = ref('')
  const operateMinWith = ref('240')
  // 侧栏折叠态（运行时态，不持久化），供 vertical 布局的固定侧栏与外壳共享
  const sideCollapse = ref(false)
  // 移动端（<640）左侧抽屉菜单开合态
  const mobileMenuOpen = ref(false)
  // 当前断点区间（由 useResponsive 维护）：mobile(<640) / pad(640–1024) / desktop(>=1024)
  const layoutZone = ref('desktop')
  // 平板区间强制「通栏侧边」收缩布局（不改用户所选布局）
  const isPad = computed(() => layoutZone.value === 'pad')
  // 进入平板前的侧栏折叠态快照，离开平板时据此恢复（区分「系统收缩」与「用户收缩」）
  const prePadCollapse = ref(false)

  const toggleMobileMenu = (v) => {
    mobileMenuOpen.value = typeof v === 'boolean' ? v : !mobileMenuOpen.value
  }

  // 切换断点区间：进入平板时快照当前折叠态并强制收缩，离开平板时恢复进入前的折叠态，
  // 从而只撤销系统自己施加的收缩，不覆盖用户在桌面手动收缩的意图。
  const setLayoutZone = (zone) => {
    const prev = layoutZone.value
    if (zone === 'pad' && prev !== 'pad') {
      prePadCollapse.value = sideCollapse.value
      sideCollapse.value = true
    } else if (zone !== 'pad' && prev === 'pad') {
      sideCollapse.value = prePadCollapse.value
    }
    layoutZone.value = zone
  }

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
    mobileMenuOpen,
    toggleMobileMenu,
    layoutZone,
    isPad,
    setLayoutZone,
    toggleDevice
  }
})
