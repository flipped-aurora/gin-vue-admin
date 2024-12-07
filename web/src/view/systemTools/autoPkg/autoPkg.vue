<template>
  <div>
    <warning-bar
      href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
      title="此功能为开发环境使用，不建议发布到生产，具体使用效果请看视频https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
    />
    <div class="gva-table-box">
      <div class="gva-btn-list gap-3 flex items-center">
        <el-button type="primary" icon="plus" @click="openDialog('addApi')">
          新增
        </el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column align="left" label="id" width="120" prop="ID" />
        <el-table-column
          align="left"
          label="包名"
          width="150"
          prop="packageName"
        />
        <el-table-column
          align="left"
          label="模板"
          width="150"
          prop="template"
        />
        <el-table-column align="left" label="展示名" width="150" prop="label" />
        <el-table-column
          align="left"
          label="描述"
          min-width="150"
          prop="desc"
        />

        <el-table-column align="left" label="操作" width="200">
          <template #default="scope">
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteApiFunc(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-drawer v-model="dialogFormVisible" size="40%" :show-close="false">
      <warning-bar
        title="模板package会创建集成于项目本体中的代码包，模板plugin会创建插件包"
      />
      <el-form ref="pkgForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="包名" prop="packageName">
          <el-input v-model="form.packageName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="模板" prop="template">
          <el-select v-model="form.template">
            <el-option
              v-for="template in templatesOptions"
              :label="template"
              :value="template"
              :key="template"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="展示名" prop="label">
          <el-input v-model="form.label" autocomplete="off" />
        </el-form-item>
        <el-form-item label="描述" prop="desc">
          <el-input v-model="form.desc" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">创建Package</span>
          <div>
            <el-button @click="closeDialog"> 取 消 </el-button>
            <el-button type="primary" @click="enterDialog"> 确 定 </el-button>
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
    getTemplatesApi
  } from '@/api/autoCode'
  import { ref } from 'vue'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { ElMessage, ElMessageBox } from 'element-plus'

  defineOptions({
    name: 'AutoPkg'
  })

  const form = ref({
    packageName: '',
    template: '',
    label: '',
    desc: ''
  })
  const templatesOptions = ref([])

  const getTemplates = async () => {
    const res = await getTemplatesApi()
    if (res.code === 0) {
      templatesOptions.value = res.data
    }
  }

  getTemplates()

  const validateData = (rule, value, callback) => {
    if (/[\u4E00-\u9FA5]/g.test(value)) {
      callback(new Error('不能为中文'))
    } else if (/^\d+$/.test(value[0])) {
      callback(new Error('不能够以数字开头'))
    } else {
      callback()
    }
  }

  const rules = ref({
    packageName: [
      { required: true, message: '请输入包名', trigger: 'blur' },
      { validator: validateData, trigger: 'blur' }
    ],
    template: [
      { required: true, message: '请选择模板', trigger: 'change' },
      { validator: validateData, trigger: 'blur' }
    ]
  })

  const dialogFormVisible = ref(false)
  const openDialog = () => {
    dialogFormVisible.value = true
  }

  const closeDialog = () => {
    dialogFormVisible.value = false
    form.value = {
      packageName: '',
      template: '',
      label: '',
      desc: ''
    }
  }

  const pkgForm = ref(null)
  const enterDialog = async () => {
    pkgForm.value.validate(async (valid) => {
      if (valid) {
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
    })
  }

  const tableData = ref([])
  const getTableData = async () => {
    const table = await getPackageApi()
    if (table.code === 0) {
      tableData.value = table.data.pkgs
    }
  }

  const deleteApiFunc = async (row) => {
    ElMessageBox.confirm(
      '此操作仅删除数据库中的pkg存储，后端相应目录结构请自行删除与数据库保持一致！',
      '提示',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    ).then(async () => {
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
