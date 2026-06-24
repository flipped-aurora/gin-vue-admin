<template>
  <el-drawer
    v-model="visible"
    :title="`调用场景编排 - ${cli.name || cli.command || ''}`"
    direction="rtl"
    size="80%"
    :append-to-body="false"
    @open="onOpen"
    @closed="onClosed"
  >
    <div class="scenario-flow">
      <div class="scenario-bar">
        <el-select v-model="activeIndex" placeholder="选择场景" style="width: 200px">
          <el-option v-for="(s, i) in scenarios" :key="i" :label="s.name || '未命名'" :value="i" />
        </el-select>
        <el-button icon="plus" @click="addScenario">新建场景</el-button>
        <el-button icon="delete" @click="removeScenario" :disabled="scenarios.length === 0">删除当前</el-button>
        <el-input v-if="current" v-model="current.name" placeholder="场景名" style="width: 180px" />
        <el-input v-if="current" v-model="current.description" placeholder="场景说明" style="width: 260px" />
      </div>

      <div class="scenario-toolbar" v-if="current">
        <el-button icon="plus" @click="addNode('command')">命令节点</el-button>
        <el-button icon="plus" @click="addNode('decision')">判断节点</el-button>
        <span class="tip">拖节点锚点手动连线；点节点编辑详情；点连线编辑条件</span>
      </div>

      <div ref="canvasRef" class="scenario-canvas"></div>

      <div class="scenario-detail" v-if="selectedNode">
        <div class="detail-title">
          <el-tag :type="selectedNode.type === 'decision' ? 'warning' : 'success'" size="small">{{ selectedNode.type }}</el-tag>
          <span>编辑节点</span>
          <el-button icon="delete" link type="danger" style="margin-left: auto" @click="removeSelected">删除节点</el-button>
        </div>
        <div class="alias-row">
          <el-input v-model="selectedNode.alias" style="width: 360px" :class="{ 'alias-dup': isAliasDup }" placeholder="如 getUser，场景内唯一；用于 别名.字段 引用" @input="syncSelectedToLf">
            <template #prepend>别名</template>
          </el-input>
          <span v-if="isAliasDup" class="alias-warn">场景内已存在</span>
        </div>
        <div class="alias-hint io-line muted" v-if="aliasHints.length">
          <span class="io-label">可用别名：</span>
          <span v-for="h in aliasHints" :key="h.alias" class="alias-chip">{{ h.alias }}({{ h.label }}): {{ h.fields }}</span>
        </div>
        <template v-if="selectedNode.type === 'command'">
          <div class="detail-row">
            <el-select v-model="selectedNode.commandName" placeholder="选择命令" style="width: 220px" @change="syncSelectedToLf">
              <el-option v-for="c in commandOptions" :key="c" :label="c" :value="c" />
            </el-select>
            <el-input v-model="selectedNode.note" placeholder="说明" style="width: 220px" @input="syncSelectedToLf" />
            <el-input v-model="selectedNode.inputNote" placeholder="入参来源(如:create.id + query.token)" style="width: 300px" @input="syncSelectedToLf" />
          </div>
          <div class="io-block" v-if="selectedNode.commandName && commandMap[selectedNode.commandName]">
            <div class="io-line"><span class="io-label">入参：</span>{{ formatParams(commandMap[selectedNode.commandName].parameters) }}</div>
            <div class="io-line"><span class="io-label">出参：</span>{{ formatFields(commandMap[selectedNode.commandName].response) }}</div>
          </div>
          <div class="io-line muted" v-else-if="selectedNode.commandName">该命令暂无参数信息</div>
        </template>
        <template v-else>
          <el-input v-model="selectedNode.note" type="textarea" :rows="2" placeholder="合并判断描述（如：基于上游 A.x、B.y、C.z 合并，判断...）" @input="syncSelectedToLf" />
          <div class="io-line muted">判断节点不执行命令；出边带条件实现分支（点画布连线编辑条件）。</div>
        </template>
      </div>
      <div class="scenario-detail empty" v-else>点击画布节点编辑详情，点空白取消选择</div>
    </div>

    <template #footer>
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" :loading="saving" @click="onSave">保存</el-button>
    </template>
  </el-drawer>
</template>

<script setup>
import { ref, computed, nextTick, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import LogicFlow, { RectNode, RectNodeModel } from '@logicflow/core'
import '@logicflow/core/dist/index.css'
import { getCliDetail, updateCli, previewManifest } from '@/plugin/auto/api/cli'

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
  lf = new LogicFlow({ container: canvasRef.value, grid: true, height: canvasRef.value.clientHeight || 360 })
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
    // 用户新建的边（无 condition 属性）默认显示提示，引导点击编辑
    if (!data.properties || data.properties.condition === undefined) {
      lf.setProperties(data.id, { condition: '', note: '' })
      try { lf.updateText(data.id, '默认流转') } catch (e) { /* ignore */ }
    }
  })
  lf.on('edge:click', async ({ data }) => {
    try {
      const { value } = await ElMessageBox.prompt('分支条件（留空=默认流转；用自然语言描述，可引用 别名.字段）', '编辑连线条件', {
        inputValue: (data.properties && data.properties.condition) || '',
        inputPlaceholder: '如：当 create.status 为 active 时'
      })
      lf.setProperties(data.id, { condition: value })
      try { lf.updateText(data.id, value || '默认流转') } catch (e) { /* ignore */ }
    } catch (e) { /* cancel */ }
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
.scenario-bar { display: flex; gap: 8px; align-items: center; margin-bottom: 12px; flex-wrap: wrap; }
.scenario-toolbar { display: flex; gap: 8px; align-items: center; margin-bottom: 8px; }
.scenario-toolbar .tip { color: #909399; font-size: 12px; }
:deep(.el-drawer__body) { display: flex; flex-direction: column; overflow: auto; }
.scenario-flow { flex: 1; display: flex; flex-direction: column; min-height: 0; }
.scenario-canvas { border: 1px solid #ebeef5; flex: 1; min-height: 0; background: #fafafa; border-radius: 4px; }
.scenario-detail { margin-top: 12px; border: 1px solid #ebeef5; border-radius: 4px; padding: 12px; background: #fff; }
.scenario-detail.empty { color: #909399; font-size: 13px; text-align: center; }
.detail-title { display: flex; gap: 8px; align-items: center; margin-bottom: 10px; font-weight: 600; }
.detail-row { display: flex; gap: 8px; align-items: center; flex-wrap: wrap; margin-bottom: 8px; }
.io-block { background: #f5f7fa; border-radius: 4px; padding: 8px 10px; }
.io-line { font-size: 13px; line-height: 1.8; }
.io-line.muted { color: #909399; }
.io-label { color: #606266; font-weight: 600; }
.alias-row { display: flex; align-items: center; gap: 8px; margin-bottom: 10px; }
.alias-dup :deep(.el-input__wrapper) { box-shadow: 0 0 0 1px #f56c6c inset; }
.alias-dup :deep(.el-input-group__prepend) { color: #f56c6c; }
.alias-warn { color: #f56c6c; font-size: 12px; }
.alias-hint { margin-bottom: 10px; }
.alias-chip { display: inline-block; margin-right: 12px; }
</style>
