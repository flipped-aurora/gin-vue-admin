<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      {{- if .GvaModel }}
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
      {{ end -}}
           {{- range .Fields}}  {{- if .FieldSearchType}} {{- if eq .FieldType "bool" }}
            <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
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
            <el-input v-model.number="searchInfo.start{{.FieldName}}" placeholder="最小值" />
            —
            <el-input v-model.number="searchInfo.end{{.FieldName}}" placeholder="最大值" />
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
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="{{.PrimaryField.FieldJson}}"
        @selection-change="handleSelectionChange"
        {{- if .NeedSort}}
        @sort-change="sortChange"
        {{- end}}
        >
        <el-table-column type="selection" width="55" />
        {{ if .GvaModel }}
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ "{{ formatDate(scope.row.CreatedAt) }}" }}</template>
        </el-table-column>
        {{ end }}
        {{- range .FrontFields}}
        {{- if .CheckDataSource }}
        <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
          <template #default="scope">
                {{if eq .DataSource.Association 2}}
                    <el-tag v-for="(item,key) in filterDataSource(dataSource.{{.FieldJson}},scope.row.{{.FieldJson}})" :key="key">
                        {{ "{{ item }}" }}
                    </el-tag>
                {{ else }}
                    <span>{{"{{"}} filterDataSource(dataSource.{{.FieldJson}},scope.row.{{.FieldJson}}) {{"}}"}}</span>
                {{ end }}
         </template>
         </el-table-column>
        {{- else if .DictType}}
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
         <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="180">
            <template #default="scope">{{"{{"}} formatDate(scope.row.{{.FieldJson}}) {{"}}"}}</template>
         </el-table-column>
          {{- else if eq .FieldType "picture" }}
          <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="200">
              <template #default="scope">
                <el-image style="width: 100px; height: 100px" :src="getUrl(scope.row.{{.FieldJson}})" fit="cover"/>
              </template>
          </el-table-column>
           {{- else if eq .FieldType "pictures" }}
           <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image v-for="(item,index) in scope.row.{{.FieldJson}}" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
           {{- else if eq .FieldType "video" }}
           <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="200">
              <template #default="scope">
               <video
                  style="width: 100px; height: 100px"
                  muted
                  preload="metadata"
                  >
                    <source :src="getUrl(scope.row.{{.FieldJson}}) + '#t=1'">
                  </video>
              </template>
           </el-table-column>
           {{- else if eq .FieldType "richtext" }}
                      <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="200">
                         <template #default="scope">
                            [富文本内容]
                         </template>
                      </el-table-column>
           {{- else if eq .FieldType "file" }}
                    <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="200">
                        <template #default="scope">
                             <div class="file-list">
                               <el-tag v-for="file in scope.row.{{.FieldJson}}" :key="file.uid">{{"{{"}}file.name{{"}}"}}</el-tag>
                             </div>
                        </template>
                    </el-table-column>
         {{- else if eq .FieldType "json" }}
          <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="200">
              <template #default="scope">
                  [JSON]
              </template>
          </el-table-column>
        {{- else }}
        <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120" />
        {{- end }}
        {{- end }}
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="update{{.StructName}}Func(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{"{{"}}type==='create'?'添加':'修改'{{"}}"}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        {{- range .FrontFields}}
            <el-form-item label="{{.FieldDesc}}:"  prop="{{.FieldJson}}" >
          {{- if .CheckDataSource}}
            <el-select {{if eq .DataSource.Association 2}} multiple {{ end }} v-model="formData.{{.FieldJson}}" placeholder="请选择{{.FieldDesc}}" style="width:100%" :clearable="{{.Clearable}}" >
              <el-option v-for="(item,key) in dataSource.{{.FieldJson}}" :key="key" :label="item.label" :value="item.value" />
            </el-select>
          {{- else }}
          {{- if eq .FieldType "bool" }}
              <el-switch v-model="formData.{{.FieldJson}}" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
          {{- end }}
          {{- if eq .FieldType "string" }}
          {{- if .DictType}}
              <el-select v-model="formData.{{ .FieldJson }}" placeholder="请选择{{.FieldDesc}}" style="width:100%" :clearable="{{.Clearable}}" >
                <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
              </el-select>
          {{- else }}
              <el-input v-model="formData.{{.FieldJson}}" :clearable="{{.Clearable}}"  placeholder="请输入{{.FieldDesc}}" />
          {{- end }}
          {{- end }}
          {{- if eq .FieldType "richtext" }}
              <RichEdit v-model="formData.{{.FieldJson}}"/>
          {{- end }}
          {{- if eq .FieldType "json" }}
              // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.{{.FieldJson}} 后端会按照json的类型进行存取
              {{"{{"}} formData.{{.FieldJson}} {{"}}"}}
          {{- end }}
           {{- if eq .FieldType "array" }}
           <el-tag v-for="(item,key) in formData.{{.FieldJson}}" :key="key">
              {{ "{{ item }}" }}
           </el-tag>
           {{- end }}
          {{- if eq .FieldType "int" }}
              <el-input v-model.number="formData.{{ .FieldJson }}" :clearable="{{.Clearable}}" placeholder="请输入{{.FieldDesc}}" />
          {{- end }}
          {{- if eq .FieldType "time.Time" }}
              <el-date-picker v-model="formData.{{ .FieldJson }}" type="date" style="width:100%" placeholder="选择日期" :clearable="{{.Clearable}}"  />
          {{- end }}
          {{- if eq .FieldType "float64" }}
              <el-input-number v-model="formData.{{ .FieldJson }}"  style="width:100%" :precision="2" :clearable="{{.Clearable}}"  />
          {{- end }}
          {{- if eq .FieldType "enum" }}
                <el-select v-model="formData.{{ .FieldJson }}" placeholder="请选择{{.FieldDesc}}" style="width:100%" :clearable="{{.Clearable}}" >
                   <el-option v-for="item in [{{.DataTypeLong}}]" :key="item" :label="item" :value="item" />
                </el-select>
          {{- end }}
          {{- if eq .FieldType "picture" }}
                <SelectImage
                 v-model="formData.{{ .FieldJson }}"
                 file-type="image"
                />
          {{- end }}
          {{- if eq .FieldType "pictures" }}
                <SelectImage
                 multiple
                 v-model="formData.{{ .FieldJson }}"
                 file-type="image"
                 />
          {{- end }}
          {{- if eq .FieldType "video" }}
                <SelectImage
                v-model="formData.{{ .FieldJson }}"
                file-type="video"
                />
           {{- end }}
          {{- if eq .FieldType "file" }}
                <SelectFile v-model="formData.{{ .FieldJson }}" />
          {{- end }}
          {{- end }}
            </el-form-item>
          {{- end }}
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  {{- if .HasDataSource }}
    get{{.StructName}}DataSource,
  {{- end }}
  create{{.StructName}},
  delete{{.StructName}},
  delete{{.StructName}}ByIds,
  update{{.StructName}},
  find{{.StructName}},
  get{{.StructName}}List
} from '@/api/{{.Package}}/{{.PackageName}}'

