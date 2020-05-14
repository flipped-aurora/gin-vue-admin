<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        此处请使用表单生成器生成form填充 表单默认绑定 formData 如手动修改过请自行修改key
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border ref="multipleTable" stripe style="width: 100%" tooltip-effect="dark">
      <!--<el-table-column type="selection" width="55"></el-table-column>-->
      <el-table-column label="ID" prop="ID" width="55" align="center"></el-table-column>

      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>

      <el-table-column label="分类名称" prop="c_name" width="120"></el-table-column>

      <el-table-column label="关键字" prop="key_word" width="120"></el-table-column>

      <el-table-column label="描述" prop="desc" width="120"></el-table-column>

      <el-table-column label="ico地址" prop="ico_image" width="180"></el-table-column>

      <el-table-column label="背景图片" prop="bg_image" width="180"></el-table-column>

      <el-table-column label="排序" prop="sort" width="60" align="center"></el-table-column>

      <el-table-column label="跳转连接" prop="href" width="180"></el-table-column>

      <!-- 
     <el-table-column label="是否启动" prop="status" width="120"></el-table-column>
      -->

      <el-table-column label="是否启动" prop="status" width="120" align="center">
        <template slot-scope="scope">
          <p>{{ scope.row.status==1?'启动':'禁用' }}</p>
        </template>
      </el-table-column>

      <el-table-column label="操作" width="160" align="center">
        <template slot-scope="scope">
          <el-button @click="updateCategory(scope.row)" size="small" type="text">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteCategory(scope.row)">确定</el-button>
            </div>
            <el-button type="text" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}" :total="total" @current-change="handleCurrentChange"
      @size-change="handleSizeChange" layout="total, sizes, prev, pager, next, jumper"></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="添加分类">
      <!--此处请使用表单生成器生成form填充 表单默认绑定 formData 如手动修改过请自行修改key-->
      <template>
        <div>
          <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
            <el-form-item label="分类名称" prop="c_name">
              <el-input v-model="formData.c_name" placeholder="请输入分类名称" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
            <el-form-item label="关键字" prop="key_word">
              <el-input v-model="formData.key_word" placeholder="请输入关键字" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
            <el-form-item label="描述" prop="desc">
              <el-input v-model="formData.desc" placeholder="请输入描述" clearable :style="{width: '100%'}"></el-input>
            </el-form-item>
            <el-form-item label="ico" prop="ico_image">
              <el-input v-model="formData.ico_image" placeholder="请输入ico" clearable :style="{width: '100%'}"></el-input>
            </el-form-item>
            <el-form-item label="背景图片" prop="bg_image">
              <el-input v-model="formData.bg_image" placeholder="请输入背景图片" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
            <el-form-item label="排序" prop="sort">
              <el-input-number v-model="formData.sort" placeholder="请输入排序" step-strictly controls-position=right>
              </el-input-number>
            </el-form-item>
            <el-form-item label="跳转链接" prop="href">
              <el-input v-model="formData.href" placeholder="请输入跳转链接" clearable :style="{width: '100%'}"></el-input>
            </el-form-item>
            <el-form-item label="是否显示" prop="status">
              <el-switch v-model="formData.status" :active-value='1' :inactive-value='2'></el-switch>
            </el-form-item>


            <!--
            <el-form-item size="large">
              <el-button type="primary" @click="submitForm">提交</el-button>
               <el-button @click="resetForm">重置</el-button>
             </el-form-item>
            -->
          </el-form>
        </div>
      </template>



      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import {
    createCategory,
    deleteCategory,
    updateCategory,
    findCategory,
    getCategoryList
  } from "@/api/category"; //  此处请自行替换地址
  import {
    formatTimeToStr
  } from "@/utils/data";
  import infoList from "@/components/mixins/infoList";

  export default {
    name: "Category",
    mixins: [infoList],
    data() {
      return {
        listApi: getCategoryList,
        dialogFormVisible: false,
        visible: false,
        type: "",
        formData: {
          c_name: null,
          key_word: null,
          desc: null,
          ico_image: null,
          bg_image: null,
          sort: 50,
          href: null,
          status: 1,
        },
        rules: {
          c_name: [{
              required: true,
              message: '请输入分类名称',
              trigger: 'blur'
            },
            {
              min: 3,
              max: 15,
              message: '长度在 3 到 15 个字符',
              trigger: 'blur'
            }
          ]
        }
      };
    },

    filters: {
      formatDate: function (time) {
        if (time != null && time != "") {
          var date = new Date(time);
          return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
        } else {
          return "";
        }
      }
    },
    methods: {
      async updateCategory(row) {
        const res = await findCategory({
          ID: row.ID
        });
        this.type = "update";
        if (res.code == 0) {
          this.formData = res.data.recate;
          this.dialogFormVisible = true;
        }
      },
      closeDialog() {
        this.dialogFormVisible = false;
        this.formData = {
          c_name: null,
          key_word: null,
          desc: null,
          ico_image: null,
          bg_image: null,
          sort: 50,
          href: null,
          status: 1,
        };
      },
      async deleteCategory(row) {
        this.visible = false;
        const res = await deleteCategory({
          ID: row.ID
        });
        if (res.code == 0) {
          this.$message({
            type: "success",
            message: "删除成功"
          });
          this.getTableData();
        }
      },
      async enterDialog() {
        this.$refs.elForm.validate(async (v) => {
          if (v) {
            let res;
            switch (this.type) {
              case "create":
                res = await createCategory(this.formData);
                break;
              case "update":
                res = await updateCategory(this.formData);
                break;
              default:
                res = await createCategory(this.formData);
                break;
            }
            if (res.code == 0) {
              this.closeDialog();
              this.getTableData();
            }
          } else {
            this.$message({
              type: "error",
              message: "请正确填写信息",
              showClose: true,
            });
            return false;
          }
        });

      },
      openDialog() {
        this.type = "create";
        this.dialogFormVisible = true;
      }
    },
    created() {
      this.getTableData();
    }
  };
</script>

<style>
</style>