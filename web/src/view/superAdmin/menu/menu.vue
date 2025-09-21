<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addMenu(0)">
          {{ t('view.superAdmin.menu.addRootMenu') }}
        </el-button>
      </div>

      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="100" prop="ID" />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.menu.displayName')"
          min-width="160"
          prop="authorityName"
        >
          <template #default="scope">
            <span>{{ scope.row.meta.title }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('view.superAdmin.menu.icon')"
          min-width="140"
          prop="authorityName"
        >
          <template #default="scope">
            <div v-if="scope.row.meta.icon" class="icon-column">
              <el-icon>
                <component :is="scope.row.meta.icon" />
              </el-icon>
              <span>{{ scope.row.meta.icon }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('view.superAdmin.menu.routeName')"
          show-overflow-tooltip
          min-width="160"
          prop="name"
        />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.menu.routePath')"
          show-overflow-tooltip
          min-width="160"
          prop="path"
        />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.menu.visibility')"
          min-width="100"
          prop="hidden"
        >
          <template #default="scope">
            <span>{{
              scope.row.hidden
                ? t('view.superAdmin.menu.hide')
                : t('view.superAdmin.menu.show')
            }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('view.superAdmin.menu.parent')"
          min-width="90"
          prop="parentId"
        />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.menu.sort')"
          min-width="70"
          prop="sort"
        />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.menu.filePath')"
          min-width="360"
          prop="component"
        />
        <el-table-column align="left" fixed="right" :label="t('general.operations')" :min-width="appStore.operateMinWith" width="280px">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="plus"
              @click="addMenu(scope.row.ID)"
            >
              {{ t('view.superAdmin.menu.addSubMenu') }}
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >
              {{ t('general.edit') }}
            </el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteMenu(scope.row.ID)"
            >
              {{ t('general.delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-drawer
      v-model="dialogFormVisible"
      :size="appStore.drawerSize"
      :before-close="handleClose"
      :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ dialogTitle }}</span>
          <div>
            <el-button @click="closeDialog">
              {{ t('general.close') }}
            </el-button>
            <el-button type="primary" @click="enterDialog">
              {{ t('general.confirm') }}
            </el-button>
          </div>
        </div>
      </template>

      <warning-bar :title="t('view.superAdmin.menu.newMenuNote')" />
      
      <!-- 基础信息区域 -->
      <div class="border-b border-gray-200">
        <h3 class="font-semibold text-gray-700 mb-4">基础信息</h3>
        <el-form
          v-if="dialogFormVisible"
          ref="menuForm"
          :inline="true"
          :model="form"
          :rules="rules"
          label-position="top"
        >
          <el-row class="w-full">
            <el-col :span="24">
              <el-form-item :label="t('view.superAdmin.menu.filePath')" prop="component">
                <components-cascader
                  :component="form.component"
                  @change="fmtComponent"
                />
                <div class="form-tip">
                  <el-icon><InfoFilled /></el-icon>
                  <span>{{ t('view.superAdmin.menu.subMenuNote') }}</span>
                  <el-button
                    size="small"
                    type="text"
                    @click="form.component = 'view/routerHolder.vue'"
                  >
                    {{ t('view.superAdmin.menu.clickMe') }}
                  </el-button>
                </div>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row class="w-full">
            <el-col :span="12">
              <el-form-item :label="t('view.superAdmin.menu.displayName')" prop="meta.title">
                <el-input 
                  v-model="form.meta.title" 
                  autocomplete="off" 
                  :placeholder="t('view.superAdmin.menu.titleNote')"
                />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item :label="t('view.superAdmin.menu.routeName')" prop="path">
                <el-input
                  v-model="form.name"
                  autocomplete="off"
                  :placeholder="t('view.superAdmin.menu.routeNameNote')"
                  @change="changeName"
                />
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </div>
       
      <!-- 路由配置区域 -->
      <div class="border-b border-gray-200">
        <h3 class="font-semibold text-gray-700 mb-4">路由配置</h3>
        <el-form
          :inline="true"
          :model="form"
          :rules="rules"
          label-position="top"
        >
           <el-row class="w-full">
             <el-col :span="12">
               <el-form-item :label="t('view.superAdmin.menu.parentId')">
                 <el-cascader
                   v-model="form.parentId"
                   style="width: 100%"
                   :disabled="!isEdit"
                   :options="menuOption"
                   :props="{
                     checkStrictly: true,
                     label: 'title',
                     value: 'ID',
                     disabled: 'disabled',
                     emitPath: false
                   }"
                   :show-all-levels="false"
                   filterable
                   placeholder="请选择父节点"
                 />
               </el-form-item>
             </el-col>
             <el-col :span="12">
               <el-form-item prop="path">
                 <template #label>
                  <div class="inline-flex items-center h-4">
                     <span>{{ t('view.superAdmin.menu.routePath') }}</span>
                     <el-checkbox
                       class="ml-2"
                       v-model="checkFlag"
                       >{{ t('view.superAdmin.menu.addParameter') }}</el-checkbox
                     >
                    </div>
                 </template>
                 <el-input
                   v-model="form.path"
                   :disabled="!checkFlag"
                   autocomplete="off"
                   :placeholder="t('view.superAdmin.menu.routePathNote')"
                 />
               </el-form-item>
             </el-col>
           </el-row>
        </el-form>
      </div>
       
      <!-- 显示设置区域 -->
      <div class="border-b border-gray-200">
        <h3 class="font-semibold text-gray-700 mb-4">显示设置</h3>
        <el-form
          :inline="true"
          :model="form"
          :rules="rules"
          label-position="top"
        >
           <el-row class="w-full">
              <el-col :span="8">
                <el-form-item :label="t('view.superAdmin.menu.icon')" prop="meta.icon">
                  <icon v-model="form.meta.icon" />
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item :label="t('general.order')" prop="sort">
                  <el-input 
                    v-model.number="form.sort" 
                    autocomplete="off" 
                    placeholder="请输入排序数字"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item :label="t('view.superAdmin.menu.visibility')">
                  <el-select
                    v-model="form.hidden"
                    style="width: 100%"
                    :placeholder="t('view.superAdmin.menu.visibilityNote')"
                  >
                    <el-option :value="false" :label="t('general.no')" />
                    <el-option :value="true" :label="t('general.yes')" />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
        </el-form>
      </div>
        
      <!-- 高级配置区域 -->
      <div class="border-b border-gray-200">
        <h3 class="font-semibold text-gray-700 mb-4">高级配置</h3>
        <el-form
          :inline="true"
          :model="form"
          :rules="rules"
          label-position="top"
        >
            <el-row class="w-full">
              <el-col :span="12">
                <el-form-item prop="meta.activeName">
                  <template #label>
                    <div class="label-with-tooltip">
                      <span>{{ t('view.superAdmin.menu.highlightMenu') }}</span>
                      <el-tooltip
                        :content="t('view.superAdmin.menu.highlightMenu')"
                        placement="top"
                        effect="light"
                      >
                        <el-icon><QuestionFilled /></el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                  <el-input
                    v-model="form.meta.activeName"
                    :placeholder="form.name || '请输入高亮菜单名称'"
                    autocomplete="off"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="KeepAlive" prop="meta.keepAlive">
                  <el-select
                    v-model="form.meta.keepAlive"
                    style="width: 100%"
                    :placeholder="t('view.superAdmin.menu.keepAliveNote')"
                  >
                    <el-option :value="false" :label="t('general.no')" />
                    <el-option :value="true" :label="t('general.yes')" />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
             <el-row class="w-full">
               <el-col :span="8">
                 <el-form-item label="CloseTab" prop="meta.closeTab">
                   <el-select
                     v-model="form.meta.closeTab"
                     style="width: 100%"
                     :placeholder="t('view.superAdmin.menu.closeTabNote')"
                   >
                    <el-option :value="false" :label="t('general.no')" />
                    <el-option :value="true" :label="t('general.yes')" />
                   </el-select>
                 </el-form-item>
               </el-col>
               <el-col :span="8">
                 <el-form-item>
                   <template #label>
                     <div class="label-with-tooltip">
                       <span>{{ t('view.superAdmin.menu.basicPage') }}</span>
                       <el-tooltip
                         :content="t('view.superAdmin.menu.basicPageNote')"
                         placement="top"
                         effect="light"
                       >
                         <el-icon><QuestionFilled /></el-icon>
                       </el-tooltip>
                     </div>
                   </template>
                   <el-select
                     v-model="form.meta.defaultMenu"
                     style="width: 100%"
                     :placeholder="t('view.superAdmin.menu.basicPage')"
                   >
                    <el-option :value="false" :label="t('general.no')" />
                    <el-option :value="true" :label="t('general.yes')" />
                   </el-select>
                 </el-form-item>
               </el-col>
               <el-col :span="8">
                 <el-form-item>
                   <template #label>
                     <div class="label-with-tooltip">
                       <span>路由切换动画</span>
                       <el-tooltip
                         content="如果设置了路由切换动画，在本路由下的动画优先级高于全局动画切换优先级。"
                         placement="top"
                         effect="light"
                       >
                         <el-icon><QuestionFilled /></el-icon>
                       </el-tooltip>
                     </div>
                   </template>
                   <el-select
                     v-model="form.meta.transitionType"
                     style="width: 100%"
                     placeholder="跟随全局"
                     clearable
                   >
                     <el-option value="fade" label="淡入淡出" />
                     <el-option value="slide" label="滑动" />
                     <el-option value="zoom" label="缩放" />
                     <el-option value="none" label="无动画" />
                   </el-select>
                 </el-form-item>
               </el-col>
             </el-row>
        </el-form>
      </div>
          
      <!-- 菜单参数配置区域 -->
      <div class="border-b border-gray-200">
        <div class="flex justify-between items-center mb-4">
          <h3 class="font-semibold text-gray-700">菜单参数配置</h3>
          <el-button type="primary" size="small" @click="addParameter(form)">
            {{ t('view.superAdmin.menu.addMenuParameters') }}
          </el-button>
        </div>
            <el-table 
              :data="form.parameters" 
              style="width: 100%"
              class="parameter-table"
            >
              <el-table-column
                align="center"
                prop="type"
                :label="t('view.superAdmin.menu.parameterType')"
                width="150"
              >
                <template #default="scope">
                  <el-select 
                    v-model="scope.row.type" 
                    :placeholder="t('general.pleaseSelect')"
                    size="small"
                  >
                    <el-option key="query" value="query" label="query" />
                    <el-option key="params" value="params" label="params" />
                  </el-select>
                </template>
              </el-table-column>
              <el-table-column align="center" prop="key" :label="t('view.superAdmin.menu.paremeterKey')" width="150">
                <template #default="scope">
                  <el-input 
                    v-model="scope.row.key" 
                    size="small"
                    placeholder="请输入参数key"
                  />
                </template>
              </el-table-column>
              <el-table-column align="center" prop="value" :label="t('view.superAdmin.menu.parameterValue')">
                <template #default="scope">
                  <el-input 
                    v-model="scope.row.value" 
                    size="small"
                    :placeholder="t('view.superAdmin.params.enterParamValue')"
                  />
                </template>
              </el-table-column>
              <el-table-column align="center" :label="t('general.operations')" width="100">
                <template #default="scope">
                  <el-button
                    type="danger"
                    size="small"
                    @click="deleteParameter(form.parameters, scope.$index)"
                  >
                    <el-icon><Delete /></el-icon>
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
      </div>
           
      <!-- 可控按钮配置区域 -->
      <div class="mb-2 mt-2">
        <div class="flex justify-between items-center mb-4">
          <h3 class="font-semibold text-gray-700">可控按钮配置</h3>
          <div class="flex items-center gap-2">
            <el-button type="primary" size="small" @click="addBtn(form)">
              {{ t('view.superAdmin.menu.addButton') }}
            </el-button>
            <el-tooltip
              content="点击查看按钮权限配置文档"
              placement="top"
              effect="light"
            >
              <el-icon
                class="cursor-pointer text-blue-500 hover:text-blue-700"
                @click="toDoc('https://www.gin-vue-admin.com/guide/web/button-auth.html')"
              >
                <QuestionFilled />
              </el-icon>
            </el-tooltip>
          </div>
        </div>
             <el-table 
               :data="form.menuBtn" 
               style="width: 100%"
               class="button-table"
             >
               <el-table-column
                 align="center"
                 prop="name"
                 label="按钮名称"
                 width="150"
               >
                 <template #default="scope">
                   <el-input 
                     v-model="scope.row.name" 
                     size="small"
                     placeholder="请输入按钮名称"
                   />
                 </template>
               </el-table-column>
               <el-table-column align="center" prop="desc" :label="t('view.superAdmin.menu.comments')">
                 <template #default="scope">
                   <el-input 
                     v-model="scope.row.desc" 
                     size="small"
                     placeholder="请输入按钮备注"
                   />
                 </template>
               </el-table-column>
               <el-table-column align="center" :label="t('general.operations')" width="100">
                 <template #default="scope">
                   <el-button
                     type="danger"
                     size="small"
                     @click="deleteBtn(form.menuBtn, scope.$index)"
                   >
                     <el-icon><Delete /></el-icon>
                   </el-button>
                 </template>
               </el-table-column>
             </el-table>
       </div>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    updateBaseMenu,
    getMenuList,
    addBaseMenu,
    deleteBaseMenu,
    getBaseMenuById
  } from '@/api/menu'
  import icon from '@/view/superAdmin/menu/icon.vue'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { canRemoveAuthorityBtnApi } from '@/api/authorityBtn'
  import { reactive, ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { QuestionFilled, InfoFilled, Delete } from '@element-plus/icons-vue'
  import { toDoc } from '@/utils/doc'
  import { toLowerCase } from '@/utils/stringFun'
  import ComponentsCascader from '@/view/superAdmin/menu/components/components-cascader.vue'

  import pathInfo from '@/pathInfo.json'
  import { useAppStore } from "@/pinia";
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual

  const { t } = useI18n() // added by mohamed hassan to support multilingual

  defineOptions({
    name: 'Menus'
  })

  const appStore = useAppStore()

  const rules = reactive({
    path: [
      {
        required: true,
        message: t('view.superAdmin.menu.enterMenuNameNote'),
        trigger: 'blur'
      }
    ],
    component: [
      {
        required: true,
        message: t('view.superAdmin.menu.enterFilePathNote'),
        trigger: 'blur'
      }
    ],
    'meta.title': [
      {
        required: true,
        message: t('view.superAdmin.menu.enterMenuDisplayNameNote'),
        trigger: 'blur'
      }
    ]
  })

  const tableData = ref([])
  // 查询
  const getTableData = async () => {
    const table = await getMenuList()
    if (table.code === 0) {
      tableData.value = table.data
    }
  }

  getTableData()

  // 新增参数
  const addParameter = (form) => {
    if (!form.parameters) {
      form.parameters = []
    }
    form.parameters.push({
      type: 'query',
      key: '',
      value: ''
    })
  }

  const fmtComponent = (component) => {
    form.value.component = component.replace(/\\/g, '/')
    form.value.name = toLowerCase(pathInfo['/src/' + component])
    form.value.path = form.value.name
  }

  // 删除参数
  const deleteParameter = (parameters, index) => {
    parameters.splice(index, 1)
  }

  // 新增可控按钮
  const addBtn = (form) => {
    if (!form.menuBtn) {
      form.menuBtn = []
    }
    form.menuBtn.push({
      name: '',
      desc: ''
    })
  }
  // 删除可控按钮
  const deleteBtn = async (btns, index) => {
    const btn = btns[index]
    if (btn.ID === 0) {
      btns.splice(index, 1)
      return
    }
    const res = await canRemoveAuthorityBtnApi({ id: btn.ID })
    if (res.code === 0) {
      btns.splice(index, 1)
    }
  }

  const form = ref({
    ID: 0,
    path: '',
    name: '',
    hidden: false,
    parentId: 0,
    component: '',
    meta: {
      activeName: '',
      title: '',
      icon: '',
      defaultMenu: false,
      closeTab: false,
      keepAlive: false
    },
    parameters: [],
    menuBtn: []
  })
  const changeName = () => {
    form.value.path = form.value.name
  }

  const handleClose = (done) => {
    initForm()
    done()
  }
  // 删除菜单
  const deleteMenu = (ID) => {
    ElMessageBox.confirm(
      t('view.superAdmin.menu.deleteAllRolesConfirm'),
      t('general.hint'),
      {
        confirmButtonText: t('general.confirm'),
        cancelButtonText: t('general.cancel'),
        type: 'warning'
      }
    )
      .then(async () => {
        const res = await deleteBaseMenu({ ID })
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: t('general.deleteSuccess')
          })

          getTableData()
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: t('general.undeleted')
        })
      })
  }
  // 初始化弹窗内表格方法
  const menuForm = ref(null)
  const checkFlag = ref(false)
  const initForm = () => {
    checkFlag.value = false
    menuForm.value.resetFields()
    form.value = {
      ID: 0,
      path: '',
      name: '',
      hidden: false,
      parentId: 0,
      component: '',
      meta: {
        title: '',
        icon: '',
        defaultMenu: false,
        closeTab: false,
        keepAlive: false
      }
    }
  }
  // 关闭弹窗

  const dialogFormVisible = ref(false)
  const closeDialog = () => {
    initForm()
    dialogFormVisible.value = false
  }
  // 添加menu
  const enterDialog = async () => {
    menuForm.value.validate(async (valid) => {
      if (valid) {
        let res
        if (isEdit.value) {
          res = await updateBaseMenu(form.value)
        } else {
          res = await addBaseMenu(form.value)
        }
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: isEdit.value ? t('general.editSuccess') : '添加成功，请到角色管理页面分配权限'
          })
          getTableData()
        }
        initForm()
        dialogFormVisible.value = false
      }
    })
  }

  const menuOption = ref([
    {
      ID: '0',
      title: t('view.superAdmin.menu.rootMenu')
    }
  ])
  const setOptions = () => {
    menuOption.value = [
      {
        ID: '0',
        title: t('view.superAdmin.menu.rootDirctory')
      }
    ]
    setMenuOptions(tableData.value, menuOption.value, false)
  }
  const setMenuOptions = (menuData, optionsData, disabled) => {
    menuData &&
      menuData.forEach((item) => {
        if (item.children && item.children.length) {
          const option = {
            title: item.meta.title,
            ID: item.ID,
            disabled: disabled || item.ID === form.value.ID,
            children: []
          }
          setMenuOptions(
            item.children,
            option.children,
            disabled || item.ID === form.value.ID
          )
          optionsData.push(option)
        } else {
          const option = {
            title: item.meta.title,
            ID: item.ID,
            disabled: disabled || item.ID === form.value.ID
          }
          optionsData.push(option)
        }
      })
  }

  // 添加菜单方法，id为 0则为添加根菜单
  const isEdit = ref(false)
  const dialogTitle = ref(t('view.superAdmin.menu.addMenu'))
  const addMenu = (id) => {
    dialogTitle.value = t('view.superAdmin.menu.addMenu')
    form.value.parentId = id
    isEdit.value = false
    setOptions()
    dialogFormVisible.value = true
  }
  // 修改菜单方法
  const editMenu = async (id) => {
    dialogTitle.value = t('view.superAdmin.menu.editMenu')
    const res = await getBaseMenuById({ id })
    form.value = res.data.menu
    isEdit.value = true
    setOptions()
    dialogFormVisible.value = true
  }
</script>

<style scoped lang="scss">
  .warning {
    color: #dc143c;
  }
  .icon-column {
    display: flex;
    align-items: center;
    .el-icon {
      margin-right: 8px;
    }
  }


  
  .form-tip {
    margin-top: 8px;
    font-size: 12px;
    color: #909399;
    display: flex;
    align-items: center;
    gap: 8px;
    
    .el-icon {
      color: #409eff;
    }
  }
  
  .label-with-tooltip {
    display: flex;
    align-items: center;
    gap: 6px;
    
    .el-icon {
      color: #909399;
      cursor: help;
      
      &:hover {
        color: #409eff;
      }
    }
  }
  
  .parameter-table,
  .button-table {
    border: 1px solid #ebeef5;
    border-radius: 6px;
    
    :deep(.el-table__header) {
      background-color: #fafafa;
    }
    
    :deep(.el-table__body) {
      .el-table__row {
        &:hover {
          background-color: #f5f7fa;
        }
      }
    }
  }
</style>
