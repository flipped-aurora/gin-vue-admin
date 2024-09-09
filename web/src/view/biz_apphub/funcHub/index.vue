<script setup>

import {onMounted, ref} from 'vue';
import {getBizCloudFunctionList} from "@/api/biz_apphub/bizCloudFunction";
const funcList =ref([])
import { useRouter } from 'vue-router';
const router=useRouter()
onMounted(() => {
  // 在这里编写组件挂载后要执行的逻辑
  getFuncList()
})



const getFuncList = async () => {
  let bizCloudFunctionList  = await getBizCloudFunctionList();
  funcList.value=bizCloudFunctionList.data.list
  console.log("funcList:",funcList.value)
}

function funcDetail(){
  router.push('/func_hub/detail');

}

</script>

<template>
<div>
  <div v-for="(v,i) in funcList">
      <a @click="funcDetail">{{v.cn_name}}</a>
  </div>
</div>
</template>

<style scoped lang="scss">

</style>