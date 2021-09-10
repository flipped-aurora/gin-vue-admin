<template>
  <el-menu-item :index="routerInfo.name" :route="{parameters:routerInfo.parameters}">
    <i :class="'el-icon-'+routerInfo.meta.icon" />
    <template #title>
      <span>{{ routerInfo.meta.title }}</span>
    </template>
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
