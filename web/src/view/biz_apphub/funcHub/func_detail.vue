<script setup>
import {onMounted, reactive,ref} from 'vue';
import { useRoute } from 'vue-router';
import {findBizCloudFunction} from "@/api/biz_apphub/bizCloudFunction";
import axios from "axios";
const route=useRoute()
const func =reactive({detail:{}})
const funcDefine=ref("")
// const userId = route.state.aaa;

// const state=reactive({
//   detail:detail,
// })
// console.log(state.detail)


const funcCall=async ()=>{
  let bd=reactive({})
  for (let i = 0; i < func.detail.param.length; i++) {
    if (func.detail.param[i].mode==="in"){
      bd[func.detail.param[i].code]= func.detail.param[i].value
    }
  }
  if (func.detail.api_config.method==="post"){
    let res =await axios.post(func.detail.api_config.path,bd)
    if(res.data.code===0){
      for (const key in res.data.data) {
        if (res.data.data.hasOwnProperty(key)) {
          const value = res.data.data[key];
          for (let i = 0; i < func.detail.param.length; i++) {
            console.log("func.detail.param[i]",func.detail.param[i])
            if (func.detail.param[i].code===key){
              func.detail.param[i].value=value
            }
          }
        }
      }
    }else {
      
    }
  }else {
    
  }
}

onMounted(() => {

  let fn = async ()=>{
    const res=await findBizCloudFunction({ID: route.query.id})
    if (res.code===0){
      func.detail=res.data
    }
    getFuncInfo()
  }
  fn()
})
function getFuncInfo(){
  let inp=[]
  let outp=[]
  console.log("func.detail:",func.detail)
  for (let i = 0; i < func.detail.param.length; i++) {
    let item=func.detail.param[i]
    if (item.mode==="in"){
      inp.push(item.desc+" " + item.type)
    }else{
      outp.push(item.desc+" " + item.type)
    }
  }
  funcDefine.value= func.detail.code_name+"("+inp.join(",")+")"+"->"+"("+outp.join(",")+")"
}
// function getFuncInfo() {
//   // return func.detail.code_name+"("
// }
let id = 0;
const getId = () => {
  return ++id;
};
const studentTab = ref(1);
const formData = reactive({
  school: 1,
  students: [
    {
      id: getId(),
      label: '运行',
      name: 'run',
      courseType: 'wenke',
      course: ['1'],
    },
    {
      id: getId(),
      label: '示例',
      name: 'demo',
      courseType: 'wenke',
      course: [],
    },
  ],
});
</script>

<template>
<!--<div>测试：{{detail.cn_name}}</div>-->

  <div>
    <t-descriptions bordered  :title="func.detail.title">
      <t-descriptions-item label="中文名">{{func.detail.cn_name}}</t-descriptions-item>
      <t-descriptions-item label="函数名称">{{func.detail.code_name}}</t-descriptions-item>
      <t-descriptions-item label="函数定义"
      >{{funcDefine}}</t-descriptions-item
      >
      <t-descriptions-item label="归属"
      >{{func.detail.classify}}</t-descriptions-item
      >
    </t-descriptions>
    <h2>
      函数名称：{{func.detail.cn_name}}
    </h2>
    <t-form
        ref="form"
        :data="formData"
        reset-type="initial"
        style="max-width: 100%"
        @reset="onReset"
        @submit="onSubmit"
    >

      <t-tabs
          v-model="studentTab"
          theme="card"
          addable
          style="margin-left: 30px; border: 1px solid var(--td-component-stroke)"
          @add="onAddStudent"
      >
        <t-tab-panel
            v-for="(student, index) in formData.students"
            :key="student.id"
            :value="student.id"
            :label="student.label"
            :destroy-on-hide="false"
        >
          <div style="padding: 24px 24px 24px 0">
            <!-- 重点阅读：数组里面，注意 name 定义，用于区分不同的字段 -->
            <!-- 重点阅读：name 表示当前字段在 formData 中的路径，ruleName 表示当前字段在 rules 中的名称 -->
<!--            <t-form-item label="姓名" :name="`students[${index}].name`" :label-width="80">-->
<!--              <t-input v-model="formData.students[index].name" placeholder="请输入内容"></t-input>-->
<!--            </t-form-item>-->

            <t-form-item v-for="(v,i) in func.detail.param" v-show="v.mode==='in'" :label="v.desc" :label-width="150">
              <el-input
                  v-model="v.value"
                  v-if="v.input_mode==='text_field'"

                  style="width: 80%"
                  :autosize="{ minRows: 3, maxRows: 30 }"
                  :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"
                  show-word-limit
                  type="textarea"
              />
<!--              <t-textarea  allowInputOverMax autosize v-if="v.input_mode==='text_field'" v-model="v.value" :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"></t-textarea>-->
              <t-input v-else v-model="v.value" :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"></t-input>

            </t-form-item>

            <t-button @click="funcCall" style="margin: 10px" block theme="primary" variant="base">运行</t-button>

            <t-form-item v-for="(v,i) in func.detail.param" v-show="v.mode==='out'" :label="v.desc" :label-width="150">
              <el-input
                  v-model="v.value"
                  v-if="v.input_mode==='text_field'"
                  style="width: 80%"
                  :autosize="{ minRows: 3, maxRows: 30 }"

                  :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"
                  show-word-limit
                  type="textarea"
              />
<!--              <t-textarea  allow-input-over-max  autosize v-if="v.input_mode==='text_field'"  v-model="v.value" :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"></t-textarea>-->
              <t-input v-else v-model="v.value" :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"></t-input>

            </t-form-item>
          </div>
        </t-tab-panel>

      </t-tabs>

    </t-form>


  </div>
</template>

<style scoped lang="scss">

</style>