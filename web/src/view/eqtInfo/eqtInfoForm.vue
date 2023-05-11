<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="设备类型:" prop="eqtClass">
          <el-select v-model="formData.eqtClass" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in eqtClassOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="卡片码:" prop="eqtCardNo">
          <el-input v-model="formData.eqtCardNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="SN号:" prop="eqtSn">
          <el-input v-model="formData.eqtSn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="设备名:" prop="eqtName">
          <el-input v-model="formData.eqtName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="设备型号:" prop="eqtModel">
          <el-input v-model="formData.eqtModel" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="设备状态:" prop="eqtStatus">
          <el-select v-model="formData.eqtStatus" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in reservationStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="库存状态:" prop="eqtStockStatus">
          <el-select v-model="formData.eqtStockStatus" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in eqtStockOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="当前科室:" prop="currentDept">
          <el-input v-model.number="formData.currentDept" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="serviceTime字段:" prop="serviceTime">
          <el-date-picker v-model="formData.serviceTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  name: 'EqtInfo'
}
</script>

<script setup>
import {
  createEqtInfo,
  updateEqtInfo,
  findEqtInfo
} from '@/api/eqtInfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const eqtClassOptions = ref([])
const reservationStatusOptions = ref([])
const eqtStockOptions = ref([])
const formData = ref({
            eqtClass: undefined,
            eqtCardNo: '',
            eqtSn: '',
            eqtName: '',
            eqtModel: '',
            eqtStatus: undefined,
            eqtStockStatus: undefined,
            currentDept: 0,
            serviceTime: new Date(),
        })
// 验证规则
const rule = reactive({
               eqtClass : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               eqtName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               eqtModel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               eqtStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               eqtStockStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEqtInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reeqtInfo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    eqtClassOptions.value = await getDictFunc('eqtClass')
    reservationStatusOptions.value = await getDictFunc('reservationStatus')
    eqtStockOptions.value = await getDictFunc('eqtStock')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createEqtInfo(formData.value)
               break
             case 'update':
               res = await updateEqtInfo(formData.value)
               break
             default:
               res = await createEqtInfo(formData.value)
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
