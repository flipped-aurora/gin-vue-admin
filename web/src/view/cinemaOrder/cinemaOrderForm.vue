<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="座位ID:" prop="seatId">
          <el-input v-model.number="formData.seatId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电影院:" prop="filmId">
          <el-input v-model.number="formData.filmId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="影厅:" prop="filmHall">
          <el-switch v-model="formData.filmHall" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="选座:" prop="filmSeat">
          <el-input v-model="formData.filmSeat" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电影价格:" prop="filmName">
          <el-input v-model="formData.filmName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电影类型:" prop="filmType">
          <el-input v-model="formData.filmType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="播放时间:" prop="playTime">
          <el-date-picker v-model="formData.playTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="价格:" prop="filmPrice">
          <el-input-number v-model="formData.filmPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="状态 1. 正常 0.退款:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createCinemaOrder,
  updateCinemaOrder,
  findCinemaOrder
} from '@/api/cinemaOrder'

defineOptions({
    name: 'CinemaOrderForm'
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
            seatId: 0,
            filmId: 0,
            filmHall: false,
            filmSeat: '',
            filmName: '',
            filmType: '',
            playTime: new Date(),
            filmPrice: 0,
            status: false,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCinemaOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recinemaOrder
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
               res = await createCinemaOrder(formData.value)
               break
             case 'update':
               res = await updateCinemaOrder(formData.value)
               break
             default:
               res = await createCinemaOrder(formData.value)
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
