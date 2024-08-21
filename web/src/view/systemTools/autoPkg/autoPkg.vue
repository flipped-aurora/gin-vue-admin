<template>
  <div>
    <warning-bar
      href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
      :title="t('view.systemTools.autoPkg.autoPkgNote')"
    />
    <div class="gva-table-box">
      <div class="gva-btn-list gap-3 flex items-center">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog('addApi')"
        >
          {{ t('general.add') }}
        </el-button>

      </div>
      <el-table :data="tableData">
        <el-table-column
          align="left"
          label="ID"
          width="60"
          prop="ID"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.autoPkg.packageName')"
          width="150"
          prop="packageName"
        />
        <el-table-column
            align="left"
            :label="t('view.systemTools.autoPkg.template')"
            width="150"
            prop="template"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.autoPkg.displayName')"
          width="150"
          prop="label"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.autoPkg.description')"
          min-width="150"
          prop="desc"
        />

        <el-table-column
          align="left"
          :label="t('general.operations')"
          width="200"
        >
          <template #default="scope">
            <el-button
              icon="delete"

              type="primary"
              link
              @click="deleteApiFunc(scope.row)"
            >
              {{ t('general.delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-drawer
      v-model="dialogFormVisible"
      size="40%"
      :show-close="false"
    >
      <warning-bar :title="t('view.systemTools.autoPkg.templatePackageNote')" />
      <el-form
        ref="pkgForm"
        :model="form"
        :rules="rules"
        label-width="120px"
      >
        <el-form-item
          :label="t('view.systemTools.autoPkg.packageName')"
          prop="packageName"
        >
          <el-input
            v-model="form.packageName"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          :label="t('view.systemTools.autoPkg.template')"
          prop="template"
        >
          <el-select
            v-model="form.template"
          >
            <el-option v-for="template in templatesOptions" :label="template" :value="template" :key="template"/>
          </el-select>
        </el-form-item>

        <el-form-item
          :label="t('view.systemTools.autoPkg.displayName')"
          prop="label"
        >
          <el-input
            v-model="form.label"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          :label="t('view.systemTools.autoPkg.description')"
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
          <span class="text-lg">{{ t('view.systemTools.autoPkg.creatingPackage') }}</span>
          <div>
            <el-button @click="closeDialog">
              {{ t('general.close') }}
            </el-button>
            <el-button
              type="primary"
              @click="enterDialog"
            >
              {{ t('general.confirm') }}
            </el-button>
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
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

defineOptions({
  name: 'AutoPkg',
})

const form = ref({
  packageName: '',
  template: '',
  label: '',
  desc: '',
})
const templatesOptions = ref([])

const getTemplates = async ()=>{
  const res = await getTemplatesApi()
  if (res.code === 0){
    templatesOptions.value = res.data
  }
}

getTemplates()

const validateNum = (rule, value, callback) => {
  if ((/^\d+$/.test(value[0]))) {
    callback(new Error(t('view.systemTools.autoPkg.cannotStartWithNumberNote')))
  } else {
    callback()
  }
}

const rules = ref({
  packageName: [
    { required: true, message: t('view.systemTools.autoPkg.enterPackageNameNote'), trigger: 'blur' },
    { validator: validateNum, trigger: 'blur' }
  ],
  template:[
    { required: true, message: t('view.systemTools.autoPkg.selectTemplateNote'), trigger: 'change' },
    { validator: validateNum, trigger: 'blur' }
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
    desc: '',
  }
}

const pkgForm = ref(null)
const enterDialog = async() => {
  pkgForm.value.validate(async valid => {
    if (valid) {
      const res = await createPackageApi(form.value)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('view.systemTools.autoPkg.addSuccess'),
          showClose: true
        })
      }
      getTableData()
      closeDialog()
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
  ElMessageBox.confirm(t('view.systemTools.autoPkg.deletePackageNote'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  })
    .then(async() => {
      const res = await deletePackageApi(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess')
        })
        getTableData()
      }
    })
}

getTableData()
</script>
