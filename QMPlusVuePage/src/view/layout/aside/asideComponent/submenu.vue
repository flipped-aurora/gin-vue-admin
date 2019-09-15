<template>
  <el-submenu :index="routerInfo.name">
    <template slot="title">
      <i class="el-icon-location"></i>
      <span slot="title">{{routerInfo.meta.title}}</span>
    </template>
    <slot></slot>
  </el-submenu>
</template>

<script>
export default {
  name: 'SubMenu',
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