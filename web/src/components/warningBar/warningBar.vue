<template>
  <div
    class="px-1.5 py-2 flex items-center bg-amber-50 rounded gap-2 mb-3 text-amber-500"
    :class="href&&'cursor-pointer'"
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
