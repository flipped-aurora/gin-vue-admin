<template>
  <el-container class="layout-cont">
    <el-container>
      <el-aside class="main-cont main-left">
        <Aside class="aside" />
      </el-aside>
      <!-- 分块滑动功能 -->
      <el-main class="main-cont main-right">
        <el-header class="header-cont">
          <div @click="totalCollapse" class="menu-total">
            <i class="el-icon-s-unfold" v-if="isCollapse"></i>
            <i class="el-icon-s-fold" v-else></i>
          </div>
          <h1 class="admin-title">Gin-Vue-Admin</h1>
          <div class="fl-right right-box">
            <el-dropdown>
              <span class="el-dropdown-link">
                <img :src="userInfo.headerImg" height="30" width="30" />
                {{userInfo.nickName}}
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
        <el-breadcrumb class="breadcrumb" separator-class="el-icon-arrow-right">
          <el-breadcrumb-item
            :key="item.path"
            v-for="item in matched.slice(1,matched.length)"
          >{{item.meta.title}}</el-breadcrumb-item>
        </el-breadcrumb>
        <transition mode="out-in" name="el-fade-in-linear">
          <router-view class="admin-box"></router-view>
        </transition>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import Aside from '@/view/layout/aside'
import { mapGetters, mapActions } from 'vuex'
export default {
  name: 'Layout',
  data() {
    return {
      isCollapse: false
    }
  },
  components: {
    Aside
  },
  methods: {
    ...mapActions('user', ['LoginOut']),
    totalCollapse() {
      this.isCollapse = !this.isCollapse
      this.$bus.emit('totalCollapse')
    },
    toPerson() {
      this.$router.push({ name: 'person' })
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
$mainHight: 100vh;
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
    box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
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
      background: #fff;
    }
    height: $mainHight !important;
    overflow: visible;
    position: relative;
    .menu-total {
      // z-index: 5;
      // position: absolute;
      // top: 10px;
      // right: -35px;
      margin-left: -10px;
      float: left;
      margin-top: 10px;
      width: 30px;
      height: 30px;
      line-height: 30px;
      font-size: 30px;
      // border: 0 solid #ffffff;
      // border-radius: 50%;
      // background: #fff;
    }
    .aside {
      overflow: auto;
      background: #fff;
      &::-webkit-scrollbar {
        display: none;
      }
    }

    .el-menu-vertical {
      height: 100vh !important;
      visibility: auto;
      &:not(.el-menu--collapse) {
        width: 220px;
      }
    }

    &::-webkit-scrollbar {
      display: none;
    }
    &.main-left {
      width: auto !important;
    }
    &.main-right {
      .admin-title {
        float: left;
        font-size: 16px;
        vertical-align: middle;
        margin-left: 20px;
        img {
          vertical-align: middle;
        }
        &.collapse {
          width: 53px;
        }
      }
    }
  }
}
</style>