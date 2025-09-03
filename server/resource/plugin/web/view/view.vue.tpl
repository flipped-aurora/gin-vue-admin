{{- $global := . }}
{{- $templateID := printf "%s_%s" .Package .StructName }}
{{- if .IsAdd }}
// 请在搜索条件中增加如下代码
{{- range .Fields}}
    {{- if .FieldSearchType}}
{{ GenerateSearchFormItem .}}
    {{ end }}
{{ end }}


// 表格增加如下列代码

{{- range .Fields}}
    {{- if .Table}}
       {{ GenerateTableColumn . }}
    {{- end }}
{{- end }}

// 新增表单中增加如下代码
{{- range .Fields}}
   {{- if .Form}}
     {{ GenerateFormItem . }}
   {{- end }}
{{- end }}

// 查看抽屉中增加如下代码

{{- range .Fields}}
              {{- if .Desc }}
    {{ GenerateDescriptionItem . }}
              {{- end }}
            {{- end }}

// 字典增加如下代码
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}

// setOptions方法中增加如下调用

{{- range $index, $element := .DictTypes }}
    {{ $element }}Options.value = await getDictFunc('{{$element}}')
{{- end }}

// 基础formData结构（变量处和关闭表单处）增加如下字段
{{- range .Fields}}
          {{- if .Form}}
            {{ GenerateDefaultFormValue . }}
          {{- end }}
        {{- end }}
// 验证规则中增加如下字段

{{- range .Fields }}
        {{- if .Form }}
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
    {{- end }}



{{- if .HasDataSource }}
// 请引用
get{{.StructName}}DataSource,

//  获取数据源
const dataSource = ref({})
const getDataSourceFunc = async()=>{
  const res = await get{{.StructName}}DataSource()
  if (res.code === 0) {
    dataSource.value = res.data || []
  }
}
getDataSourceFunc()
{{- end }}

{{- else }}

