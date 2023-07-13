<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建人:" prop="createdBy">
          <el-input v-model="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="抄送人:" prop="copyTo">
          <el-input v-model="formData.copyTo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="航线名称:" prop="missionName">
          <el-input v-model="formData.missionName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="航线id:" prop="missionid">
          <el-input v-model="formData.missionid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="计划时间:" prop="planAt">
          <el-date-picker v-model="formData.planAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="作业ID:" prop="executeId">
          <el-input v-model="formData.executeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="作作业人:" prop="executeAt">
          <el-input v-model="formData.executeAt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="飞行秒:" prop="flyInSecond">
          <el-input v-model.number="formData.flyInSecond" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="全景链接:" prop="panoramaLink">
          <el-input v-model="formData.panoramaLink" :clearable="true" placeholder="请输入" />
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
  name: 'NestExecRecord'
}
</script>

<script setup>
import {
  createNestExecRecord,
  updateNestExecRecord,
  findNestExecRecord
} from '@/api/nestExecRecord'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            status: '',
            createdBy: '',
            copyTo: '',
            missionName: '',
            missionid: '',
            planAt: new Date(),
            type: '',
            executeId: '',
            executeAt: '',
            flyInSecond: 0,
            panoramaLink: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findNestExecRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reNtERecord
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
               res = await createNestExecRecord(formData.value)
               break
             case 'update':
               res = await updateNestExecRecord(formData.value)
               break
             default:
               res = await createNestExecRecord(formData.value)
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
