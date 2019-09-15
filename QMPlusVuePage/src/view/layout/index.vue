<template>
  <el-container class="layout-cont">
    <el-header class="header-cont">顶部</el-header>
    <el-container>
      <el-aside class="main-cont main-left">
        <Aside />
      </el-aside>
      <el-main class="main-cont main-right">
        <transition name="el-fade-in-linear" mode="out-in">
          <router-view></router-view>
        </transition>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import Aside from '@/view/layout/aside'
export default {
  name: 'Layout',
  components: {
    Aside
  }
}
</script>

<style lang="scss">
$headerHigh: 52px;
$mainHight: calc(100vh - 52px);
.layout-cont {
  .header-cont {
    height: $headerHigh !important;
    background: palevioletred;
  }
  .main-cont {
    .el-menu-vertical {
      min-height: $mainHight !important;
    }
    &::-webkit-scrollbar {
      display: none;
    }
    &.main-left {
      width: auto !important;
    }
    background: blueviolet;
    height: $mainHight !important;
  }
}
</style>