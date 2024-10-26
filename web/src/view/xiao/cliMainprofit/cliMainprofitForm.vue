
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入用户地址" />
       </el-form-item>
        <el-form-item label="结算次数:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="静态收益:" prop="static">
          <el-input-number v-model="formData.static" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="直推收益:" prop="pull">
          <el-input-number v-model="formData.pull" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="间接收益:" prop="indirect">
          <el-input-number v-model="formData.indirect" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="团队收益:" prop="team">
          <el-input-number v-model="formData.team" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="累积金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
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
  createCliMainprofit,
  updateCliMainprofit,
  findCliMainprofit
} from '@/api/xiao/cliMainprofit'

defineOptions({
    name: 'CliMainprofitForm'
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
            num: undefined,
            static: 0,
            pull: 0,
            indirect: 0,
            team: 0,
            amount: 0,
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
      const res = await findCliMainprofit({ ID: route.query.id })
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
               res = await createCliMainprofit(formData.value)
               break
             case 'update':
               res = await updateCliMainprofit(formData.value)
               break
             default:
               res = await createCliMainprofit(formData.value)
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
