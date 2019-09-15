<template>
  <el-menu-item :index="routerInfo.name">
    <i class="el-icon-document"></i>
    <span slot="title">{{routerInfo.meta.title}}</span>
  </el-menu-item>
</template>

<script>
export default {
  name: 'MenuItem',
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