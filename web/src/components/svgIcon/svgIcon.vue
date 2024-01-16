<template>
  <svg
    :class="svgClass"
    v-bind="$attrs"
    :style="{ color: color }"
  >
    <use
      :xlink:href="iconName"
      rel="external nofollow"
    />
  </svg>
</template>
<script setup>
import { defineProps, computed } from 'vue'
const props = defineProps({
  name: {
    type: String,
    required: true
  },
  color: {
    type: String,
    default: ''
  }
})
const iconName = computed(() => `#icon-${props.name}`)
const svgClass = computed(() => {
  if (props.name) {
    return `svg-icon ${iconName.value}`
  }
  return 'svg-icon'
})
</script>
<style scoped>
.svg-icon {
  @apply w-4 h-4;
  fill: currentColor;
  vertical-align: middle;
}
</style>
