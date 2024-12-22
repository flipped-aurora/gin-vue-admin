<template>
  <div>
    <warning-bar :title="t('view.systemTools.autoCode.fieldDialog.note')" />
    <el-form
      ref="fieldDialogForm"
      :model="middleDate"
      label-width="180px"
      label-position="right"
      :rules="rules"
      class="grid grid-cols-2"
    >
      <el-form-item
        :label="t('view.systemTools.autoCode.fieldName')"
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
          <span style="font-size: 12px">{{ t('view.systemTools.autoCode.fieldDialog.autoFill') }}</span>
        </el-button>
      </el-form-item>
      <el-form-item
        :label="t('view.systemTools.autoCode.fieldDesc')"
        prop="fieldDesc"
      >
        <el-input
          v-model="middleDate.fieldDesc"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        :label="t('view.systemTools.autoCode.fieldJson')"
        prop="fieldJson"
      >
        <el-input
          v-model="middleDate.fieldJson"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        :label="t('view.systemTools.autoCode.columnName')"
        prop="columnName"
      >
        <el-input
          v-model="middleDate.columnName"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        :label="t('view.systemTools.autoCode.comment')"
        prop="comment"
      >
        <el-input
          v-model="middleDate.comment"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        :label="t('view.systemTools.autoCode.fieldDataType')"
        prop="fieldType"
      >
        <el-select
          v-model="middleDate.fieldType"
          style="width:100%"
          :placeholder="t('view.systemTools.autoCode.fieldDialog.selectDataType')"
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
        :label="middleDate.fieldType === 'enum' ? t('view.systemTools.autoCode.fieldDialog.enumValue') : t('view.systemTools.autoCode.fieldDialog.typeLength')"
        prop="dataTypeLong"
      >
        <el-input
          v-model="middleDate.dataTypeLong"
          :placeholder="middleDate.fieldType === 'enum'? t('view.systemTools.autoCode.fieldDialog.enumExample') : t('view.systemTools.autoCode.fieldDialog.dataTypeLength')"
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
          :placeholder="t('view.systemTools.autoCode.fieldDialog.selectSearchType')"
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
        :label="t('view.systemTools.autoCode.fieldDialog.associativeDictionary')"
        prop="dictType"
      >
        <el-select
          v-model="middleDate.dictType"
          style="width:100%"
          :disabled="middleDate.fieldType!=='string'"
          :placeholder="t('view.systemTools.autoCode.fieldDialog.selectDictionary')"
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
      <el-form-item :label="t('view.systemTools.autoCode.defaultValue')">
        <el-input
          v-model="middleDate.defaultValue"
          :placeholder="t('view.systemTools.autoCode.fieldDialog.enterDefaultValueNote')"
        />
      </el-form-item>
      <el-form-item :label="t('view.systemTools.autoCode.primaryKey')">
        <el-checkbox v-model="middleDate.primaryKey" />
      </el-form-item>
      <el-form-item
        :label="t('view.systemTools.autoCode.indexType')"
        prop="fieldIndexType"
      >
        <el-select
          v-model="middleDate.fieldIndexType"
          :disabled="middleDate.fieldType === 'json'"
          style="width:100%"
          :placeholder="t('view.systemTools.autoCode.selectIndexType')"
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
      <el-form-item :label="t('view.systemTools.autoCode.fieldDialog.frontendCreateEdit')">
        <el-switch v-model="middleDate.form" />
      </el-form-item>
      <el-form-item :label="t('view.systemTools.autoCode.fieldDialog.frontendTableColmuns')">
        <el-switch v-model="middleDate.table" />
      </el-form-item>
      <el-form-item :label="t('view.systemTools.autoCode.fieldDialog.frontendDetails')">
        <el-switch v-model="middleDate.desc" />
      </el-form-item>
      <el-form-item :label="t('view.systemTools.autoCode.importExport')">
        <el-switch v-model="middleDate.excel" />
      </el-form-item>
      <el-form-item :label="t('view.systemTools.autoCode.fieldDialog.sort')">
        <el-switch v-model="middleDate.sort" />
      </el-form-item>
      <el-form-item :label="t('view.systemTools.autoCode.fieldDialog.required')">
        <el-switch v-model="middleDate.require" />
      </el-form-item>
      <el-form-item :label="t('view.systemTools.autoCode.fieldDialog.canBeCleared')">
        <el-switch v-model="middleDate.clearable" />
      </el-form-item>
      <el-form-item :label="t('view.systemTools.autoCode.fieldDialog.hideSearch')">
        <el-switch v-model="middleDate.fieldSearchHide" :disabled="!middleDate.fieldSearchType" />
      </el-form-item>
      <el-form-item :label="t('view.systemTools.autoCode.fieldDialog.verificationError')">
        <el-input v-model="middleDate.errorText" />
      </el-form-item>
    </el-form>
    <el-collapse v-model="activeNames">
      <el-collapse-item
        :title="t('view.systemTools.autoCode.fieldDialog.dataSourceConfigNote')"
        name="1"
      >
        <el-row :gutter="8">
          <el-col
              :span="4"
          >
            <el-select
                v-model="middleDate.dataSource.dbName"
                :placeholder="t('view.systemTools.autoCode.fieldDialog.dataSourceNameNote')"
                @change="dbNameChange"
                clearable
            >
              <el-option
                  v-for="item in dbList"
                  :key="item.aliasName"
                  :value="item.aliasName"
                  :label="item.aliasName"
                  :disabled="item.disable"
              >
                <div>
                  <span>{{ item.aliasName }}</span>
                  <span style="float:right;color:#8492a6;font-size:13px">{{ item.dbName }}</span>
                </div>
              </el-option>
            </el-select>
          </el-col>
          <el-col
            :span="4"
          >
            <el-select
              v-model="middleDate.dataSource.association"
              :placeholder="t('view.systemTools.autoCode.fieldDialog.associationMode')"
              @change="associationChange"
            >
              <el-option
                :label="t('view.systemTools.autoCode.fieldDialog.oneToOne')"
                :value="1"
              />
              <el-option
                :label="t('view.systemTools.autoCode.fieldDialog.oneToMany')"
                :value="2"
              />
            </el-select>
          </el-col>
          <el-col :span="5">
            <el-select
              v-model="middleDate.dataSource.table" :placeholder="t('view.systemTools.autoCode.fieldDialog.selectDataSourceTable')"
              filterable allow-create @focus="getDBTableList" @change="selectDB"
            >
              <el-option
                v-for="item in dbTableList" :key="item.tableName" :label="item.tableName"
                :value="item.tableName"
              />
            </el-select>
          </el-col>
          <el-col :span="5">
            <el-select v-model="middleDate.dataSource.value" :placeholder="t('view.systemTools.autoCode.fieldDialog.selectDataToStore')">
              <template #label="{ value }">
                <span>{{ t('view.systemTools.autoCode.fieldDialog.storage') }}</span>
                <span style="font-weight: bold">{{ value }}</span>
              </template>
              <el-option v-for="item in dbColumnList" :key="item.columnName" :value="item.columnName">
                <span style="float: left"> <el-tag :type="item.isPrimary ? 'primary' : 'info'">
                  {{ item.isPrimary ? t('view.systemTools.autoCode.primaryKey') : t('view.systemTools.autoCode.nonPrimaryKey') }}
                </el-tag> {{ item.columnName }}</span>
                <span
                  style="
          float: right;
          margin-left:5px;
          color: var(--el-text-color-secondary);
          font-size: 13px;
        "
                >
                  {{ t('view.systemTools.autoCode.fieldDialog.type') }}{{ item.type }} <block v-if="item.comment != ''">{{ t('view.systemTools.autoCode.fieldDialog.fileDesc') }}{{ item.comment }}</block>
                </span>
              </el-option>
            </el-select>
          </el-col>
          <el-col :span="5">
            <el-select v-model="middleDate.dataSource.label" :placeholder="t('view.systemTools.autoCode.fieldDialog.selectDataToDisplay')">
              <template #label="{ value }">
                <span>{{ t('view.systemTools.autoCode.fieldDialog.display') }}</span>
                <span style="font-weight: bold">{{ value }}</span>
              </template>
              <el-option v-for="item in dbColumnList" :key="item.columnName" :value="item.columnName">
                <span style="float: left"> <el-tag :type="item.isPrimary ? 'primary' : 'info'">
                  {{ item.isPrimary ? t('view.systemTools.autoCode.primaryKey') : t('view.systemTools.autoCode.nonPrimaryKey') }}
                </el-tag> {{ item.columnName }}</span>
                <span
                    style="
          float: right;
          margin-left:5px;
          color: var(--el-text-color-secondary);
          font-size: 13px;
        "
                >
                {{ t('view.systemTools.autoCode.fieldDialog.type') }}{{ item.type }} <span v-if="item.comment != ''">{{ t('view.systemTools.autoCode.fieldDialog.fileDesc') }}{{ item.comment }}</span>
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
import { ref,onMounted } from 'vue'
import { ElMessageBox } from 'element-plus'
import {getColumn, getDB, getTable} from "@/api/autoCode";
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
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

