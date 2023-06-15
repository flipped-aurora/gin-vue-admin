<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="uuid字段:" prop="uuid">
        </el-form-item>
        <el-form-item label="username字段:" prop="username">
        </el-form-item>
        <el-form-item label="password字段:" prop="password">
        </el-form-item>
        <el-form-item label="nickName字段:" prop="nickName">
        </el-form-item>
        <el-form-item label="sideMode字段:" prop="sideMode">
        </el-form-item>
        <el-form-item label="headerImg字段:" prop="headerImg">
        </el-form-item>
        <el-form-item label="baseColor字段:" prop="baseColor">
        </el-form-item>
        <el-form-item label="activeColor字段:" prop="activeColor">
        </el-form-item>
        <el-form-item label="authorityId字段:" prop="authorityId">
        </el-form-item>
        <el-form-item label="phone字段:" prop="phone">
        </el-form-item>
        <el-form-item label="email字段:" prop="email">
        </el-form-item>
        <el-form-item label="enable字段:" prop="enable">
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
  name: 'SysUsers'
}
</script>

<script setup>
import {
  createSysUsers,
  updateSysUsers,
  findSysUsers
} from '@/api/sysUsers'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSysUsers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resysUsers
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
               res = await createSysUsers(formData.value)
               break
             case 'update':
               res = await updateSysUsers(formData.value)
               break
             default:
               res = await createSysUsers(formData.value)
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
