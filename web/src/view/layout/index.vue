<template>
  <div class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-700 w-screen h-screen">
    <gva-header></gva-header>
    <div class="flex flex-row w-full gva-container  pt-16">
      <gva-aside v-model:collapse="isCollapse" />
      <div class="flex-1  w-0 h-full">
        <history-component></history-component>
        <router-view v-if="reloadFlag" v-slot="{ Component }">
          <div id="gva-base-load-dom" class="bg-gray-50 dark:bg-gray-700 p-1 gva-container2 overflow-auto">
            <transition mode="out-in" name="el-fade-in-linear">
              <keep-alive :include="routerStore.keepAliveRouters">
                <component :is="Component" />
              </keep-alive>
            </transition>
          </div>
        </router-view>
        <BottomInfo />
      </div>
    </div>
  </div>
</template>

<script setup>
import GvaAside from "@/view/layout/aside/index.vue";
import GvaHeader from "@/view/layout/header/index.vue";
import useResponsive  from "@/hooks/responsive";
import HistoryComponent from "@/view/layout/aside/historyComponent/history.vue";
import BottomInfo from "@/view/layout/bottomInfo/bottomInfo.vue";
import { setUserAuthority } from "@/api/user";
import { emitter } from "@/utils/bus.js";
import { computed, ref, onMounted, nextTick, provide, watchEffect } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useRouterStore } from "@/pinia/modules/router";
import { fmtTitle } from "@/utils/fmtRouterTitle";
import { useUserStore } from "@/pinia/modules/user";

defineOptions({
  name: "GvaLayout",
});

useResponsive(true)

const router = useRouter();
const route = useRoute();
const routerStore = useRouterStore();






onMounted(() => {
  // 挂载一些通用的事件
  emitter.on("reload", reload);
  if (userStore.loadingInstance) {
    userStore.loadingInstance.close();
  }
});

const userStore = useUserStore();


const matched = computed(() => route.meta.matched);

const changeUserAuth = async (id) => {
  const res = await setUserAuthority({
    authorityId: id,
  });
  if (res.code === 0) {
    window.sessionStorage.setItem("needCloseAll", "true");
    window.location.reload();
  }
};

const reloadFlag = ref(true);
let reloadTimer = null;
const reload = async () => {
  if (reloadTimer) {
    window.clearTimeout(reloadTimer);
  }
  reloadTimer = window.setTimeout(async () => {
    if (route.meta.keepAlive) {
      reloadFlag.value = false;
      await nextTick();
      reloadFlag.value = true;
    } else {
      const title = route.meta.title;
      router.push({ name: "Reload", params: { title } });
    }
  }, 400);
};

const isShadowBg = ref(false);
const totalCollapse = () => {
  isCollapse.value = !isCollapse.value;
  isSider.value = !isCollapse.value;
  if(isMobile.value){
    isShadowBg.value = !isCollapse.value;
  }
};


const changeShadow = () => {
  isShadowBg.value = !isShadowBg.value;
  isSider.value = !!isCollapse.value;
  totalCollapse();
};


</script>

<style lang="scss">
.gva-container{
  height: calc(100% - 4rem);
}
.gva-container2{
  height: calc(100% - 8rem);
}
</style>
