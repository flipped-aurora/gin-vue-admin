<script setup>
import { ThumbUpIcon, ChatIcon, ShareIcon, HeartIcon, MoreIcon, UserIcon } from 'tdesign-icons-vue-next';
import { MessagePlugin } from 'tdesign-vue-next';
import {onMounted, ref} from 'vue';
import {getBizCloudFunctionList} from "@/api/biz_apphub/bizCloudFunction";
const funcList =ref([])
const functionsList =ref([]) //函数列表二维数组，控制每行显示


import { useRouter } from 'vue-router';
const router=useRouter()
onMounted(() => {
  // 在这里编写组件挂载后要执行的逻辑
  getFuncList()
})

function chunkArray(array, n) {
  let result = [];
  for (let i = 0; i < array.length; i += n) {
    result.push(array.slice(i, i + n));
  }
  return result;
}
const cover = 'https://tdesign.gtimg.com/site/source/card-demo.png';
const getFuncList = async () => {
  let bizCloudFunctionList  = await getBizCloudFunctionList();
  if (bizCloudFunctionList.code!==0){
    MessagePlugin.error(bizCloudFunctionList.msg);
    return
  }
  functionsList.value=[]
  funcList.value=bizCloudFunctionList.data.list
  let list=bizCloudFunctionList.data.list
  functionsList.value=chunkArray(list,6)
  // console.log("funcList:",funcList.value)
  console.log("functionsList",functionsList.value)
}
const options = [
  {
    content: '操作一',
    value: 1
  },
  {
    content: '操作二',
    value: 2,
  },
];

function getCovers(func){
  console.log("getCovers",func.covers.split(";"))
  return func.covers.split(";")
}
function getFuncInfo(func) {
  let inp = []
  let outp = []
  console.log("func:", func)
  for (let i = 0; i < func.param.length; i++) {
    let item = func.param[i]
    if (item.mode === "in") {
      inp.push(item.desc + " " + item.type)
    } else {
      outp.push(item.desc + " " + item.type)
    }
  }
  return  func.code_name + "(" + inp.join(",") + ")" + "->" + "(" + outp.join(",") + ")"
}

function getFuncInMockData(func) {
  let desc={}
  for (let i = 0; i < func.param.length; i++) {
    let item = func.param[i]
    if (item.mode === "in") {
      desc[item.desc] =  item.mock_data
    }
  }
  return  desc
}
function getFuncOutMockData(func) {
  // let inp = []
  let desc={}
  for (let i = 0; i < func.param.length; i++) {
    let item = func.param[i]
    if (item.mode === "out") {
      desc[item.desc] =  item.mock_data
    }
  }
  return  desc
}

function getMockData(func){
  console.log("getCovers",func.covers.split(";"))
  return func.covers.split(";")
}

const clickHandler = (data) => {
  MessagePlugin.success(`选中【${data.content}】`);
};
function funcDetail(v){
  console.log("vvvvvv:",v)
  router.push({
    path:'/func_hub/detail',
    query:{
      id:v.ID
    }
  });
}

</script>

<template>
<!--<div style="display: flex;justify-content: left;align-items: flex-start">-->
<!--  <div v-for="(v,i) in funcList">-->
<!--      <a target="_blank" @click="funcDetail(v)">{{v.cn_name}}</a>-->

<p class="function-line" v-for="(functions,i) in functionsList">

  <div style="display: flex;justify-content: left;align-items: flex-start">
  <el-card shadow="hover" :body-style="{padding:'0px',width:'300px'}" v-for="(v,i1) in functions" @click="funcDetail(v)" style="max-width: 480px;margin: 15px">
    <template #header>
      <div  class="card-header">
        <!--        <span>{{v.cn_name}}</span>-->
        <h4>{{v.cn_name}}</h4>
      </div>
    </template>
    <div >
<!--      <p v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</p>-->

      <el-carousel interval="3000" >
        <el-carousel-item v-if="v.covers!==''" v-for="(img,ii) in getCovers(v)" :key="img">
<!--          <h3 class="small justify-center" text="2xl">{{ item }}</h3>-->
          <img style="height: 300px" :src="img">
          <
        </el-carousel-item>
        <div v-else>
<!--          <h3 class="small justify-center" text="2xl">{{ getFuncInfo(v) }}</h3>-->
          <h2>输入参数</h2>
<!--          {{getFuncInMockData(v)}}-->
          <p v-for="(val,key) in getFuncInMockData(v)">
            <span>{{key}}: </span>
            <el-text type="success">{{val}} </el-text>
          </p>

          <h2>输出参数</h2>
          <p v-for="(val,key) in getFuncOutMockData(v)">
            <span>{{key}}: </span>
            <el-text type="success">{{val}} </el-text>
          </p>
        </div>
      </el-carousel>
    </div>
    <template #footer>
      <div >
<!--        <h4>{{v.cn_name}}</h4>-->
        <p>{{v.code_name}}</p>
      </div>
    </template>
  </el-card>
  </div>
</p>

<!--    <t-card v-for="(v,i) in funcList" @click="funcDetail(v)" :cover="cover" theme="poster2" style="width:250px;margin: 10px ">-->
<!--      <template #footer>-->
<!--        <t-button variant="text" shape="square" :style="{ 'margin-right': '8px' }">-->
<!--          <heart-icon />-->
<!--        </t-button>-->
<!--        <span>-->
<!--          {{v.cn_name}}-->
<!--        </span>-->
<!--&lt;!&ndash;        <t-button variant="text" shape="square" :style="{ 'margin-right': '8px' }">&ndash;&gt;-->
<!--&lt;!&ndash;          <chat-icon />&ndash;&gt;-->
<!--&lt;!&ndash;        </t-button>&ndash;&gt;-->
<!--        <t-button variant="text" shape="square">-->
<!--          <share-icon />-->
<!--        </t-button>-->
<!--      </template>-->
<!--      <template #actions>-->
<!--        <t-dropdown :options="options" :min-column-width="112" @click="clickHandler">-->
<!--          <t-button variant="text" shape="square">-->
<!--            <more-icon />-->
<!--          </t-button>-->
<!--        </t-dropdown>-->
<!--      </template>-->
<!--    </t-card>-->

<!--  </div>-->
<!--</div>-->
</template>

<style scoped lang="scss">

</style>

<style>
.demonstration {
  color: var(--el-text-color-secondary);
}

.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 150px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>
