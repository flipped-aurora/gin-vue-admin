<script setup>
import { computed } from 'vue'
import { SelectRoot } from 'reka-ui'
import SelectTrigger from './SelectTrigger.vue'
import SelectContent from './SelectContent.vue'
import SelectItem from './SelectItem.vue'

defineOptions({ name: 'UiSelect' })

/** @typedef {{ label: string, value: string | number | boolean, disabled?: boolean }} SelectOption */

const props = defineProps({
  // 便捷模式：传入 options 即自动渲染 trigger / content / items；
  // 不传 options 则回退到默认插槽（granular 模式：自行组合 SelectTrigger / SelectContent / SelectItem）。
  /** @type {SelectOption[]} */
  options: { type: Array, default: () => [] },
  placeholder: { type: String, default: '请选择' },
  disabled: { type: Boolean, default: false },
  // 透传到触发器，便于外部传布局 class（如 min-w-24）
  class: { type: null, default: '' },
  contentClass: { type: null, default: '' }
})

// 双向绑定统一走 defineModel；放宽到 string | number | boolean，回写时保留 option.value 的原类型。
const modelValue = defineModel({ type: [String, Number, Boolean], default: undefined })

const hasOptions = computed(() => props.options.length > 0)

// reka SelectRoot / SelectItem 只按字符串匹配（item 的 value 也被 String 化），
// 这里维护 String(value) -> 原始 value 映射：入参把 modelValue 串化喂给 reka，
// 回写时再用映射还原成原始类型，避免数字 / 布尔 value 被抹平成字符串。
const valueMap = computed(() => new Map(props.options.map((o) => [String(o.value), o.value])))
const innerValue = computed(() => (modelValue.value == null ? undefined : String(modelValue.value)))
const onUpdate = (next) => {
  // granular 模式下 options 为空、映射命不中时原样透传（保持既有字符串语义）
  modelValue.value = valueMap.value.has(next) ? valueMap.value.get(next) : next
}

// 便捷模式下用本地 options 直接算出当前项文案，
// 避免依赖 reka-ui SelectValue 在面板首次挂载前回填选中文本的时机问题。
const currentLabel = computed(
  () => props.options.find((o) => String(o.value) === String(modelValue.value))?.label
)
</script>

<template>
  <SelectRoot
    :model-value="innerValue"
    :disabled="disabled"
    @update:model-value="onUpdate"
  >
    <template v-if="hasOptions">
      <SelectTrigger :class="props.class">
        <span v-if="currentLabel != null" class="truncate">{{ currentLabel }}</span>
        <span v-else class="truncate text-muted-foreground">{{ placeholder }}</span>
      </SelectTrigger>
      <SelectContent :class="contentClass">
        <SelectItem
          v-for="opt in options"
          :key="String(opt.value)"
          :value="String(opt.value)"
          :disabled="opt.disabled"
        >
          {{ opt.label }}
        </SelectItem>
      </SelectContent>
    </template>
    <slot v-else />
  </SelectRoot>
</template>
