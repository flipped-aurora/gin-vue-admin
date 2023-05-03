<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="裁剪单id:" prop="croppingID">
          <el-input v-model.number="formData.croppingID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户id:" prop="userID">
          <el-input v-model.number="formData.userID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="工序名称:" prop="processName">
          <el-input v-model="formData.processName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="单价:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="申请数量:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="预估收入:" prop="income">
          <el-input-number v-model="formData.income" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="真实数量:" prop="realQuantity">
          <el-input v-model.number="formData.realQuantity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="真实收入:" prop="realIncome">
          <el-input-number v-model="formData.realIncome" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="步骤:" prop="step">
          <el-select v-model="formData.step" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in stepOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="工作类型:" prop="jobType">
          <el-select v-model="formData.jobType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in jobTypeOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'Job'
}
</script>

<script setup>
import {
  createJob,
  updateJob,
  findJob
} from '@/api/job'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const stepOptions = ref([])
const jobTypeOptions = ref([])
const formData = ref({
            croppingID: 0,
            userID: 0,
            processName: '',
            price: 0,
            quantity: 0,
            income: 0,
            realQuantity: 0,
            realIncome: 0,
            step: undefined,
            jobType: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findJob({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rejob
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    stepOptions.value = await getDictFunc('step')
    jobTypeOptions.value = await getDictFunc('jobType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createJob(formData.value)
               break
             case 'update':
               res = await updateJob(formData.value)
               break
             default:
               res = await createJob(formData.value)
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
