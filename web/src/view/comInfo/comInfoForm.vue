<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="比赛标题:" prop="comTitle">
          <el-input v-model="formData.comTitle" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="比赛简介:" prop="comIntroduction">
          <el-input v-model="formData.comIntroduction" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="比赛图片:" prop="comPicture">
          <SelectImage v-model="formData.comPicture" file-type="image"/>
       </el-form-item>
        <el-form-item label="比赛等级:" prop="comModel">
          <el-select v-model="formData.comModel" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in comModelOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="比赛开始时间:" prop="comStart">
          <el-date-picker v-model="formData.comStart" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="比赛结束时间:" prop="comEnd">
          <el-date-picker v-model="formData.comEnd" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="比赛参加人数:" prop="comNumber">
          <el-input v-model.number="formData.comNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="比赛类型:" prop="comType">
          <el-select v-model="formData.comType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in comTypeOptions" :key="key" :label="item.label" :value="item.value" />
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

<script setup>
import {
  createComInfo,
  updateComInfo,
  findComInfo
} from '@/api/comInfo'

defineOptions({
    name: 'ComInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const comModelOptions = ref([])
const comTypeOptions = ref([])
const formData = ref({
            comTitle: '',
            comIntroduction: '',
            comPicture: "",
            comModel: undefined,
            comStart: new Date(),
            comEnd: new Date(),
            comNumber: 0,
            comType: undefined,
        })
// 验证规则
const rule = reactive({
               comTitle : [{
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
      const res = await findComInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recomData
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    comModelOptions.value = await getDictFunc('comModel')
    comTypeOptions.value = await getDictFunc('comType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createComInfo(formData.value)
               break
             case 'update':
               res = await updateComInfo(formData.value)
               break
             default:
               res = await createComInfo(formData.value)
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
