<template>
  <div class="authority">
    <div class="button-box clearflex">
      <el-button @click="addAuthority('0')" type="primary">新增角色</el-button>
    </div>
    <el-table
      :data="tableData"
      :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
      border
      row-key="authorityId"
      stripe
      style="width: 100%"
    >
      <el-table-column label="角色id" min-width="180" prop="authorityId"></el-table-column>
      <el-table-column label="角色名称" min-width="180" prop="authorityName"></el-table-column>
      <el-table-column fixed="right" label="操作" width="460">
        <template slot-scope="scope">
          <el-button @click="opdendrawer(scope.row)" size="small" type="primary">设置权限</el-button>
          <el-button
            @click="addAuthority(scope.row.authorityId)"
            icon="el-icon-plus"
            size="small"
            type="primary"
          >新增子角色</el-button>
          <el-button
            @click="copyAuthority(scope.row)"
            icon="el-icon-copy-document"
            size="small"
            type="primary"
          >拷贝</el-button>
          <el-button
            @click="editAuthority(scope.row)"
            icon="el-icon-edit"
            size="small"
            type="primary"
          >编辑</el-button>
          <el-button
            @click="deleteAuth(scope.row)"
            icon="el-icon-delete"
            size="small"
            type="danger"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 新增角色弹窗 -->
    <el-dialog :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :model="form" :rules="rules" ref="authorityForm">
        <el-form-item label="父级角色" prop="parentId">
          <el-cascader
            :disabled="dialogType=='add'"
            :options="AuthorityOption"
            :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
            v-model="form.parentId"
          ></el-cascader>
        </el-form-item>
        <el-form-item label="角色ID" prop="authorityId">
          <el-input :disabled="dialogType=='edit'" autocomplete="off" v-model="form.authorityId"></el-input>
        </el-form-item>
        <el-form-item label="角色姓名" prop="authorityName">
          <el-input autocomplete="off" v-model="form.authorityName"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-drawer :visible.sync="drawer" :with-header="false" size="40%" title="角色配置" v-if="drawer">
      <el-tabs :before-leave="autoEnter" class="role-box" type="border-card">
        <el-tab-pane label="角色菜单">
          <Menus :row="activeRow" ref="menus" />
        </el-tab-pane>
        <el-tab-pane label="角色api">
          <apis :row="activeRow" ref="apis" />
        </el-tab-pane>
        <el-tab-pane label="资源权限">
          <Datas :authority="tableData" :row="activeRow" ref="datas" />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成

import {
  getAuthorityList,
  deleteAuthority,
  createAuthority,
  updateAuthority,
  copyAuthority
} from "@/api/authority";

import Menus from "@/view/superAdmin/authority/components/menus";
import Apis from "@/view/superAdmin/authority/components/apis";
import Datas from "@/view/superAdmin/authority/components/datas";

