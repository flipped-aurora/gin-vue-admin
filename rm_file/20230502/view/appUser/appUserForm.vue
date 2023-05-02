<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户名:" prop="userName">
          <el-input v-model="formData.userName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="角色id:" prop="roleID">
          <el-input v-model.number="formData.roleID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公司id:" prop="companyID">
          <el-input v-model.number="formData.companyID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="工资:" prop="wages">
          <el-input-number v-model="formData.wages" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="工作计算方式:" prop="workType">
          <el-select v-model="formData.workType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in workTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AppUser'
}
</script>

<script setup>
import {
  createAppUser,
  updateAppUser,
  findAppUser
} from '@/api/appUser'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const workTypeOptions = ref([])
const statusOptions = ref([])
const formData = ref({
            userName: '',
            password: '',
            roleID: 0,
            nickname: '',
            companyID: 0,
            wages: 0,
            workType: undefined,
            status: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAppUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappUser
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    workTypeOptions.value = await getDictFunc('workType')
    statusOptions.value = await getDictFunc('status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppUser(formData.value)
               break
             case 'update':
               res = await updateAppUser(formData.value)
               break
             default:
               res = await createAppUser(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
