<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="作业ID:" prop="executeId">
          <el-input v-model="formData.executeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件名称:" prop="fileName">
          <el-input v-model="formData.fileName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="序号:" prop="fileOrder">
          <el-input v-model.number="formData.fileOrder" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="拍摄位置:" prop="location">
          <el-input v-model="formData.location" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="拍摄时间:" prop="timestamp">
          <el-date-picker v-model="formData.timestamp" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="下载完成:" prop="downloaded">
          <el-input v-model.number="formData.downloaded" :clearable="true" placeholder="请输入" />
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
  name: 'FlyResult'
}
</script>

<script setup>
import {
  createFlyResult,
  updateFlyResult,
  findFlyResult
} from '@/api/flyResult'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            executeId: '',
            fileName: '',
            fileOrder: 0,
            type: 0,
            location: '',
            timestamp: new Date(),
            downloaded: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFlyResult({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reFlyRt
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
               res = await createFlyResult(formData.value)
               break
             case 'update':
               res = await updateFlyResult(formData.value)
               break
             default:
               res = await createFlyResult(formData.value)
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