{{- if or .HasPic .HasFile}}
import { getUrl } from '@/utils/image'
{{- end }}
{{- if .HasPic }}
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
{{- end }}

{{- if .HasRichText }}
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
{{- end }}


{{- if .HasFile }}
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'
{{- end }}

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: '{{.StructName}}'
})

// 自动化生成的字典（可能为空）以及字段
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}
const formData = ref({
        {{- range .FrontFields}}
        {{- if eq .FieldType "bool" }}
        {{.FieldJson}}: false,
        {{- end }}
        {{- if eq .FieldType "string" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "richtext" }}
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
        {{- if eq .FieldType "picture" }}
        {{.FieldJson}}: "",
        {{- end }}
        {{- if eq .FieldType "video" }}
        {{.FieldJson}}: "",
        {{- end }}
        {{- if eq .FieldType "pictures" }}
        {{.FieldJson}}: [],
        {{- end }}
        {{- if eq .FieldType "file" }}
        {{.FieldJson}}: [],
        {{- end }}
        {{- if eq .FieldType "json" }}
        {{.FieldJson}}: {},
        {{- end }}
        {{- if eq .FieldType "array" }}
        {{.FieldJson}}: [],
        {{- end }}
        {{- end }}
        })

{{- if .HasDataSource }}
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await get{{.StructName}}DataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()
{{- end }}



// 验证规则
const rule = reactive({
    {{- range .FrontFields }}
            {{- if eq .Require true }}
               {{.FieldJson }} : [{
                   required: true,
                   message: '{{ .ErrorText }}',
                   trigger: ['input','blur'],
               },
               {{- if eq .FieldType "string" }}
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              {{- end }}
              ],
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
  {{- range .FrontFields }}
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
  const sortMap = {
    {{- range .FrontFields}}
      {{- if and .Sort}}
        {{- if not (eq .ColumnName "")}}
            {{.FieldJson}}: '{{.ColumnName}}',
        {{- end}}
      {{- end}}
    {{- end}}
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
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
    {{- range .FrontFields}}{{- if eq .FieldType "bool" }}
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
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            delete{{.StructName}}Func(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const {{.PrimaryField.FieldJson}}s = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          {{.PrimaryField.FieldJson}}s.push(item.{{.PrimaryField.FieldJson}})
        })
      const res = await delete{{.StructName}}ByIds({ {{.PrimaryField.FieldJson}}s })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === {{.PrimaryField.FieldJson}}s.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const update{{.StructName}}Func = async(row) => {
    const res = await find{{.StructName}}({ {{.PrimaryField.FieldJson}}: row.{{.PrimaryField.FieldJson}} })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.re{{.Abbreviation}}
        dialogFormVisible.value = true
    }
}


// 删除行
const delete{{.StructName}}Func = async (row) => {
    const res = await delete{{.StructName}}({ {{.PrimaryField.FieldJson}}: row.{{.PrimaryField.FieldJson}} })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
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
    {{- range .FrontFields}}
        {{- if eq .FieldType "bool" }}
        {{.FieldJson}}: false,
        {{- end }}
        {{- if eq .FieldType "string" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "richtext" }}
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
        {{- if eq .FieldType "picture" }}
        {{.FieldJson}}: "",
        {{- end }}
        {{- if eq .FieldType "video" }}
        {{.FieldJson}}: "",
        {{- end }}
        {{- if eq .FieldType "pictures" }}
        {{.FieldJson}}: [],
        {{- end }}
        {{- if eq .FieldType "file" }}
        {{.FieldJson}}: [],
        {{- end }}
        {{- if eq .FieldType "json" }}
        {{.FieldJson}}: {},
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
                  message: '创建/更改成功'
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

.fileBtn{
  margin-bottom: 10px;
}

.fileBtn:last-child{
  margin-bottom: 0;
}
{{end}}
</style>
