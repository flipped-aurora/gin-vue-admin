<template>
  <el-container class="layout-cont">
    <el-container :class="[isSider?'openside':'hideside',isMobile ? 'mobile': '']">
      <el-row
        :class="[isShadowBg && isMobile?'bg-black opacity-30 w-full h-full absolute top-0 left-0 z-[1001]':'']"
        @click="changeShadow()"
      />
      <el-aside
        class="main-cont gva-aside"
        :style="{width:asideWidth()}"
      >
        <div
          class="min-h-[60px] text-center transition-all duration-300 flex items-center justify-center gap-2"
          :style="{background: backgroundColor}"
        >
          <img
            alt
            class="w-9 h-9 p-1 bg-white rounded-full"
            :src="$GIN_VUE_ADMIN.appLogo"
          >
          <div
            v-if="isSider"
            class="inline-flex font-bold text-2xl"
            :style="{color:textColor}"
          >{{ $GIN_VUE_ADMIN.appName }}</div>
        </div>
        <Aside class="aside" />
      </el-aside>
      <!-- 分块滑动功能 -->
      <el-main class="main-cont main-right">
        <transition
          :duration="{ enter: 800, leave: 100 }"
          mode="out-in"
          name="el-fade-in-linear"
        >
          <div
            :style="{width: `calc(100% - ${getAsideWidth()})`}"
            class="fixed top-0 box-border z-50"
          >
            <el-row>
              <el-col>
                <el-header class="header-cont">
                  <el-row class="p-0 h-full">
                    <el-col
                      :xs="2"
                      :lg="1"
                      :md="1"
                      :sm="1"
                      :xl="1"
                      class="z-50 flex items-center pl-3"
                    >
                      <div
                        class="text-black cursor-pointer text-lg leading-5"
                        @click="totalCollapse"
                      >
                        <div
                          v-if="isCollapse"
                          class="gvaIcon gvaIcon-arrow-double-right"
                        />
                        <div
                          v-else
                          class="gvaIcon gvaIcon-arrow-double-left"
                        />
                      </div>
                    </el-col>
                    <el-col
                      :xs="10"
                      :lg="14"
                      :md="14"
                      :sm="9"
                      :xl="14"
                      :pull="1"
                      class="flex items-center"
                    >
                      <!-- 修改为手机端不显示顶部标签 -->
                      <el-breadcrumb
                        v-show="!isMobile"
                        class="breadcrumb"
                      >
                        <el-breadcrumb-item
                          v-for="item in matched.slice(1,matched.length)"
                          :key="item.path"
                        >{{ fmtTitle(item.meta.title,route) }}</el-breadcrumb-item>
                      </el-breadcrumb>
                    </el-col>
                    <el-col
                      :xs="12"
                      :lg="9"
                      :md="9"
                      :sm="14"
                      :xl="9"
                      class="flex items-center justify-end"
                    >
                      <div class="mr-1.5 flex items-center">
                        <Search />
                        <el-dropdown>
                          <div class="flex justify-center items-center h-full w-full">
                            <span class="cursor-pointer flex justify-center items-center">
                              <CustomPic />
                              <span
                                v-show="!isMobile"
                                style="margin-left: 5px"
                              >{{ userStore.userInfo.nickName }}</span>
                              <el-icon>
                                <arrow-down />
                              </el-icon>
                            </span>
                          </div>
                          <template #dropdown>
                            <el-dropdown-menu>
                              <el-dropdown-item>
                                <span class="font-bold">
                                  当前角色：{{ userStore.userInfo.authority.authorityName }}
                                </span>
                              </el-dropdown-item>
                              <template v-if="userStore.userInfo.authorities">
                                <el-dropdown-item
                                  v-for="item in userStore.userInfo.authorities.filter(i=>i.authorityId!==userStore.userInfo.authorityId)"
                                  :key="item.authorityId"
                                  @click="changeUserAuth(item.authorityId)"
                                >
                                  <span>
                                    切换为：{{ item.authorityName }}
                                  </span>
                                </el-dropdown-item>
                              </template>
                              <el-dropdown-item icon="avatar">
                                <div
                                  class="command-box"
                                  style="display: flex"
                                  @click="handleCommand"
                                >
                                  <div>指令菜单</div>
                                  <div style="margin-left: 8px">
                                    <span class="button">{{ first }}</span>
                                    +
                                    <span class="button">K</span>
                                  </div>
                                </div>
                              </el-dropdown-item>
                              <el-dropdown-item
                                icon="avatar"
                                @click="toPerson"
                              >个人信息</el-dropdown-item>
                              <el-dropdown-item
                                icon="reading-lamp"
                                @click="userStore.LoginOut"
                              >登 出</el-dropdown-item>
                            </el-dropdown-menu>
                          </template>
                        </el-dropdown>
                      </div>
                    </el-col>
                  </el-row>
                </el-header>
              </el-col>
            </el-row>
            <!-- 当前面包屑用路由自动生成可根据需求修改 -->
            <!--
            :to="{ path: item.path }" 暂时注释不用-->
            <HistoryComponent ref="layoutHistoryComponent" />
          </div>
        </transition>
        <router-view
          v-if="reloadFlag"
          v-slot="{ Component }"
          class="admin-box"
        >
          <div
            id="gva-base-load-dom"
          >
            <transition
              mode="out-in"
              name="el-fade-in-linear"
            >
              <keep-alive :include="routerStore.keepAliveRouters">
                <component :is="Component" />
              </keep-alive>
            </transition>
          </div>
        </router-view>
        <BottomInfo />
        <setting />
        <CommandMenu ref="command" />
      </el-main>
    </el-container>

  </el-container>
