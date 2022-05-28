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

.gva-menu-item-title {
    flex:1;
  }
}

.el-menu--collapse{
  .el-menu-item.is-active{
    color: v-bind(activeBackground);
  }
}

.el-menu-item{
  color: v-bind(normalText);
}

.el-menu-item.is-active{
  .gva-menu-item{
    background: v-bind(activeBackground) !important;
    border-radius: 4px;
    box-shadow: 0 0 2px 1px v-bind(activeBackground) !important;
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
    border-radius: 4px;
    box-shadow: 0 0 2px 1px v-bind(hoverBackground);
    color: v-bind(hoverText);
  }
}

</style>
