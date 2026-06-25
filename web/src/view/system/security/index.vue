<template>
  <div class="p-4">
    <el-card>
      <template #header>
        <div class="flex items-center justify-between">
          <span class="text-base font-medium">安全配置</span>
          <el-button type="primary" :loading="saving" @click="onSave">保存</el-button>
        </div>
      </template>

      <el-tabs v-model="activeTab">
        <el-tab-pane label="验证码" name="captcha">
          <el-form :model="form" label-width="160px" class="max-w-2xl">
            <el-form-item label="错误N次后出验证码">
              <el-input-number v-model="form.captchaOpen" :min="0" />
              <span class="ml-2 text-gray-400 text-sm">0 = 每次都需要验证码</span>
            </el-form-item>
            <el-form-item label="计数缓存超时(秒)">
              <el-input-number v-model="form.captchaTimeout" :min="1" />
            </el-form-item>
            <el-form-item label="验证码长度">
              <el-input-number v-model="form.keyLong" :min="1" />
            </el-form-item>
            <el-form-item label="验证码宽度">
              <el-input-number v-model="form.imgWidth" :min="1" />
            </el-form-item>
            <el-form-item label="验证码高度">
              <el-input-number v-model="form.imgHeight" :min="1" />
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <el-tab-pane label="密码复杂度" name="password">
          <el-form :model="form" label-width="160px" class="max-w-2xl">
            <el-form-item label="最小长度">
              <el-input-number v-model="form.pwdMinLength" :min="1" />
            </el-form-item>
            <el-form-item label="需大写字母">
              <el-switch v-model="form.pwdRequireUpper" />
            </el-form-item>
            <el-form-item label="需小写字母">
              <el-switch v-model="form.pwdRequireLower" />
            </el-form-item>
            <el-form-item label="需数字">
              <el-switch v-model="form.pwdRequireDigit" />
            </el-form-item>
            <el-form-item label="需特殊字符">
              <el-switch v-model="form.pwdRequireSpecial" />
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <el-tab-pane label="限流" name="limit">
          <el-form :model="form" label-width="160px" class="max-w-2xl">
            <el-form-item label="开启限流">
              <el-switch v-model="form.limitEnable" />
            </el-form-item>
            <el-form-item label="窗口(秒)">
              <el-input-number v-model="form.limitWindow" :min="1" />
            </el-form-item>
            <el-form-item label="窗口内最大次数">
              <el-input-number v-model="form.limitCount" :min="1" />
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <el-tab-pane label="失败锁定" name="lock">
          <el-form :model="form" label-width="160px" class="max-w-2xl">
            <el-form-item label="开启失败锁定">
              <el-switch v-model="form.lockEnable" />
            </el-form-item>
            <el-form-item label="失败次数阈值">
              <el-input-number v-model="form.lockThreshold" :min="1" />
            </el-form-item>
            <el-form-item label="锁定时长(分钟)">
              <el-input-number v-model="form.lockDuration" :min="1" />
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <el-tab-pane label="密码过期" name="expire">
          <el-form :model="form" label-width="160px" class="max-w-2xl">
            <el-form-item label="开启密码过期">
              <el-switch v-model="form.pwdExpireEnable" />
            </el-form-item>
            <el-form-item label="有效天数">
              <el-input-number v-model="form.pwdExpireDays" :min="1" />
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getSecurityConfig, setSecurityConfig } from '@/api/securityConfig'

defineOptions({ name: 'SecurityConfig' })

const activeTab = ref('captcha')
const saving = ref(false)
const form = ref({
  captchaOpen: 0,
  captchaTimeout: 3600,
  keyLong: 6,
  imgWidth: 240,
  imgHeight: 80,
  pwdMinLength: 8,
  pwdRequireUpper: false,
  pwdRequireLower: false,
  pwdRequireDigit: false,
  pwdRequireSpecial: false,
  limitEnable: false,
  limitWindow: 60,
  limitCount: 30,
  lockEnable: false,
  lockThreshold: 5,
  lockDuration: 30,
  pwdExpireEnable: false,
  pwdExpireDays: 90
})

const load = async () => {
  const res = await getSecurityConfig()
  if (res.code === 0) {
    form.value = { ...form.value, ...res.data }
  }
}

const onSave = async () => {
  saving.value = true
  try {
    const res = await setSecurityConfig(form.value)
    if (res.code === 0) {
      form.value = { ...form.value, ...res.data }
      ElMessage.success('保存成功')
    }
  } finally {
    saving.value = false
  }
}

onMounted(load)
</script>
