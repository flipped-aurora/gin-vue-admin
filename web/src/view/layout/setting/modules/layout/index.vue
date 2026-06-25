<template>
  <div class="gva-theme-font">
    <!-- 布局模式 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">布局模式</span>
      </div>
      <div class="gva-theme-section-content">
        <LayoutModeCard v-model="settings.layout.mode" />
      </div>
    </div>

    <!-- 顶栏：可见性 + 配色 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">顶栏</span>
      </div>
      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg">
          <SettingItem label="显示面包屑">
            <g-switch v-model="settings.header.breadcrumb.visible" aria-label="显示面包屑" />
          </SettingItem>
          <SettingItem label="显示面包屑图标">
            <g-switch
              v-model="settings.header.breadcrumb.showIcon"
              :disabled="!settings.header.breadcrumb.visible"
              aria-label="显示面包屑图标"
            />
          </SettingItem>
          <SettingItem label="显示刷新按钮">
            <g-switch v-model="settings.header.refresh.visible" aria-label="显示刷新按钮" />
          </SettingItem>
          <SettingItem label="显示搜索按钮">
            <g-switch v-model="settings.header.search.visible" aria-label="显示搜索按钮" />
          </SettingItem>
          <SettingItem label="显示折叠按钮">
            <g-switch v-model="settings.header.collapseButton.visible" aria-label="显示折叠按钮" />
          </SettingItem>
          <SettingItem label="顶栏背景">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">留空跟随主题</span>
            </template>
            <g-color-picker
              v-model="settings.header.bg"
              alpha
              format="rgb"
              placeholder="跟随主题"
              aria-label="顶栏背景"
            />
          </SettingItem>
          <SettingItem label="顶栏阴影">
            <g-select
              v-model="settings.header.shadow"
              class="min-w-24"
              :options="[
                { label: '无', value: 'none' },
                { label: '小', value: 'sm' },
                { label: '中', value: 'md' },
                { label: '大', value: 'lg' }
              ]"
            />
          </SettingItem>
          <SettingItem label="标签栏背景">
            <g-color-picker
              v-model="settings.tab.bg"
              alpha
              format="rgb"
              placeholder="跟随主题"
              aria-label="标签栏背景"
            />
          </SettingItem>
          <div class="flex items-center gap-1.5 py-2.5 text-xs text-gray-400 dark:text-gray-500 leading-snug">
            <svg-icon icon="lucide:info" class="flex-shrink-0" />
            <span>暗色模式下将基于以上配色自动推导深色版本，无需单独设置</span>
          </div>
        </div>
      </div>
    </div>

    <!-- 界面 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">界面</span>
      </div>
      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg">
          <SettingItem label="显示标签页">
            <g-switch v-model="settings.tab.visible" aria-label="显示标签页" />
          </SettingItem>
          <SettingItem label="页面切换动画">
            <g-select
              v-model="settings.page.transition"
              class="min-w-24"
              :options="[
                { label: '淡入淡出', value: 'fade' },
                { label: '滑动', value: 'slide' },
                { label: '缩放', value: 'zoom' },
                { label: '无动画', value: 'none' }
              ]"
            />
          </SettingItem>
        </div>
      </div>
    </div>

    <!-- 侧栏尺寸 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">侧栏尺寸</span>
      </div>
      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg">
          <SettingItem label="展开宽度">
            <g-number-field v-model="settings.layout.sideWidth" :min="150" :max="400" :step="10" class="w-28" />
          </SettingItem>
          <SettingItem label="收缩宽度">
            <g-number-field v-model="settings.layout.sideCollapsedWidth" :min="60" :max="100" class="w-28" />
          </SettingItem>
          <SettingItem label="菜单项高度">
            <g-number-field v-model="settings.layout.sideItemHeight" :min="30" :max="50" class="w-28" />
          </SettingItem>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia'
import { useThemeStore } from '@/pinia'
import LayoutModeCard from '../../components/layoutModeCard.vue'
import SettingItem from '../../components/settingItem.vue'

defineOptions({
  name: 'LayoutSettings'
})

const themeStore = useThemeStore()
const { settings } = storeToRefs(themeStore)
</script>
