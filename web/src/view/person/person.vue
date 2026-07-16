<template>
  <div class="profile-container">

    <!-- 页面标题 -->
    <div class="profile-card p-4 mb-2">
      <h1 class="text-lg font-bold text-gray-900 dark:text-gray-100">个人信息</h1>
      <p class="mt-1 text-sm text-gray-400 dark:text-gray-500">
        管理你的帐号资料，安全设置与联系方式
      </p>
    </div>

    <!-- 主体两栏布局 -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-2 items-stretch">
      <!-- 左侧：个人资料卡片 -->
      <div class="lg:col-span-4 profile-card p-4">
        <div class="flex flex-col items-center">
          <div class="profile-avatar-wrapper">
            <SelectImage
              v-model="userStore.userInfo.headerImg"
              file-type="image"
              rounded
            />
          </div>

          <!-- 昵称（保留原有编辑功能） -->
          <div
            v-if="!editFlag"
            class="mt-5 text-lg font-bold flex items-center gap-2 text-gray-800 dark:text-gray-100"
          >
            {{ userStore.userInfo.nickName }}
            <el-icon
              class="cursor-pointer text-base text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 transition-colors duration-200"
              @click="openEdit"
            >
              <edit />
            </el-icon>
          </div>
          <div v-else class="mt-5 flex items-center">
            <el-input v-model="nickName" class="w-40 mr-3" />
            <el-button type="primary" plain size="small" @click="enterEdit">
              确认
            </el-button>
            <el-button type="danger" plain size="small" @click="closeEdit">
              取消
            </el-button>
          </div>

          <!-- 角色标签 -->
          <span class="role-tag">
            {{ userStore.userInfo.authority?.authorityName || '系统管理员' }}
          </span>
        </div>

        <div class="card-divider"></div>

        <!-- 基本信息列表 -->
        <div class="space-y-6">
          <div class="info-item">
            <el-icon class="info-icon"><office-building /></el-icon>
            <div>
              <div class="info-label">所属架构</div>
              <div class="info-value">
                {{ orgText }}
              </div>
            </div>
          </div>
          <div class="info-item">
            <el-icon class="info-icon"><phone /></el-icon>
            <div>
              <div class="info-label">手机号码</div>
              <div class="info-value">
                {{ userStore.userInfo.phone || '未设置' }}
              </div>
            </div>
          </div>
          <div class="info-item">
            <el-icon class="info-icon"><message /></el-icon>
            <div>
              <div class="info-label">邮箱地址</div>
              <div class="info-value">
                {{ userStore.userInfo.email || '未设置' }}
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 右侧：账号资料设置 -->
      <div class="lg:col-span-8 profile-card">
        <div class="px-4 pt-4 pb-4">
          <h2 class="text-base font-bold text-gray-800 dark:text-gray-100">
            账号资料设置
          </h2>
          <p class="mt-1.5 text-sm text-gray-400 dark:text-gray-500">
            统一管理头像、联系方式、密码与账号标识
          </p>
        </div>
        <div class="section-divider"></div>

        <div class="px-4">
          <!-- 头像 -->
          <div class="setting-row">
            <div class="setting-label">
              <div class="setting-title">头像</div>
              <div class="setting-desc">用于系统内个人资料展示</div>
            </div>
            <div class="setting-value">
              <img
                v-if="userStore.userInfo.headerImg"
                :src="userStore.userInfo.headerImg"
                class="w-11 h-11 rounded-full object-cover"
                alt="avatar"
              />
            </div>
            <!-- TODO: 原型图中的「更换头像」按钮，当前更换头像入口为点击左侧大头像 -->
            <el-button class="setting-action">更换头像</el-button>
          </div>

          <!-- 手机号码 -->
          <div class="setting-row">
            <div class="setting-label">
              <div class="setting-title">手机号码</div>
              <div class="setting-desc">用于安全验证与通知接收</div>
            </div>
            <div class="setting-value">
              {{ userStore.userInfo.phone || '未设置' }}
            </div>
            <el-button class="setting-action" @click="changePhoneFlag = true">
              修改手机
            </el-button>
          </div>

          <!-- 邮箱地址 -->
          <div class="setting-row">
            <div class="setting-label">
              <div class="setting-title">邮箱地址</div>
              <div class="setting-desc">用于账号通知与登录验证</div>
            </div>
            <div class="setting-value">
              {{ userStore.userInfo.email || '未设置' }}
            </div>
            <el-button class="setting-action" @click="changeEmailFlag = true">
              修改邮箱
            </el-button>
          </div>

          <!-- 账号密码 -->
          <div class="setting-row last-row">
            <div class="setting-label">
              <div class="setting-title">账号密码</div>
              <div class="setting-desc">建议定期更新密码以保障账号安全</div>
            </div>
            <div class="setting-value">已设置</div>
            <el-button class="setting-action" @click="showPassword = true">
              修改密码
            </el-button>
          </div>
        </div>
      </div>
    </div>

    <!-- 弹窗 -->
    <el-dialog
      v-model="showPassword"
      title="修改密码"
      width="400px"
      class="custom-dialog"
      @close="clearPassword"
    >
      <el-form
        ref="modifyPwdForm"
        :model="pwdModify"
        :rules="rules"
        label-width="90px"
        class="py-4"
      >
        <el-form-item :minlength="6" label="原密码" prop="password">
          <el-input v-model="pwdModify.password" show-password />
        </el-form-item>
        <el-form-item :minlength="6" label="新密码" prop="newPassword">
          <el-input v-model="pwdModify.newPassword" show-password />
        </el-form-item>
        <el-form-item :minlength="6" label="确认密码" prop="confirmPassword">
          <el-input v-model="pwdModify.confirmPassword" show-password />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showPassword = false">取 消</el-button>
          <el-button type="primary" @click="savePassword">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="changePhoneFlag"
      title="修改手机号"
      width="400px"
      class="custom-dialog"
    >
      <el-form :model="phoneForm" label-width="80px" class="py-4">
        <el-form-item label="手机号">
          <el-input v-model="phoneForm.phone" placeholder="请输入新的手机号码">
            <template #prefix>
              <el-icon><phone /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="验证码">
          <div class="flex gap-4">
            <el-input
              v-model="phoneForm.code"
              placeholder="请输入验证码[模拟]"
              class="flex-1"
            >
              <template #prefix>
                <el-icon><key /></el-icon>
              </template>
            </el-input>
            <el-button
              type="primary"
              :disabled="time > 0"
              class="w-32"
              @click="getCode"
            >
              {{ time > 0 ? `${time}s` : '获取验证码' }}
            </el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeChangePhone">取 消</el-button>
          <el-button type="primary" @click="changePhone">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="changeEmailFlag"
      title="修改邮箱"
      width="400px"
      class="custom-dialog"
    >
      <el-form :model="emailForm" label-width="80px" class="py-4">
        <el-form-item label="邮箱">
          <el-input v-model="emailForm.email" placeholder="请输入新的邮箱地址">
            <template #prefix>
              <el-icon><message /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="验证码">
          <div class="flex gap-4">
            <el-input
              v-model="emailForm.code"
              placeholder="请输入验证码[模拟]"
              class="flex-1"
            >
              <template #prefix>
                <el-icon><key /></el-icon>
              </template>
            </el-input>
            <el-button
              type="primary"
              :disabled="emailTime > 0"
              class="w-32"
              @click="getEmailCode"
            >
              {{ emailTime > 0 ? `${emailTime}s` : '获取验证码' }}
            </el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeChangeEmail">取 消</el-button>
          <el-button type="primary" @click="changeEmail">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
  import { setSelfInfo, changePassword } from '@/api/user.js'
  import { computed, reactive, ref, watch } from 'vue'
  import { ElMessage } from 'element-plus'
  import { useUserStore } from '@/pinia/modules/user'
  import SelectImage from '@/components/selectImage/selectImage.vue'
  defineOptions({
    name: 'Person'
  })

  const userStore = useUserStore()

  // 所属架构:当前用户各部门的「公司/部门」全路径, 多部门用「｜」分隔, 无部门显示「未定义部门」
  const orgText = computed(() => {
    const list = userStore.userInfo.departments || []
    const text = list
      .map((d) => d.namePath || d.name)
      .filter(Boolean)
      .join(' ｜ ')
    return text || '未定义部门'
  })

  const modifyPwdForm = ref(null)
  const showPassword = ref(false)
  const pwdModify = ref({})
  const nickName = ref('')
  const editFlag = ref(false)

  const rules = reactive({
    password: [
      { required: true, message: '请输入密码', trigger: 'blur' },
      { min: 6, message: '最少6个字符', trigger: 'blur' }
    ],
    newPassword: [
      { required: true, message: '请输入新密码', trigger: 'blur' },
      { min: 6, message: '最少6个字符', trigger: 'blur' }
    ],
    confirmPassword: [
      { required: true, message: '请输入确认密码', trigger: 'blur' },
      { min: 6, message: '最少6个字符', trigger: 'blur' },
      {
        validator: (rule, value, callback) => {
          if (value !== pwdModify.value.newPassword) {
            callback(new Error('两次密码不一致'))
          } else {
            callback()
          }
        },
        trigger: 'blur'
      }
    ]
  })

  const savePassword = async () => {
    modifyPwdForm.value.validate((valid) => {
      if (valid) {
        changePassword({
          password: pwdModify.value.password,
          newPassword: pwdModify.value.newPassword
        }).then((res) => {
          if (res.code === 0) {
            ElMessage.success('修改密码成功！')
          }
          showPassword.value = false
        })
      }
    })
  }

  const clearPassword = () => {
    pwdModify.value = {
      password: '',
      newPassword: '',
      confirmPassword: ''
    }
    modifyPwdForm.value?.clearValidate()
  }

  const openEdit = () => {
    nickName.value = userStore.userInfo.nickName
    editFlag.value = true
  }

  const closeEdit = () => {
    nickName.value = ''
    editFlag.value = false
  }

  const enterEdit = async () => {
    const res = await setSelfInfo({
      nickName: nickName.value
    })
    if (res.code === 0) {
      userStore.ResetUserInfo({ nickName: nickName.value })
      ElMessage.success('修改成功')
    }
    nickName.value = ''
    editFlag.value = false
  }

  const changePhoneFlag = ref(false)
  const time = ref(0)
  const phoneForm = reactive({
    phone: '',
    code: ''
  })

  const getCode = async () => {
    time.value = 60
    let timer = setInterval(() => {
      time.value--
      if (time.value <= 0) {
        clearInterval(timer)
        timer = null
      }
    }, 1000)
  }

  const closeChangePhone = () => {
    changePhoneFlag.value = false
    phoneForm.phone = ''
    phoneForm.code = ''
  }

  const changePhone = async () => {
    const res = await setSelfInfo({ phone: phoneForm.phone })
    if (res.code === 0) {
      ElMessage.success('修改成功')
      userStore.ResetUserInfo({ phone: phoneForm.phone })
      closeChangePhone()
    }
  }

  const changeEmailFlag = ref(false)
  const emailTime = ref(0)
  const emailForm = reactive({
    email: '',
    code: ''
  })

  const getEmailCode = async () => {
    emailTime.value = 60
    let timer = setInterval(() => {
      emailTime.value--
      if (emailTime.value <= 0) {
        clearInterval(timer)
        timer = null
      }
    }, 1000)
  }

  const closeChangeEmail = () => {
    changeEmailFlag.value = false
    emailForm.email = ''
    emailForm.code = ''
  }

  const changeEmail = async () => {
    const res = await setSelfInfo({ email: emailForm.email })
    if (res.code === 0) {
      ElMessage.success('修改成功')
      userStore.ResetUserInfo({ email: emailForm.email })
      closeChangeEmail()
    }
  }

  watch(() => userStore.userInfo.headerImg, async(val) => {
    const res = await setSelfInfo({ headerImg: val })
    if (res.code === 0) {
      userStore.ResetUserInfo({ headerImg: val })
      ElMessage({
        type: 'success',
        message: '设置成功',
      })
    }
  })
