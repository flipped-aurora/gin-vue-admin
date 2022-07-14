<template>
  <SysDialog v-model:value="dialogVisible" :title="dialogTitle" alert-title="表单验证规则只会写入前端验证" height="200px" @close="handleClose">
    <template #body>
      <el-form
        :model="formData"
        label-width="100px"
        class="grid-form"
      >
        <el-form-item label="是否必填">
          <el-switch v-model="formData.require" />
        </el-form-item>
        <el-form-item label="显示类型">
          <el-select v-model="formData.visibleType" size="small" placeholder="请选择显示类型">
            <el-option label="字符" value="text" />
            <el-option label="密码" value="password" />
            <el-option label="数字" value="number" />
          </el-select>
        </el-form-item>
        <el-form-item label="校验失败内容">
          <el-input v-model="formData.errorText" />
        </el-form-item>
        <el-form-item label="是否可清除">
          <el-switch v-model="formData.clearable" />
        </el-form-item>

      </el-form>
    </template>
    <template #actions>
      <div>
        <el-button size="small" @click="handleClose">取消</el-button>
        <el-button size="small" type="primary" @click="handleSubmit">确 定</el-button>
      </div>
    </template>
  </SysDialog>
</template>

<script setup>
import SysDialog from '@/components/sysDialog/index.vue'
import { reactive, ref, watch, toRaw } from 'vue'
const dialogVisible = ref(false)
const dialogTitle = ref('验证规则')
const emit = defineEmits(['close', 'confirm'])
const prop = defineProps({
  value: {
    type: Boolean,
    default: false,
  },
})
const formData = reactive({
  visibleType: 'text', // 显示类型
  require: true, // 是否必填
  errorText: '验证失败！', // 验证失败提示
  clearable: true, // 是否可清除
})


watch(
  () => prop.value,
  (val) => {
    dialogVisible.value = val
  }, {
    immediate: true,
  })

const handleClose = () => {
  dialogVisible.value = false
  emit('close')
}

const handleSubmit = () => {
  emit('confirm', {
    ...toRaw(formData)
  })
  handleClose()
}

const initDialog = ({
  title = '验证规则',
  value
}) => {
  dialogTitle.value = title
  if (value) {
    console.log('传入的value', value)
    formData.visibleType = value.visibleType
    formData.require = value.require
    formData.errorText = value.errorText
    formData.ruleType = value.ruleType
    formData.clearable = value.clearable
    formData.verificationRules = value.verificationRules.toString()
    formData.lenLimit = value.lenLimit
  }
}

defineExpose({
  initDialog
})


</script>

<style scoped>
.grid-form{
  display: grid;
  grid-template-columns: 1fr 1fr;
  align-items: center;

}
:deep(.el-input__inner),
:deep(.el-input--suffix){
  height: 30px !important;
}
.flex{
  display: flex;
  align-items: center;
}
</style>
