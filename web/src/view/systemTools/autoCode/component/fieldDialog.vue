<template>
    <div>
        <span style="color:red">搜索时如果条件为LIKE只支持字符串</span>
        <el-form :model="dialogMiddle" ref="fieldDialogFrom" label-width="120px" label-position="left" :rules="rules">
            <el-form-item label="Field名称" prop="fieldName">
                <el-col :span="6">
                <el-input v-model="dialogMiddle.fieldName" autocomplete="off"></el-input>
                </el-col>
            </el-form-item>
            <el-form-item label="Field中文名" prop="fieldDesc">
                <el-col :span="6">
                <el-input v-model="dialogMiddle.fieldDesc" autocomplete="off"></el-input>
                </el-col>
            </el-form-item>
            <el-form-item label="FieldJSON" prop="fieldJson">
                <el-col :span="6">
                <el-input v-model="dialogMiddle.fieldJson" autocomplete="off"></el-input>
                </el-col>
            </el-form-item>
            <el-form-item label="数据库字段名" prop="columnName">
                <el-col :span="6">
                <el-input v-model="dialogMiddle.columnName" autocomplete="off"></el-input>
                </el-col>
            </el-form-item>
            <el-form-item label="数据库字段描述" prop="comment">
                <el-col :span="6">
                <el-input v-model="dialogMiddle.comment" autocomplete="off"></el-input>
                </el-col>
            </el-form-item>
            <el-form-item label="Field数据类型" prop="fieldType">
                <el-col :span="8">
            <el-select v-model="dialogMiddle.fieldType" placeholder="请选择field数据类型" @change="getDbfdOptions" clearable >
                <el-option
                    v-for="item in typeOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                </el-option>
            </el-select>
                </el-col>
            </el-form-item>

            <el-form-item label="数据库字段类型" prop="dataType">
                <el-col :span="8">
            <el-select :disabled="!dialogMiddle.fieldType" v-model="dialogMiddle.dataType" placeholder="请选择数据库字段类型" clearable >
                <el-option
                    v-for="item in dbfdOptions"
                    :key="item.label"
                    :label="item.label"
                    :value="item.label">
                </el-option>
            </el-select>
                </el-col>
            </el-form-item>
            <el-form-item label="数据库字段长度" prop="dataTypeLong">
                <el-col :span="6">
                <el-input :disabled="!dialogMiddle.dataType" v-model="dialogMiddle.dataTypeLong"></el-input>
                </el-col>
            </el-form-item>
            <el-form-item label="Field查询条件" prop="fieldSearchType">
                <el-col :span="8">
            <el-select v-model="dialogMiddle.fieldSearchType" placeholder="请选择Field查询条件" clearable >
                <el-option
                    v-for="item in typeSearchOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                </el-option>
            </el-select>
                </el-col>
            </el-form-item>
        </el-form>
    </div>
</template>
<script>
import {getDict} from '@/utils/dictionary'
export default {
    name:"FieldDialog",
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
            dbfdOptions:[],
            visible:false,
            typeSearchOptions:[
                {
                    label:"=",
                    value:"="
                },
                {
                    label:"<>",
                    value:"<>"
                },
                {
                    label:">",
                    value:">"
                },
                {
                    label:"<",
                    value:"<"
                },
                {
                    label:"LIKE",
                    value:"LIKE"
                },
            ],
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
            rules:{
                fieldName:[ { required: true, message: '请输入field英文名', trigger: 'blur' }],
                fieldDesc:[ { required: true, message: '请输入field中文名', trigger: 'blur' }],
                fieldJson:[ { required: true, message: '请输入field格式化json', trigger: 'blur' }],
                columnName:[{ required: true, message: '请输入数据库字段', trigger: 'blur' }],
                fieldType:[ { required: true, message: '请选择field数据类型', trigger: 'blur' }]
            }

        }
    },
    methods: {
        async getDbfdOptions(){
            if(this.dialogMiddle.fieldType){
                 const res = await getDict(this.dialogMiddle.fieldType)
                 this.dbfdOptions = res
            }
        }
    },
    created() {
        this.getDbfdOptions()
    },
}
</script>
<style lang="scss">

</style>
