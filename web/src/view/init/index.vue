<template>
  <div id="userLayout" class="init-layout relative w-full overflow-hidden bg-white dark:bg-slate-900">
    <!-- 右侧蓝色斜切 banner（仅桌面端显示） -->
    <div class="banner-oblique absolute inset-y-0 right-0 hidden w-[56%] overflow-hidden bg-[#2264f2] md:block">
      <img class="absolute right-0 top-0 h-full w-auto max-w-none" src="@/assets/cover-redesign.svg" alt="banner" />
    </div>

    <!-- 左侧内容区 -->
    <div class="relative z-10 h-full w-full md:w-1/2">
      <div class="flex h-full min-h-0 items-center justify-center p-6 md:p-8">
        <div class="flex max-h-full min-h-0 w-4/5 flex-col overflow-hidden md:w-96">
          <!-- 顶部 logo + 大写字标（无副标题） -->
          <EntryBrand class="mb-6 shrink-0 md:mb-9" />
          <div v-if="!page.showForm" class="anim-fade-up flex min-h-0 flex-1 flex-col overflow-hidden">
            <h2 id="init-notice-title" class="mb-6 shrink-0 text-2xl font-bold dark:text-white">初始化须知</h2>
            <div
              ref="noticeScrollRef"
              class="min-h-0 flex-1 overflow-y-auto pr-2"
              role="region"
              aria-labelledby="init-notice-title"
              tabindex="0"
              @scroll="updateNoticeReadState"
            >
              <div ref="noticeContentRef">
                <div class="space-y-5">
              <div class="notice-group">
                <h4 class="mb-2 flex items-center gap-1.5 text-sm font-semibold text-[#2264f2]">
                  <el-icon><Reading /></el-icon>使用前准备
                </h4>
                <div class="space-y-3 rounded-xl bg-slate-50 p-4 dark:bg-slate-800/50">
                  <p class="flex gap-2.5 text-gray-600 dark:text-gray-300">
                    <span class="mt-[10px] h-1 w-1 shrink-0 rounded-full bg-[#2264f2]" />
                    <span>
                      Gin-Vue-Admin 是一个前后端分离开发框架，建议您具备一定的
                      <span class="font-bold text-[#2264f2]">
                        Vue、Go、数据库及服务器部署
                      </span>
                      基础后再进行使用。
                    </span>
                  </p>
                  <p class="flex gap-2.5 text-gray-600 dark:text-gray-300">
                    <span class="mt-[10px] h-1 w-1 shrink-0 rounded-full bg-[#2264f2]" />
                    <span>
                      请确认您已经阅读
                      <a class="mx-1 font-bold text-[#2264f2]" href="https://www.gin-vue-admin.com" target="_blank">
                        官方文档
                      </a>
                      及
                      <a class="font-bold text-[#2264f2]" href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=2"
                        target="_blank">
                        初始化视频
                      </a>
                      ，并了解项目初始化及部署流程。
                    </span>
                  </p>
                </div>
              </div>

              <div class="notice-group">
                <h4 class="mb-2 flex items-center gap-1.5 text-sm font-semibold text-[#2264f2]">
                  <el-icon><Coin /></el-icon>数据库要求
                </h4>
                <div class="space-y-3 rounded-xl bg-slate-50 p-4 dark:bg-slate-800/50">
                  <p class="flex gap-2.5 text-gray-600 dark:text-gray-300">
                    <span class="mt-[10px] h-1 w-1 shrink-0 rounded-full bg-[#2264f2]" />
                    <span>
                      如果您使用 MySQL 数据库，请确认数据库存储引擎为
                      <span class="font-bold text-[#2264f2]">
                        InnoDB
                      </span>
                      ，否则可能导致数据库功能异常。
                    </span>
                  </p>
                </div>
              </div>

              <div class="notice-group">
                <h4 class="mb-2 flex items-center gap-1.5 text-sm font-semibold text-[#2264f2]">
                  <el-icon><Lock /></el-icon>授权与许可
                </h4>
                <div class="space-y-3 rounded-xl bg-slate-50 p-4 dark:bg-slate-800/50">
                  <p class="flex gap-2.5 text-gray-600 dark:text-gray-300">
                    <span class="mt-[10px] h-1 w-1 shrink-0 rounded-full bg-[#2264f2]" />
                    <span>
                      本项目基于
                      <span class="font-bold text-[#2264f2]">
                        Business Source License（BSL）
                      </span>
                      协议发布。个人学习、技术研究、测试开发等符合协议约定的用途可以免费使用；
                      企业生产环境部署、商业项目开发、软件交付及商业服务等用途需要取得对应商业授权。
                    </span>
                  </p>
                  <p class="flex gap-2.5 text-gray-600 dark:text-gray-300">
                    <span class="mt-[10px] h-1 w-1 shrink-0 rounded-full bg-[#2264f2]" />
                    <span>
                      为保护软件版权及授权权益，系统可能包含授权验证机制，
                      用于验证当前环境授权状态、防止未经授权的商业部署。
                    </span>
                  </p>
                  <p class="flex gap-2.5 text-gray-600 dark:text-gray-300">
                    <span class="mt-[10px] h-1 w-1 shrink-0 rounded-full bg-[#2264f2]" />
                    <span>
                      当系统检测到当前环境未获得有效商业授权时，
                      系统可能自动进入
                      <span class="font-bold text-[#2264f2]">
                        未授权模式
                      </span>
                      ，展示授权提醒信息、水印标识。
                    </span>
                  </p>
                  <p class="flex gap-2.5 text-gray-600 dark:text-gray-300">
                    <span class="mt-[10px] h-1 w-1 shrink-0 rounded-full bg-[#2264f2]" />
                    <span>
                      如果您认为授权检测结果存在异常，或已经购买商业授权，
                      请通过官方授权管理页面提交授权信息进行核验。
                    </span>
                  </p>
                </div>
              </div>
                </div>

                <p class="mt-4 text-sm text-gray-400 dark:text-gray-500">
                  注：官方文档及社区主要用于项目学习与使用指导。
                  商业部署支持、技术服务及定制开发需求，请联系官方获取对应服务。
                </p>
              </div>
            </div>
            <p
              id="init-notice-hint"
              class="mt-2 flex h-5 shrink-0 items-center gap-1 text-xs text-gray-500 dark:text-gray-400"
              aria-live="polite"
            >
              <template v-if="noticeHint">
                <el-icon aria-hidden="true"><Bottom /></el-icon>
                <span>{{ noticeHint }}</span>
              </template>
            </p>
            <div class="mt-2 flex shrink-0 gap-4">
              <el-button type="primary" size="large" class="flex-1" @click="goDoc">
                阅读文档
              </el-button>
              <el-button
                type="primary"
                plain
                size="large"
                class="btn-hollow flex-1"
                :disabled="!canConfirmNotice"
                :title="canConfirmNotice ? '' : '请先阅读至底部'"
                aria-describedby="init-notice-hint"
                @click="showNext"
              >
                我已确认
              </el-button>
            </div>
          </div>
          <div v-if="page.showForm" class="anim-fade-up min-h-0 flex-1 overflow-y-auto pr-2">
            <el-form ref="formRef" :model="form" label-position="top" size="large" class="init-form">
              <el-form-item label="管理员密码">
                <el-input v-model="form.adminPassword" show-password placeholder="admin账号的默认密码" />
              </el-form-item>
              <el-form-item label="数据库类型">
                <el-select v-model="form.dbType" placeholder="请选择" class="w-full" @change="changeDB">
                  <el-option key="mysql" label="mysql" value="mysql" />
                  <el-option key="pgsql" label="pgsql" value="pgsql" />
                  <el-option key="oracle" label="oracle" value="oracle" />
                  <el-option key="mssql" label="mssql" value="mssql" />
                  <el-option key="sqlite" label="sqlite" value="sqlite" />
                </el-select>
              </el-form-item>
              <div class="flex gap-4">
                <el-form-item v-if="form.dbType !== 'sqlite'" label="host" class="flex-[7]">
                  <el-input v-model="form.host" placeholder="请输入数据库链接" />
                </el-form-item>
                <el-form-item v-if="form.dbType !== 'sqlite'" label="port" class="flex-[3]">
                  <el-input v-model="form.port" placeholder="请输入数据库端口" />
                </el-form-item>
              </div>
              <div class="flex gap-4">
                <el-form-item v-if="form.dbType !== 'sqlite'" label="userName" class="flex-[7]">
                  <el-input v-model="form.userName" placeholder="请输入数据库用户名" />
                </el-form-item>
                <el-form-item label="dbName" class="flex-[3]">
                  <el-input v-model="form.dbName" placeholder="请输入数据库名称" />
                </el-form-item>
              </div>
              <el-form-item v-if="form.dbType !== 'sqlite'" label="password">
                <el-input v-model="form.password" placeholder="请输入数据库密码（没有则为空）" />
              </el-form-item>
              <el-form-item v-if="form.dbType === 'sqlite'" label="dbPath">
                <el-input v-model="form.dbPath" placeholder="请输入sqlite数据库文件存放路径" />
              </el-form-item>
              <el-form-item v-if="form.dbType === 'pgsql'" label="template">
                <el-input v-model="form.template" placeholder="请输入postgresql指定template" />
              </el-form-item>
              <el-form-item style="padding-top: 12px">
                <el-button type="primary" size="large" class="w-full" @click="onSubmit">
                  立即初始化
                </el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// @ts-ignore
