<template>
  <div>
    <warning-bar
        href="https://www.gin-vue-admin.com/empower/"
        title="此功能只针对授权用户开放，点我【购买授权】"
    />
    <div class="gva-search-box">
      <div class="text-xl mb-2 text-gray-600">
        AI前端工程师<a
          class="text-blue-600 text-sm ml-4"
          href="https://plugin.gin-vue-admin.com/#/layout/userInfo/center"
          target="_blank"
      >获取AiPath</a
      >
      </div>
      
      <!-- 选项模式 -->
      <div class="mb-4">
        <div class="mb-3">
          <div class="text-base font-medium mb-2">页面用途</div>
          <el-radio-group v-model="pageType" class="mb-2" @change="handlePageTypeChange">
            <el-radio label="企业官网">企业官网</el-radio>
            <el-radio label="电商页面">电商页面</el-radio>
            <el-radio label="个人博客">个人博客</el-radio>
            <el-radio label="产品介绍">产品介绍</el-radio>
            <el-radio label="活动落地页">活动落地页</el-radio>
            <el-radio label="其他">其他</el-radio>
          </el-radio-group>
          <el-input v-if="pageType === '其他'" v-model="pageTypeCustom" placeholder="请输入页面用途" class="w-full" />
        </div>
        
        <div class="mb-3">
          <div class="text-base font-medium mb-2">主要内容板块</div>
          <el-checkbox-group v-model="contentBlocks" class="flex flex-wrap gap-2 mb-2">
            <el-checkbox label="Banner轮播图">Banner轮播图</el-checkbox>
            <el-checkbox label="产品/服务介绍">产品/服务介绍</el-checkbox>
            <el-checkbox label="功能特点展示">功能特点展示</el-checkbox>
            <el-checkbox label="客户案例">客户案例</el-checkbox>
            <el-checkbox label="团队介绍">团队介绍</el-checkbox>
            <el-checkbox label="联系表单">联系表单</el-checkbox>
            <el-checkbox label="新闻/博客列表">新闻/博客列表</el-checkbox>
            <el-checkbox label="价格表">价格表</el-checkbox>
            <el-checkbox label="FAQ/常见问题">FAQ/常见问题</el-checkbox>
            <el-checkbox label="用户评价">用户评价</el-checkbox>
            <el-checkbox label="数据统计">数据统计</el-checkbox>
            <el-checkbox label="商品列表">商品列表</el-checkbox>
            <el-checkbox label="商品卡片">商品卡片</el-checkbox>
            <el-checkbox label="购物车">购物车</el-checkbox>
            <el-checkbox label="结算页面">结算页面</el-checkbox>
            <el-checkbox label="订单跟踪">订单跟踪</el-checkbox>
            <el-checkbox label="商品分类">商品分类</el-checkbox>
            <el-checkbox label="热门推荐">热门推荐</el-checkbox>
            <el-checkbox label="限时特惠">限时特惠</el-checkbox>
            <el-checkbox label="其他">其他</el-checkbox>
          </el-checkbox-group>
          <el-input v-if="contentBlocks.includes('其他')" v-model="contentBlocksCustom" placeholder="请输入其他内容板块" class="w-full" />
        </div>
        
        <div class="mb-3">
          <div class="text-base font-medium mb-2">风格偏好</div>
          <el-radio-group v-model="stylePreference" class="mb-2">
            <el-radio label="简约">简约</el-radio>
            <el-radio label="科技感">科技感</el-radio>
            <el-radio label="温馨">温馨</el-radio>
            <el-radio label="专业">专业</el-radio>
            <el-radio label="创意">创意</el-radio>
            <el-radio label="复古">复古</el-radio>
            <el-radio label="奢华">奢华</el-radio>
            <el-radio label="其他">其他</el-radio>
          </el-radio-group>
          <el-input v-if="stylePreference === '其他'" v-model="stylePreferenceCustom" placeholder="请输入风格偏好" class="w-full" />
        </div>
        
        <div class="mb-3">
          <div class="text-base font-medium mb-2">设计布局</div>
          <el-radio-group v-model="layoutDesign" class="mb-2">
            <el-radio label="单栏布局">单栏布局</el-radio>
            <el-radio label="双栏布局">双栏布局</el-radio>
            <el-radio label="三栏布局">三栏布局</el-radio>
            <el-radio label="网格布局">网格布局</el-radio>
            <el-radio label="画廊布局">画廊布局</el-radio>
            <el-radio label="瀑布流">瀑布流</el-radio>
            <el-radio label="卡片式">卡片式</el-radio>
            <el-radio label="侧边栏+内容布局">侧边栏+内容布局</el-radio>
            <el-radio label="分屏布局">分屏布局</el-radio>
            <el-radio label="全屏滚动布局">全屏滚动布局</el-radio>
            <el-radio label="混合布局">混合布局</el-radio>
            <el-radio label="响应式">响应式</el-radio>
            <el-radio label="其他">其他</el-radio>
          </el-radio-group>
          <el-input v-if="layoutDesign === '其他'" v-model="layoutDesignCustom" placeholder="请输入设计布局" class="w-full" />
        </div>
        
        <div class="mb-3">
          <div class="text-base font-medium mb-2">配色方案</div>
          <el-radio-group v-model="colorScheme" class="mb-2">
            <el-radio label="蓝色系">蓝色系</el-radio>
            <el-radio label="绿色系">绿色系</el-radio>
            <el-radio label="红色系">红色系</el-radio>
            <el-radio label="黑白灰">黑白灰</el-radio>
            <el-radio label="纯黑白">纯黑白</el-radio>
            <el-radio label="暖色调">暖色调</el-radio>
            <el-radio label="冷色调">冷色调</el-radio>
            <el-radio label="其他">其他</el-radio>
          </el-radio-group>
          <el-input v-if="colorScheme === '其他'" v-model="colorSchemeCustom" placeholder="请输入配色方案" class="w-full" />
        </div>
      </div>
      
      <!-- 详细描述输入框 -->
      <div class="relative">
        <div class="text-base font-medium mb-2">详细描述（可选）</div>
        <el-input
            v-model="prompt"
            :maxlength="2000"
            :placeholder="placeholder"
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
                此功能仅针对授权用户开放，前往<a
                  class="text-blue-600"
                  href="https://www.gin-vue-admin.com/empower/"
                  target="_blank"
              >购买授权</a
              >
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
      <div v-if="outPut && htmlFromLLM">
        <el-tabs type="border-card">
          <el-tab-pane label="页面预览">
            <div class="h-[500px] overflow-auto bg-gray-50 p-4 rounded">
              <div v-if="!loadedComponents" class="text-gray-500 text-center py-4">
                组件加载中...
              </div>
              <component
                v-else
                :is="loadedComponents" 
                class="vue-component-container w-full"
              />
            </div>
          </el-tab-pane>
          <el-tab-pane label="源代码">
            <div class="relative h-[500px] overflow-auto bg-gray-50 p-4 rounded">
              <el-button 
                type="primary" 
                :icon="DocumentCopy" 
                class="absolute top-2 right-2 px-2 py-1" 
                @click="copySnippet(htmlFromLLM)" 
                plain
              >
                复制
              </el-button>
              <pre class="mt-10 whitespace-pre-wrap">{{ htmlFromLLM }}</pre>
            </div>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
  </div>
