<template>
  <el-menu-item :index="routerInfo.name" class="h-12">
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
