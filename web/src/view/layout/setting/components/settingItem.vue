<template>
  <div class="gva-theme-setting-item">
    <div class="flex items-center gap-2">
      <span class="gva-theme-setting-label">{{ label }}</span>
      <slot name="suffix"></slot>
    </div>
    <div class="flex items-center">
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
// 抽屉内的表单控件已全部切到 reka-ui 基础组件层（按主题 token 上色），
// 这里不再需要针对 Element Plus 控件的覆盖样式与 --setting-* 变量。
defineOptions({
  name: 'SettingItem'
})

defineProps({
  label: {
    type: String,
    required: true
  }
})
</script>