</template>

<script setup>
import Aside from '@/view/layout/aside/index.vue'
import HistoryComponent from '@/view/layout/aside/historyComponent/history.vue'
import Search from '@/view/layout/search/search.vue'
import BottomInfo from '@/view/layout/bottomInfo/bottomInfo.vue'
import CustomPic from '@/components/customPic/index.vue'
import CommandMenu from '@/components/commandMenu/index.vue'
import Setting from './setting/index.vue'
import { setUserAuthority } from '@/api/user'
import { emitter } from '@/utils/bus.js'
import { computed, ref, onMounted, nextTick } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useRouterStore } from '@/pinia/modules/router'
import { fmtTitle } from '@/utils/fmtRouterTitle'
import { useUserStore } from '@/pinia/modules/user'

defineOptions({
  name: 'Layout',
})

const router = useRouter()
const route = useRoute()
const routerStore = useRouterStore()
// 三种窗口适配
const isCollapse = ref(false)
const isSider = ref(true)
const isMobile = ref(false)

const first = ref('')
const initPage = () => {
  // 判断当前用户的操作系统
  if (window.localStorage.getItem('osType') === 'WIN') {
    first.value = 'Ctrl'
  } else {
    first.value = '⌘'
  }
  // 当用户同时按下ctrl和k键的时候
  const handleKeyDown = (e) => {
    if (e.ctrlKey && e.key === 'k') {
      // 阻止浏览器默认事件
      e.preventDefault()
      handleCommand()
    }
  }
  window.addEventListener('keydown', handleKeyDown)

  const screenWidth = document.body.clientWidth
  if (screenWidth < 1000) {
    isMobile.value = true
    isSider.value = false
    isCollapse.value = true
  } else if (screenWidth >= 1000 && screenWidth < 1200) {
    isMobile.value = false
    isSider.value = false
    isCollapse.value = true
  } else {
    isMobile.value = false
    isSider.value = true
    isCollapse.value = false
  }
}

initPage()

const command = ref()
const handleCommand = () => {
  command.value.open()
}

onMounted(() => {
  // 挂载一些通用的事件
  emitter.emit('collapse', isCollapse.value)
  emitter.emit('mobile', isMobile.value)
  emitter.on('reload', reload)
  window.onresize = () => {
    return (() => {
      initPage()
      emitter.emit('collapse', isCollapse.value)
      emitter.emit('mobile', isMobile.value)
    })()
  }
  if (userStore.loadingInstance) {
    userStore.loadingInstance.close()
  }
})

const userStore = useUserStore()

const asideWidth = () => {
  if (isMobile.value) {
    return isCollapse.value ? '0px' : '220px'
  }
  return isCollapse.value ? '54px' : '220px'
}

const getAsideWidth = () => {
  if (isMobile.value) return '0px'
  return isCollapse.value ? '54px' : '220px'
}

const textColor = computed(() => {
  if (userStore.sideMode === 'dark') {
    return '#fff'
  } else if (userStore.sideMode === 'light') {
    return '#191a23'
  } else {
    return userStore.baseColor
  }
})

const backgroundColor = computed(() => {
  if (userStore.sideMode === 'dark') {
    return '#191a23'
  } else if (userStore.sideMode === 'light') {
    return '#fff'
  } else {
    return userStore.sideMode
  }
})

const matched = computed(() => route.meta.matched)

const changeUserAuth = async(id) => {
  const res = await setUserAuthority({
    authorityId: id
  })
  if (res.code === 0) {
    window.sessionStorage.setItem('needCloseAll', 'true')
    window.location.reload()
  }
}

const reloadFlag = ref(true)
let reloadTimer = null
const reload = async() => {
  if (reloadTimer) {
    window.clearTimeout(reloadTimer)
  }
  reloadTimer = window.setTimeout(async() => {
    if (route.meta.keepAlive) {
      reloadFlag.value = false
      await nextTick()
      reloadFlag.value = true
    } else {
      const title = route.meta.title
      router.push({ name: 'Reload', params: { title }})
    }
  }, 400)
}

const isShadowBg = ref(false)
const totalCollapse = () => {
  isCollapse.value = !isCollapse.value
  isSider.value = !isCollapse.value
  isShadowBg.value = !isCollapse.value
  emitter.emit('collapse', isCollapse.value)
}

const toPerson = () => {
  router.push({ name: 'person' })
}
const changeShadow = () => {
  isShadowBg.value = !isShadowBg.value
  isSider.value = !!isCollapse.value
  totalCollapse()
}
</script>

<style lang="scss">
.button {
  font-size: 12px;
  color: #666;
  background:	rgb(250,250,250);
  width: 25px!important;
  padding: 4px 8px !important;
  border: 1px solid #eaeaea;
  margin-right: 4px;
  border-radius: 4px;
}
:deep .el-overlay {
  background-color: hsla(0,0%,100%,.9) !important;
}

</style>
