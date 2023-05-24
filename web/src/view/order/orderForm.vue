<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单号:" prop="orderNo">
          <el-input v-model="formData.orderNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付单号:" prop="payNo">
          <el-input v-model="formData.payNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公司id:" prop="companyID">
          <el-input v-model.number="formData.companyID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户id:" prop="userID">
          <el-input v-model.number="formData.userID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="实付金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="订单状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in orderStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="支付状态:" prop="payStatus">
          <el-select v-model="formData.payStatus" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in payStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="天数:" prop="day">
          <el-input v-model.number="formData.day" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="职员数量:" prop="clerkCount">
          <el-input v-model.number="formData.clerkCount" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付时间:" prop="payAt">
          <el-date-picker v-model="formData.payAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  name: 'Order'
}
</script>

<script setup>
import {
  createOrder,
  updateOrder,
  findOrder
} from '@/api/order'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const orderStatusOptions = ref([])
const payStatusOptions = ref([])
const formData = ref({
            orderNo: '',
            payNo: '',
            companyID: 0,
            userID: 0,
            price: 0,
            amount: 0,
            status: undefined,
            payStatus: undefined,
            day: 0,
            clerkCount: 0,
            payAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reorder
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    orderStatusOptions.value = await getDictFunc('orderStatus')
    payStatusOptions.value = await getDictFunc('payStatus')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createOrder(formData.value)
               break
             case 'update':
               res = await updateOrder(formData.value)
               break
             default:
               res = await createOrder(formData.value)
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
