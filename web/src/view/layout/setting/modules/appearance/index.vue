<template>
  <div class="gva-theme-font">
    <!-- 主题模式 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">主题模式</span>
      </div>
      <div class="gva-theme-section-content">
        <ThemeModeSelector
          :model-value="settings.themeScheme"
          @update:modelValue="themeStore.setThemeScheme"
        />
      </div>
    </div>

    <!-- 主题色 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">主题色</span>
      </div>
      <div class="gva-theme-section-content">
        <ThemeColorPicker
          :model-value="settings.themeColor"
          @update:modelValue="(value) => themeStore.updateThemeColors('primary', value)"
        />
      </div>
    </div>

    <!-- 菜单风格 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">菜单风格</span>
      </div>
      <div class="gva-theme-section-content">
        <MenuThemeSelector v-model="settings.menu.theme" />
      </div>
    </div>

    <!-- 标签栏风格 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">标签栏风格</span>
      </div>
      <div class="gva-theme-section-content">
        <TabModeSelector v-model="settings.tab.mode" />
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
            <RadiusSelector v-model="settings.themeRadius" />
          </SettingItem>
          <SettingItem label="卡片样式">
            <CardModeSelector v-model="settings.card.mode" />
          </SettingItem>
          <SettingItem label="推荐色阶">
            <g-switch v-model="settings.recommendColor" aria-label="推荐色阶" />
          </SettingItem>
          <SettingItem label="信息色跟随主色">
            <g-switch v-model="settings.isInfoFollowPrimary" aria-label="信息色跟随主色" />
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
            <g-select
              v-model="settings.size"
              class="min-w-24"
              :options="[
                { label: '默认', value: 'default' },
                { label: '大', value: 'large' },
                { label: '小', value: 'small' }
              ]"
            />
          </SettingItem>
          <SettingItem label="灰色模式">
            <g-switch v-model="settings.grayscale" aria-label="灰色模式" />
          </SettingItem>
          <SettingItem label="色弱模式">
            <g-switch v-model="settings.colourWeakness" aria-label="色弱模式" />
          </SettingItem>
          <SettingItem label="显示水印">
            <g-switch v-model="settings.watermark.visible" aria-label="显示水印" />
          </SettingItem>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia'
import { useThemeStore } from '@/pinia'
import ThemeModeSelector from '../../components/themeModeSelector.vue'
import ThemeColorPicker from '../../components/themeColorPicker.vue'
import MenuThemeSelector from '../../components/menuThemeSelector.vue'
import TabModeSelector from '../../components/tabModeSelector.vue'
import SemanticColorPicker from '../../components/semanticColorPicker.vue'
import RadiusSelector from '../../components/radiusSelector.vue'
import CardModeSelector from '../../components/cardModeSelector.vue'
import SettingItem from '../../components/settingItem.vue'

defineOptions({
  name: 'AppearanceSettings'
})

const themeStore = useThemeStore()
const { settings } = storeToRefs(themeStore)
</script>
