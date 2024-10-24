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
      <el-tooltip effect="light" :content="item.title" placement="top">
        <div class="text-xs text-gray-700 dark:text-gray-300 line-clamp-1">
          {{ item.title }}
        </div>
      </el-tooltip>
    </div>
  </el-scrollbar>
</template>

<script setup>
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
  const notices = [
    {
      type : 'primary',
      typeTitle : t('view.dashboard.announcement'),
      title : t('view.dashboard.noticeList.notice1'),
    },
    {
      type : 'success',
      typeTitle : t('view.dashboard.notification'),
      title : t('view.dashboard.noticeList.notice2'),
    },
    {
      type : 'warning',
      typeTitle : t('general.warning'),
      title : t('view.dashboard.noticeList.notice3'),
    },
    {
      type : 'danger',
      typeTitle : t('view.dashboard.violation'),
      title : t('view.dashboard.noticeList.notice4'),
    },
    {
      type : 'info',
      typeTitle : t('view.dashboard.information'),
      title : t('view.dashboard.noticeList.notice5'),
    },
    {
      type : 'primary',
      typeTitle : t('view.dashboard.announcement'),
      title : t('view.dashboard.noticeList.notice6'),
    },
    {
      type : 'success',
      typeTitle : t('view.dashboard.notification'),
      title : t('view.dashboard.noticeList.notice7'),
    },
    {
      type : 'warning',
      typeTitle : t('general.warning'),
      title : t('view.dashboard.noticeList.notice8'),
    },
    {
      type : 'danger',
      typeTitle : t('view.dashboard.error'),
      title : t('view.dashboard.noticeList.notice9'),
    },
    {
      type : 'info',
      typeTitle : t('view.dashboard.information'),
      title : t('view.dashboard.noticeList.notice10'),
    }
  ]
</script>

<style scoped lang="scss">

</style>
