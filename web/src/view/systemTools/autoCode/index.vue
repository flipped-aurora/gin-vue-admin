<template>
  <div>
    <!-- 从数据库直接获取字段 -->
    <div class="gva-search-box">
      <el-collapse v-model="activeNames" style="margin-bottom:12px">
        <el-collapse-item name="1">
          <template #title>
            <div :style="{fontSize:'16px',paddingLeft:'20px'}">
              点这里从现有数据库创建代码
              <i class="header-icon el-icon-thumb" />
            </div>
          </template>
          <el-form ref="getTableForm" style="margin-top:24px" :inline="true" :model="dbform" label-width="120px">
            <el-form-item label="数据库名" prop="structName">
              <el-select v-model="dbform.dbName" filterable placeholder="请选择数据库" @change="getTable">
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
              <el-button size="mini" type="primary" @click="getColumn">使用此表创建</el-button>
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
          <el-input v-model="form.packageName" placeholder="生成文件的默认名称(建议为驼峰格式,首字母小写,如sysXxxXxxx)" />
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
            <el-tooltip content="注：自动迁移生成的文件到ymal配置的对应位置" placement="bottom" effect="light">
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
        <el-button size="mini" type="primary" @click="editAndAddField()">新增Field</el-button>
      </div>
      <el-table :data="form.fields">
        <el-table-column align="left" type="index" label="序列" width="100" />
        <el-table-column align="left" prop="fieldName" label="Field名" />
        <el-table-column align="left" prop="fieldDesc" label="中文名" />
        <el-table-column align="left" prop="fieldJson" label="FieldJson" />
        <el-table-column align="left" prop="fieldType" label="Field数据类型" width="130" />
        <el-table-column align="left" prop="dataType" label="数据库字段类型" width="130" />
        <el-table-column align="left" prop="dataTypeLong" label="数据库字段长度" width="130" />
        <el-table-column align="left" prop="columnName" label="数据库字段" width="130" />
        <el-table-column align="left" prop="comment" label="数据库字段描述" width="130" />
        <el-table-column align="left" prop="fieldSearchType" label="搜索条件" width="130" />
        <el-table-column align="left" prop="dictType" label="字典" width="130" />
        <el-table-column align="left" label="操作" width="300">
          <template #default="scope">
            <el-button
              size="mini"
              type="text"
              icon="el-icon-edit"
              @click="editAndAddField(scope.row)"
            >编辑</el-button>
            <el-button
              size="mini"
              type="text"
              :disabled="scope.$index === 0"
              @click="moveUpField(scope.$index)"
            >上移</el-button>
            <el-button
              size="mini"
              type="text"
              :disabled="(scope.$index + 1) === form.fields.length"
              @click="moveDownField(scope.$index)"
            >下移</el-button>
            <el-popover :visible="scope.row.visible" placement="top">
              <p>确定删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="mini" @click="deleteField(scope.$index)">确定</el-button>
              </div>
              <template #reference>
                <el-button size="mini" type="text" icon="el-icon-delete">删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <!-- 组件列表 -->
      <div class="gva-btn-list justify-content-flex-end auto-btn-list">
        <el-button size="mini" type="primary" @click="enterForm(true)">预览代码</el-button>
        <el-button size="mini" type="primary" @click="enterForm(false)">生成代码</el-button>
      </div>
    </div>
    <!-- 组件弹窗 -->
    <el-dialog v-model="dialogFlag" title="组件内容">
      <FieldDialog v-if="dialogFlag" ref="fieldDialog" :dialog-middle="dialogMiddle" />
      <template #footer>
        <div class="dialog-footer">
          <el-button size="mini" @click="closeDialog">取 消</el-button>
          <el-button size="mini" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="previewFlag">
      <template #title>
        <div class="previewCodeTool">
          <p>操作栏：</p>
          <el-button size="mini" type="primary" @click="selectText">全选</el-button>
          <el-button size="mini" type="primary" @click="copy">复制</el-button>
        </div>
      </template>
      <PreviewCodeDialog v-if="previewFlag" ref="preview" :preview-code="preViewCode" />
      <template #footer>
        <div class="dialog-footer" style="padding-top:14px;padding-right:14px">
          <el-button size="small" type="primary" @click="previewFlag = false">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
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
import { toUpperCase, toHump, toSQLLine } from '@/utils/stringFun'
import { createTemp, getDB, getTable, getColumn, preview, getMeta } from '@/api/autoCode'
import { getDict } from '@/utils/dictionary'

