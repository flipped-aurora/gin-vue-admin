<script setup>
import { ref, computed, watch, onBeforeUnmount } from 'vue'
import {
  PopoverRoot,
  PopoverTrigger,
  PopoverPortal,
  PopoverContent,
  ColorAreaRoot,
  ColorAreaArea,
  ColorAreaThumb,
  ColorSliderRoot,
  ColorSliderTrack,
  ColorSliderThumb,
  ColorFieldRoot,
  ColorFieldInput,
  ColorSwatchPickerRoot,
  ColorSwatchPickerItem,
  parseColor,
  colorToString
} from 'reka-ui'
import { themeSettings } from '@/theme/settings'
import { cn, FOCUS_RING, FOCUS_RING_WITHIN } from '../utils'
import { COLOR_FORMATS } from './index'

defineOptions({ name: 'UiColorPicker' })

const props = defineProps({
  // 是否提供透明度通道（顶栏 / 标签背景需要 rgba）
  alpha: { type: Boolean, default: false },
  // 'hex' | 'rgb'：输出格式。语义 / 主题色用 hex；带 alpha 的背景用 rgb。
  format: { type: String, default: 'hex', validator: (v) => COLOR_FORMATS.includes(v) },
  // 预设色卡（十六进制数组）
  swatches: { type: Array, default: () => [] },
  // 空值时触发器上的占位文案（如顶栏背景"留空跟随主题"）
  placeholder: { type: String, default: '默认' },
  // 触发器是否展示颜色文本：语义色那种一排小色块用 false（只显示色块）
  showValue: { type: Boolean, default: true },
  // 原生 title（hover 提示）；无障碍场景下同时作为可访问名兜底
  title: { type: String, default: undefined },
  // 纯图形触发器的可访问名（屏幕阅读器），优先级高于 title / placeholder
  ariaLabel: { type: String, default: undefined },
  class: { type: null, default: '' }
})

// 双向绑定统一走 defineModel，取代手写 modelValue prop + emit
const modelValue = defineModel({ type: String, default: '' })

// area / hue 用 hsb 取「饱和度·明度」面板；rgb 模式（带 alpha 的背景）整套走 rgb。
const colorSpace = computed(() => (props.format === 'rgb' ? 'rgb' : 'hsb'))
// alpha 模式输出带透明度通道的格式（rgba / hexa），无 alpha 时输出 rgb / hex
const outFormat = computed(() =>
  props.format === 'rgb' ? (props.alpha ? 'rgba' : 'rgb') : props.alpha ? 'hexa' : 'hex'
)
// 空值兜底色取自默认主题色单一真源（settings.themeColor），不再手抄字面量 #2264f2 / rgba(93,135,255,1)。
// alpha 模式补足透明度通道（满不透明）；派生失败则回落到 hex 本身（仍可被 parseColor 解析）。
const fallback = computed(() => {
  const base = themeSettings.themeColor
  if (!props.alpha) return base
  try {
    return colorToString(parseColor(base), 'rgba')
  } catch {
    return base
  }
})

const safeParse = (value) => {
  try {
    return parseColor(value)
  } catch {
    return parseColor(fallback.value)
  }
}

// 关键：所有颜色部件共享同一个 reka Color「对象」（而非各自解析字符串），
// 这样 area 的饱和度/明度、hue 条、alpha 条、数值框、色卡始终基于同一份颜色，互不丢失分量。
const color = ref(safeParse(modelValue.value || fallback.value))

const toOut = (c) => {
  try {
    return colorToString(c, outFormat.value)
  } catch {
    return modelValue.value
  }
}

// 拖动色相/面板时 @update:color 会高频触发，父级（全局主题应用）开销大导致卡顿。
// 本地 color 立即更新保证滑块跟手；对外写回 modelValue 防抖，停手 100ms 后才应用主题。
let emitTimer = null
let pendingColor = null
const cancelEmit = () => {
  if (emitTimer) {
    clearTimeout(emitTimer)
    emitTimer = null
  }
}
// flush：立刻把挂起的最终颜色写回（卸载 / 收尾时调用），
// 避免「拖完颜色 <100ms 内就关抽屉 / 切 tab / 重置」导致最终值被丢弃、所见≠所存。
const flushEmit = () => {
  if (!emitTimer) return
  cancelEmit()
  if (pendingColor != null) {
    modelValue.value = toOut(pendingColor)
    pendingColor = null
  }
}
const onColorUpdate = (next) => {
  color.value = next
  pendingColor = next
  cancelEmit()
  emitTimer = setTimeout(() => {
    emitTimer = null
    pendingColor = null
    modelValue.value = toOut(next)
  }, 100)
}
onBeforeUnmount(flushEmit)

// 预设色卡：ColorSwatchPicker 的 v-model 是十六进制字符串
const hexValue = computed(() => {
  try {
    return colorToString(color.value, 'hex')
  } catch {
    return ''
  }
})
const onSwatchPick = (hex) => {
  if (!hex) return
  onColorUpdate(safeParse(hex))
}

// 外部 modelValue 变化（如点主题色预设圆点）时同步进来；与当前一致则跳过，避免回环。
watch(modelValue, (value) => {
  if (!value) return
  if (value !== toOut(color.value)) color.value = safeParse(value)
})

