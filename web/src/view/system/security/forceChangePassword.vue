<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <el-card class="w-96">
      <template #header>
        <span class="text-base font-medium">密码已过期，请修改密码</span>
      </template>
      <el-form ref="formRef" :model="form" :rules="rules" label-width="90px">
        <el-form-item label="原密码" prop="password">
          <el-input v-model="form.password" type="password" show-password />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input v-model="form.newPassword" type="password" show-password />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirm">
          <el-input v-model="form.confirm" type="password" show-password />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="loading" @click="onSubmit">提交并重新登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { changePassword } from '@/api/user'
import { useUserStore } from '@/pinia/modules/user'

defineOptions({ name: 'ForceChangePassword' })

const userStore = useUserStore()
const formRef = ref(null)
const loading = ref(false)
const form = ref({ password: '', newPassword: '', confirm: '' })

const rules = {
  password: [{ required: true, message: '请输入原密码', trigger: 'blur' }],
  newPassword: [{ required: true, message: '请输入新密码', trigger: 'blur' }],
  confirm: [
    { required: true, message: '请确认新密码', trigger: 'blur' },
    {
      validator: (rule, value, cb) =>
        value === form.value.newPassword ? cb() : cb(new Error('两次密码不一致')),
      trigger: 'blur'
    }
  ]
}

const onSubmit = () => {
  formRef.value.validate(async (valid) => {
    if (!valid) return
    loading.value = true
    try {
      const res = await changePassword({
        password: form.value.password,
        newPassword: form.value.newPassword
      })
      if (res.code === 0) {
        ElMessage.success('修改成功，请重新登录')
        await userStore.LoginOut()
      }
    } finally {
      loading.value = false
    }
  })
}
</script>
