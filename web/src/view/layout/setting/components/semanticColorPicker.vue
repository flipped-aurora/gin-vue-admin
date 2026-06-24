<template>
  <SettingItem label="语义色">
    <div class="flex items-center gap-2">
      <el-color-picker
        v-for="c in colors"
        :key="c.key"
        :model-value="settings.otherColor[c.key]"
        :title="c.label"
        size="small"
        @update:modelValue="(val) => themeStore.updateThemeColors(c.key, val)"
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
</script>
