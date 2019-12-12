<template>
  <div>
    <el-upload
      :before-upload="checkFile"
      :headers="{'x-token':token}"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      action="/api/fileUploadAndDownload/upload"
    >
      <el-button size="small" type="primary">点击上传</el-button>
      <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
    </el-upload>
    <el-table :data="tableData" border stripe>
      <el-table-column label="预览" width="100">
        <template slot-scope="scope">
          <img :alt="scope.row.alt" :src="scope.row.url" height="80" width="80" />
        </template>
      </el-table-column>
      <el-table-column label="日期" prop="UpdatedAt" width="180">
        <template slot-scope="scope">
          <div>{{scope.row.UpdatedAt|formatDate}}</div>
        </template>
      </el-table-column>
      <el-table-column label="文件名" prop="name" width="180"></el-table-column>
      <el-table-column label="链接" prop="url"></el-table-column>
      <el-table-column label="标签" prop="tag" width="100">
        <template slot-scope="scope">
          <el-tag
            :type="scope.row.tag === 'jpg' ? 'primary' : 'success'"
            disable-transitions
          >{{scope.row.tag}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="100">
        <template slot-scope="scope">
          <el-button @click="downloadFile(scope.row)" size="small" type="text">下载</el-button>
          <el-button @click="deleteFile(scope.row)" size="small" type="text">删除</el-button>
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
import { mapGetters } from 'vuex'
import infoList from '@/components/mixins/infoList'
import { getFileList, deleteFile } from '@/api/fileUploadAndDownload'
import { downloadImage } from '@/utils/downloadImg'
import { formatTimeToStr } from '@/utils/data'
export default {
  name: 'Upload',
  mixins: [infoList],
  data() {
    return {
      listApi: getFileList,
      listKey: 'list',
      tableData: [
        {
          UpdatedAt: '2019-10-25',
          name: '文件名.jpg',
          url: 'http://qmplusimg.henrongyi.top/1571321688timg.jpg',
          tag: 'jpg'
        },
        {
          UpdatedAt: '2019-10-25',
          name: '文件名.jpg',
          url: 'http://qmplusimg.henrongyi.top/157162774820191015140921496.gif',
          tag: 'gif'
        }
      ]
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    }
  },
  methods: {
    async deleteFile(row) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteFile(row)
          if (res.success) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    checkFile(file) {
      const isJPG = file.type === 'image/jpeg'
      const isPng = file.type === 'image/png'
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isJPG && !isPng) {
        this.$message.error('上传头像图片只能是 JPG或png 格式!')
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!')
      }
      return (isPng || isJPG) && isLt2M
    },
    uploadSuccess(res) {
      this.$message({
        type: 'success',
        message: '上传成功'
      })
      if (res.success) {
        this.getTableData()
      }
    },
    uploadError() {
      this.$message({
        type: 'error',
        message: '上传失败'
      })
    },
    downloadFile(row) {
      downloadImage(row.url, row.name)
    }
  }
}
</script>