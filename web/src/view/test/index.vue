<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <!-- 此处可使用表单生成器生成搜索条件 -->
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增</el-button>
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
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
     <el-table-column label="用户名" prop="name" width="120"></el-table-column>
    
     <el-table-column label="年龄" prop="age" width="120"></el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateTestTest(scope.row)" size="small" type="text">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteTestTest(scope.row)">确定</el-button>
            </div>
            <el-button type="text" size="mini" slot="reference">删除</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <!-- 此处请使用表单生成器生成form填充  -->


    <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
      <el-form-item label="用户名" prop="user_name">
        <el-input v-model="formData.user_name" placeholder="请输入用户名" clearable :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="密码" prop="pwd">
        <el-input v-model="formData.pwd" placeholder="请输入密码" clearable show-password :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="性别" prop="sex">
        <el-radio-group v-model="formData.sex" size="medium">
          <el-radio v-for="(item, index) in sexOptions" :key="index" :label="item.value"
            :disabled="item.disabled">{{item.label}}</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="内容" prop="content">
        <el-input v-model="formData.content" type="textarea" placeholder="请输入内容"
          :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
      </el-form-item>
      <!-- <el-form-item size="large">
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item> -->
    </el-form>
      
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createTestTest,
    deleteTestTest,
    updateTestTest,
    findTestTest,
    getTestTestList
} from "@/api/test";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  components: {},
  props: [],
  name: "TestTest",
  mixins: [infoList],
  data() {
    return {
      listApi: getTestTestList,
      dialogFormVisible: false,
      visible: false,
      type: "",
    //   form: {
    //     name:null,age:null,
    //   },
      formData: {
        user_name: undefined,
        pwd: undefined,
        sex: 1,
        content: undefined,
      },
      rules: {
        user_name: [{
          required: true,
          message: '请输入用户名',
          trigger: 'blur'
        }],
        pwd: [{
          required: true,
          message: '请输入密码',
          trigger: 'blur'
        }],
        sex: [{
          required: true,
          message: '性别不能为空',
          trigger: 'change'
        }],
        content: [{
          required: true,
          message: '请输入内容',
          trigger: 'blur'
        }],
      },
      sexOptions: [{
        "label": "男",
        "value": 1
      }, {
        "label": "女",
        "value": 2
      }],


      
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
    }
  },
  methods: {
    async updateTestTest(row) {
      const res = await findTestTest({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.form = res.data.ret;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.form = {
        
          user_name:null,
          pwd:null,
          sex:null,
          content:null,
      };
    },
    async deleteTestTest(row) {
      this.visible = false;
      const res = await deleteTestTest({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createTestTest(this.formData);
          break;
        case "update":
          res = await updateTestTest(this.formData);
          break;
        default:
          res = await createTestTest(this.formData);
          break;
      }

      if (res.code == 0) {
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },

submitForm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        // TODO 提交表单
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
  },
  created() {
    this.getTableData();
  }
};
</script>

<style>
</style>