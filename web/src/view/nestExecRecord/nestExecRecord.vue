<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
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
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="120" />
        <el-table-column align="left" label="创建人" prop="createdBy" width="120" />
        <el-table-column align="left" label="抄送人" prop="copyTo" width="120" />
        <el-table-column align="left" label="航线名称" prop="missionName" width="120" />
        <el-table-column align="left" label="航线id" prop="missionid" width="120" />
         <el-table-column align="left" label="计划时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.planAt) }}</template>
         </el-table-column>
        <el-table-column align="left" label="类型" prop="type" width="120" />
        <el-table-column align="left" label="作业ID" prop="executeId" width="120" />
        <el-table-column align="left" label="作作业人" prop="executeAt" width="120" />
        <el-table-column align="left" label="飞行秒" prop="flyInSecond" width="120" />
        <el-table-column align="left" label="全景链接" prop="panoramaLink" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateNestExecRecordFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="状态:"  prop="status" >
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人:"  prop="createdBy" >
          <el-input v-model="formData.createdBy" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="抄送人:"  prop="copyTo" >
          <el-input v-model="formData.copyTo" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="航线名称:"  prop="missionName" >
          <el-input v-model="formData.missionName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="航线id:"  prop="missionid" >
          <el-input v-model="formData.missionid" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="计划时间:"  prop="planAt" >
          <el-date-picker v-model="formData.planAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="类型:"  prop="type" >
          <el-input v-model="formData.type" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="作业ID:"  prop="executeId" >
          <el-input v-model="formData.executeId" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="作作业人:"  prop="executeAt" >
          <el-input v-model="formData.executeAt" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="飞行秒:"  prop="flyInSecond" >
          <el-input v-model.number="formData.flyInSecond" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="全景链接:"  prop="panoramaLink" >
          <el-input v-model="formData.panoramaLink" :clearable="true"  placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'NestExecRecord'
}
</script>

<script setup>
import {
  createNestExecRecord,
  deleteNestExecRecord,
  deleteNestExecRecordByIds,
  updateNestExecRecord,
  findNestExecRecord,
  getNestExecRecordList
} from '@/api/nestExecRecord'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        status: '',
        createdBy: '',
        copyTo: '',
        missionName: '',
        missionid: '',
        planAt: new Date(),
        type: '',
        executeId: '',
        executeAt: '',
        flyInSecond: 0,
        panoramaLink: '',
        })

// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
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
  elSearchFormRef.value?.validate(async(valid) => {
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
const getTableData = async() => {
  const table = await getNestExecRecordList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const setOptions = async () =>{
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
            deleteNestExecRecordFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteNestExecRecordByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateNestExecRecordFunc = async(row) => {
    const res = await findNestExecRecord({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reNtERecord
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteNestExecRecordFunc = async (row) => {
    const res = await deleteNestExecRecord({ ID: row.ID })
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

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        status: '',
        createdBy: '',
        copyTo: '',
        missionName: '',
        missionid: '',
        planAt: new Date(),
        type: '',
        executeId: '',
        executeAt: '',
        flyInSecond: 0,
        panoramaLink: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createNestExecRecord(formData.value)
                  break
                case 'update':
                  res = await updateNestExecRecord(formData.value)
                  break
                default:
                  res = await createNestExecRecord(formData.value)
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

<style>

</style>
