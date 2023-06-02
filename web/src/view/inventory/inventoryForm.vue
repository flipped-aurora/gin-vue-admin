<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="裁剪单id:" prop="croppingID">
          <el-input v-model.number="formData.croppingID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="尺码表id:" prop="sizeID">
          <el-input v-model.number="formData.sizeID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="工序id:" prop="processID">
          <el-input v-model.number="formData.processID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="总量:" prop="total">
          <el-input v-model.number="formData.total" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="余量:" prop="margin">
          <el-input v-model.number="formData.margin" :clearable="true" placeholder="请输入" />
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
  name: 'Inventory'
}
</script>

<script setup>
import {
  createInventory,
  updateInventory,
  findInventory
} from '@/api/inventory'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            croppingID: 0,
            sizeID: 0,
            processID: 0,
            total: 0,
            margin: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findInventory({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reinventory
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
               res = await createInventory(formData.value)
               break
             case 'update':
               res = await updateInventory(formData.value)
               break
             default:
               res = await createInventory(formData.value)
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
