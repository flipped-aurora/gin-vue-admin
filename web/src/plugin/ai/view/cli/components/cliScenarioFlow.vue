<template>
  <el-drawer
    v-model="visible"
    direction="rtl"
    size="80%"
    :show-close="false"
    :append-to-body="false"
    @open="onOpen"
    @opened="onOpened"
    @closed="onClosed"
  >
    <template #header>
      <div class="flex justify-between items-center w-full">
        <span class="text-lg">调用场景编排 - {{ cli.name || cli.command || '' }}</span>
        <div>
          <el-button @click="visible = false">取 消</el-button>
          <el-button type="primary" :loading="saving" @click="onSave">确 定</el-button>
        </div>
      </div>
    </template>

    <div class="flex-1 flex flex-col min-h-0">
      <div class="scenario-bar flex items-center gap-2 flex-wrap mb-3">
        <el-select v-model="activeIndex" placeholder="选择场景" class="!w-[200px]">
          <el-option v-for="(s, i) in scenarios" :key="i" :label="s.name || '未命名'" :value="i" />
        </el-select>
        <el-button icon="plus" @click="addScenario">新建场景</el-button>
        <el-button icon="delete" @click="removeScenario" :disabled="scenarios.length === 0">删除当前</el-button>
        <el-input v-if="current" v-model="current.name" placeholder="场景名" class="!w-[180px]" />
        <el-input v-if="current" v-model="current.description" placeholder="场景说明" class="!w-[260px]" />
      </div>

      <div class="scenario-toolbar flex items-center gap-2 flex-wrap mb-2" v-if="current">
        <el-button icon="plus" @click="addNode('command')">命令节点</el-button>
        <el-button icon="plus" @click="addNode('decision')">判断节点</el-button>
        <el-divider direction="vertical" />
        <el-button-group>
          <el-button icon="full-screen" @click="fitView">适应画布</el-button>
          <el-button icon="aim" @click="centerView">居中</el-button>
          <el-button icon="zoom-in" @click="zoomIn" />
          <el-button icon="zoom-out" @click="zoomOut" />
          <el-button icon="refresh-right" @click="resetView">重置</el-button>
        </el-button-group>
        <el-dropdown v-if="current.nodes && current.nodes.length" trigger="click" @command="focusNode">
          <el-button icon="position">定位节点<el-icon class="el-icon--right"><arrow-down /></el-icon></el-button>
          <template #dropdown>
            <el-dropdown-menu class="max-h-[320px] overflow-y-auto">
              <el-dropdown-item v-for="n in current.nodes" :key="n.id" :command="n.id">{{ nodeLabel(n) }}</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
        <span class="text-xs text-[#909399]">拖节点锚点手动连线；点节点编辑详情；点连线编辑条件</span>
      </div>

      <div class="scenario-main flex gap-3 min-h-0" :style="{ height: canvasHeight + 'px' }">
        <div ref="canvasRef" class="flex-1 min-w-0 h-full border border-solid border-[#ebeef5] bg-[#fafafa] rounded box-border overflow-hidden"></div>

        <div class="scenario-detail" v-if="selectedNode">
          <div class="flex items-center gap-2 mb-2.5 font-semibold">
            <el-tag :type="selectedNode.type === 'decision' ? 'warning' : 'success'" size="small">{{ selectedNode.type }}</el-tag>
            <span>编辑节点</span>
            <el-button icon="delete" link type="danger" class="ml-auto" @click="removeSelected">删除节点</el-button>
          </div>
          <div class="mb-2.5">
            <el-input v-model="selectedNode.alias" class="w-full" :class="{ 'alias-dup': isAliasDup }" placeholder="如 getUser，场景内唯一；用于 别名.字段 引用" @input="syncSelectedToLf">
              <template #prepend>别名</template>
            </el-input>
            <div v-if="isAliasDup" class="text-[#f56c6c] text-xs mt-1">场景内已存在</div>
          </div>
          <div class="mb-2.5 text-[13px] leading-[1.8] text-[#909399]" v-if="aliasHints.length">
            <span class="text-[#606266] font-semibold">可用别名：</span>
            <span v-for="h in aliasHints" :key="h.alias" class="inline-block mr-3">{{ h.alias }}({{ h.label }}): {{ h.fields }}</span>
          </div>
          <template v-if="selectedNode.type === 'command'">
            <div class="mb-2.5">
              <el-select v-model="selectedNode.commandName" placeholder="选择命令" class="w-full" @change="syncSelectedToLf">
                <el-option v-for="c in commandOptions" :key="c" :label="c" :value="c" />
              </el-select>
            </div>
            <div class="mb-2.5">
              <el-input v-model="selectedNode.note" type="textarea" :rows="3" class="w-full" placeholder="说明" @input="syncSelectedToLf" />
            </div>
            <div class="mb-2.5">
              <el-input v-model="selectedNode.inputNote" placeholder="入参来源(如:create.id + query.token)" class="w-full" @input="syncSelectedToLf" />
            </div>
            <div class="bg-[#f5f7fa] rounded px-2.5 py-2" v-if="selectedNode.commandName && commandMap[selectedNode.commandName]">
              <div class="text-[13px] leading-[1.8]"><span class="text-[#606266] font-semibold">入参：</span>{{ formatParams(commandMap[selectedNode.commandName].parameters) }}</div>
              <div class="text-[13px] leading-[1.8]"><span class="text-[#606266] font-semibold">出参：</span>{{ formatFields(commandMap[selectedNode.commandName].response) }}</div>
            </div>
            <div class="text-[13px] leading-[1.8] text-[#909399]" v-else-if="selectedNode.commandName">该命令暂无参数信息</div>
          </template>
          <template v-else>
            <div class="mb-2.5">
              <el-input v-model="selectedNode.note" type="textarea" :rows="3" class="w-full" placeholder="合并判断描述（如：基于上游 A.x、B.y、C.z 合并，判断...）" @input="syncSelectedToLf" />
            </div>
            <div class="text-[13px] leading-[1.8] text-[#909399]">判断节点不执行命令；出边带条件实现分支（在下方添加连线并填条件）。</div>
          </template>

          <el-divider content-position="left">出向连线</el-divider>
          <div class="mb-2.5">
            <el-select v-model="connectedTargetIds" multiple filterable collapse-tags collapse-tags-tooltip placeholder="选择要连线的目标节点（可多选，已连的会标记）" class="w-full" @change="onEdgeTargetsChange">
              <el-option v-for="o in edgeTargetOptions" :key="o.id" :label="o.label" :value="o.id" />
            </el-select>
          </div>
          <div v-if="outgoingEdges.length" class="flex flex-col gap-2.5">
            <div v-for="e in outgoingEdges" :key="e.id" class="border border-solid border-[#ebeef5] rounded p-2 bg-[#fafafa]">
              <div class="flex items-center gap-1.5 mb-1.5 text-[13px]">
                <span class="text-[#409eff] font-semibold">→</span>
                <span class="flex-1 truncate" :title="nodeLabel(nodeById(e.to))">{{ nodeLabel(nodeById(e.to)) }}</span>
                <el-button icon="delete" link type="danger" size="small" @click="deleteEdge(e)" />
              </div>
              <el-input v-model="e.condition" size="small" class="w-full" placeholder="分支条件（留空=默认流转）" @input="updateEdgeCondition(e)" />
            </div>
          </div>
          <div v-else class="text-[13px] leading-[1.8] text-[#909399]">暂无出向连线，可在上方选择目标节点添加</div>
        </div>
        <div class="scenario-detail empty" v-else>点击画布节点编辑详情，点空白取消选择</div>
      </div>
    </div>
  </el-drawer>
