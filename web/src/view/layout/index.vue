<template>
  <el-container class="layout-cont">
    <el-container :class="[isSider?'openside':'hideside',isMobile ? 'mobile': '']">
      <el-row :class="[isShadowBg?'shadowBg':'']" @click="changeShadow()" />
      <el-aside class="main-cont main-left gva-aside">
        <div class="tilte" :style="{background: backgroundColor}">
          <img alt class="logoimg" :src="$GIN_VUE_ADMIN.appLogo">
          <div v-if="isSider" class="tit-text" :style="{color:textColor}">{{ $GIN_VUE_ADMIN.appName }}</div>
        </div>
        <Aside class="aside" />
      </el-aside>
      <!-- 分块滑动功能 -->
      <el-main class="main-cont main-right">
        <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
          <div
            :style="{width: `calc(100% - ${isMobile?'0px':isCollapse?'54px':'220px'})`}"
            class="topfix"
          >
            <el-row>
              <el-col>
                <el-header class="header-cont">
                  <el-row class="pd-0">
                    <el-col :xs="2" :lg="1" :md="1" :sm="1" :xl="1" style="z-index:100">
                      <div class="menu-total" @click="totalCollapse">
                        <div v-if="isCollapse" class="gvaIcon gvaIcon-arrow-double-right" />
                        <div v-else class="gvaIcon gvaIcon-arrow-double-left" />
                      </div>
                    </el-col>
                    <el-col :xs="10" :lg="14" :md="14" :sm="9" :xl="14" :pull="1">
                      <el-breadcrumb class="breadcrumb">
                        <el-breadcrumb-item
                          v-for="item in matched.slice(1,matched.length)"
                          :key="item.path"
                        >{{ fmtTitle(item.meta.title,route) }}</el-breadcrumb-item>
                      </el-breadcrumb>
                    </el-col>
                    <el-col :xs="12" :lg="9" :md="9" :sm="14" :xl="9">
                      <div class="right-box">
                        <Search />
                        <el-dropdown>
                          <div class="dp-flex justify-content-center align-items height-full width-full">
                            <span class="header-avatar" style="cursor: pointer">
                              <CustomPic />
                              <span style="margin-left: 5px">{{ userStore.userInfo.nickName }}</span>
                              <el-icon>
                                <arrow-down />
                              </el-icon>
                            </span>
                          </div>
                          <template #dropdown>
                            <el-dropdown-menu class="dropdown-group">
                              <el-dropdown-item>
                                <span style="font-weight: 600;">
                                  当前角色：{{ userStore.userInfo.authority.authorityName }}
                                </span>
                              </el-dropdown-item>
                              <template v-if="userStore.userInfo.authorities">
                                <el-dropdown-item v-for="item in userStore.userInfo.authorities.filter(i=>i.authorityId!==userStore.userInfo.authorityId)" :key="item.authorityId" @click="changeUserAuth(item.authorityId)">
                                  <span>
                                    切换为：{{ item.authorityName }}
                                  </span>
                                </el-dropdown-item>
                              </template>
                              <el-dropdown-item icon="avatar" @click="toPerson">个人信息</el-dropdown-item>
                              <el-dropdown-item icon="reading-lamp" @click="userStore.LoginOut">登 出</el-dropdown-item>
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
          v-loading="loadingFlag"
          element-loading-text="正在加载中"
          class="admin-box"
        >
          <div>
            <transition mode="out-in" name="el-fade-in-linear">
              <keep-alive :include="routerStore.keepAliveRouters">
                <component :is="Component" />
              </keep-alive>
            </transition>
          </div>
        </router-view>
        <BottomInfo />
        <setting />
      </el-main>
    </el-container>

  </el-container>
</template>

<script>
export default {
  name: 'Layout',
}
</script>

<script setup>
import Aside from '@/view/layout/aside/index.vue'
import HistoryComponent from '@/view/layout/aside/historyComponent/history.vue'
import Search from '@/view/layout/search/search.vue'
import BottomInfo from '@/view/layout/bottomInfo/bottomInfo.vue'
import CustomPic from '@/components/customPic/index.vue'
import Setting from './setting/index.vue'
import { setUserAuthority } from '@/api/user'
import { emitter } from '@/utils/bus.js'
import { computed, ref, onMounted, nextTick } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { useRouterStore } from '@/pinia/modules/router'
import { fmtTitle } from '@/utils/fmtRouterTitle'

const router = useRouter()
const route = useRoute()
const routerStore = useRouterStore()
// 三种窗口适配
const isCollapse = ref(false)
const isSider = ref(true)
const isMobile = ref(false)

const initPage = () => {
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

const loadingFlag = ref(false)
onMounted(() => {
  // 挂载一些通用的事件
  emitter.emit('collapse', isCollapse.value)
  emitter.emit('mobile', isMobile.value)
  emitter.on('reload', reload)
  emitter.on('showLoading', () => {
    loadingFlag.value = true
  })
  emitter.on('closeLoading', () => {
    loadingFlag.value = false
  })
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
    emitter.emit('closeAllPage')
    setTimeout(() => {
      window.location.reload()
    }, 50)
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
@import '@/style/mobile.scss';

.dark{
  background-color: #191a23 !important;
  color: #fff !important;
}
.light{
  background-color: #fff !important;
  color: #000 !important;
}
</style>
