<template>
  <div class="gva-theme-font">
    <div class="mb-10">
      <div class="gva-theme-section-header">
        <div class="gva-theme-divider"></div>
        <span class="gva-theme-section-title">系统信息</span>
        <div class="gva-theme-divider"></div>
      </div>

      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg">
          <div class="grid grid-cols-2 gap-4 text-sm">
            <div class="flex justify-between items-center py-3 border-b border-gray-200 dark:border-gray-600">
              <span class="gva-theme-text-sub font-medium">版本</span>
              <span class="font-mono gva-theme-text-main font-semibold">v2.7.4</span>
            </div>
            <div class="flex justify-between items-center py-3 border-b border-gray-200 dark:border-gray-600">
              <span class="gva-theme-text-sub font-medium">前端框架</span>
              <span class="font-mono gva-theme-text-main font-semibold">Vue 3</span>
            </div>
            <div class="flex justify-between items-center py-3 border-b border-gray-200 dark:border-gray-600">
              <span class="gva-theme-text-sub font-medium">UI 组件库</span>
              <span class="font-mono gva-theme-text-main font-semibold">Element Plus</span>
            </div>
            <div class="flex justify-between items-center py-3 border-b border-gray-200 dark:border-gray-600">
              <span class="gva-theme-text-sub font-medium">构建工具</span>
              <span class="font-mono gva-theme-text-main font-semibold">Vite</span>
            </div>
            <div class="flex justify-between items-center py-3">
              <span class="gva-theme-text-sub font-medium">浏览器</span>
              <span class="font-mono gva-theme-text-main font-semibold">{{ browserInfo }}</span>
            </div>
            <div class="flex justify-between items-center py-3">
              <span class="gva-theme-text-sub font-medium">屏幕分辨率</span>
              <span class="font-mono gva-theme-text-main font-semibold">{{ screenResolution }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mb-10">
      <div class="gva-theme-section-header">
        <div class="gva-theme-divider"></div>
        <span class="gva-theme-section-title">配置管理</span>
        <div class="gva-theme-divider"></div>
      </div>

      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg">
          <div
            class="gva-theme-card-white flex items-center justify-between">
            <div class="flex items-center gap-4">
              <div
                class="w-12 h-12 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-xl flex items-center justify-center text-red-600 dark:text-red-400 text-xl">
                🔄
              </div>
              <div>
                <h4 class="text-sm font-semibold gva-theme-text-main">重置配置</h4>
                <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">将所有设置恢复为默认值（导入/导出已迁移至「预设」）</p>
              </div>
            </div>
            <g-button
              variant="destructive"
              size="sm"
              class="rounded-lg font-medium hover:-translate-y-0.5"
              @click="handleResetConfig"
            >
              重置配置
            </g-button>
          </div>
        </div>
      </div>
    </div>

    <div class="mb-10">
      <div class="gva-theme-section-header">
        <div class="gva-theme-divider"></div>
        <span class="gva-theme-section-title">关于项目</span>
        <div class="gva-theme-divider"></div>
      </div>

      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg">
          <div class="flex items-start gap-5">
            <div
              class="w-16 h-16 bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-xl flex items-center justify-center flex-shrink-0 shadow-sm">
              <Logo />
            </div>
            <div class="flex-1">
              <h4 class="text-xl font-semibold gva-theme-text-main mb-3">Gin-Vue-Admin</h4>
              <p class="text-sm gva-theme-text-sub mb-5 leading-relaxed">
                基于 Vue3 + Gin 的全栈开发基础平台，提供完整的后台管理解决方案
              </p>
              <div class="flex items-center gap-3 text-sm">
                <a href="https://github.com/flipped-aurora/gin-vue-admin" target="_blank"
                  class="font-medium transition-colors duration-150 hover:underline"
                  :style="{ color: settings.themeColor }">
                  GitHub 仓库
                </a>
                <span class="text-gray-400 dark:text-gray-500">·</span>
                <a href="https://www.gin-vue-admin.com/" target="_blank"
                  class="font-medium transition-colors duration-150 hover:underline"
                  :style="{ color: settings.themeColor }">
                  官方文档
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { storeToRefs } from 'pinia'
import { useThemeStore } from '@/pinia'
import Logo from '@/components/logo/index.vue'

defineOptions({
  name: 'GeneralSettings'
})

const themeStore = useThemeStore()
const { settings } = storeToRefs(themeStore)

const browserInfo = ref('')
const screenResolution = ref('')

onMounted(() => {
  const userAgent = navigator.userAgent
  if (userAgent.includes('Chrome')) {
    browserInfo.value = 'Chrome'
  } else if (userAgent.includes('Firefox')) {
    browserInfo.value = 'Firefox'
  } else if (userAgent.includes('Safari')) {
    browserInfo.value = 'Safari'
  } else if (userAgent.includes('Edge')) {
    browserInfo.value = 'Edge'
  } else {
    browserInfo.value = 'Unknown'
  }

  screenResolution.value = `${screen.width}×${screen.height}`
})

const handleResetConfig = async () => {
  try {
    await ElMessageBox.confirm(
      '确定要重置所有配置吗？此操作不可撤销。',
      '重置配置',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )

    themeStore.resetConfig()
    ElMessage.success('配置已重置')
  } catch {
    // User cancelled
  }
}
</script>

