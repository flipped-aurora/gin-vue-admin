<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <el-scrollbar>
    <div v-for="(item , index) in notices" :key="index" class="flex items-center mb-1.5 gap-3">
      <el-tag :type="item.type" size="small">
        {{ item.typeTitle }}
      </el-tag>
      <div class="text-xs text-gray-700 dark:text-gray-300 line-clamp-1">
        {{ item.title }}
      </div>
    </div>
  </el-scrollbar>
</template>

<script setup>
  const notices = [
    {
      type : 'primary',
      typeTitle : "公告",
      title : "公告1公告1公告1公告1公告1公告1公告1公告1公告1公告1公告1公告1",
    },
    {
      type : 'success',
      typeTitle : "通知",
      title : "通知1通知1通知1通知1通知1通知1通知1通知1通知1通知1通知1通知1通知1通知1",
    },
    {
      type : 'warning',
      typeTitle : "警告",
      title : "警告1警告1警告1警告1警告1警告1警告1警告1警告1警告1警告1警告1警告1警告1",
    },
    {
      type : 'danger',
      typeTitle : "错误",
      title : "错误1错误1错误1错误1错误1错误1错误1错误1错误1错误1错误1错误1错误1",
    },
    {
      type : 'info',
      typeTitle : "信息",
      title : "信息1信息1信息1信息1信息1信息1信息1信息1信息1信息1信息1信息1信息1信息1信息1",
    },
    {
      type : 'primary',
      typeTitle : "公告",
      title : "公告2",
    },
    {
      type : 'success',
      typeTitle : "通知",
      title : "通知2",
    },
    {
      type : 'warning',
      typeTitle : "警告",
      title : "警告2",
    },
    {
      type : 'danger',
      typeTitle : "错误",
      title : "错误2",
    },
    {
      type : 'info',
      typeTitle : "信息",
      title : "信息2",
    }
  ]
</script>

<style scoped lang="scss">

</style>
