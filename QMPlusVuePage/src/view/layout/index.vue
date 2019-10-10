<template>
  <el-container class="layout-cont">
    <el-header class="header-cont">
      <h1 class="fl-left">QMPlus gin-vue-admin</h1>
      <div class="fl-right right-box">
        <el-dropdown>
          <span class="el-dropdown-link">
            <img :src="userInfo.headerImg" height="30" width="30" />
            {{userInfo.nickName}}
            <i class="el-icon-arrow-down"></i>
          </span>
          <el-dropdown-menu class="dropdown-group" slot="dropdown">
            <el-dropdown-item>
              <span>更多信息
                <el-badge is-dot />
              </span>
            </el-dropdown-item>
            <el-dropdown-item @click.native="LoginOut" icon="el-icon-table-lamp">登出</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </el-header>
    <el-container>
      <el-aside class="main-cont main-left">
        <Aside class="aside" />
      </el-aside>
      <!-- 分块滑动功能 -->
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
    </el-container>
  </el-container>
</template>

<script>
import Aside from '@/view/layout/aside'
import { mapGetters, mapMutations } from 'vuex'
export default {
  name: 'Layout',
  components: {
    Aside
  },
  methods: {
    ...mapMutations('user', ['LoginOut']),
    totalCollapse() {
      this.$bus.emit('totalCollapse')
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
  }
}
</script>

<style lang="scss">
$headerHigh: 52px;
$mainHight: calc(100vh - 52px);
.el-dropdown-link {
  cursor: pointer;
}
.dropdown-group {
  min-width: 100px;
}
.el-scrollbar__wrap {
  padding-bottom: 17px;
}
.layout-cont {
  .right-box {
    text-align: center;
    vertical-align: middle;
    img {
      vertical-align: middle;
      border: 1px solid #ccc;
      border-radius: 6px;
    }
  }
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    color: #eee;
    display: table-cell;
    vertical-align: middle;
  }
  .header-cont {
    height: $headerHigh !important;
    background: #fff;
    border-bottom: 1px solid #ccc;
    line-height: $headerHigh;
  }
  .main-cont {
    .breadcrumb {
      line-height: 24px;
      padding: 6px;
      border-bottom: 1px solid #eee;
      margin-bottom: 6px;
    }
    &.el-main {
      overflow: auto;
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
      overflow: auto;
      background: #fff;
      &::-webkit-scrollbar {
        display: none;
      }
    }

    .el-menu-vertical {
      height: calc(100vh - 67px) !important;
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