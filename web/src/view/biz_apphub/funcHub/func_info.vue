<script setup>
import {defineProps} from "vue";

const props = defineProps({
  func:{
    type: Object,
  }
});

function getFuncDefine() {
  let inp = []
  let outp = []
  // let func =props.func
  console.log("func.detail:", props.func.detail.param)

  if (!props.func.detail.param){
    return ""
  }
  props.func.detail.param.forEach(param=>{
    if (param.mode === "in") {
      inp.push(param.desc + " " + param.type)
    } else {
      outp.push(param.desc + " " + param.type)
    }
  })
  // for (let i = 0; i < props.func.detail.param.length; i++) {
  //   let item = props.func.detail.param[i]
  //   if (item.mode === "in") {
  //     inp.push(item.desc + " " + item.type)
  //   } else {
  //     outp.push(item.desc + " " + item.type)
  //   }
  // }
  return  props.func.detail.code_name + "(" + inp.join(",") + ")" + "->" + "(" + outp.join(",") + ")"
}

</script>

<template>
<div>
  <el-descriptions
      class="margin-top"
      title="函数详情"
      :column="3"
      border
  >
    <template #extra>
    </template>
    <el-descriptions-item>
      <template #label>
        <div class="cell-item">
          中文名
        </div>
      </template>
      {{func.detail.cn_name}}
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <div class="cell-item">
          函数名称
        </div>
      </template>
      {{func.detail.code_name}}
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <div class="cell-item">
          函数定义
        </div>
      </template>
      {{getFuncDefine()}}
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <div class="cell-item">
          归属
        </div>
      </template>
      <el-tag size="small">{{func.detail.classify}}</el-tag>
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <div class="cell-item">
          函数介绍
        </div>
      </template>
      {{func.detail.content}}
    </el-descriptions-item>
  </el-descriptions>
</div>
</template>

<style scoped lang="scss">

</style>
