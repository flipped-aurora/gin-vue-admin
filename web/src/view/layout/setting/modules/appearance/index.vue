<template>
  <div class="gva-theme-font">
    <!-- 主题模式 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">主题模式</span>
      </div>
      <div class="gva-theme-section-content">
        <ThemeModeSelector v-model="config.darkMode" @update:modelValue="appStore.toggleDarkMode" />
      </div>
    </div>

    <!-- 主题色 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">主题色</span>
      </div>
      <div class="gva-theme-section-content">
        <ThemeColorPicker v-model="config.primaryColor" @update:modelValue="appStore.togglePrimaryColor" />
      </div>
    </div>

    <!-- 菜单风格 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">菜单风格</span>
      </div>
      <div class="gva-theme-section-content">
        <MenuThemeSelector v-model="config.menu_theme" @update:modelValue="appStore.toggleMenuTheme" />
      </div>
    </div>

    <!-- 外观细节：圆角 / 卡片 / 语义色 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">外观细节</span>
      </div>
      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg">
          <SettingItem label="全局圆角">
            <RadiusSelector v-model="config.radius" @update:modelValue="appStore.toggleRadius" />
          </SettingItem>
          <SettingItem label="卡片样式">
            <CardModeSelector v-model="config.card_mode" @update:modelValue="appStore.toggleCardMode" />
          </SettingItem>
          <SemanticColorPicker />
        </div>
      </div>
    </div>

    <!-- 偏好：尺寸 / 视觉辅助 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">偏好</span>
      </div>
      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg">
          <SettingItem label="全局尺寸">
            <el-select v-model="config.global_size" class="min-w-24" size="small" @change="appStore.toggleGlobalSize">
              <el-option label="默认" value="default" />
              <el-option label="大" value="large" />
              <el-option label="小" value="small" />
            </el-select>
          </SettingItem>
          <SettingItem label="灰色模式">
            <el-switch v-model="config.grey" @change="appStore.toggleGrey" />
          </SettingItem>
          <SettingItem label="色弱模式">
            <el-switch v-model="config.weakness" @change="appStore.toggleWeakness" />
          </SettingItem>
          <SettingItem label="显示水印">
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
import MenuThemeSelector from '../../components/menuThemeSelector.vue'
import SemanticColorPicker from '../../components/semanticColorPicker.vue'
import RadiusSelector from '../../components/radiusSelector.vue'
import CardModeSelector from '../../components/cardModeSelector.vue'
import SettingItem from '../../components/settingItem.vue'

defineOptions({
  name: 'AppearanceSettings'
})

const appStore = useAppStore()
const { config } = storeToRefs(appStore)
</script>
