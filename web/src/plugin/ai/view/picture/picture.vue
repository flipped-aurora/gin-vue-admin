<template>
  <div>
    <warning-bar
        href="https://plugin.gin-vue-admin.com/license"
        title="此功能只针对授权用户开放，点我【购买授权】"
    />
<div class="gva-search-box">
      <div class="text-xl mb-2 text-gray-600 dark:text-slate-300">
        AI前端工程师<a
          class="text-blue-600 text-sm ml-4 dark:text-blue-400"
          href="https://plugin.gin-vue-admin.com/#/layout/userInfo/center"
          target="_blank"
        >获取AiPath</a
      >
      </div>
      
      <!-- 选项模式 -->
      <div class="mb-4">
        <div class="mb-3">
          <div class="text-base font-medium mb-2 dark:text-slate-200">页面用途</div>
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
          <div class="text-base font-medium mb-2 dark:text-slate-200">主要内容板块</div>
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
          <div class="text-base font-medium mb-2 dark:text-slate-200">风格偏好</div>
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
          <div class="text-base font-medium mb-2 dark:text-slate-200">是否响应式</div>
          <el-radio-group v-model="responsive" class="mb-2">
            <el-radio label="是">是</el-radio>
            <el-radio label="否">否</el-radio>
          </el-radio-group>
        </div>

        <div class="mb-3">
          <div class="text-base font-medium mb-2 dark:text-slate-200">配色方案</div>
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
    </div>
    <!-- 结果区：左右分栏（左窄右宽，窄屏自动堆叠） -->
    <div class="mt-4">
      <div
        class="grid grid-cols-1 lg:grid-cols-12 gap-4"
      >
        <!-- 左栏：对话式 AI 输出（消息历史 + 底部统一输入），占 3/12 -->
        <div class="lg:col-span-3 flex flex-col rounded-lg border border-gray-200 dark:border-slate-700 overflow-hidden h-[600px]">
          <div class="flex items-center justify-between px-4 py-2 bg-gray-50 dark:bg-slate-800 border-b border-gray-200 dark:border-slate-700">
            <span class="font-medium text-gray-700 dark:text-slate-200">
              AI 输出
              <span v-if="turnCount" class="ml-2 text-xs font-normal text-gray-400 dark:text-slate-500">
                第 {{ turnCount }} 轮
              </span>
            </span>
            <el-button
              v-if="streaming"
              type="danger"
              size="small"
              plain
              @click="stopStreaming"
            >
              停止生成
            </el-button>
          </div>

          <!-- 消息历史列表（始终展示，调整时不清理） -->
          <div
            ref="messageListRef"
            class="flex-1 overflow-auto p-4 space-y-4 bg-white dark:bg-slate-900"
          >
            <template v-if="messages.length">
              <div
                v-for="msg in messages"
                :key="msg.id"
                class="flex"
                :class="msg.role === 'user' ? 'justify-end' : 'justify-start'"
              >
                <div
                  class="max-w-[85%] rounded-lg px-3 py-2 text-sm leading-6"
                  :class="msg.role === 'user'
                    ? 'bg-blue-600 text-white'
                    : 'bg-gray-100 text-gray-800 dark:bg-slate-800 dark:text-slate-100'"
                >
                  <!-- 用户消息：直接展示文本 -->
                  <div v-if="msg.role === 'user'" class="whitespace-pre-wrap break-words">{{ msg.query }}</div>

                  <!-- 助手消息：顶部可折叠 think + 正文 -->
                  <div v-else class="space-y-2">
                    <!-- 流式进行中：点击可展开实时思考内容 -->
                    <div
                      v-if="msg.streaming"
                      class="flex items-center gap-2 text-xs text-sky-600 dark:text-sky-400"
                      :class="msg.hasThink ? 'cursor-pointer select-none' : ''"
                      @click="msg.hasThink && toggleThink(msg.id)"
                    >
                      <el-icon class="animate-spin"><RefreshRight/></el-icon>
                      <span>AI 思考中...</span>
                      <el-icon v-if="msg.hasThink" class="ml-0.5">
                        <ArrowRight v-if="msg.thinkCollapsed"/><ArrowDown v-else/>
                      </el-icon>
                      <span v-if="msg.hasThink" class="text-sky-500/70">
                        {{ msg.thinkCollapsed ? '点击查看思考过程' : '点击折叠' }}
                      </span>
                    </div>

                    <!-- 可折叠 Think 条（流式与完成后共用） -->
                    <div
                      v-if="msg.hasThink"
                      class="rounded border border-amber-200 bg-amber-50 dark:border-amber-700/50 dark:bg-amber-900/20"
                    >
                      <button
                        type="button"
                        class="w-full flex items-center gap-1 px-2 py-1 text-xs font-medium text-amber-700 dark:text-amber-400"
                        @click="toggleThink(msg.id)"
                      >
                        <el-icon><ArrowRight v-if="msg.thinkCollapsed"/><ArrowDown v-else/></el-icon>
                        <span>Think{{ msg.streaming ? '（实时）' : '' }}</span>
                        <span class="ml-auto text-amber-500/70">{{ msg.thinkCollapsed ? '展开' : '折叠' }}</span>
                      </button>
                      <pre
                        v-show="!msg.thinkCollapsed"
                        class="mx-2 mb-2 max-h-[160px] overflow-auto whitespace-pre-wrap rounded bg-slate-950 p-2 text-[11px] leading-5 text-slate-100"
                      >{{ msg.think }}</pre>
                    </div>

                    <!-- 正文（生成中时实时展示，去掉 think 后的内容） -->
                    <pre
                      v-if="msg.displayContent"
                      class="whitespace-pre-wrap break-words text-xs leading-5 max-h-[240px] overflow-auto bg-slate-950 text-slate-100 rounded p-2"
                    >{{ msg.displayContent }}</pre>
                    <div
                      v-else-if="!msg.streaming"
                      class="text-xs text-gray-400 dark:text-slate-500"
                    >
                      模型未返回可显示内容。
                    </div>

                    <!-- 操作标签 -->
                    <div class="flex flex-wrap gap-2" v-if="msg.hasThink || msg.displayContent">
                      <el-tag
                        v-if="msg.hasThink"
                        size="small"
                        effect="plain"
                        type="warning"
                        class="cursor-pointer"
                        @click="openDrawer('Think', msg.think)"
                      >完整 Think</el-tag>
                      <el-tag
                        v-if="msg.displayContent"
                        size="small"
                        effect="plain"
                        class="cursor-pointer"
                        @click="openDrawer('原始文件', msg.displayContent)"
                      >原始文件</el-tag>
                    </div>
                  </div>
                </div>
              </div>
            </template>

            <!-- 首次进入的空态 -->
            <div v-else class="h-full grid place-items-center">
              <div class="text-center text-gray-400 dark:text-slate-500">
                <el-icon :size="40" class="mb-2"><ChatDotRound/></el-icon>
                <div class="text-sm">在下方输入对页面的要求，开始生成</div>
              </div>
            </div>
          </div>

          <!-- 底部统一输入栏（首轮与调整共用），按钮悬浮在 textarea 右下角 -->
          <div class="border-t border-gray-200 dark:border-slate-700 p-3 bg-gray-50 dark:bg-slate-800">
            <div class="relative">
              <el-input
                v-model="chatInput"
                type="textarea"
                :rows="3"
                resize="none"
                :placeholder="chatPlaceholder"
                @keydown="handleChatKeydown"
                class="chat-textarea"
              />
              <el-button
                type="primary"
                :loading="streaming"
                :disabled="!chatInput.trim() && !streaming"
                @click="sendChat"
                class="chat-send-btn !absolute right-2 bottom-2"
              >
                {{ hasConversation ? '调整' : '生成' }}
              </el-button>
            </div>
            <div class="mt-1 text-[11px] text-gray-400 dark:text-slate-500">
              <span v-if="!hasConversation">首轮会自动带入上方勾选项</span>
              <span v-else>本轮仅发送文案，复用同一会话（Ctrl/Cmd + Enter 发送）</span>
            </div>
          </div>
        </div>

        <!-- 右栏：页面预览 / 源代码，占 9/12，与左栏等高 -->
        <div class="lg:col-span-9 flex flex-col h-[600px]">
          <el-tabs type="border-card" class="flex-1 flex flex-col picture-tabs">
            <el-tab-pane label="页面预览" class="picture-tab-pane">
              <div class="flex-1 overflow-auto bg-gray-50 p-4 rounded dark:bg-slate-900">
                <div
                  v-if="!loadedComponents"
                  class="text-gray-500 text-center py-4 dark:text-slate-400"
                >
                  {{ streaming ? '等待生成完成后渲染...' : '组件加载中...' }}
                </div>
                <component
                  v-else
                  :is="loadedComponents"
                  class="vue-component-container w-full"
                />
              </div>
            </el-tab-pane>
            <el-tab-pane label="源代码" class="picture-tab-pane">
              <div class="relative flex-1 overflow-auto bg-gray-50 p-4 rounded dark:bg-slate-900">
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

    <!-- Think / 原始文件 抽屉 -->
    <el-drawer
      v-model="drawerVisible"
      :title="drawerTitle"
      size="50%"
      destroy-on-close
    >
      <pre class="h-full overflow-auto whitespace-pre-wrap rounded bg-slate-950 p-4 text-xs leading-6 text-slate-100">{{ drawerContent }}</pre>
    </el-drawer>
  </div>
