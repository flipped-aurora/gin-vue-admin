<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->
<template>
  <div class="mt-8 w-full">
    <div class="grid grid-cols-2 md:grid-cols-3 3xl:grid-cols-4">
      <div
        v-for="(item, index) in shortcuts"
        :key="index"
        class="flex flex-col items-center mb-3 group cursor-pointer"
        @click="toPath(item)"
      >
        <div
          class="w-8 h-8 rounded bg-gray-200 dark:bg-slate-500 flex items-center justify-center group-hover:bg-blue-400 group-hover:text-white"
        >
          <el-icon><component :is="item.icon" /></el-icon>
        </div>
        <div class="text-xs mt-2 text-gray-700 dark:text-gray-300">
          {{ item.title }}
        </div>
      </div>
    </div>
    <div class="grid grid-cols-2 md:grid-cols-3 3xl:grid-cols-4 mt-8">
      <div
        v-for="(item, index) in recentVisits"
        :key="index"
        class="flex flex-col items-center mb-3 group cursor-pointer"
        @click="openLink(item)"
      >
        <div
          class="w-8 h-8 rounded bg-gray-200 dark:bg-slate-500 flex items-center justify-center group-hover:bg-blue-400 group-hover:text-white"
        >
          <el-icon><component :is="item.icon" /></el-icon>
        </div>
        <div class="text-xs mt-2 text-gray-700 dark:text-gray-300">
          {{ item.title }}
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
  import {
    Menu,
    Link,
    User,
    Service,
    Document,
    Reading,
    Files,
    Memo
  } from '@element-plus/icons-vue'
  import { useRouter } from 'vue-router'
  const router = useRouter()

  const toPath = (item) => {
    router.push({ name: item.path })
  }

  const openLink = (item) => {
    window.open(item.path, '_blank')
  }
  const shortcuts = [
    {
      icon: Menu,
      title: '菜单管理',
      path: 'menu'
    },
    {
      icon: Link,
      title: 'API管理',
      path: 'api'
    },
    {
      icon: Service,
      title: '角色管理',
      path: 'authority'
    },
    {
      icon: User,
      title: '用户管理',
      path: 'user'
    },
    {
      icon: Files,
      title: '自动化包',
      path: 'autoPkg'
    },
    {
      icon: Memo,
      title: '自动代码',
      path: 'autoCode'
    }
  ]

  const recentVisits = [
    {
      icon: Reading,
      title: '授权购买',
      path: 'https://gin-vue-admin.com/empower/index.html'
    },
    {
      icon: Document,
      title: '插件市场',
      path: 'https://plugin.gin-vue-admin.com/#/layout/home'
    }
  ]
</script>

<style scoped lang="scss"></style>
