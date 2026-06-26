<template>
  <div class="flex gap-1 p-1 bg-gray-100 dark:bg-gray-800 rounded-lg w-fit">
    <div
      v-for="item in modes"
      :key="item.value"
      class="px-5 py-1.5 text-sm font-medium rounded-md cursor-pointer transition-all duration-150 ease-in-out"
      :class="modelValue === item.value
        ? 'text-white shadow-sm'
        : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200'"
      :style="modelValue === item.value ? { backgroundColor: primaryColor } : {}"
      @click="emit('update:modelValue', item.value)"
    >
      {{ item.label }}
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useThemeStore } from '@/pinia'

defineOptions({
  name: 'CardModeSelector'
})

defineProps({
  modelValue: {
    type: String,
    default: 'border'
  }
})

const emit = defineEmits(['update:modelValue'])

const themeStore = useThemeStore()
const { settings } = storeToRefs(themeStore)
const primaryColor = computed(() => settings.value.themeColor)

const modes = [
  { value: 'border', label: '边框' },
  { value: 'shadow', label: '阴影' }
]
</script>
