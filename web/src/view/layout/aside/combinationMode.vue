<template>
  <div class="h-full">
    <div
      v-if="mode === 'head'"
      class="bg-white h-[calc(100%-4px)] text-slate-700 dark:text-slate-300 mx-2 dark:bg-slate-900 flex items-center w-[calc(100vw-600px)] overflow-auto"
    >
      <el-menu
        :default-active="routerStore.topActive"
        mode="horizontal"
        class="border-r-0 border-b-0 w-full flex gap-1 items-center box-border h-[calc(100%-1px)]"
        unique-opened
        @select="(index, _, ele) => selectMenuItem(index, _, ele, true)"
      >
        <template v-for="item in routerStore.topMenu">
          <aside-component
            v-if="!item.hidden"
            :key="item.name"
            :router-info="item"
            mode="horizontal"
          />
        </template>
      </el-menu>
    </div>
    <div
      v-if="mode === 'normal'"
      class="relative h-full bg-white text-slate-700 dark:text-slate-300 dark:bg-slate-900 border-r shadow dark:shadow-gray-700"
      :class="isCollapse ? '' : '  px-2'"
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
          @select="(index, _, ele) => selectMenuItem(index, _, ele, false)"
        >
          <template v-for="item in routerStore.leftMenu">
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
    name: 'GvaAside'
  })

  defineProps({
    mode: {
      type: String,
      default: 'normal'
    }
  })

  const route = useRoute()
  const router = useRouter()
  const routerStore = useRouterStore()
  const isCollapse = ref(false)
  const active = ref('')
  const layoutSideWidth = computed(() => {
    if (!isCollapse.value) {
      return config.value.layout_side_width
    } else {
      return config.value.layout_side_collapsed_width
    }
  })

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

  const selectMenuItem = (index, _, ele, top) => {
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
      if (!top) {
        router.push({ name: index, query, params })
        return
      }
      if (!routerStore.setLeftMenu(index)) {
        router.push({ name: index, query, params })
      }
    }
  }

  const toggleCollapse = () => {
    isCollapse.value = !isCollapse.value
  }
</script>
