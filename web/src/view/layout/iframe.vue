<template>
  <div
    class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-800 w-screen h-screen"
  >
    <iframe
        v-if="reloadFlag"
        id="gva-base-load-dom"
        class="gva-body-h bg-gray-50 dark:bg-slate-800 w-full border-t border-gray-200 dark:border-slate-700"
        :src="url"
    ></iframe>
  </div>
</template>

<script setup>
  import useResponsive from '@/hooks/responsive'
  import { emitter } from '@/utils/bus.js'
  import { ref, onMounted, nextTick, reactive, watchEffect } from 'vue'
  import { useRouter, useRoute } from 'vue-router'
  import { useUserStore } from '@/pinia/modules/user'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  const appStore = useAppStore()
  const { isDark } = storeToRefs(appStore)


  defineOptions({
    name: 'GvaLayoutIframe'
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

  const url = route.query.url || 'https://www.gin-vue-admin.com'

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
