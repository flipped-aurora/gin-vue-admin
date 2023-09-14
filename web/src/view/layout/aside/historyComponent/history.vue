<template>
  <div class="router-history">
    <el-tabs
      v-model="activeValue"
      :closable="!(historys.length === 1 && $route.name === defaultRouter)"
      type="card"
      @contextmenu.prevent="openContextMenu($event)"
      @tab-click="changeTab"
      @tab-remove="removeTab"
    >
      <el-tab-pane
        v-for="item in historys"
        :key="name(item)"
        :label="item.meta.title"
        :name="name(item)"
        :tab="item"
        class="gva-tab"
      >
        <template #label>
          <span
            :tab="item"
            :style="{
              color: activeValue === name(item) ? userStore.activeColor : '#333',
            }"
          ><i
             class="dot"
             :style="{
               backgroundColor:
                 activeValue === name(item) ? userStore.activeColor : '#ddd',
             }"
           />
            {{ fmtTitle(item.meta.title,item) }}</span>
        </template>
      </el-tab-pane>
    </el-tabs>

    <!--自定义右键菜单html代码-->
    <ul
      v-show="contextMenuVisible"
      :style="{ left: left + 'px', top: top + 'px' }"
      class="contextmenu"
    >
      <li @click="closeAll">关闭所有</li>
      <li @click="closeLeft">关闭左侧</li>
      <li @click="closeRight">关闭右侧</li>
      <li @click="closeOther">关闭其他</li>
    </ul>
  </div>
</template>


<script setup>
import { emitter } from '@/utils/bus.js'
import { computed, onUnmounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { fmtTitle } from '@/utils/fmtRouterTitle'

defineOptions({
  name: 'HistoryComponent',
})

const route = useRoute()
const router = useRouter()

const getFmtString = (item) => {
  return item.name + JSON.stringify(item.query) + JSON.stringify(item.params)
}

const historys = ref([])
const activeValue = ref('')
const contextMenuVisible = ref(false)

const userStore = useUserStore()

const name = (item) => {
  return (
    item.name + JSON.stringify(item.query) + JSON.stringify(item.params)
  )
}

const left = ref(0)
const top = ref(0)
const isCollapse = ref(false)
const isMobile = ref(false)
const rightActive = ref('')
const defaultRouter = computed(() => userStore.userInfo.authority.defaultRouter)
const openContextMenu = (e) => {
  if (
    historys.value.length === 1 &&
        route.name === defaultRouter.value
  ) {
    return false
  }
  let id = ''
  if (e.srcElement.nodeName === 'SPAN') {
    id = e.srcElement.offsetParent.id
  } else {
    id = e.srcElement.id
  }
  if (id) {
    contextMenuVisible.value = true
    let width
    if (isCollapse.value) {
      width = 54
    } else {
      width = 220
    }
    if (isMobile.value) {
      width = 0
    }
    left.value = e.clientX - width
    top.value = e.clientY + 10
    rightActive.value = id.substring(4)
  }
}
const closeAll = () => {
  historys.value = [
    {
      name: defaultRouter.value,
      meta: {
        title: '首页',
      },
      query: {},
      params: {},
    },
  ]
  router.push({ name: defaultRouter.value })
  contextMenuVisible.value = false
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
const isSame = (route1, route2) => {
  if (route1.name !== route2.name) {
    return false
  }
  if (Object.keys(route1.query).length !== Object.keys(route2.query).length || Object.keys(route1.params).length !== Object.keys(route2.params).length) {
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

const historyMap = ref({})

const changeTab = (TabsPaneContext) => {
  const name = TabsPaneContext?.props?.name
  if (!name) return
  const tab = historyMap.value[name]
  router.push({
    name: tab.name,
    query: tab.query,
    params: tab.params,
  })
}
const removeTab = (tab) => {
  const index = historys.value.findIndex(
    (item) => getFmtString(item) === tab
  )
  if (getFmtString(route) === tab) {
    if (historys.value.length === 1) {
      router.push({ name: defaultRouter.value })
    } else {
      if (index < historys.value.length - 1) {
        router.push({
          name: historys.value[index + 1].name,
          query: historys.value[index + 1].query,
          params: historys.value[index + 1].params,
        })
      } else {
        router.push({
          name: historys.value[index - 1].name,
          query: historys.value[index - 1].query,
          params: historys.value[index - 1].params,
        })
      }
    }
  }
  historys.value.splice(index, 1)
}

watch(() => contextMenuVisible.value, () => {
  if (contextMenuVisible.value) {
    document.body.addEventListener('click', () => {
      contextMenuVisible.value = false
    })
  } else {
    document.body.removeEventListener('click', () => {
      contextMenuVisible.value = false
    })
  }
})

watch(() => route, (to, now) => {
  if (to.name === 'Login' || to.name === 'Reload') {
    return
  }
  historys.value = historys.value.filter((item) => !item.meta.closeTab)
  setTab(to)
  sessionStorage.setItem('historys', JSON.stringify(historys.value))
  activeValue.value = window.sessionStorage.getItem('activeValue')
}, { deep: true })

watch(() => historys.value, () => {
  sessionStorage.setItem('historys', JSON.stringify(historys.value))
  historyMap.value = {}
  historys.value.forEach((item) => {
    historyMap.value[getFmtString(item)] = item
  })
  emitter.emit('setKeepAlive', historys.value)
}, {
  deep: true
})

const initPage = () => {
  // 全局监听 关闭当前页面函数
  emitter.on('closeThisPage', () => {
    removeTab(name(route))
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
  const initHistorys = [
    {
      name: defaultRouter.value,
      meta: {
        title: '首页',
      },
      query: {},
      params: {},
    },
  ]
  historys.value =
      JSON.parse(sessionStorage.getItem('historys')) || initHistorys
  if (!window.sessionStorage.getItem('activeValue')) {
    activeValue.value = getFmtString(route)
  } else {
    activeValue.value = window.sessionStorage.getItem('activeValue')
  }
  setTab(route)
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

<style lang="scss">
.contextmenu {
  width: 100px;
  margin: 0;
  border: 1px solid #ccc;
  background: #fff;
  z-index: 3000;
  position: absolute;
  list-style-type: none;
  padding: 5px 0;
  border-radius: 4px;
  font-size: 14px;
  color: #333;
  box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.2);
}
.el-tabs__item .el-icon-close {
  color: initial !important;
}
.el-tabs__item .dot {
  content: "";
  width: 9px;
  height: 9px;
  margin-right: 8px;
  display: inline-block;
  border-radius: 50%;
  transition: background-color 0.2s;
}

.contextmenu li {
  margin: 0;
  padding: 7px 16px;
}
.contextmenu li:hover {
  background: #f2f2f2;
  cursor: pointer;
}
</style>
