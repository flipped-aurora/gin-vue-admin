<template>
  <div>
    <div @click="isCollapse=!isCollapse" class="menu-total">
      <i class="el-icon-arrow-right" v-if="isCollapse"></i>
      <i class="el-icon-arrow-left" v-else></i>
    </div>
    <el-scrollbar style="height:calc(100vh - 60px)">
       
        <el-menu
          :collapse-transition="true"
          :class="['el-menu-vertical',!isCollapse&&'noCollapse']"
          :collapse="isCollapse"
          :default-active="active"
          @select="selectMenuItem"
          unique-opened
          background-color="#0F3D5F"
          text-color="#bbb"
          active-text-color="#fff"
        >
          <template v-for="item in asyncRouters[0].children">
            <aside-component :key="item.name" :routerInfo="item" v-if="!item.hidden" />
          </template>
        </el-menu>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
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
    this.$bus.on('totalCollapse', () => {
      this.isCollapse = !this.isCollapse
    })
  },
  beforeDestroy() {
    this.$bus.off('totalCollapse')
  }
}
</script>

<style lang="scss">
.el-scrollbar{
  .el-scrollbar__view{
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