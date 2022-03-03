<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="20%">
      {{- range .Fields}}
        <el-form-item label="{{.FieldDesc}}:">
      {{- if eq .FieldType "bool" }}
          <el-switch v-model="formData.{{.FieldJson}}" active-color="#13ce66" inactive-color="#ff4949" :active-text="t('general.yes')" :inactive-text="t('general.no')" clearable ></el-switch>
      {{- end }}
      {{- if eq .FieldType "string" }}
          <el-input v-model="formData.{{.FieldJson}}" clearable :placeholder="t('general.pleaseEnter')" />
      {{- end }}
      {{- if eq .FieldType "int" }}
      {{- if .DictType }}
          <el-select v-model="formData.{{ .FieldJson }}" :placeholder="t('general.pleaseSelect')" clearable>
            <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
          </el-select>
      {{- else }}
          <el-input v-model.number="formData.{{ .FieldJson }}" clearable :placeholder="t('general.pleaseEnter')" />
      {{- end }}
      {{- end }}
      {{- if eq .FieldType "time.Time" }}
          <el-date-picker v-model="formData.{{ .FieldJson }}" type="date" :placeholder="t('general.selectDate')" clearable></el-date-picker>
      {{- end }}
      {{- if eq .FieldType "float64" }}
          <el-input-number v-model="formData.{{ .FieldJson }}" :precision="2" clearable></el-input-number>
      {{- end }}
        </el-form-item>
      {{- end }}
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">{{ "{{ t('general.save') }}" }}</el-button>
          <el-button size="mini" type="primary" @click="back">{{ "{{ t('general.back') }}" }}</el-button>
        </el-form-item>
      </el-form>
    </div>
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
  update{{.StructName}},
  find{{.StructName}}
} from '@/api/{{.PackageName}}'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const route = useRoute()
const router = useRouter()
const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await find{{.StructName}}({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.re{{.Abbreviation}}
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
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
