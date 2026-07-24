<script setup>
  import { computed, ref } from 'vue'
  import { cn, FOCUS_RING_WITHIN } from '../utils'
  import { INPUT_SIZES, INPUT_VARIANTS } from './index'

  // 文本 / 密码输入框：无样式原生 <input> 外裹一层带主题描边与聚焦态的容器，
  // 颜色只用语义 token（container / border / primary / error / muted-foreground），
  // 天然跟随 themeStore 换肤与暗色模式。
  // - variant=default：盒式（描边 + container 底 + 圆角 + 聚焦环）
  // - variant=underline：下划线式（仅底边，聚焦时底边转 primary，用于极简排版）
  // password 类型可选「显示 / 隐藏」切换。
  defineOptions({ name: 'UiInput', inheritAttrs: false })

  const props = defineProps({
    // 原生 input type：text / password / email / tel ...
    type: { type: String, default: 'text' },
    size: {
      type: String,
      default: 'default',
      validator: (v) => INPUT_SIZES.includes(v)
    },
    variant: {
      type: String,
      default: 'default',
      validator: (v) => INPUT_VARIANTS.includes(v)
    },
    placeholder: { type: String, default: '' },
    disabled: { type: Boolean, default: false },
    // 前 / 后置图标：Iconify 名（优先 lucide 集，如 'lucide:user'）
    prefixIcon: { type: String, default: '' },
    suffixIcon: { type: String, default: '' },
    // 校验失败态：描边与聚焦态转 error 色
    invalid: { type: Boolean, default: false },
    // type=password 时是否显示「显示 / 隐藏密码」切换按钮
    showPassword: { type: Boolean, default: false },
    ariaLabel: { type: String, default: undefined },
    class: { type: null, default: '' }
  })

  // 双向绑定统一走 defineModel，取代手写 modelValue prop + defineEmits + @update 三段样板
  const modelValue = defineModel({ type: String, default: '' })

  const revealed = ref(false)
  const innerType = computed(() =>
    props.type === 'password' && props.showPassword && revealed.value
      ? 'text'
      : props.type
  )
  const showToggle = computed(
    () => props.type === 'password' && props.showPassword
  )

  const containerClass = computed(() => {
    const base =
      'group flex items-center gap-2 text-base-text transition-colors'
    const sizeCls = props.size === 'lg' ? 'h-11' : 'h-9'
    const disabledCls = props.disabled && 'cursor-not-allowed opacity-50'

    if (props.variant === 'underline') {
      return cn(
        base,
        sizeCls,
        'rounded-none border-0 border-b bg-transparent px-0',
        props.invalid
          ? 'border-error focus-within:border-error'
          : 'border-border focus-within:border-primary',
        disabledCls,
        props.class
      )
    }

    return cn(
      base,
      sizeCls,
      'rounded-md border bg-container px-3.5',
      props.invalid
        ? 'border-error focus-within:border-error focus-within:ring-2 focus-within:ring-error-300'
        : cn('border-border focus-within:border-primary', FOCUS_RING_WITHIN),
      disabledCls,
      props.class
    )
  })
</script>

<template>
  <div :class="containerClass">
    <svg-icon
      v-if="prefixIcon"
      :icon="prefixIcon"
      class="h-4 w-4 shrink-0 text-muted-foreground"
      aria-hidden="true"
    />
    <input
      v-bind="$attrs"
      v-model="modelValue"
      :type="innerType"
      :placeholder="placeholder"
      :disabled="disabled"
      :aria-label="ariaLabel"
      :aria-invalid="invalid || undefined"
      class="h-full w-full min-w-0 flex-1 bg-transparent text-sm text-base-text outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed"
    />
    <button
      v-if="showToggle"
      type="button"
      tabindex="-1"
      :aria-label="revealed ? '隐藏密码' : '显示密码'"
      class="flex shrink-0 items-center text-muted-foreground transition-colors hover:text-primary"
      @click="revealed = !revealed"
    >
      <svg-icon
        :icon="revealed ? 'lucide:eye' : 'lucide:eye-off'"
        class="h-4 w-4"
        aria-hidden="true"
      />
    </button>
    <svg-icon
      v-else-if="suffixIcon"
      :icon="suffixIcon"
      class="h-4 w-4 shrink-0 text-muted-foreground"
      aria-hidden="true"
    />
  </div>
</template>
