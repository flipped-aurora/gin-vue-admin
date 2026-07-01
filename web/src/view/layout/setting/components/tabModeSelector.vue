<template>
  <div class="grid grid-cols-3 gap-3 px-1">
    <div
      v-for="item in tabModes"
      :key="item.value"
      class="relative flex flex-col items-center gap-1.5 p-1.5 bg-white dark:bg-gray-700/50 border rounded-lg cursor-pointer transition-colors duration-150 ease-in-out hover:border-gray-300 dark:hover:border-gray-500"
      :class="modelValue === item.value ? '' : 'border-gray-200 dark:border-gray-600'"
      :style="
        modelValue === item.value
          ? { borderColor: primaryColor, boxShadow: `0 0 0 1px ${primaryColor}` }
          : {}
      "
      :title="item.label"
      @click="handleChange(item.value)"
    >
      <!-- 默认风格：扁平药丸 + 短竖线分隔，选中项主色弱底 -->
      <div
        v-if="item.value === 'button'"
        class="flex h-9 w-full items-center gap-1 overflow-hidden rounded bg-white px-1.5 dark:bg-gray-800"
      >
        <span class="h-4 w-9 rounded" :style="{ backgroundColor: primarySoft }" />
        <span class="h-3 w-px bg-gray-200 dark:bg-gray-600" />
        <span class="h-4 w-9 rounded bg-gray-100 dark:bg-gray-700" />
        <span class="h-3 w-px bg-gray-200 dark:bg-gray-600" />
        <span class="h-4 w-7 rounded bg-gray-100 dark:bg-gray-700" />
      </div>

      <!-- Chrome 风格：交叠标签，选中项抬升 + 主色弱底 -->
      <div
        v-else-if="item.value === 'chrome'"
        class="flex h-9 w-full items-end overflow-hidden rounded bg-gray-100 px-1 dark:bg-gray-800"
      >
        <span class="h-5 w-9 rounded-t bg-gray-200 dark:bg-gray-700" />
        <span
          class="-ml-1 h-6 w-9 rounded-t"
          :style="{ backgroundColor: primarySoft }"
        />
        <span class="-ml-1 h-5 w-9 rounded-t bg-gray-200 dark:bg-gray-700" />
      </div>

      <!-- Slider 风格：底部 2px 主色指示条 + 激活弱底 -->
      <div
        v-else-if="item.value === 'slider'"
        class="flex h-9 w-full items-end rounded bg-white px-1 dark:bg-gray-800"
      >
        <span class="mr-0.5 flex h-6 flex-1 items-center justify-center rounded-t-sm border-b-2 border-transparent bg-gray-50 dark:bg-gray-700">
          <span class="h-2 w-5 rounded-sm bg-gray-200 dark:bg-gray-600" />
        </span>
        <span
          class="mr-0.5 flex h-6 flex-1 items-center justify-center rounded-t-sm border-b-2"
          :style="{ borderColor: primaryColor, backgroundColor: primarySoft }"
        >
          <span class="h-2 w-5 rounded-sm" :style="{ backgroundColor: primaryColor }" />
        </span>
        <span class="flex h-6 flex-1 items-center justify-center rounded-t-sm border-b-2 border-transparent bg-gray-50 dark:bg-gray-700">
          <span class="h-2 w-4 rounded-sm bg-gray-200 dark:bg-gray-600" />
        </span>
      </div>

      <span
        class="text-xs"
        :style="modelValue === item.value ? { color: primaryColor } : {}"
      >
        {{ item.label }}
      </span>

      <svg-icon
        v-if="modelValue === item.value"
        icon="lucide:check"
        class="absolute top-1 right-1"
        :style="{ color: primaryColor }"
      />
    </div>
  </div>
</template>

<script setup>
  import { computed } from 'vue'
  import { storeToRefs } from 'pinia'
  import { useThemeStore } from '@/pinia'

  defineOptions({
    name: 'TabModeSelector'
  })

  defineProps({
    modelValue: {
      type: String,
      default: 'button'
    }
  })

  const emit = defineEmits(['update:modelValue'])

  const themeStore = useThemeStore()
  const { settings } = storeToRefs(themeStore)
  const primaryColor = computed(() => settings.value.themeColor)
  // 主色弱底（真实合成色，跟随换肤；预览只示意，不必严格等于实际标签底色）
  const primarySoft = 'rgb(var(--primary-color) / 0.15)'

  const tabModes = [
    { value: 'button', label: '默认' },
    { value: 'chrome', label: 'Chrome' },
    { value: 'slider', label: '指示条' }
  ]

  const handleChange = (value) => {
    emit('update:modelValue', value)
  }
</script>
