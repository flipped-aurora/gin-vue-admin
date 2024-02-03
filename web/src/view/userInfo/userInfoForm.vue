<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户openid:" prop="userWxopenid">
          <el-input v-model="formData.userWxopenid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户昵称:" prop="userNickname">
          <el-input v-model="formData.userNickname" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户头像:" prop="userAvatarUrl">
          <SelectImage v-model="formData.userAvatarUrl" file-type="image"/>
       </el-form-item>
        <el-form-item label="用户性别:" prop="userGender">
          <el-select v-model="formData.userGender" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="用户年级:" prop="userGrade">
          <el-input v-model.number="formData.userGrade" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户专业:" prop="userProfession">
          <el-input v-model="formData.userProfession" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户简介:" prop="userIntroduction">
          <el-input v-model="formData.userIntroduction" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户标签:" prop="userLabel">
          <el-input v-model="formData.userLabel" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户城市:" prop="userCity">
          <el-input v-model="formData.userCity" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户获赞数量:" prop="loveNumber">
          <el-input v-model.number="formData.loveNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户身份:" prop="userModel">
          <el-select v-model="formData.userModel" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in userModelOptions" :key="key" :label="item.label" :value="item.value" />
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

<script setup>
import {
  createUserInfo,
  updateUserInfo,
  findUserInfo
} from '@/api/userInfo'

defineOptions({
    name: 'UserInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const genderOptions = ref([])
const userModelOptions = ref([])
const formData = ref({
            userWxopenid: '',
            userNickname: '',
            userAvatarUrl: "",
            userGender: undefined,
            userGrade: 0,
            userProfession: '',
            userIntroduction: '',
            userLabel: '',
            userCity: '',
            loveNumber: 0,
            userModel: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reuserData
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    genderOptions.value = await getDictFunc('gender')
    userModelOptions.value = await getDictFunc('userModel')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createUserInfo(formData.value)
               break
             case 'update':
               res = await updateUserInfo(formData.value)
               break
             default:
               res = await createUserInfo(formData.value)
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
