<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="应用名称（中文）:" prop="appName">
          <el-input v-model="formData.appName" :clearable="true"  placeholder="请输入应用名称（中文）" />
       </el-form-item>
        <el-form-item label="应用名称（英文标识）:" prop="appCode">
          <el-input v-model="formData.appCode" :clearable="true"  placeholder="请输入应用名称（英文标识）" />
       </el-form-item>
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入标题" />
       </el-form-item>
        <el-form-item label="应用介绍:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入应用介绍" />
       </el-form-item>
        <el-form-item label="分类:" prop="classify">
          <el-input v-model="formData.classify" :clearable="true"  placeholder="请输入分类" />
       </el-form-item>
        <el-form-item label="应用版本:" prop="version">
          <el-input v-model="formData.version" :clearable="true"  placeholder="请输入应用版本" />
       </el-form-item>
        <el-form-item label="收费模式:" prop="mode">
           <el-select v-model="formData.mode" placeholder="请选择收费模式" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in price_modeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="后续迭代:" prop="developMode">
           <el-select v-model="formData.developMode" placeholder="请选择后续迭代" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dev_modeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="封面地址:" prop="cover">
          <el-input v-model="formData.cover" :clearable="true"  placeholder="请输入封面地址" />
       </el-form-item>
        <el-form-item label="应用标签:" prop="tags">
          <el-input v-model="formData.tags" :clearable="true"  placeholder="请输入应用标签" />
       </el-form-item>
        <el-form-item label="介绍视频:" prop="video">
          <el-input v-model="formData.video" :clearable="true"  placeholder="请输入介绍视频" />
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
  createBizAppHub,
  updateBizAppHub,
  findBizAppHub
} from '@/api/biz_apphub/biz_apphub'

defineOptions({
    name: 'BizAppHubForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const price_modeOptions = ref([])
const dev_modeOptions = ref([])
const formData = ref({
            appName: '',
            appCode: '',
            title: '',
            desc: '',
            classify: '',
            version: '',
            mode: '',
            developMode: '',
            cover: '',
            tags: '',
            video: '',
        })
// 验证规则
const rule = reactive({
               appName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               appCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               desc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               classify : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               version : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               developMode : [{
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
      const res = await findBizAppHub({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    price_modeOptions.value = await getDictFunc('price_mode')
    dev_modeOptions.value = await getDictFunc('dev_mode')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createBizAppHub(formData.value)
               break
             case 'update':
               res = await updateBizAppHub(formData.value)
               break
             default:
               res = await createBizAppHub(formData.value)
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
