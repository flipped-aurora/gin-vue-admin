<template>
  <div>
    <warning-bar
        href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
        title="此功能为开发环境使用，不建议发布到生产，具体使用效果请点我观看。"
    />
    <div class="gva-search-box">
      <div class="text-lg mb-2 text-gray-600">
        使用AI创建<a
          class="text-blue-600 text-sm ml-4"
          href="https://plugin.gin-vue-admin.com/#/layout/userInfo/center"
          target="_blank"
      >获取AiPath</a
      >
      </div>
      <div class="relative">
        <el-input
            v-model="prompt"
            :maxlength="2000"
            :placeholder="`请大体描述您希望生成的页面`"
            :rows="5"
            resize="none"
            type="textarea"
            @blur="handleBlur"
            @focus="handleFocus"
        />
        <div class="flex absolute right-2 bottom-2">
          <el-tooltip effect="light">
            <template #content>
              <div>
                【完全免费】前往<a
                  class="text-blue-600"
                  href="https://plugin.gin-vue-admin.com/#/layout/userInfo/center"
                  target="_blank"
              >插件市场个人中心</a
              >申请AIPath，填入config.yaml的ai-path属性即可使用。
              </div>
            </template>
            <el-button
                type="primary"
                @click="llmAutoFunc()"
            >
              <el-icon size="18">
                <ai-gva/>
              </el-icon>
              生成
            </el-button>
          </el-tooltip>
        </div>
      </div>
    </div>
    <div>
      <div v-if="!outPut">
        <el-empty :image-size="200"/>
      </div>
      <div v-if="outPut">
        <div v-for="(snippet, index) in htmlFromLLM" :key="index" class="mb-6 p-4 border">
          <el-button type="primary" :icon="Upload" class="px-2 py-1" @click="copySnippet(snippet)" plain>复制</el-button>
          <div class="mt-2">
            <iframe-renderer :html="snippet" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  eye, createWeb
} from '@/api/autoCode'
import {ref} from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'

defineOptions({
  name: 'Picture'
})

import IframeRenderer from '@/view/systemTools/autoCode/component/iframeRenderer.vue'

const handleFocus = () => {
  document.addEventListener('keydown', handleKeydown);
}

const handleBlur = () => {
  document.removeEventListener('keydown', handleKeydown);
}

const handleKeydown = (event) => {
  if ((event.ctrlKey || event.metaKey) && event.key === 'Enter') {
    llmAutoFunc()
  }
}

const prompt = ref('')

// 判断是否返回的标志
const outPut = ref(false)
// 容纳llm返回的html
const htmlFromLLM = ref([])

const llmAutoFunc = async () => {
  const res = await createWeb({web: prompt.value, command: 'createWeb'})
  if (res.code === 0) {
    outPut.value = true
    // 使用mock数据模拟大模型返回值
    htmlFromLLM.value.push(res.data)
  }
}



</script>
