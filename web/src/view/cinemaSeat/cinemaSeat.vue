<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="影厅" prop="hall" @click="getFilms">
          <el-select v-model="searchInfo.hall" placeholder="选择影厅" style="width: 200px;">
            <el-option v-for="dict in hallOptions" :label="dict.label" :value="dict.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="电影" prop="filmId">
          <el-select v-model="searchInfo.filmId" placeholder="选择电影"  style="width: 200px;">
            <el-option v-for="dict in filmOptions" :label="dict.label" :value="dict.value" />
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
        :filmId="String(searchInfo.filmId)"
        :seatInfo="seatInfo"
        :hallId="searchInfo.hall" 
        @printSeatSave="printSeatSave"
        @printSeatDel="printSeatDel"
      />
    </div>
  </div>
</template>

<script setup>
import {
  createCinemaSeat,
  deleteCinemaSeat,
  getCinemaSeatList
} from '@/api/cinemaSeat'

import {
  getCinemaFilmList
} from '@/api/cinemaFilm'

// 全量引入格式化工具 请按需保留
import { ref, watch } from 'vue'
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
const seatInfo = ref([])
const elSearchFormRef = ref()
const searchInfo = ref({
  date: new Date(),
  filmId: '',
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

// 监听影厅
watch(() => searchInfo.value.hall, (v) => {
  filmOptions.value = []
  searchInfo.value.filmId = ''
  seatInfo.value = []
  getFilms()
}, { immediate: true })

// 获取座位
const getSeats = async () => {
  const seatList = await getCinemaSeatList({ page: 1, pageSize: 100, ...searchInfo.value })
  if (seatList.code === 0) {
    if (seatList.data.list === null) {
      seatInfo.value = []
      return
    }
    seatInfo.value = seatList.data.list
  }
}

watch(() => searchInfo.value.filmId, (v) => {
  if (v === '') return
  getSeats()
}, { immediate: true })

watch(() => searchInfo.value.date, (v) => {
  if (v === '') return
  getSeats()
}, { immediate: true })

const printSeatSave = async (seats)=>{
  try {
    console.log('save_seats', seats)
    const params = {
      filmId: searchInfo.value.filmId,
      date: searchInfo.value.date,
      positions: seats
    }
    await createCinemaSeat(params)
  } catch (error) {
    console.log('save_seats_error', error)
  }
}

const printSeatDel = async (seat)=>{
  try {
    console.log('del_seats', seat)
    seat = seatInfo.value.find(item => item.position === seat)
    await deleteCinemaSeat({ ID: seat.ID })
  } catch (error) {
    console.log('del_seats_error', error)
  }
}

</script>

<style></style>
