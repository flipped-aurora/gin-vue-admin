
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户地址:" prop="address">
          <el-input v-model="formData.address" :clearable="false"  placeholder="请输入用户地址" />
       </el-form-item>
        <el-form-item label="父级节点:" prop="father">
          <el-input v-model="formData.father" :clearable="true"  placeholder="请输入父级节点" />
       </el-form-item>
        <el-form-item label="邀请节点:" prop="invite">
          <el-input v-model="formData.invite" :clearable="true"  placeholder="请输入邀请节点" />
       </el-form-item>
        <el-form-item label="左值:" prop="leftval">
          <el-input v-model.number="formData.leftval" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="右值:" prop="rightval">
          <el-input v-model.number="formData.rightval" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="层级:" prop="level">
          <el-input v-model.number="formData.level" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户等级:" prop="vip">
          <el-input v-model.number="formData.vip" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入备注" />
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
  createCliTree,
  updateCliTree,
  findCliTree
} from '@/api/xiao/cliTree'

defineOptions({
    name: 'CliTreeForm'
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
            father: '',
            invite: '',
            leftval: undefined,
            rightval: undefined,
            level: undefined,
            vip: undefined,
            description: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCliTree({ ID: route.query.id })
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
               res = await createCliTree(formData.value)
               break
             case 'update':
               res = await updateCliTree(formData.value)
               break
             default:
               res = await createCliTree(formData.value)
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
