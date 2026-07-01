<script setup>
import { cn } from '../utils'

// slider 风格：底部 2px 主色指示条 + 激活弱底，扁平现代。
// 弱底用 muted（暗色安全的语义 token），指示与文字走 primary。
defineOptions({ name: 'PageTabSlider' })

const props = defineProps({
  active: { type: Boolean, default: false },
  class: { type: null, default: '' }
})
</script>

<template>
  <div
    :class="
      cn(
        'group relative inline-flex h-full cursor-pointer select-none items-center justify-center gap-1.5 whitespace-nowrap border-b-2 px-3 py-1 text-[13px] transition-colors duration-200',
        active
          ? 'border-primary bg-muted text-primary'
          : 'border-transparent text-base-text hover:text-primary',
        props.class
      )
    "
  >
    <slot name="prefix" />
    <slot />
    <slot name="suffix" />
  </div>
</template>
