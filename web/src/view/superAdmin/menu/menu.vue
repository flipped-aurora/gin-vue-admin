<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="addMenu('0')">{{ t('menu.addRootMenu') }}</el-button>
      </div>

      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="100" prop="ID" />
        <el-table-column align="left" :label="t('menu.routeName')" show-overflow-tooltip min-width="160" prop="name" />
        <el-table-column align="left" :label="t('menu.routePath')" show-overflow-tooltip min-width="160" prop="path" />
        <el-table-column align="left" :label="t('menu.visibility')" min-width="100" prop="hidden">
          <template #default="scope">
            <span>{{ scope.row.hidden? t('menu.hide') : t('menu.show') }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('menu.parent')" min-width="90" prop="parentId" />
        <el-table-column align="left" :label="t('menu.sort')" min-width="70" prop="sort" />
        <el-table-column align="left" :label="t('menu.filePath')" min-width="360" prop="component" />
        <el-table-column align="left" :label="t('menu.displayName')" min-width="120" prop="authorityName">
          <template #default="scope">
            <span>{{ scope.row.meta.title }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('menu.icon')" min-width="140" prop="authorityName">
          <template #default="scope">
            <div class="icon-column">
              <el-icon>
                <component :is="scope.row.meta.icon" />
              </el-icon>
              <span>{{ scope.row.meta.icon }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" fixed="right" :label="t('general.operations')" width="300">
          <template #default="scope">
            <el-button
              size="mini"
              type="text"
              icon="plus"
              @click="addMenu(scope.row.ID)"
            >{{ t('menu.addSubMenu') }}</el-button>
            <el-button
              size="mini"
              type="text"
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >{{ t('general.edit') }}</el-button>
            <el-button
              size="mini"
              type="text"
              icon="delete"
              @click="deleteMenu(scope.row.ID)"
            >{{ t('general.delete') }}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="handleClose" :title="dialogTitle" width="800px">
      <warning-bar :title="t('menu.newMenuNote')" />
      <el-form
        v-if="dialogFormVisible"
        ref="menuForm"
        :inline="true"
        :model="form"
        :rules="rules"
        label-position="top"
        label-width="100px"
      >
        <el-form-item :label="t('menu.routeName')" prop="path" style="width:30%">
          <el-input
            v-model="form.name"
            autocomplete="off"
            :placeholder="t('menu.routeNameNote')"
            @change="changeName"
          />
        </el-form-item>
        <el-form-item prop="path" style="width:30%">
          <template #label>
            <div style="display:inline-flex">
              {{ t('menu.routePath') }}
              <el-checkbox v-model="checkFlag" style="float:right;margin-left:20px;">{{ t('menu.addParameter') }}</el-checkbox>
            </div>
          </template>

          <el-input
            v-model="form.path"
            :disabled="!checkFlag"
            autocomplete="off"
            :placeholder="t('menu.routePathNote')"
          />
        </el-form-item>
        <el-form-item :label="t('menu.visibility')" style="width:20%">
          <el-select v-model="form.hidden" :placeholder="t('menu.visibilityNote')">
            <el-option :value="false" :label="t('general.no')" />
            <el-option :value="true" :label="t('general.yes')" />
          </el-select>
        </el-form-item>
        <el-form-item :label="t('menu.parentId')" style="width:30%">
          <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :disabled="!isEdit"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item :label="t('menu.filePath')" prop="component" style="width:60%">
          <el-input v-model="form.component" autocomplete="off" />
          <span style="font-size:12px;margin-right:12px;">{{ t('menu.subMenuNote') }}</span><el-button size="mini" @click="form.component = 'view/routerHolder.vue'">{{ t('menu.clickMe') }}</el-button>
        </el-form-item>
        <el-form-item :label="t('menu.displayName')" prop="meta.title" style="width:30%">
          <el-input v-model="form.meta.title" autocomplete="off" />
        </el-form-item>
        <el-form-item :label="t('menu.icon')" prop="meta.icon" style="width:30%">
          <icon :meta="form.meta" style="width:100%" />
        </el-form-item>
        <el-form-item :label="t('general.order')" prop="sort" style="width:30%">
          <el-input v-model.number="form.sort" autocomplete="off" />
        </el-form-item>
        <el-form-item label="KeepAlive" prop="meta.keepAlive" style="width:30%">
          <el-select v-model="form.meta.keepAlive" style="width:100%" :placeholder="t('menu.keepAliveNote')">
            <el-option :value="false" :label="t('general.no')" />
            <el-option :value="true" :label="t('general.yes')" />
          </el-select>
        </el-form-item>
        <el-form-item label="CloseTab" prop="meta.closeTab" style="width:30%">
          <el-select v-model="form.meta.closeTab" style="width:100%" :placeholder="t('menu.closeTabNote')">
            <el-option :value="false" :label="t('general.no')" />
            <el-option :value="true" :label="t('general.yes')" />
          </el-select>
        </el-form-item>
      </el-form>
      <div>
        <el-button
          size="small"
          type="primary"
          icon="edit"
          @click="addParameter(form)"
        >{{ t('menu.addMenuParameters') }}</el-button>
        <el-table :data="form.parameters" style="width: 100%">
          <el-table-column align="left" prop="type" :label="t('menu.parameterType')" width="180">
            <template #default="scope">
              <el-select v-model="scope.row.type" :placeholder="t('general.pleaseSelect')">
                <el-option key="query" value="query" label="query" />
                <el-option key="params" value="params" label="params" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="key" :label="t('menu.paremeterKey')" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.key" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="value" :label="t('menu.parameterValue')">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.value" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"
                  size="small"
                  icon="delete"
                  @click="deleteParameter(form.parameters,scope.$index)"
                >{{ t('general.delete') }}</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>

        <el-button
          style="margin-top:12px"
          size="small"
          type="primary"
          icon="edit"
          @click="addBtn(form)"
        >{{ t('menu.addButton') }}</el-button>
        <el-table :data="form.menuBtn" style="width: 100%">
          <el-table-column align="left" prop="name" :label="t('menu.buttonName')" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.name" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="name" :label="t('menu.comments')" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.desc" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"
                  size="small"
                  icon="delete"
                  @click="deleteBtn(form.menuBtn,scope.$index)"
                >{{ t('general.delete') }}</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">{{ t('general.cancel') }}</el-button>
          <el-button size="small" type="primary" @click="enterDialog">{{ t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>
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
import warningBar from '@/components/warningBar/warningBar.vue'
import { canRemoveAuthorityBtnApi } from '@/api/authorityBtn'
import { reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const rules = reactive({
  path: [{ required: true, message: t('menu.enterMenuNameNote'), trigger: 'blur' }],
  component: [
    { required: true, message: t('menu.enterFilePathNote'), trigger: 'blur' }
  ],
  'meta.title': [
    { required: true, message: t('menu.enterMenuDisplayNameNote'), trigger: 'blur' }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})
// 查询
const getTableData = async() => {
  const table = await getMenuList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
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
// 删除参数
const deleteParameter = (parameters, index) => {
  parameters.splice(index, 1)
}

// 新增可控按钮
const addBtn = (form) => {
  console.log(form)
  if (!form.menuBtn) {
    form.menuBtn = []
  }
  form.menuBtn.push({
    name: '',
    desc: '',
  })
}
// 删除可控按钮
const deleteBtn = async(btns, index) => {
  const btn = btns[index]
  if (btn.ID === 0) {
    btns.splice(index, 1)
    return
  }
  const res = await canRemoveAuthorityBtnApi({ id: btn.ID })
  if (res.code === 0) {
    btns.splice(index, 1)
    return
  }
}

const form = ref({
  ID: 0,
  path: '',
  name: '',
  hidden: '',
  parentId: '',
  component: '',
  meta: {
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
  ElMessageBox.confirm(t('menu.deleteAllRolesConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteBaseMenu({ ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess')
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
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
    hidden: '',
    parentId: '',
    component: '',
    meta: {
      title: '',
      icon: '',
      defaultMenu: false,
      keepAlive: ''
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
const enterDialog = async() => {
  menuForm.value.validate(async valid => {
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
          message: isEdit.value ? t('general.editSuccess') : t('general.addSuccess')
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
    title: t('menu.rootMenu')
  }
])
const setOptions = () => {
  menuOption.value = [
    {
      ID: '0',
      title: t('menu.rootDirctory')
    }
  ]
  setMenuOptions(tableData.value, menuOption.value, false)
}
const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
        menuData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
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
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID
            }
            optionsData.push(option)
          }
        })
}

// 添加菜单方法，id为 0则为添加根菜单
const isEdit = ref(false)
const dialogTitle = ref(t('menu.addMenu'))
const addMenu = (id) => {
  dialogTitle.value = t('menu.addMenu')
  form.value.parentId = String(id)
  isEdit.value = false
  setOptions()
  dialogFormVisible.value = true
}
// 修改菜单方法
const editMenu = async(id) => {
  dialogTitle.value = t('menu.editMenu')
  const res = await getBaseMenuById({ id })
  form.value = res.data.menu
  isEdit.value = true
  setOptions()
  dialogFormVisible.value = true
}

</script>

<script>
export default {
  name: 'Menus',
}
</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
.icon-column{
  display: flex;
  align-items: center;
  .el-icon{
    margin-right: 8px;
  }
}
</style>
