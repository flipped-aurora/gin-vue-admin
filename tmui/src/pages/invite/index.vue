<template>
  <tm-app ref="app" darkColor="#000" :bgImg="bgimg" style="  background-size: cover;background-repeat: no-repeat;" >
    <!-- 顶部导航 -->
    <tm-navbar :title="''"  :height="50" :shadow="1" hide-home  :followTheme="true"  color="#181940" >
      <template v-slot:left>
        <view class="flex flex-center flex-row">
          <view>
            <tm-text class="text-weight-b">邀请中心</tm-text>
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
            <tm-text>直推成员</tm-text>
          </view>
          <view class="flex-center mt-10">
            <tm-text class="text-weight-b">{{ myuser.pullnum }}</tm-text>
          </view>
        </view>
        <tm-divider vertical :height="80" class="mx-n20"></tm-divider>
        <view>
          <view class="flex-center">
            <tm-text>团队成员</tm-text>
          </view>
          <view class="flex-center mt-10">
            <tm-text class="text-weight-b">{{myuser.teamnum}}</tm-text>
          </view>
        </view>
      </view>
    </tm-sheet>
    <view class="flex-center mx-30">
      <view style="width: 50%">
        <tm-button @click="btn1" :color="btncolor" class="round-l-4 round-r-0" block>直推数据</tm-button>
      </view>
      <view style="width: 50%">
        <tm-button @click="btn2"  :color="btncolor2"  class="round-r-4 round-l-0" block>团队数据</tm-button>
      </view>
    </view>
    <tm-sheet v-if="show" transprent :margin="[30, 20]" :padding="[10, 10]" class="  round-4" style="border: #6A39F2 1px solid;">
      <view class=""  v-for="(item,index) in items" :key="index">
        <view class="items itemround ">
          <view class="pa-20 flex flex-row-top-start">
            <image :src="item.description"  style="width: 120rpx;height: 120rpx;border-radius: 50%" />
            <view class="ml-20 mt-5">

              <tm-text class="mt-5" :fontSize="24">昵称： {{item.father}}</tm-text>
              <tm-text class="mt-5" :fontSize="24">地址： {{item.address.substring(0, 8)+'***'+item.address.slice(-8)}}</tm-text>
              <tm-text class="mt-5" :fontSize="24">时间： {{item.CreatedAt.substring(0, 16)}}</tm-text>
            </view>
          </view>
          <view v-show ="index<items.length-1">
            <tm-divider  :height="1" color="#6A39F2"></tm-divider>
          </view>
        </view>
      </view>
    </tm-sheet>
    <tm-sheet v-if="!show" transprent :margin="[30, 20]" :padding="[10, 10]" class="  round-4" style="border: #6A39F2 1px solid;">
      <view class=""  v-for="(item,index) in team" :key="index">
        <view class="items itemround ">
          <view class="pa-20 flex flex-row-top-start">
            <image :src="item.description"  style="width: 120rpx;height: 120rpx;border-radius: 50%" />
            <view class="ml-20 mt-5">

              <tm-text class="mt-5" :fontSize="24">昵称： {{item.father}}</tm-text>
              <tm-text class="mt-5" :fontSize="24">地址： {{item.address.substring(0, 8)+'***'+item.address.slice(-8)}}</tm-text>
              <tm-text class="mt-5" :fontSize="24">时间： {{item.CreatedAt.substring(0, 16)}}</tm-text>
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
  import {getInviteListApi} from "@/api";
  import {onShow} from "@dcloudio/uni-app";
  const store = useTmpiniaStore()

  onShow(()=>{
    getInviteList(1)
  })

  const items = ref(
      [
        {
          "ID": 4,
          "CreatedAt": "2024-10-29T15:38:45.431+08:00",
          "UpdatedAt": "2024-10-29T17:59:09.605+08:00",
          "address": "001",
          "father": "尹方财",
          "invite": "0x154b8BB871b72C501aE45765d945A16b8659F417",
          "leftval": 5,
          "rightval": 8,
          "level": 3,
          "vip": 0,
          "description": "https://api.multiavatar.com/9.png"
        }
      ]
  )
  const team = ref(
      [
        {
          "ID": 4,
          "CreatedAt": "2024-10-29T15:38:45.431+08:00",
          "UpdatedAt": "2024-10-29T17:59:09.605+08:00",
          "address": "001",
          "father": "尹方财",
          "invite": "0x154b8BB871b72C501aE45765d945A16b8659F417",
          "leftval": 5,
          "rightval": 8,
          "level": 3,
          "vip": 0,
          "description": "https://api.multiavatar.com/9.png"
        }
      ]
  )
  const myuser = ref({
    "ID": 4,
    "CreatedAt": "2024-10-27T10:35:03.944+08:00",
    "UpdatedAt": "2024-10-29T17:59:09.613+08:00",
    "address": "0x154b8BB871b72C501aE45765d945A16b8659F417",
    "nickname": "申钰伦",
    "avatarurl": "https://api.multiavatar.com/91.png",
    "parent": "root",
    "pullnum": 2,
    "teamnum": 12,
    "mypull": [
      "001"
    ],
    "myUp": null,
    "vips": null,
    "status": "正常",
    "desc": "文本备注",
    "desnum": "0"
  })

  //查询邀请列表
  const getInviteList = async (page: number) => {
    const params = {
      address: store.tmStore.userInfo.wallet?.address,
      page: page,
      size: 10
    }
    const res = await getInviteListApi(params)
    if (res.data.code === 0) {
      console.log(res)
      items.value = res.data.data.Pulltrees
      team.value = res.data.data.Teamtrees
      myuser.value = res.data.data.Cliuser
    }
  }
  const show = ref(true)
  const btncolor2 = ref('rgba(237,237,237,0.66)')
  const btncolor = ref('#6A39F2')

  const activeColor = '#6A39F2'
  const inactiveColor = 'rgba(237,237,237,0.66)'

  const toggleBtn = (isActive: boolean) => {
    show.value = isActive
    btncolor2.value = !isActive ? activeColor : inactiveColor
    btncolor.value = !isActive ? inactiveColor : activeColor
  }

  const btn1 = () => toggleBtn(true)
  const btn2 = () => toggleBtn(false)

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