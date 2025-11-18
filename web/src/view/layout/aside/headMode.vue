<template>
  <div
    class="bg-white h-[calc(100%-4px)] text-slate-700 dark:text-slate-300 mx-2 dark:bg-slate-900 flex items-center w-[calc(100vw-600px)] overflow-auto"
    ref="menuContainer"
  >
    <el-menu
      :default-active="active"
      mode="horizontal"
      class="!border-r-0 w-full flex gap-1 items-center box-border h-[calc(100%-1px)]"
      unique-opened
      :ellipsis="shouldEllipsis"
      @select="selectMenuItem"
      ref="menuRef"
    >
      <template v-for="item in routerStore.asyncRouters[0].children">
        <aside-component
          v-if="!item.hidden"
          :key="item.name"
          :router-info="item"
          mode="horizontal"
        />
      </template>
    </el-menu>
  </div>
</template>

<script setup>
  import AsideComponent from '@/view/layout/aside/asideComponent/index.vue'
  import { ref, provide, watchEffect, onMounted, nextTick } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  const appStore = useAppStore()
  const { device } = storeToRefs(appStore)

  defineOptions({
    name: 'GvaAside'
  })
  const route = useRoute()
  const router = useRouter()
  const routerStore = useRouterStore()
  const isCollapse = ref(false)
  const active = ref('')
  const menuRef = ref(null)
  const menuContainer = ref(null)
  const shouldEllipsis = ref(false)

  // 计算是否需要启用省略功能
  const calculateEllipsis = async () => {
    await nextTick()
    if (!menuRef.value || !menuContainer.value) return

    const menuItems = menuRef.value.$el.querySelectorAll('.el-menu-item, .el-sub-menu')
    let totalWidth = 0

    menuItems.forEach(item => {
      totalWidth += item.offsetWidth
    })

    const containerWidth = menuContainer.value.offsetWidth
    shouldEllipsis.value = totalWidth > containerWidth
  }

  watchEffect(() => {
    if (route.name === 'gvaLayoutIframe') {
      active.value = decodeURIComponent(route.query.url)
      return
    }
    active.value = route.meta.activeName || route.name
  })

  watchEffect(() => {
    if (device.value === 'mobile') {
      isCollapse.value = true
    } else {
      isCollapse.value = false
    }
    // 设备变化时重新计算
    calculateEllipsis()
  })

  // 当路由变化时重新计算
  watchEffect(() => {
    if (route.name) {
      nextTick(calculateEllipsis)
    }
  })

  provide('isCollapse', isCollapse)

  onMounted(() => {
    calculateEllipsis()
    window.addEventListener('resize', calculateEllipsis)
  })

  const selectMenuItem = (index) => {
    const query = {}
    const params = {}
    routerStore.routeMap[index]?.parameters &&
      routerStore.routeMap[index]?.parameters.forEach((item) => {
        if (item.type === 'query') {
          query[item.key] = item.value
        } else {
          params[item.key] = item.value
        }
      })
    if (index === route.name) return
    if (index.indexOf('http://') > -1 || index.indexOf('https://') > -1) {
        window.open(index, '_blank')
        return
    }
    if (index === 'gvaLayoutIframe') {
      query.url = decodeURIComponent(index)
    }
    router.push({ name: index, query, params })
  }
</script>

<style lang="scss">
  .el-menu--horizontal.el-menu,
  .el-menu--horizontal > .el-menu-item.is-active {
    border-bottom: none !important;
  }

  .el-menu--horizontal>.el-sub-menu.is-active .el-sub-menu__title {
    border-bottom: none !important;
  }

  .el-menu-item.is-active {
    background-color: var(--el-color-primary-light-8) !important;
  }

  .dark {
    .el-menu-item.is-active {
      background-color: var(--el-color-primary-bg) !important;
    }
  }
</style>
