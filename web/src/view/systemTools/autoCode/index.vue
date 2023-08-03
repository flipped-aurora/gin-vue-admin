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
            <el-form-item label="业务库" prop="selectDBtype">
              <template #label>
                <el-tooltip content="注：需要提前到db-list自行配置多数据库，如未配置需配置后重启服务方可使用。（此处可选择对应库表，可理解为从哪个库选择表）" placement="bottom" effect="light">
                  <div> 业务库 <el-icon><QuestionFilled /></el-icon> </div>
                </el-tooltip>
              </template>
              <el-select v-model="dbform.businessDB" clearable style="width:194px" placeholder="选择业务库" @change="getDbFunc">
                <el-option
                  v-for="item in dbList"
                  :key="item.aliasName"
                  :value="item.aliasName"
                  :label="item.aliasName"
                  :disabled="item.disable"
                >
                  <div>
                    <span>{{ item.aliasName }}</span>
                    <span style="float:right;color:#8492a6;font-size:13px">{{ item.dbName }}</span>
                  </div>
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item :label="t('autoCode.dbName')" prop="structName">
              <el-select v-model="dbform.dbName" clearable filterable :placeholder="t('autoCode.selectDB')" @change="getTableFunc">
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
              <el-button type="primary" @click="getColumnFunc">{{ t('autoCode.createUsingTable') }}</el-button>
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
        <el-form-item label="业务库" prop="businessDB">
          <template #label>
            <el-tooltip content="注：需要提前到db-list自行配置多数据库，此项为空则会使用gva本库创建自动化代码(global.GVA_DB),填写后则会创建指定库的代码(global.MustGetGlobalDBByDBName(dbname))" placement="bottom" effect="light">
              <div> 业务库 <el-icon><QuestionFilled /></el-icon> </div>
            </el-tooltip>
          </template>
          <el-select
            v-model="form.businessDB"
            style="width:194px"
            placeholder="选择业务库"
          >
            <el-option
              v-for="item in dbList"
              :key="item.aliasName"
              :value="item.aliasName"
              :label="item.aliasName"
              :disabled="item.disable"
            >
              <div>
                <span>{{ item.aliasName }}</span>
                <span style="float:right;color:#8492a6;font-size:13px">{{ item.dbName }}</span>
              </div>
            </el-option>
          </el-select>
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
        <el-button type="primary" @click="editAndAddField()">{{ t('autoCode.addField') }}</el-button>
      </div>
      <el-table :data="form.fields">
        <el-table-column align="left" type="index" :label="t('autoCode.fieldIndex')" width="60" />
        <el-table-column align="left" prop="fieldName" :label="t('autoCode.fieldName')" width="160">
          <template #default="{row}">
            <el-input v-model="row.fieldName" />
          </template>
        </el-table-column>
        <el-table-column align="left" prop="fieldDesc" :label="t('autoCode.fieldDesc')" width="160">
          <template #default="{row}">
            <el-input v-model="row.fieldDesc" />
          </template>
        </el-table-column>
        <el-table-column align="left" prop="require" label="必填">
          <template #default="{row}"> <el-checkbox v-model="row.require" /></template>
        </el-table-column>
        <el-table-column align="left" prop="sort" label="排序">
          <template #default="{row}"> <el-checkbox v-model="row.sort" /> </template>
        </el-table-column>
        <el-table-column align="left" prop="fieldJson" width="160px" label="字段Json">
          <template #default="{row}">
            <el-input v-model="row.fieldJson" />
          </template>
        </el-table-column>
        <el-table-column align="left" prop="fieldType" :label="t('autoCode.fieldDataType')" width="160">
          <template #default="{row}">
            <el-select
              v-model="row.fieldType"
              style="width:100%"
              placeholder="请选择字段类型"
              clearable
            >
              <el-option
                v-for="item in typeOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column align="left" prop="dataTypeLong" :label="t('autoCode.fieldLen')" width="160">
          <template #default="{row}">
            <el-input v-model="row.dataTypeLong" />
          </template>
        </el-table-column>
        <el-table-column align="left" prop="columnName" :label="t('autoCode.columnName')" width="160">
          <template #default="{row}">
            <el-input v-model="row.columnName" />
          </template>
        </el-table-column>
        <el-table-column align="left" prop="comment" :label="t('autoCode.comment')" width="160">
          <template #default="{row}">
            <el-input v-model="row.columnName" />
          </template>
        </el-table-column>
        <el-table-column align="left" prop="fieldSearchType" :label="t('general.searchCriteria')" width="130">
          <template #default="{row}">
            <el-select
              v-model="row.fieldSearchType"
              style="width:100%"
              placeholder="请选择字段查询条件"
              clearable
            >
              <el-option
                v-for="item in typeSearchOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                :disabled="
                  (row.fieldType!=='string'&&item.value==='LIKE')||
                    ((row.fieldType!=='int'&&row.fieldType!=='time.Time'&&row.fieldType!=='float64')&&(item.value==='BETWEEN' || item.value==='NOT BETWEEN'))"
              />
            </el-select>
          </template>

        </el-table-column>
        <el-table-column align="left" :lable="t('general.operations')" width="300" fixed="right">
          <template #default="scope">
            <el-button

              type="primary"
              link
              icon="edit"
              @click="editAndAddField(scope.row)"
            >{{ t('general.edit') }}</el-button>
            <el-button

              type="primary"
              link
              :disabled="scope.$index === 0"
              @click="moveUpField(scope.$index)"
            >{{ t('autoCode.moveUp') }}</el-button>
            <el-button

              type="primary"
              link
              :disabled="(scope.$index + 1) === form.fields.length"
              @click="moveDownField(scope.$index)"
            >{{ t('autoCode.moveDown') }}</el-button>
            <el-popover v-model="scope.row.visible" placement="top">
              <p>{{ t('autoCode.confirmDelete') }}</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="scope.row.visible = false">{{ t('general.cancel') }}</el-button>
                <el-button type="primary" @click="deleteField(scope.$index)">{{ t('general.confirm') }}</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete" @click="scope.row.visible = true">{{ t('general.delete') }}</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <!-- 组件列表 -->
      <div class="gva-btn-list justify-content-flex-end auto-btn-list">
        <el-button type="primary" @click="enterForm(true)">{{ t('autoCode.codePreview') }}</el-button>
        <el-button type="primary" @click="enterForm(false)">{{ t('autoCode.generateCode') }}</el-button>
      </div>
    </div>
    <!-- 组件弹窗 -->
    <el-dialog v-model="dialogFlag" width="70%" :title="t('autoCode.componentContent')">
      <FieldDialog v-if="dialogFlag" ref="fieldDialogNode" :dialog-middle="dialogMiddle" :typeOptions="typeOptions" :typeSearchOptions="typeSearchOptions" />
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">{{ t('general.close') }}</el-button>
          <el-button type="primary" @click="enterDialog">{{ t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="previewFlag">
      <template #header>
        <div class="previewCodeTool">
          <p>操作栏：</p>
          <el-button type="primary" @click="selectText">{{ t('general.selectAll') }}</el-button>
          <el-button type="primary" @click="copy">{{ t('autoCode.copy') }}</el-button>
        </div>
      </template>
      <PreviewCodeDialog v-if="previewFlag" ref="preview" :preview-code="preViewCode" />
      <template #footer>
        <div class="dialog-footer" style="padding-top:14px;padding-right:14px">
          <el-button type="primary" @click="previewFlag = false">{{ t('general.confirm') }}</el-button>
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
import { ref, getCurrentInstance, reactive, watch, toRaw } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const typeOptions = ref([
  {
    label: '字符串',
    value: 'string'
  },
  {
    label: '整型',
    value: 'int'
  },
  {
    label: '布尔值',
    value: 'bool'
  },
  {
    label: '浮点型',
    value: 'float64'
  },
  {
    label: '时间',
    value: 'time.Time'
  },
  {
    label: '枚举',
    value: 'enum'
  },
  {
    label: '单图片（字符串）',
    value: 'picture',
  },
  {
    label: '文件（json字符串）',
    value: 'file',
  },
  {
    label: '多图片（开发中）',
    value: 'pictures',
    disabled: true
  }
])

const typeSearchOptions = ref([
  {
    label: '=',
    value: '='
  },
  {
    label: '<>',
    value: '<>'
  },
  {
    label: '>',
    value: '>'
  },
  {
    label: '<',
    value: '<'
  },
  {
    label: 'LIKE',
    value: 'LIKE'
  },
  {
    label: 'BETWEEN',
    value: 'BETWEEN'
  },
  {
    label: 'NOT BETWEEN',
    value: 'NOT BETWEEN'
  }
])

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
  sort: false,
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
  businessDB: '',
  dbName: '',
  tableName: ''
})
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
  businessDB: '',
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

  if (form.value.package === form.value.abbreviation) {
    ElMessage({
      type: 'error',
      message: 'package和结构体简称不可同名'
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

const dbList = ref([])
const dbOptions = ref([])

const getDbFunc = async() => {
  dbform.value.dbName = ''
  dbform.value.tableName = ''
  const res = await getDB({ businessDB: dbform.value.businessDB })
  if (res.code === 0) {
    dbOptions.value = res.data.dbs
    dbList.value = res.data.dbList
  }
}
const getTableFunc = async() => {
  const res = await getTable({ businessDB: dbform.value.businessDB, dbName: dbform.value.dbName })
  if (res.code === 0) {
    tableOptions.value = res.data.tables
  }
  dbform.value.tableName = ''
}

const getColumnFunc = async() => {
  const gormModelList = ['id', 'created_at', 'updated_at', 'deleted_at']
  const res = await getColumn(dbform.value)
  if (res.code === 0) {
    let dbtype = ''
    if (dbform.value.businessDB !== '') {
      const dbtmp = dbList.value.find(item => item.aliasName === dbform.value.businessDB)
      const dbraw = toRaw(dbtmp)
      dbtype = dbraw.dbtype
    }
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
                columnName: dbtype === 'oracle' ? item.columnName.toUpperCase() : item.columnName,
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

watch(() => route.params.id, () => {
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
