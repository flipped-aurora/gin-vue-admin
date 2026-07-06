<template>
  <SettingItem label="语义色">
    <div class="flex items-center gap-2">
      <g-color-picker
        v-for="c in colors"
        :key="c.key"
        :model-value="settings.otherColor[c.key]"
        :title="c.label"
        :show-value="false"
        :swatches="presetSwatches"
        @update:model-value="(val) => themeStore.updateThemeColors(c.key, val)"
      />
    </div>
  </SettingItem>
</template>

<script setup>
import { storeToRefs } from 'pinia'
import { useThemeStore } from '@/pinia'
import SettingItem from './settingItem.vue'

defineOptions({
  name: 'SemanticColorPicker'
})

const themeStore = useThemeStore()
const { settings } = storeToRefs(themeStore)

const colors = [
  { key: 'success', label: '成功' },
  { key: 'warning', label: '警告' },
  { key: 'error', label: '危险' },
  { key: 'info', label: '信息' }
]

// 语义色常用预设，供面板里快速选取
const presetSwatches = ['#67c23a', '#e6a23c', '#f56c6c', '#909399', '#60c041', '#f9901f', '#194bfb', '#38c0fc']
</script>
