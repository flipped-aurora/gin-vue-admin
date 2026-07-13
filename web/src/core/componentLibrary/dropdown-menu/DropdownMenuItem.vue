<script setup>
import { DropdownMenuItem } from 'reka-ui'
import { cn } from '../utils'

defineOptions({ name: 'UiDropdownMenuItem', inheritAttrs: false })

// 高亮态按官方 demo 走主题色实底 + 白字（data-[highlighted] 同时覆盖鼠标悬停与键盘导航）；
// reka 的 @select 事件经 $attrs 透传给调用方
const props = defineProps({
  disabled: { type: Boolean, default: false },
  // 危险语义项（如删除 / 登出），红色文本、高亮时红色实底
  danger: { type: Boolean, default: false },
  class: { type: null, default: '' }
})
</script>

<template>
  <DropdownMenuItem
    v-bind="$attrs"
    :disabled="disabled"
    :class="cn(
      'relative flex w-full cursor-pointer select-none items-center gap-2 rounded-lg px-2.5 py-2 text-[14px] text-base-text outline-none transition-colors data-[highlighted]:bg-primary data-[highlighted]:text-white data-[disabled]:pointer-events-none data-[disabled]:opacity-50',
      danger && 'text-error data-[highlighted]:bg-error data-[highlighted]:text-white',
      props.class
    )"
  >
    <slot />
  </DropdownMenuItem>
</template>
