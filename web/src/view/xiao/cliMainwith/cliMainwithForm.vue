
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入用户地址" />
       </el-form-item>
        <el-form-item label="提币次数:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="可提金额:" prop="withable">
          <el-input-number v-model="formData.withable" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="已提金额:" prop="withed">
          <el-input-number v-model="formData.withed" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="累计总额:" prop="total">
          <el-input-number v-model="formData.total" :precision="2" :clearable="true"></el-input-number>
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
  createCliMainwith,
  updateCliMainwith,
  findCliMainwith
} from '@/api/xiao/cliMainwith'

defineOptions({
    name: 'CliMainwithForm'
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
            withable: 0,
            withed: 0,
            total: 0,
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
      const res = await findCliMainwith({ ID: route.query.id })
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
               res = await createCliMainwith(formData.value)
               break
             case 'update':
               res = await updateCliMainwith(formData.value)
               break
             default:
               res = await createCliMainwith(formData.value)
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
