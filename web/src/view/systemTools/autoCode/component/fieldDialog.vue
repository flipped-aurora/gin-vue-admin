<template>
  <div>
    <warning-bar :title="t('fieldDialog.note')" />
    <el-form
      ref="fieldDialogForm"
      :model="middleDate"
      label-width="120px"
      label-position="right"
      :rules="rules"
      class="grid grid-cols-2"
    >
      <el-form-item
        :label="t('autoCode.fieldName')"
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
          <span style="font-size: 12px">{{ t('fieldDialog.autoFill') }}</span>
        </el-button>
      </el-form-item>
      <el-form-item
        :label="t('autoCode.fieldDesc')"
        prop="fieldDesc"
      >
        <el-input
          v-model="middleDate.fieldDesc"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        :label="t('autoCode.fieldJson')"
        prop="fieldJson"
      >
        <el-input
          v-model="middleDate.fieldJson"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        :label="t('autoCode.columnName')"
        prop="columnName"
      >
        <el-input
          v-model="middleDate.columnName"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        :label="t('autoCode.comment')"
        prop="comment"
      >
        <el-input
          v-model="middleDate.comment"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        :label="t('autoCode.fieldDataType')"
        prop="fieldType"
      >
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
        :label="t('general.searchCriteria')"
        prop="fieldSearchType"
      >
        <el-select
          v-model="middleDate.fieldSearchType"
          :disabled="middleDate.fieldType === 'json'"
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
      <el-form-item
        :label="t('fieldDialog.associativeDictionary')"
        prop="dictType"
      >
        <el-select
          v-model="middleDate.dictType"
          style="width:100%"
          :disabled="middleDate.fieldType!=='string'"
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
      <el-form-item label="默认值">
        <el-input
          v-model="middleDate.defaultValue"
          placeholder="请输入默认值"
        />
      </el-form-item>
      <el-form-item label="主键">
        <el-checkbox v-model="middleDate.primaryKey" />
      </el-form-item>
      <el-form-item
        label="索引类型"
        prop="fieldIndexType"
      >
        <el-select
          v-model="middleDate.fieldIndexType"
          :disabled="middleDate.fieldType === 'json'"
          style="width:100%"
          placeholder="请选择字段索引类型"
          clearable
        >
          <el-option
            v-for="item in typeIndexOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            :disabled="canSelect(item.value)"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="前端新建/编辑">
        <el-switch v-model="middleDate.form" />
      </el-form-item>
      <el-form-item label="前端表格列">
        <el-switch v-model="middleDate.table" />
      </el-form-item>
      <el-form-item label="前端详情">
        <el-switch v-model="middleDate.desc" />
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
      <el-form-item label="隐藏查询条件">
        <el-switch v-model="middleDate.fieldSearchHide" :disabled="!middleDate.fieldSearchType" />
      </el-form-item>
      <el-form-item label="校验失败文案">
        <el-input v-model="middleDate.errorText" />
      </el-form-item>
    </el-form>
    <el-collapse v-model="activeNames">
      <el-collapse-item
        title="数据源配置（此配置为高级配置，如编程基础不牢，可能导致自动化代码不可用）"
        name="1"
      >
        <el-row :gutter="8">
          <el-col
            :span="3"
          >
            <el-select
              v-model="middleDate.dataSource.association"
              placeholder="关联模式"
              @change="associationChange"
            >
              <el-option
                label="一对一"
                :value="1"
              />
              <el-option
                label="一对多"
                :value="2"
              />
            </el-select>
          </el-col>

          <el-col :span="7">
            <el-select
              v-model="middleDate.dataSource.table" placeholder="请选择数据源表"
              filterable allow-create @focus="getDBTableList" @change="selectDB"
            >
              <el-option
                v-for="item in dbTableList" :key="item.tableName" :label="item.tableName"
                :value="item.tableName"
              />
            </el-select>
            <!-- <el-input v-model="middleDate.dataSource.table" placeholder="数据源表" /> -->
          </el-col>
          <el-col :span="7">
            <el-select v-model="middleDate.dataSource.value" placeholder="请先选择需要存储的数据">
              <template #label="{ value }">
                <span>存储: </span>
                <span style="font-weight: bold">{{ value }}</span>
              </template>
              <el-option v-for="item in dbColumnList" :key="item.columnName" :value="item.columnName">
                <span style="float: left"> <el-tag :type="item.isPrimary ? 'primary' : 'info'">
                  {{ item.isPrimary ? "主&emsp;键" : "非主键" }}
                </el-tag> {{ item.columnName }}</span>
                <span
                  style="
          float: right;
          margin-left:5px;
          color: var(--el-text-color-secondary);
          font-size: 13px;
        "
                >
                  类型：{{ item.type }} <block v-if="item.comment != ''">，字段说明：{{ item.comment }}</block>
                </span>
              </el-option>
            </el-select>
            <!-- <el-input v-model="middleDate.dataSource.value" placeholder="存储用字段" /> -->
          </el-col>
          <el-col :span="7">
            <el-select v-model="middleDate.dataSource.label" placeholder="请先选择需要展示的数据">
              <template #label="{ value }">
                <span>展示: </span>
                <span style="font-weight: bold">{{ value }}</span>
              </template>
              <el-option v-for="item in dbColumnList" :key="item.columnName" :value="item.columnName">
                <span style="float: left"> <el-tag :type="item.isPrimary ? 'primary' : 'info'">
                  {{ item.isPrimary ? "主&emsp;键" : "非主键" }}
                </el-tag> {{ item.columnName }}</span>
                <span
                    style="
          float: right;
          margin-left:5px;
          color: var(--el-text-color-secondary);
          font-size: 13px;
        "
                >
                  类型：{{ item.type }} <span v-if="item.comment != ''">，字段说明：{{ item.comment }}</span>
                </span>
              </el-option>
            </el-select>
            <!-- <el-input v-model="middleDate.dataSource.label" placeholder="展示用字段" /> -->
          </el-col>
        </el-row>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script setup>
