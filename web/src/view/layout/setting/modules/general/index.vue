<template>
  <div class="font-inter">
    <div class="mb-10">
      <div class="flex items-center justify-center mb-6">
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
        <span class="px-6 text-lg font-semibold text-gray-700 dark:text-gray-300">{{ t('layout.setting.modules.general.systemInfo') }}</span>
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
      </div>

      <div class="section-content">
        <div class="bg-gray-50 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl p-6 shadow-sm">
          <div class="grid grid-cols-2 gap-4 text-sm">
            <div class="flex justify-between items-center py-3 border-b border-gray-200 dark:border-gray-600">
              <span class="text-gray-600 dark:text-gray-400 font-medium">{{ t('layout.setting.modules.general.version') }}</span>
              <span class="font-mono text-gray-900 dark:text-white font-semibold">v2.7.4</span>
            </div>
            <div class="flex justify-between items-center py-3 border-b border-gray-200 dark:border-gray-600">
              <span class="text-gray-600 dark:text-gray-400 font-medium">{{ t('layout.setting.modules.general.frontendFramework') }}</span>
              <span class="font-mono text-gray-900 dark:text-white font-semibold">Vue 3</span>
            </div>
            <div class="flex justify-between items-center py-3 border-b border-gray-200 dark:border-gray-600">
              <span class="text-gray-600 dark:text-gray-400 font-medium">{{ t('layout.setting.modules.general.uiComponentLib') }}</span>
              <span class="font-mono text-gray-900 dark:text-white font-semibold">Element Plus</span>
            </div>
            <div class="flex justify-between items-center py-3 border-b border-gray-200 dark:border-gray-600">
              <span class="text-gray-600 dark:text-gray-400 font-medium">{{ t('layout.setting.modules.general.buildTools') }}</span>
              <span class="font-mono text-gray-900 dark:text-white font-semibold">Vite</span>
            </div>
            <div class="flex justify-between items-center py-3">
              <span class="text-gray-600 dark:text-gray-400 font-medium">{{ t('layout.setting.modules.general.browser') }}</span>
              <span class="font-mono text-gray-900 dark:text-white font-semibold">{{ browserInfo }}</span>
            </div>
            <div class="flex justify-between items-center py-3">
              <span class="text-gray-600 dark:text-gray-400 font-medium">{{ t('layout.setting.modules.general.screenResolution') }}</span>
              <span class="font-mono text-gray-900 dark:text-white font-semibold">{{ screenResolution }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mb-10">
      <div class="flex items-center justify-center mb-6">
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
        <span class="px-6 text-lg font-semibold text-gray-700 dark:text-gray-300">{{ t('layout.setting.modules.general.configManagement') }}</span>
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
      </div>

      <div class="section-content">
        <div class="bg-gray-50 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl p-6 shadow-sm">
          <div class="space-y-5">
            <div class="bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-lg p-5 flex items-center justify-between hover:shadow-md transition-all duration-150 ease-in-out hover:-translate-y-0.5">
              <div class="flex items-center gap-4">
                <div class="w-12 h-12 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-xl flex items-center justify-center text-red-600 dark:text-red-400 text-xl">
                  🔄
                </div>
                <div>
                  <h4 class="text-sm font-semibold text-gray-900 dark:text-white">{{ t('layout.setting.resetConfig') }}</h4>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">{{ t('layout.setting.modules.general.resetConfigNote') }}</p>
                </div>
              </div>
              <el-button
                type="danger"
                size="small"
                class="rounded-lg font-medium transition-all duration-150 ease-in-out hover:-translate-y-0.5"
                @click="handleResetConfig"
              >
                {{ t('layout.setting.resetConfig') }}
              </el-button>
            </div>

            <div class="bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-lg p-5 flex items-center justify-between hover:shadow-md transition-all duration-150 ease-in-out hover:-translate-y-0.5">
              <div class="flex items-center gap-4">
                <div class="w-12 h-12 bg-blue-50 dark:bg-blue-900/20 border border-blue-200 dark:border-blue-800 rounded-xl flex items-center justify-center text-blue-600 dark:text-blue-400 text-xl">
                  📤
                </div>
                <div>
                  <h4 class="text-sm font-semibold text-gray-900 dark:text-white">{{ t('layout.setting.modules.general.exportConfig') }}</h4>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">{{ t('layout.setting.modules.general.exportConfigNote') }}</p>
                </div>
              </div>
              <el-button
                type="primary"
                size="small"
                class="rounded-lg font-medium transition-all duration-150 ease-in-out hover:-translate-y-0.5"
                :style="{ backgroundColor: config.primaryColor, borderColor: config.primaryColor }"
                @click="handleExportConfig"
              >
                {{ t('layout.setting.modules.general.exportConfig') }}
              </el-button>
            </div>

            <div class="bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-lg p-5 flex items-center justify-between hover:shadow-md transition-all duration-150 ease-in-out hover:-translate-y-0.5">
              <div class="flex items-center gap-4">
                <div class="w-12 h-12 bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-800 rounded-xl flex items-center justify-center text-green-600 dark:text-green-400 text-xl">
                  📥
                </div>
                <div>
                  <h4 class="text-sm font-semibold text-gray-900 dark:text-white">{{ t('layout.setting.modules.general.importConfig') }}</h4>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">{{ t('layout.setting.modules.general.importConfigNote') }}</p>
                </div>
              </div>
              <el-upload
                ref="uploadRef"
                :auto-upload="false"
                :show-file-list="false"
                accept=".json"
                @change="handleImportConfig"
              >
                <el-button
                  type="success"
                  size="small"
                  class="rounded-lg font-medium transition-all duration-150 ease-in-out hover:-translate-y-0.5"
                >
                  {{ t('layout.setting.modules.general.importConfig') }}
                </el-button>
              </el-upload>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mb-10">
      <div class="flex items-center justify-center mb-6">
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
        <span class="px-6 text-lg font-semibold text-gray-700 dark:text-gray-300">{{ t('layout.setting.modules.general.about') }}</span>
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
      </div>

      <div class="section-content">
        <div class="bg-gray-50 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl p-6 shadow-sm">
          <div class="flex items-start gap-5">
            <div class="w-16 h-16 bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-xl flex items-center justify-center flex-shrink-0 shadow-sm">
              <img
                src="/logo.png"
                alt="Gin-Vue-Admin Logo"
                class="w-10 h-10 object-contain"
                @error="handleLogoError"
              />
            </div>
            <div class="flex-1">
              <h4 class="text-xl font-semibold text-gray-900 dark:text-white mb-3">Gin-Vue-Admin</h4>
              <p class="text-sm text-gray-600 dark:text-gray-400 mb-5 leading-relaxed">
                {{ t('layout.setting.modules.general.aboutNote') }}
              </p>
              <div class="flex items-center gap-3 text-sm">
                <a
                  href="https://github.com/flipped-aurora/gin-vue-admin"
                  target="_blank"
                  class="font-medium transition-colors duration-150 hover:underline"
                  :style="{ color: config.primaryColor }"
                >
                  {{ t('layout.setting.modules.general.githubRepo') }}
                </a>
                <span class="text-gray-400 dark:text-gray-500">·</span>
                <a
                  href="https://www.gin-vue-admin.com/"
                  target="_blank"
                  class="font-medium transition-colors duration-150 hover:underline"
                  :style="{ color: config.primaryColor }"
                >
                  {{ t('layout.setting.modules.general.officialDocs') }}
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
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { storeToRefs } from 'pinia'
import { useAppStore } from '@/pinia'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual

const { t } = useI18n() // added by mohamed hassan to support multilingual

defineOptions({
  name: 'GeneralSettings'
})

const appStore = useAppStore()
const { config } = storeToRefs(appStore)
const uploadRef = ref()

const browserInfo = ref('')
const screenResolution = ref('')
const logoUrl = ref('')

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

const handleLogoError = () => {
  logoUrl.value = ''
}

const handleResetConfig = async () => {
  try {
    await ElMessageBox.confirm(
      t('layout.setting.modules.general.resetAllConfigNote'),
      t('layout.setting.resetConfig'),
      {
        confirmButtonText: t('general.confirm'),
        cancelButtonText: t('general.cancel'),
        type: 'warning'
      }
    )
    
    appStore.resetConfig()
    ElMessage.success(t('layout.setting.modules.general.configReset'))
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
  
  ElMessage.success(t('layout.setting.modules.general.configExported'))
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
      
      ElMessage.success(t('layout.setting.modules.general.configImported'))
    } catch (error) {
      ElMessage.error(t('layout.setting.modules.general.configFileFormatError'))
    }
  }
  reader.readAsText(file.raw)
}
</script>

<style scoped>
.font-inter {
  font-family: 'Inter', sans-serif;
}

.section-content {
  animation: fadeInUp 0.3s ease;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
