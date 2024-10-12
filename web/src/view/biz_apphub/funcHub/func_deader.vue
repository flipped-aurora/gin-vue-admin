<script setup>
import {ref} from "vue";

const select = ref('')
const input3 = ref('')

import {Search} from "@element-plus/icons-vue";
</script>

<template>
  <div style="display: flex;justify-content: center;align-items:start" class="mt-4">
    <div style="flex: 2;display: flex;align-items: start;">
      <el-text style="font-size: 24px" class="mx-1" type="primary">云函数</el-text>

    </div>
    <div style="flex: 5">
      <el-input
          v-model="input3"
          style="max-width: 600px"
          placeholder="请输入要搜索的内容"
          class="input-with-select"
      >
        <template #prepend>
          <el-select v-model="select" placeholder="云函数" style="width: 115px">
            <el-option label="Restaurant" value="1" />
            <el-option label="Order No." value="2" />
            <el-option label="Tel" value="3" />
          </el-select>
        </template>
        <template #append>
          <el-button :icon="Search" />
        </template>
      </el-input></div>
    <div style="display: flex;flex: 2;justify-content: end;align-items:start">
      <div><el-avatar :size="30" src="https://img1.baidu.com/it/u=1090452517,2487311686&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=501" /></div>
      <div style="display: flex;justify-content: center;align-items: center;margin-left: 10px">
        <Edit style="width: 30px; height: 30px;" />
        <span>发布函数</span>
      </div>
      <!--            </el-icon>-->

    </div>
  </div>

</template>

<style scoped lang="scss">

</style>
