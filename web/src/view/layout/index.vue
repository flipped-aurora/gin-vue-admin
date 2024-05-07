<template>
  <div class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-700 w-screen h-screen">
    <gva-header></gva-header>
    <div class="flex flex-row w-full gva-container pt-16">
      <gva-aside v-model:collapse="isCollapse" />
      <div class="flex-1  w-0 h-full">
        <history-component></history-component>
        <div class="overflow-auto gva-body-container">
          <div class="gva-body-h">
            <router-view v-if="reloadFlag" v-slot="{ Component }">
              <div id="gva-base-load-dom" class="bg-gray-50 dark:bg-gray-700 p-1">
                <transition mode="out-in" name="el-fade-in-linear">
                  <keep-alive :include="routerStore.keepAliveRouters">
                    <component :is="Component" />
                  </keep-alive>
                </transition>
              </div>
            </router-view>
          </div>
        <BottomInfo />
        </div>
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
import { emitter } from "@/utils/bus.js";
import { computed, ref, onMounted, nextTick} from "vue";
import { useRouter, useRoute } from "vue-router";
import { useRouterStore } from "@/pinia/modules/router";
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




</script>

<style lang="scss">
.gva-container{
  height: calc(100% - 4rem);
}
.gva-body-container{
  height: calc(100% - 3rem);
}

.gva-body-h{
  min-height: calc(100% - 3.5rem);
}
</style>
