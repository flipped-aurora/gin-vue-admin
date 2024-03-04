<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="影厅" prop="hall" @click="getFilms">
          <el-select v-model="searchInfo.hall" placeholder="选择影厅" style="width: 200px;">
            <el-option v-for="dict in hallOptions" :key="dict.value" :label="dict.label" :value="dict.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="电影" prop="filmId">
          <el-select v-model="searchInfo.filmId" placeholder="选择电影" style="width: 200px;">
            <el-option v-for="dict in filmOptions" :key="dict.value" :label="dict.label" :value="dict.value" />
          </el-select>
        </el-form-item>
        <!-- // 默认时间是当天 -->
        <el-form-item label="打印日期" prop="date">
          <el-date-picker v-model="searchInfo.date" type="date" placeholder="打印日期" />
        </el-form-item>
        <!-- <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item> -->
      </el-form>
    </div>
    <div class="gva-table-box">
      <SeatSelect 
        :propFilmOptions="filmOptions"
         :filmId="searchInfo.filmId"
         :hallId="searchInfo.hall" 
         @printSeatSave="printSeatSave" 
         @printSeatDel="printSeatDel"/>
    </div>
  </div>
</template>

<script setup>
import {
  createCinemaSeat,
  deleteCinemaSeat,
  deleteCinemaSeatByIds,
  updateCinemaSeat,
  findCinemaSeat,
  getCinemaSeatList
} from '@/api/cinemaSeat'

import {
  getCinemaFilmList
} from '@/api/cinemaFilm'

// 全量引入格式化工具 请按需保留
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch } from 'vue'
import SeatSelect from './cinemaSeatSelect.vue'

defineOptions({
  name: 'CinemaSeat'
})

const hallOptions = ref([
  { label: '1号厅', value: 1 },
  { label: '2号厅', value: 2 },
  { label: '3号厅', value: 3 },
  { label: '4号厅', value: 4 },
  { label: '5号厅', value: 5 },
  { label: '6号厅', value: 6 },
  { label: '7号厅', value: 7 },
])

const filmOptions = ref([])
const elSearchFormRef = ref()
const searchInfo = ref({
  date: new Date(),
  hall: 1,
})

// 获取当前影厅的电影
const getFilms = async () => {
  const filmList = await getCinemaFilmList({ page: 1, pageSize: 100, ...searchInfo.value })
  if (filmList.code === 0) {
    filmList.data.list.forEach(item => {
      filmOptions.value.push({
        info: item,
        label: `${item.name}(${item.playTime})`,
        value: item.ID,
      })
    })
  }
}

getFilms()
// 监听影厅
watch(() => searchInfo.value.hall, (v) => {
  filmOptions.value = []
  searchInfo.value.filmId = ''
  getFilms()
}, { immediate: true })

const printSeatSave = async (seats)=>{
  try {
    console.log('save_seats', seats)
    const params={
      filmId: searchInfo.value.film,
      date: searchInfo.value.date,
      position: seats
    }
    await createCinemaSeat(params)
  } catch (error) {
    console.log('save_seats_error', error)
  }
}

const printSeatDel = async (seat)=>{
  try {
    console.log('del_seats', seat)
    await deleteCinemaSeat({ seatId: seat.ID })
  } catch (error) {
    console.log('del_seats_error', error)
  }
}

</script>

<style></style>
