<template>
  <div>
    <warning-bar
      href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
      title="此功能为开发环境使用，不建议发布到生产，具体使用效果请看视频https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
    />
    <div class="gva-table-box">
      <div class="gva-btn-list gap-3 flex items-center">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog('addPackage')"
        >新增</el-button>
        <el-icon
          class="cursor-pointer"
          @click="toDoc('https://www.bilibili.com/video/BV1kv4y1g7nT?p=3&vd_source=f2640257c21e3b547a790461ed94875e')"
        ><VideoCameraFilled /></el-icon>
      </div>
      <el-table :data="tableData">
        <el-table-column
          align="left"
          label="id"
          width="60"
          prop="ID"
        />
        <el-table-column
          align="left"
          label="包名"
          width="150"
          prop="packageName"
        />
        <el-table-column
          align="left"
          label="展示名"
          width="150"
          prop="label"
        />
        <el-table-column
          align="left"
          label="描述"
          min-width="150"
          prop="desc"
        />

        <el-table-column
          align="left"
          label="操作"
          width="200"
        >
          <template #default="scope">
            <el-button
              icon="edit"
              type="primary"
              link
              @click="editPackageFunc(scope.row)"
            >编辑</el-button>
            <el-button icon="delete"

              type="primary"
              link
              @click="deleteApiFunc(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>

    </div>

    <el-drawer
      v-model="dialogFormVisible"
      size="40%"
      :show-close="false"
    >
      <warning-bar title="新增Pkg用于自动化代码使用" />
      <el-form
        ref="pkgForm"
        :model="form"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item
          label="包名"
          prop="packageName"
        >
          <el-input
              v-model="form.packageName"
              autocomplete="off"
              :disabled="type === 'edit'"
          />
        </el-form-item>
        <el-form-item
          label="展示名"
          prop="label"
        >
          <el-input
            v-model="form.label"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="描述"
          prop="desc"
        >
          <el-input
            v-model="form.desc"
            autocomplete="off"
          />
        </el-form-item>
      </el-form>
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">创建Package</span>
          <div>
            <el-button @click="closeDialog">取 消</el-button>
            <el-button
                type="primary"
                @click="enterDialog"
            >确 定</el-button>
          </div>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createPackageApi,
  getPackageApi,
  deletePackageApi,
  updatePackageDetail,
  getPackageById
} from '@/api/autoCode'
import { ref } from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { toDoc } from '@/utils/doc'
import { VideoCameraFilled } from '@element-plus/icons-vue'

defineOptions({
  name: 'AutoPkg',
})

const form = ref({
  packageName: '',
  label: '',
  desc: '',
})

const validateNum = (rule, value, callback) => {
  if ((/^\d+$/.test(value[0]))) {
    callback(new Error('不能够以数字开头'))
  } else {
    callback()
  }
}
const type = ref('')
const rules = ref({
  packageName: [
    { required: true, message: '请输入包名', trigger: 'blur' },
    { validator: validateNum, trigger: 'blur' }
  ],
})

// 编辑弹窗相关
const pkgForm = ref(null)
const initForm = () => {
  pkgForm.value.resetFields()
  form.value = {
    packageName: '',
    desc: '',
    label: '',
  }
}
const dialogTitle = ref('新增Package')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addPackage':
      dialogTitle.value = '新增Package'
      break
    case 'edit':
      dialogTitle.value = '编辑Package'
      break
    default:
      break
  }
  type.value = key
  dialogFormVisible.value = true
}

const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}

const editPackageFunc = async (row) => {
  const res = await getPackageById({id: row.ID})
  form.value = res.data.pkg
  openDialog('edit')
}

const enterDialog = async () => {
  pkgForm.value.validate(async valid => {
    if (valid) {
      switch (type.value) {
        case 'addPackage': {
          const res = await createPackageApi(form.value)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '添加成功',
              showClose: true
            })
          }
          getTableData()
          closeDialog()
        }
          break
        case 'edit': {
          const res = await updatePackageDetail(form.value)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '编辑成功',
              showClose: true

            })
          }
          getTableData()
          closeDialog()
        }
          break
        default:
          // eslint-disable-next-line no-lone-blocks
        {
          ElMessage({
            type: 'error',
            message: '未知操作',
            showClose: true
          })
        }
          break
      }
    }
  })
}

const tableData = ref([])
const getTableData = async() => {
  const table = await getPackageApi()
  if (table.code === 0) {
    tableData.value = table.data.pkgs
  }
}

const deleteApiFunc = async(row) => {
  ElMessageBox.confirm('此操作仅删除数据库中的pkg存储，后端相应目录结构请自行删除与数据库保持一致！', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deletePackageApi(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        getTableData()
      }
    })
}

getTableData()
</script>
