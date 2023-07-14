<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="航线id:" prop="missionid">
          <el-input v-model="formData.missionid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="航线名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="航线类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="自动飞行速度:" prop="autoFlightSpeed">
          <el-input v-model.number="formData.autoFlightSpeed" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="起飞模式:" prop="gotoFirstWaypointMode">
          <el-input v-model.number="formData.gotoFirstWaypointMode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结束模式:" prop="finishAction">
          <el-input v-model.number="formData.finishAction" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="路径模式:" prop="flightPathMode">
          <el-input v-model.number="formData.flightPathMode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="朝向模式:" prop="headingMode">
          <el-input v-model.number="formData.headingMode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="参数json体:" prop="param">
          <el-input v-model="formData.param" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="安全:" prop="safealt">
          <el-input v-model="formData.safealt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="kml:" prop="kml">
          <el-input v-model="formData.kml" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="gps:" prop="gps">
          <el-input v-model="formData.gps" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="站点id:" prop="station">
          <el-input v-model="formData.station" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="明确定位:" prop="clearHomeLocation">
          <el-input v-model="formData.clearHomeLocation" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="制作人:" prop="producer">
          <el-input v-model="formData.producer" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="制作单位:" prop="productionUnit">
          <el-input v-model="formData.productionUnit" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="isactive:" prop="isActive">
          <el-input v-model="formData.isActive" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="固定返航点:" prop="fixedReturnPoint">
          <el-input v-model="formData.fixedReturnPoint" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机巢id:" prop="nestId">
          <el-input v-model="formData.nestId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入" />
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
  name: 'NestAirline'
}
</script>

<script setup>
import {
  createNestAirline,
  updateNestAirline,
  findNestAirline
} from '@/api/nestAirline'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            missionid: '',
            name: '',
            type: '',
            autoFlightSpeed: 0,
            gotoFirstWaypointMode: 0,
            finishAction: 0,
            flightPathMode: 0,
            headingMode: 0,
            param: '',
            safealt: '',
            kml: '',
            gps: '',
            station: '',
            clearHomeLocation: '',
            producer: '',
            productionUnit: '',
            isActive: '',
            fixedReturnPoint: '',
            nestId: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findNestAirline({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reNtAirline
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
               res = await createNestAirline(formData.value)
               break
             case 'update':
               res = await updateNestAirline(formData.value)
               break
             default:
               res = await createNestAirline(formData.value)
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
