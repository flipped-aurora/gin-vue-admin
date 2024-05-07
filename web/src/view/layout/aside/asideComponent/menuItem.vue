<template>
  <el-menu-item :index="routerInfo.name" :style="{
    height : config.layout_side_item_height+ 'px',
  }">
    <el-icon v-if="routerInfo.meta.icon">
      <component :is="routerInfo.meta.icon" />
    </el-icon>
    <template #title>
      {{ routerInfo.meta.title }}
    </template>
  </el-menu-item>
</template>


<script setup>
import { inject } from 'vue'
import { useAppStore } from '@/pinia'
import { storeToRefs } from 'pinia'
const appStore = useAppStore()
const { config } = storeToRefs(appStore)

defineOptions({
  name: 'MenuItem',
})

defineProps({
  routerInfo: {
    default: function() {
      return null
    },
    type: Object
  },
})

const isCollapse = inject('isCollapse', {
  default: false,
})

</script>

<style lang="scss">

</style>
