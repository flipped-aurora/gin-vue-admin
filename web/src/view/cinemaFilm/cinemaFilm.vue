<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <!-- <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item> -->

        <el-form-item label="影厅" prop="hall">
          <el-select v-model="searchInfo.hall" placeholder="选择影厅" style="width: 100px;">
            <el-option v-for="dict in hallOptions" :key="dict.value" :label="dict.label" :value="dict.value" />
          </el-select>
        </el-form-item>

        <el-form-item label="电影名字" prop="name">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />

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

        <!-- <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->

        <el-table-column align="left" label="影厅" prop="hall" width="120" />
        <el-table-column align="left" label="电影名字" prop="name" width="120" />
        <el-table-column align="left" label="价格" prop="price" width="120" />
        <el-table-column align="left" label="播放时间" width="180" >
          <template #default="scope">{{ scope.row.playTime }}</template>
        </el-table-column>
        <el-table-column align="left" label="电影票类型" prop="type" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <!-- <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              查看详情
            </el-button> -->
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateCinemaFilmFunc(scope.row)">变更</el-button>
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
        <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
          <el-form-item label="影厅:" prop="hall">
            <el-select v-model="formData.hall" placeholder="选择影厅">
              <el-option v-for="dict in hallOptions" :key="dict.value" :label="dict.label" :value="dict.value" />
            </el-select>
          </el-form-item>
          <el-form-item label="电影名字:" prop="name">
            <el-input v-model="formData.name" :clearable="true" placeholder="请输入电影名字" />
          </el-form-item>
          <el-form-item label="价格:" prop="price">
            <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="true" />
          </el-form-item>
          <el-form-item label="播放时间:" prop="playTime">
            <el-time-picker v-model="formData.playTime" :format="'HH:mm'" value-format="HH:mm" placeholder="请输入播放时间" />
          </el-form-item>
          <el-form-item label="电影票类型:" prop="type">
            <el-input v-model="formData.type" :clearable="true" placeholder="请输入电影票类型" />
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
          <el-descriptions-item label="影厅">
            {{ formData.hall }}
          </el-descriptions-item>
          <el-descriptions-item label="电影名字">
            {{ formData.name }}
          </el-descriptions-item>
          <el-descriptions-item label="价格">
            {{ formData.price }}
          </el-descriptions-item>
          <el-descriptions-item label="播放时间">
            {{  forformData.playTime }}
          </el-descriptions-item>
          <el-descriptions-item label="电影票类型">
            {{ formData.type }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCinemaFilm,
  deleteCinemaFilm,
  deleteCinemaFilmByIds,
  updateCinemaFilm,
  findCinemaFilm,
  getCinemaFilmList
} from '@/api/cinemaFilm'

// 全量引入格式化工具 请按需保留
// import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'CinemaFilm'
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
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  hall: 1,
  name: '',
  price: 0,
  playTime: '',
  type: '',
})

// 验证规则
const rule = reactive({
  hall: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  price: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  playTime: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
})

const elFormRef = ref()
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
  const table = await getCinemaFilmList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteCinemaFilmFunc(row)
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
  const res = await deleteCinemaFilmByIds({ IDs })
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
const updateCinemaFilmFunc = async (row) => {
  const res = await findCinemaFilm({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.recinemaFilm
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteCinemaFilmFunc = async (row) => {
  const res = await deleteCinemaFilm({ ID: row.ID })
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
  const res = await findCinemaFilm({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recinemaFilm
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    hall: 0,
    name: '',
    price: 0,
    playTime: new Date(),
    type: '',
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
    hall: 0,
    name: '',
    price: 0,
    playTime: new Date(),
    type: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createCinemaFilm(formData.value)
        break
      case 'update':
        res = await updateCinemaFilm(formData.value)
        break
      default:
        res = await createCinemaFilm(formData.value)
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
