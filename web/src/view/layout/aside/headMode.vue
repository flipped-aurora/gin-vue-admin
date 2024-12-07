<template>
  <div
    class="bg-white h-[calc(100%-4px)] text-slate-700 dark:text-slate-300 mx-2 dark:bg-slate-900 flex items-center w-[calc(100vw-600px)] overflow-auto"
  >
    <el-menu
      :default-active="active"
      mode="horizontal"
      class="border-r-0 w-full flex gap-1 items-center box-border h-[calc(100%-1px)]"
      unique-opened
      @select="selectMenuItem"
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
  import { ref, provide, watchEffect } from 'vue'
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
  watchEffect(() => {
    active.value = route.meta.activeName || route.name
  })

  watchEffect(() => {
    if (device.value === 'mobile') {
      isCollapse.value = true
    } else {
      isCollapse.value = false
    }
  })

  provide('isCollapse', isCollapse)

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
      window.open(index)
    } else {
      router.push({ name: index, query, params })
    }
  }
</script>

<style lang="scss" scoped>
  .el-menu--horizontal.el-menu,
  .el-menu--horizontal > .el-menu-item.is-active {
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
