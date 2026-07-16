<template>
  <div class="gva-form-box">
    <div class="p-4 bg-white dark:bg-slate-900">
      <WarningBar
        title="仅支持「全栈标准插件」（由插件模板生成、同时含 web 与 server 目录）。下面依次为插件定义要随包一起安装的菜单 / API / 字典，点击「生成」会写入该插件 initialize 目录下对应的初始化代码，安装插件时自动注册、无需再手动配置。纯前端 / 纯后端 / 非标准插件请自行打包。"
      />

      <!-- 第一步：选择目标插件 -->
      <div class="flex items-center gap-3 mt-4">
        <span
          class="text-sm text-gray-600 dark:text-gray-300 whitespace-nowrap"
        >
          目标插件
        </span>
        <el-select
          v-model="plugName"
          filterable
          placeholder="请选择要打包的插件（无需手动输入插件名）"
          class="flex-1"
          @change="onPluginChange"
        >
          <el-option
            v-for="item in pluginList"
            :key="item.pluginName"
            :label="item.pluginName"
            :value="item.pluginName"
            :disabled="item.pluginType !== 'full'"
          >
            <div class="flex items-center justify-between">
              <span>{{ item.pluginName }}</span>
              <el-tag
                size="small"
                :type="pluginTagType[item.pluginType]"
                class="ml-3"
              >
                {{ pluginTypeMap[item.pluginType] || item.pluginType }}
              </el-tag>
            </div>
          </el-option>
        </el-select>
        <el-button icon="refresh" @click="loadBaseData"> 刷新 </el-button>
      </div>
      <div
        v-if="!pluginList.length"
        class="text-xs text-gray-400 dark:text-gray-500 mt-2"
      >
        暂无可选插件。请确认已在 server/plugin 与 web/plugin
        下生成对应的插件目录（本工具用于打包全栈标准插件）。
      </div>
      <div
        v-else-if="!hasFullPlugin"
        class="text-xs text-gray-400 dark:text-gray-500 mt-2"
      >
        当前无「全栈插件」可选。菜单 / API / 字典初始化会写入 server 端，打包也要求
        web 与 server 目录同时存在，故仅全栈插件可在此操作。
      </div>

      <!-- 步骤列表：菜单 / API / 字典初始化 -->
      <div class="mt-4 flex flex-col gap-3">
        <div
          v-for="(step, idx) in steps"
          :key="step.key"
          class="flex items-center gap-4 p-4 rounded-md border border-gray-200 dark:border-slate-700 bg-gray-50 dark:bg-slate-800"
        >
          <div
            class="flex-none w-8 h-8 rounded-full bg-blue-500 text-white flex items-center justify-center text-sm"
          >
            {{ idx + 1 }}
          </div>
          <div class="flex-1 min-w-0">
            <div class="text-sm font-medium text-gray-800 dark:text-gray-100">
              {{ step.title }}
            </div>
            <div class="text-xs text-gray-500 dark:text-gray-400 mt-1">
              {{ step.desc }}
            </div>
          </div>
          <el-tag size="small" :type="step.tagType" class="flex-none">
            {{ step.tagText }}
          </el-tag>
          <el-button
            type="primary"
            plain
            class="flex-none"
            :disabled="!plugName"
            @click="openDrawer(step.key)"
          >
            配置
          </el-button>
        </div>
      </div>

      <!-- 打包 -->
      <div class="flex justify-end mt-6">
        <el-button type="primary" :disabled="!plugName" @click="pubPlugin">
          打包插件
        </el-button>
      </div>
    </div>

    <!-- 菜单初始化抽屉 -->
    <el-drawer
      v-model="menuDrawer"
      :size="appStore.drawerSize"
      :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-base">定义安装菜单</span>
          <div>
            <el-button @click="menuDrawer = false"> 取 消 </el-button>
            <el-button type="primary" @click="fmtInitMenu">
              生成菜单初始化
            </el-button>
          </div>
        </div>
      </template>
      <WarningBar
        title="穿梭框请只选择「子级菜单」。下方「菜单组名」会作为父级菜单标题，插件安装后所选子菜单会自动挂到该分组下。点击「生成」会覆盖 server/plugin/<插件>/initialize/menu.go。"
      />
      <el-input
        v-model="parentMenu"
        placeholder="请输入菜单组名，例：公告管理"
        class="mb-3"
      />
      <el-transfer
        v-model="menus"
        :props="{ key: 'ID' }"
        class="plugin-transfer"
        :data="menusData"
        filterable
        :filter-method="filterMenuMethod"
        filter-placeholder="请输入菜单名称/路径"
        :titles="['可选菜单', '使用菜单']"
        :button-texts="['移除', '选中']"
      >
        <template #default="{ option }">
          {{ option.meta?.title }} {{ option.component }}
        </template>
      </el-transfer>
    </el-drawer>

    <!-- API 初始化抽屉 -->
    <el-drawer
      v-model="apiDrawer"
      :size="appStore.drawerSize"
      :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-base">定义安装 API</span>
          <div>
            <el-button @click="apiDrawer = false"> 取 消 </el-button>
            <el-button type="primary" @click="fmtInitAPI">
              生成 API 初始化
            </el-button>
          </div>
        </div>
      </template>
      <WarningBar
        title="选择该插件依赖的后端接口。插件安装后这些 API 会自动注册到系统，仍需在「角色管理」内为对应角色授权后才能调用。点击「生成」会覆盖 server/plugin/<插件>/initialize/api.go。"
      />
      <el-transfer
        v-model="apis"
        :props="{ key: 'ID' }"
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
    </el-drawer>

    <!-- 字典初始化抽屉 -->
    <el-drawer
      v-model="dictDrawer"
      :size="appStore.drawerSize"
      :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-base">定义安装字典</span>
          <div>
            <el-button @click="dictDrawer = false"> 取 消 </el-button>
            <el-button type="primary" @click="fmtInitDictionary">
              生成字典初始化
            </el-button>
          </div>
        </div>
      </template>
      <WarningBar
        title="选择该插件依赖的数据字典，插件安装后会自动写入系统字典（含字典项）。点击「生成」会覆盖 server/plugin/<插件>/initialize/dictionary.go。"
      />
      <el-transfer
        v-model="dictionaries"
        :props="{ key: 'ID' }"
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
    </el-drawer>
  </div>
