<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <el-carousel class="-mt-2">
    <el-carousel-item
      class="cursor-pointer lg:h-40"
      v-for="(item, index) in banners"
      :key="index"
      @click="openLink(item.link)"
    >
      <el-image class="h-full w-full" :src="item.img" fit="fill"></el-image>
    </el-carousel-item>
  </el-carousel>
</template>

<script setup>
  import banner from '@/assets/banner.jpg'
  import banner2 from '@/assets/banner2.jpg'

  const openLink = (link) => {
    window.open(link, '_blank')
  }

  const banners = [
    {
      img: banner,
      link: 'https://gin-vue-admin.com/empower/index.html'
    },
    {
      img: banner2,
      link: 'https://plugin.gin-vue-admin.com'
    },
    {
      img: 'https://qmplusimg.henrongyi.top/gvaDemo/k8s.jpg',
      link: 'https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=42'
    }
  ]
</script>

<style scoped lang="scss"></style>
