<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div class="grid grid-cols-2 gap-2">
    <a v-for="item in wikis" :key="item.url" :href="item.url" class="text-sm  text-gray-700 dark:text-gray-300 no-underline hover:text-active " target="_blank">
      {{ item.title }}
    </a>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const wikis = [
  {
    title: t('view.dashboard.grids.ginDoc'),
    url : 'https://gin-gonic.com/'
  },
  {
    title : t('view.dashboard.grids.gvaDoc'),
    url : 'https://www.gin-vue-admin.com/'
  },
  {
    title : t('view.dashboard.grids.githubRepo'),
    url : 'https://github.com/flipped-aurora/gin-vue-admin'
  },{
    title : t('view.dashboard.grids.pluginMarket'),
    url : 'https://plugin.gin-vue-admin.com/'
  },
  {
    title: 'Vue3',
    url : 'https://v3.cn.vuejs.org/guide/introduction.html'
  },

]
</script>

<style scoped lang="scss">

</style>
