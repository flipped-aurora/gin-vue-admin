<template>
  <el-menu
  mode="vertical"
    :collapse="isCollapse"
    :default-active="active"
    @select="selectMenuItem"
    :class="['el-menu-vertical',!isCollapse&&'noCollapse']"
  >
  <AsideComponent :routerInfo="asyncRouters[0]"/>
  </el-menu>
</template>

<script>
import {mapGetters} from 'vuex'
import AsideComponent from '@/view/layout/aside/asideComponent'
export default {
  name: 'Aside',
  data() {
    return {
      active: '',
      isCollapse: true
    }
  },
  methods: {
    selectMenuItem(index) {
      this.$router.push({name:index})
    }
  },
  computed:{
    ...mapGetters('router',['asyncRouters'])
  },
  components: {
    AsideComponent
  },
  created() {
    this.active = this.$route.name
  }
}
</script>
