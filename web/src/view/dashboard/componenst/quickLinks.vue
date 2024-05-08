<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div class="mt-8 w-full">
    <div class="grid grid-cols-2 md:grid-cols-3 3xl:grid-cols-4 border-b border-solid border-gray-200">
      <div v-for="(item, index ) in shortcuts" :key="index" class="flex flex-col items-center mb-3 group ">
        <div class="w-8 h-8 rounded bg-gray-200 dark:bg-slate-500 flex items-center justify-center group-hover:bg-blue-400 cursor-pointer group-hover:text-white">
          <el-icon><Message /></el-icon>
        </div>
        <div class="text-xs mt-2  text-gray-700 dark:text-gray-300">{{ item.title }}</div>
      </div>
    </div>
    <div class="grid grid-cols-2 md:grid-cols-3 3xl:grid-cols-4 mt-8">
      <div v-for="(item, index ) in recentVisits" :key="index" class="flex flex-col items-center mb-3 group">
        <div class="w-8 h-8 rounded bg-gray-200 dark:bg-slate-500 flex items-center justify-center group-hover:bg-blue-400 cursor-pointer group-hover:text-white">
          <el-icon><Message /></el-icon>
        </div>
        <div class="text-xs mt-2  text-gray-700 dark:text-gray-300">{{ item.title }}</div>
      </div>
    </div>
  </div>
</template>

<script setup>
  const shortcuts = [
    {
      icon : "",
      title : "快捷链接1",
    },
    {
      icon : "",
      title : "快捷链接2",
    },
    {
      icon : "",
      title : "快捷链接3",
    },
    {
      icon : "",
      title : "快捷链接4",
    },
    {
      icon : "",
      title : "快捷链接5",
    },
    {
      icon : "",
      title : "快捷链接6",
    }
  ]

  const recentVisits = [
    {
      icon : "",
      title : "快捷链接1",
    },
    {
      icon : "",
      title : "快捷链接2",
    }
  ]
</script>

<style scoped lang="scss">

</style>
