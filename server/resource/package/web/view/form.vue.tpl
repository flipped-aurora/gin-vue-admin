{{- if .IsAdd }}
// 新增表单中增加如下代码
{{- range .Fields}}
     {{- if .Form}}
        {{ GenerateFormItem . }}
     {{- end }}
{{- end }}

// 字典增加如下代码
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}

// init方法中增加如下调用

{{- range $index, $element := .DictTypes }}
    {{ $element }}Options.value = await getDictFunc('{{$element}}')
{{- end }}

// 基础formData结构增加如下字段
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
const dataSource = ref([])
const getDataSourceFunc = async()=>{
  const res = await get{{.StructName}}DataSource()
  if (res.code === 0) {
    dataSource.value = res.data
  }
}
getDataSourceFunc()
{{- end }}
{{- else }}
{{- if not .OnlyTemplate }}
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
      {{- if .Form }}
        {{ GenerateFormItem . }}
      {{- end }}
      {{- end }}
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  {{- if .HasDataSource }}
    get{{.StructName}}DataSource,
  {{- end }}
  {{- if .IsTree }}
    get{{.StructName}}List,
  {{- end }}
  create{{.StructName}},
  update{{.StructName}},
  find{{.StructName}}
} from '@/api/{{.Package}}/{{.PackageName}}'

defineOptions({
    name: '{{.StructName}}Form'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
{{- if .HasPic }}
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
{{- end }}

{{- if .HasFile }}
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'
{{- end }}

{{- if .HasRichText }}
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
{{- end }}

{{- if .HasArray}}
// 数组控制组件
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'
{{- end }}


const route = useRoute()
const router = useRouter()

{{- if .IsTree }}
const tableData = ref([])

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

const getTableData = async() => {
  const table = await get{{.StructName}}List()
  if (table.code === 0) {
    tableData.value = table.data || []
  }
}

getTableData()

{{- end }}

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}
const formData = ref({
        {{- if .IsTree }}
            parentID: undefined,
        {{- end }}
        {{- range .Fields}}
          {{- if .Form }}
            {{ GenerateDefaultFormValue . }}
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
               }],
            {{- end }}
    {{- end }}
})

const elFormRef = ref()

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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await find{{.StructName}}({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    {{- range $index, $element := .DictTypes }}
    {{ $element }}Options.value = await getDictFunc('{{$element}}')
    {{- end }}
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
{{- else }}
<template>
<div>form</div>
</template>
<script setup>
</script>
<style>
</style>
{{- end }}
{{- end }}
