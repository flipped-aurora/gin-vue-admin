<template>
  <div style="height:80vh">
    <iframe width="100%" height="100%" :src="`${basePath}:${basePort}/form-generator/#/`" frameborder="0" />
  </div>
</template>

<script>
var path = import.meta.env.VITE_BASE_PATH
var port = import.meta.env.VITE_SERVER_PORT
export default {
  name: 'FormGenerator',
  data() {
    return {
      basePath: path,
      basePort: port
    }
  }
}
</script>
