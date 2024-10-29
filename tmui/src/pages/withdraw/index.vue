<template>
  <tm-app ref="app" darkColor="#000" :bgImg="bgimg" style="  background-size: cover;background-repeat: no-repeat;" >
    <!-- 顶部导航 -->
    <tm-navbar :title="''"  :height="50" :shadow="1" hide-home  :followTheme="true"  color="#181940" >
      <template v-slot:left>
        <view class="flex flex-center flex-row">
          <view>
            <tm-text class="text-weight-b">提币中心</tm-text>
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
    <tm-sheet style="background: linear-gradient(107deg,#6a39f2,#7e67ff);" class="round-4">
      <view class="flex flex-center ">
        <view>
          <view class="flex-center">
            <tm-text>待提金额</tm-text>
          </view>
          <view class="flex-center mt-10">
            <tm-text class="text-weight-b">{{ mainwith.withable }}</tm-text>
          </view>
        </view>
        <tm-divider vertical :height="80" class="mx-n20"></tm-divider>
        <view>
          <view class="flex-center">
            <tm-text>已提金额</tm-text>
          </view>
          <view class="flex-center mt-10">
            <tm-text class="text-weight-b">{{ mainwith.withed }}</tm-text>
          </view>
        </view>
      </view>
    </tm-sheet>
    <tm-sheet transprent :margin="[30, 20]" :padding="[10, 10]" class="  round-4" style="border: #6A39F2 1px solid;">
      <view class=""  v-for="(item,index) in items" :key="index">
        <view class="items itemround ">
          <view class="pa-20 flex flex-row-top-start">
<!--            <image :src="item.avatarurl"  style="width: 150rpx;height: 150rpx;border-radius: 50%" />-->
            <view class="ml-15 mt-5">
<!--              <tm-text class="mt-5" :fontSize="24">地址： {{item.address.substring(0, 8)+'***'+item.address.slice(-8)}}</tm-text>-->
              <tm-text class="mt-5" :fontSize="24">时间： {{item.CreatedAt.substring(0, 16)}}</tm-text>
              <tm-text class="mt-5" :fontSize="24">状态： {{item.status}}</tm-text>
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
import {getWithdrawListApi} from "@/api";
import {onShow} from "@dcloudio/uni-app";
const store = useTmpiniaStore()

onShow(()=>{
  getWithlist()
})
const items = ref(
    [
      {
        "ID": 1,
        "CreatedAt": "2024-10-29T20:08:16.649+08:00",
        "UpdatedAt": "2024-10-29T20:08:16.649+08:00",
        "address": "0x154b8BB871b72C501aE45765d945A16b8659F417",
        "amount": "10",
        "text": "",
        "status": "自动通过",
        "desc": "备注",
        "descnum": "0"
      }
    ]
)
const mainwith = ref(
    {
      "ID": 3,
      "CreatedAt": "2024-10-27T10:35:03.944+08:00",
      "UpdatedAt": "2024-10-27T10:35:03.944+08:00",
      "address": "0x154b8BB871b72C501aE45765d945A16b8659F417",
      "num": 0,
      "withable": "0",
      "withed": "0",
      "total": "0",
      "desc": "备注",
      "descnum": "0"
    }
)

//查询提币订单列表
const getWithlist = async () => {
  const params = {
    address: store.tmStore.userInfo.cliUser?.address,
    page: 1,
    pageSize: 10,
  }
  const res = await getWithdrawListApi(params)
  console.log(res)
  if (res.data.code === 0) {
    items.value = res.data.data.CliWithdraws
    mainwith.value = res.data.data.CliMainwiths
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