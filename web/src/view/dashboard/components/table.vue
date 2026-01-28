<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" label="排名" width="80" align="center" />
      <el-table-column prop="message" label="更新内容" show-overflow-tooltip />
      <el-table-column prop="author" label="提交人" width="140" />
      <el-table-column prop="date" label="时间" width="180" />
    </el-table>
  </div>
</template>

<script setup>
  import { formatTimeToStr } from '@/utils/date'
  import { ref, onMounted } from 'vue'
  import axios from 'axios'

  const service = axios.create()

  const tableData = ref([])

  const Commits =(page) => {
   return service({
    url:
      'https://api.github.com/repos/flipped-aurora/gin-vue-admin/commits?page=' +
      page,
    method: 'get'
  })
}

  const loadCommits = async () => {
    const { data } = await Commits(1)
    tableData.value = data.slice(0, 5).map((item, index) => {
      return {
        ranking: index + 1,
        message: item.commit.message,
        author: item.commit.author.name,
        date: formatTimeToStr(item.commit.author.date, 'yyyy-MM-dd hh:mm:ss')
      }
    })
  }

  onMounted(() => {
    loadCommits()
  })
</script>

<style scoped lang="scss"></style>
