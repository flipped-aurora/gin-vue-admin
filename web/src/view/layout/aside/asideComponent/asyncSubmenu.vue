<template>
  <el-sub-menu ref="subMenu" :index="routerInfo.name">
    <template #title>
      <i :class="'el-icon-'+routerInfo.meta.icon" />
      <span>{{ routerInfo.meta.title }}</span>
    </template>
    <slot />
  </el-sub-menu>
</template>

<script>
export default {
  name: 'AsyncSubmenu',
  props: {
    routerInfo: {
      default: function() {
        return null
      },
      type: Object
    }
  }
}
</script>
