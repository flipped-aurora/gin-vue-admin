<template>
  <el-menu-item :index="routerInfo.name">
    <template v-if="isCollapse">
      <el-tooltip
        class="box-item"
        effect="light"
        :content="routerInfo.meta.title"
        placement="right"
      >
        <el-icon v-if="routerInfo.meta.icon">
          <component :is="routerInfo.meta.icon" />
        </el-icon>
      </el-tooltip>
    </template>
    <template v-else>
      <div class="gva-menu-item">
        <el-icon v-if="routerInfo.meta.icon">
          <component :is="routerInfo.meta.icon" />
        </el-icon>
        <span class="gva-menu-item-title">{{ routerInfo.meta.title }}</span>
      </div>
    </template>
  </el-menu-item>
</template>

<script>
export default {
  name: 'MenuItem',
}
</script>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  routerInfo: {
    default: function() {
      return null
    },
    type: Object
  },
  isCollapse: {
    default: function() {
      return false
    },
    type: Boolean
  },
  theme: {
    default: function() {
      return {}
    },
    type: Object
  }
})

const activeBackground = ref(props.theme.activeBackground)
const activeText = ref(props.theme.activeText)
const normalText = ref(props.theme.normalText)
const hoverBackground = ref(props.theme.hoverBackground)
const hoverText = ref(props.theme.hoverText)

watch(() => props.theme, () => {
  activeBackground.value = props.theme.activeBackground
  activeText.value = props.theme.activeText
  normalText.value = props.theme.normalText
  hoverBackground.value = props.theme.hoverBackground
  hoverText.value = props.theme.hoverText
})

</script>

<style lang="scss" scoped>

 .gva-menu-item{
   width: 100%;
    flex:1;
    height: 44px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    padding-left: 4px;

.gva-menu-item-title {
    flex:1;
  }
}

.el-menu--collapse{
  width: auto;
  margin: 0;
  .el-menu-item.is-active{
    color: #fff;
    background-color:v-bind(activeBackground) ;
    border-radius: 10px;
    width:40px;
    height: 40px;
    left: 50%;
    transform: translateX(-50%);
    box-shadow: -0.9765625px 4.015625px 15px -2px v-bind(activeBackground) !important;
    z-index: 9998;
  }
}

.el-menu-item{
  color: v-bind(normalText);
}

.el-menu-item.is-active{
  .gva-menu-item{
    background: v-bind(activeBackground) !important;
    height: 38px;
    border-radius: 8px;
    box-shadow: -3.9765625px 12.00390625px 25px -11px v-bind(activeBackground) !important;
    z-index: 9999;
    i{
      color: v-bind(activeText);
    }
    span{
      color: v-bind(activeText);
    }
  }
}


.el-menu-item:hover{
  .gva-menu-item{
    background: v-bind(hoverBackground);
    border-radius: 8px;
    box-shadow: 0 0 2px 1px v-bind(hoverBackground);
    color: v-bind(hoverText);
  }
}

</style>
