<template>
  <svg
    :class="svgClass"
    v-bind="$attrs"
    :color="color"
  >
    <use
      :xlink:href="'#'+name"
      rel="external nofollow"
    />
  </svg>
</template>
<script setup>
import { computed } from 'vue'
const props = defineProps({
  name: {
    type: String,
    required: true
  },
  color: {
    type: String,
    default: 'currentColor'
  }
})

const svgClass = computed(() => {
  if (props.name) {
    return `svg-icon ${props.name}`
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
