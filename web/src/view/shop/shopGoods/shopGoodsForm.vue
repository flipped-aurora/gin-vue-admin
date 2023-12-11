<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="店铺ID:" prop="goodsShopid">
          <el-input v-model="formData.goodsShopid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备ID:" prop="goodsMacid">
          <el-input v-model="formData.goodsMacid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品名称:" prop="goodsName">
          <el-input v-model="formData.goodsName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品描述:" prop="goodsDesc">
          <el-input v-model="formData.goodsDesc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品价格:" prop="goodsPrice">
          <el-input v-model.number="formData.goodsPrice" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="市场价:" prop="goodsMaketPrice">
          <el-input v-model.number="formData.goodsMaketPrice" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品库存:" prop="goodsStock">
          <el-input v-model.number="formData.goodsStock" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="图片链接:" prop="goodsImg">
          <el-input v-model="formData.goodsImg" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="上架时间:" prop="goodsSellTime">
          <el-date-picker v-model="formData.goodsSellTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="商品状态:" prop="goodsStatus">
          <el-select v-model="formData.goodsStatus" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in itemStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="商品邮费:" prop="goodsPostfee">
          <el-input v-model.number="formData.goodsPostfee" :clearable="true" placeholder="请输入" />
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
  createShopGoods,
  updateShopGoods,
  findShopGoods
} from '@/api/shopGoods'

defineOptions({
    name: 'ShopGoodsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const itemStatusOptions = ref([])
const formData = ref({
            goodsShopid: '',
            goodsMacid: '',
            goodsName: '',
            goodsDesc: '',
            goodsPrice: 0,
            goodsMaketPrice: 0,
            goodsStock: 0,
            goodsImg: '',
            goodsSellTime: new Date(),
            goodsStatus: undefined,
            goodsPostfee: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findShopGoods({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reshopGoods
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    itemStatusOptions.value = await getDictFunc('itemStatus')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createShopGoods(formData.value)
               break
             case 'update':
               res = await updateShopGoods(formData.value)
               break
             default:
               res = await createShopGoods(formData.value)
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
