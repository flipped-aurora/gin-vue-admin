<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="招聘职位:" prop="job">
          <el-input v-model="formData.job" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="招聘人数:" prop="count">
          <el-input v-model="formData.count" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="工作地点:" prop="address">
          <el-input v-model="formData.address" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="岗位描述:" prop="content">
          <el-input v-model="formData.content" :clearable="true" placeholder="请输入" />
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
  name: 'Recruit'
}
</script>

<script setup>
import {
  createRecruit,
  updateRecruit,
  findRecruit
} from '@/api/recruit'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            job: '',
            count: '',
            address: '',
            content: '',
        })
// 验证规则
const rule = reactive({
               job : [{
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
      const res = await findRecruit({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rerecruit
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
               res = await createRecruit(formData.value)
               break
             case 'update':
               res = await updateRecruit(formData.value)
               break
             default:
               res = await createRecruit(formData.value)
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
