<script setup>
import {onMounted, reactive,ref} from 'vue';
import { useRoute } from 'vue-router';
import {findBizCloudFunction} from "@/api/biz_apphub/bizCloudFunction";
import axios from "axios";
const route=useRoute()
const func =reactive({detail:{}})
const funcDefine=ref("")

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

</script>

<template>

  <div>

    <el-descriptions
        class="margin-top"
        title="With border"
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
        {{funcDefine}}
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
    <h2>
      <div style="">函数名称：{{func.detail.cn_name}}</div>
    </h2>
    <t-form ref="form" reset-type="initial" style="max-width: 100%">
      <el-tabs type="border-card">
        <el-tab-pane label="运行">
          <div style="padding: 24px 24px 24px 0">
            <t-form-item v-for="(v,i) in func.detail.param" v-show="v.mode==='in'" :label="v.desc" :label-width="150">
              <el-input v-model="v.value" v-if="v.input_mode==='text_field'" style="width: 80%"
                  :autosize="{ minRows: 3, maxRows: 30 }"
                  :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"
                  show-word-limit
                  type="textarea"
              />
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
              <t-input v-else v-model="v.value" :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"></t-input>

            </t-form-item>
          </div>

        </el-tab-pane>
        <el-tab-pane label="案例">Config</el-tab-pane>
      </el-tabs>


    </t-form>


  </div>
</template>

<style scoped lang="scss">

</style>