<template>
    <div>
        <el-form label-width="100px" :model="form" ref="form">
            <el-form-item label="工作流名称" >
                <el-input type="text" v-model="form.workflowNickName"></el-input>
            </el-form-item>
            <el-form-item label="工作流英文id" >
                <el-input type="text" v-model="form.workflowName"></el-input>
            </el-form-item>
            <el-form-item label="工作流描述" >
                <el-input type="text" v-model="form.workflowDescription"></el-input>
            </el-form-item>
        </el-form>
        <el-table
                :data="tableData"
                border
                style="width: 100%">
            <el-table-column
                    prop="isEnd"
                    label="是否是完结流节点">
                <template scope="scope">
                    <el-select v-model="scope.row.isEnd" placeholder="请选择">
                        <el-option v-for="item in options"
                        :label="item.name"
                        :value="item.value">
                        </el-option>
                    </el-select>
                </template>
            </el-table-column>
            <el-table-column
                    prop="isStrat"
                    label="是否是开始流节点">
                <template scope="scope">
                    <el-select v-model="scope.row.isStrat" placeholder="请选择">
                        <el-option v-for="item in options"
                                   :label="item.name"
                                   :value="item.value">
                        </el-option>
                    </el-select>
                </template>

            </el-table-column>
            <el-table-column
                    prop="stepAuthorityID"
                    label="操作者级别id">
                <template scope="scope">
                    <el-input type="text" placeholder="请输入" v-model="scope.row.stepAuthorityID"></el-input>
                </template>
            </el-table-column>
            <el-table-column
                    prop="stepName"
                    label="工作流名称">
                <template scope="scope">
                    <el-input type="text" placeholder="请输入" v-model="scope.row.stepName"></el-input>
                </template>
            </el-table-column>
            <el-table-column
                    prop="stepNo"
                    label="步骤id">
                <template scope="scope">
                    <el-input type="text" placeholder="请输入" v-model="scope.row.stepNo"></el-input>
                </template>
            </el-table-column>
        </el-table>
        <el-button type="primary" @click="submit">提交</el-button>
    </div>
</template>

<script>
import { createWorkFlow } from '@/api/workflow'
    export default {
        name: "workflow",
        data() {
            return {
                form: {
                    workflowName:'',
                    workflowDescription:'',
                    workflowNickName:''
                },
                tableData: [{
                    isEnd: '',
                    isStrat: '',
                    stepAuthorityID: '',
                    stepName:'',
                    stepNo:''
                }],
                options:[
                    {
                        name:'是',
                        value:true
                    },
                    {
                        name:'否',
                        value:false
                    }
                ]
            }
        },
        component: {},
        methods: {
            async submit() {
                let params = {
                        "workflowDescription": this.form.workflowDescription,
                            "workflowName": this.form.workflowName,
                            "workflowNickName": this.form.workflowNickName,
                            "workflowStep": [
                        {
                            "isEnd": this.tableData[0].isEnd,
                            "isStrat": this.tableData[0].isStrat,
                            "stepAuthorityID": this.tableData[0].stepAuthorityID,
                            "stepName": this.tableData[0].stepName,
                            "stepNo": this.tableData[0].stepNo
                        }]
                }
                const res = await createWorkFlow(params)
                console.log(res)
            }
        }
    }
</script>
<style scoped>

</style>
