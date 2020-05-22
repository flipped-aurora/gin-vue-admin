<template>
  <div style="height:80vh">
    <!-- <iframe width="10%" height="100%" src="http://127.0.0.1:8887/form-generator/#/" frameborder="0"></iframe> -->
    <iframe width="100%" height="100%" :src="basePath+'/form-generator/#/'" frameborder="0"></iframe>
  </div>
</template>

<script>
  var path = process.env.VUE_APP_BASE_API
  export default {
    name: "FormGenerator",
    data() {
      return {
        basePath: path
      }
    },
  };
</script>

<style lang="scss">
</style>