</template>

<script setup>
import { ref, computed, nextTick, watch, onMounted, onUnmounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import LogicFlow, { RectNode, RectNodeModel } from '@logicflow/core'
import '@logicflow/core/dist/index.css'
import { getCliDetail, updateCli, previewManifest } from '@/plugin/ai/api/cli'

class CommandModel extends RectNodeModel {
  initNodeData(data) {
    super.initNodeData(data)
    this.width = 200
    this.height = 44
  }
  getNodeStyle() {
    const style = super.getNodeStyle()
    style.fill = '#f0f9eb'
    style.stroke = '#67c23a'
    style.radius = 8
    return style
  }
}
class DecisionModel extends RectNodeModel {
  initNodeData(data) {
    super.initNodeData(data)
    this.width = 200
    this.height = 44
  }
  getNodeStyle() {
    const style = super.getNodeStyle()
    style.fill = '#fdf6ec'
    style.stroke = '#e6a23c'
    style.radius = 8
    return style
  }
}

const props = defineProps({
  modelValue: Boolean,
  cli: { type: Object, default: () => ({}) }
})
const emit = defineEmits(['update:modelValue'])
const visible = computed({
  get: () => props.modelValue,
  set: (v) => emit('update:modelValue', v)
})

const scenarios = ref([])
const activeIndex = ref(undefined)
const current = computed(() => scenarios.value[activeIndex.value])
const commandOptions = ref([])
const commandMap = ref({})
const canvasRef = ref(null)
const canvasHeight = ref(400)
const saving = ref(false)
const cliData = ref({})
const selectedNodeId = ref(null)
const selectedNode = computed(() => {
  if (!current.value || !selectedNodeId.value) return null
  return current.value.nodes.find((n) => n.id === selectedNodeId.value) || null
})
const isAliasDup = computed(() => {
  const n = selectedNode.value
  if (!n) return false
  const alias = (n.alias || '').trim()
  if (!alias) return false
  return (current.value.nodes || []).some((o) => o.id !== n.id && (o.alias || '').trim() === alias)
})
const aliasHints = computed(() => {
  if (!current.value) return []
  return (current.value.nodes || [])
    .filter((n) => n.type === 'command' && (n.alias || '').trim())
    .map((n) => {
      const cmd = n.commandName && commandMap.value[n.commandName]
      const fields = cmd ? formatFields(cmd.response) : ''
      return { alias: n.alias.trim(), label: n.commandName || '未选命令', fields: fields || '无' }
    })
})
const defaultAlias = (kind) => {
  const prefix = kind === 'decision' ? 'check' : 'step'
  let max = 0
  ;(current.value.nodes || []).forEach((n) => {
    const a = n.alias || ''
    if (a.startsWith(prefix)) {
      const num = parseInt(a.slice(prefix.length), 10)
      if (!isNaN(num) && num > max) max = num
    }
  })
  return prefix + (max + 1)
}
const findDupAlias = () => {
  for (const s of scenarios.value) {
    const seen = {}
    for (const n of (s.nodes || [])) {
      const a = (n.alias || '').trim()
      if (!a) continue
      if (seen[a]) return { scenario: s.name || '未命名', alias: a }
      seen[a] = true
    }
  }
  return null
}
let lf = null

const MIN_CANVAS_HEIGHT = 320
const CANVAS_BOTTOM_GAP = 8 // 底部留一点空隙，保证画布下边框可见
// JS 计算画布高度：用抽屉内容区(el-drawer__body)的真实底部，减去画布距顶部的距离，
// 这样无论上方工具栏占多高、有没有 footer，都能精确铺满剩余高度，不留空隙。
const computeCanvasHeight = () => {
  if (!canvasRef.value) return
  const top = canvasRef.value.getBoundingClientRect().top
  const body = canvasRef.value.closest('.el-drawer__body')
  let bottom = window.innerHeight
  if (body) {
    const rect = body.getBoundingClientRect()
    const padBottom = parseFloat(getComputedStyle(body).paddingBottom) || 0
    bottom = rect.bottom - padBottom
  }
  const h = bottom - top - CANVAS_BOTTOM_GAP
  canvasHeight.value = h > MIN_CANVAS_HEIGHT ? Math.floor(h) : MIN_CANVAS_HEIGHT
}
// 把画布真实内部尺寸同步给 LogicFlow（clientHeight 不含边框，画好的图不会盖住下边框）
const syncCanvasSize = () => {
  if (!lf || !canvasRef.value) return
  const w = canvasRef.value.clientWidth
  const h = canvasRef.value.clientHeight
  if (w > 0 && h > 0) lf.resize(w, h)
}
const onWindowResize = () => {
  if (!visible.value) return
  computeCanvasHeight()
  nextTick(syncCanvasSize)
}

// 画布视图操作
const fitView = () => { if (lf) lf.fitView(30, 30) }            // 缩放平移到刚好容纳全部节点
const centerView = () => { if (lf) lf.translateCenter() }       // 图形整体居中
const zoomIn = () => { if (lf) lf.zoom(true) }                  // 放大
const zoomOut = () => { if (lf) lf.zoom(false) }               // 缩小
const resetView = () => { if (lf) { lf.resetZoom(); lf.translateCenter() } } // 重置缩放并居中
const focusNode = (id) => { if (lf && id) lf.focusOn({ id }) }  // 定位到指定节点
const nodeLabel = (n) => {
  if (!n) return '未知节点'
  const alias = (n.alias || '').trim()
  const name = n.type === 'decision' ? ('判断：' + (n.note || '未描述').slice(0, 12)) : (n.commandName || '未选命令')
  return alias ? `${alias} · ${name}` : name
}
const nodeById = (id) => (current.value && current.value.nodes || []).find((n) => n.id === id)

// 连线编辑：画布是连线的真实数据源，这里维护一份响应式镜像供右侧面板读写
const edgeList = ref([])
const connectedTargetIds = ref([]) // 多选框的值=当前节点已连线的目标，已连的会被标记选中，不可重复连
const refreshEdges = () => {
  if (!lf) { edgeList.value = []; return }
  let g
  try { g = lf.getGraphData() } catch (e) { edgeList.value = []; return }
  edgeList.value = (g.edges || []).map((e) => ({
    id: e.id,
    from: e.sourceNodeId,
    to: e.targetNodeId,
    condition: (e.properties && e.properties.condition) || '',
    note: (e.properties && e.properties.note) || ''
  }))
}
// 当前选中节点的出向连线
const outgoingEdges = computed(() => edgeList.value.filter((e) => e.from === selectedNodeId.value))
// 多选框始终跟随真实连线（手动拖线、切换节点、增删都会同步勾选状态）
watch(outgoingEdges, (edges) => { connectedTargetIds.value = [...new Set(edges.map((e) => e.to))] }, { immediate: true })
// 可连线的目标节点（排除自己）
const edgeTargetOptions = computed(() =>
  (current.value && current.value.nodes || [])
    .filter((n) => n.id !== selectedNodeId.value)
    .map((n) => ({ id: n.id, label: nodeLabel(n) }))
)
// 多选变化时，对比已有连线，新增的勾选建连线、取消的勾选删连线（同一目标不会重复连）
const onEdgeTargetsChange = (newIds) => {
  if (!lf || !selectedNodeId.value) return
  const currentIds = outgoingEdges.value.map((e) => e.to)
  const toAdd = newIds.filter((id) => !currentIds.includes(id))
  const toRemove = currentIds.filter((id) => !newIds.includes(id))
  toAdd.forEach((tid) => {
    try {
      lf.addEdge({
        sourceNodeId: selectedNodeId.value,
        targetNodeId: tid,
        type: 'polyline',
        text: '默认流转',
        properties: { condition: '', note: '' }
      })
    } catch (e) { /* ignore */ }
  })
  toRemove.forEach((tid) => {
    const edge = outgoingEdges.value.find((e) => e.to === tid)
    if (edge) { try { lf.deleteEdge(edge.id) } catch (e) { /* ignore */ } }
  })
  refreshEdges()
}
const updateEdgeCondition = (edge) => {
  if (!lf) return
  try {
    lf.setProperties(edge.id, { condition: edge.condition, note: edge.note })
    lf.updateText(edge.id, edge.condition || '默认流转')
  } catch (e) { /* ignore */ }
}
const deleteEdge = (edge) => {
  if (!lf) return
  try { lf.deleteEdge(edge.id) } catch (e) { /* ignore */ }
  refreshEdges()
}

onMounted(() => window.addEventListener('resize', onWindowResize))
onUnmounted(() => window.removeEventListener('resize', onWindowResize))

const uid = () => 'n_' + Date.now().toString(36) + Math.random().toString(36).slice(2, 5)
const formatParams = (params) => {
  if (!params || !params.length) return '无'
  return params.map((p) => `${p.flag || p.name}${p.required ? '*' : ''}:${p.type || 'string'}`).join('  ')
}
const formatFields = (fields) => {
  if (!fields || !fields.length) return '无'
  return fields.map((f) => f.name).join(', ')
}
const buildText = (n) => {
  if (!n) return ''
  if (n.type === 'decision') return '判断：' + (n.note || '未描述').slice(0, 20)
  return n.commandName || '未选命令'
}

const onOpen = async () => {
  if (!props.cli.id && !props.cli.ID) {
    ElMessage.warning('请先保存 CLI 基础信息')
    visible.value = false
    return
  }
  const id = props.cli.ID || props.cli.id
  const res = await getCliDetail({ id })
  if (res.code !== 0) return
  cliData.value = res.data.cli || {}
  const manifestRes = await previewManifest({ cliId: id })
  const commands = manifestRes.code === 0 ? (manifestRes.data.commands || []) : []
  commandOptions.value = commands.map((c) => c.name).filter(Boolean)
  const map = {}
  commands.forEach((c) => { map[c.name] = { parameters: c.parameters || [], response: c.response || [] } })
  commandMap.value = map
  let parsed = []
  try {
    parsed = JSON.parse(res.data.cli.scenariosJson || '[]')
  } catch (e) {
    parsed = []
  }
  scenarios.value = (Array.isArray(parsed) ? parsed : []).map((s) => ({
    name: s.name || '未命名',
    description: s.description || '',
    sort: s.sort || 0,
    nodes: Array.isArray(s.nodes) ? s.nodes : [],
    edges: Array.isArray(s.edges) ? s.edges : []
  }))
  activeIndex.value = undefined
  selectedNodeId.value = null
  await nextTick()
  renderCanvas()
}

const renderCanvas = () => {
  if (!canvasRef.value) return
  if (lf) { try { lf.destroy && lf.destroy() } catch (e) { /* ignore */ }; lf = null }
  canvasRef.value.innerHTML = ''
  computeCanvasHeight()
  lf = new LogicFlow({ container: canvasRef.value, grid: true, height: canvasHeight.value })
  lf.register({ type: 'command-node', view: RectNode, model: CommandModel })
  lf.register({ type: 'decision-node', view: RectNode, model: DecisionModel })
  const s = current.value || { nodes: [], edges: [] }
  const nodes = s.nodes.map((n) => ({
    id: n.id,
    type: n.type === 'decision' ? 'decision-node' : 'command-node',
    x: n.x || 140,
    y: n.y || 60,
    text: buildText(n),
    properties: { commandName: n.commandName || '', inputNote: n.inputNote || '', note: n.note || '', alias: n.alias || '' }
  }))
  const edges = s.edges.map((e) => ({
    sourceNodeId: e.from,
    targetNodeId: e.to,
    type: 'polyline',
    text: e.condition || '默认流转',
    properties: { condition: e.condition || '', note: e.note || '' }
  }))
  lf.render({ nodes, edges })
  lf.on('node:click', ({ data }) => { selectedNodeId.value = data.id })
  lf.on('blank:click', () => { selectedNodeId.value = null })
  lf.on('edge:add', ({ data }) => {
    // 同一对节点之间不允许重复连线（手动拖线也拦截）
    let g
    try { g = lf.getGraphData() } catch (e) { g = { edges: [] } }
    const dup = (g.edges || []).filter((e) => e.sourceNodeId === data.sourceNodeId && e.targetNodeId === data.targetNodeId)
    if (dup.length > 1) {
      try { lf.deleteEdge(data.id) } catch (e) { /* ignore */ }
      ElMessage.warning('该连线已存在，不能重复连线')
      refreshEdges()
      return
    }
    // 用户新建的边（无 condition 属性）默认显示提示，引导点击编辑
    if (!data.properties || data.properties.condition === undefined) {
      lf.setProperties(data.id, { condition: '', note: '' })
      try { lf.updateText(data.id, '默认流转') } catch (e) { /* ignore */ }
    }
    refreshEdges()
  })
  lf.on('edge:delete', () => { refreshEdges() })
  lf.on('edge:click', async ({ data }) => {
    try {
      const { value } = await ElMessageBox.prompt('分支条件（留空=默认流转；用自然语言描述，可引用 别名.字段）', '编辑连线条件', {
        inputValue: (data.properties && data.properties.condition) || '',
        inputPlaceholder: '如：当 create.status 为 active 时'
      })
      lf.setProperties(data.id, { condition: value })
      try { lf.updateText(data.id, value || '默认流转') } catch (e) { /* ignore */ }
      refreshEdges()
    } catch (e) { /* cancel */ }
  })
  refreshEdges()
  nextTick(() => {
    computeCanvasHeight()
    syncCanvasSize()
    // 渲染后自动适应一次，保证打开/切换场景时所有节点都在可视区域内
    if (current.value && current.value.nodes && current.value.nodes.length) fitView()
  })
}

const syncFromCanvas = (idx) => {
  if (idx === undefined) idx = activeIndex.value
  if (!lf || !scenarios.value[idx]) return
  let g
  try { g = lf.getGraphData() } catch (e) { return }
  // nodes：业务字段（命令名/说明等）以 scenarios 为准（详情面板直接编辑），画布只同步坐标
  const existing = {}
  scenarios.value[idx].nodes.forEach((n) => { existing[n.id] = n })
  scenarios.value[idx].nodes = (g.nodes || []).map((ln) => {
    const old = existing[ln.id]
    if (old) {
      old.x = ln.x
      old.y = ln.y
      return old
    }
    return {
      id: ln.id,
      type: ln.type === 'decision-node' ? 'decision' : 'command',
      commandName: (ln.properties && ln.properties.commandName) || '',
      inputNote: (ln.properties && ln.properties.inputNote) || '',
      note: (ln.properties && ln.properties.note) || '',
      alias: (ln.properties && ln.properties.alias) || '',
      x: ln.x,
      y: ln.y
    }
  })
  // edges：手动连线只在画布，从画布提取
  scenarios.value[idx].edges = (g.edges || []).map((e) => ({
    from: e.sourceNodeId,
    to: e.targetNodeId,
    condition: (e.properties && e.properties.condition) || '',
    note: (e.properties && e.properties.note) || ''
  }))
}

const syncSelectedToLf = () => {
  if (!lf || !selectedNode.value) return
  const n = selectedNode.value
  try {
    lf.setProperties(n.id, { commandName: n.commandName || '', inputNote: n.inputNote || '', note: n.note || '', alias: n.alias || '' })
    lf.updateText(n.id, buildText(n))
  } catch (e) { /* ignore */ }
}

const addScenario = () => {
  scenarios.value.push({ name: '新场景', description: '', sort: scenarios.value.length, nodes: [], edges: [] })
  activeIndex.value = scenarios.value.length - 1
  selectedNodeId.value = null
}
const removeScenario = () => {
  if (activeIndex.value == null) return
  scenarios.value.splice(activeIndex.value, 1)
  if (activeIndex.value > 0) activeIndex.value -= 1
  selectedNodeId.value = null
  nextTick(renderCanvas)
}
const addNode = (kind) => {
  if (!current.value) return
  const id = uid()
  const alias = defaultAlias(kind)
  const node = { id, type: kind, commandName: '', inputNote: '', note: '', alias, x: 140, y: 60 + current.value.nodes.length * 72 }
  current.value.nodes.push(node)
  selectedNodeId.value = id
  if (lf) {
    lf.addNode({ id, type: kind === 'decision' ? 'decision-node' : 'command-node', x: node.x, y: node.y, text: buildText(node), properties: { commandName: '', inputNote: '', note: '', alias } })
  }
}
const removeSelected = () => {
  if (!selectedNodeId.value) return
  const id = selectedNodeId.value
  if (lf) { try { lf.deleteNode(id) } catch (e) { /* ignore */ } }
  if (current.value) {
    current.value.nodes = current.value.nodes.filter((n) => n.id !== id)
    current.value.edges = current.value.edges.filter((e) => e.from !== id && e.to !== id)
  }
  selectedNodeId.value = null
}

watch(activeIndex, (n, o) => {
  if (!visible.value) return
  if (lf && scenarios.value[o] != null) syncFromCanvas(o)
  selectedNodeId.value = null
  nextTick(renderCanvas)
})

// 抽屉滑入动画结束后布局才稳定，此时再按真实位置重算一次高度
const onOpened = () => {
  computeCanvasHeight()
  syncCanvasSize()
}

// 右侧详情面板显示/隐藏会改变画布宽度，需让 LogicFlow 重新适配
watch(() => !!selectedNode.value, () => {
  nextTick(syncCanvasSize)
})

const onClosed = () => {
  if (lf) { try { lf.destroy && lf.destroy() } catch (e) { /* ignore */ }; lf = null }
  selectedNodeId.value = null
}

const onSave = async () => {
  saving.value = true
  try {
    syncFromCanvas()
    const dup = findDupAlias()
    if (dup) {
      ElMessage.error(`场景「${dup.scenario}」存在重复别名：${dup.alias}，请修改后再保存`)
      return
    }
    const c = cliData.value
    const res = await updateCli({
      id: c.ID || c.id,
      name: c.name,
      command: c.command,
      displayName: c.displayName,
      version: c.version,
      description: c.description,
      status: c.status,
      skillName: c.skillName,
      skillDescription: c.skillDescription,
      scenariosJson: JSON.stringify(scenarios.value)
    })
    if (res.code === 0) {
      ElMessage.success('场景已保存')
      visible.value = false
    } else {
      ElMessage.error(res.msg || '保存失败')
    }
  } finally {
    saving.value = false
  }
}
</script>

<style scoped>
/* 仅保留无法用 UnoCSS 原子类表达的样式：:deep() 覆盖 Element Plus 内部样式、复用的成组复杂样式 */
:deep(.el-drawer__body) { display: flex; flex-direction: column; overflow: auto; }
/* 容器已用 gap 控距，去掉 el-button 相邻默认 margin，避免间距叠加变大 */
.scenario-bar :deep(.el-button + .el-button),
.scenario-toolbar :deep(.el-button + .el-button) { margin-left: 0; }
.scenario-toolbar :deep(.el-divider--vertical) { margin: 0 4px; }
/* 右侧详情面板：选中态与空态复用同一组复杂样式 */
.scenario-detail { width: 360px; flex: none; height: 100%; overflow-y: auto; border: 1px solid #ebeef5; border-radius: 4px; padding: 12px; background: #fff; box-sizing: border-box; }
.scenario-detail.empty { color: #909399; font-size: 13px; display: flex; align-items: center; justify-content: center; text-align: center; }
.alias-dup :deep(.el-input__wrapper) { box-shadow: 0 0 0 1px #f56c6c inset; }
.alias-dup :deep(.el-input-group__prepend) { color: #f56c6c; }
</style>
