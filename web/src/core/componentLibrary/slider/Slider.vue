<script setup>
import { computed } from 'vue'
import { SliderRoot, SliderTrack, SliderRange, SliderThumb } from 'reka-ui'
import { cn, FOCUS_RING } from '../utils'

defineOptions({ name: 'UiSlider' })

/** @typedef {Record<string | number, string>} SliderMarks 刻度：键为刻度值、值为下方文字 */

// 对外用「单值 number」表达，内部按 reka 约定包成数组。
const props = defineProps({
  min: { type: Number, default: 0 },
  max: { type: Number, default: 100 },
  step: { type: Number, default: 1 },
  // 形如 { 0: '0', 0.5: '0.5', 1: '1' }，在轨道下方渲染刻度文字
  marks: { type: Object, default: null },
  disabled: { type: Boolean, default: false },
  // 拨块的可访问名（屏幕阅读器），调用方按语义传入，如「圆角大小」
  ariaLabel: { type: String, default: undefined },
  class: { type: null, default: '' }
})

// 双向绑定走 defineModel（对外单值 number），用可写 computed 桥到 reka 要求的数组形态
const modelValue = defineModel({ type: Number, default: 0 })
const sliderValue = computed({
  get: () => [modelValue.value],
  set: (val) => {
    modelValue.value = Array.isArray(val) ? val[0] : val
  }
})

const markList = computed(() => {
  if (!props.marks) return []
  const span = props.max - props.min || 1
  return Object.entries(props.marks).map(([value, label]) => ({
    value: Number(value),
    label,
    percent: ((Number(value) - props.min) / span) * 100
  }))
})
</script>

<template>
  <div :class="cn('w-full', props.class)">
    <SliderRoot
      v-model="sliderValue"
      :min="min"
      :max="max"
      :step="step"
      :disabled="disabled"
      class="relative cursor-pointer flex h-5 w-full touch-none select-none items-center disabled:opacity-50"
    >
      <SliderTrack class="relative h-1.5 w-full grow overflow-hidden rounded-full bg-control-track">
        <SliderRange class="absolute h-full rounded-full bg-primary" />
      </SliderTrack>
      <SliderThumb
        :aria-label="ariaLabel"
        :class="cn('block h-4 w-4 rounded-full border-2 border-primary bg-container shadow transition-colors', FOCUS_RING)"
      />
    </SliderRoot>
    <div v-if="markList.length" class="relative mt-1 h-3.5">
      <span
        v-for="m in markList"
        :key="m.value"
        class="absolute -translate-x-1/2 text-[10px] leading-none text-muted-foreground"
        :style="{ left: m.percent + '%' }"
      >{{ m.label }}</span>
    </div>
  </div>
</template>
