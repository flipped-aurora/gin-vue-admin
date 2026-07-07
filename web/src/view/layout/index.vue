<template>
  <div
    class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-800 w-screen h-screen"
  >
    <el-watermark
      v-if="settings.watermark.visible"
      :font="font"
      :z-index="9999"
      :gap="[180, 150]"
      class="!absolute !inset-0 !pointer-events-none"
      :content="userStore.userInfo.nickName"
    />
    <gva-header />
    <div
      class="flex flex-row w-full gva-container pt-16 box-border !h-full"
      :style="
        effectiveMode === 'vertical' && !isMobile
          ? { paddingLeft: verticalSideWidth + 'px' }
          : {}
      "
    >
      <gva-aside
        v-if="
          !isMobile &&
          (effectiveMode === 'normal' ||
            effectiveMode === 'sidebar' ||
            effectiveMode === 'vertical')
        "
      />
      <gva-aside
        v-if="!isMobile && effectiveMode === 'combination'"
        mode="normal"
        class="z-10"
      />
      <mobile-menu-drawer v-if="isMobile" />
      <div class="flex-1 w-0 h-full">
        <gva-tabs v-if="settings.tab.visible" />
        <div
          class="overflow-auto px-2"
          :class="settings.tab.visible ? 'gva-container2' : 'gva-container pt-1'"
        >
          <router-view v-if="reloadFlag" v-slot="{ Component, route }">
            <div
              id="gva-base-load-dom"
              class="gva-body-h bg-gray-50 dark:bg-slate-800"
            >
              <transition
                mode="out-in"
                :name="route.meta.transitionType || settings.page.transition"
              >
                <keep-alive :include="routerStore.keepAliveRouters">
                  <component :is="Component" :key="route.fullPath" />
                </keep-alive>
              </transition>
            </div>
          </router-view>
          <BottomInfo />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
  import GvaAside from '@/view/layout/aside/index.vue'
  import MobileMenuDrawer from '@/view/layout/aside/MobileMenuDrawer.vue'
  import GvaHeader from '@/view/layout/header/index.vue'
  import { useLayoutMode } from '@/hooks/useLayoutMode'
  import { useSideWidth } from '@/hooks/useSideWidth'
  import useResponsive from '@/hooks/responsive'
  import GvaTabs from './tabs/index.vue'
  import BottomInfo from '@/components/bottomInfo/bottomInfo.vue'
  import { emitter } from '@/utils/bus.js'
  import { computed, ref, onMounted, nextTick, reactive, watchEffect } from 'vue'
  import { useRouter, useRoute } from 'vue-router'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useUserStore } from '@/pinia/modules/user'
  import { useAppStore, useThemeStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  import '@/style/transition.scss'
  const appStore = useAppStore()
  const themeStore = useThemeStore()
  const { device } = storeToRefs(appStore)
  const { settings, isDark } = storeToRefs(themeStore)
  const isMobile = computed(() => device.value === 'mobile')
  const { effectiveMode } = useLayoutMode()
  const { sideWidth: verticalSideWidth } = useSideWidth()

  defineOptions({
    name: 'GvaLayout'
  })

  useResponsive(true)
  const font = reactive({
    color: 'rgba(0, 0, 0, .15)'
  })

  watchEffect(() => {
    font.color = isDark.value ? 'rgba(255,255,255, .15)' : 'rgba(0, 0, 0, .15)'
  })

  const router = useRouter()
  const route = useRoute()
  const routerStore = useRouterStore()

  onMounted(() => {
    // 挂载一些通用的事件
    emitter.on('reload', reload)
    if (userStore.loadingInstance) {
      userStore.loadingInstance.close()
    }
  })

  const userStore = useUserStore()

  const reloadFlag = ref(true)
  let reloadTimer = null
  const reload = async () => {
    if (reloadTimer) {
      window.clearTimeout(reloadTimer)
    }
    reloadTimer = window.setTimeout(async () => {
      if (route.meta.keepAlive) {
        reloadFlag.value = false
        await nextTick()
        reloadFlag.value = true
      } else {
        const title = route.meta.title
        router.push({ name: 'Reload', params: { title } })
      }
    }, 400)
  }
</script>

<style lang="scss"></style>
