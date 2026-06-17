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
      <el-divider content-position="left">AI Skill（供 Claude Code 等 AI 助手使用）</el-divider>
      <el-form-item label="Skill名"><el-input v-model="localForm.skillName" placeholder="留空默认为 <主命令>-cli" /></el-form-item>
      <el-form-item label="Skill描述"><el-input v-model="localForm.skillDescription" type="textarea" placeholder="告诉 AI 何时使用这个 CLI，例如：需要在终端查询或操作用户数据时使用" /></el-form-item>
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
  description: '',
  skillName: '',
  skillDescription: ''
})

watch(() => props.formData, (value) => {
  localForm.id = value?.ID || value?.id
  localForm.name = value?.name || ''
  localForm.command = value?.command || ''
  localForm.displayName = value?.displayName || ''
  localForm.version = value?.version || 'v1'
  localForm.status = value?.status || 'enabled'
  localForm.description = value?.description || ''
  localForm.skillName = value?.skillName || ''
  localForm.skillDescription = value?.skillDescription || ''
}, { immediate: true, deep: true })

const submit = () => {
  emit('submit', { ...localForm })
}
</script>
