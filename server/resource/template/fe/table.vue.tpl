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
         <template slot-scope="scope">{{ "{{scope.row.CreatedAt|formatDate}}" }}</template>
    </el-table-column>
    {{range .Fields}}
     <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120"></el-table-column>
    {{ end }}
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="update{{.StructName}}(scope.row)" size="small" type="text">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="delete{{.StructName}}(scope.row)">确定</el-button>
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
      此处请使用表单生成器生成form填充 表单默认绑定 formData 如手动修改过请自行修改key
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    create{{.StructName}},
    delete{{.StructName}},
    update{{.StructName}},
    find{{.StructName}},
    get{{.StructName}}List
} from "@/api/{{.PackageName}}";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "{{.StructName}}",
  mixins: [infoList],
  data() {
    return {
      listApi: get{{.StructName}}List,
      dialogFormVisible: false,
      visible: false,
      type: "",
      formData: {
        {{range .Fields}}{{.FieldJson}}:null,{{ end }}
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
    }
  },
  methods: {
    async update{{.StructName}}(row) {
      const res = await find{{.StructName}}({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.re{{.Abbreviation}};
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        {{range .Fields}}
          {{.FieldJson}}:null,{{ end }}
      };
    },
    async delete{{.StructName}}(row) {
      this.visible = false;
      const res = await delete{{.StructName}}({ ID: row.ID });
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
          res = await create{{.StructName}}(this.formData);
          break;
        case "update":
          res = await update{{.StructName}}(this.formData);
          break;
        default:
          res = await create{{.StructName}}(this.formData);
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
    }
  },
  created() {
    this.getTableData();
  }
};
</script>

<style>
</style>