const dbList = ref([])

const getDbFunc = async() => {
  const res = await getDB()
  if (res.code === 0) {
    dbList.value = res.data.dbList
  }
}

const validateDataTypeLong = (rule, value, callback) => {
  const regex = /^('([^']*)'(?:,'([^']+)'*)*)$/;
  if (middleDate.value.fieldType == "enum" && !regex.test(value)) {
    callback(new Error(t('view.systemTools.autoCode.fieldDialog.enumValueValidationError')));
  } else {
    callback();
  }
};

const rules = ref({
  fieldName: [
    { required: true, message: t('view.systemTools.autoCode.fieldDialog.entFieldName'), trigger: 'blur' }
  ],
  fieldDesc: [
    { required: true, message: t('view.systemTools.autoCode.fieldDialog.entFieldDesc'), trigger: 'blur' }
  ],
  fieldJson: [
    { required: true, message: t('view.systemTools.autoCode.fieldDialog.entFieldJson'), trigger: 'blur' }
  ],
  columnName: [
    { required: true, message: t('view.systemTools.autoCode.fieldDialog.entColumnName'), trigger: 'blur' }
  ],
  fieldType: [
    { required: true, message: t('view.systemTools.autoCode.fieldDialog.entFieldDataType'), trigger: 'blur' }
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
      t('view.systemTools.autoCode.fieldDialog.oneToManyNote'),
      t('general.hint'),
      {
        confirmButtonText: t('general.confirm'),
        cancelButtonText: t('general.cancel'),
        type: 'warning'
      }
    ).then(() => {
      middleDate.value.fieldType = 'array'
    }).catch(() => {
      middleDate.value.dataSource.association = 1
    })
  }
}


