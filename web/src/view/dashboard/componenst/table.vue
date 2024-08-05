<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" :label="t('view.dashboard.components.table.ranking')" width="80" align="center" />
      <el-table-column prop="title" :label="t('view.dashboard.components.table.contentTitle')" show-overflow-tooltip />
      <el-table-column prop="click_num" :label="t('view.dashboard.components.table.attention')" width="100" />
      <el-table-column prop="hot" label="热度值" width="100"  />
    </el-table>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const tableData = [
  {
    ranking: 1,
    title : t('view.dashboard.components.table.content1'),
    click_num : 523,
    hot : 263
  },
  {
    ranking: 2,
    title : t('view.dashboard.components.table.content2'),
    click_num : 416,
    hot : 223
  },
  {
    ranking: 3,
    title : t('view.dashboard.components.table.content3'),
    click_num : 337,
    hot : 176
  },
  {
    ranking: 4,
    title : t('view.dashboard.components.table.content4'),
    click_num : 292,
    hot : 145
  },
  {
    ranking: 5,
    title : t('view.dashboard.components.table.content5'),
    click_num : 173,
    hot : 110
  },
]
</script>

<style scoped lang="scss">

</style>
