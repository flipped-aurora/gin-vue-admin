<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="机巢id:" prop="nestid">
          <el-input v-model="formData.nestid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机巢坐标:" prop="nestLocation">
          <el-input v-model="formData.nestLocation" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="飞机推流地址:" prop="aircraftVideoPushURL">
          <el-input v-model="formData.aircraftVideoPushURL" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="飞机推流播放地址:" prop="aircraftVideoURL">
          <el-input v-model="formData.aircraftVideoURL" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机巢监控播放地址:" prop="nestVideoURL">
          <el-input v-model="formData.nestVideoURL" :clearable="true" placeholder="请输入" />
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
  name: 'NestInfoBAC'
}
</script>

<script setup>
import {
  createNestInfo,
  updateNestInfo,
  findNestInfo
} from '@/api/nestInfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            nestid: '',
            nestLocation: '',
            aircraftVideoPushURL: '',
            aircraftVideoURL: '',
            nestVideoURL: '',
        })
// 验证规则
const rule = reactive({
               nestid : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nestLocation : [{
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
      const res = await findNestInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.renestinfo
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
               res = await createNestInfo(formData.value)
               break
             case 'update':
               res = await updateNestInfo(formData.value)
               break
             default:
               res = await createNestInfo(formData.value)
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
