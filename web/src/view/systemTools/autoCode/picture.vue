<template>
  <div>
    <warning-bar
        href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
        title="此功能为开发环境使用，不建议发布到生产，具体使用效果请点我观看。"
    />
    <div v-if="!isAdd" class="gva-search-box">
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
            :placeholder="`现已完全免费\n试试复制一张图片然后按下ctrl+v或者commend+v\n试试描述你的表，让AI帮你完成。\n此功能需要到插件市场个人中心获取自己的AI-Path，把AI-Path填入config.yaml下的autocode-->ai-path，重启项目即可使用。\n按下 Ctrl+Enter 或 Cmd+Enter 直接生成`"
            :rows="5"
            resize="none"
            type="textarea"
            @blur="handleBlur"
            @focus="handleFocus"
        />

        <div class="flex absolute right-28 bottom-2">
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
                :disabled="form.onlyTemplate"
                type="primary"
                @click="eyeFunc()"
            >
              <el-icon size="18">
                <ai-gva/>
              </el-icon>
              识图
            </el-button>
          </el-tooltip>
        </div>

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
                :disabled="form.onlyTemplate"
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
        <div
            v-for="(snippet, index) in htmlFromLLM"
            :key="index"
            class="mb-6 p-4 border"
        >
          <el-button type="primary" :icon="Upload" class="px-2 py-1" @click="copySnippet(snippet)" plain>复制</el-button>
          <!-- 用 v-html 渲染该片段 -->
          <div class="mt-2" v-html="snippet"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  getDB,
  getMeta,
  getPackageApi,
  llmAuto, eye
} from '@/api/autoCode'
import {getDict} from '@/utils/dictionary'
import {ref, watch, nextTick} from 'vue'
import {useRoute, useRouter} from 'vue-router'
import {ElMessage, ElMessageBox} from 'element-plus'
import WarningBar from '@/components/warningBar/warningBar.vue'

const handleFocus = () => {
  document.addEventListener('keydown', handleKeydown);
  document.addEventListener('paste', handlePaste);
}

const handleBlur = () => {
  document.removeEventListener('keydown', handleKeydown);
  document.removeEventListener('paste', handlePaste);
}

const handleKeydown = (event) => {
  if ((event.ctrlKey || event.metaKey) && event.key === 'Enter') {
    llmAutoFunc()
  }
}

const handlePaste = (event) => {
  const items = event.clipboardData.items;
  for (let i = 0; i < items.length; i++) {
    if (items[i].type.indexOf('image') !== -1) {
      const file = items[i].getAsFile();
      const reader = new FileReader();
      reader.onload = async (e) => {
        const base64String = e.target.result;
        const res = await eye({picture: base64String, command: 'eye'})
        if (res.code === 0) {
          prompt.value = res.data
          llmAutoFunc()
        }
      };
      reader.readAsDataURL(file);
    }
  }
};


const prompt = ref('')

const eyeFunc = async () => {
  const input = document.createElement('input');
  input.type = 'file';
  input.accept = 'image/*';

  input.onchange = (event) => {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = async (e) => {
        const base64String = e.target.result;

        const res = await eye({picture: base64String, command: 'eye'})
        if (res.code === 0) {
          prompt.value = res.data
          llmAutoFunc()
        }
      };
      reader.readAsDataURL(file);
    }
  };

  input.click();
}

// 判断是否返回的标志
const outPut = ref(false)
// 容纳llm返回的html
const htmlFromLLM = ref([])

