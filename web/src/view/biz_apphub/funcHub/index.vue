<script setup>
import { ThumbUpIcon, ChatIcon, ShareIcon, HeartIcon, MoreIcon, UserIcon } from 'tdesign-icons-vue-next';
import { MessagePlugin } from 'tdesign-vue-next';
import {onMounted, ref} from 'vue';
import {getBizCloudFunctionList} from "@/api/biz_apphub/bizCloudFunction";
const funcList =ref([])


import { useRouter } from 'vue-router';
const router=useRouter()
onMounted(() => {
  // 在这里编写组件挂载后要执行的逻辑
  getFuncList()
})


const cover = 'https://tdesign.gtimg.com/site/source/card-demo.png';
const getFuncList = async () => {
  let bizCloudFunctionList  = await getBizCloudFunctionList();
  funcList.value=bizCloudFunctionList.data.list
  console.log("funcList:",funcList.value)
}
const options = [
  {
    content: '操作一',
    value: 1,
  },
  {
    content: '操作二',
    value: 2,
  },
];

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
<div>
  <div v-for="(v,i) in funcList">
<!--      <a target="_blank" @click="funcDetail(v)">{{v.cn_name}}</a>-->


    <t-card @click="funcDetail(v)" :cover="cover" theme="poster2" :style="{ width: '250px' }">
      <template #footer>
        <t-button variant="text" shape="square" :style="{ 'margin-right': '8px' }">
          <heart-icon />
        </t-button>
        <span>
          {{v.cn_name}}
        </span>
<!--        <t-button variant="text" shape="square" :style="{ 'margin-right': '8px' }">-->
<!--          <chat-icon />-->
<!--        </t-button>-->
        <t-button variant="text" shape="square">
          <share-icon />
        </t-button>
      </template>
      <template #actions>
        <t-dropdown :options="options" :min-column-width="112" @click="clickHandler">
          <t-button variant="text" shape="square">
            <more-icon />
          </t-button>
        </t-dropdown>
      </template>
    </t-card>

  </div>
</div>
</template>

<style scoped lang="scss">

</style>