const dbNameChange = () => {
  getDBTableList()
  middleDate.value.dataSource.table = ''
  middleDate.value.dataSource.value = ''
  middleDate.value.dataSource.label = ''
}


const dbTableList = ref([])

const getDBTableList = async () => {
  const res = await getTable({ businessDB: middleDate.value.dataSource.dbName })
  if (res.code === 0) {
    let list = res.data.tables; // 确保这里正确获取到 tables 数组
    dbTableList.value = list.map(item => ({
      tableName: item.tableName,
      value: item.tableName // 这里假设 value 也是 tableName，如果不同请调整
    }));
  }
  middleDate.value.dataSource.value = ''
  middleDate.value.dataSource.label = ''
}

const dbColumnList = ref([])
const selectDB = async (val,isInit) => {
  middleDate.value.dataSource.hasDeletedAt = false
  middleDate.value.dataSource.table = val
  const res = await getColumn({
    businessDB: middleDate.value.dataSource.dbName,
    tableName: val
  })

  if (res.code === 0) {
    let list = res.data.columns; // 确保这里正确获取到 tables 数组
    dbColumnList.value = list.map(item => {
        if(item.columnName === 'deleted_at'){
          middleDate.value.dataSource.hasDeletedAt = true
        }
        return{
            columnName: item.columnName,
            value: item.columnName,
            type: item.dataType,
            isPrimary: item.primaryKey,
            comment: item.columnComment
        }
    });
    if (dbColumnList.value.length > 0 && !isInit) {
      middleDate.value.dataSource.label = dbColumnList.value[0].columnName
      middleDate.value.dataSource.value = dbColumnList.value[0].columnName
    }
  }
}


const fieldDialogForm = ref(null)
defineExpose({ fieldDialogForm })

onMounted(()=>{
  getDbFunc()
  if(middleDate.value.dataSource.table){
    selectDB(middleDate.value.dataSource.table,true)
  }
})
</script>
