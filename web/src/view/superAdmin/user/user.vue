<template>
  <div>
    <warning-bar :title="t('authority.authorityNote')" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="addUser">{{ t('user.addUser') }}</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column align="left" :label="t('user.avatar')" min-width="50">
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="UUID" min-width="250" prop="uuid" />
        <el-table-column align="left" :label="t('user.userName')" min-width="150" prop="userName" />
        <el-table-column align="left" :label="t('user.nickName')" min-width="100" prop="nickName">
          <template #default="scope">
            <p v-if="!scope.row.editFlag" class="nickName">{{ scope.row.nickName }}
              <el-icon class="pointer" color="#66b1ff" @click="openEidt(scope.row)">
                <edit />
              </el-icon>
            </p>
            <p v-if="scope.row.editFlag" class="nickName">
              <el-input v-model="scope.row.nickName" />
              <el-icon class="pointer" color="#67c23a" @click="enterEdit(scope.row)">
                <check />
              </el-icon>
              <el-icon class="pointer" color="#f23c3c" @click="closeEdit(scope.row)">
                <close />
              </el-icon>
            </p>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('user.userRole')" min-width="150">
          <template #default="scope">
            <el-cascader
              v-model="scope.row.authorityIds"
              :options="authOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
              @visible-change="(flag)=>{changeAuthority(scope.row,flag)}"
              @remove-tag="()=>{changeAuthority(scope.row,false)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('general.operations')" min-width="150">
          <template #default="scope">
            <el-popover v-model:visible="scope.row.visible" placement="top" width="160">
              <p>{{ t('user.deleteUserConfrim') }}</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="scope.row.visible = false">>{{ t('general.cancel') }}</el-button>
                <el-button type="primary" size="mini" @click="deleteUserFunc(scope.row)">{{ t('general.confirm') }}</el-button>
              </div>
              <template #reference>
                <el-button type="text" icon="delete" size="mini">{{ t('general.delete') }}</el-button>
              </template>
            </el-popover>
            <el-button type="text" icon="magic-stick" size="mini" @click="resetPasswordFunc(scope.row)">{{ t('user.resetPassword') }}</el-button>
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
    <el-dialog v-model="addUserDialog" custom-class="user-dialog" :title="t('user.addUser')">
      <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="90px">
        <el-form-item :label="t('user.userName')" prop="username">
          <el-input v-model="userInfo.username" />
        </el-form-item>
        <el-form-item :label="t('user.password')" prop="password">
          <el-input v-model="userInfo.password" />
        </el-form-item>
        <el-form-item :label="t('user.nickName')" prop="nickName">
          <el-input v-model="userInfo.nickName" />
        </el-form-item>
        <el-form-item :label="t('user.userRole')" prop="authorityId">
          <el-cascader
            v-model="userInfo.authorityIds"
            style="width:100%"
            :options="authOptions"
            :show-all-levels="false"
            :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
            :clearable="false"
          />
        </el-form-item>
        <el-form-item :label="t('user.avatar')" label-width="80px">
          <div style="display:inline-block" @click="openHeaderChange">
            <img v-if="userInfo.headerImg" class="header-img-box" :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg">
            <div v-else class="header-img-box">{{ t('user.mediaLibrary') }}</div>
          </div>
        </el-form-item>

      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeAddUserDialog">{{ t('general.close') }}</el-button>
          <el-button size="small" type="primary" @click="enterAddUserDialog">{{ t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`" />
  </div>
</template>

<script>
export default {
  name: 'User',
}
</script>

<script setup>
import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser
} from '@/api/user'

import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import warningBar from '@/components/warningBar/warningBar.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const path = ref(import.meta.env.VITE_BASE_API)
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
        AuthorityData.forEach(item => {
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
const getTableData = async() => {
  const table = await getUserList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

watch(tableData, () => {
  setAuthorityIds()
})

const initPage = async() => {
  getTableData()
  const res = await getAuthorityList({ page: 1, pageSize: 999 })
  setOptions(res.data.list)
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm(
    t('user.resetPasswordConfrim'),
    t('general.warning'),
    {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning',
    }
  ).then(async() => {
    const res = await resetPassword({
      ID: row.ID,
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg,
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg,
      })
    }
  })
}
const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    const authorityIds = user.authorities && user.authorities.map(i => {
      return i.authorityId
    })
    user.authorityIds = authorityIds
  })
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  setAuthorityOptions(authData, authOptions.value)
}

const backNickName = ref('')
const openEidt = (row) => {
  if (tableData.value.some(item => item.editFlag)) {
    ElMessage(t('user.anotherUserEdit'))
    return
  }
  backNickName.value = row.nickName
  row.editFlag = true
}

const enterEdit = async(row) => {
  const res = await setUserInfo({ nickName: row.nickName, ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: t('general.setupSuccess')
    })
  }
  backNickName.value = ref('')
  row.editFlag = false
}

const closeEdit = (row) => {
  row.nickName = backNickName.value
  backNickName.value = ''
  row.editFlag = false
}

const deleteUserFunc = async(row) => {
  const res = await deleteUser({ id: row.ID })
  if (res.code === 0) {
    ElMessage.success(t('general.deleteSuccess'))
    await getTableData()
    row.visible = false
  }
}

// 弹窗相关
const userInfo = ref({
  username: '',
  password: '',
  nickName: '',
  headerImg: '',
  authorityId: '',
  authorityIds: []
})

const rules = ref({
  username: [
    { required: true, message: t('user.userNameNote'), trigger: 'blur' },
    { min: 5, message: t('user.userNameLenNote'), trigger: 'blur' }
  ],
  password: [
    { required: true, message: t('user.passwordNote'), trigger: 'blur' },
    { min: 6, message: t('user.passwordLenNote'), trigger: 'blur' }
  ],
  nickName: [
    { required: true, message: t('user.nickNameNote'), trigger: 'blur' }
  ],
  authorityId: [
    { required: true, message: t('user.userRoleNote'), trigger: 'blur' }
  ]
})
const userForm = ref(null)
const enterAddUserDialog = async() => {
  userInfo.value.authorityId = userInfo.value.authorityIds[0]
  userForm.value.validate(async valid => {
    if (valid) {
      const res = await register(userInfo.value)
      if (res.code === 0) {
        ElMessage({ type: 'success', message: t('user.userAddedNote') })
      }
      await getTableData()
      closeAddUserDialog()
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
const addUser = () => {
  addUserDialog.value = true
}
const changeAuthority = async(row, flag) => {
  if (flag) {
    return
  }

  await nextTick()
  const res = await setUserAuthorities({
    ID: row.ID,
    authorityIds: row.authorityIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: t('user.roleSetNote') })
  }
}
</script>

<style lang="scss">
.user-dialog {
  .header-img-box {
  width: 200px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
.nickName{
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.pointer{
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}
</style>
