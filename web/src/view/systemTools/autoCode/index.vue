<template>
  <div>
    <!-- 从数据库直接获取字段 -->
    <div class="gva-search-box">
      <el-collapse v-model="activeNames" style="margin-bottom:12px">
        <el-collapse-item name="1">
          <template #title>
            <div :style="{fontSize:'16px',paddingLeft:'20px'}">
              点这里从现有数据库创建代码
              <el-icon class="header-icon ">
                <pointer />
              </el-icon>
            </div>
          </template>
          <el-form ref="getTableForm" style="margin-top:24px" :inline="true" :model="dbform" label-width="120px">
            <el-form-item label="数据库名" prop="structName">
              <el-select v-model="dbform.dbName" filterable placeholder="请选择数据库" @change="getTableFunc">
                <el-option
                  v-for="item in dbOptions"
                  :key="item.database"
                  :label="item.database"
                  :value="item.database"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="表名" prop="structName">
              <el-select
                v-model="dbform.tableName"
                :disabled="!dbform.dbName"
                filterable
                placeholder="请选择表"
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
              <el-button size="small" type="primary" @click="getColumnFunc">使用此表创建</el-button>
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
    </div>
    <div class="gva-search-box">
      <!-- 初始版本自动化代码工具 -->
      <el-form ref="autoCodeForm" :rules="rules" :model="form" label-width="120px" :inline="true">
        <el-form-item label="Struct名称" prop="structName">
          <el-input v-model="form.structName" placeholder="首字母自动转换大写" />
        </el-form-item>
        <el-form-item label="TableName" prop="tableName">
          <el-input v-model="form.tableName" placeholder="指定表名（非必填）" />
        </el-form-item>
        <el-form-item label="Struct简称" prop="abbreviation">
          <el-input v-model="form.abbreviation" placeholder="简称会作为入参对象名和路由group" />
        </el-form-item>
        <el-form-item label="Struct中文名称" prop="description">
          <el-input v-model="form.description" placeholder="中文描述作为自动api描述" />
        </el-form-item>
        <el-form-item label="文件名称" prop="packageName">
          <el-input v-model="form.packageName" placeholder="生成文件的默认名称(建议为驼峰格式,首字母小写,如sysXxxXxxx)" @blur="toLowerCaseFunc(form,'packageName')" />
        </el-form-item>
        <el-form-item label="Package（包）" prop="packageName">
          <el-select v-model="form.package" style="width:194px">
            <el-option v-for="item in pkgs" :key="item.ID" :value="item.packageName" :label="item.packageName" />
          </el-select>
          <el-icon class="auto-icon" @click="getPkgs"><refresh /></el-icon>
          <el-icon class="auto-icon" @click="goPkgs"><document-add /></el-icon>
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip content="注：把自动生成的API注册进数据库" placement="bottom" effect="light">
              <div> 自动创建API </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoCreateApiToSql" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip content="注：自动迁移生成的文件到yaml配置的对应位置" placement="bottom" effect="light">
              <div> 自动移动文件 </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoMoveFile" />
        </el-form-item>
      </el-form>
    </div>
    <!-- 组件列表 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" @click="editAndAddField()">新增Field</el-button>
      </div>
      <el-table :data="form.fields">
        <el-table-column align="left" type="index" label="序列" width="100" />
        <el-table-column align="left" prop="fieldName" label="Field名" />
        <el-table-column align="left" prop="fieldDesc" label="中文名" />
        <el-table-column align="left" prop="fieldJson" label="FieldJson" />
        <el-table-column align="left" prop="fieldType" label="Field数据类型" width="130" />
        <el-table-column align="left" prop="dataTypeLong" label="数据库字段长度" width="130" />
        <el-table-column align="left" prop="columnName" label="数据库字段" width="130" />
        <el-table-column align="left" prop="comment" label="数据库字段描述" width="130" />
        <el-table-column align="left" prop="fieldSearchType" label="搜索条件" width="130" />
        <el-table-column align="left" prop="dictType" label="字典" width="130" />
        <el-table-column align="left" label="操作" width="300">
          <template #default="scope">
            <el-button
              size="small"
              type="text"
              icon="edit"
              @click="editAndAddField(scope.row)"
            >编辑</el-button>
            <el-button
              size="small"
              type="text"
              :disabled="scope.$index === 0"
              @click="moveUpField(scope.$index)"
            >上移</el-button>
            <el-button
              size="small"
              type="text"
              :disabled="(scope.$index + 1) === form.fields.length"
              @click="moveDownField(scope.$index)"
            >下移</el-button>
            <el-popover v-model:visible="scope.row.visible" placement="top">
              <p>确定删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="small" @click="deleteField(scope.$index)">确定</el-button>
              </div>
              <template #reference>
                <el-button size="small" type="text" icon="delete" @click="scope.row.visible = true">删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <!-- 组件列表 -->
      <div class="gva-btn-list justify-content-flex-end auto-btn-list">
        <el-button size="small" type="primary" @click="enterForm(true)">预览代码</el-button>
        <el-button size="small" type="primary" @click="enterForm(false)">生成代码</el-button>
      </div>
    </div>
    <!-- 组件弹窗 -->
    <el-dialog v-model="dialogFlag" title="组件内容">
      <FieldDialog v-if="dialogFlag" ref="fieldDialogNode" :dialog-middle="dialogMiddle" />
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="previewFlag">
      <template #title>
        <div class="previewCodeTool">
          <p>操作栏：</p>
          <el-button size="small" type="primary" @click="selectText">全选</el-button>
          <el-button size="small" type="primary" @click="copy">复制</el-button>
        </div>
      </template>
      <PreviewCodeDialog v-if="previewFlag" ref="previewNode" :preview-code="preViewCode" />
      <template #footer>
        <div class="dialog-footer" style="padding-top:14px;padding-right:14px">
          <el-button size="small" type="primary" @click="previewFlag = false">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

