<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="布料宽:" prop="clothWidth">
          <el-input-number v-model="formData.clothWidth" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="斜条长度:" prop="barLength">
          <el-input-number v-model="formData.barLength" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="斜条宽:" prop="barWidth">
          <el-input-number v-model="formData.barWidth" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="件数:" prop="countNum">
          <el-input v-model.number="formData.countNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="单根协调长度45:" prop="barLength45">
          <el-input-number v-model="formData.barLength45" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="90:" prop="barLength90">
          <el-input-number v-model="formData.barLength90" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="180:" prop="barLength180">
          <el-input-number v-model="formData.barLength180" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="所需总布料长度45:" prop="clothLength45">
          <el-input-number v-model="formData.clothLength45" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="90:" prop="clothLength90">
          <el-input-number v-model="formData.clothLength90" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="180:" prop="clothLength180">
          <el-input-number v-model="formData.clothLength180" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="用户id:" prop="userID">
          <el-input v-model.number="formData.userID" :clearable="true" placeholder="请输入" />
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
  name: 'Computation'
}
</script>

<script setup>
import {
  createComputation,
  updateComputation,
  findComputation
} from '@/api/computation'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            clothWidth: 0,
            barLength: 0,
            barWidth: 0,
            countNum: 0,
            barLength45: 0,
            barLength90: 0,
            barLength180: 0,
            clothLength45: 0,
            clothLength90: 0,
            clothLength180: 0,
            userID: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findComputation({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recomputation
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
               res = await createComputation(formData.value)
               break
             case 'update':
               res = await updateComputation(formData.value)
               break
             default:
               res = await createComputation(formData.value)
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
