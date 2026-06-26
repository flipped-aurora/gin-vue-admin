<template>
  <el-menu-item
    :index="routerInfo.name"
    :style="{
          height: sideHeight
        }"
  >
    <el-icon v-if="routerInfo.meta.icon">
      <component :is="routerInfo.meta.icon" />
    </el-icon>
    <template v-else>
      {{ isCollapse ? routerInfo.meta.title[0] : "" }}
    </template>
    <template #title>
      {{ routerInfo.meta.title }}
    </template>
  </el-menu-item>
</template>

<script setup>
import {computed, inject} from 'vue'
  import { useThemeStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  const themeStore = useThemeStore()
  const { settings } = storeToRefs(themeStore)

  defineOptions({
    name: 'MenuItem'
  })

  defineProps({
    routerInfo: {
      default: function () {
        return null
      },
      type: Object
    }
  })

const isCollapse = inject('isCollapse', {
  default: false
})

  const sideHeight = computed(() => {
    return settings.value.layout.sideItemHeight + 'px'
  })
</script>

<style lang="scss"></style>
