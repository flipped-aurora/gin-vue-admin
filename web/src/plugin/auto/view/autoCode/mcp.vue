<template>
  <div class="gva-form-box">
    <el-form :model="form" ref="formRef" label-width="100px" :rules="rules">
      <el-form-item label="工具名称" prop="name">
        <el-input v-model="form.name" placeholder="例:CurrentTime" />
      </el-form-item>
      <el-form-item label="工具描述" prop="description">
        <el-input type="textarea" v-model="form.description" placeholder="请输入工具描述" />
      </el-form-item>
      <el-form-item label="参数列表">
        <el-table :data="form.params"  style="width: 100%">
          <el-table-column prop="name" label="参数名" width="120">
            <template #default="scope">
              <el-input v-model="scope.row.name" placeholder="参数名" />
            </template>
          </el-table-column>
          <el-table-column prop="description" label="描述" min-width="180">
            <template #default="scope">
              <el-input v-model="scope.row.description" placeholder="描述" />
            </template>
          </el-table-column>
          <el-table-column prop="type" label="类型" width="120">
            <template #default="scope">
              <el-select v-model="scope.row.type" placeholder="类型">
                <el-option label="string" value="string" />
                <el-option label="number" value="number" />
                <el-option label="boolean" value="boolean" />
                <el-option label="object" value="object" />
                <el-option label="array" value="array" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column label="默认值" width="300">
            <template #default="scope">
              <el-input :disabled="scope.row.type === 'object'" v-model="scope.row.default" />
            </template>
          </el-table-column>
          <el-table-column prop="required" label="必填" width="80">
            <template #default="scope">
              <el-checkbox v-model="scope.row.required" />
            </template>
          </el-table-column>
          <el-table-column label="操作" width="80">
            <template #default="scope">
              <el-button type="text" @click="removeParam(scope.$index)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-form-item>
      <div class="flex justify-end">
        <el-button type="primary" icon="plus" @click="addParam" style="margin-top: 10px;">添加参数</el-button>
      </div>
      <el-form-item label="返回参数">
        <el-table :data="form.response" style="width: 100%">
          <el-table-column prop="type" label="类型" min-width="120">
            <template #default="scope">
              <el-select v-model="scope.row.type" placeholder="类型">
                <el-option label="text" value="text" />
                <el-option label="image" value="image" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="80">
            <template #default="scope">
              <el-button type="text" @click="removeResponse(scope.$index)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-form-item>
      <div class="flex justify-end">
        <el-button type="primary" icon="plus" @click="addResponse" style="margin-top: 10px;">添加返回参数</el-button>
      </div>

      <div class="flex justify-end mt-8">
        <el-button type="primary" @click="submit">生成</el-button>
      </div>
    </el-form>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { mcp } from '@/api/autoCode'

defineOptions({
  name: 'MCP'
})

const formRef = ref(null)
const form = reactive({
  name: '',
  description: '',
  type: '',
  params: [],
  response: []
})

const rules = {
  name: [{ required: true, message: '请输入工具名称', trigger: 'blur' }],
  description: [{ required: true, message: '请输入工具描述', trigger: 'blur' }],
  type: [{ required: true, message: '请选择类型', trigger: 'change' }]
}

function addParam() {
  form.params.push({
    name: '',
    description: '',
    type: '',
    required: false
  })
}

function removeParam(index) {
  form.params.splice(index, 1)
}

function addResponse() {
  form.response.push({
    type: ''
  })
}

function removeResponse(index) {
  form.response.splice(index, 1)
}

function submit() {
  formRef.value.validate(async (valid) => {
    if (!valid) return
    // 简单校验参数
    for (const p of form.params) {
      if (!p.name || !p.description || !p.type) {
        ElMessage.error('请完善所有参数信息')
        return
      }
    }
    // 校验返回参数
    for (const r of form.response) {
      if (!r.type) {
        ElMessage.error('请完善所有返回参数类型')
        return
      }
    }
      const res = await mcp(form)
      if (res.code === 0) {
        ElMessage.success(res.msg)
      }
  })
}
</script>
