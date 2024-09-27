<template>
  <div class="gva-form-box">
    <div class="p-4 bg-white dark:bg-slate-900">
      <WarningBar :title="t('view.plugins.pluginSupport')" />
      <div class="flex items-center gap-3">
        <el-input
          v-model="plugName"
          :placeholder="t('view.plugins.pluginNameInput')"
        />
      </div>
      <el-card class="mt-2 text-center">
        <WarningBar :title="t('view.plugins.menuSelectionNote')" />
        <el-input v-model="parentMenu" :placeholder="t('view.plugins.menuGroupNameInput')" class="mb-2"></el-input>
        <el-transfer
          v-model="menus"
          :props="{
            key: 'ID',
          }"
          class="plugin-transfer"
          :data="menusData"
          filterable
          :filter-method="filterMenuMethod"
          :filter-placeholder="t('view.plugins.menuNamePathInput')"
          :titles="[t('view.plugins.optionalMenu'),t('view.plugins.useMenu')]"
          :button-texts="[t('view.plugins.remove'), t('view.plugins.selected')]"
        >
          <template #default="{option}">
            {{ option.meta.title }} {{ option.component }}
          </template>
        </el-transfer>
        <div class="flex justify-end mt-2">
          <el-button
            type="primary"
            @click="fmtInitMenu"
          >
            {{ t('view.plugins.defineMenuInstall') }}
          </el-button>
        </div>
      </el-card>
      <el-card class="mt-2 text-center">
        <el-transfer
          v-model="apis"
          :props="{
            key: 'ID',
          }"
          class="plugin-transfer"
          :data="apisData"
          filterable
          :filter-method="filterApiMethod"
          :filter-placeholder="t('view.plugins.apiDescriptionPathInput')"
          :titles="[t('view.plugins.optionalAPI'),t('view.plugins.useAPI')]"
          :button-texts="[t('view.plugins.remove'), t('view.plugins.selected')]"
        >
          <template #default="{option}">
            {{ option.description }} {{ option.path }}
          </template>
        </el-transfer>
        <div class="flex justify-end mt-2">
          <el-button
            type="primary"
            @click="fmtInitAPI"
          >
            {{ t('view.plugins.defineAPIInstall') }}
          </el-button>
        </div>
      </el-card>
    </div>
    <div class="flex justify-end">
      <el-button
        type="primary"
        @click="pubPlugin"
      >
        {{ t('view.plugins.packagePlugin') }}
      </el-button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { pubPlug,initMenu,initAPI } from '@/api/autoCode.js'
import {ElMessage, ElMessageBox} from 'element-plus'
import {getAllApis} from "@/api/api";
import {getMenuList} from "@/api/menu";

const plugName = ref('')

const menus = ref([])
const menusData = ref([])
const apis = ref([])
const apisData = ref([])
const parentMenu = ref('')
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const fmtMenu = (menus) => {
  // 如果menu存在children，递归展开到一级
  const res = []
  menus.forEach(item => {
    if (item.children) {
      res.push(...fmtMenu(item.children))
    } else {
      res.push(item)
    }
  })
  return res
}

const initData = async() => {
  const menuRes = await getMenuList()
  if (menuRes.code === 0) {
    menusData.value = fmtMenu(menuRes.data)
  }
  const apiRes = await getAllApis()
  if (apiRes.code === 0) {
    apisData.value = apiRes.data.apis
  }
}

const filterMenuMethod = (query, item) => {
  return item.meta.title.indexOf(query) > -1 || item.component.indexOf(query) > -1
}

const filterApiMethod = (query, item) => {
  return item.description.indexOf(query) > -1 || item.path.indexOf(query) > -1
}

initData()

const pubPlugin = async() => {

  ElMessageBox.confirm(
      `请检查server下的/plugin/${plugName.value}/plugin.go是否已放开需要的 initialize.Api(ctx) 和 initialize.Menu(ctx)?`,
      t('view.plugins.package'),
      {
        confirmButtonText: t('view.plugins.package'),
        cancelButtonText: t('general.cancel'),
        type: 'warning',
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
          message: t('view.plugins.closePackage'),
        })
      })

}

const fmtInitMenu = () => {
  if (!parentMenu.value) {
    ElMessage.error(t('view.plugins.enterMenuGroupName'))
    return
  }
  if (menus.value.length === 0) {
    ElMessage.error(t('view.plugins.selectAtLeastOneMenu'))
    return
  }
  if (plugName.value === '') {
    ElMessage.error(t('view.plugins.enterPluginName'))
    return
  }
  ElMessageBox.confirm(
      `点击后将会覆盖server下的/plugin/${plugName.value}/initialize/menu. 是否继续?`,
      t('view.plugins.generateInitialMenu'),
      {
        confirmButtonText: t('general.generate'),
        cancelButtonText: t('general.cancel'),
        type: 'warning',
      }
  )
      .then(() => {
        const req = {
          plugName: plugName.value,
          parentMenu: parentMenu.value,
          menus: menus.value
        }
        initMenu(req)
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: t('view.plugins.closeGenerateMenu'),
        })
      })
}
const fmtInitAPI = () => {
  if (apis.value.length === 0) {
    ElMessage.error(t('view.plugins.selectAtLeastOneAPI'))
    return
  }
  if (plugName.value === '') {
    ElMessage.error(t('view.plugins.enterPluginNameAgain'))
    return
  }
  ElMessageBox.confirm(
      `点击后将会覆盖server下的/plugin/${plugName.value}/initialize/api. 是否继续?`,
      '生成初始API',
      {
        confirmButtonText: t('general.generate'),
        cancelButtonText: t('general.cancel'),
        type: 'warning',
      }
  )
      .then(() => {
        const req = {
          plugName: plugName.value,
          apis: apis.value
        }
        initAPI(req)
        console.log(req)
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: t('view.plugins.closeGenerateAPI'),
        })
      })
}

</script>

<style lang="scss">
.plugin-transfer{
  .el-transfer-panel{
    width: 400px !important;
  }
}
</style>
