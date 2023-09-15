<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="机器简称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="项目编号:" prop="code">
          <el-input v-model="formData.code" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="删除状态:" prop="isRemoved">
          <el-switch v-model="formData.isRemoved" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="ip地址:" prop="ip">
          <el-input v-model="formData.ip" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="端口号:" prop="port">
          <el-input v-model="formData.port" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="账号:" prop="userName">
          <el-input v-model="formData.userName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSysRemotes,
  updateSysRemotes,
  findSysRemotes
} from '@/api/sysRemotes'

defineOptions({
    name: 'SysRemotesForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            code: '',
            status: false,
            isRemoved: false,
            ip: '',
            port: '',
            userName: '',
            password: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               code : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               isRemoved : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               ip : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               port : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               password : [{
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
      const res = await findSysRemotes({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resysRemotes
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
               res = await createSysRemotes(formData.value)
               break
             case 'update':
               res = await updateSysRemotes(formData.value)
               break
             default:
               res = await createSysRemotes(formData.value)
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
