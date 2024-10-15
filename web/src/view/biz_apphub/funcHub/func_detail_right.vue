<script setup>

import {formatDate} from "@/utils/format";
import {defineProps} from "vue";
const props = defineProps({
  func:{
    type: Object,
  }
});
function getTags() {
  if (props.func.detail.tags) {
    return props.func.detail.tags.split(";")
  }
  return []

}


</script>

<template>
  <div>
    <div style="padding: 20px 5px 5px 0px">
      <div>
        <el-row>
          <el-col :span="6"><div class="grid-content ep-bg-purple" />
            <el-avatar :size="50" src="https://img1.baidu.com/it/u=1465664392,2808406094&fm=253&fmt=auto&app=120&f=JPEG?w=800&h=800" />
          </el-col>
          <el-col :span="18">
            <div>
              <span style="font-size: 18px">这个网站的老大</span>
              <!--                  <p>出生在中国，爱好写代码</p>-->
              <p style="margin: 6px 0px">
                <el-tooltip
                    class="box-item"
                    effect="dark"
                    content="云函数站点头目，00后独立开发者"
                    placement="top-start"
                >
                  <el-text truncated size="small" class="mx-1">云函数站点头目，00后独立开发者</el-text>
                </el-tooltip>
              </p>
            </div>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="6"><div class="grid-content ep-bg-purple" />
          </el-col>
          <el-col :span="9">
            <el-button size="small" type="primary">+ 关注 988</el-button>
          </el-col>
          <el-col :span="9">
            <el-button size="small" type="primary">打赏项目</el-button>
          </el-col>
        </el-row>
      </div>

      <el-divider border-style="double" />
      <h2>
        {{func.detail.cn_name}}
      </h2>
      <el-row>
        <el-col :span="6">
          <div>{{func.detail.classify}}</div>
        </el-col>
        <el-col :span="6">
          <div style="display: flex;align-items: center;">
            <View  style="height: 16px"/>
            <span style="margin-left: 2px">1200</span>
          </div>
        </el-col>

        <el-col :span="12"><div>
          <!--                CreatedAt-->
          {{ formatDate(func.detail.CreatedAt) }}
        </div></el-col>
      </el-row>
      <h3>函数简介</h3>
      <el-text size="small" class="mx-1">{{func.detail.content}}</el-text>

      <!--            <span>{{func.detail.content}}</span>-->
      <h3>相关Tag</h3>
      <div><el-tag effect="plain" style="margin: 3px" :key="i" v-for="(v,i) in getTags()">{{v}}</el-tag></div>
    </div>

    <el-row style="margin-top: 10px">
      <el-col :span="6">
        <div style="display: flex;align-items: center;">
          <el-tooltip
              class="box-item"
              effect="dark"
              content="300星标"
              placement="top-start"
          >
            <Star style="height: 30px"/>
          </el-tooltip>
          <span style="margin-left: 3px">200</span>
        </div>
      </el-col>
      <el-col :span="6">

        <div style="display: flex;align-items: center;">
          <!--                <Coin style="height: 30px"/><span style="margin-left: 3px">300</span>-->
          <el-tooltip
              class="box-item"
              effect="dark"
              content="300打赏"
              placement="top-start"
          >
            <Coin style="height: 30px"/>
          </el-tooltip>
          <span style="margin-left: 3px">200</span>
        </div>

      </el-col>

      <el-col :span="6"><div>

        <div style="display: flex;align-items: center;">
          <!--                <el-icon><Cpu /></el-icon>-->
          <!--                <CollectionTag style="height: 30px"/><span style="margin-left: 3px">300</span>-->

          <el-tooltip
              class="box-item"
              effect="dark"
              content="300收藏"
              placement="top-start"
          >
            <CollectionTag style="height: 30px"/>
          </el-tooltip>
          <span style="margin-left: 3px">200</span>
        </div>
      </div>
      </el-col>

      <el-col :span="6"><div>

        <div style="display: flex;align-items: center;">
          <!--                <Cpu style="height: 30px"/><span style="margin-left: 3px">300</span>-->

          <el-tooltip
              class="box-item"
              effect="dark"
              content="300次运行"
              placement="top-start"
          >
            <Cpu style="height: 30px"/>
          </el-tooltip>
          <span style="margin-left: 3px">200</span>
        </div>
      </div>
      </el-col>

    </el-row>
  </div>
</template>

<style scoped lang="scss">

</style>
