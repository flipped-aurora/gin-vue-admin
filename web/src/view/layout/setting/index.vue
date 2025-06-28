<template>
  <el-drawer
    v-model="drawer"
    title="系统配置"
    direction="rtl"
    :size="width"
    :show-close="false"
    class="theme-config-drawer"
  >
    <template #header>
      <div class="flex items-center justify-between w-full px-6 py-4 bg-white dark:bg-gray-900 border-b border-gray-200 dark:border-gray-700">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white font-inter">系统配置</h2>
        <el-button
          type="primary"
          size="small"
          class="reset-btn"
          :style="{ backgroundColor: config.primaryColor, borderColor: config.primaryColor }"
          @click="resetConfig"
        >
          重置配置
        </el-button>
      </div>
    </template>

    <div class="h-full bg-white dark:bg-gray-900">
      <div class="px-8 pt-4 pb-6 border-b border-gray-200 dark:border-gray-700">
        <div class="flex justify-center">
          <div class="inline-flex bg-gray-100 dark:bg-gray-800 rounded-xl p-1.5 border border-gray-200 dark:border-gray-700 shadow-sm">
            <button
              v-for="tab in tabs"
              :key="tab.key"
              class="px-6 py-3 text-base font-medium rounded-lg transition-all duration-150 ease-in-out min-w-[80px]"
              :class="[
                activeTab === tab.key
                  ? 'text-white shadow-md transform -translate-y-0.5'
                  : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200 hover:bg-gray-50 dark:hover:bg-gray-700'
              ]"
              :style="activeTab === tab.key ? { backgroundColor: config.primaryColor } : {}"
              @click="activeTab = tab.key"
            >
              {{ tab.label }}
            </button>
          </div>
        </div>
      </div>

      <div class="pb-8 h-full overflow-y-auto">
        <div class="transition-all duration-300 ease-in-out">
          <AppearanceSettings v-if="activeTab === 'appearance'" />
          <LayoutSettings v-else-if="activeTab === 'layout'" />
          <GeneralSettings v-else-if="activeTab === 'general'" />
        </div>
      </div>
    </div>
  </el-drawer>
</template>

<script setup>
  import { ref, computed, watch } from 'vue'
  import { storeToRefs } from 'pinia'
  import { ElMessage } from 'element-plus'
  import { useAppStore } from '@/pinia'
  import { setSelfSetting } from '@/api/user'
  import AppearanceSettings from './modules/appearance/index.vue'
  import LayoutSettings from './modules/layout/index.vue'
  import GeneralSettings from './modules/general/index.vue'

  defineOptions({
    name: 'GvaSetting'
  })

  const appStore = useAppStore()
  const { config, device } = storeToRefs(appStore)

  const activeTab = ref('appearance')

  const tabs = [
    { key: 'appearance', label: '外观' },
    { key: 'layout', label: '布局' },
    { key: 'general', label: '通用' }
  ]

  const width = computed(() => {
    return device.value === 'mobile' ? '100%' : '500px'
  })

  const drawer = defineModel('drawer', {
    default: true,
    type: Boolean
  })

  const saveConfig = async () => {
    const res = await setSelfSetting(config.value)
    if (res.code === 0) {
      localStorage.setItem('originSetting', JSON.stringify(config.value))
      ElMessage.success('保存成功')
    }
  }

  const resetConfig = () => {
    appStore.resetConfig()
  }

  watch(config, async () => {
    await saveConfig();
  }, { deep: true });
</script>

<style lang="scss" scoped>
.theme-config-drawer {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;

  ::v-deep(.el-drawer) {
    background: white;
  }

  ::v-deep(.el-drawer__header) {
    padding: 0;
    border: 0;
  }

  ::v-deep(.el-drawer__body) {
    padding: 0;
  }
}

.dark .theme-config-drawer {
  ::v-deep(.el-drawer) {
    background: #111827;
  }
}

.font-inter {
  font-family: 'Inter', sans-serif;
}

.reset-btn {
  border-radius: 0.5rem;
  font-weight: 500;
  transition: all 150ms ease-in-out;

  &:hover {
    transform: translateY(-2px);
    filter: brightness(0.9);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  }
}

/* Custom scrollbar for webkit browsers */
::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: #f3f4f6;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 3px;

  &:hover {
    background: #9ca3af;
  }
}

.dark ::-webkit-scrollbar-track {
  background: #1f2937;
}

.dark ::-webkit-scrollbar-thumb {
  background: #4b5563;

  &:hover {
    background: #6b7280;
  }
}
</style>
