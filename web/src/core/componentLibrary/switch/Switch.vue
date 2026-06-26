<script setup>
import { SwitchRoot, SwitchThumb } from 'reka-ui'
import { cn, FOCUS_RING } from '../utils'

defineOptions({ name: 'UiSwitch' })

const props = defineProps({
  disabled: { type: Boolean, default: false },
  // 纯图形控件的可访问名：屏幕阅读器据此朗读该开关含义，调用方按语义传入
  ariaLabel: { type: String, default: undefined },
  class: { type: null, default: '' }
})

// 双向绑定统一走 defineModel，取代手写 modelValue prop + defineEmits + @update 三段样板
const modelValue = defineModel({ type: Boolean, default: false })
</script>

<template>
  <SwitchRoot
    v-model="modelValue"
    :disabled="disabled"
    :aria-label="ariaLabel"
    :class="cn(
      'relative inline-flex h-5 w-9 shrink-0 cursor-pointer items-center rounded-full border border-transparent transition-colors duration-200 bg-control-track data-[state=checked]:bg-primary disabled:cursor-not-allowed disabled:opacity-50',
      FOCUS_RING,
      props.class
    )"
  >
    <!-- 旋钮固定浅色：需在 off=control-track / on=primary 两种轨道色上都清晰可见，
         故不随暗色翻深、不走 bg-container（旋钮无描边，仅靠填充色与轨道对比） -->
    <SwitchThumb
      class="pointer-events-none block h-4 w-4 translate-x-0.5 rounded-full bg-white shadow transition-transform duration-200 data-[state=checked]:translate-x-[14px]"
    />
  </SwitchRoot>
</template>
