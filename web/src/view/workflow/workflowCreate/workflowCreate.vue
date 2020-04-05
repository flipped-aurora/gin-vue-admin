<template>
  <div>
    <el-form :model="form" label-width="100px" ref="form">
      <el-form-item label="工作流名称">
        <el-input type="text" v-model="form.workflowNickName"></el-input>
      </el-form-item>
      <el-form-item label="工作流英文id">
        <el-input type="text" v-model="form.workflowName"></el-input>
      </el-form-item>
      <el-form-item label="工作流描述">
        <el-input type="text" v-model="form.workflowDescription"></el-input>
      </el-form-item>
    </el-form>
    <el-button @click="createWorkflowStep" class="fl-right mg" type="primary">新增</el-button>
    <el-table :data="form.workflowStep" border style="width: 100%">
      <el-table-column label="是否是完结流节点" prop="isEnd">
        <template slot-scope="scope">
          <el-select placeholder="请选择" v-model="scope.row.isEnd">
            <el-option
              :key="key"
              :label="item.name"
              :value="item.value"
              v-for="(item,key) in options"
            ></el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column label="是否是开始流节点" prop="isStrat">
        <template slot-scope="scope">
          <el-select placeholder="请选择" v-model="scope.row.isStrat">
            <el-option
              :key="key"
              :label="item.name"
              :value="item.value"
              v-for="(item,key) in options"
            ></el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column label="操作者级别id" prop="stepAuthorityID">
        <template slot-scope="scope">
          <el-input placeholder="请输入" type="text" v-model="scope.row.stepAuthorityID"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="工作流名称" prop="stepName">
        <template slot-scope="scope">
          <el-input placeholder="请输入" type="text" v-model="scope.row.stepName"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="步骤id" prop="stepNo">
        <template slot-scope="scope">
          <el-input placeholder="请输入" type="text" v-model="scope.row.stepNo"></el-input>
        </template>
      </el-table-column>
    </el-table>
    <el-button @click="submit" type="primary" class="fl-right mg">提交</el-button>
  </div>
</template>

<script>
import { createWorkFlow } from '@/api/workflow'
export default {
  name: 'Workflow',
  data() {
    return {
      form: {
        workflowName: '',
        workflowDescription: '',
        workflowNickName: '',
        workflowStep: [
          {
            isEnd: false,
            isStrat: true,
            stepAuthorityID: '',
            stepName: '',
            stepNo: ''
          }
        ]
      },
      options: [
        {
          name: '是',
          value: true
        },
        {
          name: '否',
          value: false
        }
      ]
    }
  },
  methods: {
    createWorkflowStep() {
      this.form.workflowStep.push({
        isEnd: false,
        isStrat: false,
        stepAuthorityID: '',
        stepName: '',
        stepNo: ''
      })
    },
    async submit() {
      const res = await createWorkFlow(this.form)
      if(res.code == 0){
          this.$message({
              message:"创建成功",
              type:"success"
          })
      }
    }
  }
}
</script>
<style>
</style>
