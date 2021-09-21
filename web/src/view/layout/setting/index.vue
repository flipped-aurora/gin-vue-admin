<template>
  <div>
    <el-button type="primary" class="drawer-container" icon="el-icon-setting" @click="showSettingDrawer" />
    <el-drawer
      v-model="drawer"
      title="系统配置"
      :direction="direction"
      :before-close="handleClose"
    >
      <div class="setting_body">
        <div class="setting_card">
          <div class="setting_content">
            <div class="theme-box">
              <div class="item" @click="changeMode('light')">
                <div class="item-top">
                  <i v-if="mode === 'light'" class="el-icon-check check" />
                  <img src="https://gw.alipayobjects.com/zos/antfincdn/NQ%24zoisaD2/jpRkZQMyYRryryPNtyIC.svg">
                </div>
                <p>
                  简约白
                </p>
              </div>
              <div class="item" @click="changeMode('dark')">
                <div class="item-top">
                  <i v-if="mode === 'dark'" class="el-icon-check check" />
                  <img src="https://gw.alipayobjects.com/zos/antfincdn/XwFOFbLkSM/LCkqqYNmvBEbokSDscrm.svg">
                </div>
                <p>
                  商务黑
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </el-drawer>

  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  data() {
    return {
      drawer: false,
      direction: 'rtl'
    }
  },
  computed: {
    ...mapGetters('user', ['mode'])
  },
  methods: {
    handleClose() {
      this.drawer = false
    },
    showSettingDrawer() {
      this.drawer = true
    },
    changeMode(e) {
      if (e === null) {
        this.$store.dispatch('user/changeSideMode', 'dark')
        return
      }
      this.$store.dispatch('user/changeSideMode', e)
    },
  }
}
</script>

<style lang="scss" scoped>
.drawer-container {
  transition: all 0.2s;
  &:hover{
    right: 0
  }
  position: fixed;
  right: -20px;
  bottom: 15%;
  height: 40px;
  width: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
  color: #fff;
  border-radius: 4px 0 0 4px;
  cursor: pointer;
  -webkit-box-shadow: inset 0 0 6px rgba(0 ,0 ,0, 10%);
}
.setting_body{
  padding: 20px;
  .setting_card{
    margin-bottom: 20px;
  }
  .setting_content{
    margin-top: 20px;
    display: flex;
    flex-direction: column;
    >.theme-box{
     display: flex;
    }
    >.color-box{
      div{
        display: flex;
        flex-direction: column;
      }
    }
    .item{
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
      margin-right: 20px;
      .item-top{
        position: relative;
      }
      .check{
        position: absolute;
        font-size: 20px;
        color: #00afff;
        right:10px;
        bottom: 10px;
      }
      p{
        text-align: center;
        font-size: 12px;
      }
    }
  }
}

</style>
