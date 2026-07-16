<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="任务名">
          <el-input v-model="searchInfo.name" placeholder="搜索任务名" />
        </el-form-item>
        <el-form-item label="执行器">
          <el-select v-model="searchInfo.executorType" placeholder="请选择" clearable style="width: 120px">
            <el-option label="注册方法" value="method" />
            <el-option label="HTTP" value="http" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openForm()">新增任务</el-button>
      </div>
      <el-table :data="tableData" style="width: 100%" row-key="ID">
        <el-table-column align="left" label="ID" prop="ID" width="70" />
        <el-table-column align="left" label="任务名" prop="name" min-width="140" show-overflow-tooltip />
        <el-table-column align="left" label="说明" prop="description" min-width="180" show-overflow-tooltip />
        <el-table-column align="left" label="cron" prop="spec" width="120" />
        <el-table-column align="left" label="执行器" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.executorType === 'method' ? 'primary' : 'warning'">
              {{ scope.row.executorType === 'method' ? '注册方法' : 'HTTP' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="启用" width="90">
          <template #default="scope">
            <el-switch :model-value="scope.row.enabled" @change="(v) => onToggle(scope.row, v)" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="下次执行" width="170">
          <template #default="scope">
            {{ scope.row.nextRunAt ? formatDate(scope.row.nextRunAt) : '—' }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="280" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="video-play" @click="onTrigger(scope.row)">触发</el-button>
            <el-button type="primary" link icon="document" @click="openLogs(scope.row)">日志</el-button>
            <el-button type="primary" link icon="edit" @click="openForm(scope.row)">编辑</el-button>
            <el-popover v-model:visible="scope.row.deleteVisible" placement="top" width="170">
              <p>确定删除该任务吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button size="small" type="primary" link @click="scope.row.deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="danger" link icon="delete" @click="scope.row.deleteVisible = true">删除</el-button>
              </template>
            </el-popover>
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

    <!-- 任务表单抽屉(标准抽屉格式) -->
    <el-drawer v-model="formVisible" :size="appStore.drawerSize" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-base">{{ formTitle }}</span>
          <div>
            <el-button @click="formVisible = false">取 消</el-button>
            <el-button type="primary" @click="submitForm">确 定</el-button>
          </div>
        </div>
      </template>
      <el-form ref="formRef" :model="form" label-width="110px">
        <el-form-item label="任务名" required>
          <el-input v-model="form.name" placeholder="唯一任务名" />
        </el-form-item>
        <el-form-item label="说明">
          <el-input v-model="form.description" type="textarea" :rows="2" placeholder="任务用途说明" />
        </el-form-item>
        <el-form-item label="cron 表达式" required>
          <div class="flex gap-2 w-full">
            <el-input v-model="form.spec" placeholder="如 @daily 或 */30 * * * *" />
            <el-select placeholder="常用模板" style="width: 160px" @change="(v) => (form.spec = v)">
              <el-option label="每天零点 @daily" value="@daily" />
              <el-option label="每小时 @hourly" value="@hourly" />
              <el-option label="每周 @weekly" value="@weekly" />
              <el-option label="每30分钟" value="*/30 * * * *" />
              <el-option label="每分钟" value="* * * * *" />
              <el-option label="工作日9点" value="0 9 * * 1-5" />
            </el-select>
          </div>
          <div class="text-xs text-gray-400 mt-1">合法性由后端校验；保存失败时错误会回显</div>
        </el-form-item>
        <el-form-item label="含秒位">
          <el-switch v-model="form.withSeconds" />
          <span class="text-xs text-gray-400 ml-2">开启后为六位表达式(秒 分 时 日 月 周)</span>
        </el-form-item>
        <el-form-item label="执行器" required>
          <el-radio-group v-model="form.executorType">
            <el-radio value="method">调用已注册方法</el-radio>
            <el-radio value="http">HTTP 调用</el-radio>
          </el-radio-group>
        </el-form-item>

        <template v-if="form.executorType === 'method'">
          <el-form-item label="方法" required>
            <el-select v-model="form.methodName" placeholder="选择已注册方法" style="width: 100%">
              <el-option v-for="m in methodOptions" :key="m.name" :label="`${m.name} — ${m.description}`" :value="m.name" />
            </el-select>
            <div class="text-xs text-gray-400 mt-1 flex items-center flex-wrap">
              <span>没有想要的方法？需在后端 server/initialize/timer.go 的 Timer() 中注册，重启服务后即可选择</span>
              <el-popover placement="top" :width="560" trigger="click">
                <template #reference>
                  <el-button type="primary" link size="small">查看注册模板</el-button>
                </template>
                <pre class="text-xs leading-5 whitespace-pre overflow-x-auto m-0 p-2 rounded bg-gray-50 dark:bg-gray-800">{{ methodRegisterTemplate }}</pre>
                <div class="text-right mt-2">
                  <el-button size="small" type="primary" @click="copyRegisterTemplate">复制模板</el-button>
                </div>
              </el-popover>
            </div>
          </el-form-item>
          <el-form-item label="参数(JSON)">
            <el-input v-model="paramsText" type="textarea" :rows="4" placeholder='自由 JSON, 由方法自行解析, 如 {"days": 30}; 留空则不传' />
          </el-form-item>
        </template>

        <template v-if="form.executorType === 'http'">
          <el-form-item label="URL" required>
            <el-input v-model="form.httpUrl" placeholder="https://example.com/hook" />
          </el-form-item>
          <el-form-item label="Method">
            <el-select v-model="form.httpMethod" style="width: 140px">
              <el-option label="GET" value="GET" />
              <el-option label="POST" value="POST" />
              <el-option label="PUT" value="PUT" />
              <el-option label="DELETE" value="DELETE" />
            </el-select>
          </el-form-item>
          <el-form-item label="Header(JSON)">
            <el-input v-model="headerText" type="textarea" :rows="2" placeholder='{"Authorization":"Bearer xxx"}; 留空则不加' />
          </el-form-item>
          <el-form-item label="Body">
            <el-input v-model="form.httpBody" type="textarea" :rows="3" placeholder="请求体原文" />
          </el-form-item>
          <el-form-item label="允许内网">
            <el-switch v-model="form.httpAllowPrivate" />
            <span class="text-xs text-gray-400 ml-2">默认禁止访问内网/环回地址(SSRF 防护)，仅确有需要时开启</span>
          </el-form-item>
        </template>

        <el-form-item label="启用">
          <el-switch v-model="form.enabled" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 执行日志抽屉(标准抽屉格式) -->
    <el-drawer v-model="logVisible" :size="appStore.drawerSize" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-base">执行日志：{{ logTaskName }}</span>
          <div>
            <el-button @click="logVisible = false">关 闭</el-button>
          </div>
        </div>
      </template>
      <el-table :data="logData" style="width: 100%">
        <el-table-column type="expand">
          <template #default="scope">
            <div class="px-4 py-2">
              <p v-if="scope.row.errorMsg" class="text-red-500 break-all">错误：{{ scope.row.errorMsg }}</p>
              <p v-if="scope.row.output" class="break-all mt-1">输出：{{ scope.row.output }}</p>
              <p v-if="!scope.row.errorMsg && !scope.row.output" class="text-gray-400">无详情</p>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="触发" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.triggerType === 'auto' ? 'info' : 'warning'">
              {{ scope.row.triggerType === 'auto' ? '调度' : '手动' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" width="90">
          <template #default="scope">
            <el-tag :type="scope.row.status === 'success' ? 'success' : 'danger'">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="开始时间" width="170">
          <template #default="scope">{{ formatDate(scope.row.startedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="耗时(ms)" prop="durationMs" width="100" />
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="logPage"
          :page-size="logPageSize"
          :total="logTotal"
          layout="total, prev, pager, next"
          @current-change="handleLogPageChange"
        />
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createTimedTask,
  updateTimedTask,
  deleteTimedTask,
  toggleTimedTask,
  triggerTimedTask,
  getTimedTaskList,
  getTimedTaskLogList,
  getRegisteredMethods
} from '@/api/timedTask'
import { useTimedTaskAlert } from './useAlertStream'
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'
import { useAppStore } from '@/pinia'

defineOptions({ name: 'TimedTask' })

const appStore = useAppStore()
useTimedTaskAlert() // 页内订阅失败告警

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const getTableData = async () => {
  const res = await getTimedTaskList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}

const onSubmit = () => {
  page.value = 1
  getTableData()
}
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// ------- 表单抽屉 -------
const formVisible = ref(false)
const formRef = ref(null)
const emptyForm = () => ({
  ID: 0,
  name: '',
  description: '',
  spec: '',
  withSeconds: false,
  executorType: 'method',
  methodName: '',
  params: null,
  httpUrl: '',
  httpMethod: 'GET',
  httpHeader: null,
  httpBody: '',
  httpAllowPrivate: false,
  enabled: true
})
const form = ref(emptyForm())
const paramsText = ref('')
const headerText = ref('')
const methodOptions = ref([])
const formTitle = computed(() => (form.value.ID ? '编辑定时任务' : '新增定时任务'))

// 注册方法代码模板: 引导二开在 server/initialize/timer.go 的 Timer() 中追加 task.Register
const methodRegisterTemplate = `// 位置: server/initialize/timer.go -> Timer()
task.Register("MyTask", "任务说明(展示在方法下拉中)", func(ctx context.Context, params json.RawMessage) error {
	// params 为面板「参数(JSON)」原文, 按需解析; 无参数可忽略
	// var p struct {
	// 	Days int \`json:"days"\`
	// }
	// if len(params) > 0 {
	// 	if err := json.Unmarshal(params, &p); err != nil {
	// 		return err
	// 	}
	// }
	// TODO: 业务逻辑; 数据库操作请使用 global.GVA_DB.WithContext(ctx)
	return nil
})`

const copyRegisterTemplate = () => {
  const input = document.createElement('textarea')
  input.value = methodRegisterTemplate
  document.body.appendChild(input)
  input.select()
  document.execCommand('copy')
  document.body.removeChild(input)
  ElMessage.success('已复制, 请粘贴到 server/initialize/timer.go 的 Timer() 中, 重启后生效')
}

const openForm = async (row) => {
  form.value = row ? { ...emptyForm(), ...row } : emptyForm()
  paramsText.value = row && row.params ? JSON.stringify(row.params, null, 2) : ''
  headerText.value = row && row.httpHeader ? JSON.stringify(row.httpHeader, null, 2) : ''
  formVisible.value = true
  if (methodOptions.value.length === 0) {
    const res = await getRegisteredMethods()
    if (res.code === 0) {
      methodOptions.value = res.data.methods || []
    }
  }
}

// parseJsonField 本地 JSON.parse 校验; 空串返回 null
const parseJsonField = (text, label) => {
  if (!text || !text.trim()) return { ok: true, value: null }
  try {
    return { ok: true, value: JSON.parse(text) }
  } catch {
    ElMessage.error(`${label} 不是合法 JSON`)
    return { ok: false, value: null }
  }
}

const submitForm = async () => {
  if (!form.value.name || !form.value.spec) {
    ElMessage.warning('任务名与 cron 表达式必填')
    return
  }
  const p = parseJsonField(paramsText.value, '参数')
  if (!p.ok) return
  const h = parseJsonField(headerText.value, 'Header')
  if (!h.ok) return
  const payload = { ...form.value, params: p.value, httpHeader: h.value }
  const res = form.value.ID ? await updateTimedTask(payload) : await createTimedTask(payload)
  if (res.code === 0) {
    ElMessage.success(form.value.ID ? '更新成功' : '创建成功')
    formVisible.value = false
    getTableData()
  }
}

// ------- 行操作 -------
const onToggle = async (row, enabled) => {
  const res = await toggleTimedTask({ ID: row.ID, enabled })
  if (res.code === 0) {
    ElMessage.success(enabled ? '已启用' : '已停用')
    getTableData()
  }
}

const onTrigger = async (row) => {
  const res = await triggerTimedTask({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('已触发, 稍后可在日志中查看结果')
  }
}

const onDelete = async (row) => {
  row.deleteVisible = false
  const res = await deleteTimedTask({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    if (tableData.value.length === 1 && page.value > 1) page.value--
    getTableData()
  }
}

// ------- 日志抽屉 -------
const logVisible = ref(false)
const logData = ref([])
const logTotal = ref(0)
const logPage = ref(1)
const logPageSize = ref(10)
const logTaskId = ref(0)
const logTaskName = ref('')

const loadLogs = async () => {
  const res = await getTimedTaskLogList({ page: logPage.value, pageSize: logPageSize.value, taskId: logTaskId.value })
  if (res.code === 0) {
    logData.value = res.data.list
    logTotal.value = res.data.total
  }
}

const openLogs = (row) => {
  logTaskId.value = row.ID
  logTaskName.value = row.name
  logPage.value = 1
  logVisible.value = true
  loadLogs()
}

const handleLogPageChange = (val) => {
  logPage.value = val
  loadLogs()
}

getTableData()
</script>

<style scoped></style>
