<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="资料名称:">
          <el-input v-model="formData.profile_name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="formData.username" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="年龄:">
          <el-input v-model.number="formData.age" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:">
          <el-select v-model="formData.sex" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="工作经验:">
          <el-input v-model.number="formData.work_experience" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="手机号:">
          <el-input v-model="formData.phone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="简历文件:">
          <el-input v-model="formData.resume_file" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="隐身:">
          <el-switch v-model="formData.stealth" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="公司名称:">
          <el-input v-model="formData.company_name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="简介:">
          <el-input v-model="formData.intro" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图标:">
          <el-input v-model="formData.logo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮箱:">
          <el-input v-model="formData.email" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="网站:">
          <el-input v-model="formData.website" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户ID:">
          <el-input v-model.number="formData.user_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户名:">
          <el-input v-model="formData.user_name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Profile'
}
</script>

<script setup>
import {
  createProfile,
  updateProfile,
  findProfile
} from '@/api/profile'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const genderOptions = ref([])
const formData = ref({
  profile_name: '',
  username: '',
  age: 0,
  sex: undefined,
  work_experience: 0,
  phone: '',
  resume_file: '',
  stealth: false,
  company_name: '',
  intro: '',
  logo: '',
  email: '',
  website: '',
  user_id: 0,
  user_name: '',
})

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findProfile({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reprofile
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
  genderOptions.value = await getDictFunc('gender')
}

init()
// 保存按钮
const save = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createProfile(formData.value)
      break
    case 'update':
      res = await updateProfile(formData.value)
      break
    default:
      res = await createProfile(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
  }
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
