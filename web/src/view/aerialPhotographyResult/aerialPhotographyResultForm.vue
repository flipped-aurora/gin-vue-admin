<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="拍摄时间:" prop="photographyCreatetime">
          <el-date-picker v-model="formData.photographyCreatetime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="上传人:" prop="uploadBy">
          <el-input v-model="formData.uploadBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="航拍文件:" prop="aerialPhotographyFile">
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="坐标:" prop="position">
          <el-input v-model="formData.position" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否加载:" prop="loadOrNot">
          <el-input v-model.number="formData.loadOrNot" :clearable="true" placeholder="请输入" />
       </el-form-item>
       <el-form-item label="机巢id集:"  prop="nestIds" >
          <el-input v-model="formData.nestIds" :clearable="true" placeholder="请输入" />
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
  name: 'AerialPhotographyResult'
}
</script>

<script setup>
import {
  createAerialPhotographyResult,
  updateAerialPhotographyResult,
  findAerialPhotographyResult
} from '@/api/aerialPhotographyResult'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectFile from '@/components/selectFile/selectFile.vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            photographyCreatetime: new Date(),
            uploadBy: '',
            type: 0,
            status: 0,
            position: '',
            loadOrNot: 0,
            nestIds: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAerialPhotographyResult({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reALPhotographyResult
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
               res = await createAerialPhotographyResult(formData.value)
               break
             case 'update':
               res = await updateAerialPhotographyResult(formData.value)
               break
             default:
               res = await createAerialPhotographyResult(formData.value)
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
