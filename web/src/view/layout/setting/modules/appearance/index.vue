<template>
  <div class="gva-theme-font">
    <!-- Theme Mode Section -->
    <div class="mb-10">
      <div class="gva-theme-section-header">
        <div class="gva-theme-divider"></div>
        <span class="gva-theme-section-title">主题模式</span>
        <div class="gva-theme-divider"></div>
      </div>

      <div class="gva-theme-section-content">
        <ThemeModeSelector v-model="config.darkMode" @update:modelValue="appStore.toggleDarkMode" />
      </div>
    </div>

    <!-- Theme Color Section -->
    <div class="mb-10">
      <div class="gva-theme-section-header">
        <div class="gva-theme-divider"></div>
        <span class="gva-theme-section-title">主题颜色</span>
        <div class="gva-theme-divider"></div>
      </div>

      <div class="gva-theme-section-content">
        <ThemeColorPicker v-model="config.primaryColor" @update:modelValue="appStore.togglePrimaryColor" />
      </div>
    </div>

    <!-- Global Size Section -->
    <div class="mb-10">
      <div class="gva-theme-section-header">
        <div class="gva-theme-divider"></div>
        <span class="gva-theme-section-title">全局尺寸</span>
        <div class="gva-theme-divider"></div>
      </div>

      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg">
          <SettingItem label="全局尺寸">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">设置全局组件尺寸</span>
            </template>
            <div class="w-39">
              <el-select v-model="config.global_size" placeholder="请选择" @change="appStore.toggleGlobalSize">
                <el-option label="默认就好了" value="default" />
                <el-option label="大点好" value="large" />
                <el-option label="小的也不错" value="small" />
              </el-select>
            </div>
          </SettingItem>
        </div>
      </div>
    </div>

    <!-- Visual Accessibility Section -->
    <div class="mb-10">
      <div class="gva-theme-section-header">
        <div class="gva-theme-divider"></div>
        <span class="gva-theme-section-title">视觉辅助</span>
        <div class="gva-theme-divider"></div>
      </div>

      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg">
          <SettingItem label="灰色模式">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">降低色彩饱和度</span>
            </template>
            <el-switch v-model="config.grey" @change="appStore.toggleGrey" />
          </SettingItem>

          <SettingItem label="色弱模式">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">优化色彩对比度</span>
            </template>
            <el-switch v-model="config.weakness" @change="appStore.toggleWeakness" />
          </SettingItem>

          <SettingItem label="显示水印">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">在页面显示水印标识</span>
            </template>
            <el-switch v-model="config.show_watermark" @change="appStore.toggleConfigWatermark" />
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