const fieldTemplate = {
  fieldName: '',
  fieldDesc: '',
  fieldType: '',
  dataType: '',
  fieldJson: '',
  columnName: '',
  dataTypeLong: '',
  comment: '',
  fieldSearchType: '',
  dictType: ''
}

import FieldDialog from '@/view/systemTools/autoCode/component/fieldDialog.vue'
import PreviewCodeDialog from '@/view/systemTools/autoCode/component/previewCodeDialg.vue'
import { toUpperCase, toHump, toSQLLine, toLowerCase } from '@/utils/stringFun'
import { createTemp, getDB, getTable, getColumn, preview, getMeta, getPackageApi } from '@/api/autoCode'
import { getDict } from '@/utils/dictionary'
import { ref, getCurrentInstance, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

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
  autoCreateApiToSql: false,
  autoMoveFile: false,
  fields: []
})
const rules = ref({
  structName: [
    { required: true, message: '请输入结构体名称', trigger: 'blur' }
  ],
  abbreviation: [
    { required: true, message: '请输入结构体简称', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入结构体描述', trigger: 'blur' }
  ],
  packageName: [
    {
      required: true,
      message: '文件名称：sysXxxxXxxx',
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
      message: '请填写至少一个field'
    })
    return false
  }
  if (
    form.value.fields.some(item => item.fieldName === form.value.structName)
  ) {
    ElMessage({
      type: 'error',
      message: '存在与结构体同名的字段'
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
          message: 'structName和struct简称不能相同'
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
              message: '自动化代码创建成功，自动移动成功'
            })
            return
          }
          ElMessage({
            type: 'success',
            message: '自动化代码创建成功，正在下载'
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
    form.value.description = tbHump + '表'
    form.value.autoCreateApiToSql = true
    form.value.fields = []
    res.data.columns &&
          res.data.columns.forEach(item => {
            if (!gormModelList.some(gormfd => gormfd === item.columnName)) {
              const fbHump = toHump(item.columnName)
              form.value.fields.push({
                fieldName: toUpperCase(fbHump),
                fieldDesc: item.columnComment || fbHump + '字段',
                fieldType: fdMap.value[item.dataType],
                dataType: item.dataType,
                fieldJson: fbHump,
                dataTypeLong: item.dataTypeLong && item.dataTypeLong.split(',')[0],
                columnName: item.columnName,
                comment: item.columnComment,
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

</script>

<script>

export default {
  name: 'AutoCode'
}
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
