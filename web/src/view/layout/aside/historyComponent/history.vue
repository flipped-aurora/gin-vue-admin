<template>
  <div class="router-history">
    <el-tabs
      :closable="!(historys.length==1&&this.$route.name==defaultRouter)"
      @contextmenu.prevent.native="openContextMenu($event)"
      @tab-click="changeTab"
      @tab-remove="removeTab"
      type="card"
      v-model="activeValue"
    >
      <el-tab-pane
        :key="item.name + JSON.stringify(item.query)+JSON.stringify(item.params)"
        :label="item.meta.title"
        :name="item.name + JSON.stringify(item.query)+JSON.stringify(item.params)"
        :tab="item"
        v-for="item in historys"
      ></el-tab-pane>
    </el-tabs>

    <!--自定义右键菜单html代码-->
    <ul :style="{left:left+'px',top:top+'px'}" class="contextmenu" v-show="contextMenuVisible">
      <li @click="closeAll">关闭所有</li>
      <li @click="closeLeft">关闭左侧</li>
      <li @click="closeRight">关闭右侧</li>
      <li @click="closeOther">关闭其他</li>
    </ul>
  </div>
</template>
<script>
import { mapGetters } from "vuex";

export default {
  name: "HistoryComponent",
  data() {
    return {
      historys: [],
      activeValue: "",
      contextMenuVisible: false,
      left: 0,
      top: 0,
      isCollapse: false,
      isMobile: false,
      rightActive: ""
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo"]),
    defaultRouter() {
      return this.userInfo.authority.defaultRouter;
    }
  },
  created() {
    this.$bus.on("mobile", isMobile => {
      this.isMobile = isMobile;
    });
    this.$bus.on("collapse", isCollapse => {
      this.isCollapse = isCollapse;
    });
    const initHistorys = [
      {
        name: this.defaultRouter,
        meta: {
          title: "首页"
        },
        query: {},
        params: {}
      }
    ];
    this.historys =
      JSON.parse(sessionStorage.getItem("historys")) || initHistorys;
      if(!window.sessionStorage.getItem("activeValue")){
        this.activeValue = this.$route.name + JSON.stringify(this.$route.query)+JSON.stringify(this.$route.params)
      }else{
        this.activeValue = window.sessionStorage.getItem("activeValue");
      }
    this.setTab(this.$route);
  },

  beforeDestroy() {
    this.$bus.off("collapse");
    this.$bus.off("mobile");
  },
  methods: {
    openContextMenu(e) {
      if (this.historys.length == 1 && this.$route.name == this.defaultRouter) {
        return false;
      }
      if (e.srcElement.id) {
        this.contextMenuVisible = true;
        let width;
        if (this.isCollapse) {
          width = 54;
        } else {
          width = 220;
        }
        if (this.isMobile) {
          width = 0;
        }
        this.left = e.clientX - width;
        this.top = e.clientY + 10;
        this.rightActive = e.srcElement.id.split("-")[1];
      }
    },
    closeAll() {
      this.historys = [
        {
          name: this.defaultRouter,
          meta: {
            title: "首页"
          },
          query: {},
          params: {}
        }
      ];
      this.$router.push({ name: this.defaultRouter });
      this.contextMenuVisible = false;
      sessionStorage.setItem("historys", JSON.stringify(this.historys));
    },
    closeLeft() {
      let right;
      const rightIndex = this.historys.findIndex(item => {
        if (
          item.name +
            JSON.stringify(item.query) +
            JSON.stringify(item.params) ==
          this.rightActive
        ) {
          right = item;
        }
        return (
          item.name +
            JSON.stringify(item.query) +
            JSON.stringify(item.params) ==
          this.rightActive
        );
      });
      const activeIndex = this.historys.findIndex(
        item =>
          item.name +
            JSON.stringify(item.query) +
            JSON.stringify(item.params) ==
          this.activeValue
      );
      this.historys.splice(0, rightIndex);
      if (rightIndex > activeIndex) {
        this.$router.push(right);
      }
      sessionStorage.setItem("historys", JSON.stringify(this.historys));
    },
    closeRight() {
      let right;
      const leftIndex = this.historys.findIndex(item => {
        if (
          item.name +
            JSON.stringify(item.query) +
            JSON.stringify(item.params) ==
          this.rightActive
        ) {
          right = item;
        }
        return (
          item.name +
            JSON.stringify(item.query) +
            JSON.stringify(item.params) ==
          this.rightActive
        );
      });
      const activeIndex = this.historys.findIndex(
        item =>
          item.name +
            JSON.stringify(item.query) +
            JSON.stringify(item.params) ==
          this.activeValue
      );
      this.historys.splice(leftIndex + 1, this.historys.length);
      if (leftIndex < activeIndex) {
        this.$router.push(right);
      }
      sessionStorage.setItem("historys", JSON.stringify(this.historys));
    },
    closeOther() {
      let right;
      this.historys = this.historys.filter(item => {
        if (
          item.name +
            JSON.stringify(item.query) +
            JSON.stringify(item.params) ==
          this.rightActive
        ) {
          right = item;
        }
        return (
          item.name +
            JSON.stringify(item.query) +
            JSON.stringify(item.params) ==
          this.rightActive
        );
      });
      this.$router.push(right);
      sessionStorage.setItem("historys", JSON.stringify(this.historys));
    },
    isSame(route1, route2) {
      if (route1.name != route2.name) {
        return false;
      }
      for (let key in route1.query) {
        if (route1.query[key] != route2.query[key]) {
          return false;
        }
      }
      for (let key in route1.params) {
        if (route1.params[key] != route2.params[key]) {
          return false;
        }
      }
      return true;
    },
    setTab(route) {
      if (!this.historys.some(item => this.isSame(item, route))) {
        const obj = {};
        obj.name = route.name;
        obj.meta = route.meta;
        obj.query = route.query;
        obj.params = route.params;
        this.historys.push(obj);
      }
      window.sessionStorage.setItem(
        "activeValue",
        this.$route.name +
          JSON.stringify(this.$route.query) +
          JSON.stringify(this.$route.params)
      );
    },
    changeTab(component) {
      const tab = component.$attrs.tab;
      this.$router.push({
        name: tab.name,
        query: tab.query,
        params: tab.params
      });
    },
    removeTab(tab) {
      const index = this.historys.findIndex(
        item =>
          item.name +
            JSON.stringify(item.query) +
            JSON.stringify(item.params) ==
          tab
      );
      if (
        this.$route.name +
          JSON.stringify(this.$route.query) +
          JSON.stringify(this.$route.params) ==
        tab
      ) {
        if (this.historys.length == 1) {
          this.$router.push({ name: this.defaultRouter });
        } else {
          if (index < this.historys.length - 1) {
            this.$router.push({
              name: this.historys[index + 1].name,
              query: this.historys[index + 1].query,
              params: this.historys[index + 1].params
            });
          } else {
            this.$router.push({
              name: this.historys[index - 1].name,
              query: this.historys[index - 1].query,
              params: this.historys[index - 1].params
            });
          }
        }
      }
      this.historys.splice(index, 1);
    }
  },
  watch: {
    contextMenuVisible() {
      if (this.contextMenuVisible) {
        document.body.addEventListener("click", () => {
          this.contextMenuVisible = false;
        });
      } else {
        document.body.removeEventListener("click", () => {
          this.contextMenuVisible = false;
        });
      }
    },
    $route(to, now) {
      this.historys = this.historys.filter(item => !item.meta.closeTab);
      this.setTab(to);
      sessionStorage.setItem("historys", JSON.stringify(this.historys));
      this.activeValue = window.sessionStorage.getItem("activeValue");
      if (now && to && now.name == to.name) {
        this.$bus.$emit("reload");
      }
    }
  }
};
</script>
<style lang="scss">
.contextmenu {
  width: 100px;
  margin: 0;
  border: 1px solid #ccc;
  background: #fff;
  z-index: 3000;
  position: absolute;
  list-style-type: none;
  padding: 5px 0;
  border-radius: 4px;
  font-size: 14px;
  color: #333;
  box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.2);
}
.contextmenu li {
  margin: 0;
  padding: 7px 16px;
}
.contextmenu li:hover {
  background: #f2f2f2;
  cursor: pointer;
}
</style>