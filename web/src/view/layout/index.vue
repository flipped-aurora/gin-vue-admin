<template>
  <div
    class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-800 w-screen h-screen"
  >
    <el-watermark
      v-if="config.show_watermark"
      :font="font"
      :z-index="9999"
      :gap="[180, 150]"
      class="absolute inset-0 pointer-events-none"
      :content="userStore.userInfo.nickName"
    />
    <gva-header />
    <div class="flex flex-row w-full gva-container pt-16 box-border h-full">
      <gva-aside
        v-if="
          config.side_mode === 'normal' ||
          (device === 'mobile' && config.side_mode == 'head') ||
          (device === 'mobile' && config.side_mode == 'combination')
        "
      />
      <gva-aside
        v-if="config.side_mode === 'combination' && device !== 'mobile'"
        mode="normal"
      />
      <div class="flex-1 p-2 w-0 h-full">
        <gva-tabs v-if="config.showTabs" />
        <div
          class="overflow-auto"
          :class="config.showTabs ? 'gva-container2' : 'gva-container pt-1'"
        >
          <router-view v-if="reloadFlag" v-slot="{ Component, route }">
            <div
              id="gva-base-load-dom"
              class="gva-body-h bg-gray-50 dark:bg-slate-800"
            >
              <transition mode="out-in" name="el-fade-in-linear">
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
  import GvaHeader from '@/view/layout/header/index.vue'
  import useResponsive from '@/hooks/responsive'
  import GvaTabs from './tabs/index.vue'
  import BottomInfo from '@/components/bottomInfo/bottomInfo.vue'
  import { emitter } from '@/utils/bus.js'
  import { ref, onMounted, nextTick, reactive, watchEffect } from 'vue'
  import { useRouter, useRoute } from 'vue-router'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useUserStore } from '@/pinia/modules/user'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  const appStore = useAppStore()
  const { config, isDark, device } = storeToRefs(appStore)

  defineOptions({
    name: 'GvaLayout'
  })

  useResponsive(true)
  const font = reactive({
    color: 'rgba(0, 0, 0, .15)'
  })

  watchEffect(() => {
    font.color =
      isDark.value ? 'rgba(255,255,255, .15)' : 'rgba(0, 0, 0, .15)'
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
