<template>
  <div style="height:80vh">
    <iframe width="100%" height="100%" :src="url" frameborder="0"></iframe>
  </div>
</template>

<script>
var path = process.env.VUE_APP_BASE_API
export default {
    name:"FormGenerator",
    data(){
      return{
        url:'http://127.0.0.1:' + buildConf.goServerPort + '/form-generator/#/',
        basePath:path
      }
    },
};
</script>

<style lang="scss">
</style>