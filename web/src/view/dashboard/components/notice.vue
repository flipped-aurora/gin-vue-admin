<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <el-scrollbar>
    <div
      v-for="(item, index) in notices"
      :key="index"
      class="flex items-center mb-1.5 gap-3"
    >
      <el-tag :type="item.type" size="small">
        {{ item.typeTitle }}
      </el-tag>
      <el-tooltip effect="light" :content="item.title" placement="top">
        <div class="text-xs text-gray-700 dark:text-gray-300 line-clamp-1">
          {{ item.title }}
        </div>
      </el-tooltip>
    </div>
  </el-scrollbar>
</template>

<script setup>
  const notices = [
    {
      type: 'primary',
      typeTitle: '公告',
      title: '授权费将在从六月一日起结束第一价格梯度，进入第二价格梯度。'
    },
    {
      type: 'success',
      typeTitle: '通知',
      title: '授权后将进入专属飞书群，获取官方辅助。'
    },
    {
      type: 'warning',
      typeTitle: '警告',
      title: '授权可获得插件市场极大优惠价格。'
    },
    {
      type: 'danger',
      typeTitle: '违规',
      title: '未授权商用将有可能被资源采集工具爬取并追责。'
    },
    {
      type: 'info',
      typeTitle: '信息',
      title: '再次感谢您对开源事业的支持'
    },
    {
      type: 'primary',
      typeTitle: '公告',
      title: '让创意更有价值。'
    },
    {
      type: 'success',
      typeTitle: '通知',
      title: '让劳动更有意义。'
    },
    {
      type: 'warning',
      typeTitle: '警告',
      title: '让思维更有深度。'
    },
    {
      type: 'danger',
      typeTitle: '错误',
      title: '让生活更有趣味。'
    },
    {
      type: 'info',
      typeTitle: '信息',
      title: '让公司更有活力。'
    }
  ]
</script>

<style scoped lang="scss"></style>
