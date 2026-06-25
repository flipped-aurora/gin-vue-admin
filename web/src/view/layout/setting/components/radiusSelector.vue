<template>
  <div class="flex items-center gap-4 w-full min-w-[220px]">
    <g-slider
      v-model="modelValue"
      :min="0"
      :max="1"
      :step="0.05"
      :marks="marks"
      aria-label="全局圆角"
      class="flex-1"
    />
    <span class="text-xs font-mono gva-theme-text-main w-14 text-right">{{ modelValue }}rem</span>
  </div>
</template>

<script setup>
defineOptions({
  name: 'RadiusSelector'
})

// 双向绑定统一走 defineModel，取代手写 modelValue prop + emit
const modelValue = defineModel({
  type: Number,
  default: 0.75
})

const marks = {
  0: '0',
  0.5: '0.5',
  1: '1'
}
</script>
