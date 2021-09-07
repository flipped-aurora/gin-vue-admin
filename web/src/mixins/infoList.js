import { getDict } from '@/utils/dictionary'
import { formatTimeToStr } from '@/utils/date'

export default {
  data() {
    return {
      page: 1,
      total: 10,
      pageSize: 10,
      tableData: [],
      searchInfo: {}
    }
  },
  methods: {
    formatBoolean: function(bool) {
      if (bool !== null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    },
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    filterDict(value, type) {
      const rowLabel = this[type + 'Options'] && this[type + 'Options'].filter(item => item.value === value)
      return rowLabel && rowLabel[0] && rowLabel[0].label
    },
    async getDict(type) {
      const dicts = await getDict(type)
      this[type + 'Options'] = dicts
      return dicts
    },
    handleSizeChange(val) {
      this.pageSize = val
      this.getTableData()
    },
    handleCurrentChange(val) {
      this.page = val
      this.getTableData()
    },
    // @params beforeFunc function 请求发起前执行的函数 默认为空函数
    // @params afterFunc function 请求完成后执行的函数 默认为空函数
    async getTableData(beforeFunc = () => {}, afterFunc = () => {}) {
      beforeFunc()
      const table = await this.listApi({ page: this.page, pageSize: this.pageSize, ...this.searchInfo })
      if (table.code === 0) {
        this.tableData = table.data.list
        this.total = table.data.total
        this.page = table.data.page
        this.pageSize = table.data.pageSize
      }
      afterFunc()
    }
  }
}