export default {
  name: 'AutoCode',
  components: {
    FieldDialog,
    PreviewCodeDialog
  },
  data() {
    return {
      activeNames: [''],
      preViewCode: {},
      dbform: {
        dbName: '',
        tableName: ''
      },
      dbOptions: [],
      tableOptions: [],
      addFlag: '',
      fdMap: {},
      form: {
        structName: '',
        tableName: '',
        packageName: '',
        abbreviation: '',
        description: '',
        autoCreateApiToSql: false,
        autoMoveFile: false,
        fields: []
      },
      rules: {
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
        ]
      },
      dialogMiddle: {},
      bk: {},
      dialogFlag: false,
      previewFlag: false
    }
  },
  created() {
    this.getDb()
    this.setFdMap()
    const id = this.$route.params.id
    if (id) {
      this.getAutoCodeJson(id)
    }
  },
  methods: {
    selectText() {
      this.$refs.preview.selectText()
    },
    copy() {
      this.$refs.preview.copy()
    },
    editAndAddField(item) {
      this.dialogFlag = true
      if (item) {
        this.addFlag = 'edit'
        this.bk = JSON.parse(JSON.stringify(item))
        this.dialogMiddle = item
      } else {
        this.addFlag = 'add'
        this.dialogMiddle = JSON.parse(JSON.stringify(fieldTemplate))
      }
    },
    moveUpField(index) {
      if (index === 0) {
        return
      }
      const oldUpField = this.form.fields[index - 1]
      this.form.fields.splice(index - 1, 1)
      this.form.fields.splice(index, 0, oldUpField)
    },
    moveDownField(index) {
      const fCount = this.form.fields.length
      if (index === fCount - 1) {
        return
      }
      const oldDownField = this.form.fields[index + 1]
      this.form.fields.splice(index + 1, 1)
      this.form.fields.splice(index, 0, oldDownField)
    },
    enterDialog() {
      this.$refs.fieldDialog.$refs.fieldDialogFrom.validate(valid => {
        if (valid) {
          this.dialogMiddle.fieldName = toUpperCase(
            this.dialogMiddle.fieldName
          )
          if (this.addFlag === 'add') {
            this.form.fields.push(this.dialogMiddle)
          }
          this.dialogFlag = false
        } else {
          return false
        }
      })
    },
    closeDialog() {
      if (this.addFlag === 'edit') {
        this.dialogMiddle = this.bk
      }
      this.dialogFlag = false
    },
    deleteField(index) {
      this.form.fields.splice(index, 1)
    },
    async enterForm(isPreview) {
      if (this.form.fields.length <= 0) {
        this.$message({
          type: 'error',
          message: '请填写至少一个field'
        })
        return false
      }
      if (
        this.form.fields.some(item => item.fieldName === this.form.structName)
      ) {
        this.$message({
          type: 'error',
          message: '存在与结构体同名的字段'
        })
        return false
      }
      this.$refs.autoCodeForm.validate(async valid => {
        if (valid) {
          this.form.structName = toUpperCase(this.form.structName)
          if (this.form.tableName) { this.form.tableName = this.form.tableName.replace(' ', '') }
          if (this.form.structName === this.form.abbreviation) {
            this.$message({
              type: 'error',
              message: 'structName和struct简称不能相同'
            })
            return false
          }
          this.form.humpPackageName = toSQLLine(this.form.packageName)
          if (isPreview) {
            const data = await preview(this.form)
            this.preViewCode = data.data.autoCode
            this.previewFlag = true
          } else {
            const data = await createTemp(this.form)
            if (data.headers?.success === 'false') {
              return
            } else {
              if (this.form.autoMoveFile) {
                this.$message({
                  type: 'success',
                  message: '自动化代码创建成功，自动移动成功'
                })
                return
              }
              this.$message({
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
    },
    async getDb() {
      const res = await getDB()
      if (res.code === 0) {
        this.dbOptions = res.data.dbs
      }
    },
    async getTable() {
      const res = await getTable({ dbName: this.dbform.dbName })
      if (res.code === 0) {
        this.tableOptions = res.data.tables
      }
      this.dbform.tableName = ''
    },
    async getColumn() {
      const gormModelList = ['id', 'created_at', 'updated_at', 'deleted_at']
      const res = await getColumn(this.dbform)
      if (res.code === 0) {
        const tbHump = toHump(this.dbform.tableName)
        this.form.structName = toUpperCase(tbHump)
        this.form.tableName = this.dbform.tableName
        this.form.packageName = tbHump
        this.form.abbreviation = tbHump
        this.form.description = tbHump + '表'
        this.form.autoCreateApiToSql = true
        this.form.fields = []
        res.data.columns &&
          res.data.columns.forEach(item => {
            if (!gormModelList.some(gormfd => gormfd === item.columnName)) {
              const fbHump = toHump(item.columnName)
              this.form.fields.push({
                fieldName: toUpperCase(fbHump),
                fieldDesc: item.columnComment || fbHump + '字段',
                fieldType: this.fdMap[item.dataType],
                dataType: item.dataType,
                fieldJson: fbHump,
                dataTypeLong: item.dataTypeLong,
                columnName: item.columnName,
                comment: item.columnComment,
                fieldSearchType: '',
                dictType: ''
              })
            }
          })
      }
    },
    async setFdMap() {
      const fdTypes = ['string', 'int', 'bool', 'float64', 'time.Time']
      fdTypes.forEach(async fdtype => {
        const res = await getDict(fdtype)
        res && res.forEach(item => {
          this.fdMap[item.label] = fdtype
        })
      })
    },
    async getAutoCodeJson(id) {
      const res = await getMeta({ id: Number(id) })
      if (res.code === 0) {
        this.form = JSON.parse(res.data.meta)
      }
    }
  }
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
</style>
