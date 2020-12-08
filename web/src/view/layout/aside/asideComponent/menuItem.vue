<template>
  <el-menu-item  :index="routerInfo.name" :route="{parameters:routerInfo.parameters}">
    <div class="el-menu-item-div">
      <i :class="'el-icon-'+routerInfo.meta.icon"></i>
      <span slot="title">{{routerInfo.meta.title}}</span>
    </div>
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