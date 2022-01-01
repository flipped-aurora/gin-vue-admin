<template>
  <div>
    <warning-bar :title="$t('fieldDialog.note')" />
    <el-form
      ref="fieldDialogFrom"
      :model="middleDate"
      label-width="160px"
      label-position="left"
      :rules="rules"
    >
      <el-form-item :label="$t('autoCode.fieldName')" prop="fieldName">
        <el-input v-model="middleDate.fieldName" autocomplete="off" style="width:80%" />
        <el-button size="mini" style="width:18%;margin-left:2%" @click="autoFill">{{ $t('fieldDialog.autoFill') }}</el-button>
      </el-form-item>
      <el-form-item :label="$t('autoCode.fieldDesc')" prop="fieldDesc">
        <el-input v-model="middleDate.fieldDesc" autocomplete="off" />
      </el-form-item>
      <el-form-item :label="$t('autoCode.fieldJson')" prop="fieldJson">
        <el-input v-model="middleDate.fieldJson" autocomplete="off" />
      </el-form-item>
      <el-form-item :label="$t('autoCode.columnName')" prop="columnName">
        <el-input v-model="middleDate.columnName" autocomplete="off" />
      </el-form-item>
      <el-form-item :label="$t('autoCode.comment')" prop="comment">
        <el-input v-model="middleDate.comment" autocomplete="off" />
      </el-form-item>
      <el-form-item :label="$t('autoCode.fieldDataType')" prop="fieldType">
        <el-select
          v-model="middleDate.fieldType"
          style="width:100%"
          :placeholder="$t('fieldDialog.selectDataType')"
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
      <el-form-item :label="$t('autoCode.fieldLen')" prop="dataTypeLong">
        <el-input v-model="middleDate.dataTypeLong" :placeholder="$t('fieldDialog.dataTypeNote')" />
      </el-form-item>
      <el-form-item :label="$t('autoCode.searchType')" prop="fieldSearchType">
        <el-select
          v-model="middleDate.fieldSearchType"
          style="width:100%"
          :placeholder="$t('fieldDialog.selectSearchType')"
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

      <el-form-item :label="$t('fieldDialog.associativeDictionary')" prop="dictType">
        <el-select
          v-model="middleDate.dictType"
          style="width:100%"
          :disabled="middleDate.fieldType!=='int'"
          :placeholder="$t('fieldDialog.selectDictionary')"
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

<script>
import { toLowerCase, toSQLLine } from '@/utils/stringFun'
import { getSysDictionaryList } from '@/api/sysDictionary'
import warningBar from '@/components/warningBar/warningBar.vue'

export default {
  name: 'FieldDialog',
  components: { warningBar },
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
          label: this.$t('fieldDialog.string'),
          value: 'string'
        },
        {
          label: this.$t('fieldDialog.integer'),
          value: 'int'
        },
        {
          label: this.$t('fieldDialog.boolean'),
          value: 'bool'
        },
        {
          label: this.$t('fieldDialog.float'),
          value: 'float64'
        },
        {
          label: this.$t('fieldDialog.time'),
          value: 'time.Time'
        }
      ],
      rules: {
        fieldName: [
          { required: true, message: this.$t('fieldDialog.entFieldName'), trigger: 'blur' }
        ],
        fieldDesc: [
          { required: true, message: this.$t('fieldDialog.entFieldDesc'), trigger: 'blur' }
        ],
        fieldJson: [
          { required: true, message: this.$t('fieldDialog.entFieldJson'), trigger: 'blur' }
        ],
        columnName: [
          { required: true, message: this.$t('fieldDialog.entColumnName'), trigger: 'blur' }
        ],
        fieldType: [
          { required: true, message: this.$t('fieldDialog.entFieldDataType'), trigger: 'blur' }
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
  }
}
</script>
