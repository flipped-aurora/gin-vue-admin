<template>
  <div>
    <warning-bar :title="t('fieldDialog.note')" />
    <el-form
      ref="fieldDialogFrom"
      :model="middleDate"
      label-width="120px"
      label-position="right"
      :rules="rules"
      class="grid-form"
    >
      <el-form-item :label="t('autoCode.fieldName')" prop="fieldName">
        <el-input v-model="middleDate.fieldName" autocomplete="off" style="width:80%" />
        <el-button style="width:18%;margin-left:2%" @click="autoFill">
          <span style="font-size: 12px">{{ t('fieldDialog.autoFill') }}</span>
        </el-button>
      </el-form-item>
      <el-form-item :label="t('autoCode.fieldDesc')" prop="fieldDesc">
        <el-input v-model="middleDate.fieldDesc" autocomplete="off" />
      </el-form-item>
      <el-form-item :label="t('autoCode.fieldJson')" prop="fieldJson">
        <el-input v-model="middleDate.fieldJson" autocomplete="off" />
      </el-form-item>
      <el-form-item :label="t('autoCode.columnName')" prop="columnName">
        <el-input v-model="middleDate.columnName" autocomplete="off" />
      </el-form-item>
      <el-form-item :label="t('autoCode.comment')" prop="comment">
        <el-input v-model="middleDate.comment" autocomplete="off" />
      </el-form-item>
      <el-form-item :label="t('autoCode.fieldDataType')" prop="fieldType">
        <el-select
          v-model="middleDate.fieldType"
          style="width:100%"
          :placeholder="t('fieldDialog.selectDataType')"
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
      <el-form-item :label="middleDate.fieldType === 'enum' ? '枚举值' : '类型长度'" prop="dataTypeLong">
        <el-input v-model="middleDate.dataTypeLong" :placeholder="middleDate.fieldType === 'enum'?`例:'北京','天津'`:'数据库类型长度'" />
      </el-form-item>
      <el-form-item :label="t('general.searchCriteria')" prop="fieldSearchType">
        <el-select
          v-model="middleDate.fieldSearchType"
          style="width:100%"
          :placeholder="t('fieldDialog.selectSearchType')"
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

      <el-form-item :label="t('fieldDialog.associativeDictionary')" prop="dictType">
        <el-select
          v-model="middleDate.dictType"
          style="width:100%"
          :disabled="middleDate.fieldType!=='int'"
          :placeholder="t('fieldDialog.selectDictionary')"
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
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

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
    { required: true, message: t('fieldDialog.entFieldName'), trigger: 'blur' }
  ],
  fieldDesc: [
    { required: true, message: t('fieldDialog.entFieldDesc'), trigger: 'blur' }
  ],
  fieldJson: [
    { required: true, message: t('fieldDialog.entFieldJson'), trigger: 'blur' }
  ],
  columnName: [
    { required: true, message: t('fieldDialog.entColumnName'), trigger: 'blur' }
  ],
  fieldType: [
    { required: true, message: t('fieldDialog.entFieldDataType'), trigger: 'blur' }
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
</style>