import { initDB } from '@/api/initdb'
import { computed, nextTick, onBeforeUnmount, onMounted, reactive, ref } from 'vue'
import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
import { Bottom, Coin, Lock, Reading } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import EntryBrand from '@/components/entryBrand/index.vue'
import { getScrollGateHint, isScrollAtBottom } from './scroll-gate'

defineOptions({
  name: 'Init'
})

const router = useRouter()

const page = reactive({
  showReadme: false,
  showForm: false
})

const noticeScrollRef = ref(null)
const noticeContentRef = ref(null)
const canConfirmNotice = ref(false)
const noticeHint = computed(() => getScrollGateHint(canConfirmNotice.value))
let noticeResizeObserver

const updateNoticeReadState = () => {
  if (!noticeScrollRef.value) {
    canConfirmNotice.value = false
    return
  }
  canConfirmNotice.value = isScrollAtBottom(noticeScrollRef.value)
}

const disconnectNoticeObserver = () => {
  noticeResizeObserver?.disconnect()
  noticeResizeObserver = undefined
}

onMounted(async () => {
  await nextTick()
  updateNoticeReadState()

  noticeResizeObserver = new ResizeObserver(updateNoticeReadState)
  noticeResizeObserver.observe(noticeScrollRef.value)
  noticeResizeObserver.observe(noticeContentRef.value)
})

