// table 纯前端示例
<template>
  <div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column label="日期" width="120">
        <template slot-scope="scope">{{ scope.row.date }}</template>
      </el-table-column>
      <el-table-column label="姓名" prop="name" width="120"></el-table-column>
      <el-table-column label="年龄" prop="age" width="120"></el-table-column>
      <el-table-column label="住址" prop="address" min-width="200" show-overflow-tooltip></el-table-column>
      <el-table-column label="是否禁用" prop="switch" width="180">
        <template slot-scope="scope">
          <el-switch active-text="开启" inactive-text="禁用" v-model="scope.row.switch"></el-switch>
        </template>
      </el-table-column>
      <el-table-column label="按钮组" width="200">
        <template slot-scope="scope" >
            <el-button type="text" size="small" @click="toggleSelection([scope.row])">按钮1</el-button>
            <el-button type="text" size="small" @click="toggleSelection([scope.row])">按钮2</el-button>
            <el-button type="text" size="small" @click="toggleSelection([scope.row])">按钮3</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div style="margin-top: 20px">
      <el-button @click="toggleSelection([tableData[1], tableData[2]])">切换第二、第三行的选中状态</el-button>
      <el-button @click="toggleSelection()">取消选择</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Table',
  data() {
    return {
      tableData: [
        {
          date: '2016-05-03',
          name: '王小虎',
          age: 12,
          address: '上海市普陀区金沙江路 1518 弄',
          switch:true
        },
        {
          date: '2016-05-02',
          name: '王小虎',
          age: 12,
          address: '上海市普陀区金沙江路 1518 弄',
          switch:true
        },
        {
          date: '2016-05-04',
          name: '王小虎',
          age: 12,
          address: '上海市普陀区金沙江路 1518 弄',
          switch:true
        },
        {
          date: '2016-05-01',
          name: '王小虎',
          age: 12,
          address: '上海市普陀区金沙江路 1518 弄',
          switch:false
        },
        {
          date: '2016-05-08',
          name: '王小虎',
          age: 12,
          address: '上海市普陀区金沙江路 1518 弄',
          switch:true
        },
        {
          date: '2016-05-06',
          name: '王小虎',
          age: 12,
          address: '上海市普陀区金沙江路 1518 弄',
          switch:true
        },
        {
          date: '2016-05-07',
          name: '王小虎',
          age: 12,
          address: '上海市普陀区金沙江路 1518 弄',
          switch:false
        }
      ],
      multipleSelection: []
    }
  },

  methods: {
    toggleSelection(rows) {
      if (rows) {
        rows.forEach(row => {
          this.$refs.multipleTable.toggleRowSelection(row)
        })
      } else {
        this.$refs.multipleTable.clearSelection()
      }
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    }
  }
}
</script>