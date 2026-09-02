<template>
  <div
    class="px-1.5 py-2 flex items-center rounded-sm mt-2 bg-amber-50 gap-2 mb-3 text-amber-500 dark:bg-amber-700 dark:text-gray-200"
    :class="href && 'cursor-pointer'"
    @click="open"
  >
    <el-icon class="text-xl">
      <warning-filled />
    </el-icon>
    <span>
      {{ title }}
    </span>
  </div>
</template>
<script setup>
  import { WarningFilled } from '@element-plus/icons-vue'
  const prop = defineProps({
    title: {
      type: String,
      default: ''
    },
    href: {
      type: String,
      default: ''
    }
  })

  const open = () => {
    if (prop.href) {
      window.open(prop.href)
    }
  }
</script>