onBeforeUnmount(disconnectNoticeObserver)

const showNext = () => {
  if (!canConfirmNotice.value) return

  disconnectNoticeObserver()
  page.showReadme = false
  setTimeout(() => {
    page.showForm = true
  }, 20)
}

const goDoc = () => {
  window.open('https://www.gin-vue-admin.com/guide/start-quickly/env.html')
}

const out = ref(false)

const form = reactive({
  adminPassword: '123456',
  dbType: 'mysql',
  host: '127.0.0.1',
  port: '3306',
  userName: 'root',
  password: '',
  dbName: 'gva',
  dbPath: ''
})

const changeDB = (val) => {
  switch (val) {
    case 'mysql':
      Object.assign(form, {
        adminPassword: '123456',
        reAdminPassword: '',
        dbType: 'mysql',
        host: '127.0.0.1',
        port: '3306',
        userName: 'root',
        password: '',
        dbName: 'gva',
        dbPath: ''
      })
      break
    case 'pgsql':
      Object.assign(form, {
        adminPassword: '123456',
        dbType: 'pgsql',
        host: '127.0.0.1',
        port: '5432',
        userName: 'postgres',
        password: '',
        dbName: 'gva',
        dbPath: '',
        template: 'template0'
      })
      break
    case 'oracle':
      Object.assign(form, {
        adminPassword: '123456',
        dbType: 'oracle',
        host: '127.0.0.1',
        port: '1521',
        userName: 'oracle',
        password: '',
        dbName: 'gva',
        dbPath: ''
      })
      break
    case 'mssql':
      Object.assign(form, {
        adminPassword: '123456',
        dbType: 'mssql',
        host: '127.0.0.1',
        port: '1433',
        userName: 'mssql',
        password: '',
        dbName: 'gva',
        dbPath: ''
      })
      break
    case 'sqlite':
      Object.assign(form, {
        adminPassword: '123456',
        dbType: 'sqlite',
        host: '',
        port: '',
        userName: '',
        password: '',
        dbName: 'gva',
        dbPath: ''
      })
      break
    default:
      Object.assign(form, {
        adminPassword: '123456',
        dbType: 'mysql',
        host: '127.0.0.1',
        port: '3306',
        userName: 'root',
        password: '',
        dbName: 'gva',
        dbPath: ''
      })
  }
}
const onSubmit = async () => {
  if (form.adminPassword.length < 6) {
    ElMessage({
      type: 'error',
      message: '密码长度不能小于6位'
    })
    return
  }

  const loading = ElLoading.service({
    lock: true,
    text: '正在初始化数据库，请稍候',
    spinner: 'loading',
    background: 'rgba(0, 0, 0, 0.7)'
  })
  try {
    const res = await initDB(form)
    if (res.code === 0) {
      out.value = true
      ElMessage({
        type: 'success',
        message: res.msg
      })

      // 显示AI助手配置提示弹窗
      ElMessageBox.confirm(
        '已经完成基础数据库初始化！建议先进行编辑器AI助手配置，以获得更好的开发体验。',
        '配置完成',
        {
          confirmButtonText: '查看AI配置文档',
          cancelButtonText: '稍后配置',
          type: 'success',
          center: true
        }
      ).then(() => {
        // 点击确认按钮，打开AI配置文档
        window.open('https://www.gin-vue-admin.com/guide/server/mcp.html', '_blank')
        router.push({ name: 'Login' })
      }).catch(() => {
        // 点击取消按钮或关闭弹窗，直接跳转到登录页
        router.push({ name: 'Login' })
      })
    }
    loading.close()
  } catch (_) {
    loading.close()
  }
}
</script>