</script>

<style scoped lang="scss">
  .profile-container {
    @apply py-2 min-h-screen bg-gray-50 dark:bg-slate-900;

    /* ===== 页面头部 ===== */
    .page-header {
      @apply flex items-start justify-between mb-6 lg:mb-8;
    }

    .page-title {
      @apply text-lg font-bold text-gray-900 dark:text-gray-100;
    }

    .page-subtitle {
      @apply mt-1 text-sm text-gray-400 dark:text-gray-500;
    }

    .save-btn {
      @apply px-6 rounded-lg;
    }

    /* ===== 卡片通用 ===== */
    .profile-card {
      @apply bg-white dark:bg-slate-800 border border-gray-100 dark:border-slate-700 shadow-sm;
      border-radius: var(--gva-radius, 0.75rem);
    }

    /* ===== 左侧卡片 ===== */
    .profile-avatar-wrapper {
      @apply flex justify-center;
    }

    .role-tag {
      @apply mt-3 inline-block px-3 py-1 rounded-md text-sm
        text-blue-600 bg-blue-50 dark:text-blue-300 dark:bg-blue-500/10;
    }

    .card-divider {
      @apply my-6 border-t border-gray-100 dark:border-slate-700;
    }

    .info-item {
      @apply flex items-start gap-3;
    }

    .info-icon {
      @apply mt-0.5 text-lg text-gray-400 dark:text-gray-500;
    }

    .info-label {
      @apply text-sm text-gray-400 dark:text-gray-500;
    }

    .info-value {
      @apply mt-1 text-sm text-gray-800 dark:text-gray-200;
    }

    /* ===== 右侧卡片 ===== */
    .section-divider {
      @apply border-t border-gray-100 dark:border-slate-700;
    }

    .setting-row {
      @apply flex flex-col lg:flex-row lg:items-center gap-3 lg:gap-4 py-4
        border-b border-gray-100 dark:border-slate-700;

      &.last-row {
        @apply border-b-0;
      }
    }

    .setting-label {
      @apply w-full lg:w-2/5 flex-shrink-0;
    }

    .setting-title {
      @apply text-sm font-medium text-gray-800 dark:text-gray-200;
    }

    .setting-desc {
      @apply mt-1.5 text-sm text-gray-400 dark:text-gray-500;
    }

    .setting-value {
      @apply flex-1 text-sm text-gray-800 dark:text-gray-200;
    }

    .setting-action {
      @apply flex-shrink-0 rounded-lg w-24;
    }

    /* ===== 弹窗（原样保留） ===== */
    .custom-dialog {
      :deep(.el-dialog__header) {
        @apply mb-0 pb-4 border-b border-gray-100 dark:border-gray-700;
      }
      :deep(.el-dialog__footer) {
        @apply mt-0 pt-4 border-t border-gray-100 dark:border-gray-700;
      }
      :deep(.el-input__prefix) {
        @apply mr-2;
      }
    }
  }
</style>