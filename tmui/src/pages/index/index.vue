<template>
  <tm-app ref="app" darkColor="#000" :bgImg="bgimg" >
    <!-- 顶部导航 -->
    <tm-navbar :title="''"  :height="60" :shadow="1" hide-home :followTheme="true"  transprent>
      <template v-slot:left>
        <view class="flex flex-center flex-row">
          <view class="mx-20">
            <tm-image src="/static/home/logo.8f451d87.png"  :width="60" :height="60"></tm-image>
          </view>
          <view>
            <tm-image src="/static/home/soll.7664f7cf.svg"  :width="260" :height="80"></tm-image>
          </view>

        </view>
      </template>
      <template v-slot:right>
        <view class="flex flex-center flex-row mr-20">
<!-- <view class="mr-10"><tm-text >{{store.tmStore.userInfo.wallet.address.substring(0,6)+'...'+store.tmStore.userInfo.wallet.address.slice(-6) }}</tm-text></view> -->
          <view @click=""  class="">
            <tm-button style="width: 90px;height: 35px; box-shadow: inset 0 0 8px 0 #6158ff;border: 1px solid #706da5;background: linear-gradient(0deg,#100842,#100842),linear-gradient(0deg,#706da5,#706da5);"
                       :block='true' @click="onChangeDark" :round="10" size="normal" >連接錢包
            </tm-button>
          </view>
          <tm-icon @click="onChangeDark()" :color="store.tmStore.dark?'white':''" _class="pl-15"  :font-size="40" name="tmicon-ios-sunny"></tm-icon>
        </view>
      </template>
    </tm-navbar>
    <view class="relative ">
      <view class="absolute flex-center" style="width:55%;height: 85%;top:50%;left:50%;border-radius: 999px;overflow: hidden;transform: translate(-50%,-50%) scaleX(1) scaleY(1) ">
        <video :src="globviedo" autoplay="true" muted="true" loop="true" playsinline></video>
      </view>
      <view class="flex-center" style="">
        <image  class="loader"  :src="roundimg"></image>
      </view>
    </view>
    <tm-sheet darkBgColor="#391c8c66" style="transform: translateY(-30px);border-width: 1px;border-color: #fff3;" class="" :round="4" :margin="[20,20]">
      <view class="flex flex-row-center-between">
        <view class="flex">
          <view class="flex-col-center-center mx-10">
            <svg xmlns="http://www.w3.org/2000/svg" width="12" height="13" viewBox="0 0 12 13" fill="none"><path d="M12 10.5C12 11.8807 9.31371 13 6 13C2.68629 13 0 11.8807 0 10.5C0 10.0447 0 8.62502 0 8.62502H12C12 8.62502 12 9.1193 12 10.5Z" fill="#9281FF"></path><ellipse cx="6" cy="8.62499" rx="6" ry="2.5" fill="black"></ellipse><path d="M12 7.5C12 8.88072 9.31371 10 6 10C2.68629 10 0 8.88072 0 7.5C0 7.04464 0 5.625 0 5.625H12C12 5.625 12 6.11929 12 7.5Z" fill="#9281FF"></path><ellipse cx="6" cy="5.625" rx="6" ry="2.5" fill="black"></ellipse><path d="M12 4.37502C12 5.75573 9.31371 6.87502 6 6.87502C2.68629 6.87502 0 5.75573 0 4.37502C0 3.91966 0 2.50002 0 2.50002H12C12 2.50002 12 2.9943 12 4.37502Z" fill="#9281FF"></path><ellipse cx="6" cy="2.5" rx="6" ry="2.5" fill="#6754E0"></ellipse></svg>
          </view>
          <tm-text  class="flex-col-center-center ">累積獎勵 (SOLG)</tm-text>
        </view>
        <view class="flex py-10  px-20 round-10" style="background-color: #d9d9d91a">
          <tm-text class="flex-col-center-center  ">Depin 獎勵</tm-text>
          <view class="flex-col-center-center ml-10 ">
            <svg xmlns="http://www.w3.org/2000/svg" width="5" height="9" viewBox="0 0 5 9" fill="none"><path d="M0.939362 1L4.12134 4.18198C4.3166 4.37724 4.3166 4.69383 4.12134 4.88909L0.939362 8.07107" stroke="white" stroke-linecap="round"></path></svg>
          </view>
        </view>
      </view>
      <view class="flex flex-row-center-between">
        <view class="ml-n15">
          <tm-text :fontSize="60">0</tm-text>
          <tm-text color="green">+0</tm-text>
        </view>
        <view>
          <tm-image :src="homecoin"  :width="200" :height="200"></tm-image>
        </view>
      </view>
    </tm-sheet>
    <tm-sheet class="relative " darkBgColor="#4b21e5" style="background-image: url('@/static/home/home_bg_5.png'); transform: translateY(-30px);border-width: 1px;border-color: #fff3;"  :round="4" :margin="[20,20]">
      6666666

<!--      <image :src="homebgimg"  style="width: 100%;height: 500rpx;"></image>-->

    </tm-sheet>
    <view>
      <tm-button style="  box-shadow: inset 0 0 8px 0 #6158ff;border: 1px solid #706da5;background: linear-gradient(0deg,#100842,#100842),linear-gradient(0deg,#706da5,#706da5);" :block='true' @click="onChangeDark">嗨~TMUI</tm-button>
    </view>
  </tm-app>
</template>
<script lang="ts" setup>
	import { ref } from "vue"
  import bgimg from "@/static/home/bg_6.3d11a96f.png"
  import roundimg from "@/static/home/borderound.svg"
  import globviedo from "@/static/home/global.mp4"
  import homecoin from '@/static/home/home_coin.svg'
  import homebgimg from "@/static/home/home_bg_5.png"
  import { useTmpiniaStore } from '@/tmui/tool/lib/tmpinia'
  import {onShow} from "@dcloudio/uni-app";
  const store = useTmpiniaStore()

  onShow(()=>{

  })


  function onChangeDark() {
    console.log('Set dark')
    if( store.tmStore.dark){
      store.setTmVuetifyDark(false)
    }else{
      store.setTmVuetifyDark(true)
    }
  }
</script>
<style>

.btn-gradient {
  box-shadow: inset 0 0 8px 0 #6158ff;
  border: 1px solid #706da5;
  background: linear-gradient(0deg,#100842,#100842),linear-gradient(0deg,#706da5,#706da5);
}
.bg-primary{
  text-transform: uppercase;
  --tw-bg-opacity: 1;
  background-color: rgb(25 8 75/var(--tw-bg-opacity));
}
.loader{
  animation: rotation 30s linear infinite;
}
@keyframes rotation {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(1turn);
  }
}
</style>