const hasValue = computed(() => Boolean(modelValue.value))
const previewStyle = computed(() => ({ backgroundColor: modelValue.value || 'transparent' }))

// 透明区域用的棋盘格底纹：格子色走 border token（亮 / 暗自适应），不再硬编码裸 hex。
const CHECKER =
  'linear-gradient(45deg,rgb(var(--border-color)) 25%,transparent 25%),linear-gradient(-45deg,rgb(var(--border-color)) 25%,transparent 25%),linear-gradient(45deg,transparent 75%,rgb(var(--border-color)) 75%),linear-gradient(-45deg,transparent 75%,rgb(var(--border-color)) 75%)'
const checkerStyle = {
  backgroundImage: CHECKER,
  backgroundSize: '8px 8px',
  backgroundPosition: '0 0,0 4px,4px -4px,-4px 0'
}
</script>

<template>
  <PopoverRoot>
    <PopoverTrigger as-child>
      <button
        type="button"
        :aria-label="ariaLabel || title || placeholder"
        :title="title"
        :class="cn(
          'inline-flex h-8 items-center gap-2 rounded-md border border-border bg-container text-[13px] text-base-text transition-colors hover:border-primary',
          FOCUS_RING,
          showValue ? 'px-2' : 'w-8 justify-center px-0',
          props.class
        )"
      >
        <span
          class="relative shrink-0 overflow-hidden rounded border border-border"
          :class="showValue ? 'h-4 w-4' : 'h-5 w-5'"
          :style="checkerStyle"
        >
          <span class="absolute inset-0" :style="previewStyle" />
        </span>
        <template v-if="showValue">
          <span v-if="hasValue" class="font-mono text-xs text-muted-foreground">{{ modelValue }}</span>
          <span v-else class="text-xs text-muted-foreground">{{ placeholder }}</span>
        </template>
      </button>
    </PopoverTrigger>

    <PopoverPortal>
      <PopoverContent
        :side-offset="6"
        align="end"
        class="z-[3000] w-60 rounded-lg border border-border bg-container p-3 text-base-text shadow-lg"
      >
        <!-- 饱和度 / 明度面板 -->
        <ColorAreaRoot
          v-slot="{ style }"
          :model-value="color"
          :color-space="colorSpace"
          x-channel="saturation"
          y-channel="brightness"
          class="relative h-32 w-full overflow-hidden rounded-md"
          @update:color="onColorUpdate"
        >
          <ColorAreaArea :style="style" class="h-full w-full">
            <ColorAreaThumb class="h-3.5 w-3.5 rounded-full border-2 border-white shadow ring-1 ring-black/25" />
          </ColorAreaArea>
        </ColorAreaRoot>

        <!-- 色相条 -->
        <ColorSliderRoot
          v-slot="{ style }"
          :model-value="color"
          channel="hue"
          :color-space="colorSpace"
          class="relative mt-3 flex h-4 w-full items-center"
          @update:color="onColorUpdate"
        >
          <ColorSliderTrack :style="style" class="h-2 w-full rounded-full" />
          <ColorSliderThumb class="block h-4 w-4 rounded-full border-2 border-white shadow ring-1 ring-black/25" />
        </ColorSliderRoot>

        <!-- 透明度条 -->
        <ColorSliderRoot
          v-if="alpha"
          v-slot="{ style }"
          :model-value="color"
          channel="alpha"
          :color-space="colorSpace"
          class="relative mt-3 flex h-4 w-full items-center"
          @update:color="onColorUpdate"
        >
          <div class="relative h-2 w-full overflow-hidden rounded-full" :style="checkerStyle">
            <ColorSliderTrack :style="style" class="absolute inset-0 rounded-full" />
          </div>
          <ColorSliderThumb class="block h-4 w-4 rounded-full border-2 border-white shadow ring-1 ring-black/25" />
        </ColorSliderRoot>

        <!-- 数值输入 -->
        <ColorFieldRoot
          :model-value="color"
          :class="cn('mt-3 flex h-8 items-center rounded-md border border-border bg-container px-2 transition-colors focus-within:border-primary', FOCUS_RING_WITHIN)"
          @update:color="onColorUpdate"
        >
          <ColorFieldInput class="w-full bg-transparent text-center font-mono text-xs text-base-text outline-none" />
        </ColorFieldRoot>

        <!-- 预设色卡 -->
        <ColorSwatchPickerRoot
          v-if="swatches.length"
          :model-value="hexValue"
          class="mt-3 flex flex-wrap gap-1.5"
          @update:model-value="onSwatchPick"
        >
          <ColorSwatchPickerItem
            v-for="s in swatches"
            :key="s"
            :value="s"
            :class="cn('h-5 w-5 cursor-pointer overflow-hidden rounded border border-border transition-transform hover:scale-110 data-[state=checked]:ring-2 data-[state=checked]:ring-primary data-[state=checked]:ring-offset-1 data-[state=checked]:ring-offset-container', FOCUS_RING)"
          >
            <!-- Swatch 自动取色（Alpha 组件渲染不稳定），直接用 hex 画背景就完事了 -->
            <span class="block h-full w-full" :style="{ backgroundColor: s }" />
          </ColorSwatchPickerItem>
        </ColorSwatchPickerRoot>
      </PopoverContent>
    </PopoverPortal>
  </PopoverRoot>
</template>
