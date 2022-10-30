<template>
  <div>
    <warning-bar title="id , created_at , updated_at , deleted_at 会自动生成请勿重复创建。搜索时如果条件为LIKE只支持字符串" />
    <el-form
      ref="fieldDialogFrom"
      :model="middleDate"
      label-width="120px"
      label-position="right"
      :rules="rules"
      class="grid-form"
    >
      <el-form-item label="Field名称" prop="fieldName">
        <el-input v-model="middleDate.fieldName" autocomplete="off" style="width:80%" />
        <el-button size="small" style="width:18%;margin-left:2%" @click="autoFill">
          <span style="font-size: 12px">自动填充</span>
        </el-button>
      </el-form-item>
      <el-form-item label="Field中文名" prop="fieldDesc">
        <el-input v-model="middleDate.fieldDesc" autocomplete="off" />
      </el-form-item>
      <el-form-item label="FieldJSON" prop="fieldJson">
        <el-input v-model="middleDate.fieldJson" autocomplete="off" />
      </el-form-item>
      <el-form-item label="数据库字段名" prop="columnName">
        <el-input v-model="middleDate.columnName" autocomplete="off" />
      </el-form-item>
      <el-form-item label="数据库字段描述" prop="comment">
        <el-input v-model="middleDate.comment" autocomplete="off" />
      </el-form-item>
      <el-form-item label="Field数据类型" prop="fieldType">
        <el-select
          v-model="middleDate.fieldType"
          style="width:100%"
          placeholder="请选择field数据类型"
          clearable
          @change="clearOther"
        >
          <el-option
            v-for="item in typeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item :label="middleDate.fieldType === 'enum' ? '枚举值' : '类型长度'" prop="dataTypeLong">
        <el-input v-model="middleDate.dataTypeLong" :placeholder="middleDate.fieldType === 'enum'?`例:'北京','天津'`:'数据库类型长度'" />
      </el-form-item>
      <el-form-item label="Field查询条件" prop="fieldSearchType">
        <el-select
          v-model="middleDate.fieldSearchType"
          style="width:100%"
          placeholder="请选择Field查询条件"
          clearable
        >
          <el-option
            v-for="item in typeSearchOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            :disabled="
              (middleDate.fieldType!=='string'&&item.value==='LIKE')||
                ((middleDate.fieldType!=='int'&&middleDate.fieldType!=='time.Time'&&middleDate.fieldType!=='float64')&&(item.value==='BETWEEN' || item.value==='NOT BETWEEN'))
            "
          />
        </el-select>
      </el-form-item>
      <el-form-item label="关联字典" prop="dictType">
        <el-select
          v-model="middleDate.dictType"
          style="width:100%"
          :disabled="middleDate.fieldType!=='int'"
          placeholder="请选择字典"
          clearable
        >
          <el-option
            v-for="item in dictOptions"
            :key="item.type"
            :label="`${item.type}(${item.name})`"
            :value="item.type"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="是否排序">
        <el-switch v-model="middleDate.sort" />
      </el-form-item>
      <el-form-item label="是否必填">
        <el-switch v-model="middleDate.require" />
      </el-form-item>
      <el-form-item label="是否可清空">
        <el-switch v-model="middleDate.clearable" />
      </el-form-item>
      <el-form-item label="校验失败文案">
        <el-input v-model="middleDate.errorText" />
      </el-form-item>

    </el-form>
  </div>
</template>

<script setup>
import { toLowerCase, toSQLLine } from '@/utils/stringFun'
import { getSysDictionaryList } from '@/api/sysDictionary'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'

const props = defineProps({
  dialogMiddle: {
    type: Object,
    default: function() {
      return {}
    }
  }
})

const middleDate = ref({})
const dictOptions = ref([])
const typeSearchOptions = ref([
  {
    label: '=',
    value: '='
  },
  {
    label: '<>',
    value: '<>'
  },
  {
    label: '>',
    value: '>'
  },
  {
    label: '<',
    value: '<'
  },
  {
    label: 'LIKE',
    value: 'LIKE'
  },
  {
    label: 'BETWEEN',
    value: 'BETWEEN'
  },
  {
    label: 'NOT BETWEEN',
    value: 'NOT BETWEEN'
  }
])
const typeOptions = ref([
  {
    label: '字符串',
    value: 'string'
  },
  {
    label: '整型',
    value: 'int'
  },
  {
    label: '布尔值',
    value: 'bool'
  },
  {
    label: '浮点型',
    value: 'float64'
  },
  {
    label: '时间',
    value: 'time.Time'
  },
  {
    label: '枚举',
    value: 'enum'
  }
])
const rules = ref({
  fieldName: [
    { required: true, message: '请输入field英文名', trigger: 'blur' }
  ],
  fieldDesc: [
    { required: true, message: '请输入field中文名', trigger: 'blur' }
  ],
  fieldJson: [
    { required: true, message: '请输入field格式化json', trigger: 'blur' }
  ],
  columnName: [
    { required: true, message: '请输入数据库字段', trigger: 'blur' }
  ],
  fieldType: [
    { required: true, message: '请选择field数据类型', trigger: 'blur' }
  ]
})

const init = async() => {
  middleDate.value = props.dialogMiddle
  const dictRes = await getSysDictionaryList({
    page: 1,
    pageSize: 999999
  })

  dictOptions.value = dictRes.data.list
}
init()

const autoFill = () => {
  middleDate.value.fieldJson = toLowerCase(middleDate.value.fieldName)
  middleDate.value.columnName = toSQLLine(middleDate.value.fieldJson)
}

const clearOther = () => {
  middleDate.value.fieldSearchType = ''
  middleDate.value.dictType = ''
}

const fieldDialogFrom = ref(null)
defineExpose({ fieldDialogFrom })
</script>

<script>

export default {
  name: 'FieldDialog'
}
</script>
<style scoped>
.grid-form{
  display: grid;
  grid-template-columns: 1fr 1fr;
}
.click-text{
  color: #0d84ff;
  font-size: 13px;
  cursor: pointer;
  user-select: none;
}
</style>
