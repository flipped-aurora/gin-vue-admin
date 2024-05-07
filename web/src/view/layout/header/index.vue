<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/7
!-->

<template>
  <div class="fixed top-0 left-0 right-0 z-10 h-16 bg-white text-slate-700 dark:text-slate-300  dark:bg-slate-900 shadow dark:shadow-gray-700 flex items-center px-2">
    <div
      class="flex items-center "
      :class="isMobile ? 'mx-2' : 'mx-8'"
    >
      <img
        alt
        class="h-12 bg-white rounded-full"
        :src="$GIN_VUE_ADMIN.appLogo"
      >
      <div
        v-if="!isMobile"
        class="inline-flex font-bold text-2xl ml-2"
      >
        {{ $GIN_VUE_ADMIN.appName }}
      </div>


      <el-breadcrumb v-show="!isMobile" class="ml-4">
        <el-breadcrumb-item
          v-for="item in matched.slice(1, matched.length)"
          :key="item.path"
        >
          {{
            fmtTitle(item.meta.title, route)
          }}
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>

    <div class="ml-auto flex items-center">
      <tools />
      <el-dropdown>
        <div
          class="flex justify-center items-center h-full w-full"
        >
          <span
            class="cursor-pointer flex justify-center items-center text-black dark:text-gray-100"
          >
            <CustomPic />
            <span
              v-show="!isMobile"
              style="margin-left: 5px"
            >{{ userStore.userInfo.nickName }}</span>
            <el-icon>
              <arrow-down />
            </el-icon>
          </span>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>
              <span class="font-bold">
                当前角色：{{
                  userStore.userInfo.authority.authorityName
                }}
              </span>
            </el-dropdown-item>
            <template v-if="userStore.userInfo.authorities">
              <el-dropdown-item
                v-for="item in userStore.userInfo.authorities.filter(
                  (i) =>
                    i.authorityId !==
                    userStore.userInfo.authorityId
                )"
                :key="item.authorityId"
                @click="changeUserAuth(item.authorityId)"
              >
                <span>
                  切换为：{{ item.authorityName }}
                </span>
              </el-dropdown-item>
            </template>
            <el-dropdown-item icon="avatar">
              <div
                class="command-box"
                style="display: flex"
                @click="handleCommand"
              >
                <div>指令菜单</div>
                <div style="margin-left: 8px">
                  <span class="button">{{ first }}</span>
                  +
                  <span class="button">K</span>
                </div>
              </div>
            </el-dropdown-item>
            <el-dropdown-item icon="avatar" @click="toPerson">
              个人信息
            </el-dropdown-item>
            <el-dropdown-item
              icon="reading-lamp"
              @click="userStore.LoginOut"
            >
              登 出
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
    <command-menu ref="command" />
  </div>
</template>

<script setup>
import  tools from "./tools.vue"
import CustomPic from '@/components/customPic/index.vue'
import CommandMenu from "@/components/commandMenu/index.vue";
import { useUserStore } from "@/pinia/modules/user";
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from "@/pinia"
import { storeToRefs } from "pinia"
import { computed, ref } from 'vue'
import { setUserAuthority } from '@/api/user'
import { fmtTitle } from "@/utils/fmtRouterTitle";
const userStore = useUserStore();
const router = useRouter()
const route = useRoute()
const appStore = useAppStore()
const { device } = storeToRefs(appStore)
const isMobile = computed(() =>{
  return device.value  === 'mobile'
})
const toPerson = () => {
  router.push({ name: "person" });
};
const matched = computed(() => route.meta.matched);

const first = ref("");
const command = ref();
const initPage = () => {
  // 判断当前用户的操作系统
  if (window.localStorage.getItem("osType") === "WIN") {
    first.value = "Ctrl";
  } else {
    first.value = "⌘";
  }
  // 当用户同时按下ctrl和k键的时候
  const handleKeyDown = (e) => {
    if (e.ctrlKey && e.key === "k") {
      // 阻止浏览器默认事件
      e.preventDefault();
      handleCommand();
    }
  };
  window.addEventListener("keydown", handleKeyDown);
};

initPage();

const handleCommand = () => {
  command.value.open();
};


const changeUserAuth = async (id) => {
  const res = await setUserAuthority({
    authorityId: id,
  });
  if (res.code === 0) {
    window.sessionStorage.setItem("needCloseAll", "true");
    window.location.reload();
  }
};


</script>

<style scoped lang="scss">

</style>
