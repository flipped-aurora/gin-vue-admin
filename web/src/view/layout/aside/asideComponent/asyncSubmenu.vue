<template>
  <el-sub-menu
    ref="subMenu"
    :index="routerInfo.name"
    :style="menuStyleVars"
    class="gva-sub-menu dark:text-slate-300 relative"
  >
    <template #title>
      <div
        v-if="!isCollapse"
        class="flex items-center"
        :style="{
          height: sideHeight
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
  import { inject, computed } from 'vue'
  import { useThemeStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  const themeStore = useThemeStore()
  const { settings } = storeToRefs(themeStore)

  defineOptions({
    name: 'AsyncSubmenu'
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

  const menuStyleVars = computed(() => ({
    '--gva-side-height': sideHeight.value
  }))
</script>

<style lang="scss">
  .gva-sub-menu {
    .el-sub-menu__title {
      height: var(--gva-side-height) !important;
    }
  }
</style>
