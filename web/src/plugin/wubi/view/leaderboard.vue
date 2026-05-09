<template>
  <div class="wubi-leaderboard p-4">
    <el-card shadow="never">
      <template #header>
        <div class="flex items-center justify-between flex-wrap gap-2">
          <div class="text-base font-semibold">五笔排行榜</div>
          <div class="flex items-center gap-2 flex-wrap">
            <el-select v-model="days" style="width: 130px" @change="load">
              <el-option label="近 7 天" :value="7" />
              <el-option label="近 30 天" :value="30" />
              <el-option label="全部" :value="0" />
            </el-select>
            <el-select v-model="order" style="width: 130px" @change="load">
              <el-option label="按 WPM" value="wpm" />
              <el-option label="按准确率" value="accuracy" />
            </el-select>
            <el-button :loading="loading" @click="load">刷新</el-button>
          </div>
        </div>
      </template>

      <el-table :data="list" v-loading="loading" stripe>
        <el-table-column type="index" label="名次" width="80" />
        <el-table-column prop="nickName" label="用户" min-width="160" />
        <el-table-column prop="mode" label="模式" width="100" />
        <el-table-column prop="totalCount" label="题数" width="100" align="right" />
        <el-table-column prop="accuracy" label="准确率" width="120" align="right">
          <template #default="{ row }">{{ Number(row.accuracy).toFixed(1) }}%</template>
        </el-table-column>
        <el-table-column prop="wpm" label="字/分钟" width="120" align="right">
          <template #default="{ row }">{{ Number(row.wpm).toFixed(1) }}</template>
        </el-table-column>
        <el-table-column prop="durationSecond" label="用时(秒)" width="120" align="right" />
        <el-table-column prop="CreatedAt" label="时间" min-width="180">
          <template #default="{ row }">{{ formatDate(row.CreatedAt) }}</template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-card shadow="never" class="mt-4">
      <template #header>
        <div class="text-base font-semibold">我的成绩</div>
      </template>
      <el-table :data="myList" v-loading="myLoading" stripe>
        <el-table-column prop="mode" label="模式" width="100" />
        <el-table-column prop="totalCount" label="题数" width="100" align="right" />
        <el-table-column prop="accuracy" label="准确率" width="120" align="right">
          <template #default="{ row }">{{ Number(row.accuracy).toFixed(1) }}%</template>
        </el-table-column>
        <el-table-column prop="wpm" label="字/分钟" width="120" align="right">
          <template #default="{ row }">{{ Number(row.wpm).toFixed(1) }}</template>
        </el-table-column>
        <el-table-column prop="durationSecond" label="用时(秒)" width="120" align="right" />
        <el-table-column prop="CreatedAt" label="时间" min-width="180">
          <template #default="{ row }">{{ formatDate(row.CreatedAt) }}</template>
        </el-table-column>
      </el-table>
      <el-pagination
        class="mt-3"
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :total="total"
        :page-sizes="[10, 20, 50]"
        layout="total, sizes, prev, pager, next"
        @size-change="loadMine"
        @current-change="loadMine"
      />
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { formatDate } from '@/utils/format'
import { getLeaderboard, getMyScores } from '@/plugin/wubi/api/score'

const list = ref([])
const loading = ref(false)
const days = ref(0)
const order = ref('wpm')

const myList = ref([])
const myLoading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

const load = async () => {
  loading.value = true
  try {
    const res = await getLeaderboard({ limit: 20, days: days.value, order: order.value })
    if (res.code === 0) list.value = res.data || []
  } finally {
    loading.value = false
  }
}

const loadMine = async () => {
  myLoading.value = true
  try {
    const res = await getMyScores({ page: page.value, pageSize: pageSize.value })
    if (res.code === 0) {
      myList.value = res.data?.list || []
      total.value = res.data?.total || 0
    }
  } finally {
    myLoading.value = false
  }
}

onMounted(() => {
  load()
  loadMine()
})
</script>