</template>

<script setup>
  import { ref, computed, watch, nextTick } from 'vue'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import {
    pubPlug,
    initMenu,
    initAPI,
    initDictionary,
    getPluginList
  } from '@/api/autoCode.js'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { getAllApis } from '@/api/api'
  import { getMenuList } from '@/api/menu'
  import { getSysDictionaryList } from '@/api/sysDictionary'
  import { useAppStore } from '@/pinia'

  defineOptions({ name: 'PubPlug' })

  const appStore = useAppStore()

  const plugName = ref('')
  const pluginList = ref([])

  const menus = ref([])
  const menusData = ref([])
  const apis = ref([])
  const apisData = ref([])
  const dictionaries = ref([])
  const dictionariesData = ref([])
  const parentMenu = ref('')

  const menuDrawer = ref(false)
  const apiDrawer = ref(false)
  const dictDrawer = ref(false)

  // 各步骤是否已点击「生成」写入 initialize（预勾选=插件已有数据视为已生成）
  const generated = ref({ menu: false, api: false, dictionary: false })
  // 预勾选期间抑制 watch，避免把程序化赋值误判为「用户改动导致待重新生成」
  let applyingPreset = false
  watch(menus, () => {
    if (!applyingPreset) generated.value.menu = false
  })
  watch(apis, () => {
    if (!applyingPreset) generated.value.api = false
  })
  watch(dictionaries, () => {
    if (!applyingPreset) generated.value.dictionary = false
  })
  // 菜单组名也是 menu.go 的写入内容，改动后同样需重新生成
  watch(parentMenu, () => {
    if (!applyingPreset) generated.value.menu = false
  })

  const pluginTypeMap = {
    server: '后端插件',
    web: '前端插件',
    full: '全栈插件'
  }
  const pluginTagType = {
    server: 'warning',
    web: 'success',
    full: 'primary'
  }

  const hasFullPlugin = computed(() =>
    pluginList.value.some((p) => p.pluginType === 'full')
  )

  const steps = computed(() =>
    [
      {
        key: 'menu',
        title: '菜单初始化',
        desc: '定义随插件一起安装的菜单，安装后自动挂到指定菜单分组下',
        count: menus.value.length,
        done: generated.value.menu
      },
      {
        key: 'api',
        title: 'API 初始化',
        desc: '定义插件依赖的后端接口，安装后自动注册，需在角色管理内授权后使用',
        count: apis.value.length,
        done: generated.value.api
      },
      {
        key: 'dictionary',
        title: '字典初始化',
        desc: '定义插件依赖的数据字典，安装后自动写入系统字典',
        count: dictionaries.value.length,
        done: generated.value.dictionary
      }
    ].map((s) => ({
      ...s,
      tagType: s.done ? 'success' : s.count === 0 ? 'info' : 'warning',
      tagText: s.done
        ? s.count > 0
          ? `已生成 ${s.count} 项`
          : '已生成'
        : s.count === 0
          ? '未配置'
          : `已选 ${s.count} 项 · 待生成`
    }))
  )

  const openDrawer = (key) => {
    if (key === 'menu') menuDrawer.value = true
    else if (key === 'api') apiDrawer.value = true
    else if (key === 'dictionary') dictDrawer.value = true
  }

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

  // 只加载可选底表与插件列表，不改动用户当前选择（刷新按钮复用）
  const loadBaseData = async () => {
    try {
      const [menuRes, apiRes, dictionaryRes, pluginRes] = await Promise.all([
        getMenuList(),
        getAllApis(),
        getSysDictionaryList({ page: 1, pageSize: 9999 }),
        getPluginList()
      ])
      if (menuRes.code === 0) {
        menusData.value = fmtMenu(menuRes.data)
      }
      if (apiRes.code === 0) {
        apisData.value = apiRes.data.apis
      }
      if (dictionaryRes.code === 0) {
        dictionariesData.value = dictionaryRes.data
      }
      if (pluginRes.code === 0) {
        pluginList.value = pluginRes.data || []
      }
    } catch (e) {
      ElMessage.error('加载插件基础数据失败')
    }
  }

  // 选中插件后，按该插件「已定义」的菜单/API/字典自动预勾选
  const onPluginChange = (name) => {
    applyingPreset = true
    menus.value = []
    apis.value = []
    dictionaries.value = []
    parentMenu.value = ''
    generated.value = { menu: false, api: false, dictionary: false }

    const info = pluginList.value.find((p) => p.pluginName === name)
    if (info) {
      // 已定义但未匹配到当前系统数据的项（可能丢失，需提醒）
      const missing = []

      // 菜单：父级分组(routerHolder)只用于回填「菜单组名」，不参与子菜单匹配
      const childMenus = (info.menus || []).filter(
        (m) => m.component !== 'view/routerHolder.vue'
      )
      if (info.menus && info.menus.length) {
        const group = info.menus.find(
          (m) => m.component === 'view/routerHolder.vue'
        )
        if (group) {
          parentMenu.value = group.meta?.title || ''
        }
        const menuNames = new Set(childMenus.map((m) => m.name))
        menus.value = menusData.value
          .filter((m) => menuNames.has(m.name))
          .map((m) => m.ID)
        if (menus.value.length < childMenus.length) missing.push('菜单')
      }

      // API：按 path + method 匹配
      if (info.apis && info.apis.length) {
        const apiKeys = new Set(info.apis.map((a) => `${a.path}|${a.method}`))
        apis.value = apisData.value
          .filter((a) => apiKeys.has(`${a.path}|${a.method}`))
          .map((a) => a.ID)
        if (apis.value.length < info.apis.length) missing.push('API')
      }

      // 字典：按 type 匹配
      if (info.dictionaries && info.dictionaries.length) {
        const dictTypes = new Set(info.dictionaries.map((d) => d.type))
        dictionaries.value = dictionariesData.value
          .filter((d) => dictTypes.has(d.type))
          .map((d) => d.ID)
        if (dictionaries.value.length < info.dictionaries.length)
          missing.push('字典')
      }

      // generated 反映「插件 initialize 目录是否已有数据」（源数据本身），
      // 而非匹配到的当前库实体数量，避免库中实体被删时误判为「未配置」
      generated.value = {
        menu: childMenus.length > 0,
        api: !!(info.apis && info.apis.length),
        dictionary: !!(info.dictionaries && info.dictionaries.length)
      }

      if (missing.length) {
        ElMessage.warning(
          `插件已定义的部分「${missing.join(
            '/'
          )}」未匹配到当前系统数据（可能插件未安装或数据已被改动）。直接点「生成」会以当前勾选覆盖 initialize，可能丢失未匹配项，请先核对。`
        )
      }
    }

    nextTick(() => {
      applyingPreset = false
    })
  }

  const filterMenuMethod = (query, item) => {
    return (
      (item.meta?.title || '').indexOf(query) > -1 ||
      (item.component || '').indexOf(query) > -1
    )
  }

  const filterApiMethod = (query, item) => {
    return (
      (item.description || '').indexOf(query) > -1 ||
      (item.path || '').indexOf(query) > -1
    )
  }

  const filterDictionaryMethod = (query, item) => {
    return (
      (item.name || '').indexOf(query) > -1 ||
      (item.type || '').indexOf(query) > -1
    )
  }

  loadBaseData()

  const pubPlugin = async () => {
    if (!plugName.value) {
      ElMessage.error('请先选择插件')
      return
    }
    // 打包只打磁盘上 initialize 的现状；本页选择只有点「生成」才写入。
    // 对「本次改动未生成」的步骤（含新增未生成、已生成后又改动/清空）做提醒。
    const info = pluginList.value.find((p) => p.pluginName === plugName.value)
    const hadData = {
      menu: !!(
        info &&
        info.menus &&
        info.menus.some((m) => m.component !== 'view/routerHolder.vue')
      ),
      api: !!(info && info.apis && info.apis.length),
      dictionary: !!(info && info.dictionaries && info.dictionaries.length)
    }
    const pending = steps.value
      .filter((s) => !s.done && (s.count > 0 || hadData[s.key]))
      .map((s) => s.title)
    const pendingTip = pending.length
      ? ` ⚠️ 以下步骤有未点击「生成」的改动：${pending.join(
          '、'
        )}。打包以磁盘 initialize 现状为准，未生成的改动不会写入插件包，请确认是否需先生成。`
      : ''
    ElMessageBox.confirm(
      `请检查server下的/plugin/${plugName.value}/plugin.go是否已放开需要的 initialize.Api(ctx), initialize.Menu(ctx) 和 initialize.Dictionary(ctx)?${pendingTip}`,
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
    if (plugName.value === '') {
      ElMessage.error('请先选择插件')
      return
    }
    if (!parentMenu.value) {
      ElMessage.error('请填写菜单组名')
      return
    }
    if (menus.value.length === 0) {
      ElMessage.error('请至少选择一个菜单')
      return
    }
    ElMessageBox.confirm(
      `点击后将会覆盖server下的/plugin/${plugName.value}/initialize/menu.go. 是否继续?`,
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
          generated.value.menu = true
          menuDrawer.value = false
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
    if (plugName.value === '') {
      ElMessage.error('请先选择插件')
      return
    }
    if (apis.value.length === 0) {
      ElMessage.error('请至少选择一个API')
      return
    }
    ElMessageBox.confirm(
      `点击后将会覆盖server下的/plugin/${plugName.value}/initialize/api.go. 是否继续?`,
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
          generated.value.api = true
          apiDrawer.value = false
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
    if (plugName.value === '') {
      ElMessage.error('请先选择插件')
      return
    }
    if (dictionaries.value.length === 0) {
      ElMessage.error('请至少选择一个字典')
      return
    }
    ElMessageBox.confirm(
      `点击后将会覆盖server下的/plugin/${plugName.value}/initialize/dictionary.go. 是否继续?`,
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
          generated.value.dictionary = true
          dictDrawer.value = false
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

<style scoped lang="scss">
  .plugin-transfer :deep(.el-transfer-panel) {
    width: 320px !important;
  }
</style>
