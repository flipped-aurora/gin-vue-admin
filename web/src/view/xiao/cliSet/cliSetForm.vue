
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="静态收益:" prop="static">
          <el-input-number v-model="formData.static" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="直推收益:" prop="straight">
          <el-input-number v-model="formData.straight" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="间接推荐:" prop="inderict">
          <el-input-number v-model="formData.inderict" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="提币手续费:" prop="fee">
          <el-input-number v-model="formData.fee" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="产品类型:" prop="assets">
          <ArrayCtrl v-model="formData.assets" editable/>
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入状态" />
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
  createCliSet,
  updateCliSet,
  findCliSet
} from '@/api/xiao/cliSet'

defineOptions({
    name: 'CliSetForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 数组控制组件
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            static: 0,
            straight: 0,
            inderict: 0,
            fee: 0,
            assets: [],
            status: '',
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
      const res = await findCliSet({ ID: route.query.id })
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
               res = await createCliSet(formData.value)
               break
             case 'update':
               res = await updateCliSet(formData.value)
               break
             default:
               res = await createCliSet(formData.value)
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
