<template>
  <div>
    <!-- 从数据库直接获取字段 -->
    <div class="gva-search-box">
      <el-collapse v-model="activeNames" style="margin-bottom:12px">
        <el-collapse-item name="1">
          <template #title>
            <div :style="{fontSize:'16px',paddingLeft:'20px'}">
              {{ $t('autoCode.existDB') }}
              <el-icon class="header-icon ">
                <pointer />
              </el-icon>
            </div>
          </template>
          <el-form ref="getTableForm" style="margin-top:24px" :inline="true" :model="dbform" label-width="120px">
            <el-form-item :label="$t('autoCode.dbName')" prop="structName">
              <el-select v-model="dbform.dbName" filterable :placeholder="$t('autoCode.selectDB')" @change="getTable">
                <el-option
                  v-for="item in dbOptions"
                  :key="item.database"
                  :label="item.database"
                  :value="item.database"
                />
              </el-select>
            </el-form-item>
            <el-form-item :label="$t('autoCode.tableName')" prop="structName">
              <el-select
                v-model="dbform.tableName"
                :disabled="!dbform.dbName"
                filterable
                :placeholder="$t('autoCode.selectTable')"
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
              <el-button size="mini" type="primary" @click="getColumn">{{ $t('autoCode.createUsingTable') }}</el-button>
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
    </div>
    <div class="gva-search-box">
      <!-- 初始版本自动化代码工具 -->
      <el-form ref="autoCodeForm" :rules="rules" :model="form" label-width="180px" :inline="true">
        <el-form-item :label="$t('autoCode.structName')" prop="structName">
          <el-input v-model="form.structName" :placeholder="$t('autoCode.structNameNote')" />
        </el-form-item>
        <el-form-item :label="$t('autoCode.tableName')" prop="tableName">
          <el-input v-model="form.tableName" :placeholder="$t('autoCode.tableNameNote')" />
        </el-form-item>
        <el-form-item :label="$t('autoCode.structAbbreviation')" prop="abbreviation">
          <el-input v-model="form.abbreviation" :placeholder="$t('autoCode.structAbbreviationNote')" />
        </el-form-item>
        <el-form-item :label="$t('autoCode.structChineseName')" prop="description">
          <el-input v-model="form.description" :placeholder="$t('autoCode.structChineseNameNote')" />
        </el-form-item>
        <el-form-item :label="$t('autoCode.fileName')" prop="packageName">
          <el-input v-model="form.packageName" :placeholder="$t('autoCode.fileNameNote')" @blur="toLowerCase(form,'packageName')" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip :content="$t('autoCode.autoAPIDBTip')" placement="bottom" effect="light">
              <div> {{ $t('autoCode.autoAPIDBCreate') }} </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoCreateApiToSql" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip :content="$t('autoCode.autoMoveFilesTip')" placement="bottom" effect="light">
              <div> {{ $t('autoCode.autoMoveFiles') }} </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoMoveFile" />
        </el-form-item>
      </el-form>
    </div>
    <!-- 组件列表 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" @click="editAndAddField()">{{ $t('autoCode.addField') }}</el-button>
      </div>
      <el-table :data="form.fields">
        <el-table-column align="left" type="index" :label="$t('autoCode.fieldIndex')" width="100" />
        <el-table-column align="left" prop="fieldName" :label="$t('autoCode.fieldName')" width="120" />
        <el-table-column align="left" prop="fieldDesc" :label="$t('autoCode.fieldDesc')" width="120" />
        <el-table-column align="left" prop="fieldJson" :label="$t('autoCode.fieldJson')" width="110" />
        <el-table-column align="left" prop="fieldType" :label="$t('autoCode.fieldDataType')" width="130" />
        <el-table-column align="left" prop="dataTypeLong" :label="$t('autoCode.fieldLen')" width="130" />
        <el-table-column align="left" prop="columnName" :label="$t('autoCode.columnName')" width="130" />
        <el-table-column align="left" prop="comment" :label="$t('autoCode.comment')" width="130" />
        <el-table-column align="left" prop="fieldSearchType" :label="$t('autoCode.searchType')" width="130" />
        <el-table-column align="left" prop="dictType" :label="$t('autoCode.dictionary')" width="130" />
        <el-table-column align="left" :lable="$t('general.operations')" width="300">
          <template #default="scope">
            <el-button
              size="mini"
              type="text"
              icon="edit"
              @click="editAndAddField(scope.row)"
            >{{ $t('general.edit') }}</el-button>
            <el-button
              size="mini"
              type="text"
              :disabled="scope.$index === 0"
              @click="moveUpField(scope.$index)"
            >{{ $t('autoCode.moveUp') }}</el-button>
            <el-button
              size="mini"
              type="text"
              :disabled="(scope.$index + 1) === form.fields.length"
              @click="moveDownField(scope.$index)"
            >{{ $t('autoCode.moveDown') }}</el-button>
            <el-popover :visible="scope.row.visible" placement="top">
              <p>{{ $t('autoCode.confirmDelete') }}</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="scope.row.visible = false">{{ $t('general.cancel') }}</el-button>
                <el-button type="primary" size="mini" @click="deleteField(scope.$index)">{{ $t('general.confirm') }}</el-button>
              </div>
              <template #reference>
                <el-button size="mini" type="text" icon="delete">{{ $t('general.delete') }}</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <!-- 组件列表 -->
      <div class="gva-btn-list justify-content-flex-end auto-btn-list">
        <el-button size="mini" type="primary" @click="enterForm(true)">{{ $t('autoCode.codePreview') }}</el-button>
        <el-button size="mini" type="primary" @click="enterForm(false)">{{ $t('autoCode.generateCode') }}</el-button>
      </div>
    </div>
    <!-- 组件弹窗 -->
    <el-dialog v-model="dialogFlag" :title="$t('autoCode.componentContent')">
      <FieldDialog v-if="dialogFlag" ref="fieldDialog" :dialog-middle="dialogMiddle" />
      <template #footer>
        <div class="dialog-footer">
          <el-button size="mini" @click="closeDialog">{{ $t('general.close') }}</el-button>
          <el-button size="mini" type="primary" @click="enterDialog">{{ $t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="previewFlag">
      <template #title>
        <div class="previewCodeTool">
          <p>{{ $t('autoCode.actionBar') }}</p>
          <el-button size="mini" type="primary" @click="selectText">{{ $t('general.selectAll') }}</el-button>
          <el-button size="mini" type="primary" @click="copy">{{ $t('autoCode.copy') }}</el-button>
        </div>
      </template>
      <PreviewCodeDialog v-if="previewFlag" ref="preview" :preview-code="preViewCode" />
      <template #footer>
        <div class="dialog-footer" style="padding-top:14px;padding-right:14px">
          <el-button size="small" type="primary" @click="previewFlag = false">{{ $t('general.confirm') }}</el-button>
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
import { toUpperCase, toHump, toSQLLine, toLowerCase } from '@/utils/stringFun'
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
          { required: true, message: this.$t('autoCode.entStructName'), trigger: 'blur' }
        ],
        abbreviation: [
          { required: true, message: this.$t('autoCode.entStructAbbreviation'), trigger: 'blur' }
        ],
        description: [
          { required: true, message: this.$t('autoCode.entStructDesc'), trigger: 'blur' }
        ],
        packageName: [
          {
            required: true,
            message: this.$t('autoCode.entFileName'),
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
    toLowerCase(form, key) {
      form[key] = toLowerCase(form[key])
    },
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
          message: this.$t('autoCode.errNoFields')
        })
        return false
      }
      if (
        this.form.fields.some(item => item.fieldName === this.form.structName)
      ) {
        this.$message({
          type: 'error',
          message: this.$t('autoCode.errSameFiledName')
        })
        return false
      }
      this.$refs.autoCodeForm.validate(async valid => {
        if (valid) {
          for (const key in this.form) {
            if (typeof this.form[key] === 'string') {
              this.form[key] = this.form[key].trim()
            }
          }
          this.form.structName = toUpperCase(this.form.structName)
          if (this.form.tableName) { this.form.tableName = this.form.tableName.replace(' ', '') }
          if (this.form.structName === this.form.abbreviation) {
            this.$message({
              type: 'error',
              message: this.$t('autoCode.errSameStructDescAbbr')
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
                  message: this.$t('autoCode.codeGenMoveSuccess')
                })
                return
              }
              this.$message({
                type: 'success',
                message: this.$t('autoCode.codeGenDownload')
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
        this.form.description = tbHump + this.$t('autoCode.table')
        this.form.autoCreateApiToSql = true
        this.form.fields = []
        res.data.columns &&
          res.data.columns.forEach(item => {
            if (!gormModelList.some(gormfd => gormfd === item.columnName)) {
              const fbHump = toHump(item.columnName)
              this.form.fields.push({
                fieldName: toUpperCase(fbHump),
                fieldDesc: item.columnComment || fbHump + this.$t('autoCode.field'),
                fieldType: this.fdMap[item.dataType],
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
