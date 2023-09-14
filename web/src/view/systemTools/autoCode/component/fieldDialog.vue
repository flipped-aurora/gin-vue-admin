<template>
  <div>
    <warning-bar title="id , created_at , updated_at , deleted_at 会自动生成请勿重复创建。搜索时如果条件为LIKE只支持字符串" />
    <el-form
      ref="fieldDialogFrom"
      :model="middleDate"
      label-width="120px"
      label-position="right"
      :rules="rules"
      class="grid grid-cols-2"
    >
      <el-form-item
        label="字段名称"
        prop="fieldName"
      >
        <el-input
          v-model="middleDate.fieldName"
          autocomplete="off"
          style="width:80%"
        />
        <el-button
          style="width:18%;margin-left:2%"
          @click="autoFill"
        >
          <span style="font-size: 12px">自动填充</span>
        </el-button>
      </el-form-item>
      <el-form-item
        label="字段中文名"
        prop="fieldDesc"
      >
        <el-input
          v-model="middleDate.fieldDesc"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        label="字段JSON"
        prop="fieldJson"
      >
        <el-input
          v-model="middleDate.fieldJson"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        label="数据库字段名"
        prop="columnName"
      >
        <el-input
          v-model="middleDate.columnName"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        label="数据库字段描述"
        prop="comment"
      >
        <el-input
          v-model="middleDate.comment"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        label="字段类型"
        prop="fieldType"
      >
        <el-select
          v-model="middleDate.fieldType"
          style="width:100%"
          placeholder="请选择字段类型"
          clearable
          @change="clearOther"
        >
          <el-option
            v-for="item in typeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        :label="middleDate.fieldType === 'enum' ? '枚举值' : '类型长度'"
        prop="dataTypeLong"
      >
        <el-input
          v-model="middleDate.dataTypeLong"
          :placeholder="middleDate.fieldType === 'enum'?`例:'北京','天津'`:'数据库类型长度'"
        />
      </el-form-item>
      <el-form-item
        label="字段查询条件"
        prop="fieldSearchType"
      >
        <el-select
          v-model="middleDate.fieldSearchType"
          style="width:100%"
          placeholder="请选择字段查询条件"
          clearable
        >
          <el-option
            v-for="item in typeSearchOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            :disabled="canSelect(item.value)"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="关联字典"
        prop="dictType"
      >
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

defineOptions({
  name: 'FieldDialog'
})

const props = defineProps({
  dialogMiddle: {
    type: Object,
    default: function() {
      return {}
    }
  },
  typeOptions: {
    type: Array,
    default: function() {
      return []
    }
  },
  typeSearchOptions: {
    type: Array,
    default: function() {
      return []
    }
  },
})

const middleDate = ref({})
const dictOptions = ref([])

const rules = ref({
  fieldName: [
    { required: true, message: '请输入字段英文名', trigger: 'blur' }
  ],
  fieldDesc: [
    { required: true, message: '请输入字段中文名', trigger: 'blur' }
  ],
  fieldJson: [
    { required: true, message: '请输入字段格式化json', trigger: 'blur' }
  ],
  columnName: [
    { required: true, message: '请输入数据库字段', trigger: 'blur' }
  ],
  fieldType: [
    { required: true, message: '请选择字段类型', trigger: 'blur' }
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

const canSelect = (item) => {
  const fieldType = middleDate.value.fieldType
  if (fieldType !== 'string' && item === 'LIKE') {
    return true
  }

  if ((fieldType !== 'int' && fieldType !== 'time.Time' && fieldType !== 'float64') && (item === 'BETWEEN' || item === 'NOT BETWEEN')) {
    return true
  }
  return false
}

const clearOther = () => {
  middleDate.value.fieldSearchType = ''
  middleDate.value.dictType = ''
}

const fieldDialogFrom = ref(null)
defineExpose({ fieldDialogFrom })
</script>
