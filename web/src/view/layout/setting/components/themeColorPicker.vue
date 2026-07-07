<template>
  <div class="gva-theme-card-bg py-3">
    <div class="flex items-center flex-wrap gap-2.5">
      <button
        v-for="colorItem in presetColors"
        :key="colorItem.color"
        type="button"
        class="relative w-6 h-6 rounded-full cursor-pointer border text-white border-black/5 transition-transform duration-150 hover:scale-110"
        :class="modelValue === colorItem.color ? 'scale-110 ring-2 ring-offset-2 ring-offset-gray-50 dark:ring-offset-gray-800' : ''"
        :style="{
          backgroundColor: colorItem.color,
          '--tw-ring-color': colorItem.color
        }"
        :title="colorItem.name"
        @click="handleColorChange(colorItem.color)"
      >
        <svg-icon
          v-if="modelValue === colorItem.color"
          icon="lucide:check"
          class="absolute inset-0 m-auto"
          style="font-size: 12px; text-shadow: 0 1px 2px rgba(0,0,0,.35)"
        />
      </button>

      <span class="w-px h-5 bg-gray-200 dark:bg-gray-600 mx-0.5"></span>

      <g-color-picker
        :model-value="modelValue"
        :show-value="false"
        :swatches="presetColors.map((item) => item.color)"
        aria-label="自定义主题色"
        @update:model-value="handleCustomColorChange"
      />

      <code class="ml-auto text-xs font-mono text-gray-400 dark:text-gray-500">{{ modelValue }}</code>
    </div>
  </div>
</template>

<script setup>
defineOptions({
  name: 'ThemeColorPicker'
})

// 双向绑定统一走 defineModel，取代手写 modelValue prop + emit
const modelValue = defineModel({
  type: String,
  default: '#2264f2'
})

const presetColors = [
  { color: '#2264f2', name: '默认' },
  { color: '#b48df3', name: '雅紫' },
  { color: '#1d84ff', name: '天蓝' },
  { color: '#60c041', name: '清新绿' },
  { color: '#38c0fc', name: '湖青' },
  { color: '#f9901f', name: '活力橙' },
  { color: '#ff80c8', name: '樱粉' }
]

const handleColorChange = (color) => {
  modelValue.value = color
}

const handleCustomColorChange = (color) => {
  if (color) {
    modelValue.value = color
  }
}
</script>
