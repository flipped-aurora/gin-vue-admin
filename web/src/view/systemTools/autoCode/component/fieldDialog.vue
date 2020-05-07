<template>
    <div>
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
            <el-form-item label="Field数据类型" prop="fieldType">
                <el-col :span="8">
            <el-select v-model="dialogMiddle.fieldType" placeholder="请选择field数据类型">
                <el-option
                    v-for="item in typeOptions"
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
            rules:{
                fieldName:[ { required: true, message: '请输入field英文名', trigger: 'blur' }],
                fieldDesc:[ { required: true, message: '请输入field中文名', trigger: 'blur' }],
                fieldJson:[ { required: true, message: '请输入field格式化json', trigger: 'blur' }],
                fieldType:[ { required: true, message: '请选择field数据类型', trigger: 'blur' }]
            }
            
        }
    }
}
</script>
<style lang="scss">
    
</style>