import infoList from "@/mixins/infoList";
export default {
  name: "Authority",
  mixins: [infoList],
  data() {
    var mustUint = (rule, value, callback) => {
      if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
        return callback(new Error("请输入正整数"));
      }
      return callback();
    };

    return {
      AuthorityOption: [
        {
          authorityId: "0",
          authorityName: "根角色"
        }
      ],
      listApi: getAuthorityList,
      drawer: false,
      dialogType: "add",
      activeRow: {},
      activeUserId: 0,
      dialogTitle: "新增角色",
      dialogFormVisible: false,
      apiDialogFlag: false,
      copyForm: {},
      form: {
        authorityId: "",
        authorityName: "",
        parentId: "0"
      },
      rules: {
        authorityId: [
          { required: true, message: "请输入角色ID", trigger: "blur" },
          { validator: mustUint, trigger: "blur" }
        ],
        authorityName: [
          { required: true, message: "请输入角色名", trigger: "blur" }
        ],
        parentId: [
          { required: true, message: "请选择请求方式", trigger: "blur" }
        ]
      }
    };
  },
  components: {
    Menus,
    Apis,
    Datas
  },
  methods: {
    autoEnter(activeName, oldActiveName) {
      const paneArr = ["menus", "apis", "datas"];
      if (oldActiveName) {
        if (this.$refs[paneArr[oldActiveName]].needConfirm) {
          this.$refs[paneArr[oldActiveName]].enterAndNext();
          this.$refs[paneArr[oldActiveName]].needConfirm = false;
        }
      }
    },
    // 拷贝角色
    copyAuthority(row) {
      this.setOptions();
      this.dialogTitle = "拷贝角色";
      this.dialogType = "copy";
      for (let k in this.form) {
        this.form[k] = row[k];
      }
      this.copyForm = row;
      this.dialogFormVisible = true;
    },
    opdendrawer(row) {
      this.drawer = true;
      this.activeRow = row;
    },
    // 删除角色
    deleteAuth(row) {
      this.$confirm("此操作将永久删除该角色, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteAuthority({ authorityId: row.authorityId });
          if (res.code == 0) {
            this.$message({
              type: "success",
              message: "删除成功!"
            });
            if (this.tableData.length == 1) {
              this.page--;
            }
            this.getTableData();
          }
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    // 初始化表单
    initForm() {
      if (this.$refs.authorityForm) {
        this.$refs.authorityForm.resetFields();
      }
      this.form = {
        authorityId: "",
        authorityName: "",
        parentId: "0"
      };
    },
    // 关闭窗口
    closeDialog() {
      this.initForm();
      this.dialogFormVisible = false;
      this.apiDialogFlag = false;
    },
    // 确定弹窗

    async enterDialog() {
      if (this.form.authorityId == "0") {
        this.$message({
          type: "error",
          message: "角色id不能为0"
        });
        return false;
      }
      this.$refs.authorityForm.validate(async valid => {
        if (valid) {
          switch (this.dialogType) {
            case "add":
              {
                const res = await createAuthority(this.form);
                if (res.code == 0) {
                  this.$message({
                    type: "success",
                    message: "添加成功!"
                  });
                  this.getTableData();
                  this.closeDialog();
                }
              }
              break;
            case "edit":
              {
                const res = await updateAuthority(this.form);
                if (res.code == 0) {
                  this.$message({
                    type: "success",
                    message: "添加成功!"
                  });
                  this.getTableData();
                  this.closeDialog();
                }
              }
              break;
            case "copy": {
              const data = {
                authority: {
                  authorityId: "string",
                  authorityName: "string",
                  datauthorityId: [],
                  parentId: "string"
                },
                oldAuthorityId: 0
              };
              data.authority.authorityId = this.form.authorityId;
              data.authority.authorityName = this.form.authorityName;
              data.authority.parentId = this.form.parentId;
              data.authority.dataAuthorityId = this.copyForm.dataAuthorityId;
              data.oldAuthorityId = this.copyForm.authorityId;
              const res = await copyAuthority(data);
              if (res.code == 0) {
                this.$message({
                  type: "success",
                  message: "复制成功！"
                });
                this.getTableData();
              }
            }
          }

          this.initForm();
          this.dialogFormVisible = false;
        }
      });
    },
    setOptions() {
      this.AuthorityOption = [
        {
          authorityId: "0",
          authorityName: "根角色"
        }
      ];
      this.setAuthorityOptions(this.tableData, this.AuthorityOption, false);
    },
    setAuthorityOptions(AuthorityData, optionsData, disabled) {
      this.form.authorityId = String(this.form.authorityId);
      AuthorityData &&
        AuthorityData.map(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              disabled: disabled || item.authorityId == this.form.authorityId,
              children: []
            };
            this.setAuthorityOptions(
              item.children,
              option.children,
              disabled || item.authorityId == this.form.authorityId
            );
            optionsData.push(option);
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              disabled: disabled || item.authorityId == this.form.authorityId
            };
            optionsData.push(option);
          }
        });
    },
    // 增加角色
    addAuthority(parentId) {
      this.initForm();
      this.dialogTitle = "新增角色";
      this.dialogType = "add";
      this.form.parentId = parentId;
      this.setOptions();
      this.dialogFormVisible = true;
    },
    // 编辑角色
    editAuthority(row) {
      this.setOptions();
      this.dialogTitle = "编辑角色";
      this.dialogType = "edit";
      for (let key in this.form) {
        this.form[key] = row[key];
      }
      this.setOptions();
      this.dialogFormVisible = true;
    }
  },
  async created() {
    this.pageSize = 999;
    await this.getTableData();
  }
};
</script>
<style lang="scss">
.authority {
  .el-input-number {
    margin-left: 15px;
    span {
      display: none;
    }
  }
  .button-box {
    padding: 10px 20px;
    .el-button {
      float: right;
    }
  }
}
.role-box {
  .el-tabs__content {
    height: calc(100vh - 150px);
    overflow: auto;
  }
}
</style>