<style lang="scss" scoped>
.init-layout {
  height: 100vh;
  height: 100dvh;
}

/* 右侧蓝色面板的左边缘斜切：顶部靠右、底部靠左，蓝色区上窄下宽 */
.banner-oblique {
  clip-path: polygon(22% 0, 100% 0, 100% 100%, 4% 100%);
}

/* 镂空按钮：常态透明底；hover 用浅色底 + 主色文字表现（不再变白字实心） */
.btn-hollow {
  --el-button-bg-color: transparent;
  --el-button-hover-bg-color: var(--el-color-primary-light-9);
  --el-button-hover-text-color: var(--el-color-primary);
  --el-button-hover-border-color: var(--el-color-primary);
  --el-button-active-bg-color: var(--el-color-primary-light-8);
  --el-button-active-text-color: var(--el-color-primary);
  --el-button-active-border-color: var(--el-color-primary);
}

/* 全页主题色统一为 #2264f2：覆盖 Element Plus 主色及其派生深浅色，
     让实心/镂空按钮、输入框聚焦、下拉选中等所有主色区域保持一致 */
#userLayout {
  --el-color-primary: #2264f2;
  --el-color-primary-light-3: #6493f6;
  --el-color-primary-light-5: #91b2f9;
  --el-color-primary-light-7: #bdd1fb;
  --el-color-primary-light-8: #d3e0fc;
  --el-color-primary-light-9: #e9f0fe;
  --el-color-primary-dark-2: #1b50c2;
}

/* 初始化表单：非 2K 屏（<2560px）每行 margin-bottom 收窄到 10px；
     2K 及以上保持 Element Plus 默认间距 */
@media (max-width: 2559.98px) {
  .init-form :deep(.el-form-item) {
    margin-bottom: 10px;
  }
}

/* 淡入上浮：替换原来的飞入动画 */
.anim-fade-up {
  animation: anim-fade-up 0.45s cubic-bezier(0.22, 0.61, 0.36, 1) both;
}

/* 须知分组：沿用同一动画，按组递增延迟逐组浮现 */
.notice-group {
  animation: anim-fade-up 0.45s cubic-bezier(0.22, 0.61, 0.36, 1) both;
}

@for $i from 1 through 3 {
  .notice-group:nth-child(#{$i}) {
    animation-delay: #{($i - 1) * 60}ms;
  }
}

@keyframes anim-fade-up {
  0% {
    opacity: 0;
    transform: translateY(18px);
  }

  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 750px) {
  .form {
    width: 94vw !important;
    padding: 0;
  }
}
</style>
