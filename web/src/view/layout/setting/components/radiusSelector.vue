<template>
  <div class="flex items-center gap-4 w-full min-w-[220px]">
    <el-slider
      :model-value="modelValue"
      :min="0"
      :max="1"
      :step="0.05"
      :marks="marks"
      class="flex-1"
      @input="onInput"
    />
    <span class="text-xs font-mono gva-theme-text-main w-14 text-right">{{ modelValue }}rem</span>
  </div>
</template>

<script setup>
defineOptions({
  name: 'RadiusSelector'
})

defineProps({
  modelValue: {
    type: Number,
    default: 0.75
  }
})

const emit = defineEmits(['update:modelValue'])

const marks = {
  0: '0',
  0.5: '0.5',
  1: '1'
}

const onInput = (val) => {
  emit('update:modelValue', val)
}
</script>
