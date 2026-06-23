<template>
  <SettingItem label="语义色">
    <div class="flex items-center gap-2">
      <el-color-picker
        v-for="c in colors"
        :key="c.key"
        :model-value="config[c.key]"
        :title="c.label"
        size="small"
        @update:modelValue="(val) => appStore.toggleSemanticColor(c.key, val)"
      />
    </div>
  </SettingItem>
</template>

<script setup>
import { storeToRefs } from 'pinia'
import { useAppStore } from '@/pinia'
import SettingItem from './settingItem.vue'

defineOptions({
  name: 'SemanticColorPicker'
})

const appStore = useAppStore()
const { config } = storeToRefs(appStore)

const colors = [
  { key: 'successColor', label: '成功' },
  { key: 'warningColor', label: '警告' },
  { key: 'dangerColor', label: '危险' },
  { key: 'infoColor', label: '信息' }
]
</script>
