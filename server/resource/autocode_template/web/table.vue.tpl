<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
           {{- range .Fields}}  {{- if .FieldSearchType}} {{- if eq .FieldType "bool" }}
            <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clearable :placeholder="t('general.pleaseSelect')">
                <el-option
                    key="true"
                    :label="t('general.yes')"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    :label="t('general.no')"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
           {{- else if .DictType}}
           <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clearable placeholder="请选择" @clear="()=>{searchInfo.{{.FieldJson}}=undefined}">
              <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            {{- else}}
        <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">


        {{- if eq .FieldType "float64" "int"}}
            {{if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
            <el-input v-model.number="searchInfo.start{{.FieldName}}" placeholder="搜索条件（起）" />
            —
            <el-input v-model.number="searchInfo.end{{.FieldName}}" placeholder="搜索条件（止）" />
           {{- else}}
             {{- if .DictType}}
              <el-select v-model="searchInfo.{{.FieldJson}}" placeholder="请选择" style="width:100%" :clearable="true" >
               <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
             </el-select>
                    {{- else}}
             <el-input v-model.number="searchInfo.{{.FieldJson}}" placeholder="搜索条件" />
                    {{- end }}
          {{- end}}
        {{- else if eq .FieldType "time.Time"}}
            {{if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
            <template #label>
            <span>
              {{.FieldDesc}}
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.start{{.FieldName}}" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.end{{.FieldName}} ? time.getTime() > searchInfo.end{{.FieldName}}.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.end{{.FieldName}}" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.start{{.FieldName}} ? time.getTime() < searchInfo.start{{.FieldName}}.getTime() : false"></el-date-picker>
           {{- else}}
           <el-date-picker v-model="searchInfo.{{.FieldJson}}" type="datetime" placeholder="搜索条件"></el-date-picker>
          {{- end}}
        {{- else}}
         <el-input v-model="searchInfo.{{.FieldJson}}" placeholder="搜索条件" />
        {{- end}}

        </el-form-item>{{ end }}{{ end }}{{ end }}
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">{{ "{{ t('general.search') }}" }}</el-button>
          <el-button icon="refresh" @click="onReset">{{ "{{ t('general.reset') }}" }}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">{{ "{{ t('general.add') }}" }}</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>{{" {{ t('general.deleteConfirm') }}" }}</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">{{ "{{ t('general.cancel') }}" }}</el-button>
                <el-button type="primary" @click="onDelete">{{ "{{ t('general.confirm') }}" }}</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">{{ "{{ t('general.delete') }}" }}</el-button>
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
        {{- if .NeedSort}}
        @sort-change="sortChange"
        {{- end}}
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" :label="t('general.createdAt')" width="180">
            <template #default="scope">{{ "{{ formatDate(scope.row.CreatedAt) }}" }}</template>
        </el-table-column>
        {{- range .Fields}}
        {{- if .DictType}}
        <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
            <template #default="scope">
            {{"{{"}} filterDict(scope.row.{{.FieldJson}},{{.DictType}}Options) {{"}}"}}
            </template>
        </el-table-column>
        {{- else if eq .FieldType "bool" }}
        <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
            <template #default="scope">{{"{{"}} formatBoolean(scope.row.{{.FieldJson}}) {{"}}"}}</template>
        </el-table-column>
         {{- else if eq .FieldType "time.Time" }}
         <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" width="180">
            <template #default="scope">{{"{{"}} formatDate(scope.row.{{.FieldJson}}) {{"}}"}}</template>
         </el-table-column>
          {{- else if eq .FieldType "picture" }}
          <el-table-column label="{{.FieldDesc}}" width="200">
              <template #default="scope">
                <el-image style="width: 100px; height: 100px" :src="getUrl(scope.row.{{.FieldJson}})" fit="cover"/>
              </template>
          </el-table-column>
           {{- else if eq .FieldType "file" }}
                    <el-table-column label="{{.FieldDesc}}" width="200">
                        <template #default="scope">
                             <div class="file-list">
                               <el-tag v-for="file in scope.row.{{.FieldJson}}" :key="file.uid">{{"{{"}}file.name{{"}}"}}</el-tag>
                             </div>
                        </template>
                    </el-table-column>
        {{- else }}
        <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120" />
        {{- end }}
        {{- end }}
        <el-table-column align="left" :label="t('general.operations')">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="update{{.StructName}}Func(scope.row)">{{ "{{ t('general.change') }}" }}</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">{{ "{{ t('general.delete') }}" }}</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
    {{- range .Fields}}
        <el-form-item label="{{.FieldDesc}}:"  prop="{{.FieldJson}}" >
      {{- if eq .FieldType "bool" }}
          <el-switch v-model="formData.{{.FieldJson}}" active-color="#13ce66" inactive-color="#ff4949" :active-text="t('general.yes')" :inactive-text="t('general.no')" clearable ></el-switch>
      {{- end }}
      {{- if eq .FieldType "string" }}
          <el-input v-model="formData.{{.FieldJson}}" :clearable="{{.Clearable}}"  :placeholder="t('general.pleaseEnter')" />
      {{- end }}
      {{- if eq .FieldType "int" }}
      {{- if .DictType}}
          <el-select v-model="formData.{{ .FieldJson }}" :placeholder="t('general.pleaseSelect')" style="width:100%" :clearable="{{.Clearable}}" >
            <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
          </el-select>
      {{- else }}
          <el-input v-model.number="formData.{{ .FieldJson }}" :clearable="{{.Clearable}}" :placeholder="t('general.pleaseEnter')" />
      {{- end }}
      {{- end }}
      {{- if eq .FieldType "time.Time" }}
          <el-date-picker v-model="formData.{{ .FieldJson }}" type="date" style="width:100%" :placeholder="t('general.selectDate')" :clearable="{{.Clearable}}"  />
      {{- end }}
      {{- if eq .FieldType "float64" }}
          <el-input-number v-model="formData.{{ .FieldJson }}"  style="width:100%" :precision="2" :clearable="{{.Clearable}}"  />
      {{- end }}
      {{- if eq .FieldType "enum" }}
            <el-select v-model="formData.{{ .FieldJson }}" placeholder="请选择" style="width:100%" :clearable="{{.Clearable}}" >
               <el-option v-for="item in [{{.DataTypeLong}}]" :key="item" :label="item" :value="item" />
            </el-select>
      {{- end }}
      {{- if eq .FieldType "picture" }}
            <SelectImage v-model="formData.{{ .FieldJson }}" />
      {{- end }}
      {{- if eq .FieldType "file" }}
            <SelectFile v-model="formData.{{ .FieldJson }}" />
      {{- end }}
        </el-form-item>
      {{- end }}
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">{{ "{{ t('general.close') }}" }}</el-button>
          <el-button type="primary" @click="enterDialog">{{ "{{ t('general.confirm') }}" }}</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: '{{.StructName}}'
}
</script>

<script setup>
import {
  create{{.StructName}},
  delete{{.StructName}},
  delete{{.StructName}}ByIds,
  update{{.StructName}},
  find{{.StructName}},
  get{{.StructName}}List
} from '@/api/{{.PackageName}}'

{{- if or .HasPic .HasFile}}
import { getUrl } from '@/utils/image'
{{- end }}
{{- if .HasPic }}
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
{{- end }}

{{- if .HasFile }}
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'
{{- end }}

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage


// 自动化生成的字典（可能为空）以及字段
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}
const formData = ref({
        {{- range .Fields}}
        {{- if eq .FieldType "bool" }}
        {{.FieldJson}}: false,
        {{- end }}
        {{- if eq .FieldType "string" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "int" }}
        {{.FieldJson}}: {{- if .DictType }} undefined{{ else }} 0{{- end }},
        {{- end }}
        {{- if eq .FieldType "time.Time" }}
        {{.FieldJson}}: new Date(),
        {{- end }}
        {{- if eq .FieldType "float64" }}
        {{.FieldJson}}: 0,
        {{- end }}
        {{- end }}
        })

// 验证规则
const rule = reactive({
    {{- range .Fields }}
            {{- if eq .Require true }}
               {{.FieldJson }} : [{
                   required: true,
                   message: '{{ .ErrorText }}',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }],
            {{- end }}
    {{- end }}
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
  {{- range .Fields }}
    {{- if .FieldSearchType}} 
      {{- if eq .FieldType "time.Time" }}
        {{.FieldJson }} : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.start{{.FieldName}} && !searchInfo.value.end{{.FieldName}}) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.start{{.FieldName}} && searchInfo.value.end{{.FieldName}}) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.start{{.FieldName}} && searchInfo.value.end{{.FieldName}} && (searchInfo.value.start{{.FieldName}}.getTime() === searchInfo.value.end{{.FieldName}}.getTime() || searchInfo.value.start{{.FieldName}}.getTime() > searchInfo.value.end{{.FieldName}}.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
      {{- end }}
    {{- end }}
  {{- end }}
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

{{- if .NeedSort}}
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}
{{- end}}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    {{- range .Fields}}{{- if eq .FieldType "bool" }}
    if (searchInfo.value.{{.FieldJson}} === ""){
        searchInfo.value.{{.FieldJson}}=null
    }{{ end }}{{ end }}
    getTableData()
  })
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
  const table = await get{{.StructName}}List({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
{{- range $index, $element := .DictTypes }}
    {{ $element }}Options.value = await getDictFunc('{{$element}}')
{{- end }}
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
    ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
        confirmButtonText: t('general.confirm'),
        cancelButtonText: t('general.cancel'),
        type: 'warning'
    }).then(() => {
            delete{{.StructName}}Func(row)
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
          message: t('general.selectDataToDelete')
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await delete{{.StructName}}ByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess')
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
const update{{.StructName}}Func = async(row) => {
    const res = await find{{.StructName}}({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.re{{.Abbreviation}}
        dialogFormVisible.value = true
    }
}


// 删除行
const delete{{.StructName}}Func = async (row) => {
    const res = await delete{{.StructName}}({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: t('general.deleteSuccess')
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
    {{- range .Fields}}
        {{- if eq .FieldType "bool" }}
        {{.FieldJson}}: false,
        {{- end }}
        {{- if eq .FieldType "string" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "int" }}
        {{.FieldJson}}: {{- if .DictType }} undefined{{ else }} 0{{- end }},
        {{- end }}
        {{- if eq .FieldType "time.Time" }}
        {{.FieldJson}}: new Date(),
        {{- end }}
        {{- if eq .FieldType "float64" }}
        {{.FieldJson}}: 0,
        {{- end }}
        {{- end }}
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await create{{.StructName}}(formData.value)
                  break
                case 'update':
                  res = await update{{.StructName}}(formData.value)
                  break
                default:
                  res = await create{{.StructName}}(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: t('general.createUpdateSuccess')
                })
                closeDialog()
                getTableData()
              }
      })
}
{{if .HasFile }}
const downloadFile = (url) => {
    window.open(getUrl(url), '_blank')
}
{{end}}
</script>

<style>
{{if .HasFile }}
.file-list{
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}
{{end}}
</style>
