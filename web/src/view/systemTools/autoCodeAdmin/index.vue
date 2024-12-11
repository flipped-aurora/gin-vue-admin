<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="goAutoCode(null)">
          新增
        </el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="id" width="60" prop="ID" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="结构体名"
          min-width="150"
          prop="structName"
        />
        <el-table-column
          align="left"
          label="结构体描述"
          min-width="150"
          prop="description"
        />
        <el-table-column
          align="left"
          label="表名称"
          min-width="150"
          prop="tableName"
        />
        <el-table-column
          align="left"
          label="回滚标记"
          min-width="150"
          prop="flag"
        >
          <template #default="scope">
            <el-tag v-if="scope.row.flag" type="danger" effect="dark">
              已回滚
            </el-tag>
            <el-tag v-else type="success" effect="dark"> 未回滚 </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="240">
          <template #default="scope">
            <div>
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="addFuncBtn(scope.row)"
              >
                增加方法
              </el-button>
              <el-button type="primary" link @click="goAutoCode(scope.row, 1)">
                增加字段
              </el-button>
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="openDialog(scope.row)"
              >
                回滚
              </el-button>
              <el-button type="primary" link @click="goAutoCode(scope.row)">
                复用
              </el-button>
              <el-button type="primary" link @click="deleteRow(scope.row)">
                删除
              </el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :title="dialogFormTitle"
      :before-close="closeDialog"
      width="600px"
    >
      <el-form :inline="true" :model="formData" label-width="80px">
        <el-form-item label="选项：">
          <el-checkbox v-model="formData.deleteApi" label="删除接口" />
          <el-checkbox v-model="formData.deleteMenu" label="删除菜单" />
          <el-checkbox
            v-model="formData.deleteTable"
            label="删除表"
            @change="deleteTableCheck"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog"> 取 消 </el-button>
          <el-popconfirm
            title="此操作将回滚生成文件和勾选项目, 是否继续?"
            @confirm="enterDialog"
          >
            <template #reference>
              <el-button type="primary"> 确 定 </el-button>
            </template>
          </el-popconfirm>
        </div>
      </template>
    </el-dialog>

    <el-drawer
      v-model="funcFlag"
      size="60%"
      :show-close="false"
      :close-on-click-modal="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">操作栏</span>
          <div>
            <el-button type="primary" @click="runFunc" :loading="aiLoading">
              生成
            </el-button>
            <el-button type="primary" @click="closeFunc" :loading="aiLoading">
              取消
            </el-button>
          </div>
        </div>
      </template>
      <div class="">
        <el-form
          v-loading="aiLoading"
          label-position="top"
          element-loading-text="小淼正在思考，请稍候..."
          :model="autoFunc"
          label-width="80px"
        >
          <el-row :gutter="12">
            <el-col :span="8">
              <el-form-item label="包名：">
                <el-input
                    v-model="autoFunc.package"
                    placeholder="请输入包名"
                    disabled
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="结构体名：">
                <el-input
                    v-model="autoFunc.structName"
                    placeholder="请输入结构体名"
                    disabled
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="前端文件名：">
                <el-input
                    v-model="autoFunc.packageName"
                    placeholder="请输入文件名"
                    disabled
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="12">
            <el-col :span="8">
              <el-form-item label="后端文件名：">
                <el-input
                    v-model="autoFunc.humpPackageName"
                    placeholder="请输入文件名"
                    disabled
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="描述：">
                <el-input
                    v-model="autoFunc.description"
                    placeholder="请输入描述"
                    disabled
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="缩写：">
                <el-input
                    v-model="autoFunc.abbreviation"
                    placeholder="请输入缩写"
                    disabled
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item label="是否AI填充：">
            <el-switch v-model="autoFunc.isAi" />
            <span class="text-sm text-red-600 p-2"
              >当前ai帮写存在不稳定因素，生成代码后请注意手动调整部分内容</span
            >
          </el-form-item>
          <template v-if="autoFunc.isAi">
            <el-form-item label="Ai帮写:">
              <div class="relative w-full">
                <el-input
                  type="textarea"
                  placeholder="AI帮写功能，输入提示信息，自动生成代码"
                  v-model="autoFunc.prompt"
                  :rows="5"
                  @input="autoFunc.router = autoFunc.router.replace(/\//g, '')"
                />
                <el-button
                  @click="aiAddFunc"
                  type="primary"
                  class="absolute right-2 bottom-2"
                  ><ai-gva />帮写</el-button
                >
              </div>
            </el-form-item>
            <el-form-item label="Api方法:">
              <v-ace-editor
                v-model:value="autoFunc.apiFunc"
                lang="golang"
                theme="github_dark"
                class="h-80 w-full"
              />
            </el-form-item>
            <el-form-item label="Server方法:">
              <v-ace-editor
                v-model:value="autoFunc.serverFunc"
                lang="golang"
                theme="github_dark"
                class="h-80 w-full"
              />
            </el-form-item>
            <el-form-item label="前端JSAPI方法:">
              <v-ace-editor
                v-model:value="autoFunc.jsFunc"
                lang="javascript"
                theme="github_dark"
                class="h-80 w-full"
              />
            </el-form-item>
          </template>

          <el-form-item label="方法介绍：">
            <div class="flex w-full gap-2">
              <el-input
                class="flex-1"
                v-model="autoFunc.funcDesc"
                placeholder="请输入方法介绍"
              />
              <el-button type="primary" @click="autoComplete"
                ><ai-gva />补全</el-button
              >
            </div>
          </el-form-item>
          <el-form-item label="方法名：">
            <el-input
              @blur="autoFunc.funcName = toUpperCase(autoFunc.funcName)"
              v-model="autoFunc.funcName"
              placeholder="请输入方法名"
            />
          </el-form-item>
          <el-form-item label="方法：">
            <el-select v-model="autoFunc.method" placeholder="请选择方法">
              <el-option
                v-for="item in ['GET', 'POST', 'PUT', 'DELETE']"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="是否鉴权：">
            <el-switch
              v-model="autoFunc.isAuth"
              active-text="是"
              inactive-text="否"
            />
          </el-form-item>
          <el-form-item label="路由path:">
            <el-input
              v-model="autoFunc.router"
              placeholder="路由path"
              @input="autoFunc.router = autoFunc.router.replace(/\//g, '')"
            />
            <div>
              API路径: [{{ autoFunc.method }}] /{{ autoFunc.abbreviation }}/{{
                autoFunc.router
              }}
            </div>
          </el-form-item>
        </el-form>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    getSysHistory,
    rollback,
    delSysHistory,
    addFunc,
    butler
  } from '@/api/autoCode.js'
  import { useRouter } from 'vue-router'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref } from 'vue'
  import { formatDate } from '@/utils/format'
  import { toUpperCase } from '@/utils/stringFun'

  import { VAceEditor } from 'vue3-ace-editor'
  import 'ace-builds/src-noconflict/mode-javascript'
  import 'ace-builds/src-noconflict/mode-golang'
  import 'ace-builds/src-noconflict/theme-github_dark'

  defineOptions({
    name: 'AutoCodeAdmin'
  })

  const aiLoading = ref(false)

  const formData = ref({
    id: undefined,
    deleteApi: true,
    deleteMenu: true,
    deleteTable: false
  })

  const router = useRouter()
  const dialogFormVisible = ref(false)
  const dialogFormTitle = ref('')

  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])

  const activeInfo = ref('')

  const autoFunc = ref({
    package: '',
    funcName: '',
    structName: '',
    packageName: '',
    description: '',
    abbreviation: '',
    humpPackageName: '',
    businessDB: '',
    method: '',
    funcDesc: '',
    isAuth: false,
    isAi: false,
    apiFunc: '',
    serverFunc: '',
    jsFunc: ''
  })

  const addFuncBtn = (row) => {
    const req = JSON.parse(row.request)
    activeInfo.value = row.request
    autoFunc.value.package = req.package
    autoFunc.value.structName = req.structName
    autoFunc.value.packageName = req.packageName
    autoFunc.value.description = req.description
    autoFunc.value.abbreviation = req.abbreviation
    autoFunc.value.humpPackageName = req.humpPackageName
    autoFunc.value.businessDB = req.businessDB
    autoFunc.value.method = ''
    autoFunc.value.funcName = ''
    autoFunc.value.router = ''
    autoFunc.value.funcDesc = ''
    autoFunc.value.isAuth = false
    autoFunc.value.isAi = false
    autoFunc.value.apiFunc = ''
    autoFunc.value.serverFunc = ''
    autoFunc.value.jsFunc = ''
    funcFlag.value = true
  }

  const funcFlag = ref(false)

  const closeFunc = () => {
    funcFlag.value = false
  }

  const runFunc = async () => {
    // 首字母自动转换为大写
    autoFunc.value.funcName = toUpperCase(autoFunc.value.funcName)

    if (!autoFunc.value.funcName) {
      ElMessage.error('请输入方法名')
      return
    }
    if (!autoFunc.value.method) {
      ElMessage.error('请选择方法')
      return
    }
    if (!autoFunc.value.router) {
      ElMessage.error('请输入路由')
      return
    }
    if (!autoFunc.value.funcDesc) {
      ElMessage.error('请输入方法介绍')
      return
    }

    if (autoFunc.value.isAi) {
      if (
        !autoFunc.value.apiFunc ||
        !autoFunc.value.serverFunc ||
        !autoFunc.value.jsFunc
      ) {
        ElMessage.error('请先使用AI帮写完成基础代码，如果生成失败请重新调用')
        return
      }
    }

    const res = await addFunc(autoFunc.value)
    if (res.code === 0) {
      ElMessage.success('增加方法成功')
      closeFunc()
    }
  }

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 查询
  const getTableData = async () => {
    const table = await getSysHistory({
      page: page.value,
      pageSize: pageSize.value
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  getTableData()

  const deleteRow = async (row) => {
    ElMessageBox.confirm('此操作将删除本历史, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const res = await delSysHistory({ id: Number(row.ID) })
      if (res.code === 0) {
        ElMessage.success('删除成功')
        getTableData()
      }
    })
  }

  // 打开弹窗
  const openDialog = (row) => {
    dialogFormTitle.value = '回滚：' + row.structName
    formData.value.id = row.ID
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      id: undefined,
      deleteApi: true,
      deleteMenu: true,
      deleteTable: false
    }
  }

  // 确认删除表
  const deleteTableCheck = (flag) => {
    if (flag) {
      ElMessageBox.confirm(
        `此操作将删除自动创建的文件和api（会删除表！！！）, 是否继续?`,
        '提示',
        {
          closeOnClickModal: false,
          distinguishCancelAndClose: true,
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )
        .then(() => {
          ElMessageBox.confirm(
            `此操作将删除自动创建的文件和api（会删除表！！！）, 请继续确认！！！`,
            '会删除表',
            {
              closeOnClickModal: false,
              distinguishCancelAndClose: true,
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }
          ).catch(() => {
            formData.value.deleteTable = false
          })
        })
        .catch(() => {
          formData.value.deleteTable = false
        })
    }
  }

  const enterDialog = async () => {
    const res = await rollback(formData.value)
    if (res.code === 0) {
      ElMessage.success('回滚成功')
      getTableData()
    }
  }

  const goAutoCode = (row, isAdd) => {
    if (row) {
      router.push({
        name: 'autoCodeEdit',
        params: {
          id: row.ID
        },
        query: {
          isAdd: isAdd
        }
      })
    } else {
      router.push({ name: 'autoCode' })
    }
  }

  const aiAddFunc = async () => {
    aiLoading.value = true
    autoFunc.value.apiFunc = ''
    autoFunc.value.serverFunc = ''
    autoFunc.value.jsFunc = ''

    if (!autoFunc.value.prompt) {
      ElMessage.error('请输入提示信息')
      return
    }

    const res = await addFunc({ ...autoFunc.value, isPreview: true })
    if (res.code !== 0) {
      aiLoading.value = false
      ElMessage.error(res.msg)
      return
    }

    const aiRes = await butler({
      structInfo: activeInfo.value,
      template: JSON.stringify(res.data),
      prompt: autoFunc.value.prompt,
      command: 'addFunc'
    })
    aiLoading.value = false
    if (aiRes.code === 0) {
      try {
        const aiData = JSON.parse(aiRes.data)
        autoFunc.value.apiFunc = aiData.api
        autoFunc.value.serverFunc = aiData.server
        autoFunc.value.jsFunc = aiData.js
        autoFunc.value.method = aiData.method
        autoFunc.value.funcName = aiData.funcName
        const routerArr = aiData.router.split('/')
        autoFunc.value.router = routerArr[routerArr.length - 1]
        autoFunc.value.funcDesc = autoFunc.value.prompt
      } catch (_) {
        ElMessage.error('小淼忙碌，请重新调用')
      }
    }
  }

  const autoComplete = async () => {
    aiLoading.value = true
    const aiRes = await butler({
      prompt: autoFunc.value.funcDesc,
      command: 'autoCompleteFunc'
    })
    aiLoading.value = false
    if (aiRes.code === 0) {
      try {
        const aiData = JSON.parse(aiRes.data)
        autoFunc.value.method = aiData.method
        autoFunc.value.funcName = aiData.funcName
        autoFunc.value.router = aiData.router
        autoFunc.value.prompt = autoFunc.value.funcDesc
      } catch (_) {
        ElMessage.error('小淼开小差了，请重新调用')
      }
    }
  }
</script>
