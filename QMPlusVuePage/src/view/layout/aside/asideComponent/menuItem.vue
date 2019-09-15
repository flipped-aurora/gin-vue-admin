<template>
    <el-menu-item :index="routerInfo.name">{{routerInfo.meta.title}}</el-menu-item>
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
  },
}
</script>
<style lang="scss">
</style>