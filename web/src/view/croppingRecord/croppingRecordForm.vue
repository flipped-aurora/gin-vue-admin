<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="公司id:" prop="companyID">
          <el-input v-model.number="formData.companyID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="款式id:" prop="styleID">
          <el-input v-model.number="formData.styleID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="颜色:" prop="color">
          <el-input v-model="formData.color" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="长度:" prop="length">
          <el-input-number v-model="formData.length" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="尺码:" prop="size">
          <el-input v-model="formData.size" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数量:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="步骤:" prop="step">
          <el-select v-model="formData.step" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in stepOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'CroppingRecord'
}
</script>

<script setup>
import {
  createCroppingRecord,
  updateCroppingRecord,
  findCroppingRecord
} from '@/api/croppingRecord'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const stepOptions = ref([])
const formData = ref({
            companyID: 0,
            styleID: 0,
            color: '',
            length: 0,
            size: '',
            quantity: 0,
            step: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCroppingRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recroppingRecord
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    stepOptions.value = await getDictFunc('step')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCroppingRecord(formData.value)
               break
             case 'update':
               res = await updateCroppingRecord(formData.value)
               break
             default:
               res = await createCroppingRecord(formData.value)
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
