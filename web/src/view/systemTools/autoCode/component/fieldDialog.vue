<template>
  <div>
    <span style="color:red">搜索时如果条件为LIKE只支持字符串</span>
    <el-form
      ref="fieldDialogFrom"
      :model="middleDate"
      label-width="120px"
      label-position="left"
      :rules="rules"
    >
      <el-form-item label="Field名称" prop="fieldName">
        <el-col :span="6">
          <el-input v-model="middleDate.fieldName" autocomplete="off" />
        </el-col>
        <el-col :offset="1" :span="2">
          <el-button size="mini" @click="autoFill">自动填充</el-button>
        </el-col>
      </el-form-item>
      <el-form-item label="Field中文名" prop="fieldDesc">
        <el-col :span="6">
          <el-input v-model="middleDate.fieldDesc" autocomplete="off" />
        </el-col>
      </el-form-item>
      <el-form-item label="FieldJSON" prop="fieldJson">
        <el-col :span="6">
          <el-input v-model="middleDate.fieldJson" autocomplete="off" />
        </el-col>
      </el-form-item>
      <el-form-item label="数据库字段名" prop="columnName">
        <el-col :span="6">
          <el-input v-model="middleDate.columnName" autocomplete="off" />
        </el-col>
      </el-form-item>
      <el-form-item label="数据库字段描述" prop="comment">
        <el-col :span="6">
          <el-input v-model="middleDate.comment" autocomplete="off" />
        </el-col>
      </el-form-item>
      <el-form-item label="Field数据类型" prop="fieldType">
        <el-col :span="8">
          <el-select
            v-model="middleDate.fieldType"
            placeholder="请选择field数据类型"
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
        </el-col>
      </el-form-item>

      <el-form-item label="数据库字段类型" prop="dataType">
        <el-col :span="8">
          <el-select
            v-model="middleDate.dataType"
            :disabled="!middleDate.fieldType"
            placeholder="请选择数据库字段类型"
            clearable
          >
            <el-option
              v-for="item in dbfdOptions"
              :key="item.label"
              :label="item.label"
              :value="item.label"
            />
          </el-select>
        </el-col>
      </el-form-item>
      <el-form-item label="数据库字段长度" prop="dataTypeLong">
        <el-col :span="8">
          <el-input v-model="middleDate.dataTypeLong" placeholder="自定义类型必须指定长度" :disabled="!middleDate.dataType" />
        </el-col>
      </el-form-item>
      <el-form-item label="Field查询条件" prop="fieldSearchType">
        <el-col :span="8">
          <el-select v-model="middleDate.fieldSearchType" placeholder="请选择Field查询条件" clearable>
            <el-option
              v-for="item in typeSearchOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-col>
      </el-form-item>

      <el-form-item label="关联字典" prop="dictType">
        <el-col :span="8">
          <el-select v-model="middleDate.dictType" :disabled="middleDate.fieldType!=='int'" placeholder="请选择字典" clearable>
            <el-option
              v-for="item in dictOptions"
              :key="item.type"
              :label="`${item.type}(${item.name})`"
              :value="item.type"
            />
          </el-select>
        </el-col>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getDict } from '@/utils/dictionary'
import { toLowerCase, toSQLLine } from '@/utils/stringFun'
import { getSysDictionaryList } from '@/api/sysDictionary'

export default {
  name: 'FieldDialog',
  props: {
    dialogMiddle: {
      type: Object,
      default: function() {
        return {}
      }
    }
  },
  data() {
    return {
      middleDate: {},
      dbfdOptions: [],
      dictOptions: [],
      typeSearchOptions: [
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
      ],
      typeOptions: [
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
        }
      ],
      rules: {
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
      }
    }
  },
  async created() {
    this.middleDate = this.dialogMiddle
    const dictRes = await getSysDictionaryList({
      page: 1,
      pageSize: 999999
    })

    this.dictOptions = dictRes.data.list
  },
  methods: {
    autoFill() {
      this.middleDate.fieldJson = toLowerCase(this.middleDate.fieldName)
      this.middleDate.columnName = toSQLLine(this.middleDate.fieldJson)
    },
    async getDbfdOptions() {
      this.middleDate.dataType = ''
      this.middleDate.dataTypeLong = ''
      this.middleDate.fieldSearchType = ''
      this.middleDate.dictType = ''
      if (this.middleDate.fieldType) {
        this.dbfdOptions = await getDict(this.middleDate.fieldType)
      }
    }
  }
}
</script>
