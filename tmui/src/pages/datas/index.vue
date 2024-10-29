<template>
  <tm-app ref="app" darkColor="#000" :bgImg="bgimg" style="  background-size: cover;background-repeat: no-repeat;" >
    <!-- 顶部导航 -->
    <tm-navbar :title="''"  :height="50" :shadow="1" hide-home  :followTheme="true"  color="#181940" >
      <template v-slot:left>
        <view class="flex flex-center flex-row">
          <view>
            <tm-text class="text-weight-b">结算中心</tm-text>
            <!--            <tm-image src="/static/home/soll.7664f7cf.svg"  :width="260" :height="80"></tm-image>-->
          </view>

        </view>
      </template>
      <template v-slot:right>
        <view class="flex flex-center flex-row mr-20">
          <!-- <view class="mr-10"><tm-text >{{store.tmStore.userInfo.wallet.address.substring(0,6)+'...'+store.tmStore.userInfo.wallet.address.slice(-6) }}</tm-text></view> -->
          <view @click=""  class="">
            <tm-button style="width: 150px;height: 35px; box-shadow: inset 0 0 8px 0 #6158ff;border: 1px solid #706da5;background: linear-gradient(0deg,#100842,#100842),linear-gradient(0deg,#706da5,#706da5);"
                       :block='true' @click="" :round="10" size="normal" >{{ store.tmStore.userInfo.wallet?.address.substring(0,6) +'...'+ store.tmStore.userInfo.wallet?.address.slice(-6) }}
            </tm-button>
          </view>
          <tm-icon @click="toggleDrawer" :color="store.tmStore.dark?'white':''" _class="pl-15"  :font-size="40" name="tmicon-menu"></tm-icon>
        </view>
      </template>
    </tm-navbar>
    <view class="flex-center pa-20 mt-20 mx-30 round-4" style="background: linear-gradient(107deg,#6a39f2,#7e67ff);">
      <view style="width: 50%" class="flex-center">
        <tm-text>累计收益</tm-text>
      </view>
      <view style="width: 50%" class="flex-center">
        <tm-text>{{ mainpro.amount }}</tm-text>
      </view>
    </view>
    <tm-sheet style="background: linear-gradient(107deg,#6a39f2,#7e67ff);" class="round-4">
      <view class="flex flex-center ">
        <view>
          <view class="flex-center">
            <tm-text>静态收益</tm-text>
          </view>
          <view class="flex-center mt-10">
            <tm-text class="text-weight-b">{{ mainpro.static }}</tm-text>
          </view>
        </view>
        <tm-divider vertical :height="80" class="mx-n20"></tm-divider>
        <view>
          <view class="flex-center">
            <tm-text>直推收益</tm-text>
          </view>
          <view class="flex-center mt-10">
            <tm-text class="text-weight-b">{{ mainpro.pull }}</tm-text>
          </view>
        </view>
      </view>
      <view class="flex flex-center mt-20 ">
        <view>
          <view class="flex-center">
            <tm-text>间推收益</tm-text>
          </view>
          <view class="flex-center mt-10">
            <tm-text class="text-weight-b">{{ mainpro.indirect }}</tm-text>
          </view>
        </view>
        <tm-divider vertical :height="80" class="mx-n20"></tm-divider>
        <view>
          <view class="flex-center">
            <tm-text>团队收益</tm-text>
          </view>
          <view class="flex-center mt-10">
            <tm-text class="text-weight-b">{{ mainpro.team }}</tm-text>
          </view>
        </view>
      </view>
    </tm-sheet>
    <tm-sheet transprent :margin="[30, 20]" :padding="[10, 10]" class="  round-4" style="border: #6A39F2 1px solid;">
      <view class=""  v-for="(item,index) in items" :key="index">
        <view class="items itemround ">
          <view class="pa-20 flex flex-row-top-start">
            <view class="ml-15 mt-5">
              <tm-text class="mt-5" :fontSize="24">时间： {{item.CreatedAt}}</tm-text>
<!--              <tm-text class="mt-5" :fontSize="24">地址： {{item.address.substring(0, 8)+'***'+item.address.slice(-8)}}</tm-text>-->
              <tm-text class="mt-5" :fontSize="24">类型： {{ item.text }}</tm-text>
              <tm-text class="mt-5" :fontSize="24">金额： {{ item.amount }}</tm-text>
            </view>
          </view>
          <view v-show ="index<items.length-1">
            <tm-divider  :height="1" color="#6A39F2"></tm-divider>
          </view>
        </view>
      </view>
    </tm-sheet>
  </tm-app>
  <draw :showleft="showdraw"  @closedraw="handleClosedraw"></draw>
</template>
<script setup lang="ts">
import bgimg from "@/static/home/home-bg.png";
import {useTmpiniaStore} from "@/tmui/tool/lib/tmpinia";
import {ref} from "vue";
import Draw from "@/components/draw.vue";
import {onShow} from "@dcloudio/uni-app";
import {getProfitListApi, getProfitMainApi} from "@/api";
const store = useTmpiniaStore()

onShow(()=>{
  getProfitMain()
  getProfitList(1)
})

const items = ref(
    [
      {
        "ID": 1,
        "CreatedAt": "2024-10-29T14:54:44.193+08:00",
        "UpdatedAt": "2024-10-29T14:54:44.193+08:00",
        "address": "0x154b8BB871b72C501aE45765d945A16b8659F417",
        "amount": "10",
        "text": "静态收益",
        "status": "正常",
        "desc": "备注",
        "descnum": "0"
      }
    ]
)
const mainpro = ref(

  {
    "ID": 3,
      "CreatedAt": "2024-10-27T10:35:03.944+08:00",
      "UpdatedAt": "2024-10-27T10:35:03.944+08:00",
      "address": "0x154b8BB871b72C501aE45765d945A16b8659F417",
      "num": 0,
      "static": "0",
      "pull": "0",
      "indirect": "0",
      "team": "0",
      "amount": "0",
      "desc": "",
      "descnum": "0"
  }
)
//查询用户结算总表
const getProfitMain = async () => {
  let params = {
    address: store.tmStore.userInfo.wallet?.address,
    page:1,
    pageSize:10
  }
  console.log(params)
  const res = await getProfitMainApi(params)
  console.log(res)
  if (res.data.code === 0) {
    mainpro.value = res.data.data.list[0]
  }
}
//查询结算详情
const getProfitList = async (page:number) => {
  let params = {
    address: store.tmStore.userInfo.wallet?.address,
    page:page,
    pageSize:10
  }
  console.log(params)
  const res = await getProfitListApi(params)
  console.log(res)
  if (res.data.code === 0) {
    items.value = res.data.data.list
  }
}


const showdraw = ref(false)
const toggleDrawer = () => {

  showdraw.value = !showdraw.value;
  console.log('draw toggle', showdraw.value)
};
const handleClosedraw = (event: boolean) => {
  showdraw.value = event;
  console.log(uni.getLocale())
  console.log('draw closed', showdraw.value)
};

</script>
<style scoped>

</style>