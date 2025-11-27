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
          <div class="space-y-5">
            <div
              class="gva-theme-card-white flex items-center justify-between">
              <div class="flex items-center gap-4">
                <div
                  class="w-12 h-12 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-xl flex items-center justify-center text-red-600 dark:text-red-400 text-xl">
                  🔄
                </div>
                <div>
                  <h4 class="text-sm font-semibold gva-theme-text-main">重置配置</h4>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">将所有设置恢复为默认值</p>
                </div>
              </div>
              <el-button type="danger" size="small"
                class="rounded-lg font-medium transition-all duration-150 ease-in-out hover:-translate-y-0.5"
                @click="handleResetConfig">
                重置配置
              </el-button>
            </div>

            <div
              class="gva-theme-card-white flex items-center justify-between">
              <div class="flex items-center gap-4">
                <div
                  class="w-12 h-12 bg-blue-50 dark:bg-blue-900/20 border border-blue-200 dark:border-blue-800 rounded-xl flex items-center justify-center text-blue-600 dark:text-blue-400 text-xl">
                  📤
                </div>
                <div>
                  <h4 class="text-sm font-semibold gva-theme-text-main">导出配置</h4>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">导出当前配置为 JSON 文件</p>
                </div>
              </div>
              <el-button type="primary" size="small"
                class="rounded-lg font-medium transition-all duration-150 ease-in-out hover:-translate-y-0.5"
                :style="{ backgroundColor: config.primaryColor, borderColor: config.primaryColor }"
                @click="handleExportConfig">
                导出配置
              </el-button>
            </div>

            <div
              class="gva-theme-card-white flex items-center justify-between">
              <div class="flex items-center gap-4">
                <div
                  class="w-12 h-12 bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-800 rounded-xl flex items-center justify-center text-green-600 dark:text-green-400 text-xl">
                  📥
                </div>
                <div>
                  <h4 class="text-sm font-semibold gva-theme-text-main">导入配置</h4>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">从 JSON 文件导入配置</p>
                </div>
              </div>
              <el-upload ref="uploadRef" :auto-upload="false" :show-file-list="false" accept=".json"
                @change="handleImportConfig">
                <el-button type="success" size="small"
                  class="rounded-lg font-medium transition-all duration-150 ease-in-out hover:-translate-y-0.5">
                  导入配置
                </el-button>
              </el-upload>
            </div>
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
                  :style="{ color: config.primaryColor }">
                  GitHub 仓库
                </a>
                <span class="text-gray-400 dark:text-gray-500">·</span>
                <a href="https://www.gin-vue-admin.com/" target="_blank"
                  class="font-medium transition-colors duration-150 hover:underline"
                  :style="{ color: config.primaryColor }">
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
import { useAppStore } from '@/pinia'
import Logo from '@/components/logo/index.vue'

defineOptions({
  name: 'GeneralSettings'
})

const appStore = useAppStore()
const { config } = storeToRefs(appStore)
const uploadRef = ref()

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

    appStore.resetConfig()
    ElMessage.success('配置已重置')
  } catch {
    // User cancelled
  }
}

const handleExportConfig = () => {
  const configData = JSON.stringify(config.value, null, 2)
  const blob = new Blob([configData], { type: 'application/json' })
  const url = URL.createObjectURL(blob)

  const link = document.createElement('a')
  link.href = url
  link.download = `gin-vue-admin-config-${new Date().toISOString().split('T')[0]}.json`
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  URL.revokeObjectURL(url)

  ElMessage.success('配置已导出')
}

const handleImportConfig = (file) => {
  const reader = new FileReader()
  reader.onload = (e) => {
    try {
      const importedConfig = JSON.parse(e.target.result)

      Object.keys(importedConfig).forEach(key => {
        if (key in config.value) {
          config.value[key] = importedConfig[key]
        }
      })

      ElMessage.success('配置已导入')
    } catch (error) {
      ElMessage.error('配置文件格式错误')
    }
  }
  reader.readAsText(file.raw)
}
</script>


