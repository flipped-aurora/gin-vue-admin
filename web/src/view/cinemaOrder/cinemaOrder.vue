<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
            :disabled-date="time => searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.orderStatus" placeholder="请选择状态" clearable style="width: 100px;">
            <el-option label="收入" value="1" />
            <el-option label="退款" value="2" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData">
        <el-table-column align="center" label="电影" prop="filmName" width="120" />
        <el-table-column align="center" label="价格" prop="filmPrice" width="120" />
        <el-table-column align="center" label="收入|退款" prop="status" width="120">
          <template #default="scope">{{ scope.row.status == 1 ? "收入" : "退款" }}</template>
        </el-table-column>
        <el-table-column align="center" label="打印时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              查看详情
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>


    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情"
      destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
          <!-- <el-descriptions-item label="座位ID">
            {{ formData.seatId }}
          </el-descriptions-item> -->
          <!-- <el-descriptions-item label="电影">
            {{ formData.filmId }}
          </el-descriptions-item> -->
          <el-descriptions-item label="影厅">
            {{ formData.filmHall + "号厅" }}
          </el-descriptions-item>
          <el-descriptions-item label="选座">
            {{ formData.filmSeat }}
          </el-descriptions-item>
          <el-descriptions-item label="电影名称">
            {{ formData.filmName }}
          </el-descriptions-item>
          <el-descriptions-item label="电影类型">
            {{ formData.filmType }}
          </el-descriptions-item>
          <el-descriptions-item label="播放时间">
            {{ formData.playTime }}
          </el-descriptions-item>
          <el-descriptions-item label="价格">
            {{ formData.filmPrice }}
          </el-descriptions-item>
          <el-descriptions-item label="状态">
            {{ formData.status == 1 ? "收入" : "退款" }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  findCinemaOrder,
  getCinemaOrderList
} from '@/api/cinemaOrder'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'CinemaOrder'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  seatId: 0,
  filmId: 0,
  filmHall: false,
  filmSeat: '',
  filmName: '',
  filmType: '',
  playTime: new Date(),
  filmPrice: 0,
  status: false,
})


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
})

const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.filmHall === "") {
      searchInfo.value.filmHall = null
    }
    if (searchInfo.value.status === "") {
      searchInfo.value.status = null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getCinemaOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCinemaOrder({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recinemaOrder
    openDetailShow()
  }
}
// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    seatId: 0,
    filmId: 0,
    filmHall: false,
    filmSeat: '',
    filmName: '',
    filmType: '',
    playTime: new Date(),
    filmPrice: 0,
    status: false,
  }
}
</script>

<style></style>
