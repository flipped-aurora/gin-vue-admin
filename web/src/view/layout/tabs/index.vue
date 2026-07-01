<template>
  <div
    ref="scrollRef"
    class="gva-tabs isolate flex items-center gap-1.5 h-11 w-full overflow-x-auto bg-[var(--gva-tabs-bg)] px-3 pt-1.5 shadow-[var(--gva-tabs-shadow)]"
    :class="containerClass"
  >
    <ContextMenuRoot
      v-for="item in historys"
      :key="getFmtString(item)"
      :modal="false"
      @update:open="(open) => onContextMenuOpen(open, item)"
    >
      <ContextMenuTrigger as-child :disabled="contextMenuDisabled">
        <g-page-tab
          :mode="tabMode"
          :active="isActive(item)"
          :closable="isClosable(item)"
          @click="switchTo(item)"
          @mousedown="middleCloseTab($event, item)"
          @close="removeTab(getFmtString(item))"
        >
          <template v-if="item.meta.icon" #prefix>
            <component :is="item.meta.icon" class="h-4 w-4 shrink-0" />
          </template>
          {{ fmtTitle(item.meta.title, item) }}
        </g-page-tab>
      </ContextMenuTrigger>

      <ContextMenuPortal>
        <ContextMenuContent
          :side-offset="4"
          class="z-[3000] min-w-[120px] overflow-hidden rounded-md border border-border bg-container py-1 text-[13px] text-base-text shadow-card"
        >
          <ContextMenuItem
            v-for="action in contextMenuActions"
            :key="action.key"
            class="cursor-pointer select-none px-4 py-1.5 outline-none transition-colors data-[highlighted]:bg-muted"
            @select="action.handler"
          >
            {{ action.label }}
          </ContextMenuItem>
        </ContextMenuContent>
      </ContextMenuPortal>
    </ContextMenuRoot>
  </div>
</template>

