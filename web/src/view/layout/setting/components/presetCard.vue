<template>
  <div
    class="relative flex flex-col gap-3 p-4 bg-white dark:bg-gray-700 border-2 border-gray-200 dark:border-gray-600 rounded-xl cursor-pointer transition-all duration-150 ease-in-out hover:-translate-y-0.5 hover:shadow-lg"
    @click="$emit('apply')"
  >
    <!-- 迷你预览：侧栏 + 主区 + 调色板 -->
    <div class="h-14 rounded-lg overflow-hidden flex border border-gray-200 dark:border-gray-600">
      <div class="w-1/4 h-full" :style="{ background: sideBg }"></div>
      <div class="flex-1 h-full p-2 flex flex-col gap-1" :style="{ background: mainBg }">
        <div class="h-2 w-2/3 rounded" :style="{ background: preset.config.primaryColor }"></div>
        <div class="flex gap-1 mt-1">
          <span class="w-3 h-3 rounded-full" :style="{ background: preset.config.primaryColor }"></span>
          <span class="w-3 h-3 rounded-full" :style="{ background: preset.config.successColor }"></span>
          <span class="w-3 h-3 rounded-full" :style="{ background: preset.config.warningColor }"></span>
          <span class="w-3 h-3 rounded-full" :style="{ background: preset.config.dangerColor }"></span>
        </div>
      </div>
    </div>
    <div class="flex items-center justify-between">
      <span class="text-sm font-semibold gva-theme-text-main truncate">{{ preset.name }}</span>
      <el-tag v-if="preset.builtin" size="small" type="info" effect="plain">内置</el-tag>
    </div>
    <el-icon
      v-if="!preset.builtin"
      class="absolute top-2 right-2 text-gray-400 hover:text-red-500"
      @click.stop="$emit('remove')"
    >
      <Delete />
    </el-icon>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Delete } from '@element-plus/icons-vue'

defineOptions({
  name: 'PresetCard'
})

const props = defineProps({
  preset: {
    type: Object,
    required: true
  }
})

defineEmits(['apply', 'remove'])

const sideBg = computed(() => {
  const t = props.preset.config?.menu_theme
  if (t === 'dark') return '#1e293b'
  if (t === 'light') return '#f8fafc'
  return '#ffffff'
})

const mainBg = computed(() =>
  props.preset.config?.darkMode === 'dark' ? '#0f172a' : '#f5f6f8'
)
</script>
