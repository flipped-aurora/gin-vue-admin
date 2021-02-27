<template>
<!-- 此路由可作为父类路由通用路由页面使用 如需自定义父类路由页面 请参考 @/view/superAdmin/index.vue -->
  <div>
    <keep-alive>
      <router-view v-if="$route.meta.keepAlive"></router-view>
    </keep-alive>
    <router-view v-if="!$route.meta.keepAlive"></router-view>
  </div>
</template>

<script>
export default {
  name: "RouterHolder",
};
</script>
<style lang="scss"></style>
