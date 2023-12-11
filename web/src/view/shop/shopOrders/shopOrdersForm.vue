<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="APPID:" prop="AppId">
          <el-input v-model="formData.AppId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商户号:" prop="MchId">
          <el-input v-model="formData.MchId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商户订单:" prop="OutTradeNo">
          <el-input v-model="formData.OutTradeNo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付订单:" prop="TransactionId">
          <el-input v-model="formData.TransactionId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交易类型:" prop="TradeType">
          <el-input v-model="formData.TradeType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交易状态:" prop="TradeState">
          <el-select v-model="formData.TradeState" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in TradeStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="银行类型:" prop="BankType">
          <el-input v-model="formData.BankType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="附加数据:" prop="Attach">
          <el-input v-model="formData.Attach" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付时间:" prop="SuccessTime">
          <el-date-picker v-model="formData.SuccessTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="OPENID:" prop="OpenId">
          <el-input v-model="formData.OpenId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="sub_appid:" prop="SubOpenId">
          <el-input v-model="formData.SubOpenId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单金额:" prop="Total">
          <el-input v-model.number="formData.Total" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付金额:" prop="PayerTotal">
          <el-input v-model.number="formData.PayerTotal" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户ID:" prop="UserId">
          <el-input v-model.number="formData.UserId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址ID:" prop="AddrId">
          <el-input v-model.number="formData.AddrId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品ID:" prop="GoodsId">
          <el-input v-model.number="formData.GoodsId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="店铺ID:" prop="ShopID">
          <el-input v-model="formData.ShopID" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="配送ID:" prop="TransportId">
          <el-input v-model.number="formData.TransportId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结束时间:" prop="Endtime">
          <el-date-picker v-model="formData.Endtime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="商品标题:" prop="GoodsTitle">
          <el-input v-model="formData.GoodsTitle" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品价格:" prop="GoodsPrice">
          <el-input v-model.number="formData.GoodsPrice" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品邮费:" prop="PostFee">
          <el-input v-model.number="formData.PostFee" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="回调地址:" prop="CallUrl">
          <el-input v-model="formData.CallUrl" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付方式:" prop="PayMent">
          <el-select v-model="formData.PayMent" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in PayMentOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="设备ID:" prop="DeviceId">
          <el-input v-model="formData.DeviceId" :clearable="true" placeholder="请输入" />
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
  createShopOrders,
  updateShopOrders,
  findShopOrders
} from '@/api/shopOrders'

defineOptions({
    name: 'ShopOrdersForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const PayMentOptions = ref([])
const TradeStateOptions = ref([])
const formData = ref({
            AppId: '',
            MchId: '',
            OutTradeNo: '',
            TransactionId: '',
            TradeType: '',
            TradeState: undefined,
            BankType: '',
            Attach: '',
            SuccessTime: new Date(),
            OpenId: '',
            SubOpenId: '',
            Total: 0,
            PayerTotal: 0,
            UserId: 0,
            AddrId: 0,
            GoodsId: 0,
            ShopID: '',
            TransportId: 0,
            Endtime: new Date(),
            GoodsTitle: '',
            GoodsPrice: 0,
            PostFee: 0,
            CallUrl: '',
            PayMent: undefined,
            DeviceId: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findShopOrders({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reshopOrders
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    PayMentOptions.value = await getDictFunc('PayMent')
    TradeStateOptions.value = await getDictFunc('TradeState')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createShopOrders(formData.value)
               break
             case 'update':
               res = await updateShopOrders(formData.value)
               break
             default:
               res = await createShopOrders(formData.value)
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
