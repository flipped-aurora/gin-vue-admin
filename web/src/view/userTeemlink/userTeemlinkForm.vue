<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户ID:" prop="userID">
          <el-input v-model.number="formData.userID" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="天翎用户ID:" prop="tlUserID">
          <el-input v-model="formData.tlUserID" :clearable="true" placeholder="请输入" />
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
  name: 'UserTeemlink'
}
</script>

<script setup>
import {
  createUserTeemlink,
  updateUserTeemlink,
  findUserTeemlink
} from '@/api/userTeemlink'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            userID: 0,
            username: '',
            tlUserID: '',
        })
// 验证规则
const rule = reactive({
               userID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               tlUserID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserTeemlink({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reusertl
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createUserTeemlink(formData.value)
               break
             case 'update':
               res = await updateUserTeemlink(formData.value)
               break
             default:
               res = await createUserTeemlink(formData.value)
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
