<script setup>




import {defineProps} from "vue";

const props = defineProps({
  content:{
    type: String,
    default:()=>""
  }
});

</script>

<template>
<div>
  <el-scrollbar height="98%">
    <div class="code-container">
      <pre id="code-pre" class="code-pre">{{props.content}}</pre>
    </div>

  </el-scrollbar>
</div>
</template>

<style scoped lang="scss">

</style>
