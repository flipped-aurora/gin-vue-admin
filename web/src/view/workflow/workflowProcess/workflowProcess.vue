<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="流程名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="流程标题">
          <el-input placeholder="搜索条件" v-model="searchInfo.label"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
              <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
            </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
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
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="流程名称" prop="name" width="120"></el-table-column>

      <el-table-column label="分类" prop="category" width="120"></el-table-column>

      <el-table-column label="类型" prop="clazz" width="120"></el-table-column>

      <el-table-column label="流程标题" prop="label" width="120"></el-table-column>

      <el-table-column label="是否隐藏图标" prop="hideIcon" width="120">
        <template slot-scope="scope">
          {{
          scope.row.hideIcon | formatBoolean
          }}
        </template>
      </el-table-column>

      <el-table-column label="详细介绍" prop="description" width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="useWorkflowProcess(scope.row)" size="success">启动</el-button>
          <el-button
            class="table-button"
            @click="updateWorkflowProcess(scope.row)"
            size="small"
            type="primary"
          >变更</el-button>
          <el-button
            class="table-button"
            @click="viewWorkflowProcess(scope.row)"
            size="small"
            type="warning"
          >查看</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteWorkflowProcess(scope.row)">确定</el-button>
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
      :style="{ float: 'right', padding: '20px' }"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>

<script>
import {
  deleteWorkflowProcess,
  deleteWorkflowProcessByIds,
  getWorkflowProcessList
} from "@/api/workflowProcess"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "WorkflowProcess",
  mixins: [infoList],
  data() {
    return {
      listApi: getWorkflowProcessList,
      dialogFormVisible: false,
      deleteVisible: false,
      multipleSelection: []
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
      if (this.searchInfo.hideIcon == "") {
        this.searchInfo.hideIcon = null;
      }
      this.getTableData();
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    async onDelete() {
      const ids = [];
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.id);
        });
      const res = await deleteWorkflowProcessByIds({ ids });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false;
        this.getTableData();
      }
    },
    async updateWorkflowProcess(row) {
      this.$router.push({
        name: "workflowCreate",
        query: {
          id: row.id,
          type: "edit"
        }
      });
    },
    async useWorkflowProcess(row) {
      this.$router.push({
        name: "workflowUse",
        query: {
          workflowId: row.id
        }
      });
    },
    async viewWorkflowProcess(row) {
      this.$router.push({
        name: "workflowCreate",
        query: {
          id: row.id,
          type: "view"
        }
      });
    },
    async deleteWorkflowProcess(row) {
      row.visible = false;
      const res = await deleteWorkflowProcess({ id: row.id });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
          this.page--;
        }
        this.getTableData();
      }
    },
    openDialog() {
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  }
};
</script>

<style>
</style>