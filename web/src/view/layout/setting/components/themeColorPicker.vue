<template>
  <div class="gva-theme-card-bg py-3">
    <div class="flex items-center flex-wrap gap-2.5">
      <button
        v-for="colorItem in presetColors"
        :key="colorItem.color"
        type="button"
        class="relative w-6 h-6 rounded-full cursor-pointer border border-black/5 transition-transform duration-150 hover:scale-110"
        :class="modelValue === colorItem.color ? 'scale-110 ring-2 ring-offset-2 ring-offset-gray-50 dark:ring-offset-gray-800' : ''"
        :style="{
          backgroundColor: colorItem.color,
          '--tw-ring-color': colorItem.color
        }"
        :title="colorItem.name"
        @click="handleColorChange(colorItem.color)"
      >
        <el-icon
          v-if="modelValue === colorItem.color"
          class="absolute inset-0 m-auto text-white"
          style="font-size: 12px; text-shadow: 0 1px 2px rgba(0,0,0,.35)"
        >
          <Check />
        </el-icon>
      </button>

      <span class="w-px h-5 bg-gray-200 dark:bg-gray-600 mx-0.5"></span>

      <el-color-picker
        v-model="customColor"
        :predefine="presetColors.map((item) => item.color)"
        @change="handleCustomColorChange"
      />

      <code class="ml-auto text-xs font-mono text-gray-400 dark:text-gray-500">{{ modelValue }}</code>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { Check } from '@element-plus/icons-vue'

defineOptions({
  name: 'ThemeColorPicker'
})

const props = defineProps({
  modelValue: {
    type: String,
    default: '#5D87FF'
  }
})

const emit = defineEmits(['update:modelValue'])

const customColor = ref(props.modelValue)

const presetColors = [
  { color: '#5D87FF', name: '默认' },
  { color: '#B48DF3', name: '雅紫' },
  { color: '#1D84FF', name: '天蓝' },
  { color: '#60C041', name: '清新绿' },
  { color: '#38C0FC', name: '湖青' },
  { color: '#F9901F', name: '活力橙' },
  { color: '#FF80C8', name: '樱粉' }
]

const handleColorChange = (color) => {
  customColor.value = color
  emit('update:modelValue', color)
}

const handleCustomColorChange = (color) => {
  if (color) {
    emit('update:modelValue', color)
  }
}

watch(
  () => props.modelValue,
  (newValue) => {
    customColor.value = newValue
  }
)
</script>