import { toLowerCase, toSQLLine } from '@/utils/stringFun'
import { getSysDictionaryList } from '@/api/sysDictionary'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessageBox } from 'element-plus'
import {getColumn, getTable} from "@/api/autoCode";
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

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
  typeIndexOptions: {
    type: Array,
    default: function() {
      return []
    }
  },
})

const activeNames = ref([])

const middleDate = ref({})
const dictOptions = ref([])

const validateDataTypeLong = (rule, value, callback) => {
  const regex = /^('([^']*)'(?:,'([^']+)'*)*)$/;
  if (middleDate.value.fieldType == "enum" && !regex.test(value)) {
    callback(new Error("枚举值校验错误"));
  } else {
    callback();
  }
};

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
  ],
  dataTypeLong: [
    { validator: validateDataTypeLong, trigger: "blur" }
  ],
})

const init = async() => {
  middleDate.value = props.dialogMiddle
  const dictRes = await getSysDictionaryList({
    page: 1,
    pageSize: 999999
  })

  dictOptions.value = dictRes.data
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

const associationChange = (val) => {
  if (val === 2) {
    ElMessageBox.confirm(
      '一对多关联模式下，数据类型会改变为数组，后端表现为json，具体表现为数组模式，是否继续？',
      '提示',
      {
        confirmButtonText: '继续',
        cancelButtonText: '取消',
        type: 'warning'
      }
    ).then(() => {
      middleDate.value.fieldType = 'array'
    }).catch(() => {
      middleDate.value.dataSource.association = 1
    })
  }
}


const dbTableList = ref([])

const getDBTableList = async () => {
  const res = await getTable()
  console.log(res);
  if (res.code === 0) {
    let list = res.data.tables; // 确保这里正确获取到 tables 数组
    dbTableList.value = list.map(item => ({
      tableName: item.tableName,
      value: item.tableName // 这里假设 value 也是 tableName，如果不同请调整
    }));
  }
}

const dbColumnList = ref([])
const selectDB = async (val) => {
  middleDate.value.dataSource.table = val
  const res = await getColumn({
    tableName: val
  })
  console.log(res)
  if (res.code === 0) {
    let list = res.data.columns; // 确保这里正确获取到 tables 数组
    dbColumnList.value = list.map(item => ({
      columnName: item.columnName,
      value: item.columnName,
      type: item.dataType,
      isPrimary: item.primaryKey,
      comment: item.columnComment
    }));
    if (dbColumnList.value.length > 0) {
      middleDate.value.dataSource.label = dbColumnList.value[0].columnName
      middleDate.value.dataSource.value = dbColumnList.value[0].columnName
    }
  }
}


const fieldDialogForm = ref(null)
defineExpose({ fieldDialogForm })
</script>
