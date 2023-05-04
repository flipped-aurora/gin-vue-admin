<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户id:" prop="userID">
          <el-input v-model.number="formData.userID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="消息类型:" prop="msgType">
          <el-select v-model="formData.msgType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in msgTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="消息id:" prop="msgID">
          <el-input v-model.number="formData.msgID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in handleStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'MsgBox'
}
</script>

<script setup>
import {
  createMsgBox,
  updateMsgBox,
  findMsgBox
} from '@/api/msgBox'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const msgTypeOptions = ref([])
const handleStatusOptions = ref([])
const formData = ref({
            userID: 0,
            msgType: undefined,
            msgID: 0,
            status: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMsgBox({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remsgBox
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    msgTypeOptions.value = await getDictFunc('msgType')
    handleStatusOptions.value = await getDictFunc('handleStatus')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createMsgBox(formData.value)
               break
             case 'update':
               res = await updateMsgBox(formData.value)
               break
             default:
               res = await createMsgBox(formData.value)
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
