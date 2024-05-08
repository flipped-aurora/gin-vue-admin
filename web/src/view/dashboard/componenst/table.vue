<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" label="排名" width="80" align="center" />
      <el-table-column prop="title" label="内容标题" show-overflow-tooltip />
      <el-table-column prop="click_num" label="点击量" width="100" />
      <el-table-column prop="daily_increase" label="日涨幅" width="100"  />
    </el-table>
  </div>
</template>

<script setup>
const tableData = [
  {
    ranking: 1,
    title : "经济日报：财政政策要精准提升效能",
    click_num : 123,
    daily_increase : 0.1
  },
  {
    ranking: 2,
    title : "经济日报：财政政策要精准提升效能",
    click_num : 123,
    daily_increase : 0.1
  },
  {
    ranking: 3,
    title : "经济日报：财政政策要精准提升效能",
    click_num : 123,
    daily_increase : 0.1
  },
  {
    ranking: 4,
    title : "经济日报：财政政策要精准提升效能",
    click_num : 123,
    daily_increase : 0.1
  },
  {
    ranking: 5,
    title : "经济日报：财政政策要精准提升效能",
    click_num : 123,
    daily_increase : 0.1
  },
]
</script>

<style scoped lang="scss">

</style>
