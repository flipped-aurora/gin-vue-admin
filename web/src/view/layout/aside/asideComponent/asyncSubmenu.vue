<template>
  <el-sub-menu
    ref="subMenu"
    :index="routerInfo.name"
  >
    <template #title>
      <div
        v-if="!isCollapse"
        class="flex items-center "
        :style="{
          height : config.layout_side_item_height+ 'px',
        }"
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
import { useAppStore } from '@/pinia'
import { storeToRefs } from 'pinia'
const appStore = useAppStore()
const { config } = storeToRefs(appStore)

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
})

const isCollapse = inject('isCollapse', {
  default: false,
})

</script>

<style lang="scss">
.el-sub-menu__title{
  @apply h-14;
}
</style>
