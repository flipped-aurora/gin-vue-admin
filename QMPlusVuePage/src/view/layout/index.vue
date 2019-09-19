<template>
  <el-container class="layout-cont">
    <el-header class="header-cont"></el-header>
    <el-container>
      <el-aside class="main-cont main-left">
        <Aside class="aside" />
      </el-aside>
      <!-- 分块滑动功能 -->
      <vue-scroll>
        <el-main class="main-cont main-right">
          <!-- 当前面包屑用路由自动生成可根据需求修改 -->
          <el-breadcrumb class="breadcrumb" separator-class="el-icon-arrow-right">
            <el-breadcrumb-item
              :key="item.path"
              :to="{ path: item.path }"
              v-for="item in matched.slice(1,matched.length)"
            >{{item.meta.title}}</el-breadcrumb-item>
          </el-breadcrumb>
          <transition mode="out-in" name="el-fade-in-linear">
            <router-view></router-view>
          </transition>
        </el-main>
      </vue-scroll>
    </el-container>
  </el-container>
</template>

<script>
import Aside from '@/view/layout/aside'
import vueScroll from 'vuescroll'
import 'vuescroll/dist/vuescroll.css'
export default {
  name: 'Layout',
  components: {
    Aside,
    vueScroll
  },
  methods: {
    totalCollapse() {
      this.$bus.emit('totalCollapse')
    }
  },
  computed: {
    title() {
      return this.$route.meta.title || '当前页面'
    },
    matched() {
      return this.$route.matched
    }
  }
}
</script>

<style lang="scss">
$headerHigh: 52px;
$mainHight: calc(100vh - 52px);
.layout-cont {
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    color: #eee;
    display: table-cell;
    vertical-align: middle;
  }
  .header-cont {
    height: $headerHigh !important;
    background: palevioletred;
  }
  .main-cont {
    .breadcrumb {
      line-height: 24px;
      padding: 6px;
      border-bottom: 1px solid #eee;
      margin-bottom: 6px;
    }
    &.el-main {
      padding: 0px 10px;
      margin: 0px 0px 0px 12px;
      background: #fff;
    }
    height: $mainHight !important;
    overflow: visible;
    position: relative;
    .menu-total {
      z-index: 5;
      position: absolute;
      top: 50%;
      margin-top: -10px;
      line-height: 20px;
      font-size: 20px;
      border: 0 solid #ffffff;
      border-radius: 50%;
      right: -10px;
      background: #fff;
    }
    .aside {
      background: #fff;
    }
    .el-menu-vertical {
      height: $mainHight !important;
      visibility: auto;
      &.noCollapse {
        width: 250px;
      }
    }
    &::-webkit-scrollbar {
      display: none;
    }
    &.main-left {
      width: auto !important;
    }
    background: blueviolet;
  }
}
</style>