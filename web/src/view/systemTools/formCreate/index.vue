<template>
  <div class="form-designer-container">
    <fc-designer ref="designer" :config="config" height="calc(100vh - 160px)">
      <template #handle>
        <el-button type="primary" size="small" plain @click="exportVueTemplate">
          解析为 Vue 原生标签
        </el-button>
      </template>
    </fc-designer>

    <el-dialog v-model="dialogVisible" title="生成的 Vue 模板代码" width="70%" top="5vh">
      <el-input 
        type="textarea" 
        :rows="25" 
        v-model="vueCode" 
        readonly 
        class="code-input"
        resize="none"
      />
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">关闭</el-button>
          <el-button type="primary" @click="copyCode">一键复制</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import FcDesigner from '@form-create/designer'

  defineOptions({
    name: 'FormGenerator'
  })

  const designer = ref(null)
  const dialogVisible = ref(false)
  const vueCode = ref('')

  const config = {
    fieldReadonly: false,
    useTemplate: true
  }

  const kebabCase = (str) => {
    return str.replace(/([A-Z])/g, '-$1').toLowerCase()
  }

  const generateVueCode = (rules, options) => {
    let formDataInit = []
    let formRules = []

    const parseRule = (rule) => {
      if (rule.type === 'row') {
        const propsStr = rule.props ? Object.entries(rule.props).map(([k, v]) => `:${k}="${v}"`).join(' ') : ''
        let childrenStr = rule.children ? rule.children.map(c => parseRule(c)).join('\n') : ''
        return `\n    <el-row ${propsStr}>${childrenStr}\n    </el-row>`
      }
      if (rule.type === 'col') {
        const propsStr = rule.props ? Object.entries(rule.props).map(([k, v]) => `:${k}="${v}"`).join(' ') : ''
        let childrenStr = rule.children ? rule.children.map(c => parseRule(c)).join('\n') : ''
        return `\n      <el-col ${propsStr}>${childrenStr}\n      </el-col>`
      }

      if (!rule.field) return ''

      let tag = rule.type
      
      const typeMap = {
        input: 'el-input',
        inputNumber: 'el-input-number',
        select: 'el-select',
        radio: 'el-radio-group',
        checkbox: 'el-checkbox-group',
        switch: 'el-switch',
        timePicker: 'el-time-picker',
        datePicker: 'el-date-picker',
        slider: 'el-slider',
        rate: 'el-rate',
        colorPicker: 'el-color-picker',
        cascader: 'el-cascader',
        upload: 'el-upload'
      }

      const elTag = typeMap[tag] || (tag.startsWith('el-') ? tag : `el-${tag}`)

      let propsStr = ''
      if (rule.props) {
        for (const [key, value] of Object.entries(rule.props)) {
          if (value === null || value === undefined) continue
          if (typeof value === 'boolean') {
            propsStr += value ? ` ${kebabCase(key)}` : ` :${kebabCase(key)}="false"`
          } else if (typeof value === 'string') {
            propsStr += ` ${kebabCase(key)}="${value}"`
          } else {
            propsStr += ` :${kebabCase(key)}='${JSON.stringify(value)}'`
          }
        }
      }

      let innerContent = ''
      if (rule.options && Array.isArray(rule.options)) {
        if (tag === 'select') {
          innerContent = rule.options.map(opt => `\n        <el-option label="${opt.label}" value="${opt.value}" />`).join('') + '\n      '
        } else if (tag === 'radio') {
          innerContent = rule.options.map(opt => `\n        <el-radio label="${opt.value}">${opt.label}</el-radio>`).join('') + '\n      '
        } else if (tag === 'checkbox') {
          innerContent = rule.options.map(opt => `\n        <el-checkbox label="${opt.value}">${opt.label}</el-checkbox>`).join('') + '\n      '
        }
      }

      let initVal = rule.value !== undefined ? rule.value : (tag === 'checkbox' ? [] : null)
      formDataInit.push(`  ${rule.field}: ${JSON.stringify(initVal)}`)

      if (rule.$required || (rule.effect && rule.effect.required)) {
        formRules.push(`  ${rule.field}: [{ required: true, message: '${rule.title}不能为空', trigger: 'blur' }]`)
      } else if (rule.validate) {
        formRules.push(`  ${rule.field}: ${JSON.stringify(rule.validate)}`)
      }

      return `
    <el-form-item label="${rule.title}" prop="${rule.field}">
      <${elTag} v-model="formData.${rule.field}"${propsStr}>${innerContent}</${elTag}>
    </el-form-item>`
    }

    const formItems = rules.map(parseRule).join('')

    const formConfig = options.form || {}
    let formPropsStr = []
    if (formConfig.labelWidth) formPropsStr.push(`label-width="${formConfig.labelWidth}"`)
    if (formConfig.size) formPropsStr.push(`size="${formConfig.size}"`)
    if (formConfig.labelPosition) formPropsStr.push(`label-position="${formConfig.labelPosition}"`)
    if (formConfig.hideRequiredAsterisk) formPropsStr.push(`hide-required-asterisk`)

    // 8. 拼装成标准的 <template> 和 <script setup> 闭环代码
    return `<template>
  <div>
    <el-form ref="formRef" :model="formData" :rules="rules" ${formPropsStr.join(' ')}>
${formItems}
      <el-form-item>
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'

const formRef = ref(null)

const formData = reactive({
${formDataInit.join(',\n')}
})

const rules = reactive({
${formRules.join(',\n')}
})

const submitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate((valid) => {
    if (valid) {
      ElMessage.success('表单校验通过，准备提交')
      console.log('提交的数据: ', formData)
    } else {
      ElMessage.error('表单校验失败')
    }
  })
}

const resetForm = () => {
  if (!formRef.value) return
  formRef.value.resetFields()
}
<\/script>
`
  }

  const exportVueTemplate = () => {
    const rules = designer.value.getRule()
    const options = designer.value.getOption()
    
    vueCode.value = generateVueCode(rules, options)
    dialogVisible.value = true
  }

  const copyCode = async () => {
    try {
      await navigator.clipboard.writeText(vueCode.value)
      ElMessage.success('代码已成功复制到剪贴板！')
      dialogVisible.value = false
    } catch (err) {
      ElMessage.error('复制失败，请手动选择复制')
    }
  }
</script>

<style scoped>

</style>

