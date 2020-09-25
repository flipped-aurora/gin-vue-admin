<template>
  <el-submenu :popper-append-to-body="false" ref="subMenu" :index="routerInfo.name">
    <template slot="title">
      <i :class="'el-icon-'+routerInfo.meta.icon"></i>
      <span slot="title">{{routerInfo.meta.title}}</span>
    </template>
    <slot></slot>
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
<style lang="scss">
</style>