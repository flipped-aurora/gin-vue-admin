<template>
  <el-drawer title="媒体库" :visible.sync="drawer">
    <div style="display:flex;justify-content:space-around;flex-wrap:wrap;padding-top:40px">
      <el-image
        class="header-img-box-list"
        :src="(item.url && item.url.slice(0, 4) !== 'http')?path+item.url:item.url"
        v-for="(item,key) in picList"
        :key="key"
        @click.native="chooseImg(item.url,target,targetKey)"
      >
        <div slot="error" class="header-img-box-list">
          <i class="el-icon-picture-outline"></i>
        </div>
      </el-image>
    </div>
  </el-drawer>
</template>

<script>
const path = process.env.VUE_APP_BASE_API
import { getFileList } from "@/api/fileUploadAndDownload";
export default {
  props: {
    target: [Object],
    targetKey: [String]
  },
  data() {
    return {
      drawer: false,
      picList: [],
      path:path
    };
  },
  methods: {
    chooseImg(url, target, targetKey) {
      if(target&&targetKey){
        target[targetKey] = url;
      }
      this.$emit("enter-img", url);
      this.drawer = false;
    },
    async open() {
      const res = await getFileList({ page: 1, pageSize: 9999 });
      this.picList = res.data.list;
      this.drawer = true;
    }
  }
};
</script>

<style lang="scss">
.header-img-box-list {
  width: 180px;
  height: 180px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 180px;
  cursor: pointer;
}
</style>