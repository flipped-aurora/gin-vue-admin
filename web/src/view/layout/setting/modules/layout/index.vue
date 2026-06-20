<template>
  <div class="gva-theme-font">
    <!-- 布局模式 -->
    <div class="mb-6">
      <div class="gva-theme-section-header">
        <span class="gva-theme-section-title">布局模式</span>
      </div>
      <div class="gva-theme-section-content">
        <LayoutModeCard v-model="config.side_mode" @update:modelValue="appStore.toggleSideMode" />
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
            <el-switch v-model="config.show_breadcrumb" />
          </SettingItem>
          <SettingItem label="显示面包屑图标">
            <el-switch v-model="config.show_breadcrumb_icon" :disabled="!config.show_breadcrumb" />
          </SettingItem>
          <SettingItem label="显示刷新按钮">
            <el-switch v-model="config.show_refresh" />
          </SettingItem>
          <SettingItem label="显示搜索按钮">
            <el-switch v-model="config.show_search" />
          </SettingItem>
          <SettingItem label="显示折叠按钮">
            <el-switch v-model="config.show_collapse_btn" />
          </SettingItem>
          <SettingItem label="顶栏背景">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">留空跟随主题</span>
            </template>
            <el-color-picker v-model="config.header_bg" show-alpha color-format="rgb" size="small" />
          </SettingItem>
          <SettingItem label="顶栏边框">
            <el-color-picker v-model="config.header_border" show-alpha color-format="rgb" size="small" />
          </SettingItem>
          <SettingItem label="标签栏背景">
            <el-color-picker v-model="config.tabs_bg" show-alpha color-format="rgb" size="small" />
          </SettingItem>
          <div class="flex items-center gap-1.5 py-2.5 text-xs text-gray-400 dark:text-gray-500 leading-snug">
            <el-icon class="flex-shrink-0">
              <InfoFilled />
            </el-icon>
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
            <el-switch v-model="config.showTabs" @change="appStore.toggleTabs" />
          </SettingItem>
          <SettingItem label="页面切换动画">
            <el-select v-model="config.transition_type" class="w-36" size="small" @change="appStore.toggleTransition">
              <el-option value="fade" label="淡入淡出" />
              <el-option value="slide" label="滑动" />
              <el-option value="zoom" label="缩放" />
              <el-option value="none" label="无动画" />
            </el-select>
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
            <el-input-number v-model="config.layout_side_width" :min="150" :max="400" :step="10" size="small" class="w-28" controls-position="right" />
          </SettingItem>
          <SettingItem label="收缩宽度">
            <el-input-number v-model="config.layout_side_collapsed_width" :min="60" :max="100" size="small" class="w-28" controls-position="right" />
          </SettingItem>
          <SettingItem label="菜单项高度">
            <el-input-number v-model="config.layout_side_item_height" :min="30" :max="50" size="small" class="w-28" controls-position="right" />
          </SettingItem>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia'
import { useAppStore } from '@/pinia'
import { InfoFilled } from '@element-plus/icons-vue'
import LayoutModeCard from '../../components/layoutModeCard.vue'
import SettingItem from '../../components/settingItem.vue'

defineOptions({
  name: 'LayoutSettings'
})

const appStore = useAppStore()
const { config } = storeToRefs(appStore)
</script>
