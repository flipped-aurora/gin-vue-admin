<template>
  <div class="font-inter">
    <div class="bg-gray-50 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl p-8 shadow-sm">
      <div class="mb-8">
        <p class="text-base font-semibold text-gray-700 dark:text-gray-300 mb-5">精选色彩</p>
        <div class="grid grid-cols-3 gap-4">
          <div
            v-for="colorItem in presetColors"
            :key="colorItem.color"
            class="flex items-center gap-4 p-4 bg-white dark:bg-gray-700 border-2 border-gray-200 dark:border-gray-600 rounded-xl cursor-pointer transition-all duration-150 ease-in-out hover:transform hover:-translate-y-1 hover:shadow-lg"
            :class="{
              'ring-2 ring-offset-2 ring-offset-gray-50 dark:ring-offset-gray-800 transform -translate-y-1 shadow-lg': modelValue === colorItem.color
            }"
            :style="modelValue === colorItem.color ? {
              borderColor: colorItem.color,
              ringColor: colorItem.color + '40'
            } : {}"
            @click="handleColorChange(colorItem.color)"
          >
            <div
              class="relative w-10 h-10 rounded-lg border border-gray-300 dark:border-gray-500 flex-shrink-0 shadow-sm"
              :style="{ backgroundColor: colorItem.color }"
            >
              <div
                v-if="modelValue === colorItem.color"
                class="absolute inset-0 flex items-center justify-center text-white text-base"
                style="text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);"
              >
                <el-icon>
                  <Check />
                </el-icon>
              </div>
            </div>
            <div class="min-w-0 flex-1">
              <span class="block text-sm font-semibold text-gray-900 dark:text-white">{{ colorItem.name }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="flex items-center justify-between p-5 bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-xl mb-6 shadow-sm">
        <div class="flex-1">
          <h4 class="text-base font-semibold text-gray-900 dark:text-white">自定义颜色</h4>
          <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">选择任意颜色作为主题色</p>
        </div>
        <el-color-picker
          v-model="customColor"
          size="large"
          :predefine="presetColors.map(item => item.color)"
          @change="handleCustomColorChange"
          class="custom-color-picker"
        />
      </div>

      <div class="bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-xl p-5 shadow-sm">
        <div class="flex items-center justify-between">
          <span class="text-base font-semibold text-gray-700 dark:text-gray-300">当前主题色</span>
          <div class="flex items-center gap-3">
            <div
              class="w-6 h-6 rounded-lg border border-gray-300 dark:border-gray-500 shadow-sm"
              :style="{ backgroundColor: modelValue }"
            ></div>
            <code class="text-sm font-mono bg-gray-100 dark:bg-gray-600 text-gray-700 dark:text-gray-300 px-3 py-2 rounded-lg border border-gray-200 dark:border-gray-500">
              {{ modelValue }}
            </code>
          </div>
        </div>
      </div>
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
    default: '#3b82f6'
  }
})

const emit = defineEmits(['update:modelValue'])

const customColor = ref(props.modelValue)

const presetColors = [
  { color: '#4E80EE', name: '默认' },
  { color: '#8bb5d1', name: '晨雾蓝' },
  { color: '#a8c8a8', name: '薄荷绿' },
  { color: '#d4a5a5', name: '玫瑰粉' },
  { color: '#c8a8d8', name: '薰衣草' },
  { color: '#f0c674', name: '暖阳黄' },
  { color: '#b8b8b8', name: '月光银' },
  { color: '#d8a8a8', name: '珊瑚橙' },
  { color: '#a8d8d8', name: '海雾青' },
  { color: '#c8c8a8', name: '橄榄绿' },
  { color: '#d8c8a8', name: '奶茶棕' },
  { color: '#a8a8d8', name: '梦幻紫' },
  { color: '#c8d8a8', name: '抹茶绿' }
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

watch(() => props.modelValue, (newValue) => {
  customColor.value = newValue
})
</script>

<style scoped>
.font-inter {
  font-family: 'Inter', sans-serif;
}

.custom-color-picker {
  ::v-deep(.el-color-picker__trigger) {
    border: 1px solid #e5e7eb;
    border-radius: 6px;
    transition: all 150ms ease-in-out;

    &:hover {
      border-color: #9ca3af;
      transform: translateY(-1px);
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    }
  }
}

.dark .custom-color-picker {
  ::v-deep(.el-color-picker__trigger) {
    border-color: #4b5563;

    &:hover {
      border-color: #6b7280;
    }
  }
}
</style>