</template>

<script setup>
import { createWeb } from '@/api/autoCode'
import { ref, reactive, markRaw } from 'vue'
import * as Vue from "vue";
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage } from 'element-plus'
import { defineAsyncComponent } from 'vue'
import { DocumentCopy } from '@element-plus/icons-vue'
import { loadModule } from "vue3-sfc-loader";

defineOptions({
  name: 'Picture'
})

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
const copySnippet = (vueString) => {
  navigator.clipboard.writeText(vueString)
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

// 选项模式相关变量
const pageType = ref('企业官网')
const pageTypeCustom = ref('')
const contentBlocks = ref(['Banner轮播图', '产品/服务介绍'])
const contentBlocksCustom = ref('')
const stylePreference = ref('简约')
const stylePreferenceCustom = ref('')
const layoutDesign = ref('响应式')
const layoutDesignCustom = ref('')
const colorScheme = ref('蓝色系')
const colorSchemeCustom = ref('')

// 页面用途与内容板块的推荐映射关系
const pageTypeContentMap = {
  '企业官网': ['Banner轮播图', '产品/服务介绍', '功能特点展示', '客户案例', '联系表单'],
  '电商页面': ['Banner轮播图', '商品列表', '商品卡片', '购物车', '商品分类', '热门推荐', '限时特惠', '结算页面', '用户评价'],
  '个人博客': ['Banner轮播图', '新闻/博客列表', '用户评价', '联系表单'],
  '产品介绍': ['Banner轮播图', '产品/服务介绍', '功能特点展示', '价格表', 'FAQ/常见问题'],
  '活动落地页': ['Banner轮播图', '功能特点展示', '联系表单', '数据统计']
}

const prompt = ref('')

// 判断是否返回的标志
const outPut = ref(false)
// 容纳llm返回的vue组件代码
const htmlFromLLM = ref("")

// 存储加载的组件
const loadedComponents = ref(null)

const loadVueComponent = async (vueCode) => {
  try {
    // 使用内存中的虚拟路径
    const fakePath = `virtual:component-0.vue`
    
    const component = defineAsyncComponent({
      loader: async () => {
        try {
          const options = {
            moduleCache: {
              vue: Vue,
            },
            getFile(url) {
              // 处理所有可能的URL格式，包括相对路径、绝对路径等
              // 提取路径的最后部分，忽略查询参数
              const fileName = url.split('/').pop().split('?')[0]
              const componentFileName = fakePath.split('/').pop()
              
              // 如果文件名包含我们的组件名称，或者url完全匹配fakePath
              if (fileName === componentFileName || url === fakePath || 
                  url === `./component/0.vue`) {
                return Promise.resolve({
                  type: '.vue',
                  getContentData: () => vueCode
                })
              }
              
              console.warn('请求未知文件:', url)
              return Promise.reject(new Error(`找不到文件: ${url}`))
            },
            addStyle(textContent) {
              // 不再将样式添加到document.head，而是返回样式内容
              // 稍后会将样式添加到Shadow DOM中
              return textContent
            },
            handleModule(type, source, path, options) {
              // 默认处理器
              return undefined
            },
            log(type, ...args) {
              console.log(`[vue3-sfc-loader] [${type}]`, ...args)
            }
          }
          
          // 尝试加载组件
          const comp = await loadModule(fakePath, options)
          return comp.default || comp
        } catch (error) {
          console.error('组件加载详细错误:', error)
          throw error
        }
      },
      loadingComponent: {
        template: '<div>加载中...</div>'
      },
      errorComponent: {
        props: ['error'],
        template: '<div>组件加载失败: {{ error && error.message }}</div>',
        setup(props) {
          console.error('错误组件收到的错误:', props.error)
          return {}
        }
      },
      // 添加超时和重试选项
      timeout: 30000,
      delay: 200,
      suspensible: false,
      onError(error, retry, fail) {
        console.error('加载错误，细节:', error)
        fail()
      }
    })

    // 创建一个包装组件，使用Shadow DOM隔离样式
    const ShadowWrapper = {
      name: 'ShadowWrapper',
      setup() {
        return {}
      },
      render() {
        return Vue.h('div', { class: 'shadow-wrapper' })
      },
      mounted() {
        // 创建Shadow DOM
        const shadowRoot = this.$el.attachShadow({ mode: 'open' })
        
        // 创建一个容器元素
        const container = document.createElement('div')
        container.className = 'shadow-container'
        shadowRoot.appendChild(container)
        
        // 提取组件中的样式
        const styleContent = vueCode.match(/<style[^>]*>([\s\S]*?)<\/style>/i)?.[1] || ''
        
        // 创建样式元素并添加到Shadow DOM
        if (styleContent) {
          const style = document.createElement('style')
          style.textContent = styleContent
          shadowRoot.appendChild(style)
        }
        
        // 创建Vue应用并挂载到Shadow DOM容器中
        const app = Vue.createApp({
          render: () => Vue.h(component)
        })
        app.mount(container)
      }
    }

    loadedComponents.value = markRaw(ShadowWrapper)
    return ShadowWrapper
  } catch (error) {
    console.error('组件创建总错误:', error)
    return null
  }
}

// 当页面用途改变时，更新内容板块的选择
const handlePageTypeChange = (value) => {
  if (value !== '其他' && pageTypeContentMap[value]) {
    contentBlocks.value = [...pageTypeContentMap[value]]
  }
}

const llmAutoFunc = async () => {
  // 构建完整的描述，包含选项模式的选择
  let fullPrompt = ''
  
  // 添加页面用途
  fullPrompt += `页面用途: ${pageType.value === '其他' ? pageTypeCustom.value : pageType.value}\n`
  
  // 添加内容板块
  fullPrompt += '主要内容板块: '
  const blocks = contentBlocks.value.filter(block => block !== '其他')
  if (contentBlocksCustom.value) {
    blocks.push(contentBlocksCustom.value)
  }
  fullPrompt += blocks.join(', ') + '\n'
  
  // 添加风格偏好
  fullPrompt += `风格偏好: ${stylePreference.value === '其他' ? stylePreferenceCustom.value : stylePreference.value}\n`
  
  // 添加设计布局
  fullPrompt += `设计布局: ${layoutDesign.value === '其他' ? layoutDesignCustom.value : layoutDesign.value}\n`
  
  // 添加配色方案
  fullPrompt += `配色方案: ${colorScheme.value === '其他' ? colorSchemeCustom.value : colorScheme.value}\n`
  
  // 添加用户的详细描述
  if (prompt.value) {
    fullPrompt += `\n详细描述: ${prompt.value}`
  }
  
  const res = await createWeb({web: fullPrompt, command: 'createWeb'})
  if (res.code === 0) {
    outPut.value = true
    // 添加返回的Vue组件代码到数组
    htmlFromLLM.value = res.data
    // 加载新生成的组件
    await loadVueComponent(res.data)
  }
}

const placeholder = ref(`补充您对页面的其他要求或特殊需求，例如：特别强调的元素、参考网站、交互效果等。`)
</script>
