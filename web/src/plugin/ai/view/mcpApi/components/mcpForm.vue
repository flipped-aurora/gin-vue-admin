<template>
  <el-drawer :model-value="modelValue" size="520px" :show-close="false" @close="emit('update:modelValue', false)">
    <template #header>
      <div class="flex justify-between items-center w-full">
        <span class="text-base">{{ localForm.id ? '编辑MCP' : '新增MCP' }}</span>
        <div>
          <el-button @click="emit('update:modelValue', false)">取消</el-button>
          <el-button type="primary" @click="submit">确定</el-button>
        </div>
      </div>
    </template>
    <el-form ref="formRef" :model="localForm" :rules="rules" label-width="100px">
      <el-form-item label="名称" prop="name">
        <el-input v-model="localForm.name" placeholder="内部唯一标识，例如 user-manager" />
      </el-form-item>
      <el-form-item label="显示名称">
        <el-input v-model="localForm.displayName" placeholder="对外展示名称" />
      </el-form-item>
      <el-form-item label="版本">
        <el-input v-model="localForm.version" placeholder="例如 v1" />
      </el-form-item>
      <el-form-item label="状态">
        <el-select v-model="localForm.status" placeholder="请选择">
          <el-option label="启用" value="enabled" />
          <el-option label="禁用" value="disabled" />
        </el-select>
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="localForm.description" type="textarea" />
      </el-form-item>
    </el-form>
  </el-drawer>
</template>

<script setup>
import { reactive, ref, watch } from 'vue'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  formData: { type: Object, default: () => ({}) }
})
const emit = defineEmits(['update:modelValue', 'submit'])

const formRef = ref(null)

const localForm = reactive({
  id: undefined,
  name: '',
  displayName: '',
  version: 'v1',
  status: 'enabled',
  description: ''
})

const rules = {
  name: [{ required: true, message: '请输入名称', trigger: 'blur' }]
}

watch(() => props.formData, (value) => {
  localForm.id = value?.ID || value?.id
  localForm.name = value?.name || ''
  localForm.displayName = value?.displayName || ''
  localForm.version = value?.version || 'v1'
  localForm.status = value?.status || 'enabled'
  localForm.description = value?.description || ''
}, { immediate: true, deep: true })

const submit = async () => {
  if (formRef.value) {
    await formRef.value.validate(valid => {
      if (valid) {
        emit('submit', { ...localForm })
      }
    })
  } else {
    emit('submit', { ...localForm })
  }
}
</script>
