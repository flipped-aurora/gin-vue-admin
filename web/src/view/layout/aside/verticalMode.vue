<template>
  <div
    class="fixed top-0 left-0 h-screen z-30 flex flex-col shadow dark:shadow-gray-700"
    :style="{
      width: sideWidth + 'px',
      background: 'var(--gva-aside-bg)',
      color: 'var(--gva-aside-text)'
    }"
  >
    <!-- Logo -->
    <div
      class="h-16 flex items-center justify-center gap-2 cursor-pointer flex-shrink-0 overflow-hidden px-2"
      @click="router.push({ path: '/' })"
    >
      <Logo />
      <span
        v-if="!sideCollapse"
        class="font-bold text-xl whitespace-nowrap"
      >
        {{ $GIN_VUE_ADMIN.appName }}
      </span>
    </div>

    <!-- 菜单 -->
    <el-scrollbar class="flex-1">
      <el-menu
        :collapse="sideCollapse"
        :collapse-transition="false"
        :default-active="active"
        class="!border-r-0 w-full"
        :class="sideCollapse || settings.menu.theme === 'design' ? '' : '!px-2'"
        unique-opened
        @select="selectMenuItem"
      >
        <template v-for="item in routerStore.asyncRouters[0]?.children || []">
          <aside-component
            v-if="!item.hidden"
            :key="item.name"
            :router-info="item"
          />
        </template>
      </el-menu>
    </el-scrollbar>

    <!-- 折叠按钮 -->
    <div
      v-if="settings.header.collapseButton.visible"
      class="h-10 flex items-center justify-center cursor-pointer flex-shrink-0 border-0 border-t border-solid border-gray-100 dark:border-slate-800"
      @click="appStore.toggleSideCollapse()"
    >
      <el-icon v-if="!sideCollapse">
        <DArrowLeft />
      </el-icon>
      <el-icon v-else>
        <DArrowRight />
      </el-icon>
    </div>
  </div>
</template>

<script setup>
  import AsideComponent from '@/view/layout/aside/asideComponent/index.vue'
  import Logo from '@/components/logo/index.vue'
  import { computed, ref, provide, watchEffect } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useAppStore, useThemeStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  import { DArrowLeft, DArrowRight } from '@element-plus/icons-vue'

  defineOptions({
    name: 'VerticalMode'
  })

  const appStore = useAppStore()
  const themeStore = useThemeStore()
  const { sideCollapse } = storeToRefs(appStore)
  const { settings } = storeToRefs(themeStore)
  const route = useRoute()
  const router = useRouter()
  const routerStore = useRouterStore()
  const active = ref('')
  const sideWidth = computed(() =>
    sideCollapse.value
      ? settings.value.layout.sideCollapsedWidth
      : settings.value.layout.sideWidth
  )

  watchEffect(() => {
    if (route.name === 'gvaLayoutIframe') {
      active.value = decodeURIComponent(route.query.url)
      return
    }
    active.value = route.meta.activeName || route.name
  })

  // 供 asideComponent / menuItem 注入折叠态
  provide('isCollapse', sideCollapse)

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
    } else {
      router.push({ name: index, query, params })
    }
  }
</script>
