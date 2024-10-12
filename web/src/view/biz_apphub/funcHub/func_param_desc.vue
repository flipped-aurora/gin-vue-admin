<script setup>
import {defineProps} from "vue";

const props = defineProps({
  func:{
    type: Object,
  },
  type:{
    type:String
  }
});
</script>

<template>
<div style="padding: 24px 24px 24px 24px">

  <el-form-item v-for="(v,i) in func.detail.param" v-show="v.mode===props.type" :label="v.desc" >
    {{v.note?v.note:""}}
  </el-form-item>
</div>
</template>

<style scoped lang="scss">

</style>
