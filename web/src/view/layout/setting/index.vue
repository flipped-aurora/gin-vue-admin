<template>
  <div>
    <el-button type="primary" class="drawer-container" icon="el-icon-setting" @click="showSettingDrawar" />
    <el-drawer
      title="系统配置"
      :visible.sync="drawer"
      :direction="direction"
      :before-close="handleClose"
    >
      <div class="setting_body">
        <div class="setting_card">
          <div class="setting_title">侧边栏主题</div>
          <div class="setting_content">
            <div class="item" @click="chageMode('light')">
              <i v-if="sideMode === 'light'" class="el-icon-check check" />
              <img src="https://gw.alipayobjects.com/zos/antfincdn/NQ%24zoisaD2/jpRkZQMyYRryryPNtyIC.svg">
            </div>
            <div class="item" @click="chageMode('dark')">
              <i v-if="sideMode === 'dark'" class="el-icon-check check" />
              <img src="https://gw.alipayobjects.com/zos/antfincdn/XwFOFbLkSM/LCkqqYNmvBEbokSDscrm.svg">
            </div>
          </div>
        </div>
        <div class="setting_card">
          <div class="setting_title">主题色</div>
          <div class="">
            <theme-change style="width: 30px;height: 30px;margin-top: 20px" @change="themeChange" />
          </div>
        </div>
      </div>
    </el-drawer>

  </div>
</template>

<script>
import themeChange from '@/components/themeChange'
export default {
  data() {
    return {
      drawer: false,
      direction: 'rtl',
      sideMode: 'dark'
    }
  },
  components: {
    themeChange
  },
  methods: {
    handleClose() {
      this.drawer = false
    },
    showSettingDrawar() {
      this.drawer = true
    },
    chageMode(e) {
      this.sideMode = e
      this.$store.dispatch('app/changeSideMode')
    },
    themeChange(val) {
      this.$store.dispatch('app/changeTheme', val)
    }
  }
}
</script>

<style lang="scss" scoped>
.drawer-container {
  position: absolute;
  right: 0;
  top: 20%;
  height: 40px;
  width: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
  color: #fff;
  border-radius: 4px 0 0 4px;
  cursor: pointer;
  -webkit-box-shadow: inset 0 0 6px rgb(0 ,0 ,0, 10%);
}
.setting_body{
  padding: 20px;
  .setting_card{
    margin-bottom: 20px;
  }
  .setting_content{
    margin-top: 20px;
    display: flex;
    .item{
      position: relative;
      display: flex;
      align-items: center;
      justify-content: center;
      .check{
        position: absolute;
        font-size: 20px;
        color: #00afff;
      }
      img{
        margin-right: 20px;
      }
    }
  }
}

</style>
