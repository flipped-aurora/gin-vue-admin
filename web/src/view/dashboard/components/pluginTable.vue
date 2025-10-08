<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" label="排名" width="80" align="center" />
      <el-table-column prop="title" label="插件标题" show-overflow-tooltip>
        <template #default="{ row }">
          <a class="text-active" :href="row.link" target="_blank">{{
            row.title
          }}</a>
        </template>
      </el-table-column>
      <el-table-column prop="click_num" label="关注度" width="100" />
      <el-table-column prop="hot" label="热度值" width="100" />
    </el-table>
  </div>
</template>

<script setup>
  const tableData = [
    {
      ranking: 1,
      title: '组织管理插件:更方便管理组织，分配资源权限。',
      click_num: 523,
      hot: 263,
      link: 'https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=36'
    },
    {
      ranking: 2,
      title:
        'Kubernetes容器管理:，Kubernetes 原生资源管理，提供炫酷的YAML 编辑，Pod 终端，方便运维兄弟管理k8s资源',
      click_num: 416,
      hot: 223,
      link: 'https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=42'
    },
    {
      ranking: 3,
      title:
        '定时任务配置化管理:本插件用于对系统内部的定时任务进行配置化管理，可以配置自定义的函数和HTTP，可以配置cron和remark等等',
      click_num: 337,
      hot: 176,
      link: 'https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=67'
    },
    {
      ranking: 4,
      title:
        '官网CMS系统：基于Gin-Vue-Admin 和 插件市场客户端开发基座开发的企业官网类（cms）系统',
      click_num: 292,
      hot: 145,
      link: 'https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=69'
    },
    {
      ranking: 5,
      title: '微信支付插件：提供扫码支付功能（需自行对接业务）',
      click_num: 173,
      hot: 110,
      link: 'https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=28'
    }
  ]
</script>

<style scoped lang="scss"></style>
