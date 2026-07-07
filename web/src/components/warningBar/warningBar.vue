<template>
  <div
    class="gva-warning-bar px-1.5 py-2 flex items-center my-2 bg-amber-50 gap-2 text-amber-500 dark:bg-amber-700 dark:text-gray-200"
    :class="href && 'cursor-pointer'"
    @click="open"
  >
    <el-icon>
      <warning-filled />
    </el-icon>
    <span class="text-sm">
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
