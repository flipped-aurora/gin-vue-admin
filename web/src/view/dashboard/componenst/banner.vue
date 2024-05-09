<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <el-carousel class=" h-40">
    <el-carousel-item v-for="(item , index) in banners" :key="index">
      <el-image :src="item" class="" fit="contain"></el-image>
    </el-carousel-item>
  </el-carousel>
</template>

<script setup>
const banners = [
    "https://qmplusimg.henrongyi.top/171163573320240328-221254.jpg",
    "https://qmplusimg.henrongyi.top/1657463878org.png"
]
</script>

<style scoped lang="scss">

</style>
