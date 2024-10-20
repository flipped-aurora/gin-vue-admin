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

const paramType={
  "string":"文本",
  "number":"数值",
  "file":"文件",
  "table":"表格",
}

function getLabel(param){
  return param.desc+"("+paramType[param.type]+")"+":"
}

</script>

<template>
  <div style="padding: 24px 24px 24px 24px">

    <el-form-item v-for="(v,i) in func.detail.param" v-show="v.mode===props.type" :label="getLabel(v)" >
      <!--      {{v.note?v.note:""}}-->
      <!--      <el-input  v-if="v.mode===props.type && v.type==='file'" :label="v.desc" />-->

      <div v-if="v.mode===props.type">
        <pre>{{v.mock_data}}</pre>
      </div>
      <!--      <el-input v-model="v.value"  style="width: 80%"-->
      <!--                :autosize="{ minRows: 1, maxRows: 5 }"-->
      <!--                :placeholder="v.mock_data"-->
      <!--                show-word-limit-->
      <!--                type="textarea"/>-->
    </el-form-item>
  </div>
</template>

<style scoped lang="scss">

</style>
