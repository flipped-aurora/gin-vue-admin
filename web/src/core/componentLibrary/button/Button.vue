<script setup>
import { computed } from 'vue'
import { Primitive } from 'reka-ui'
import { cn } from '../utils'
import { buttonVariants, BUTTON_VARIANTS, BUTTON_SIZES } from './index'

defineOptions({ name: 'UiButton' })

const props = defineProps({
  variant: { type: String, default: 'default', validator: (v) => BUTTON_VARIANTS.includes(v) },
  size: { type: String, default: 'default', validator: (v) => BUTTON_SIZES.includes(v) },
  // 透传给 reka-ui Primitive：默认渲染 <button>，asChild 时把行为合并到插槽根元素
  as: { type: [String, Object], default: 'button' },
  asChild: { type: Boolean, default: false },
  disabled: { type: Boolean, default: false },
  // 异步提交（保存等）时置 true：禁用重复点击并在文案前转圈
  loading: { type: Boolean, default: false },
  class: { type: null, default: '' }
})

const isDisabled = computed(() => props.disabled || props.loading)
</script>

<template>
  <Primitive
    :as="as"
    :as-child="asChild"
    :disabled="isDisabled || undefined"
    :class="cn(buttonVariants({ variant, size }), props.class)"
  >
    <span
      v-if="loading"
      class="h-3.5 w-3.5 shrink-0 animate-spin rounded-full border-2 border-current border-t-transparent"
    />
    <slot />
  </Primitive>
</template>