<script setup>
  import { emitter } from '@/utils/bus.js'
  import { computed, onUnmounted, ref, watch, nextTick } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { useUserStore } from '@/pinia/modules/user'
  import { useThemeStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  import { fmtTitle } from '@/utils/fmtRouterTitle'
  import {
    ContextMenuRoot,
    ContextMenuTrigger,
    ContextMenuPortal,
    ContextMenuContent,
    ContextMenuItem
  } from 'reka-ui'

  defineOptions({
    name: 'HistoryComponent'
  })

  const route = useRoute()
  const router = useRouter()
  const userStore = useUserStore()
  const themeStore = useThemeStore()
  const { settings } = storeToRefs(themeStore)

  // 标签风格（button / chrome / slider），由主题设置驱动，可换肤
  const tabMode = computed(() => settings.value.tab.mode || 'button')
  const containerClass = computed(
    () =>
      ({
        button: 'items-center',
        chrome: 'items-end gap-0',
        slider: 'items-end gap-1'
      })[tabMode.value]
  )


  const getFmtString = (item) => {
    return item.name + JSON.stringify(item.query) + JSON.stringify(item.params)
  }

  const historys = ref([])
  const activeValue = ref('')
  const rightActive = ref('')

  const isCollapse = ref(false)
  const isMobile = ref(false)
  const defaultRouter = computed(
    () => userStore.userInfo.authority.defaultRouter
  )

  // 常驻标签（首页 / 默认路由）不可关闭：不渲染关闭按钮、也不占位；其它标签始终可关。
  const isClosable = (item) => item.name !== defaultRouter.value
  // 右键菜单禁用条件：仅“唯一首页”时右键不弹出（保留原 el-tabs 行为）。
  const contextMenuDisabled = computed(
    () => historys.value.length === 1 && route.name === defaultRouter.value
  )

  const isActive = (item) => getFmtString(item) === activeValue.value

  const scrollRef = ref(null)
  // 路由变化后把激活标签滚入可视区（el-tabs 原本免费提供，这里显式补回）
  const scrollActiveIntoView = async () => {
    await nextTick()
    if (!scrollRef.value) return
    const index = historys.value.findIndex(isActive)
    scrollRef.value.children[index]?.scrollIntoView({
      behavior: 'smooth',
      inline: 'nearest',
      block: 'nearest'
    })
  }

  // 左键：切换路由（关闭按钮已 stop，不会触达这里）
  const switchTo = (item) => {
    router.push({
      name: item.name,
      query: item.query,
      params: item.params
    })
  }

  // 中键：关闭标签。保留唯一首页不可关的守卫与 preventDefault（抑制浏览器自动滚动）
  const middleCloseTab = (e, item) => {
    if (e.button !== 1) return
    if (!isClosable(item)) return
    e.preventDefault()
    removeTab(getFmtString(item))
  }

  // 右键菜单打开时记录目标标签（替代旧的 e.srcElement.id.substring(4) DOM 取值）
  const onContextMenuOpen = (open, item) => {
    if (open) {
      rightActive.value = getFmtString(item)
    }
  }

  const closeAll = () => {
    historys.value = [
      {
        name: defaultRouter.value,
        meta: {
          title: '首页'
        },
        query: {},
        params: {}
      }
    ]
    router.push({ name: defaultRouter.value })
    sessionStorage.setItem('historys', JSON.stringify(historys.value))
  }
  const closeLeft = () => {
    let right
    const rightIndex = historys.value.findIndex((item) => {
      if (getFmtString(item) === rightActive.value) {
        right = item
      }
      return getFmtString(item) === rightActive.value
    })
    const activeIndex = historys.value.findIndex(
      (item) => getFmtString(item) === activeValue.value
    )
    historys.value.splice(0, rightIndex)
    if (rightIndex > activeIndex) {
      router.push(right)
    }
    sessionStorage.setItem('historys', JSON.stringify(historys.value))
  }
  const closeRight = () => {
    let right
    const leftIndex = historys.value.findIndex((item) => {
      if (getFmtString(item) === rightActive.value) {
        right = item
      }
      return getFmtString(item) === rightActive.value
    })
    const activeIndex = historys.value.findIndex(
      (item) => getFmtString(item) === activeValue.value
    )
    historys.value.splice(leftIndex + 1, historys.value.length)
    if (leftIndex < activeIndex) {
      router.push(right)
    }
    sessionStorage.setItem('historys', JSON.stringify(historys.value))
  }
  const closeOther = () => {
    let right
    historys.value = historys.value.filter((item) => {
      if (getFmtString(item) === rightActive.value) {
        right = item
      }
      return getFmtString(item) === rightActive.value
    })
    router.push(right)
    sessionStorage.setItem('historys', JSON.stringify(historys.value))
  }

  const contextMenuActions = [
    { key: 'all', label: '关闭所有', handler: closeAll },
    { key: 'left', label: '关闭左侧', handler: closeLeft },
    { key: 'right', label: '关闭右侧', handler: closeRight },
    { key: 'other', label: '关闭其他', handler: closeOther }
  ]

  const isSame = (route1, route2) => {
    if (route1.name !== route2.name) {
      return false
    }
    if (
      Object.keys(route1.query).length !== Object.keys(route2.query).length ||
      Object.keys(route1.params).length !== Object.keys(route2.params).length
    ) {
      return false
    }
    for (const key in route1.query) {
      if (route1.query[key] !== route2.query[key]) {
        return false
      }
    }
    for (const key in route1.params) {
      if (route1.params[key] !== route2.params[key]) {
        return false
      }
    }
    return true
  }
  const setTab = (route) => {
    if (!historys.value.some((item) => isSame(item, route))) {
      const obj = {}
      obj.name = route.name
      obj.meta = { ...route.meta }
      delete obj.meta.matched
      obj.query = route.query
      obj.params = route.params
      historys.value.push(obj)
    }
    window.sessionStorage.setItem('activeValue', getFmtString(route))
  }

  const removeTab = (tab) => {
    const index = historys.value.findIndex((item) => getFmtString(item) === tab)
    if (getFmtString(route) === tab) {
      if (historys.value.length === 1) {
        router.push({ name: defaultRouter.value })
      } else {
        if (index < historys.value.length - 1) {
          router.push({
            name: historys.value[index + 1].name,
            query: historys.value[index + 1].query,
            params: historys.value[index + 1].params
          })
        } else {
          router.push({
            name: historys.value[index - 1].name,
            query: historys.value[index - 1].query,
            params: historys.value[index - 1].params
          })
        }
      }
    }
    historys.value.splice(index, 1)
  }

  watch(
    () => route,
    (to) => {
      if (to.name === 'Login' || to.name === 'Reload') {
        return
      }
      historys.value = historys.value.filter((item) => !item.meta.closeTab)
      setTab(to)
      sessionStorage.setItem('historys', JSON.stringify(historys.value))
      activeValue.value = window.sessionStorage.getItem('activeValue')
    },
    { deep: true }
  )

  watch(
    () => historys.value,
    () => {
      sessionStorage.setItem('historys', JSON.stringify(historys.value))
      emitter.emit('setKeepAlive', historys.value)
    },
    {
      deep: true
    }
  )

  // 激活标签变化时滚动到可视区
  watch(activeValue, scrollActiveIntoView)

  const initPage = () => {
    // 全局监听 关闭当前页面函数
    emitter.on('closeThisPage', () => {
      removeTab(getFmtString(route))
    })
    // 全局监听 关闭所有页面函数
    emitter.on('closeAllPage', () => {
      closeAll()
    })
    emitter.on('mobile', (data) => {
      isMobile.value = data
    })
    emitter.on('collapse', (data) => {
      isCollapse.value = data
    })

    emitter.on('setQuery', (data) => {
      const index = historys.value.findIndex(
        (item) => getFmtString(item) === activeValue.value
      )
      historys.value[index].query = data
      activeValue.value = getFmtString(historys.value[index])
      const currentUrl = window.location.href.split('?')[0]
      const currentSearchParams = new URLSearchParams(data).toString()
      window.history.replaceState(
        {},
        '',
        `${currentUrl}?${currentSearchParams}`
      )
      sessionStorage.setItem('historys', JSON.stringify(historys.value))
    })

    emitter.on('switchTab', async (data) => {
      const index = historys.value.findIndex((item) => item.name === data.name)
      if (index < 0) {
        return
      }
      for (const key in data.query) {
        data.query[key] = String(data.query[key])
      }
      for (const key in data.params) {
        data.params[key] = String(data.params[key])
      }

      historys.value[index].query = data.query || {}
      historys.value[index].params = data.params || {}
      await nextTick()
      router.push(historys.value[index])
    })
    const initHistorys = [
      {
        name: defaultRouter.value,
        meta: {
          title: '首页'
        },
        query: {},
        params: {}
      }
    ]
    setTab(route)
    historys.value =
      JSON.parse(sessionStorage.getItem('historys')) || initHistorys
    if (!window.sessionStorage.getItem('activeValue')) {
      activeValue.value = getFmtString(route)
    } else {
      activeValue.value = window.sessionStorage.getItem('activeValue')
    }
    if (window.sessionStorage.getItem('needCloseAll') === 'true') {
      closeAll()
      window.sessionStorage.removeItem('needCloseAll')
    }
  }
  initPage()

  onUnmounted(() => {
    emitter.off('collapse')
    emitter.off('mobile')
  })
</script>

<style scoped>
  /* 横向滚动条不占位（保留可滚动，仅隐藏滚动条本体） */
  .gva-tabs::-webkit-scrollbar {
    height: 0;
  }
</style>
