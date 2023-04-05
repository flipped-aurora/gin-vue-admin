<template>
  <el-dialog
      v-model="dialogVisible"
      width="30%"
      class="overlay"
  >
      <el-select
          ref="searchInput"
          filterable
          placeholder="请选择要去的页面"
          @change="changeRouter"
      >
        <el-option-group
            v-for="group in menuList"
            :key="group.label"
            :label="group.label"
        >
          <el-option
              v-for="item in group.options"
              :key="item.value"
              :label="item.meta.title"
              :value="item.path"
          />
        </el-option-group>
      </el-select>
    <div class="apps">
      <p class="title">快速应用</p>
      <ul>
        <li class="list">
          <el-row>
            <el-col :span="10" style="height: 40px;">
              <el-icon class="icon"><Switch /></el-icon>
              切换主题
            </el-col>
            <el-col :span="14">
              <el-row style="height: 40px;">
                <el-col :span="12" style="height: 40px;">
                  <div class="item" @click="changeMode('light')">
                    <div class="item-top">
                      <el-icon v-if="userStore.mode === 'light'" class="check">
                        <check />
                      </el-icon>
                      <img src="https://gw.alipayobjects.com/zos/antfincdn/NQ%24zoisaD2/jpRkZQMyYRryryPNtyIC.svg">
                    </div>
                    <p>
                      简约白
                    </p>
                  </div>
                </el-col>
                <el-col :span="12" style="height: 40px;">
                  <div class="item" @click="changeMode('dark')">
                    <div class="item-top">
                      <el-icon v-if="userStore.mode === 'dark'" class="check">
                        <check />
                      </el-icon>
                      <img src="https://gw.alipayobjects.com/zos/antfincdn/XwFOFbLkSM/LCkqqYNmvBEbokSDscrm.svg">
                    </div>
                    <p>
                      商务黑
                    </p>
                  </div>
                </el-col>
              </el-row>
            </el-col>
          </el-row>
        </li>
        <li class="list" @click="userStore.LoginOut">
          <el-icon><ReadingLamp /></el-icon>
          登出
        </li>
      </ul>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="close">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
export default {
  name: 'commandMenu',
}
</script>
<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useRouterStore } from '@/pinia/modules/router'
import { useUserStore } from '@/pinia/modules/user'
const router = useRouter()
const userStore = useUserStore()
const routerStore = useRouterStore()
const dialogVisible = ref(false)
const searchInput = ref(null)
const menus = ref(routerStore.asyncRouters[0].children)
const menuList = []

const open = () => {
  dialogVisible.value = true
  menus.value.forEach((item) => {
    if(item.children && !item.hidden ) {
      let obj = {}
      obj.label = item.meta.title
      obj.basePath = item.path
      obj.options = item.children
      menuList.push(obj)
    }
  })
}
const changeRouter = (e) => {
  if (e.indexOf('http:') > -1 || e.indexOf('https:') > -1) {
    window.open(e)
    return
  }
  dialogVisible.value = false
  router.push({ name: e })
}

const changeMode = (e) => {
  if (e === null) {
    userStore.changeSideMode('dark')
    return
  }
  userStore.changeSideMode(e)
}

const close = () =>{
  dialogVisible.value = false
}

defineExpose({ open })
</script>

<style scoped lang="scss">
.overlay {
  color: #666;
.apps{
  padding-top: 30px;
  list-style: none;
  .title{
    font-size: 12px;
  }
  .list{
    height: 48px;
    line-height: 48px;
    margin-bottom: 20px;
    cursor:pointer;
  }
  .item{
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-right: 20px;
    .item-top{
      position: relative;
      img {
        width: 36px;
        height: 36px;
      }
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
