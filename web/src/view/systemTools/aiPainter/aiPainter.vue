<template>
<div>
  <div class="gva-search-box" >
    <div class="w-full relative">
      <el-input v-model="prompt" type="textarea" :rows="8" class="w-full"></el-input>
      <div class="absolute right-2 bottom-2">
        <el-button :loading="loading" @click="aiPainter()" type="primary"><ai-gva />{{html?'修改':'生成'}}</el-button>
        <el-button :loading="loading" v-if="html" @click="aiPainter(true)" type="primary"><ai-gva />重新生成</el-button>
        <el-button :loading="loading" v-if="html" @click="getCode" type="primary">复制代码</el-button>
      </div>
    </div>
  </div>
    <div class="w-full">
      <DynamicHtmlRenderer :htmlContent="html" />
    </div>
</div>
</template>

<script setup>
import {painter} from "@/api/autoCode";
import DynamicHtmlRenderer from "@/view/systemTools/aiPainter/components/dynamicHtmlRenderer.vue";
import { ref } from 'vue'
const loading = ref(false)
const prompt = ref('')
const html = ref(``)
const aiPainter = async (flag) => {
    const res = await painter({
        html: flag?'':html.value,
        prompt: prompt.value
    })
    html.value = res.data
}
const getCode = () => {
    navigator.clipboard.writeText(html.value).then(() => {
       Message.success('复制成功');
    }, () => {
      Message.error('复制失败');
    });
}
</script>
