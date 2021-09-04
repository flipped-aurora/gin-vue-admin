<template>
  <el-drawer v-model="drawer" title="媒体库">
    <div class="media">
      <el-image
        v-for="(item,key) in picList"
        :key="key"
        class="header-img-box-list"
        :src="(item.url && item.url.slice(0, 4) !== 'http')?path+item.url:item.url"
        @click="chooseImg(item.url,target,targetKey)"
      >
        <template #error>
          <div class="header-img-box-list">
            <i class="el-icon-picture-outline" />
          </div>
        </template>
      </el-image>
    </div>
  </el-drawer>
</template>

<script>
const path = process.env.VUE_APP_BASE_API
import { getFileList } from '@/api/fileUploadAndDownload'
export default {
  props: {
    target: {
      type: Object,
      default: null
    },
    targetKey: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      drawer: false,
      picList: [],
      path: path
    }
  },
  methods: {
    chooseImg(url, target, targetKey) {
      if (target && targetKey) {
        target[targetKey] = url
      }
      this.$emit('enter-img', url)
      this.drawer = false
    },
    async open() {
      const res = await getFileList({ page: 1, pageSize: 9999 })
      this.picList = res.data.list
      this.drawer = true
    }
  }
}
</script>

<style lang="scss">
.media{
  display:flex;
  flex-wrap:wrap;
  .header-img-box-list {
    margin-top: 20px;
    width: 120px;
    margin-left: 20px;
    height: 120px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 180px;
    cursor: pointer;
}
}

</style>
