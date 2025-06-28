<template>
  <div class="flex justify-center">
    <div class="inline-flex bg-gray-100 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-lg p-1 gap-1">
      <div
        v-for="mode in themeModes"
        :key="mode.value"
        class="flex flex-col items-center justify-center px-4 py-3 rounded-md cursor-pointer transition-all duration-150 ease-in-out min-w-[64px]"
        :class="[
          modelValue === mode.value
            ? 'text-white shadow-sm transform -translate-y-0.5'
            : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200 hover:bg-gray-50 dark:hover:bg-gray-700'
        ]"
        :style="modelValue === mode.value ? { backgroundColor: primaryColor } : {}"
        @click="handleModeChange(mode.value)"
      >
        <el-icon class="text-lg mb-1">
          <component :is="mode.icon" />
        </el-icon>
        <span class="text-xs font-medium">{{ mode.label }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { Sunny, Moon, Monitor } from '@element-plus/icons-vue'
import { useAppStore } from '@/pinia'

defineOptions({
  name: 'ThemeModeSelector'
})

const props = defineProps({
  modelValue: {
    type: String,
    default: 'auto'
  }
})

const emit = defineEmits(['update:modelValue'])

const appStore = useAppStore()
const { config } = storeToRefs(appStore)

const primaryColor = computed(() => config.value.primaryColor)

const themeModes = [
  {
    value: 'light',
    label: '浅色',
    icon: Sunny
  },
  {
    value: 'dark',
    label: '深色',
    icon: Moon
  },
  {
    value: 'auto',
    label: '跟随系统',
    icon: Monitor
  }
]

const handleModeChange = (mode) => {
  emit('update:modelValue', mode)
}
</script>
