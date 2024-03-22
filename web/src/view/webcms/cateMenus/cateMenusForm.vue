<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="菜单等级:" prop="menuLevel">
          <el-input v-model.number="formData.menuLevel" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="父菜单ID:" prop="parentId">
          <el-input v-model="formData.parentId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课程名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否在列表隐藏:" prop="hidden">
          <el-switch v-model="formData.hidden" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="排序标记:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入" />
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
  name: 'CateMenus'
}
</script>

<script setup>
import {
  createCateMenus,
  updateCateMenus,
  findCateMenus
} from '@/api/cateMenus'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            menuLevel: 0,
            parentId: '',
            name: '',
            hidden: false,
            sort: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCateMenus({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recateMenus
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
               res = await createCateMenus(formData.value)
               break
             case 'update':
               res = await updateCateMenus(formData.value)
               break
             default:
               res = await createCateMenus(formData.value)
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
