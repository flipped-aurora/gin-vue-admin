<template>
  <div class="gva-form-box">
    <div class="p-4 bg-white dark:bg-slate-900">
      <WarningBar :title="t('view.systemTools.pubPlug.pluginSupport')" />
      <div class="flex items-center gap-3">
        <el-input
          v-model="plugName"
          :placeholder="t('view.systemTools.pubPlug.pluginNameInput')"
        />
      </div>
      <el-card class="mt-2 text-center">
        <WarningBar :title="t('view.systemTools.pubPlug.menuSelectionNote')" />
        <el-input v-model="parentMenu" :placeholder="t('view.systemTools.pubPlug.menuGroupNameInput')" class="mb-2"></el-input>
        <el-transfer
          v-model="menus"
          :props="{
            key: 'ID',
          }"
          class="plugin-transfer"
          :data="menusData"
          filterable
          :filter-method="filterMenuMethod"
          :filter-placeholder="t('view.systemTools.pubPlug.menuNamePathInput')"
          :titles="[t('view.systemTools.pubPlug.optionalMenu'),t('view.systemTools.pubPlug.useMenu')]"
          :button-texts="[t('view.systemTools.pubPlug.remove'), t('view.systemTools.pubPlug.selected')]"
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
            {{ t('view.systemTools.pubPlug.defineMenuInstall') }}
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
          :filter-placeholder="t('view.systemTools.pubPlug.apiDescriptionPathInput')"
          :titles="[t('view.systemTools.pubPlug.optionalAPI'),t('view.systemTools.pubPlug.useAPI')]"
          :button-texts="[t('view.systemTools.pubPlug.remove'), t('view.systemTools.pubPlug.selected')]"
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
            {{ t('view.systemTools.pubPlug.defineAPIInstall') }}
          </el-button>
        </div>
      </el-card>
    </div>
    <div class="flex justify-end">
      <el-button
        type="primary"
        @click="pubPlugin"
      >
        {{ t('view.systemTools.pubPlug.packagePlugin') }}
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
      t('view.systemTools.pubPlug.checkMsg'),
      t('view.systemTools.pubPlug.package'),
      {
        confirmButtonText: t('view.systemTools.pubPlug.package'),
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
          message: t('view.systemTools.pubPlug.closePackage'),
        })
      })

}

const fmtInitMenu = () => {
  if (!parentMenu.value) {
    ElMessage.error(t('view.systemTools.pubPlug.enterMenuGroupName'))
    return
  }
  if (menus.value.length === 0) {
    ElMessage.error(t('view.systemTools.pubPlug.selectAtLeastOneMenu'))
    return
  }
  if (plugName.value === '') {
    ElMessage.error(t('view.systemTools.pubPlug.enterPluginName'))
    return
  }
  ElMessageBox.confirm(
      t('view.systemTools.pubPlug.overwriteMessage'),
      t('view.systemTools.pubPlug.generateInitialMenu'),
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
          message: t('view.systemTools.pubPlug.closeGenerateMenu'),
        })
      })
}
const fmtInitAPI = () => {
  if (apis.value.length === 0) {
    ElMessage.error(t('view.systemTools.pubPlug.selectAtLeastOneAPI'))
    return
  }
  if (plugName.value === '') {
    ElMessage.error(t('view.systemTools.pubPlug.enterPluginNameAgain'))
    return
  }
  ElMessageBox.confirm(
      t('view.systemTools.pubPlug.overwriteWarning'),
      t('view.systemTools.pubPlug.generateInitialAPI'),
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
          message: t('view.systemTools.pubPlug.closeGenerateAPI'),
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
