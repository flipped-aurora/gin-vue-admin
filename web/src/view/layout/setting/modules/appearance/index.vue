<template>
  <div class="font-inter">
    <!-- Theme Mode Section -->
    <div class="mb-10">
      <div class="flex items-center justify-center mb-6">
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
        <span class="px-6 text-lg font-semibold text-gray-700 dark:text-gray-300">主题模式</span>
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
      </div>

      <div class="section-content">
        <ThemeModeSelector
          v-model="config.darkMode"
          @update:modelValue="appStore.toggleDarkMode"
        />
      </div>
    </div>

    <!-- Theme Color Section -->
    <div class="mb-10">
      <div class="flex items-center justify-center mb-6">
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
        <span class="px-6 text-lg font-semibold text-gray-700 dark:text-gray-300">主题颜色</span>
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
      </div>

      <div class="section-content">
        <ThemeColorPicker
          v-model="config.primaryColor"
          @update:modelValue="appStore.togglePrimaryColor"
        />
      </div>
    </div>

    <!-- Visual Accessibility Section -->
    <div class="mb-10">
      <div class="flex items-center justify-center mb-6">
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
        <span class="px-6 text-lg font-semibold text-gray-700 dark:text-gray-300">视觉辅助</span>
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
      </div>

      <div class="section-content">
        <div class="bg-gray-50 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl p-6 shadow-sm">
          <SettingItem label="灰色模式">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">降低色彩饱和度</span>
            </template>
            <el-switch
              v-model="config.grey"
              @change="appStore.toggleGrey"
            />
          </SettingItem>

          <SettingItem label="色弱模式">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">优化色彩对比度</span>
            </template>
            <el-switch
              v-model="config.weakness"
              @change="appStore.toggleWeakness"
            />
          </SettingItem>

          <SettingItem label="显示水印">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">在页面显示水印标识</span>
            </template>
            <el-switch
              v-model="config.show_watermark"
              @change="appStore.toggleConfigWatermark"
            />
          </SettingItem>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia'
import { useAppStore } from '@/pinia'
import ThemeModeSelector from '../../components/themeModeSelector.vue'
import ThemeColorPicker from '../../components/themeColorPicker.vue'
import SettingItem from '../../components/settingItem.vue'

defineOptions({
  name: 'AppearanceSettings'
})

const appStore = useAppStore()
const { config } = storeToRefs(appStore)
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
