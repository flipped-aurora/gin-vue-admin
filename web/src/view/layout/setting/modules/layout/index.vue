<template>
  <div class="font-inter">
    <div class="mb-10">
      <div class="flex items-center justify-center mb-6">
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
        <span class="px-6 text-lg font-semibold text-gray-700 dark:text-gray-300">布局模式</span>
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
      </div>

      <div class="section-content">
        <LayoutModeCard
          v-model="config.side_mode"
          @update:modelValue="appStore.toggleSideMode"
        />
      </div>
    </div>

    <div class="mb-10">
      <div class="flex items-center justify-center mb-6">
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
        <span class="px-6 text-lg font-semibold text-gray-700 dark:text-gray-300">界面配置</span>
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
      </div>

      <div class="section-content">
        <div class="bg-gray-50 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl p-6 shadow-sm">
          <SettingItem label="显示标签页">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">页面标签导航</span>
            </template>
            <el-switch
              v-model="config.showTabs"
              @change="appStore.toggleTabs"
            />
          </SettingItem>

          <SettingItem label="页面切换动画">
            <template #suffix>
              <span class="text-xs text-gray-400 dark:text-gray-500 ml-2">页面过渡效果</span>
            </template>
            <el-select
              v-model="config.transition_type"
              @change="appStore.toggleTransition"
              class="w-32"
              size="small"
            >
              <el-option value="fade" label="淡入淡出" />
              <el-option value="slide" label="滑动" />
              <el-option value="zoom" label="缩放" />
              <el-option value="none" label="无动画" />
            </el-select>
          </SettingItem>
        </div>
      </div>
    </div>

    <div class="mb-10">
      <div class="flex items-center justify-center mb-6">
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
        <span class="px-6 text-lg font-semibold text-gray-700 dark:text-gray-300">尺寸配置</span>
        <div class="h-px bg-gray-200 dark:bg-gray-700 flex-1"></div>
      </div>

      <div class="section-content">
        <div class="bg-gray-50 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl p-6 shadow-sm">
          <div class="space-y-6">
            <div class="bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-lg p-5 hover:shadow-md transition-all duration-150 ease-in-out hover:-translate-y-0.5">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="text-sm font-medium text-gray-900 dark:text-white">侧边栏展开宽度</h4>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">侧边栏完全展开时的宽度</p>
                </div>
                <div class="flex items-center gap-2">
                  <el-input-number
                    v-model="config.layout_side_width"
                    :min="150"
                    :max="400"
                    :step="10"
                    size="small"
                    class="w-24"
                  />
                  <span class="text-xs font-medium text-gray-500 dark:text-gray-400">px</span>
                </div>
              </div>
            </div>

            <div class="bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-lg p-5 hover:shadow-md transition-all duration-150 ease-in-out hover:-translate-y-0.5">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="text-sm font-medium text-gray-900 dark:text-white">侧边栏收缩宽度</h4>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">侧边栏收缩时的最小宽度</p>
                </div>
                <div class="flex items-center gap-2">
                  <el-input-number
                    v-model="config.layout_side_collapsed_width"
                    :min="60"
                    :max="100"
                    size="small"
                    class="w-24"
                  />
                  <span class="text-xs font-medium text-gray-500 dark:text-gray-400">px</span>
                </div>
              </div>
            </div>

            <div class="bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-lg p-5 hover:shadow-md transition-all duration-150 ease-in-out hover:-translate-y-0.5">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="text-sm font-medium text-gray-900 dark:text-white">菜单项高度</h4>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">侧边栏菜单项的行高</p>
                </div>
                <div class="flex items-center gap-2">
                  <el-input-number
                    v-model="config.layout_side_item_height"
                    :min="30"
                    :max="50"
                    size="small"
                    class="w-24"
                  />
                  <span class="text-xs font-medium text-gray-500 dark:text-gray-400">px</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia'
import { useAppStore } from '@/pinia'
import LayoutModeCard from '../../components/layoutModeCard.vue'
import SettingItem from '../../components/settingItem.vue'

defineOptions({
  name: 'LayoutSettings'
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
