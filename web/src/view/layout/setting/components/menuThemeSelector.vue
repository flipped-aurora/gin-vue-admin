<template>
  <div class="grid grid-cols-3 gap-3 px-1">
    <div
      v-for="item in menuThemes"
      :key="item.value"
      class="relative flex flex-col items-center p-1.5 bg-white dark:bg-gray-700/50 border rounded-lg cursor-pointer transition-colors duration-150 ease-in-out hover:border-gray-300 dark:hover:border-gray-500"
      :class="modelValue === item.value ? '' : 'border-gray-200 dark:border-gray-600'"
      :style="modelValue === item.value ? { borderColor: primaryColor, boxShadow: `0 0 0 1px ${primaryColor}` } : {}"
      :title="item.label"
      @click="handleChange(item.value)"
    >
      <!-- 迷你预览：左侧栏 + 主区 -->
      <div class="w-full h-9 rounded overflow-hidden flex border border-gray-200 dark:border-gray-600">
        <div class="w-1/3 h-full flex flex-col gap-1 p-1" :style="{ background: item.previewSide }">
          <div class="h-1.5 rounded-sm" :style="{ background: item.previewActive }"></div>
          <div class="h-1.5 rounded-sm opacity-40" :style="{ background: item.previewSideText }"></div>
        </div>
        <div class="flex-1 h-full" :style="{ background: item.previewMain }"></div>
      </div>
      <svg-icon
        v-if="modelValue === item.value"
        icon="lucide:check"
        class="absolute top-1 right-1"
        :style="{ color: primaryColor }"
      />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useThemeStore } from '@/pinia'

defineOptions({
  name: 'MenuThemeSelector'
})

defineProps({
  modelValue: {
    type: String,
    default: 'design'
  }
})

const emit = defineEmits(['update:modelValue'])

const themeStore = useThemeStore()
const { settings } = storeToRefs(themeStore)
const primaryColor = computed(() => settings.value.themeColor)

const menuThemes = [
  // 设计：浅主色药丸选中态
  { value: 'design', label: '设计', previewSide: '#ffffff', previewMain: '#f5f6f8', previewActive: 'rgb(var(--primary-color) / 0.15)', previewSideText: '#94a3b8' },
  // 亮色：主色实底选中态
  { value: 'light', label: '亮色', previewSide: '#ffffff', previewMain: '#f5f6f8', previewActive: 'rgb(var(--primary-color))', previewSideText: '#94a3b8' },
  // 分组：通栏纯侧栏，一级为可折叠分组标题，选中为浅灰药丸
  { value: 'group', label: '分组', previewSide: '#ffffff', previewMain: '#f5f6f8', previewActive: '#e2e8f0', previewSideText: '#64748b' }
]

const handleChange = (value) => {
  emit('update:modelValue', value)
}
</script>
