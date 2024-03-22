<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="课时名称:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="视频链接:" prop="videoUrl">
          <el-input v-model="formData.videoUrl" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="总时长:" prop="totalHour">
          <el-input v-model.number="formData.totalHour" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="授课教师:" prop="teacherId">
          <el-input v-model.number="formData.teacherId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="介绍:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="enable">
          <el-switch v-model="formData.enable" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  name: 'Class'
}
</script>

<script setup>
import {
  createClass,
  updateClass,
  findClass
} from '@/api/class'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            title: '',
            videoUrl: '',
            totalHour: 0,
            teacherId: 0,
            desc: '',
            enable: false,
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               videoUrl : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalHour : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               teacherId : [{
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
      const res = await findClass({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reclass
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
               res = await createClass(formData.value)
               break
             case 'update':
               res = await updateClass(formData.value)
               break
             default:
               res = await createClass(formData.value)
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
