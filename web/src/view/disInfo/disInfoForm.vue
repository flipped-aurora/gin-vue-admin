<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="帖子标题:" prop="disTitle">
          <el-input v-model="formData.disTitle" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="帖子内容:" prop="disContent">
          <RichEdit v-model="formData.disContent"/>
       </el-form-item>
        <el-form-item label="帖子关联比赛id:" prop="disComId">
          <el-input v-model.number="formData.disComId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发帖人id:" prop="disUserId">
          <el-input v-model.number="formData.disUserId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="帖子点赞人数:" prop="disLoveNumber">
          <el-input v-model.number="formData.disLoveNumber" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="帖子收藏人数:" prop="disCollectNumber">
          <el-input v-model.number="formData.disCollectNumber" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="帖子类型:" prop="disModel">
          <el-select v-model="formData.disModel" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in disModelOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="帖子图片:" prop="disPicture">
           <SelectImage v-model="formData.disPicture" multiple file-type="image"/>
       </el-form-item>
        <el-form-item label="帖子阅读量:" prop="disLookNumber">
          <el-input v-model.number="formData.disLookNumber" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="帖子状态:" prop="disStatus">
          <el-select v-model="formData.disStatus" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in DisStatuaOptions" :key="key" :label="item.label" :value="item.value" />
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
  createDisInfo,
  updateDisInfo,
  findDisInfo
} from '@/api/disInfo'

defineOptions({
    name: 'DisInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const disModelOptions = ref([])
const DisStatuaOptions = ref([])
const formData = ref({
            disTitle: '',
            disContent: '',
            disComId: 0,
            disUserId: 0,
            disLoveNumber: 0,
            disCollectNumber: 0,
            disModel: undefined,
            disPicture: [],
            disLookNumber: 0,
            disStatus: undefined,
        })
// 验证规则
const rule = reactive({
               disTitle : [{
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
      const res = await findDisInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redisData
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    disModelOptions.value = await getDictFunc('disModel')
    DisStatuaOptions.value = await getDictFunc('DisStatua')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createDisInfo(formData.value)
               break
             case 'update':
               res = await updateDisInfo(formData.value)
               break
             default:
               res = await createDisInfo(formData.value)
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