const llmAutoFunc = async (flag) => {
  const res = await llmAuto({
    prompt: flag ? '结构体名称为' + form.value.structName : prompt.value
  })
  if (res.code === 0) {
    outPut.value = true
    // 使用mock数据模拟大模型返回值
    const mock1 = `<div class="flex flex-col min-h-screen bg-white">
    <!-- 顶部导航栏 -->
    <header class="bg-white shadow-md">
        <div class="container mx-auto px-6 py-4 flex items-center justify-between">
            <div class="text-2xl font-bold text-blue-600">LOGO</div>
            <nav class="flex space-x-6">
                <a href="#" class="text-black hover:text-blue-600 transition duration-300">首页</a>
                <a href="#" class="text-black hover:text-blue-600 transition duration-300">关于我们</a>
                <a href="#" class="text-black hover:text-blue-600 transition duration-300">服务</a>
                <a href="#" class="text-black hover:text-blue-600 transition duration-300">联系我们</a>
            </nav>
        </div>
    </header>

    <!-- 中间内容区域 -->
    <main class="flex-grow container mx-auto px-6 py-8 flex">
        <!-- 左侧边栏 -->
        <aside class="w-1/4 bg-gray-50 p-6 rounded-lg shadow-sm">
            <nav class="space-y-3">
                <a href="#" class="block text-black hover:text-blue-600 transition duration-300">导航1</a>
                <a href="#" class="block text-black hover:text-blue-600 transition duration-300">导航2</a>
                <a href="#" class="block text-black hover:text-blue-600 transition duration-300">导航3</a>
                <a href="#" class="block text-black hover:text-blue-600 transition duration-300">导航4</a>
            </nav>
        </aside>

        <!-- 右侧主要内容区 -->
        <div class="w-3/4 ml-8">
            <!-- 轮播图 -->
            <div class="bg-gray-200 h-64 mb-8 rounded-lg shadow-sm flex items-center justify-center text-gray-600">
                轮播图
            </div>

            <!-- 新闻列表 -->
            <div class="mb-8">
                <h2 class="text-xl font-bold text-black mb-4">最新新闻</h2>
                <ul class="space-y-4">
                    <li class="p-4 bg-gray-50 rounded-lg shadow-sm hover:shadow-md transition duration-300">
                        <a href="#" class="text-black hover:text-blue-600 transition duration-300">新闻标题1</a>
                        <p class="text-sm text-gray-600 mt-1">新闻简介内容...</p>
                    </li>
                    <li class="p-4 bg-gray-50 rounded-lg shadow-sm hover:shadow-md transition duration-300">
                        <a href="#" class="text-black hover:text-blue-600 transition duration-300">新闻标题2</a>
                        <p class="text-sm text-gray-600 mt-1">新闻简介内容...</p>
                    </li>
                    <li class="p-4 bg-gray-50 rounded-lg shadow-sm hover:shadow-md transition duration-300">
                        <a href="#" class="text-black hover:text-blue-600 transition duration-300">新闻标题3</a>
                        <p class="text-sm text-gray-600 mt-1">新闻简介内容...</p>
                    </li>
                </ul>
            </div>

            <!-- 广告位 -->
            <div class="bg-gray-200 h-32 rounded-lg shadow-sm flex items-center justify-center text-gray-600">
                广告位
            </div>
        </div>
    </main>

    <!-- 底部 -->
    <footer class="bg-gray-800 text-white py-6">
        <div class="container mx-auto px-6 text-center">
            <p class="text-sm">© 2023 公司名称. 保留所有权利.</p>
            <p class="text-sm mt-2">联系我们: info@example.com</p>
        </div>
    </footer>
</div>`
    const mock2 = ` <div class="flex flex-col min-h-screen bg-[#e9f0fc]">
    <!-- 上布局 -->
    <div class="flex flex-1">
      <!-- 左侧登录表单区域 -->
      <div class="w-1/2 p-12 flex items-center justify-center">
        <div class="w-full max-w-md bg-white p-8 rounded-xl shadow-2xl">
          <h2 class="text-3xl font-bold text-[#2c3e50] mb-8 text-center">登录</h2>
          <form class="space-y-6">
            <div>
              <label class="block text-sm font-medium text-[#2c3e50] mb-2">用户名</label>
              <input type="text" class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#20a0ff] transition duration-300" placeholder="请输入用户名" />
            </div>
            <div>
              <label class="block text-sm font-medium text-[#2c3e50] mb-2">密码</label>
              <input type="password" class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#20a0ff] transition duration-300" placeholder="请输入密码" />
            </div>
            <div>
              <label class="block text-sm font-medium text-[#2c3e50] mb-2">验证码</label>
              <div class="flex space-x-4">
                <input type="text" class="w-3/4 px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#20a0ff] transition duration-300" placeholder="请输入验证码" />
                <button class="w-1/4 px-4 py-3 bg-[#20a0ff] text-white rounded-lg hover:bg-[#1d8fe0] transition duration-300">获取验证码</button>
              </div>
            </div>
            <button type="submit" class="w-full px-4 py-3 bg-[#20a0ff] text-white rounded-lg hover:bg-[#1d8fe0] transition duration-300">登录</button>
          </form>
        </div>
      </div>

      <!-- 右侧功能介绍区 -->
      <div class="w-1/2 p-12 flex items-center justify-center">
        <div class="relative w-full max-w-2xl">
          <!-- 复古电视机样式 -->
          <div class="relative bg-gray-800 rounded-2xl p-8 shadow-2xl">
            <div class="absolute inset-0 bg-black rounded-2xl opacity-50"></div>
            <div class="relative z-10">
              <!-- 电视屏幕 -->
              <div class="bg-gray-700 rounded-lg p-6">
                <div class="bg-gray-900 rounded-lg p-6">
                  <!-- SVG图形和动画内容 -->
                  <svg class="w-full h-64" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
                    <circle cx="50" cy="50" r="45" fill="#20a0ff" />
                    <text x="50%" y="50%" text-anchor="middle" fill="#ffffff" font-size="24" dy=".3em">TV</text>
                    <!-- 模拟电视内容 -->
                    <animateTransform attributeName="transform" type="rotate" from="0 50 50" to="360 50 50" dur="10s" repeatCount="indefinite" />
                  </svg>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 中部空白区域 -->
    <div class="flex-1 bg-[#e9f0fc] p-12">
      <!-- 广告或其他内容 -->
      <div class="bg-white p-8 rounded-xl shadow-2xl text-center">
        <h3 class="text-2xl font-bold text-[#2c3e50] mb-4">广告位</h3>
        <p class="text-gray-600">这里可以放置广告或其他内容。</p>
      </div>
    </div>

    <!-- 底部 -->
    <footer class="bg-[#2c3e50] text-white py-6">
      <div class="container mx-auto px-6 text-center">
        <p class="text-sm">© 2023 公司名称. 保留所有权利.</p>
        <p class="text-sm mt-2">联系我们: info@example.com</p>
      </div>
    </footer>
  </div>`
    htmlFromLLM.value.push(mock1, mock2)
    console.log(htmlFromLLM.value);
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

const isAdd = ref(false)


defineOptions({
  name: 'AutoCode'
})
const gormModelList = ['id', 'created_at', 'updated_at', 'deleted_at']

const dataModelList = ['created_by', 'updated_by', 'deleted_by']

const route = useRoute()
const router = useRouter()
const dbform = ref({
  businessDB: '',
  dbName: '',
  tableName: ''
})
const fdMap = ref({})
const form = ref({
  structName: '',
  tableName: '',
  packageName: '',
  package: '',
  abbreviation: '',
  description: '',
  businessDB: '',
  autoCreateApiToSql: true,
  autoCreateMenuToSql: true,
  autoCreateBtnAuth: false,
  autoMigrate: true,
  gvaModel: true,
  autoCreateResource: false,
  onlyTemplate: false,
  isTree: false,
  treeJson: "",
  fields: []
})
const rules = ref({
  structName: [
    {required: true, message: '请输入结构体名称', trigger: 'blur'}
  ],
  abbreviation: [
    {required: true, message: '请输入结构体简称', trigger: 'blur'}
  ],
  description: [
    {required: true, message: '请输入结构体描述', trigger: 'blur'}
  ],
  packageName: [
    {
      required: true,
      message: '文件名称：sysXxxxXxxx',
      trigger: 'blur'
    }
  ],
  package: [{required: true, message: '请选择package', trigger: 'blur'}]
})


const dbList = ref([])
const dbOptions = ref([])

const getDbFunc = async () => {
  dbform.value.dbName = ''
  dbform.value.tableName = ''
  const res = await getDB({businessDB: dbform.value.businessDB})
  if (res.code === 0) {
    dbOptions.value = res.data.dbs
    dbList.value = res.data.dbList
  }
}

const setFdMap = async () => {
  const fdTypes = ['string', 'int', 'bool', 'float64', 'time.Time']
  fdTypes.forEach(async (fdtype) => {
    const res = await getDict(fdtype)
    res &&
    res.forEach((item) => {
      fdMap.value[item.label] = fdtype
    })
  })
}
const getAutoCodeJson = async (id) => {
  const res = await getMeta({id: Number(id)})
  if (res.code === 0) {
    const add = route.query.isAdd
    isAdd.value = add
    form.value = JSON.parse(res.data.meta)
    if (isAdd.value) {
      form.value.fields.forEach((item) => {
        item.disabled = true
      })
    }
  }
}

const pkgs = ref([])
const getPkgs = async () => {
  const res = await getPackageApi()
  if (res.code === 0) {
    pkgs.value = res.data.pkgs
  }
}


const init = () => {
  getDbFunc()
  setFdMap()
  getPkgs()
  const id = route.params.id
  if (id) {
    getAutoCodeJson(id)
  }
}
init()

watch(
    () => route.params.id,
    () => {
      if (route.name === 'autoCodeEdit') {
        init()
      }
    }
)

const getCatch = () => {
  const data = window.sessionStorage.getItem('autoCode')
  if (data) {
    form.value = JSON.parse(data)
  }
}

const clearCatch = async () => {
  form.value = {
    structName: '',
    tableName: '',
    packageName: '',
    package: '',
    abbreviation: '',
    description: '',
    businessDB: '',
    autoCreateApiToSql: true,
    autoCreateMenuToSql: true,
    autoCreateBtnAuth: false,
    autoMigrate: true,
    gvaModel: true,
    autoCreateResource: false,
    onlyTemplate: false,
    isTree: false,
    treeJson: "",
    fields: []
  }
  await nextTick()
  window.sessionStorage.removeItem('autoCode')
}

getCatch()


watch(
    () => form.value.onlyTemplate,
    (val) => {
      if (val) {
        ElMessageBox.confirm(
            '使用基础模板将不会生成任何结构体和CURD,仅仅配置enter等属性方便自行开发非CURD逻辑',
            '注意',
            {
              confirmButtonText: '继续',
              cancelButtonText: '取消',
              type: 'warning'
            }
        )
            .then(() => {
              form.value.fields = []
            })
            .catch(() => {
              form.value.onlyTemplate = false
            })
      }
    }
)
</script>
