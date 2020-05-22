<template>
  <div>
    <el-scrollbar style="height:calc(100vh - 64px)">
      <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
        <el-menu
          :collapse="isCollapse"
          :collapse-transition="true"
          :default-active="active"
          @select="selectMenuItem"
          active-text-color="#fff"
          class="el-menu-vertical"
          text-color="rgb(191, 203, 217)"
          unique-opened
        >
          <template v-for="item in asyncRouters[0].children">
            <aside-component :key="item.name" :routerInfo="item" v-if="!item.hidden" />
          </template>
        </el-menu>
      </transition>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import AsideComponent from '@/view/layout/aside/asideComponent'
export default {
  name: 'Aside',
  data() {
    return {
      active: '',
      isCollapse: false
    }
  },
  methods: {
    ...mapMutations('history', ['addHistory']),
    selectMenuItem(index) {
      if (index === this.$route.name) return
      this.$router.push({ name: index })
    }
  },
  computed: {
    ...mapGetters('router', ['asyncRouters'])
  },
  components: {
    AsideComponent
  },
  created() {
    this.active = this.$route.name
    let screenWidth = document.body.clientWidth
    if (screenWidth < 1000) {
      this.isCollapse = !this.isCollapse
    }

    this.$bus.on('collapse', item => {
      this.isCollapse = item
    })
  },
  watch: {
    $route() {
      this.active = this.$route.name
    }
  },
  beforeDestroy() {
    this.$bus.off('collapse')
  }
}
</script>

<style lang="scss">
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