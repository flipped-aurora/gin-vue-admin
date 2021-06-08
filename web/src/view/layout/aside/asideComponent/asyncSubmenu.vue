<template>
  <el-submenu ref="subMenu" :popper-append-to-body="false" :index="routerInfo.name">
    <template slot="title">
      <i :class="'el-icon-'+routerInfo.meta.icon" />
      <span slot="title">{{ routerInfo.meta.title }}</span>
    </template>
    <slot />
  </el-submenu>
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
