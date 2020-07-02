<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="请求方法">
          <el-input placeholder="搜索条件" v-model="searchInfo.method"></el-input>
        </el-form-item>
        <el-form-item label="请求路径">
          <el-input placeholder="搜索条件" v-model="searchInfo.path"></el-input>
        </el-form-item>
        <el-form-item label="结果状态码">
          <el-input placeholder="搜索条件" v-model="searchInfo.status"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" class="table-expand">
            <el-form-item label="请求ip">
              <span>{{ props.row.ip }}</span>
            </el-form-item>
            <el-form-item label="请求日期">
              <span>{{props.row.CreatedAt|formatDate}}</span>
            </el-form-item>
            <el-form-item label="请求方法">
              <span>{{ props.row.method }}</span>
            </el-form-item>
            <el-form-item label="请求路径">
              <span>{{ props.row.path }}</span>
            </el-form-item>
            <el-form-item label="结果状态码">
              <span>{{ props.row.status }}</span>
            </el-form-item>
            <el-form-item label="latency">
              <span>{{ props.row.latency }}</span>
            </el-form-item>
            <el-form-item label="agent">
              <span>{{ props.row.agent }}</span>
            </el-form-item>
            <el-form-item label="error_message">
              <span>{{ props.row.error_message }}</span>
            </el-form-item>
            <el-form-item label="user_id">
              <span>{{ props.row.user_id }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column type="selection" width="55"></el-table-column>

      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>

      <el-table-column label="请求ip" prop="ip" width="120"></el-table-column>

      <el-table-column label="请求方法" prop="method" width="120"></el-table-column>

      <el-table-column label="请求路径" prop="path" width="240"></el-table-column>

      <el-table-column label="error_message" prop="error_message" width="120"></el-table-column>

      <el-table-column label="user_id" prop="user_id" width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteSysOperationRecord(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>

<script>
import {
  deleteSysOperationRecord,
  getSysOperationRecordList
} from "@/api/sysOperationRecord"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "SysOperationRecord",
  mixins: [infoList],
  data() {
    return {
      listApi: getSysOperationRecordList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      formData: {
        ip: null,
        method: null,
        path: null,
        status: null,
        latency: null,
        agent: null,
        error_message: null,
        user_id: null
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    }
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    async deleteSysOperationRecord(row) {
      this.visible = false;
      const res = await deleteSysOperationRecord({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    }
  },
  created() {
    this.getTableData();
  }
};
</script>

<style lang="scss">
.table-expand {
  padding-left: 60px;
  font-size: 0;
  label {
    width: 90px;
    color: #99a9bf;
    .el-form-item {
      margin-right: 0;
      margin-bottom: 0;
      width: 50%;
    }
  }
}
</style>