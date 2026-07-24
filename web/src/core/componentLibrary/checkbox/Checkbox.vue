<script setup>
  import { CheckboxRoot, CheckboxIndicator } from 'reka-ui'
  import { cn, FOCUS_RING } from '../utils'

  // 复选框：reka-ui CheckboxRoot / CheckboxIndicator 底座，颜色只用语义 token。
  // 关态描边 + container 底，选中态 primary 实底 + 白色对勾（图标走全局 svg-icon）。
  defineOptions({ name: 'UiCheckbox' })

  const props = defineProps({
    disabled: { type: Boolean, default: false },
    // 纯图形控件的可访问名：无可见文案时由调用方按语义传入
    ariaLabel: { type: String, default: undefined },
    class: { type: null, default: '' }
  })

  // reka CheckboxRoot 用 `checked` 双向绑定；对外仍走 defineModel 的 modelValue
  const modelValue = defineModel({ type: Boolean, default: false })
</script>

<template>
  <CheckboxRoot
    v-model:checked="modelValue"
    :disabled="disabled"
    :aria-label="ariaLabel"
    :class="
      cn(
        'inline-flex h-4 w-4 shrink-0 items-center justify-center rounded border border-border bg-container text-white transition-colors data-[state=checked]:border-primary data-[state=checked]:bg-primary disabled:cursor-not-allowed disabled:opacity-50',
        FOCUS_RING,
        props.class
      )
    "
  >
    <CheckboxIndicator class="flex items-center justify-center">
      <svg-icon icon="lucide:check" class="h-3 w-3" aria-hidden="true" />
    </CheckboxIndicator>
  </CheckboxRoot>
</template>
