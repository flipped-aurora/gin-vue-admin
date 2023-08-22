<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单ID:" prop="orderId">
          <el-input v-model="formData.orderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单取票号:" prop="ticketNumber">
          <el-input v-model="formData.ticketNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单联系电话:" prop="contactPhone">
          <el-input v-model="formData.contactPhone" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否换乘:" prop="isTransfer">
          <el-switch v-model="formData.isTransfer" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="是否占座:" prop="isOccupySeat">
          <el-input v-model.number="formData.isOccupySeat" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="完成状态:" prop="completeStatus">
          <el-input v-model.number="formData.completeStatus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出票失败原因:" prop="failReason">
          <el-input v-model="formData.failReason" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备名称:" prop="machineName">
          <el-input v-model="formData.machineName" :clearable="true" placeholder="请输入" />
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
  name: 'OrderInfo'
}
</script>

<script setup>
import {
  createOrderInfo,
  updateOrderInfo,
  findOrderInfo
} from '@/api/orderinfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            orderId: '',
            ticketNumber: '',
            contactPhone: '',
            isTransfer: false,
            isOccupySeat: 0,
            completeStatus: 0,
            failReason: '',
            machineName: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOrderInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reorderInfo
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
               res = await createOrderInfo(formData.value)
               break
             case 'update':
               res = await updateOrderInfo(formData.value)
               break
             default:
               res = await createOrderInfo(formData.value)
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
