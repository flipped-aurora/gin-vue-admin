<template>
  <div>
    <warning-bar :title="t('authority.authorityNote')" />
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item :label="t('view.superAdmin.user.userName')">
          <el-input
            v-model="searchInfo.username"
            :placeholder="t('view.superAdmin.user.addUser')"
          />
        </el-form-item>
        <el-form-item :label="t('view.superAdmin.user.nickName')">
          <el-input
            v-model="searchInfo.nickname"
            :placeholder="t('view.superAdmin.user.nickName')"
          />
        </el-form-item>
        <el-form-item :label="t('view.superAdmin.user.phone')">
          <el-input
            v-model="searchInfo.phone"
            :placeholder="t('view.superAdmin.user.phone')"
          />
        </el-form-item>
        <el-form-item :label="t('view.superAdmin.user.email')">
          <el-input
            v-model="searchInfo.email"
            :placeholder="t('view.superAdmin.user.email')"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            {{ t('general.search') }}
          </el-button>
          <el-button icon="refresh" @click="onReset">
            {{ t('general.reset') }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addUser">{{
          t('view.superAdmin.user.addUser')
        }}</el-button>
      </div>
      <el-table :data="tableData" row-key="ID">
        <el-table-column
          align="left"
          :label="t('view.superAdmin.user.avatar')"
          min-width="75"
        >
          <template #default="scope">
            <CustomPic style="margin-top: 8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="ID" min-width="50" prop="ID" />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.user.userName')"
          min-width="150"
          prop="userName"
        />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.user.nickName')"
          min-width="150"
          prop="nickName"
        />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.user.phone')"
          min-width="180"
          prop="phone"
        />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.user.email')"
          min-width="180"
          prop="email"
        />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.user.userRole')"
          min-width="200"
        >
          <template #default="scope">
            <el-cascader
              v-model="scope.row.authorityIds"
              :options="authOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{
                multiple: true,
                checkStrictly: true,
                label: 'authorityName',
                value: 'authorityId',
                disabled: 'disabled',
                emitPath: false
              }"
              :clearable="false"
              @visible-change="
                (flag) => {
                  changeAuthority(scope.row, flag, 0)
                }
              "
              @remove-tag="
                (removeAuth) => {
                  changeAuthority(scope.row, false, removeAuth)
                }
              "
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('view.superAdmin.user.enable')"
          min-width="150"
        >
          <template #default="scope">
            <el-switch
              v-model="scope.row.enable"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="
                () => {
                  switchEnable(scope.row)
                }
              "
            />
          </template>
        </el-table-column>

        <el-table-column :label="t('general.operations')" :min-width="appStore.operateMinWith" align="left" fixed="right" width="300px">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteUserFunc(scope.row)"
              >{{ t('general.delete') }}</el-button
            >
            <el-button
              type="primary"
              link
              icon="edit"
              @click="openEdit(scope.row)"
              >{{ t('general.edit') }}</el-button
            >
            <el-button
              type="primary"
              link
              icon="magic-stick"
              @click="resetPasswordFunc(scope.row)"
              >{{ t('view.superAdmin.user.resetPassword') }}</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <!-- 重置密码对话框 -->
    <el-dialog
      v-model="resetPwdDialog"
      title="重置密码"
      width="500px"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <el-form :model="resetPwdInfo" ref="resetPwdForm" label-width="100px">
        <el-form-item label="用户账号">
          <el-input v-model="resetPwdInfo.userName" disabled />
        </el-form-item>
        <el-form-item label="用户昵称">
          <el-input v-model="resetPwdInfo.nickName" disabled />
        </el-form-item>
        <el-form-item label="新密码">
          <div class="flex w-full">
            <el-input class="flex-1" v-model="resetPwdInfo.password" placeholder="请输入新密码" show-password />
            <el-button type="primary" @click="generateRandomPassword" style="margin-left: 10px">
              生成随机密码
            </el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeResetPwdDialog">取 消</el-button>
          <el-button type="primary" @click="confirmResetPassword">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    
    <el-drawer
      v-model="addUserDialog"
      :size="appStore.drawerSize"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ t('view.superAdmin.user.user') }}</span>
          <div>
            <el-button @click="closeAddUserDialog">{{
              t('general.close')
            }}</el-button>
            <el-button type="primary" @click="enterAddUserDialog">{{
              t('general.confirm')
            }}</el-button>
          </div>
        </div>
      </template>

      <el-form
        ref="userForm"
        :rules="rules"
        :model="userInfo"
        label-width="120px"
      >
        <el-form-item
          v-if="dialogFlag === 'add'"
          :label="t('view.superAdmin.user.userName')"
          prop="userName"
        >
          <el-input v-model="userInfo.userName" />
        </el-form-item>
        <el-form-item
          v-if="dialogFlag === 'add'"
          :label="t('view.superAdmin.user.password')"
          prop="password"
        >
          <el-input v-model="userInfo.password" />
        </el-form-item>
        <el-form-item
          :label="t('view.superAdmin.user.nickName')"
          prop="nickName"
        >
          <el-input v-model="userInfo.nickName" />
        </el-form-item>
        <el-form-item :label="t('view.superAdmin.user.phone')" prop="phone">
          <el-input v-model="userInfo.phone" />
        </el-form-item>
        <el-form-item :label="t('view.superAdmin.user.email')" prop="email">
          <el-input v-model="userInfo.email" />
        </el-form-item>
        <el-form-item
          :label="t('view.superAdmin.user.userRole')"
          prop="authorityId"
        >
          <el-cascader
            v-model="userInfo.authorityIds"
            style="width: 100%"
            :options="authOptions"
            :show-all-levels="false"
            :props="{
              multiple: true,
              checkStrictly: true,
              label: 'authorityName',
              value: 'authorityId',
              disabled: 'disabled',
              emitPath: false
            }"
            :clearable="false"
          />
        </el-form-item>
        <el-form-item :label="t('view.superAdmin.user.enable')" prop="disabled">
          <el-switch
            v-model="userInfo.enable"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>
        <el-form-item
          :label="t('view.superAdmin.user.avatar')"
          label-width="80px"
        >
          <SelectImage v-model="userInfo.headerImg" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    getUserList,
    setUserAuthorities,
    register,
    deleteUser
  } from '@/api/user'

  import { getAuthorityList } from '@/api/authority'
  import CustomPic from '@/components/customPic/index.vue'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { setUserInfo, resetPassword } from '@/api/user.js'

  import { nextTick, ref, watch } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import SelectImage from '@/components/selectImage/selectImage.vue'
  import { useAppStore } from "@/pinia";
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual

  const { t } = useI18n() // added by mohamed hassan to support multilingual

  defineOptions({
    name: 'User'
  })

  const appStore = useAppStore()

  const searchInfo = ref({
    username: '',
    nickname: '',
    phone: '',
    email: ''
  })

  const onSubmit = () => {
    page.value = 1
    getTableData()
  }

  const onReset = () => {
    searchInfo.value = {
      username: '',
      nickname: '',
      phone: '',
      email: ''
    }
    getTableData()
  }
  // 初始化相关
  const setAuthorityOptions = (AuthorityData, optionsData) => {
    AuthorityData &&
      AuthorityData.forEach((item) => {
        if (item.children && item.children.length) {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName,
            children: []
          }
          setAuthorityOptions(item.children, option.children)
          optionsData.push(option)
        } else {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName
          }
          optionsData.push(option)
        }
      })
  }

  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 查询
  const getTableData = async () => {
    const table = await getUserList({
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  watch(
    () => tableData.value,
    () => {
      setAuthorityIds()
    }
  )

  const initPage = async () => {
    getTableData()
    const res = await getAuthorityList()
    setOptions(res.data)
  }

  initPage()

  // 重置密码对话框相关
  const resetPwdDialog = ref(false)
  const resetPwdForm = ref(null)
  const resetPwdInfo = ref({
    ID: '',
    userName: '',
    nickName: '',
    password: ''
  })
  
  // 生成随机密码
  const generateRandomPassword = () => {
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*'
    let password = ''
    for (let i = 0; i < 12; i++) {
      password += chars.charAt(Math.floor(Math.random() * chars.length))
    }
    resetPwdInfo.value.password = password
    // 复制到剪贴板
    navigator.clipboard.writeText(password).then(() => {
      ElMessage({
        type: 'success',
        message: '密码已复制到剪贴板'
      })
    }).catch(() => {
      ElMessage({
        type: 'error',
        message: '复制失败，请手动复制'
      })
    })
  }
  
  // 打开重置密码对话框
  const resetPasswordFunc = (row) => {
    resetPwdInfo.value.ID = row.ID
    resetPwdInfo.value.userName = row.userName
    resetPwdInfo.value.nickName = row.nickName
    resetPwdInfo.value.password = ''
    resetPwdDialog.value = true
  }
  
  // 确认重置密码
  const confirmResetPassword = async () => {
    if (!resetPwdInfo.value.password) {
      ElMessage({
        type: 'warning',
        message: '请输入或生成密码'
      })
      return
    }
    
    const res = await resetPassword({
      ID: resetPwdInfo.value.ID,
      password: resetPwdInfo.value.password
    })
    
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg || '密码重置成功'
      })
      resetPwdDialog.value = false
    } else {
      ElMessage({
        type: 'error',
        message: res.msg || '密码重置失败'
      })
    }
  }
  
  // 关闭重置密码对话框
  const closeResetPwdDialog = () => {
    resetPwdInfo.value.password = ''
    resetPwdDialog.value = false
  }
  const setAuthorityIds = () => {
    tableData.value &&
      tableData.value.forEach((user) => {
        user.authorityIds =
          user.authorities &&
          user.authorities.map((i) => {
            return i.authorityId
          })
      })
  }

  const authOptions = ref([])
  const setOptions = (authData) => {
    authOptions.value = []
    setAuthorityOptions(authData, authOptions.value)
  }

  const deleteUserFunc = async (row) => {
    ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning'
    }).then(async () => {
      const res = await deleteUser({ id: row.ID })
      if (res.code === 0) {
        ElMessage.success(t('general.deleteSuccess'))
        await getTableData()
      }
    })
  }

  // 弹窗相关
  const userInfo = ref({
    userName: '',
    password: '',
    nickName: '',
    headerImg: '',
    authorityId: '',
    authorityIds: [],
    enable: 1
  })

  const rules = ref({
    userName: [
      {
        required: true,
        message: t('view.superAdmin.user.userNameNote'),
        trigger: 'blur'
      },
      {
        min: 5,
        message: t('view.superAdmin.user.userNameLenNote'),
        trigger: 'blur'
      }
    ],
    password: [
      {
        required: true,
        message: t('view.superAdmin.user.passwordNote'),
        trigger: 'blur'
      },
      {
        min: 6,
        message: t('view.superAdmin.user.passwordLenNote'),
        trigger: 'blur'
      }
    ],
    nickName: [
      {
        required: true,
        message: t('view.superAdmin.user.nickNameNote'),
        trigger: 'blur'
      }
    ],
    phone: [
      {
        pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/,
        message: t('view.superAdmin.user.enterPhoneNoNote'),
        trigger: 'blur'
      }
    ],
    email: [
      {
        pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g,
        message: t('view.superAdmin.user.enterEmailNote'),
        trigger: 'blur'
      }
    ],
    authorityId: [
      {
        required: true,
        message: t('view.superAdmin.user.userRoleNote'),
        trigger: 'blur'
      }
    ]
  })
  const userForm = ref(null)
  const enterAddUserDialog = async () => {
    userInfo.value.authorityId = userInfo.value.authorityIds[0]
    userForm.value.validate(async (valid) => {
      if (valid) {
        const req = {
          ...userInfo.value
        }
        if (dialogFlag.value === 'add') {
          const res = await register(req)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: t('view.superAdmin.user.userAddedNote')
            })
            await getTableData()
            closeAddUserDialog()
          }
        }
        if (dialogFlag.value === 'edit') {
          const res = await setUserInfo(req)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: t('view.superAdmin.user.userEditedNote')
            })
            await getTableData()
            closeAddUserDialog()
          }
        }
      }
    })
  }

  const addUserDialog = ref(false)
  const closeAddUserDialog = () => {
    userForm.value.resetFields()
    userInfo.value.headerImg = ''
    userInfo.value.authorityIds = []
    addUserDialog.value = false
  }

  const dialogFlag = ref('add')

  const addUser = () => {
    addUserDialog.value = true
    dialogFlag.value = 'add'
  }

  const tempAuth = {}
  const changeAuthority = async (row, flag, removeAuth) => {
    if (flag) {
      if (!removeAuth) {
        tempAuth[row.ID] = [...row.authorityIds]
      }
      return
    }
    await nextTick()
    const res = await setUserAuthorities({
      ID: row.ID,
      authorityIds: row.authorityIds
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: t('view.superAdmin.user.roleSetNote')
      })
    } else {
      if (!removeAuth) {
        row.authorityIds = [...tempAuth[row.ID]]
        delete tempAuth[row.ID]
      } else {
        row.authorityIds = [removeAuth, ...row.authorityIds]
      }
    }
  }

  const openEdit = (row) => {
    dialogFlag.value = 'edit'
    userInfo.value = JSON.parse(JSON.stringify(row))
    addUserDialog.value = true
  }

  const switchEnable = async (row) => {
    userInfo.value = JSON.parse(JSON.stringify(row))
    await nextTick()
    const req = {
      ...userInfo.value
    }
    const res = await setUserInfo(req)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: `${req.enable === 2 ? t('view.superAdmin.user.disabledSuccessfully') : t('view.superAdmin.user.enabledSuccessfully')}`
      })
      await getTableData()
      userInfo.value.headerImg = ''
      userInfo.value.authorityIds = []
    }
  }
</script>

<style lang="scss">
  .header-img-box {
    @apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
  }
</style>
