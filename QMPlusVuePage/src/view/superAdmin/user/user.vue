<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addUser" type="primary">新增用户</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="头像" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <img :src="scope.row.headerImg" height="50" width="50" />
          </div>
        </template>
      </el-table-column>
      <el-table-column label="uuid" min-width="250" prop="uuid"></el-table-column>
      <el-table-column label="用户名" min-width="150" prop="userName"></el-table-column>
      <el-table-column label="昵称" min-width="150" prop="nickName"></el-table-column>
      <el-table-column label="用户角色" min-width="150">
        <template slot-scope="scope">
          <el-select
            @change="changeAuthority(scope.row)"
            placeholder="请选择"
            v-model="scope.row.authority.authorityId"
          >
            <el-option
              :key="item.authorityId"
              :label="item.authorityName"
              :value="item.authorityId"
              v-for="item in authOptions"
            ></el-option>
          </el-select>
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

    <el-dialog :visible.sync="addUserDialog" custom-class="user-dialog" title="新增用户">
      <el-form :model="userInfo">
        <el-form-item label="用户名" label-width="80px">
          <el-input v-model="userInfo.username"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="80px">
          <el-input v-model="userInfo.password"></el-input>
        </el-form-item>
        <el-form-item label="别名" label-width="80px">
          <el-input v-model="userInfo.nickName"></el-input>
        </el-form-item>
        <el-form-item label="头像" label-width="80px">
          <el-upload
            :headers="{'x-token':token}"
            :on-success="handleAvatarSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
          >
            <img :src="userInfo.headerImg" class="avatar" v-if="userInfo.headerImg" />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="用户角色" label-width="80px">
          <el-select placeholder="请选择" v-model="userInfo.authorityId">
            <el-option
              :key="item.authorityId"
              :label="item.authorityName"
              :value="item.authorityId"
              v-for="item in authOptions"
            ></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddUserDialog">取 消</el-button>
        <el-button @click="enterAddUserDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API
import { getUserList, setUserAuthority, regist } from '@/api/user'
import { getAuthorityList } from '@/api/authority'
import infoList from '@/components/mixins/infoList'
import { mapGetters } from 'vuex'
export default {
  name: 'Api',
  mixins: [infoList],
  data() {
    return {
      listApi: getUserList,
      listKey: 'userList',
      path:path,
      authOptions: [],
      addUserDialog: false,
      userInfo: {
        username: '',
        password: '',
        nickName: '',
        headerImg: '',
        authorityId: ''
      }
    }
  },
  computed: {
    ...mapGetters('user', ['token'])
  },
  methods: {
    async enterAddUserDialog() {
      const res = await regist(this.userInfo)
      if (res.success) {
        this.$message({ type: 'success', message: '创建成功' })
      }
      await this.getTableData()
      this.closeAddUserDialog()
    },
    closeAddUserDialog() {
      this.userInfo = {
        username: '',
        password: '',
        nickName: '',
        headerImg: '',
        authorityId: ''
      }
      this.addUserDialog = false
    },
    handleAvatarSuccess(res) {
      this.userInfo.headerImg = res.data.file.url
    },
    addUser() {
      this.addUserDialog = true
    },
    async changeAuthority(row) {
      const res = await setUserAuthority({
        uuid: row.uuid,
        authorityId: row.authority.authorityId
      })
      if (res.success) {
        this.$message({ type: 'success', message: '角色设置成功' })
      }
    }
  },
  async created() {
    const res = await getAuthorityList({ page: 1, pageSize: 999 })
    this.authOptions = res.data.list
  }
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