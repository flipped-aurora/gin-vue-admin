<template>
  <div :style="{ background: theme.background }">
    <el-scrollbar style="height: calc(100vh - 60px)">
      <transition
          :duration="{ enter: 800, leave: 100 }"
          mode="out-in"
          name="el-fade-in-linear"
      >
        <el-menu
            :collapse="isCollapse"
            :collapse-transition="false"
            :default-active="active"
            :background-color="theme.background"
            :active-text-color="theme.active"
            class="el-menu-vertical"
            unique-opened
            @select="selectMenuItem"
        >
          <template v-for="item in routerStore.asyncRouters[0].children">
            <aside-component
                v-if="!item.hidden"
                :key="item.name"
                :is-collapse="isCollapse"
                :router-info="item"
                :theme="theme"
            />
          </template>
        </el-menu>
      </transition>
    </el-scrollbar>
  </div>
</template>

<script setup>
import AsideComponent from "@/view/layout/aside/asideComponent/index.vue";
import { ref, inject, watchEffect } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useRouterStore } from "@/pinia/modules/router";

defineOptions({
  name: "Aside",
});
const route = useRoute();
const router = useRouter();
const routerStore = useRouterStore();
const isCollapse = defineModel("collapse", {
  type: Boolean,
});

const theme = inject("theme", {
  background: "#fff",
  activeBackground: "var(--el-color-primary)",
  activeText: "#fff",
  normalText: "#333",
  hoverBackground: "rgba(64, 158, 255, 0.08)",
  hoverText: "#333",
  subBackgroundActive: "var(--el-color-primary)",
});
const active = ref("");

watchEffect(() => {
  active.value = route.meta.activeName || route.name;
});

const selectMenuItem = (index, _, ele, aaa) => {
  const query = {};
  const params = {};
  routerStore.routeMap[index]?.parameters &&
  routerStore.routeMap[index]?.parameters.forEach((item) => {
    if (item.type === "query") {
      query[item.key] = item.value;
    } else {
      params[item.key] = item.value;
    }
  });
  if (index === route.name) return;
  if (index.indexOf("http://") > -1 || index.indexOf("https://") > -1) {
    window.open(index);
  } else {
    router.push({ name: index, query, params });
  }
};
</script>

<style lang="scss">
.el-sub-menu__title:hover,
.el-menu-item:hover {
  background: transparent;
}

.el-scrollbar {
  .el-scrollbar__view {
    height: 100%;
  }
}
.menu-info {
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    display: table-cell;
    vertical-align: middle;
  }
}
</style>
