<template>
  <div class="commit-table text-slate-700 dark:text-slate-400 dark:bg-slate-900">
    <div class="commit-table-title">
      更新日志
    </div>
    <div class="log">
      <div
        v-for="(item,key) in dataTimeline"
        :key="key"
        class="flex mt-2 leading-7"
      >
        <div class="flex items-center justify-between line-clamp-1">
          <span
            class="rounded-full w-2 h-2 bg-gray-300 dark:bg-gray-600 text-xs text-gray-300 dark:text-gray-600 mr-1"
            :class="key<3&&'bg-blue-300 dark:bg-blue-700'"
          />
        </div>
        <div class="flex-1 w-20">
          <div class="overflow-ellipsis overflow-hidden whitespace-nowrap">
            {{ item.message }}
          </div>
        </div>
        <div class="ml-1 flex-none w-24">
          {{ item.from }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Commits } from '@/api/github'
import { formatTimeToStr } from '@/utils/date.js'
import { ref } from 'vue'


defineOptions({
  name: 'DashboardTable',
})

const loading = ref(true)
const dataTimeline = ref([])

const loadCommits = () => {
  Commits(0).then(({ data }) => {
    loading.value = false
    data.forEach((element, index) => {
      if (element.commit.message && index < 10) {
        dataTimeline.value.push({
          from: formatTimeToStr(element.commit.author.date, 'yyyy-MM-dd'),
          title: element.commit.author.name,
          showDayAndMonth: true,
          message: element.commit.message,
        })
      }
    })
  })
}

loadCommits()
</script>

<style lang="scss" scoped>
.commit-table{
    height: 400px;
    &-title{
        font-weight: 600;
        margin-bottom: 12px;
    }
}
</style>
