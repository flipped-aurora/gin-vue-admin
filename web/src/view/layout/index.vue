<template>
  <el-container class="layout-cont">
    <el-container :class="[isSider?'openside':'hideside',isMobile ? 'mobile': '']">
      <el-row :class="[isShadowBg?'shadowBg':'']" @click.native="changeShadow()"></el-row>
      <el-aside class="main-cont main-left">
        <div class="tilte">
          <img alt class="logoimg" src="http://qmplusimg.henrongyi.top/1606920521logo.jpg" />
          <h2 class="tit-text" v-if="isSider">We-亳院</h2>
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
            <el-header class="header-cont">
              <div @click="totalCollapse" class="menu-total">
                <i class="el-icon-s-unfold" v-if="isCollapse"></i>
                <i class="el-icon-s-fold" v-else></i>
              </div>
              <el-breadcrumb class="breadcrumb" separator-class="el-icon-arrow-right">
                <el-breadcrumb-item
                  :key="item.path"
                  v-for="item in matched.slice(1,matched.length)"
                >{{item.meta.title}}</el-breadcrumb-item>
              </el-breadcrumb>
              <div class="fl-right right-box">
                <Search />
                <Screenfull class="screenfull"></Screenfull>
                <el-dropdown>
                  <span class="header-avatar">
                   <CustomPic/>
                    <span style="margin-left: 5px">{{userInfo.nickName}}</span>
                    <i class="el-icon-arrow-down"></i>
                  </span>
                  <el-dropdown-menu class="dropdown-group" slot="dropdown">
                    <el-dropdown-item>
                      <span>
                        更多信息
                        <el-badge is-dot />
                      </span>
                    </el-dropdown-item>
                    <el-dropdown-item @click.native="toPerson" icon="el-icon-s-custom">个人信息</el-dropdown-item>
                    <el-dropdown-item @click.native="LoginOut" icon="el-icon-table-lamp">登 出</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </div>
            </el-header>
            <!-- 当前面包屑用路由自动生成可根据需求修改 -->
            <!--
            :to="{ path: item.path }" 暂时注释不用-->
            <HistoryComponent />
          </div>
        </transition>
        <transition mode="out-in" name="el-fade-in-linear">
          <keep-alive>
            <router-view  v-loading="loadingFlag"  element-loading-text="正在加载中" class="admin-box" v-if="$route.meta.keepAlive"></router-view>
          </keep-alive>
        </transition>
        <transition mode="out-in" name="el-fade-in-linear">
          <router-view   v-loading="loadingFlag"  element-loading-text="正在加载中" class="admin-box" v-if="!$route.meta.keepAlive &&  isRouterAlive"></router-view>
        </transition>
       <BottomInfo />
      </el-main>
    </el-container>
   
  </el-container>
</template>

<script>
import Aside from '@/view/layout/aside'
import HistoryComponent from '@/view/layout/aside/historyComponent/history'
import Screenfull from '@/view/layout/screenfull'
import Search from '@/view/layout/search/search'
import BottomInfo from '@/view/layout/bottomInfo/bottomInfo'
import { mapGetters, mapActions } from 'vuex'
import CustomPic from '@/components/customPic'
export default {
  name: 'Layout',
  provide(){
    return {
      reload : this.reload
    }
  },
  data() {
    return {
      show: false,
      isCollapse: false,
      isSider: true,
      isMobile: false,
      isShadowBg: false,
      loadingFlag:false,
      value: '',
      isRouterAlive: true
    }
  },
  components: {
    Aside,
    HistoryComponent,
    Screenfull,
    Search,
    BottomInfo,
    CustomPic
  },
  methods: {
    ...mapActions('user', ['LoginOut']),
    totalCollapse() {
      this.isCollapse = !this.isCollapse
      this.isSider = !this.isCollapse
      this.isShadowBg = !this.isCollapse
      this.$bus.emit('collapse', this.isCollapse)
    },
    toPerson() {
      this.$router.push({ name: 'person' })
    },
    changeShadow() {
      this.isShadowBg = !this.isShadowBg
      this.isSider = !!this.isCollapse
      this.totalCollapse()
    },
    reload(){
     this.isRouterAlive = false
     this.$nextTick(function(){
       this.isRouterAlive = true
     })
   }
  },
  computed: {
    ...mapGetters('user', ['userInfo']),
    title() {
      return this.$route.meta.title || '当前页面'
    },
    matched() {
      return this.$route.matched
    }
  },
  mounted() {
    let screenWidth = document.body.clientWidth
    if (screenWidth < 1000) {
      this.isMobile = true
      this.isSider = false
      this.isCollapse = true
    } else if (screenWidth >= 1000 && screenWidth < 1200) {
      this.isMobile = false
      this.isSider = false
      this.isCollapse = true
    } else {
      this.isMobile = false
      this.isSider = true
      this.isCollapse = false
    }
    this.$bus.emit('collapse', this.isCollapse)
    this.$bus.emit('mobile', this.isMobile)
    this.$bus.on("showLoading",()=>{
      this.loadingFlag = true
    })
    this.$bus.on("closeLoading",()=>{
      this.loadingFlag = false
    })
    window.onresize = () => {
      return (() => {
        let screenWidth = document.body.clientWidth
        if (screenWidth < 1000) {
          this.isMobile = true
          this.isSider = false
          this.isCollapse = true
        } else if (screenWidth >= 1000 && screenWidth < 1200) {
          this.isMobile = false
          this.isSider = false
          this.isCollapse = true
        } else {
          this.isMobile = false
          this.isSider = true
          this.isCollapse = false
        }
        this.$bus.emit('collapse', this.isCollapse)
        this.$bus.emit('mobile', this.isMobile)
      })()
    }
  }
}
</script>

<style lang="scss">

</style>
