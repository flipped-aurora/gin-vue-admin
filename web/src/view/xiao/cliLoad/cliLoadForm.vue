
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入用户地址" />
       </el-form-item>
        <el-form-item label="登录次数:" prop="loadtimes">
          <el-input v-model.number="formData.loadtimes" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最近登录:" prop="lasttime">
          <el-date-picker v-model="formData.lasttime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="登录IP:" prop="loadip">
          <el-input v-model="formData.loadip" :clearable="true"  placeholder="请输入登录IP" />
       </el-form-item>
        <el-form-item label="登录地址:" prop="loadaddr">
          <el-input v-model="formData.loadaddr" :clearable="true"  placeholder="请输入登录地址" />
       </el-form-item>
        <el-form-item label="地址余额:" prop="usdt">
          <el-input-number v-model="formData.usdt" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="当前状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入当前状态" />
       </el-form-item>
        <el-form-item label="文本备注:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入文本备注" />
       </el-form-item>
        <el-form-item label="金额备注:" prop="desnum">
          <el-input-number v-model="formData.desnum" :precision="2" :clearable="true"></el-input-number>
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
  createCliLoad,
  updateCliLoad,
  findCliLoad
} from '@/api/xiao/cliLoad'

defineOptions({
    name: 'CliLoadForm'
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
            address: '',
            loadtimes: undefined,
            lasttime: new Date(),
            loadip: '',
            loadaddr: '',
            usdt: 0,
            status: '',
            desc: '',
            desnum: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCliLoad({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
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
               res = await createCliLoad(formData.value)
               break
             case 'update':
               res = await updateCliLoad(formData.value)
               break
             default:
               res = await createCliLoad(formData.value)
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
