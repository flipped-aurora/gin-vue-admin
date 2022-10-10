<template>
  <div>
    <warning-bar href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3" title="此功能为开发环境使用，不建议发布到生产，具体使用效果请看视频https://www.bilibili.com/video/BV1kv4y1g7nT?p=3" />
    <!-- 从数据库直接获取字段 -->
    <div class="gva-search-box">
      <el-collapse v-model="activeNames" style="margin-bottom:12px">
        <el-collapse-item name="1">
          <template #title>
            <div :style="{fontSize:'16px',paddingLeft:'20px'}">
              {{ t('autoCode.existDB') }}
              <el-icon class="header-icon ">
                <pointer />
              </el-icon>
            </div>
          </template>
          <el-form ref="getTableForm" style="margin-top:24px" :inline="true" :model="dbform" label-width="120px">
            <el-form-item :label="t('autoCode.dbName')" prop="structName">
              <el-select v-model="dbform.dbName" filterable :placeholder="t('autoCode.selectDB')" @change="getTableFunc">
                <el-option
                  v-for="item in dbOptions"
                  :key="item.database"
                  :label="item.database"
                  :value="item.database"
                />
              </el-select>
            </el-form-item>
            <el-form-item :label="t('autoCode.tableName')" prop="structName">
              <el-select
                v-model="dbform.tableName"
                :disabled="!dbform.dbName"
                filterable
                :placeholder="t('autoCode.selectTable')"
              >
                <el-option
                  v-for="item in tableOptions"
                  :key="item.tableName"
                  :label="item.tableName"
                  :value="item.tableName"
                />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button size="mini" type="primary" @click="getColumnFunc">{{ t('autoCode.createUsingTable') }}</el-button>
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
    </div>
    <div class="gva-search-box">
      <!-- 初始版本自动化代码工具 -->
      <el-form ref="autoCodeForm" :rules="rules" :model="form" label-width="120px" :inline="true">
        <el-form-item :label="t('autoCode.structName')" prop="structName">
          <el-input v-model="form.structName" :placeholder="t('autoCode.structNameNote')" />
        </el-form-item>
        <el-form-item :label="t('autoCode.tableName')" prop="tableName">
          <el-input v-model="form.tableName" :placeholder="t('autoCode.tableNameNote')" />
        </el-form-item>
        <el-form-item :label="t('autoCode.structAbbreviation')" prop="abbreviation">
          <el-input v-model="form.abbreviation" :placeholder="t('autoCode.structAbbreviationNote')" />
        </el-form-item>
        <el-form-item :label="t('autoCode.structChineseName')" prop="description">
          <el-input v-model="form.description" :placeholder="t('autoCode.structChineseNameNote')" />
        </el-form-item>
        <el-form-item :label="t('autoCode.fileName')" prop="packageName">
          <el-input v-model="form.packageName" :placeholder="t('autoCode.fileNameNote')" @blur="toLowerCaseFunc(form,'packageName')" />
        </el-form-item>
        <el-form-item label="Package（包）" prop="package">
          <el-select v-model="form.package" style="width:194px">
            <el-option v-for="item in pkgs" :key="item.ID" :value="item.packageName" :label="item.packageName" />
          </el-select>
          <el-icon class="auto-icon" @click="getPkgs"><refresh /></el-icon>
          <el-icon class="auto-icon" @click="goPkgs"><document-add /></el-icon>
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip content="注：会自动在结构体添加 created_by updated_by deleted_by，方便用户进行资源权限控制" placement="bottom" effect="light">
              <div> 创建资源标识 <el-icon><QuestionFilled /></el-icon> </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoCreateResource" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip :content="t('autoCode.autoAPIDBTip')" placement="bottom" effect="light">
              <div> {{ t('autoCode.autoAPIDBCreate') }} </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoCreateApiToSql" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip :content="t('autoCode.autoMoveFilesTip')" placement="bottom" effect="light">
              <div> {{ t('autoCode.autoMoveFiles') }} </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoMoveFile" />
        </el-form-item>
      </el-form>
    </div>
    <!-- 组件列表 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" @click="editAndAddField()">{{ t('autoCode.addField') }}</el-button>
      </div>
      <el-table :data="form.fields">
        <el-table-column align="left" type="index" :label="t('autoCode.fieldIndex')" width="60" />
        <el-table-column align="left" prop="fieldName" :label="t('autoCode.fieldName')" />
        <el-table-column align="left" prop="fieldDesc" :label="t('autoCode.fieldDesc')" />
        <el-table-column align="left" prop="require" label="是否必填">
          <template #default="{row}">{{ row.require?"是":"否" }}</template>
        </el-table-column>
        <el-table-column align="left" prop="fieldJson" min-width="120px" label="FieldJson" />
        <el-table-column align="left" prop="fieldType" :label="t('autoCode.fieldDataType')" width="130" />
        <el-table-column align="left" prop="dataTypeLong" :label="t('autoCode.fieldLen')" width="130" />
        <el-table-column align="left" prop="columnName" :label="t('autoCode.columnName')" width="130" />
        <el-table-column align="left" prop="comment" :label="t('autoCode.comment')" width="130" />
        <el-table-column align="left" prop="fieldSearchType" :label="t('general.searchCriteria')" width="130" />
        <el-table-column align="left" prop="dictType" :label="t('autoCode.dictionary')" width="130" />
        <el-table-column align="left" :lable="t('general.operations')" width="300" fixed="right">
          <template #default="scope">
            <el-button
              size="small"
              type="primary"
              link
              icon="edit"
              @click="editAndAddField(scope.row)"
            >{{ t('general.edit') }}</el-button>
            <el-button
              size="small"
              type="primary"
              link
              :disabled="scope.$index === 0"
              @click="moveUpField(scope.$index)"
            >{{ t('autoCode.moveUp') }}</el-button>
            <el-button
              size="small"
              type="primary"
              link
              :disabled="(scope.$index + 1) === form.fields.length"
              @click="moveDownField(scope.$index)"
            >{{ t('autoCode.moveDown') }}</el-button>
            <el-popover v-model="scope.row.visible" placement="top">
              <p>{{ t('autoCode.confirmDelete') }}</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="scope.row.visible = false">{{ t('general.cancel') }}</el-button>
                <el-button type="primary" size="small" @click="deleteField(scope.$index)">{{ t('general.confirm') }}</el-button>
              </div>
              <template #reference>
                <el-button size="small" type="primary" link icon="delete" @click="scope.row.visible = true">{{ t('general.delete') }}</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <!-- 组件列表 -->
      <div class="gva-btn-list justify-content-flex-end auto-btn-list">
        <el-button size="mini" type="primary" @click="enterForm(true)">{{ t('autoCode.codePreview') }}</el-button>
        <el-button size="mini" type="primary" @click="enterForm(false)">{{ t('autoCode.generateCode') }}</el-button>
      </div>
    </div>
    <!-- 组件弹窗 -->
    <el-dialog v-model="dialogFlag" width="70%" :title="t('autoCode.componentContent')">
      <FieldDialog v-if="dialogFlag" ref="fieldDialogNode" :dialog-middle="dialogMiddle" />
      <template #footer>
        <div class="dialog-footer">
          <el-button size="mini" @click="closeDialog">{{ t('general.close') }}</el-button>
          <el-button size="mini" type="primary" @click="enterDialog">{{ t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="previewFlag">
      <template #header>
        <div class="previewCodeTool">
          <p>{{ t('autoCode.actionBar') }}</p>
          <el-button size="mini" type="primary" @click="selectText">{{ t('general.selectAll') }}</el-button>
          <el-button size="mini" type="primary" @click="copy">{{ t('autoCode.copy') }}</el-button>
        </div>
      </template>
      <PreviewCodeDialog v-if="previewFlag" ref="preview" :preview-code="preViewCode" />
      <template #footer>
        <div class="dialog-footer" style="padding-top:14px;padding-right:14px">
          <el-button size="small" type="primary" @click="previewFlag = false">{{ t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>

export default {
  name: 'AutoCode'
}
</script>

<script setup>
import FieldDialog from '@/view/systemTools/autoCode/component/fieldDialog.vue'
import PreviewCodeDialog from '@/view/systemTools/autoCode/component/previewCodeDialg.vue'
import { toUpperCase, toHump, toSQLLine, toLowerCase } from '@/utils/stringFun'
import { createTemp, getDB, getTable, getColumn, preview, getMeta, getPackageApi } from '@/api/autoCode'
import { getDict } from '@/utils/dictionary'
import { ref, getCurrentInstance, reactive, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const fieldTemplate = {
  fieldName: '',
  fieldDesc: '',
  fieldType: '',
  dataType: '',
  fieldJson: '',
  columnName: '',
  dataTypeLong: '',
  comment: '',
  require: false,
  errorText: '',
  clearable: true,
  fieldSearchType: '',
  dictType: ''
}

const route = useRoute()
const router = useRouter()
const activeNames = reactive([])
const preViewCode = ref({})
const dbform = ref({
  dbName: '',
  tableName: ''
})
const dbOptions = ref([])
const tableOptions = ref([])
const addFlag = ref('')
const fdMap = ref({})
const form = ref({
  structName: '',
  tableName: '',
  packageName: '',
  package: '',
  abbreviation: '',
  description: '',
  autoCreateApiToSql: true,
  autoMoveFile: true,
  autoCreateResource: false,
  fields: []
})
const rules = ref({
  structName: [
    { required: true, message: t('autoCode.entStructName'), trigger: 'blur' }
  ],
  abbreviation: [
    { required: true, message: t('autoCode.entStructAbbreviation'), trigger: 'blur' }
  ],
  description: [
    { required: true, message: t('autoCode.entStructDesc'), trigger: 'blur' }
  ],
  packageName: [
    {
      required: true,
      message: t('autoCode.entFileName'),
      trigger: 'blur'
    }
  ],
  package: [
    { required: true, message: '请选择package', trigger: 'blur' }
  ]
})

const dialogMiddle = ref({})
const bk = ref({})
const dialogFlag = ref(false)
const previewFlag = ref(false)

const toLowerCaseFunc = (form, key) => {
  form[key] = toLowerCase(form[key])
}

const previewNode = ref(null)

const selectText = () => {
  previewNode.value.selectText()
}

const copy = () => {
  previewNode.value.copy()
}

const editAndAddField = (item) => {
  dialogFlag.value = true
  if (item) {
    addFlag.value = 'edit'
    bk.value = JSON.parse(JSON.stringify(item))
    dialogMiddle.value = item
  } else {
    addFlag.value = 'add'
    dialogMiddle.value = JSON.parse(JSON.stringify(fieldTemplate))
  }
}

const moveUpField = (index) => {
  if (index === 0) {
    return
  }
  const oldUpField = form.value.fields[index - 1]
  form.value.fields.splice(index - 1, 1)
  form.value.fields.splice(index, 0, oldUpField)
}
const moveDownField = (index) => {
  const fCount = form.value.fields.length
  if (index === fCount - 1) {
    return
  }
  const oldDownField = form.value.fields[index + 1]
  form.value.fields.splice(index + 1, 1)
  form.value.fields.splice(index, 0, oldDownField)
}

const currentInstance = getCurrentInstance()

const enterDialog = () => {
  currentInstance.refs.fieldDialogNode.fieldDialogFrom.validate(valid => {
    if (valid) {
      dialogMiddle.value.fieldName = toUpperCase(
        dialogMiddle.value.fieldName
      )
      if (addFlag.value === 'add') {
        form.value.fields.push(dialogMiddle.value)
      }
      dialogFlag.value = false
    } else {
      return false
    }
  })
}
const closeDialog = () => {
  if (addFlag.value === 'edit') {
    dialogMiddle.value = bk.value
  }
  dialogFlag.value = false
}

const deleteField = (index) => {
  form.value.fields.splice(index, 1)
}

const autoCodeForm = ref(null)

const enterForm = async(isPreview) => {
  if (form.value.fields.length <= 0) {
    ElMessage({
      type: 'error',
      message: t('autoCode.errNoFields')
    })
    return false
  }
  if (
    form.value.fields.some(item => item.fieldName === form.value.structName)
  ) {
    ElMessage({
      type: 'error',
      message: t('autoCode.errSameFiledName')
    })
    return false
  }
  autoCodeForm.value.validate(async valid => {
    if (valid) {
      for (const key in form.value) {
        if (typeof form.value[key] === 'string') {
          form.value[key] = form.value[key].trim()
        }
      }
      form.value.structName = toUpperCase(form.value.structName)
      form.value.tableName = form.value.tableName.replace(' ', '')
      if (!form.value.tableName) {
        form.value.tableName = toSQLLine(toLowerCase(form.value.structName))
      }
      if (form.value.structName === form.value.abbreviation) {
        ElMessage({
          type: 'error',
          message: t('autoCode.errSameStructDescAbbr')
        })
        return false
      }
      form.value.humpPackageName = toSQLLine(form.value.packageName)
      if (isPreview) {
        const data = await preview(form.value)
        preViewCode.value = data.data.autoCode
        previewFlag.value = true
      } else {
        const data = await createTemp(form.value)
        if (data.headers?.success === 'false') {
          return
        } else {
          if (form.value.autoMoveFile) {
            ElMessage({
              type: 'success',
              message: t('autoCode.codeGenMoveSuccess')
            })
            return
          }
          ElMessage({
            type: 'success',
            message: t('autoCode.codeGenDownload')
          })
        }
        const blob = new Blob([data])
        const fileName = 'ginvueadmin.zip'
        if ('download' in document.createElement('a')) {
          // 不是IE浏览器
          const url = window.URL.createObjectURL(blob)
          const link = document.createElement('a')
          link.style.display = 'none'
          link.href = url
          link.setAttribute('download', fileName)
          document.body.appendChild(link)
          link.click()
          document.body.removeChild(link) // 下载完成移除元素
          window.URL.revokeObjectURL(url) // 释放掉blob对象
        } else {
          // IE 10+
          window.navigator.msSaveBlob(blob, fileName)
        }
      }
    } else {
      return false
    }
  })
}
const getDbFunc = async() => {
  const res = await getDB()
  if (res.code === 0) {
    dbOptions.value = res.data.dbs
  }
}
const getTableFunc = async() => {
  const res = await getTable({ dbName: dbform.value.dbName })
  if (res.code === 0) {
    tableOptions.value = res.data.tables
  }
  dbform.value.tableName = ''
}
const getColumnFunc = async() => {
  const gormModelList = ['id', 'created_at', 'updated_at', 'deleted_at']
  const res = await getColumn(dbform.value)
  if (res.code === 0) {
    const tbHump = toHump(dbform.value.tableName)
    form.value.structName = toUpperCase(tbHump)
    form.value.tableName = dbform.value.tableName
    form.value.packageName = tbHump
    form.value.abbreviation = tbHump
    form.value.description = tbHump + t('autoCode.table')
    form.value.autoCreateApiToSql = true
    form.value.autoMoveFile = true
    form.value.fields = []
    res.data.columns &&
          res.data.columns.forEach(item => {
            if (!gormModelList.some(gormfd => gormfd === item.columnName)) {
              const fbHump = toHump(item.columnName)
              form.value.fields.push({
                fieldName: toUpperCase(fbHump),
                fieldDesc: item.columnComment || fbHump + t('autoCode.field'),
                fieldType: fdMap.value[item.dataType],
                dataType: item.dataType,
                fieldJson: fbHump,
                dataTypeLong: item.dataTypeLong && item.dataTypeLong.split(',')[0],
                columnName: item.columnName,
                comment: item.columnComment,
                require: false,
                errorText: '',
                clearable: true,
                fieldSearchType: '',
                dictType: ''
              })
            }
          })
  }
}
const setFdMap = async() => {
  const fdTypes = ['string', 'int', 'bool', 'float64', 'time.Time']
  fdTypes.forEach(async fdtype => {
    const res = await getDict(fdtype)
    res && res.forEach(item => {
      fdMap.value[item.label] = fdtype
    })
  })
}
const getAutoCodeJson = async(id) => {
  const res = await getMeta({ id: Number(id) })
  if (res.code === 0) {
    form.value = JSON.parse(res.data.meta)
  }
}

const pkgs = ref([])
const getPkgs = async() => {
  const res = await getPackageApi()
  if (res.code === 0) {
    pkgs.value = res.data.pkgs
  }
}

const goPkgs = () => {
  router.push({ name: 'autoPkg' })
}

const init = () => {
  getDbFunc()
  setFdMap()
  getPkgs()
  const id = route.params.id
  if (id) {
    getAutoCodeJson(id)
  }
}
init()

watch(() => route.params.id, (id) => {
  if (route.name === 'autoCodeEdit') {
    init()
  }
})

</script>

<style scoped lang="scss">
  .previewCodeTool {
    display: flex;
    align-items: center;
    padding: 5px 0;
  }
.button-box {
  padding: 10px 20px;
  .el-button {
    margin-right: 20px;
    float: right;
  }
}
.auto-btn-list{
  margin-top: 16px;
}
.auto-icon{
  margin-left: 6px;
  color: #666;
  cursor: pointer;
}

</style>