{{- if not .OnlyTemplate}}
<template>
  <div>
  {{- if not .IsTree }}
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      {{- if .GvaModel }}
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
         <el-date-picker
                  v-model="searchInfo.createdAtRange"
                  class="!w-380px"
                  type="datetimerange"
                  range-separator="至"
                  start-placeholder="开始时间"
                  end-placeholder="结束时间"
                />
       </el-form-item>
      {{ end -}}
           {{- range .Fields}} {{- if .FieldSearchType}} {{- if not .FieldSearchHide }}
            {{ GenerateSearchFormItem .}}
           {{ end }}{{ end }}{{ end }}
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          {{- range .Fields}}  {{- if .FieldSearchType}} {{- if .FieldSearchHide }}
          {{ GenerateSearchFormItem .}}
          {{ end }}{{ end }}{{ end }}
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
  {{- end }}
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.add"{{ end }} type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.batchDelete"{{ end }} icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            {{ if .HasExcel -}}
            <ExportTemplate {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.exportTemplate"{{ end }} template-id="{{$templateID}}" />
            <ExportExcel {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.exportExcel"{{ end }} template-id="{{$templateID}}" filterDeleted/>
            <ImportExcel {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.importExcel"{{ end }} template-id="{{$templateID}}" @on-success="getTableData" />
            {{- end }}
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
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" {{ if .IsTree -}} min-{{- end -}}width="180">
            <template #default="scope">{{ "{{ formatDate(scope.row.CreatedAt) }}" }}</template>
        </el-table-column>
        {{ end }}
        {{- range .Fields}}
        {{- if .Table}}
            {{ GenerateTableColumn . }}
        {{- end }}
        {{- end }}
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            {{- if .IsTree }}
            <el-button {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.add"{{ end }} type="primary" link class="table-button" @click="openDialog(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>新增子节点</el-button>
            {{- end }}
            <el-button {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.info"{{ end }} type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.edit"{{ end }} type="primary" link icon="edit" class="table-button" @click="update{{.StructName}}Func(scope.row)">编辑</el-button>
            <el-button {{ if .IsTree }}v-if="!scope.row.children?.length"{{ end }} {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.delete"{{ end }} type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
                <span class="text-lg">{{"{{"}}type==='create'?'新增':'编辑'{{"}}"}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
          {{- if .IsTree }}
            <el-form-item label="父节点:" prop="parentID" >
                <el-tree-select
                    v-model="formData.parentID"
                    :data="[rootNode,...tableData]"
                    check-strictly
                    :render-after-expand="false"
                    :props="defaultProps"
                    clearable
                    style="width: 240px"
                    placeholder="根节点"
                />
            </el-form-item>
          {{- end }}
        {{- range .Fields}}
          {{- if .Form}}
             {{ GenerateFormItem . }}
          {{- end }}
       {{- end }}
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
            {{- if .IsTree }}
            <el-descriptions-item label="父节点">
                <el-tree-select
                  v-model="detailForm.parentID"
                  :data="[rootNode,...tableData]"
                  check-strictly
                  disabled
                  :render-after-expand="false"
                  :props="defaultProps"
                  clearable
                  style="width: 240px"
                  placeholder="根节点"
                />
            </el-descriptions-item>
            {{- end }}
            {{- range .Fields}}
              {{- if .Desc }}
                 {{ GenerateDescriptionItem . }}
              {{- end }}
            {{- end }}
            </el-descriptions>
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
} from '@/plugin/{{.Package}}/api/{{.PackageName}}'

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
import RichView from '@/components/richtext/rich-view.vue'
{{- end }}

{{- if .HasFile }}
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'
{{- end }}

{{- if .HasArray}}
// 数组控制组件
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'
{{- end }}

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
{{- if .AutoCreateBtnAuth }}
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
{{- end }}

{{if .HasExcel -}}
// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'
{{- end}}


defineOptions({
    name: '{{.StructName}}'
})

{{- if .AutoCreateBtnAuth }}
// 按钮权限实例化
    const btnAuth = useBtnAuth()
{{- end }}

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}
const formData = ref({
        {{- if .IsTree }}
            parentID:undefined,
        {{- end }}
        {{- range .Fields}}
          {{- if .Form}}
            {{ GenerateDefaultFormValue . }}
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
    {{- range .Fields }}
        {{- if .Form }}
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
    CreatedAt:"created_at",
    ID:"id",
    {{- range .Fields}}
     {{- if .Table}}
      {{- if and .Sort}}
        {{- if not (eq .ColumnName "")}}
            {{.FieldJson}}: '{{.ColumnName}}',
        {{- end}}
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

{{- if not .IsTree }}
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
{{- else }}
// 树选择器配置
const defaultProps = {
  children: "children",
  label: "{{ .TreeJson }}",
  value: "{{ .PrimaryField.FieldJson }}"
}

const rootNode = {
  {{ .PrimaryField.FieldJson }}: 0,
  {{ .TreeJson }}: '根节点',
  children: []
}

// 查询
const getTableData = async() => {
  const table = await get{{.StructName}}List()
  if (table.code === 0) {
    tableData.value = table.data || []
  }
}
{{- end }}

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
        formData.value = res.data
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
const openDialog = ({{- if .IsTree -}}row{{- end -}}) => {
    type.value = 'create'
    {{- if .IsTree }}
    formData.value.parentID = row ? row.{{.PrimaryField.FieldJson}} : undefined
    {{- end }}
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
    {{- range .Fields}}
      {{- if .Form}}
        {{ GenerateDefaultFormValue . }}
      {{- end }}
    {{- end }}
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
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
              btnLoading.value = false
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

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await find{{.StructName}}({ {{.PrimaryField.FieldJson}}: row.{{.PrimaryField.FieldJson}} })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


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
{{- else}}
<template>
<div>form</div>
</template>
<script setup>
defineOptions({
  name: '{{.StructName}}'
})
</script>
<style>
</style>
{{- end }}

{{- end }}
