<script setup lang="ts">

import {computed, ref} from "vue";
import {language} from "../tmui/tool/lib/language";

const props = defineProps({
  showleft:{
    type:Boolean,
    default:false
  }
})
const emit = defineEmits(['closedraw']);
const updateValue = () => {
  emit('closedraw', 'new value');
};
const closedraw = () => {
  emit('closedraw', !props.showleft);
};
function toLinks(url : string) {
  closedraw()
  let urls = '/pages/'+url
  console.log(urls)
  uni.navigateTo({ url: urls })
}
function seLocal() {
  if (language('language') == 'English-US') {
    uni.setLocale('zh-Hans')
  } else {
    uni.setLocale('en')
  }
}
</script>

<template>
  <tm-drawer ok-text="" :hideHeader="true" transprent :width="750" :show="props.showleft" :mask="false"  placement="right" :round="0">
    <view style="height: 45px" class="flex-row-center-end  pr-20">
      <tm-icon _class="pl-15"  :font-size="40" name="tmicon-menu" @click="closedraw"></tm-icon>
    </view>
    <view style="height: 95vh;background-color: rgb(5 5 5 / 80%)">
      <view class="flex-col flex-col-center-center">
        <view style="height: 100rpx"></view>
        <tm-text class="text-weight-b text-size-g  my-20" color="white" @click="toLinks('index/index')">回到首页</tm-text>
        <tm-text class="text-weight-b text-size-g  my-20" color="white"  @click="toLinks('team/index')">团队中心</tm-text>
        <tm-text class="text-weight-b text-size-g  my-20" color="white"  @click="toLinks('invite/index')">邀请中心</tm-text>
        <tm-text class="text-weight-b text-size-g  my-20" color="white" @click="toLinks('datas/index')">结算中心</tm-text>
<!--        <tm-text class="text-weight-b text-size-g my-20" @click="toLinks('/pages/rank/index')">{{ language('index.com.fd') }}</tm-text>-->
        <tm-text class="text-weight-b text-size-g  my-20" color="white"  @click="toLinks('withdraw/index')">提币中心</tm-text>
      </view>
    </view>
  </tm-drawer>
</template>

<style scoped>

</style>