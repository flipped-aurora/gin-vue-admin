<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="用户名">
         <el-input v-model="searchInfo.userName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="角色id">
            
             <el-input v-model.number="searchInfo.roleID" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="昵称">
         <el-input v-model="searchInfo.nickname" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="公司id">
            
             <el-input v-model.number="searchInfo.companyID" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="工资">
            
            <el-input v-model.number="searchInfo.startWages" placeholder="搜索条件（起）" />
            —
            <el-input v-model.number="searchInfo.endWages" placeholder="搜索条件（止）" />

        </el-form-item>
           <el-form-item label="工作计算方式" prop="workType">
            <el-select v-model="searchInfo.workType" clearable placeholder="请选择" @clear="()=>{searchInfo.workType=undefined}">
              <el-option v-for="(item,key) in workTypeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
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
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="用户名" prop="userName" width="120" />
        <el-table-column align="left" label="密码" prop="password" width="120" />
        <el-table-column align="left" label="角色id" prop="roleID" width="120" />
        <el-table-column align="left" label="昵称" prop="nickname" width="120" />
        <el-table-column align="left" label="公司id" prop="companyID" width="120" />
        <el-table-column sortable align="left" label="工资" prop="wages" width="120" />
        <el-table-column align="left" label="工作计算方式" prop="workType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.workType,workTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateAppUserFunc(scope.row)">变更</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="用户名:"  prop="userName" >
          <el-input v-model="formData.userName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:"  prop="password" >
          <el-input v-model="formData.password" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="角色id:"  prop="roleID" >
          <el-input v-model.number="formData.roleID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="昵称:"  prop="nickname" >
          <el-input v-model="formData.nickname" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公司id:"  prop="companyID" >
          <el-input v-model.number="formData.companyID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="工资:"  prop="wages" >
          <el-input-number v-model="formData.wages"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="工作计算方式:"  prop="workType" >
          <el-select v-model="formData.workType" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in workTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态:"  prop="status" >
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'AppUser'
}
</script>

<script setup>
import {
  createAppUser,
  deleteAppUser,
  deleteAppUserByIds,
  updateAppUser,
  findAppUser,
  getAppUserList
} from '@/api/appUser'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const workTypeOptions = ref([])
const statusOptions = ref([])
const formData = ref({
        userName: '',
        password: '',
        roleID: 0,
        nickname: '',
        companyID: 0,
        wages: 0,
        workType: undefined,
        status: undefined,
        })

// 验证规则
const rule = reactive({
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
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
  const table = await getAppUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    workTypeOptions.value = await getDictFunc('workType')
    statusOptions.value = await getDictFunc('status')
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
            deleteAppUserFunc(row)
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
      const res = await deleteAppUserByIds({ ids })
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
const updateAppUserFunc = async(row) => {
    const res = await findAppUser({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reappUser
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAppUserFunc = async (row) => {
    const res = await deleteAppUser({ ID: row.ID })
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
        userName: '',
        password: '',
        roleID: 0,
        nickname: '',
        companyID: 0,
        wages: 0,
        workType: undefined,
        status: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createAppUser(formData.value)
                  break
                case 'update':
                  res = await updateAppUser(formData.value)
                  break
                default:
                  res = await createAppUser(formData.value)
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
