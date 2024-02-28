<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline"
        @keyup.enter="onSubmit">
        <el-form-item label="影厅" prop="hall" @click="getFilms">
          <el-select v-model="searchInfo.hall" placeholder="选择影厅" style="width: 100px;">
            <el-option v-for="dict in hallOptions" :key="dict.value" :label="dict.label" :value="dict.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="电影" prop="film">
          <el-select v-model="searchInfo.film" placeholder="选择电影" style="width: 100px;">
            <el-option v-for="dict in filmOptions" :key="dict.value" :label="dict.label" :value="dict.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="打印日期" prop="date">
          <!-- // 默认时间是当天 -->
          <el-date-picker v-model="searchInfo.date" type="date" placeholder="打印日期" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
              @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="电影院" prop="filmId" width="120" />
        <el-table-column align="left" label="打印日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.date) }}</template>
        </el-table-column>
        <el-table-column align="left" label="几排几座" prop="position" width="120" />
        <el-table-column align="left" label="状态1: 正常 2: 重打 0:退款" prop="status" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateCinemaSeatFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type === 'create' ? '添加' : '修改'"
      destroy-on-close>
      <el-scrollbar height="500px">
        <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
          <el-form-item label="电影院:" prop="filmId">
            <el-input v-model.number="formData.filmId" :clearable="true" placeholder="请输入电影院" />
          </el-form-item>
          <el-form-item label="打印日期:" prop="date">
            <el-date-picker v-model="formData.date" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
          </el-form-item>
          <el-form-item label="几排几座:" prop="position">
            <el-input v-model="formData.position" :clearable="true" placeholder="请输入几排几座" />
          </el-form-item>
          <el-form-item label="状态1: 正常 2: 重打 0:退款:" prop="status">
            <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
              inactive-text="否" clearable></el-switch>
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情"
      destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
          <el-descriptions-item label="电影院">
            {{ formData.filmId }}
          </el-descriptions-item>
          <el-descriptions-item label="打印日期">
            {{ formatDate(formData.date) }}
          </el-descriptions-item>
          <el-descriptions-item label="几排几座">
            {{ formData.position }}
          </el-descriptions-item>
          <el-descriptions-item label="状态1: 正常 2: 重打 0:退款">
            {{ formatBoolean(formData.status) }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
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
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'CinemaSeat'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  filmId: 0,
  date: new Date(),
  position: '',
  status: false,
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

const filmOptions = ref()

// 验证规则
const rule = reactive({
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(100)
const tableData = ref([])
const searchInfo = ref({
  date: new Date(),
  hall: 1,
})

// 重置
const onReset = () => {
  searchInfo.value = {
    date: new Date(),
    hall: 1,
  }
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
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

const getFilms = async () => {
  const filmList = await getCinemaFilmList({ page: 1, pageSize: 100, ...searchInfo.value })
  if (filmList.code === 0) {
    filmList.data.list.forEach(item => {
      filmOptions.value.push({
        label: item.name,
        value: item.ID,
      })
    })
  }
}
// 查询
const getTableData = async () => {
  const table = await getCinemaSeatList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteCinemaSeatFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
  const IDs = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      IDs.push(item.ID)
    })
  const res = await deleteCinemaSeatByIds({ IDs })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === IDs.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCinemaSeatFunc = async (row) => {
  const res = await findCinemaSeat({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.recinemaSeat
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteCinemaSeatFunc = async (row) => {
  const res = await deleteCinemaSeat({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCinemaSeat({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recinemaSeat
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    filmId: 0,
    date: new Date(),
    position: '',
    status: false,
  }
}


// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    filmId: 0,
    date: new Date(),
    position: '',
    status: false,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createCinemaSeat(formData.value)
        break
      case 'update':
        res = await updateCinemaSeat(formData.value)
        break
      default:
        res = await createCinemaSeat(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style></style>