</template>

<script setup>
import { createWebStream } from '@/api/autoCode'
import { ref, reactive, markRaw, computed, nextTick } from 'vue'
import * as Vue from "vue";
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage } from 'element-plus'
import { defineAsyncComponent } from 'vue'
import { DocumentCopy, RefreshRight, ArrowRight, ArrowDown, ChatDotRound } from '@element-plus/icons-vue'
import { loadModule } from "vue3-sfc-loader";

defineOptions({
  name: 'Picture'
})

// 统一输入栏的 Ctrl/Cmd+Enter 发送
const handleChatKeydown = (event) => {
  if ((event.ctrlKey || event.metaKey) && event.key === 'Enter') {
    event.preventDefault()
    sendChat()
  }
}

// 历史消息滚动到底
const scrollMessagesToBottom = () => {
  nextTick(() => {
    if (messageListRef.value) {
      messageListRef.value.scrollTop = messageListRef.value.scrollHeight
    }
  })
}

// 折叠/展开某条助手消息的 think
const toggleThink = (id) => {
  const msg = messages.value.find((m) => m.id === id)
  if (msg) msg.thinkCollapsed = !msg.thinkCollapsed
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
// 是否响应式（替代原「设计布局」）
const responsive = ref('是')
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

// 容纳llm返回的vue组件代码（当前最新一版）
const htmlFromLLM = ref("")

// 存储加载的组件
const loadedComponents = ref(null)

// === 流式相关状态 ===
const streaming = ref(false)
const messageListRef = ref(null)
let abortController = null

// 对话消息列表：{ id, role: 'user'|'assistant', query, think, displayContent,
//                hasThink, thinkCollapsed, streaming }
const messages = ref([])
let messageSeq = 0
const newId = () => ++messageSeq

// === 多轮对话相关状态 ===
// 第一轮拿到 conversation_id 后保存，后续轮次带上它以复用同一个 Dify 会话
const conversationId = ref('')
// 是否已开启会话（决定输入框是「生成」还是「调整」）
const hasConversation = computed(() => Boolean(conversationId.value))
// 已完成的 AI 回合数（用于头部展示）
const turnCount = computed(
  () => messages.value.filter((m) => m.role === 'assistant' && !m.streaming).length
)
// 输入栏占位文案
const chatPlaceholder = computed(() =>
  hasConversation.value
    ? '继续调整：例如「把主色换成绿色」「顶部加一个轮播」「去掉价格表」'
    : '描述你想要的页面，或直接发送（会自动带上上方勾选项）'
)
// 统一输入栏内容（首轮与调整共用）
const chatInput = ref('')

// Think 抽屉（查看完整内容）
const drawerVisible = ref(false)
const drawerTitle = ref('')
const drawerContent = ref('')
const openDrawer = (title, content) => {
  drawerTitle.value = title
  drawerContent.value = content
  drawerVisible.value = true
}

// 文本规整
const normalizeText = (value) =>
  String(value || '')
    .replace(/\r\n/g, '\n')
    .trim()

// 把流式累计文本拆成 think（<think>...</think> 块）和正文两部分
// 兼容流式中尚未闭合的 <think>：开头标签之后的内容视作正在进行的思考
const parseDisplay = (content) => {
  const text = normalizeText(content)
  const thinkParts = []
  let rest = text

  // 1. 先提取并剥离所有已闭合的 think 块
  const closedRe = /<think(?:ing)?>\s*([\s\S]*?)<\/think(?:ing)?>/gi
  let m
  while ((m = closedRe.exec(rest)) !== null) {
    if (m[1] && m[1].trim()) thinkParts.push(m[1])
  }
  rest = rest.replace(/<think(?:ing)?>\s*[\s\S]*?<\/think(?:ing)?>/gi, '')

  // 2. 处理未闭合的 think（流式中）：把开头标签之后的所有内容当作思考
  const openMatch = rest.match(/<think(?:ing)?>\s*([\s\S]*)$/i)
  if (openMatch) {
    if (openMatch[1] && openMatch[1].trim()) thinkParts.push(openMatch[1])
    rest = rest.replace(/<think(?:ing)?>\s*[\s\S]*$/i, '')
  }

  return { think: normalizeText(thinkParts.join('\n')), content: normalizeText(rest) }
}

// 从正文里提取 Vue SFC 代码（剥 markdown 代码围栏）
const extractVueCode = (text) => {
  if (!text) return ''
  // 优先匹配 ``` 代码围栏
  const fenced = text.match(/```(?:vue|html|sfc)?\s*([\s\S]*?)```/i)
  if (fenced && fenced[1]) return fenced[1].trim()
  // 否则尝试截取 <template> 到 </style>
  const startIdx = text.search(/<template>/i)
  const endIdx = text.search(/<\/style>/i)
  if (startIdx >= 0 && endIdx >= 0 && endIdx > startIdx) {
    return text.slice(startIdx, endIdx + '</style>'.length).trim()
  }
  return text.trim()
}

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

// 停止流式生成
const stopStreaming = () => {
  if (abortController) {
    abortController.abort()
  }
}

// 实时更新流式中的助手消息：把累计文本拆成 think / 正文写回该消息
const updateStreamingMessage = (msg, rawText) => {
  const { think, content } = parseDisplay(rawText)
  const wasThinking = msg.hasThink
  msg.think = think
  msg.hasThink = Boolean(think)
  // 进入思考阶段的瞬间自动展开，让用户立刻看到实时思考
  if (!wasThinking && msg.hasThink) msg.thinkCollapsed = false
  // 流式中展示去掉 think 后的正文（实时展开）
  msg.displayContent = content
  scrollMessagesToBottom()
}

// 流式发送内核：
// - isInitial=true 为首轮（重置会话、清掉组件、清空历史消息）
// - isInitial=false 为后续轮（保留历史、复用 conversation_id、仅刷新右栏组件）
const sendQuery = async (query, isInitial) => {
  if (isInitial) {
    messages.value = []
    loadedComponents.value = null
    htmlFromLLM.value = ''
    conversationId.value = ''
  } else {
    // 后续轮：刷新前先把旧组件置空，避免显示上一版
    loadedComponents.value = null
  }

  // 追加用户消息
  const userMsg = {
    id: newId(),
    role: 'user',
    query
  }
  messages.value.push(userMsg)

  // 追加占位助手消息（流式中实时填充）
  const assistantMsg = {
    id: newId(),
    role: 'assistant',
    think: '',
    displayContent: '',
    hasThink: false,
    thinkCollapsed: true, // 默认折叠，只占顶部一小部分
    streaming: true
  }
  messages.value.push(assistantMsg)

  streaming.value = true
  scrollMessagesToBottom()
  abortController = new AbortController()

  // 组装请求体：首轮只带 query；后续轮带上 conversation_id 复用同一 Dify 会话
  const payload = { query }
  if (conversationId.value) {
    payload.conversation_id = conversationId.value
  }

  // 记录本轮流式累计文本（取消时仍可用来拆分渲染）
  let rawText = ''

  try {
    const result = await createWebStream(payload, {
      controller: abortController,
      onMessage: (payload) => {
        if (payload.text) {
          rawText = payload.text
          updateStreamingMessage(assistantMsg, rawText)
        }
        // 上游可能在事件里返回 conversation_id，及时保存以便后续轮复用
        if (payload.conversationId && !conversationId.value) {
          conversationId.value = payload.conversationId
        }
      }
    })

    const finalText =
      (result && (result.answer || result.text)) || rawText || ''
    updateStreamingMessage(assistantMsg, finalText)

    // 保存/更新会话 id，保证下一轮能续上
    const newIdVal =
      result && (result.conversation_id || result.conversationId)
    if (newIdVal) conversationId.value = newIdVal

    // 提取 Vue SFC 代码并渲染到右栏
    const vueCode = extractVueCode(
      assistantMsg.displayContent || finalText
    )
    if (vueCode) {
      htmlFromLLM.value = vueCode
      await loadVueComponent(vueCode)
    }
  } catch (error) {
    // 用户主动取消时给出温和提示
    if (abortController?.signal?.aborted) {
      ElMessage({ message: '已停止生成', type: 'info' })
    } else {
      ElMessage.error(error?.message || 'AI 请求失败')
    }
    // 仍尝试用已收到的内容做拆分与渲染
    updateStreamingMessage(assistantMsg, rawText)
    const vueCode = extractVueCode(assistantMsg.displayContent)
    if (vueCode) {
      htmlFromLLM.value = vueCode
      await loadVueComponent(vueCode)
    }
  } finally {
    assistantMsg.streaming = false
    streaming.value = false
    abortController = null
  }
}

// 首轮：拼接所有勾选项 + 用户输入
const buildInitialPrompt = (userText) => {
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

  // 添加是否响应式
  fullPrompt += `是否响应式: ${responsive.value}\n`

  // 添加配色方案
  fullPrompt += `配色方案: ${colorScheme.value === '其他' ? colorSchemeCustom.value : colorScheme.value}\n`

  // 添加用户的输入文案
  if (userText) {
    fullPrompt += `\n详细描述: ${userText}`
  }

  return fullPrompt
}

// 统一发送入口：首轮带勾选项，后续轮只发文案、复用会话
const sendChat = async () => {
  const userText = chatInput.value.trim()
  if (!userText || streaming.value) return

  const isInitial = !hasConversation.value
  const query = isInitial ? buildInitialPrompt(userText) : userText

  chatInput.value = ''
  await sendQuery(query, isInitial)
}
</script>

<style scoped>
/* 让 textarea 右下角留出空间，避免文字被悬浮按钮挡住 */
.chat-textarea :deep(.el-textarea__inner) {
  padding-right: 96px;
  padding-bottom: 48px;
}
/* 右栏 tabs 撑满高度，使内容区与左栏等高 */
.picture-tabs :deep(.el-tabs__content) {
  flex: 1;
  overflow: hidden;
}
.picture-tab-pane {
  height: 100%;
}
</style>
