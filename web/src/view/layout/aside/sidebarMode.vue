<template>
  <div class="flex h-full">
    <!-- 一级菜单常驻侧边栏 -->
    <div
      class="relative h-full bg-white text-slate-700 dark:text-slate-300 dark:bg-slate-900 border-r shadow dark:shadow-gray-700"
      :style="{
        width: config.layout_side_collapsed_width + 'px'
      }"
    >
      <el-scrollbar>
        <el-menu
          :collapse="true"
          :collapse-transition="false"
          :default-active="topActive"
          class="border-r-0 w-full"
          unique-opened
          @select="selectTopMenuItem"
        >
          <template v-for="item in routerStore.asyncRouters[0]?.children || []">
            <el-menu-item 
              v-if="!item.hidden && (!item.children || item.children.length === 0)" 
              :key="item.name"
              :index="item.name"
              class="dark:text-slate-300 overflow-hidden"
              :style="{
                height: config.layout_side_item_height + 'px'
              }"
            >
              <el-icon v-if="item.meta.icon">
                <component :is="item.meta.icon" />
              </el-icon>
              <template v-else>
                {{ item.meta.title[0] }}
                </template>
              <template #title>
                {{ item.meta.title }}
              </template>
            </el-menu-item>
            <template v-else-if="!item.hidden" >
             <el-menu-item 
              :key="item.name"
              :index="item.name"
              :class="{'is-active': topActive === item.name}"
              class="dark:text-slate-300 overflow-hidden"
              :style="{
                height: config.layout_side_item_height + 'px'
              }"
            >
              <el-icon v-if="item.meta.icon">
                <component :is="item.meta.icon" />
              </el-icon>
              <template v-else>
                {{ item.meta.title[0] }}
                </template>
              <template #title>
                {{ item.meta.title }}
              </template>
            </el-menu-item>
            </template>
            
          </template>
        </el-menu>
      </el-scrollbar>
    </div>

    <!-- 二级菜单并列显示 -->
    <div
      class="relative h-full bg-white text-slate-700 dark:text-slate-300 dark:bg-slate-900 border-r shadow dark:shadow-gray-700 px-2"
      :style="{
        width: layoutSideWidth + 'px'
      }"
    >
      <el-scrollbar>
        <el-menu
          :collapse="isCollapse"
          :collapse-transition="false"
          :default-active="active"
          class="border-r-0 w-full"
          unique-opened
          @select="selectMenuItem"
        >
          <template v-for="item in secondLevelMenus">
            <aside-component
              v-if="!item.hidden"
              :key="item.name"
              :router-info="item"
            />
          </template>
        </el-menu>
      </el-scrollbar>
      <div
        class="absolute bottom-8 right-2 w-8 h-8 bg-gray-50 dark:bg-slate-800 flex items-center justify-center rounded cursor-pointer"
        :class="isCollapse ? 'right-0 left-0 mx-auto' : 'right-2'"
        @click="toggleCollapse"
      >
        <el-icon v-if="!isCollapse">
          <DArrowLeft />
        </el-icon>
        <el-icon v-else>
          <DArrowRight />
        </el-icon>
      </div>
    </div>
  </div>
</template>

<script setup>
  import AsideComponent from '@/view/layout/aside/asideComponent/index.vue'
  import { ref, provide, watchEffect, computed } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  
  const appStore = useAppStore()
  const { device, config } = storeToRefs(appStore)

  defineOptions({
    name: 'SidebarMode'
  })
  
  const route = useRoute()
  const router = useRouter()
  const routerStore = useRouterStore()
  const isCollapse = ref(false)
  const active = ref('')
  const topActive = ref('')
  const secondLevelMenus = ref([])
  
  const layoutSideWidth = computed(() => {
    if (!isCollapse.value) {
      return config.value.layout_side_width
    } else {
      return config.value.layout_side_collapsed_width
    }
  })


  provide('isCollapse', isCollapse)
  
  // 更新二级菜单
  const updateSecondLevelMenus = (menuName) => {
    const menu = routerStore.asyncRouters[0]?.children.find(item => item.name === menuName)
    if (menu && menu.children && menu.children.length > 0) {
      secondLevelMenus.value = menu.children
    } else {
      secondLevelMenus.value = []
    }
  }
  
  // 选择一级菜单
  const selectTopMenuItem = (index) => {
    topActive.value = index
    updateSecondLevelMenus(index)
    
    // 如果选中的一级菜单有子菜单，则导航到第一个子菜单
    const menu = routerStore.asyncRouters[0]?.children.find(item => item.name === index)
    if (menu && menu.children && menu.children.length > 0) {
      const firstVisibleChild = menu.children.find(child => !child.hidden)
      if (firstVisibleChild) {
        navigateToMenuItem(firstVisibleChild.name)
      }
    } else {
      // 如果没有子菜单，直接导航到该菜单
      navigateToMenuItem(index)
    }
  }
  
  // 选择二级或更深层级的菜单
  const selectMenuItem = (index) => {
    navigateToMenuItem(index)
  }
  
  // 导航到指定菜单
  const navigateToMenuItem = (index) => {
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
      if (index === 'Iframe') {
        query.url = decodeURIComponent(index)
        router.push({
          name: 'Iframe',
          query,
          params
        })
        return
      } else {
        window.open(index, '_blank')
        return
      }
    } else {
      router.push({ name: index, query, params })
    }
  }

  const toggleCollapse = () => {
    isCollapse.value = !isCollapse.value
  }


    
  watchEffect(() => {
    if (route.name === 'Iframe') {
      active.value = decodeURIComponent(route.query.url)
      return
    }
    active.value = route.meta.activeName || route.name
    
    // 找到当前路由所属的一级菜单
    const findParentMenu = () => {
      for (const topMenu of routerStore.asyncRouters[0]?.children || []) {
        if (topMenu.hidden) continue
        
        // 检查当前路由是否是这个一级菜单的子菜单
        if (topMenu.children && topMenu.children.some(child => child.name === route.name)) {
          return topMenu.name
        }
        
        // 递归检查更深层级
        const checkChildren = (items) => {
          for (const item of items || []) {
            if (item.name === route.name) {
              return true
            }
            if (item.children && checkChildren(item.children)) {
              return true
            }
          }
          return false
        }
        
        if (topMenu.children && checkChildren(topMenu.children)) {
          return topMenu.name
        }
      }
      return null
    }
    
    const parentMenu = findParentMenu()
    if (parentMenu) {
      topActive.value = parentMenu
      updateSecondLevelMenus(parentMenu)
    } else if (routerStore.asyncRouters[0]?.children?.length > 0) {
      // 如果没有找到父菜单，默认选择第一个有子菜单的一级菜单
      const firstMenuWithChildren = routerStore.asyncRouters[0].children.find(
        item => !item.hidden && item.children && item.children.length > 0
      )
      if (firstMenuWithChildren) {
        topActive.value = firstMenuWithChildren.name
        updateSecondLevelMenus(firstMenuWithChildren.name)
      }
    }
  })

  watchEffect(() => {
    if (device.value === 'mobile') {
      isCollapse.value = true
    } else {
      isCollapse.value = false
    }
  })
</script>

<style lang="scss" scoped>
  :deep(.el-menu--vertical) {
    .is-active {
      background-color: var(--el-color-primary-light-8) !important;
    }
  }
</style>