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
            :placeholder="placeholder"
            :rows="8"
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
import { createWeb
} from '@/api/autoCode'
import {ref} from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage } from 'element-plus'


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

// 复制方法：把某个字符串写进剪贴板
const copySnippet = (htmlString) => {
  navigator.clipboard.writeText(htmlString)
      .then(() => {
        ElMessage({
          message: '复制成功',
          type: 'success',
        })
      })
      .catch(err => {
        ElMessage({
          message: '复制失败',
          type: 'warning',
        })
      })
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

const placeholder = ref(`"✨ 请详细描述您想要的页面，例如：
• 页面用途（企业官网/电商页面/个人博客等）
• 需要包含的主要内容板块
• 偏好的风格（简约/科技感/温馨/专业等）
• 需要特别强调的元素
• 参考网站或配色建议

示例：'需要一个科技公司的产品介绍页，包含banner轮播图、三栏功能特点展示、客户案例模块，喜欢深蓝色调，参考苹果官网的简洁风格'"`)



</script>
