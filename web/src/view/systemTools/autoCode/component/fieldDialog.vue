<template>
  <div>
    <warning-bar :title="t('fieldDialog.note')" />
    <el-form
      ref="fieldDialogFrom"
      :model="middleDate"
      label-width="160px"
      label-position="left"
      :rules="rules"
    >
      <el-form-item :label="t('autoCode.fieldName')" prop="fieldName">
        <el-input v-model="middleDate.fieldName" autocomplete="off" style="width:80%" />
        <el-button size="mini" style="width:18%;margin-left:2%" @click="autoFill">{{ t('fieldDialog.autoFill') }}</el-button>
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
          @change="getDbfdOptions"
        >
          <el-option
            v-for="item in typeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item :label="t('autoCode.fieldLen')" prop="dataTypeLong">
        <el-input v-model="middleDate.dataTypeLong" :placeholder="t('fieldDialog.dataTypeNote')" />
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
    </el-form>
  </div>
</template>

<script setup>

import { toLowerCase, toSQLLine } from '@/utils/stringFun'
import { getSysDictionaryList } from '@/api/sysDictionary'
import warningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

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
  }
])
const typeOptions = ref([
  {
    label: t('fieldDialog.string'),
    value: 'string'
  },
  {
    label: t('fieldDialog.integer'),
    value: 'int'
  },
  {
    label: t('fieldDialog.boolean'),
    value: 'bool'
  },
  {
    label: t('fieldDialog.float'),
    value: 'float64'
  },
  {
    label: t('fieldDialog.time'),
    value: 'time.Time'
  }
])
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

const fieldDialogFrom = ref(null)
defineExpose({ fieldDialogFrom })
</script>

<script>

export default {
  name: 'FieldDialog'
}
</script>
