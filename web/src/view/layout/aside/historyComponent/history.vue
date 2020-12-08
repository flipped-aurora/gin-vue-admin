<template>
  <div class="router-history">
    <el-tabs
      :closable="!(historys.length==1&&this.$route.name=='dashboard')"
      @contextmenu.prevent.native="openContextMenu($event)"
      @tab-click="changeTab"
      @tab-remove="removeTab"
      type="card"
      v-model="activeValue"
    >
      <el-tab-pane
        :key="item.name"
        :label="item.meta.title"
        :name="item.name"
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
export default {
  name: 'HistoryComponent',
  data() {
    return {
      historys: [],
      activeValue: 'dashboard',
      contextMenuVisible: false,
      left: 0,
      top: 0,
      isCollapse: false,
      isMobile: false,
      rightActive: ''
    }
  },
  created() {
    this.$bus.on('mobile', isMobile => {
      this.isMobile = isMobile
    })
    this.$bus.on('collapse', isCollapse => {
      this.isCollapse = isCollapse
    })
    const initHistorys = [
      {
        name: 'dashboard',
        meta: {
          title: '仪表盘'
        }
      }
    ]
    this.historys =
      JSON.parse(sessionStorage.getItem('historys')) || initHistorys
    this.setTab(this.$route)
  },

  beforeDestroy() {
    this.$bus.off('collapse')
    this.$bus.off('mobile')
  },
  methods: {
    openContextMenu(e) {
      if (this.historys.length == 1 && this.$route.name == 'dashboard') {
        return false
      }
      if (e.srcElement.id) {
        this.contextMenuVisible = true
        let width
        if (this.isCollapse) {
          width = 54
        } else {
          width = 220
        }
        if (this.isMobile) {
          width = 0
        }
        this.left = e.clientX - width
        this.top = e.clientY + 10
        this.rightActive = e.srcElement.id.split('-')[1]
      }
    },
    closeAll() {
      this.historys = [
        {
          name: 'dashboard',
          meta: {
            title: '仪表盘'
          }
        }
      ]
      this.$router.push({ name: 'dashboard' })
      this.contextMenuVisible = false
      sessionStorage.setItem('historys', JSON.stringify(this.historys))
    },
    closeLeft() {
      let right
      const rightIndex = this.historys.findIndex(
        item => {
          if(item.name == this.rightActive){
            right = item
          }
          return item.name == this.rightActive
        }
      )
      const activeIndex = this.historys.findIndex(
        item => item.name == this.activeValue
      )
      this.historys.splice(0, rightIndex)
      if (rightIndex > activeIndex) {
        this.$router.push(right)
      }
      sessionStorage.setItem('historys', JSON.stringify(this.historys))
    },
    closeRight() {
      let right
      const leftIndex = this.historys.findIndex(
        item => {
          if(item.name == this.rightActive){
            right = item
          }
          return item.name == this.rightActive
        }
      )
      const activeIndex = this.historys.findIndex(
        item => item.name == this.activeValue
      )
      this.historys.splice(leftIndex + 1, this.historys.length)
      if (leftIndex < activeIndex) {
        this.$router.push(right)
      }
      sessionStorage.setItem('historys', JSON.stringify(this.historys))
    },
    closeOther() {
      let right
      this.historys = this.historys.filter(
        item => {
          if(item.name == this.rightActive){
            right = item
          }
          return item.name == this.rightActive
        }
      )
      this.$router.push(right)
      sessionStorage.setItem('historys', JSON.stringify(this.historys))
    },
    setTab(route) {
      if (!this.historys.some(item => item.name == route.name)) {
        const obj = {}
        obj.name = route.name
        obj.meta = route.meta
        obj.query = route.query
        obj.params = route.params
        this.historys.push(obj)
      }
      this.activeValue = this.$route.name
    },
    changeTab(component) {
      const tab = component.$attrs.tab
      this.$router.push({ name: tab.name,query:tab.query,params:tab.params })
    },
    removeTab(tab) {
      const index = this.historys.findIndex(item => item.name == tab)
      if (this.$route.name == tab) {
        if (this.historys.length == 1) {
          this.$router.push({ name: 'dashboard' })
        } else {
          if (index < this.historys.length - 1) {
            this.$router.push({ name: this.historys[index + 1].name,query:this.historys[index + 1].query,params:this.historys[index + 1].params })
          } else {
            this.$router.push({ name: this.historys[index - 1].name,query:this.historys[index - 1].query,params:this.historys[index - 1].params })
          }
        }
      }
      this.historys.splice(index, 1)
    }
  },
  watch: {
    contextMenuVisible() {
      if (this.contextMenuVisible) {
        document.body.addEventListener('click', () => {
          this.contextMenuVisible = false
        })
      } else {
        document.body.removeEventListener('click', () => {
          this.contextMenuVisible = false
        })
      }
    },
    $route(to) {
      this.historys = this.historys.filter(item => !item.meta.hidden)
      this.setTab(to)
      sessionStorage.setItem('historys', JSON.stringify(this.historys))
    }
  }
}
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