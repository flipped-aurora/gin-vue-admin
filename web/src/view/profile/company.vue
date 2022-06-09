<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="small"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除
            </el-button>
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
        <el-table-column align="left" label="资料名称" prop="profile_name" width="160" />
        <el-table-column align="left" label="公司名称" prop="name" width="120" />
        <el-table-column align="left" label="简介" prop="intro" width="120" />
        <el-table-column align="left" label="网站" prop="website" width="180" />
        <el-table-column align="left" label="图标" prop="logo" width="120" />
        <el-table-column align="left" label="手机号" prop="phone" width="120" />
        <el-table-column align="left" label="邮箱" prop="email" width="180" />
        <el-table-column align="left" label="隐身" prop="stealth" width="60">
          <template #default="scope">{{ formatBoolean(scope.row.stealth) }}</template>
        </el-table-column>
        <el-table-column align="left" fixed="right" label="按钮组" width="140">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              size="small"
              class="table-button"
              @click="updateProfileFunc(scope.row)"
            >变更
            </el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
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
      <el-form ref="form" :model="formData" :rules="rules" label-position="right" label-width="80px">
        <el-form-item label="资料名称:" prop="profile_name">
          <el-input v-model="formData.profile_name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公司名称:" prop="name">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="简介:" prop="intro">
          <el-input v-model="formData.intro" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="网站:" prop="website">
          <el-input v-model="formData.website" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图标:" prop="logo">
          <el-input v-model="formData.logo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="手机号:" prop="phone">
          <el-input v-model="formData.phone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="隐身:" prop="stealth">
          <el-switch
            v-model="formData.stealth"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Profile',
}
</script>

<script setup>
import {
  createCompanyProfile,
  deleteCompanyProfile,
  deleteCompanyProfileByIds,
  updateCompanyProfile,
  findCompanyProfile,
  getCompanyProfileList,
} from '@/api/profile'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const genderOptions = ref([])
const form = ref(null)
const formData = ref({
  profile_name: '老王科技',
  name: '老王科技',
  intro: '老王科技要做更好的自己',
  website: 'https://laowang.org',
  logo: 'https://laowang.org',
  phone: '18888888888',
  email: 'laowang@laowang.org',
  stealth: false,
})

// TODO 添加 个性化展示页面

// 格式限制
const rules = reactive({
  profile_name: [
    { required: true, message: '请输入资料名称', trigger: 'blur' },
    { min: 2, max: 100, message: '长度在 2 到 100 个字符', trigger: 'blur' },
  ],
  name: [
    { required: true, message: '请输入公司名称', trigger: 'blur' },
    { min: 2, max: 100, message: '长度在 2 到 100 个字符', trigger: 'blur' },
  ],
  intro: [
    { required: true, message: '请输入简介', trigger: 'blur' },
    { min: 2, max: 5000, message: '长度在 2 到 5000 个字符', trigger: 'blur' },
  ],
  website: [
    { required: true, message: '请输入网站地址', trigger: 'blur' },
    { type: 'url', min: 2, max: 100, message: '请输入正确的链接地址, 长度在 2 到 100 个字符', trigger: 'blur' },
  ],
  logo: [
    { required: true, message: '请输入图标地址', trigger: 'blur' },
    { type: 'url', min: 2, max: 100, message: '请输入正确的链接地址, 长度在 2 到 100 个字符', trigger: 'blur' },
  ],
  phone: [
    { required: true, message: '请输入联系电话', trigger: 'blur' },
    { min: 11, max: 15, message: '长度在 11 到 15 个字符', trigger: 'blur' },
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', min: 6, max: 50, message: '请输入正确的邮箱地址', trigger: 'blur' },
  ],
  stealth: [
    { required: true, message: '请选择是否公开资料', trigger: 'blur', type: 'enum', enum: [true, false] },
  ],
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.stealth === '') {
    searchInfo.value.stealth = null
  }
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
  const table = await getCompanyProfileList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const setOptions = async() => {
  genderOptions.value = await getDictFunc('gender')
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
    type: 'warning',
  }).then(() => {
    deleteProfileFunc(row)
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
      message: '请选择要删除的数据',
    })
    return
  }
  multipleSelection.value &&
  multipleSelection.value.map(item => {
    ids.push(item.ID)
  })
  const res = await deleteCompanyProfileByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
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
const updateProfileFunc = async(row) => {
  const res = await findCompanyProfile({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reprofile
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteProfileFunc = async(row) => {
  const res = await deleteCompanyProfile({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
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
    profile_name: '',
    username: '',
    age: 0,
    sex: undefined,
    work_experience: 0,
    phone: '',
    resume_file: '',
    stealth: false,
    company_name: '',
    intro: '',
    logo: '',
    email: '',
    website: '',
    user_id: 0,
    user_name: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  form.value.validate(async(v) => {
    if (v) {
      let res
      switch (type.value) {
        case 'create':
          res = await createCompanyProfile(formData.value)
          break
        case 'update':
          res = await updateCompanyProfile(formData.value)
          break
        default:
          res = await createCompanyProfile(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功',
        })
        closeDialog()
        getTableData()
      }
    }
  })
}
</script>

<style>
</style>
