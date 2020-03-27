<template>
    <div>
        <!-- 开发中功能，若您发现这块代码可以研究，可以无视 -->
        <!-- 此版本为简单版 -->
        <!-- 结构体基础配置 -->
        <!-- develop分支中开发此功能 -->
        <el-form ref="form" :model="form" label-width="100px" :inline="true">
            <el-form-item label="结构名称" :span="8">
                <el-input v-model="form.structName"></el-input>
            </el-form-item>
            <el-form-item label="结构类型" :span="8">
               <el-select v-model="form.structType" multiple placeholder="请选择结构类型(多选)">
                    <el-option
                    v-for="item in options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                    </el-option>
                </el-select>
            </el-form-item>
        </el-form>
        <!-- 组件列表 -->
        <div class="button-box clearflex">
            <el-button @click="editAndAddComponent()" type="primary">新增组件</el-button>
        </div>
         <el-table
            :data="form.components"
             border stripe>
             <el-table-column
                type="index"
                label="序列"
                width="180">
            </el-table-column>
            <el-table-column
                prop="componentName"
                label="组件名"
                width="180">
            </el-table-column>
            <el-table-column
                prop="componentType"
                label="组件数据类型"
                width="180">
            </el-table-column>
            <el-table-column
                prop="componentShowType"
                label="组件展示类型">
            </el-table-column>
            <el-table-column
                prop="dictionaryName"
                label="字典名称（选）">
            </el-table-column>
            <el-table-column
                label="操作">
                <template slot-scope="scope">
                    <el-button @click="editAndAddComponent(scope.row)">编辑</el-button>
                    <el-popover
                    placement="top"
                    width="160"
                    v-model="scope.row.visible">
                    <p>这是一段内容这是一段内容确定删除吗？</p>
                    <div style="text-align: right; margin: 0">
                        <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                        <el-button type="primary" size="mini"  @click="deleteComponent(scope.$index)">确定</el-button>
                    </div>
                    <el-button slot="reference">删除</el-button>
                    </el-popover>
                </template>
            </el-table-column>
            </el-table>
            <!-- 组件列表 -->
        <div class="button-box clearflex">
            <el-button @click="enterForm" type="primary">提交</el-button>
        </div>

    <!-- 组件弹窗 -->
            <el-dialog title="组件内容" :visible.sync="dialogFlag">
                <ComponentDialog :dialogMiddle="dialogMiddle" />
                <div slot="footer" class="dialog-footer">
                    <el-button @click="closeDialog">取 消</el-button>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                </div>
            </el-dialog>
    </div>
</template>
<script>
const componentTemplate={
            componentName:"",
            componentType:"",
            componentShowType:"",
            dictionaryName:"",
            isMultiple:false,
            nideDictionary:false,
            visible:false,
            componentDictionary:[]
        }

import ComponentDialog from "@/view/superAdmin/autoCode/component/componentDialog.vue"
export default {
    name:"autoCode",
    data(){
        return{
            addFlag:"",
            form:{
                structName:"",
                structType:[],
                components:[]
            },
            options:[
                {label:"表格类型",value:"grid"},
                {label:"表单类型",value:"form"}
            ],
            dialogMiddle:{},
            bk:{},
            dialogFlag:false
        }
    },
    components:{
        ComponentDialog
    },
    methods:{
        editAndAddComponent(item){
            this.dialogFlag = true
            if(item){
                this.addFlag="edit"
                 this.bk=JSON.parse(JSON.stringify(item))
                 this.dialogMiddle = item
            }else{
                this.addFlag="add"
                this.dialogMiddle = JSON.parse(JSON.stringify(componentTemplate))
            }
        },
        enterDialog(){
            if(this.addFlag=="add"){
                this.form.components.push(this.dialogMiddle)
            }
            this.dialogFlag = false
        },
        closeDialog(){
            if(this.addFlag=="edit"){
                this.dialogMiddle = this.bk
            }
            this.dialogFlag = false
        },
        deleteComponent(index){
            this.form.components.splice(index,1)
        },
        enterForm(){
            console.log(this.form)
        }
    }
}
</script>
<style scope lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
</style>