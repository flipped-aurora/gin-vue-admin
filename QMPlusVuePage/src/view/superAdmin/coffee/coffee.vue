<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addCoffee" type="primary">新增咖啡</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="图片" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <img :src="scope.row.img" height="50" width="50" />
          </div>
        </template>
      </el-table-column>
      <el-table-column label="咖啡编号" min-width="250" prop="uuid"></el-table-column>
      <el-table-column label="咖啡名称" min-width="150" prop="name"></el-table-column>
      <el-table-column label="价格" min-width="150" prop="value"></el-table-column>
      <el-table-column label="咖啡类型" min-width="150" prop="code">
        <template slot-scope="scope">
            {{type[scope.row.code]}}
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button @click="deleteCoffee(scope.row)" size="small" type="text">删除咖啡</el-button>
          <el-button @click="editCoffee(scope.row)" size="small" type="text">编辑咖啡</el-button>
          <el-button @click="getCoffeeSpec(scope.row)" size="small" type="text">查看咖啡规格</el-button>
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

    <el-dialog :visible.sync="addCoffeeDialog" custom-class="user-dialog" :title="titleMap[dialogTitle]">
      <el-form :model="coffeeInfo" :rules="rules" ref="coffeeInfo">
        <el-form-item label="咖啡名称" label-width="80px" required prop="name">
          <el-input v-model="coffeeInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="价格" label-width="80px" required prop="value">
          <el-input v-model="coffeeInfo.value"></el-input>
        </el-form-item>
        <el-form-item label="描述" label-width="80px" required prop="des">
        <quill-editor 
            v-model="coffeeInfo.des" 
            ref="myQuillEditor" 
            :options="editorOption" 
            @blur="onEditorBlur($event)" @focus="onEditorFocus($event)"
            @change="onEditorChange($event)">
        </quill-editor>
        </el-form-item>
        <el-form-item label="图片" label-width="80px">
          <el-upload
            :on-success="handleAvatarSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
          >
            <img :src="coffeeInfo.img" class="avatar" v-if="coffeeInfo.img" />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="咖啡类型" label-width="80px" required prop="code">
          <el-select placeholder="请选择" v-model="coffeeInfo.code">
            <el-option
              :key="item.code"
              :label="item.name"
              :value="item.code"
              v-for="item in typeOptions"
            ></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddCoffeeDialog">取 消</el-button>
        <el-button @click="enterAddCoffeeDialog('coffeeInfo')" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog title="规格" :visible.sync="coffeeSpecDialog" custom-class="user-dialog">
      <el-table :data="Spec" size="mini">
        <el-table-column prop="spec_id" label="规格编号" width="180"></el-table-column>
        <el-table-column prop="name" label="规格名称" width="180"></el-table-column>
        <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button @click="enterCoffeeSpecDetailDialog(scope.row)" size="small" type="text">查看详细信息</el-button>
        </template>
      </el-table-column>
      </el-table>
    </el-dialog>

  <el-dialog
  title="咖啡规格详细信息"
  :visible.sync="coffeeSpecDetailDialog"
  width="30%"
  :before-close="CloseSpecDetailDialog">
      <el-table :data="specDetail" size="mini">
        <el-table-column prop="coffee_id" label="咖啡编号" width="180"></el-table-column>
        <el-table-column prop="value" label="规格名称" width="180"></el-table-column>
        <el-table-column prop="price_incre" label="价格增加" width="180"></el-table-column>
  <span slot="footer" class="dialog-footer">
    <el-button @click="dialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
  </span>
        </el-table>
  </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getCoffeeList,
  addCoffee,
  getCoffeeTypeList,
  updateCoffeeType,
  updateCoffee,
  delCoffee,
  getCoffeeById,
  getCoffeeSpecByCoffeeId,
  getCoffeeSpecDetail
} from "@/api/coffee";
import infoList from "@/components/mixins/infoList";
import { quillEditor } from "vue-quill-editor"; //调用编辑器
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css';
import 'quill/dist/quill.bubble.css';
export default {
  name: "coffee",
  mixins: [infoList],
  data() {
    const checkname = (rule, value, callback) => {
      if(!value) {
        return callback(new Error('咖啡名不能为空'))
      } else {
        callback()
      }
    }
    const checkvalue = (rule, value, callback) => {
      if(!value) {
        return callback(new Error('咖啡价格不能为空'))
      } else {
        callback()
      }
    }
    const checkdes = (rule, value, callback) => {
      if(!value) {
        return callback(new Error('咖啡描述不能为空'))
      } else {
        callback()
      }
    }
    const checkcode = (rule, value, callback) => {
      if(!value) {
        return callback(new Error('咖啡类型编号不能为空'))
      } else {
        callback()
      }
    }

    return {
      listApi: getCoffeeList,
      listKey: "coffeeList",
      path: path,
      addCoffeeDialog: false,
      coffeeSpecDialog: false,
      coffeeSpecDetailDialog: false,
      dialogTitle: '',
      titleMap : {
        addData : "添加咖啡",
        updateData : "修改咖啡"
      },
      isEdit: false,
      coffeeInfo: {
        name: "",
        value: 0.0,
        des: "",
        img: "",
        code: ""
      },
      rules: {
        name: [{required: true, validator: checkname, trigger: 'blur' }],
        value: [{required: true, validator: checkvalue, trigger: 'blur' }],
        des:  [{required: true, validator: checkdes, trigger: 'blur' }],
        code:  [{required: true, validator: checkcode, trigger: 'blur' }]
      },
      typeOptions: [],
      Spec: [],
      specDetail: [],
      editorOption: {},
      type: {}
    };
  },
  methods: {
    async enterAddCoffeeDialog(formName) {
      // eslint-disable-next-line no-console
      //console.log(this.coffeeInfo)
      this.$refs[formName].validate(async v => {
      if(v) {
        this.coffeeInfo.value = Number(this.coffeeInfo.value);
        let res;
        if (this.isEdit) {
          res = await updateCoffee(this.coffeeInfo);
        } else {
          res = await addCoffee(this.coffeeInfo);
        }

        if (res.success) {
          this.$message({ type: "success", message: "创建成功" });
        }
        await this.getTableData();
        this.closeAddCoffeeDialog();
      }else {
          this.$message({
            type: 'error',
            message: '请正确填写信息',
            showClose: true
          })
          return false
      }
    })
    },
    closeAddCoffeeDialog() {
      this.coffeeInfo = {
        name: "",
        value: 0.0,
        des: "",
        img: "",
        code: ""
      };
      this.addCoffeeDialog = false;
    },
    handleAvatarSuccess(res) {
      this.coffeeInfo.img = res.data.file.url;
    },
    addCoffee() {
      this.coffeeInfo = {
        name: "",
        value: 0.0,
        code: "",
        des: "",
        img: ""
      };
      this.dialogTitle = "addData";
      this.isEdit = false;
      this.addCoffeeDialog = true;
    },
    async changeCoffeeType(row) {
      const res = await updateCoffeeType({
        uuid: row.uuid,
        code: row.code
      });
      if (res.success) {
        this.$message({ type: "success", message: "咖啡类型设置成功" });
      }
    },
    async editCoffee(row) {
      const res = await getCoffeeById({ uuid: row.uuid });
      this.coffeeInfo = res.data.coffee;
      this.isEdit = true;
      this.dialogTitle = "updateData";
      this.addCoffeeDialog = true;

    },
    async deleteCoffee(row) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async ()=> {
          const res = await delCoffee({ uuid: row.uuid });
          if (res.success) {
            this.$message({ type: "success", message: "删除咖啡成功" });
          }
          await this.getTableData();    
      }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
      })

    },
    async getCoffeeSpec(row) {
      const res = await getCoffeeSpecByCoffeeId({ uuid: row.uuid });
      if (res.success) {
        this.$message({ type: "success", message: "获取成功" });
      }
      this.Spec = res.data.coffeeSpec;
      
      this.coffeeSpecDialog = true
    },
    async enterCoffeeSpecDetailDialog(row) {
      const res = await getCoffeeSpecDetail({spec_id : row.spec_id})
      this.specDetail = res.data.coffeeSpecDetail
      console.log(this.specDetail)
      this.coffeeSpecDetailDialog = true
    },
    CloseSpecDetailDialog() {
      this.specDetail = []
      this.coffeeSpecDetailDialog = false
    },
    onEditorChange({ editor, html, text }) {
      this.coffeeInfo.des = html;
    }
  },
  async created() {
    this.page = 1;
    this.pageSize = 999;
    const res = await getCoffeeTypeList({ page: 1, pageSize: 999 });
    this.typeOptions = res.data.coffeetype;
    this.typeOptions.forEach((item) => {
      this.type[item.code] = item.name
    })
    // console.log(this.type)
    // console.log(this.typeOptions)
  },
}
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}

.user-dialog {
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
</style>