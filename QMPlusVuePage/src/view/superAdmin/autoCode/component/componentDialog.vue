<template>
    <div>
        <el-form :model="dialogMiddle"  :inline="true" >
            <el-form-item label="组件名称" label-width="80" :span="8">
                <el-input v-model="dialogMiddle.componentName" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="数据类型" label-width="80" :span="8">
            <el-select v-model="dialogMiddle.componentType" placeholder="请选择活动区域">
                <el-option
                    v-for="item in typeOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                </el-option>
            </el-select>
            </el-form-item>
            <el-form-item label="展示类型" label-width="80" :span="8">
            <el-select v-model="dialogMiddle.componentShowType" placeholder="请选择活动区域">
                <el-option
                    v-for="item in showTypeOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                    </el-option>
            </el-select>
            </el-form-item>
            <el-form-item label="字典Key" label-width="80" :span="8">
                <el-input v-model="dialogMiddle.dictionaryName" autocomplete="off"></el-input>
            </el-form-item>
            <div>
            <el-form-item label="是否多选" label-width="80">
                <el-switch
                    v-model="dialogMiddle.isMultiple"
                    active-text="多选"
                    inactive-text="单选">
                </el-switch>
            </el-form-item>
            </div>
            <div>
            <el-form-item label="是否使用字典" label-width="80">
                <el-switch
                    v-model="dialogMiddle.nideDictionary"
                    active-text="使用"
                    inactive-text="不使用">
                </el-switch>
            </el-form-item>
            </div>
        </el-form>
        <div class="button-box clearflex">
            <el-button @click="addDictionary" type="primary">新增字典</el-button>
        </div>
         <el-table
            :data="dialogMiddle.componentDictionary"
            stripe
            style="width: 100%">
            <el-table-column
            label="展示值"
            >
            <template slot-scope="scope">
                <el-input v-model="scope.row.label"></el-input>
            </template>
            </el-table-column>
            <el-table-column
            label="交互值"
            >
            <template slot-scope="scope">
                <el-input v-model="scope.row.value"></el-input>
            </template>
            </el-table-column>
             <el-table-column
            label="操作"
            >
            <template slot-scope="scope">
                <el-popover
                    placement="top"
                    width="160"
                    v-model="scope.row.visible">
                    <p>这是一段内容这是一段内容确定删除吗？</p>
                    <div style="text-align: right; margin: 0">
                        <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                        <el-button type="primary" size="mini" @click="deleteRow(scope)">确定</el-button>
                    </div>
                    <el-button type="text" slot="reference">删除</el-button>
                    </el-popover>
            </template>
            </el-table-column>
        </el-table>
    </div>
</template>
<script>
const dictionaryTemplate = {
    label:"",
    value:"",
    visible:false
}
export default {
    name:"ComponentDialog",
    props:{
        dialogMiddle:{
            type:Object,
            default:function(){
                return {}
            }
        }
    },
    data(){
        return{
            visible:false,
            typeOptions:[
                {
                    label:"字符串",
                    value:"string"
                },
                {
                    label:"整型",
                    value:"int"
                },
                {
                    label:"布尔值",
                    value:"bool"
                },
                {
                    label:"浮点型",
                    value:"float64"
                },
                {
                    label:"时间",
                    value:"time.Time"
                },
            ],
            showTypeOptions:[
                {
                    label:"单选框",
                    value:"radio"
                },
                {
                    label:"多选框",
                    value:"checkBox"
                },
                {
                    label:"输入框",
                    value:"input"
                },
                {
                    label:"计数器",
                    value:"inputNumber"
                },
                {
                    label:"选择器",
                    value:"select"
                },
                {
                    label:"级联选择器",
                    value:"cascader"
                },
                {
                    label:"开关",
                    value:"switch"
                },
                {
                    label:"时间日期选择器",
                    value:"datePicker"
                }
            ]
        }
    },
    methods:{
        addDictionary(){
            this.dialogMiddle.componentDictionary.push({...dictionaryTemplate})
        },
        deleteRow(row){
            this.dialogMiddle.componentDictionary.splice(row.$index,1)
        }
    }
}
</script>
<style lang="scss">
    
</style>