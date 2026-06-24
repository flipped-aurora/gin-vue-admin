<template>
  <el-drawer
    v-model="drawer"
    title="系统配置"
    direction="rtl"
    :size="width"
    :show-close="false"
    class="gva-theme-drawer"
  >
    <template #header>
      <div class="flex items-center justify-between w-full px-6 py-4 bg-white dark:bg-gray-900 border-b border-gray-200 dark:border-gray-700">
        <h2 class="text-xl font-semibold gva-theme-text-main gva-theme-font">系统配置</h2>
        <div class="flex items-center gap-2">
          <el-button
            type="primary"
            size="small"
            class="reset-btn"
            :style="{ backgroundColor: settings.themeColor, borderColor: settings.themeColor }"
            @click="resetConfig"
          >
            重置配置
          </el-button>
          <el-button
            size="small"
            circle
            :icon="Close"
            aria-label="关闭系统配置"
            @click="closeDrawer"
          />
        </div>
      </div>
    </template>

    <div class="bg-white dark:bg-gray-900 px-5">
      <div class="pt-3 pb-4">
        <div class="flex justify-center">
          <div class="inline-flex bg-gray-100 dark:bg-gray-800 rounded-lg p-1 gap-0.5">
            <div
              v-for="tab in tabs"
              :key="tab.key"
              class="px-4 py-1.5 text-[13px] text-center cursor-pointer font-medium rounded-md transition-colors duration-150 ease-in-out min-w-[60px]"
              :class="[
                activeTab === tab.key
                  ? 'text-white shadow-sm'
                  : 'text-gray-500 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200'
              ]"
              :style="activeTab === tab.key ? { backgroundColor: settings.themeColor } : {}"
              @click="activeTab = tab.key"
            >
              {{ tab.label }}
            </div>
          </div>
        </div>
      </div>

      <div class="pb-8 h-full overflow-y-auto">
        <div class="transition-all duration-300 ease-in-out">
          <AppearanceSettings v-if="activeTab === 'appearance'" />
          <LayoutSettings v-else-if="activeTab === 'layout'" />
          <PresetSettings v-else-if="activeTab === 'presets'" />
          <GeneralSettings v-else-if="activeTab === 'general'" />
        </div>
      </div>
    </div>
  </el-drawer>
</template>

<script setup>
  import { ref, computed } from 'vue'
  import { storeToRefs } from 'pinia'
  import { Close } from '@element-plus/icons-vue'
  import { useAppStore, useThemeStore } from '@/pinia'
  import AppearanceSettings from './modules/appearance/index.vue'
  import LayoutSettings from './modules/layout/index.vue'
  import PresetSettings from './modules/presets/index.vue'
  import GeneralSettings from './modules/general/index.vue'

  defineOptions({
    name: 'GvaSetting'
  })

  const appStore = useAppStore()
  const themeStore = useThemeStore()
  const { device } = storeToRefs(appStore)
  const { settings } = storeToRefs(themeStore)

  const activeTab = ref('appearance')

  const tabs = [
    { key: 'appearance', label: '外观' },
    { key: 'layout', label: '布局' },
    { key: 'presets', label: '预设' },
    { key: 'general', label: '通用' }
  ]

  const width = computed(() => {
    return device.value === 'mobile' ? '100%' : '440px'
  })

  const drawer = defineModel('drawer', {
    default: true,
    type: Boolean
  })

  const resetConfig = () => {
    themeStore.resetConfig()
  }

  const closeDrawer = () => {
    drawer.value = false
  }

  // 主题的本地缓存与远端保存已下沉到 themeStore，抽屉只负责触发交互，无需再监听持久化
</script>

<style lang="scss">
.gva-theme-drawer {
  .el-drawer {
    @apply bg-white dark:bg-gray-900;
  }

  .el-drawer__header {
    @apply p-0 border-0;
  }

  .el-drawer__body {
    @apply p-0;
  }
}

.gva-theme-font {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}

.gva-theme-card-bg {
  @apply bg-gray-50/70 dark:bg-gray-800/40 border border-gray-100 dark:border-gray-700/60 rounded-lg px-3.5 py-0.5;
}

.gva-theme-card-white {
  @apply bg-white dark:bg-gray-700/60 border border-gray-100 dark:border-gray-600/60 rounded-lg px-3 py-2.5 transition-colors duration-150;
}

.gva-theme-section-header {
  @apply flex items-center mb-2 mt-0.5;
}

.gva-theme-section-title {
  @apply text-xs font-semibold tracking-wide text-gray-400 dark:text-gray-500;
}

.gva-theme-divider {
  @apply hidden;
}

.gva-theme-text-main {
  @apply text-gray-900 dark:text-white;
}

.gva-theme-text-sub {
  @apply text-gray-600 dark:text-gray-400;
}

.gva-theme-section-content {
  animation: fadeInUp 0.3s ease;
}

.gva-theme-setting-item {
  @apply flex items-center justify-between py-2.5 gva-theme-font border-0 border-b border-solid border-gray-100/80 dark:border-gray-700/50 last:border-b-0;
}

.gva-theme-setting-label {
  @apply text-[13px] font-medium gva-theme-text-main;
}

.gva-theme-mode-selector {
  @apply inline-flex bg-gray-100 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-lg p-1 gap-1;
}

.gva-theme-mode-item {
  @apply flex flex-col items-center justify-center px-3 py-2 rounded-md cursor-pointer transition-all duration-150 ease-in-out min-w-[64px];
}

.gva-theme-layout-card {
  @apply bg-white dark:bg-gray-700/50 border border-gray-200 dark:border-gray-600 rounded-lg p-2 cursor-pointer transition-colors duration-150 ease-in-out hover:border-gray-300 dark:hover:border-gray-500;
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

/* Custom scrollbar for webkit browsers */
.gva-theme-drawer ::-webkit-scrollbar {
  width: 6px;
}

.gva-theme-drawer ::-webkit-scrollbar-track {
  background: #f3f4f6;
  border-radius: 3px;
}

.gva-theme-drawer ::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 3px;

  &:hover {
    background: #9ca3af;
  }
}

.dark .gva-theme-drawer ::-webkit-scrollbar-track {
  background: #1f2937;
}

.dark .gva-theme-drawer ::-webkit-scrollbar-thumb {
  background: #4b5563;

  &:hover {
    background: #6b7280;
  }
}
</style>

<style lang="scss" scoped>
.reset-btn {
  @apply rounded-lg font-medium transition-all duration-150 ease-in-out hover:-translate-y-0.5 hover:brightness-90 hover:shadow-lg;
}
</style>
