<template>
  <div class="gva-form-box">
    <div class="p-4 bg-white dark:bg-slate-900">
      <WarningBar
        title="目前只支持标准插件（通过插件模板生成的标准目录插件），非标准插件请自行打包"
      />
      <div class="flex items-center gap-3">
        <el-input v-model="plugName" placeholder="插件模板处填写的【插件名】" />
      </div>
      <el-card class="mt-2 text-center">
        <WarningBar title="穿梭框请只选择子级菜单即可" />
        <el-input
          v-model="parentMenu"
          placeholder="请输入菜单组名，例：公告管理"
          class="mb-2"
        ></el-input>
        <el-transfer
          v-model="menus"
          :props="{
            key: 'ID'
          }"
          class="plugin-transfer"
          :data="menusData"
          filterable
          :filter-method="filterMenuMethod"
          filter-placeholder="请输入菜单名称/路径"
          :titles="['可选菜单', '使用菜单']"
          :button-texts="['移除', '选中']"
        >
          <template #default="{ option }">
            {{ option.meta.title }} {{ option.component }}
          </template>
        </el-transfer>
        <div class="flex justify-end mt-2">
          <el-button type="primary" @click="fmtInitMenu">
            定义安装菜单
          </el-button>
        </div>
      </el-card>
      <el-card class="mt-2 text-center">
        <el-transfer
          v-model="apis"
          :props="{
            key: 'ID'
          }"
          class="plugin-transfer"
          :data="apisData"
          filterable
          :filter-method="filterApiMethod"
          filter-placeholder="请输入API描述/PATH"
          :titles="['可选API', '使用API']"
          :button-texts="['移除', '选中']"
        >
          <template #default="{ option }">
            {{ option.description }} {{ option.path }}
          </template>
        </el-transfer>
        <div class="flex justify-end mt-2">
          <el-button type="primary" @click="fmtInitAPI">
            定义安装API
          </el-button>
        </div>
      </el-card>
      <el-card class="mt-2 text-center">
        <el-transfer
          v-model="dictionaries"
          :props="{
            key: 'ID'
          }"
          class="plugin-transfer"
          :data="dictionariesData"
          filterable
          :filter-method="filterDictionaryMethod"
          filter-placeholder="请输入字典名称/Type"
          :titles="['可选字典', '使用字典']"
          :button-texts="['移除', '选中']"
        >
          <template #default="{ option }">
            {{ option.name }} {{ option.type }}
          </template>
        </el-transfer>
        <div class="flex justify-end mt-2">
          <el-button type="primary" @click="fmtInitDictionary">
            定义安装字典
          </el-button>
        </div>
      </el-card>
    </div>
    <div class="flex justify-end">
      <el-button type="primary" @click="pubPlugin"> 打包插件 </el-button>
    </div>
  </div>
</template>

<script setup>
  import { ref } from 'vue'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { pubPlug, initMenu, initAPI, initDictionary } from '@/api/autoCode.js'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { getAllApis } from '@/api/api'
  import { getMenuList } from '@/api/menu'
  import { getSysDictionaryList } from '@/api/sysDictionary'

  const plugName = ref('')

  const menus = ref([])
  const menusData = ref([])
  const apis = ref([])
  const apisData = ref([])
  const dictionaries = ref([])
  const dictionariesData = ref([])
  const parentMenu = ref('')

  const fmtMenu = (menus) => {
    // 如果menu存在children，递归展开到一级
    const res = []
    menus.forEach((item) => {
      if (item.children) {
        res.push(...fmtMenu(item.children))
      } else {
        res.push(item)
      }
    })
    return res
  }

  const initData = async () => {
    const menuRes = await getMenuList()
    if (menuRes.code === 0) {
      menusData.value = fmtMenu(menuRes.data)
    }
    const apiRes = await getAllApis()
    if (apiRes.code === 0) {
      apisData.value = apiRes.data.apis
    }
    const dictionaryRes = await getSysDictionaryList({
      page: 1,
      pageSize: 9999
    })
    if (dictionaryRes.code === 0) {
      dictionariesData.value = dictionaryRes.data
    }
  }

  const filterMenuMethod = (query, item) => {
    return (
      item.meta.title.indexOf(query) > -1 || item.component.indexOf(query) > -1
    )
  }

  const filterApiMethod = (query, item) => {
    return item.description.indexOf(query) > -1 || item.path.indexOf(query) > -1
  }

  const filterDictionaryMethod = (query, item) => {
    return item.name.indexOf(query) > -1 || item.type.indexOf(query) > -1
  }

  initData()


  const pubPlugin = async () => {
    ElMessageBox.confirm(
      `请检查server下的/plugin/${plugName.value}/plugin.go是否已放开需要的 initialize.Api(ctx), initialize.Menu(ctx) 和 initialize.Dictionary(ctx)?`,
      '打包',
      {
        confirmButtonText: '打包',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
      .then(async () => {
        const res = await pubPlug({ plugName: plugName.value })
        if (res.code === 0) {
          ElMessage.success(res.msg)
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '关闭打包'
        })
      })
  }

  const fmtInitMenu = () => {
    if (!parentMenu.value) {
      ElMessage.error('请填写菜单组名')
      return
    }
    if (menus.value.length === 0) {
      ElMessage.error('请至少选择一个菜单')
      return
    }
    if (plugName.value === '') {
      ElMessage.error('请填写插件名')
      return
    }
    ElMessageBox.confirm(
      `点击后将会覆盖server下的/plugin/${plugName.value}/initialize/menu. 是否继续?`,
      '生成初始菜单',
      {
        confirmButtonText: '生成',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
      .then(async () => {
        const req = {
          plugName: plugName.value,
          parentMenu: parentMenu.value,
          menus: menus.value
        }
        const res = await initMenu(req)
        if (res.code === 0) {
          ElMessage.success('菜单注入成功')
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '关闭生成菜单'
        })
      })
  }
  const fmtInitAPI = () => {
    if (apis.value.length === 0) {
      ElMessage.error('请至少选择一个API')
      return
    }
    if (plugName.value === '') {
      ElMessage.error('请填写插件名')
      return
    }
    ElMessageBox.confirm(
      `点击后将会覆盖server下的/plugin/${plugName.value}/initialize/api. 是否继续?`,
      '生成初始API',
      {
        confirmButtonText: '生成',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
      .then(async () => {
        const req = {
          plugName: plugName.value,
          apis: apis.value
        }
        const res = await initAPI(req)
        if (res.code === 0) {
          ElMessage.success('API注入成功')
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '关闭生成API'
        })
      })
  }

  const fmtInitDictionary = () => {
    if (dictionaries.value.length === 0) {
      ElMessage.error('请至少选择一个字典')
      return
    }
    if (plugName.value === '') {
      ElMessage.error('请填写插件名')
      return
    }
    ElMessageBox.confirm(
      `点击后将会覆盖server下的/plugin/${plugName.value}/initialize/dictionary. 是否继续?`,
      '生成初始字典',
      {
        confirmButtonText: '生成',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
      .then(async () => {
        const req = {
          plugName: plugName.value,
          dictionaries: dictionaries.value
        }
        const res = await initDictionary(req)
        if (res.code === 0) {
          ElMessage.success('字典注入成功')
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '关闭生成字典'
        })
      })
  }
</script>

<style lang="scss">
  .plugin-transfer {
    .el-transfer-panel {
      width: 400px !important;
    }
  }
</style>
