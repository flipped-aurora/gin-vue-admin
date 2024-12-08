<template>
  <div>
    <warning-bar
      title="id , created_at , updated_at , deleted_at 会自动生成请勿重复创建。搜索时如果条件为LIKE只支持字符串"
    />
    <el-form
      ref="fieldDialogForm"
      :model="middleDate"
      label-width="120px"
      label-position="right"
      :rules="rules"
      class="grid grid-cols-2"
    >
      <el-form-item label="字段名称" prop="fieldName">
        <el-input
          v-model="middleDate.fieldName"
          autocomplete="off"
          style="width: 80%"
        />
        <el-button style="width: 18%; margin-left: 2%" @click="autoFill">
          <span style="font-size: 12px">自动填充</span>
        </el-button>
      </el-form-item>
      <el-form-item label="字段中文名" prop="fieldDesc">
        <el-input v-model="middleDate.fieldDesc" autocomplete="off" />
      </el-form-item>
      <el-form-item label="字段JSON" prop="fieldJson">
        <el-input v-model="middleDate.fieldJson" autocomplete="off" />
      </el-form-item>
      <el-form-item label="数据库字段名" prop="columnName">
        <el-input v-model="middleDate.columnName" autocomplete="off" />
      </el-form-item>
      <el-form-item label="数据库字段描述" prop="comment">
        <el-input v-model="middleDate.comment" autocomplete="off" />
      </el-form-item>
      <el-form-item label="字段类型" prop="fieldType">
        <el-select
          v-model="middleDate.fieldType"
          style="width: 100%"
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
          :placeholder="
            middleDate.fieldType === 'enum'
              ? `例:'北京','天津'`
              : '数据库类型长度'
          "
        />
      </el-form-item>
      <el-form-item label="字段查询条件" prop="fieldSearchType">
        <el-select
          v-model="middleDate.fieldSearchType"
          :disabled="middleDate.fieldType === 'json'"
          style="width: 100%"
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
      <el-form-item label="关联字典" prop="dictType">
        <el-select
          v-model="middleDate.dictType"
          style="width: 100%"
          :disabled="middleDate.fieldType !== 'string' && middleDate.fieldType !== 'array'"
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
      <el-form-item label="默认值">
        <el-input
          v-model="middleDate.defaultValue"
          placeholder="请输入默认值"
        />
      </el-form-item>
      <el-form-item label="主键">
        <el-checkbox v-model="middleDate.primaryKey" />
      </el-form-item>
      <el-form-item label="索引类型" prop="fieldIndexType">
        <el-select
          v-model="middleDate.fieldIndexType"
          :disabled="middleDate.fieldType === 'json'"
          style="width: 100%"
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
      <el-form-item label="导入/导出">
        <el-switch v-model="middleDate.excel" />
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
        <el-switch
          v-model="middleDate.fieldSearchHide"
          :disabled="!middleDate.fieldSearchType"
        />
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
          <el-col :span="4">
            <el-select
              v-model="middleDate.dataSource.dbName"
              placeholder="数据库【不填则为GVA库】"
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
                  <span style="float: right; color: #8492a6; font-size: 13px">{{
                    item.dbName
                  }}</span>
                </div>
              </el-option>
            </el-select>
          </el-col>
          <el-col :span="4">
            <el-select
              v-model="middleDate.dataSource.association"
              placeholder="关联模式"
              @change="associationChange"
            >
              <el-option label="一对一" :value="1" />
              <el-option label="一对多" :value="2" />
            </el-select>
          </el-col>
          <el-col :span="5">
            <el-select
              v-model="middleDate.dataSource.table"
              placeholder="请选择数据源表"
              filterable
              allow-create
              clearable
              @focus="getDBTableList"
              @change="selectDB"
              @clear="clearAccress"
            >
              <el-option
                v-for="item in dbTableList"
                :key="item.tableName"
                :label="item.tableName"
                :value="item.tableName"
              />
            </el-select>
          </el-col>
          <el-col :span="5">
            <el-select
              v-model="middleDate.dataSource.value"
              placeholder="请先选择需要存储的数据"
            >
              <template #label="{ value }">
                <span>存储: </span>
                <span style="font-weight: bold">{{ value }}</span>
              </template>
              <el-option
                v-for="item in dbColumnList"
                :key="item.columnName"
                :value="item.columnName"
              >
                <span style="float: left">
                  <el-tag :type="item.isPrimary ? 'primary' : 'info'">
                    {{ item.isPrimary ? '主&emsp;键' : '非主键' }}
                  </el-tag>
                  {{ item.columnName }}</span
                >
                <span
                  style="
                    float: right;
                    margin-left: 5px;
                    color: var(--el-text-color-secondary);
                    font-size: 13px;
                  "
                >
                  类型：{{ item.type }}
                  <block v-if="item.comment != ''"
                    >，字段说明：{{ item.comment }}</block
                  >
                </span>
              </el-option>
            </el-select>
          </el-col>
          <el-col :span="5">
            <el-select
              v-model="middleDate.dataSource.label"
              placeholder="请先选择需要展示的数据"
            >
              <template #label="{ value }">
                <span>展示: </span>
                <span style="font-weight: bold">{{ value }}</span>
              </template>
              <el-option
                v-for="item in dbColumnList"
                :key="item.columnName"
                :value="item.columnName"
              >
                <span style="float: left">
                  <el-tag :type="item.isPrimary ? 'primary' : 'info'">
                    {{ item.isPrimary ? '主&emsp;键' : '非主键' }}
                  </el-tag>
                  {{ item.columnName }}</span
                >
                <span
                  style="
                    float: right;
                    margin-left: 5px;
                    color: var(--el-text-color-secondary);
                    font-size: 13px;
                  "
                >
                  类型：{{ item.type }}
                  <span v-if="item.comment != ''"
                    >，字段说明：{{ item.comment }}</span
                  >
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
  import { ref, onMounted } from 'vue'
  import { ElMessageBox } from 'element-plus'
  import { getColumn, getDB, getTable } from '@/api/autoCode'

  defineOptions({
    name: 'FieldDialog'
  })

  const props = defineProps({
    dialogMiddle: {
      type: Object,
      default: function () {
        return {}
      }
    },
    typeOptions: {
      type: Array,
      default: function () {
        return []
      }
    },
    typeSearchOptions: {
      type: Array,
      default: function () {
        return []
      }
    },
    typeIndexOptions: {
      type: Array,
      default: function () {
        return []
      }
    }
  })

  const activeNames = ref([])

  const middleDate = ref({})
  const dictOptions = ref([])

  const dbList = ref([])

  const getDbFunc = async () => {
    const res = await getDB()
    if (res.code === 0) {
      dbList.value = res.data.dbList
    }
  }

  const validateDataTypeLong = (rule, value, callback) => {
    const regex = /^('([^']*)'(?:,'([^']+)'*)*)$/
    if (middleDate.value.fieldType == 'enum' && !regex.test(value)) {
      callback(new Error('枚举值校验错误'))
    } else {
      callback()
    }
  }

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
    fieldType: [{ required: true, message: '请选择字段类型', trigger: 'blur' }],
    dataTypeLong: [{ validator: validateDataTypeLong, trigger: 'blur' }]
  })

  const init = async () => {
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

    if (
      fieldType !== 'int' &&
      fieldType !== 'time.Time' &&
      fieldType !== 'float64' &&
      (item === 'BETWEEN' || item === 'NOT BETWEEN')
    ) {
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
      )
        .then(() => {
          middleDate.value.fieldType = 'array'
        })
        .catch(() => {
          middleDate.value.dataSource.association = 1
        })
    }
  }

  const clearAccress = () => {
    middleDate.value.dataSource.value = ''
    middleDate.value.dataSource.label = ''
  }

  const clearDataSourceTable = () => {
    middleDate.value.dataSource.table = ''
  }

  const dbNameChange = () => {
    getDBTableList()
    clearDataSourceTable()
    clearAccress()
  }

  const dbTableList = ref([])

  const getDBTableList = async () => {
    const res = await getTable({
      businessDB: middleDate.value.dataSource.dbName
    })
    if (res.code === 0) {
      let list = res.data.tables // 确保这里正确获取到 tables 数组
      dbTableList.value = list.map((item) => ({
        tableName: item.tableName,
        value: item.tableName // 这里假设 value 也是 tableName，如果不同请调整
      }))
    }
    clearAccress()
  }

  const dbColumnList = ref([])
  const selectDB = async (val, isInit) => {
    middleDate.value.dataSource.hasDeletedAt = false
    middleDate.value.dataSource.table = val
    const res = await getColumn({
      businessDB: middleDate.value.dataSource.dbName,
      tableName: val
    })

    if (res.code === 0) {
      let list = res.data.columns // 确保这里正确获取到 tables 数组
      dbColumnList.value = list.map((item) => {
        if (item.columnName === 'deleted_at') {
          middleDate.value.dataSource.hasDeletedAt = true
        }
        return {
          columnName: item.columnName,
          value: item.columnName,
          type: item.dataType,
          isPrimary: item.primaryKey,
          comment: item.columnComment
        }
      })
      if (dbColumnList.value.length > 0 && !isInit) {
        middleDate.value.dataSource.label = dbColumnList.value[0].columnName
        middleDate.value.dataSource.value = dbColumnList.value[0].columnName
      }
    }
  }

  const fieldDialogForm = ref(null)
  defineExpose({ fieldDialogForm })

  onMounted(() => {
    getDbFunc()
    if (middleDate.value.dataSource.table) {
      selectDB(middleDate.value.dataSource.table, true)
    }
  })
</script>
