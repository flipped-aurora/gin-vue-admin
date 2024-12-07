<template>
  <div @click="clickFull">
    <div v-if="isShow" class="gvaIcon gvaIcon-fullscreen-expand" />
    <div v-else class="gvaIcon gvaIcon-fullscreen-shrink" />
  </div>
</template>

<script setup>
  import screenfull from 'screenfull' // 引入screenfull
  import { onMounted, onUnmounted, ref } from 'vue'

  defineOptions({
    name: 'Screenfull'
  })

  defineProps({
    width: {
      type: Number,
      default: 22
    },
    height: {
      type: Number,
      default: 22
    },
    fill: {
      type: String,
      default: '#48576a'
    }
  })

  onMounted(() => {
    if (screenfull.isEnabled) {
      screenfull.on('change', changeFullShow)
    }
  })

  onUnmounted(() => {
    screenfull.off('change')
  })

  const clickFull = () => {
    if (screenfull.isEnabled) {
      screenfull.toggle()
    }
  }

  const isShow = ref(true)
  const changeFullShow = () => {
    isShow.value = !screenfull.isFullscreen
  }
</script>

<style scoped lang="scss">
  .screenfull-svg {
    width: 16px;
    height: 16px;
    cursor: pointer;
    vertical-align: middle;
    margin-right: 32px;
    fill: rgba(0, 0, 0, 0.45);
  }
</style>
