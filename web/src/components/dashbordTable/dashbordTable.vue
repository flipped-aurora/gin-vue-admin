<template>
  <div class="commit-table">
    <div class="commit-table-title">Gva 仓库commit 记录</div>
    <el-table v-loading="loading" :data="dataTimeline" style="width: 100%">
      <el-table-column prop="from" show-overflow-tooltip label="日期" width="180" />
      <el-table-column prop="title" show-overflow-tooltip label="推送者" width="120" />
      <el-table-column prop="message" show-overflow-tooltip label="commit 信息" />
    </el-table>
  </div>
</template>

<script>
import { Commits } from '@/api/github'
export default {
  data() {
    return {
      loading: true,
      dataTimeline: [],
    }
  },
  created() {
    this.loadCommits()
  },
  methods: {
    loadCommits() {
      Commits(0).then(({ data }) => {
        this.loading = false
        data.forEach((element, index) => {
          if (element.commit.message && index < 6) {
            this.dataTimeline.push({
              from: new Date(element.commit.author.date),
              title: element.commit.author.name,
              showDayAndMonth: true,
              message: element.commit.message,
            })
          }
        })
      })
    },
  }
}
</script>
<style lang="scss" scoped>
.commit-table{
    padding: 20px;
    background-color: #fff;
    height: 400px;
    &-title{
        color: rgb(56,137,206);
        font-size: 18px;
        margin-bottom: 20px;
    }
}
</style>
