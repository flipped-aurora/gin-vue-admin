<template>
  <el-sub-menu
    ref="subMenu"
    :index="routerInfo.name"
  >

    <template #title>
      <div
        v-if="!isCollapse"
      >
        <el-icon v-if="routerInfo.meta.icon">
          <component :is="routerInfo.meta.icon" />
        </el-icon>
        <span>{{ routerInfo.meta.title }}</span>
      </div>
      <template v-else>
        <el-icon v-if="routerInfo.meta.icon">
          <component :is="routerInfo.meta.icon" />
        </el-icon>
        <span>{{ routerInfo.meta.title }}</span>
      </template>
    </template>
    <slot />
  </el-sub-menu>
</template>

<script setup>
import { inject } from 'vue'

defineOptions({
  name: 'AsyncSubmenu',
})

 defineProps({
  routerInfo: {
    default: function() {
      return null
    },
    type: Object
  },

  theme: {
    default: function() {
      return {}
    },
    type: Object
  }
})

const isCollapse = inject('isCollapse', {
  default: false,
})

</script>

<style lang="scss">
.el-sub-menu__title{
  @apply h-12;
}
</style>
