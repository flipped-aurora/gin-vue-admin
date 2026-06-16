<template>
  <el-drawer :model-value="modelValue" size="520px" :show-close="false" @close="emit('update:modelValue', false)">
    <template #header>
      <div class="flex justify-between items-center w-full">
        <span class="text-lg">{{ localForm.id ? '编辑CLI' : '新增CLI' }}</span>
        <div>
          <el-button @click="emit('update:modelValue', false)">取消</el-button>
          <el-button type="primary" @click="submit">确定</el-button>
        </div>
      </div>
    </template>
    <el-form :model="localForm" label-width="100px">
      <el-form-item label="CLI名称"><el-input v-model="localForm.name" placeholder="内部唯一标识，例如 user-manager" /></el-form-item>
      <el-form-item label="主命令"><el-input v-model="localForm.command" placeholder="用户实际输入的命令，例如 opsctl；留空默认同CLI名称" /></el-form-item>
      <el-form-item label="显示名称"><el-input v-model="localForm.displayName" /></el-form-item>
      <el-form-item label="版本"><el-input v-model="localForm.version" /></el-form-item>
      <el-form-item label="状态">
        <el-select v-model="localForm.status" placeholder="请选择">
          <el-option label="启用" value="enabled" />
          <el-option label="禁用" value="disabled" />
        </el-select>
      </el-form-item>
      <el-form-item label="描述"><el-input v-model="localForm.description" type="textarea" /></el-form-item>
    </el-form>
  </el-drawer>
</template>

<script setup>
import { reactive, watch } from 'vue'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  formData: { type: Object, default: () => ({}) }
})
const emit = defineEmits(['update:modelValue', 'submit'])

const localForm = reactive({
  id: undefined,
  name: '',
  command: '',
  displayName: '',
  version: 'v1',
  status: 'enabled',
  description: ''
})

watch(() => props.formData, (value) => {
  localForm.id = value?.ID || value?.id
  localForm.name = value?.name || ''
  localForm.command = value?.command || ''
  localForm.displayName = value?.displayName || ''
  localForm.version = value?.version || 'v1'
  localForm.status = value?.status || 'enabled'
  localForm.description = value?.description || ''
}, { immediate: true, deep: true })

const submit = () => {
  emit('submit', { ...localForm })
}
</script>
