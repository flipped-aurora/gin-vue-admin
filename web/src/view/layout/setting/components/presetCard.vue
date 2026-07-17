<template>
  <div
    class="relative flex flex-col gap-3 p-4 bg-container border-2 border-border rounded-xl cursor-pointer transition-all duration-150 ease-in-out hover:-translate-y-0.5 hover:shadow-lg"
    :class="{ 'opacity-60': !compatible }"
    @click="$emit('apply')"
  >
    <!-- 迷你预览：侧栏 + 主区 + 调色板 -->
    <div class="h-14 rounded-lg overflow-hidden flex border border-border">
      <div class="w-1/4 h-full" :style="{ background: sideBg }"></div>
      <div class="flex-1 h-full p-2 flex flex-col gap-1" :style="{ background: mainBg }">
        <div class="h-2 w-2/3 rounded" :style="{ background: preview.primaryColor }"></div>
        <div class="flex gap-1 mt-1">
          <span class="w-3 h-3 rounded-full" :style="{ background: preview.primaryColor }"></span>
          <span class="w-3 h-3 rounded-full" :style="{ background: preview.successColor }"></span>
          <span class="w-3 h-3 rounded-full" :style="{ background: preview.warningColor }"></span>
          <span class="w-3 h-3 rounded-full" :style="{ background: preview.dangerColor }"></span>
        </div>
      </div>
    </div>
    <div class="flex items-center justify-between gap-2">
      <span class="text-sm font-semibold gva-theme-text-main truncate">{{ preset.name }}</span>
      <el-tag
        v-if="!compatible"
        size="small"
        type="warning"
        effect="plain"
        :title="`需要 GVA ≥ ${preset.minMainVersion}`"
      >
        需 v{{ preset.minMainVersion }}+
      </el-tag>
      <el-tag v-else-if="preset.builtin" size="small" type="info" effect="plain">内置</el-tag>
    </div>
    <button
      v-if="!preset.builtin"
      type="button"
      aria-label="删除预设"
      class="absolute top-2 right-2 p-0 m-0 bg-transparent border-0 cursor-pointer leading-none text-muted-foreground hover:text-red-500 [font:inherit]"
      @click.stop="$emit('remove')"
    >
      <el-icon>
        <Delete />
      </el-icon>
    </button>
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
  },
  compatible: {
    type: Boolean,
    default: true
  }
})

defineEmits(['apply', 'remove'])

const preview = computed(() => {
  const theme = props.preset.theme || {}
  return {
    primaryColor: theme.themeColor,
    successColor: theme.otherColor?.success,
    warningColor: theme.otherColor?.warning,
    dangerColor: theme.otherColor?.error,
    menuTheme: theme.menu?.theme,
    themeScheme: theme.themeScheme
  }
})

const sideBg = computed(() => {
  const t = preview.value.menuTheme
  if (t === 'dark') return '#1e293b'
  if (t === 'light') return '#f8fafc'
  return '#ffffff'
})

const mainBg = computed(() =>
  preview.value.themeScheme === 'dark' ? '#0f172a' : '#f5f6f8'
)
</script>
