<template>
  <div>
    <template>
      <el-upload
        :file-list="fileList"
        :headers="{'x-token':token}"
        :show-file-list="false"
        @before-upload="checkFile"
        @on-error="uploadError"
        @on-success="uploadSuccess"
        action="https://jsonplaceholder.typicode.com/posts/"
      >
        <el-button size="small" type="primary">点击上传</el-button>
        <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
      </el-upload>
      <el-table :data="tableData" border stripe>
          <el-table-column label="预览"  width="100">
          <template slot-scope="scope">
            <img :alt="scope.row.alt" width="80" height="80" :src="scope.row.url" />
          </template>
        </el-table-column>
        <el-table-column label="日期" prop="date" width="180"></el-table-column>
        <el-table-column label="文件名" prop="name" width="180">
        </el-table-column>
        <el-table-column :formatter="formatter" label="链接" prop="url"></el-table-column>
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
            <el-button size="small" type="text" @click="downloadFile(scope.row)">下载</el-button>
          </template>
        </el-table-column>
      </el-table>
    </template>
  </div>
</template>
   
<script>
import { mapGetters } from 'vuex'
import infoList from '@/components/mixins/infoList'
export default {
  name: 'Upload',
  mixins: [infoList],
  data() {
    return {
      tableData: [
        {
          date: '2019-10-25',
          name: '文件名.jpg',
          url: 'http://qmplusimg.henrongyi.top/1571321688timg.jpg',
          tag: 'jpg'
        },
        {
          date: '2019-10-25',
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
  methods: {
    checkFile(file) {
      const isJPG = file.type === 'image/jpeg'
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!')
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!')
      }
      return isJPG && isLt2M
    },
    downloadFile(){
    }
  }
}
</script>