
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Vip级别:" prop="levels">
          <el-input v-model="formData.levels" :clearable="true"  placeholder="请输入Vip级别" />
       </el-form-item>
        <el-form-item label="团队要求:" prop="team">
          <el-input v-model="formData.team" :clearable="true"  placeholder="请输入团队要求" />
       </el-form-item>
        <el-form-item label="业绩要求:" prop="come">
          <el-input-number v-model="formData.come" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="收益比率:" prop="rate">
          <el-input-number v-model="formData.rate" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="文本备注:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入文本备注" />
       </el-form-item>
        <el-form-item label="金额备注:" prop="descnum">
          <el-input-number v-model="formData.descnum" :precision="2" :clearable="true"></el-input-number>
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
  createCliSetvip,
  updateCliSetvip,
  findCliSetvip
} from '@/api/xiao/cliSetvip'

defineOptions({
    name: 'CliSetvipForm'
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
            levels: '',
            team: '',
            come: 0,
            rate: 0,
            desc: '',
            descnum: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCliSetvip({ ID: route.query.id })
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
               res = await createCliSetvip(formData.value)
               break
             case 'update':
               res = await updateCliSetvip(formData.value)
               break
             default:
               res = await createCliSetvip(formData.value)
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
