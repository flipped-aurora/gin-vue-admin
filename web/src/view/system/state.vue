<template>
  <div class="space-y-2 pt-2">
    <!-- 首载骨架 -->
    <div v-if="!loaded" class="gva-card bg-container p-4">
      <el-skeleton :rows="6" animated />
    </div>

    <template v-else>
      <!-- KPI 行 -->
      <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-3">
        <div v-for="tile in kpiTiles" :key="tile.key" class="gva-card bg-container px-4 py-3">
          <div class="text-xs text-muted-foreground tracking-wide">{{ tile.label }}</div>
          <div class="mt-0.5 text-2xl font-semibold leading-snug" :class="tile.valueClass">
            {{ tile.value }}<span v-if="tile.suffix" class="text-xs font-normal text-muted-foreground ml-1">{{
              tile.suffix }}</span>
          </div>
          <!-- 趋势微线: 本地累积快照, 不足 2 个时留白占位(高度不塌) -->
          <svg v-if="tile.points" class="block mt-1.5 w-full" :class="tile.lineClass" height="22" viewBox="0 0 120 22"
            preserveAspectRatio="none" aria-hidden="true">
            <polyline :points="tile.points" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
              stroke-linejoin="round" vector-effect="non-scaling-stroke" />
          </svg>
          <div v-else class="mt-1.5 h-[22px]"></div>
        </div>
      </div>

      <!-- 双栏面板 -->
      <div class="grid grid-cols-1 lg:grid-cols-5 gap-3">
        <!-- CPU 每核负载 -->
        <div class="lg:col-span-3 gva-card bg-container px-4 py-3">
          <div class="flex items-baseline justify-between mb-2.5">
            <span class="text-sm font-semibold tracking-tight text-base-text">CPU 每核负载</span>
            <span class="text-xs text-muted-foreground">{{ coreSummary }}</span>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-6">
            <div v-for="core in cores" :key="core.index" class="flex items-center gap-2 mb-1.5">
              <span class="text-xs text-muted-foreground w-11 shrink-0">core {{ core.index }}</span>
              <div class="flex-1 h-2 rounded bg-control-track overflow-hidden">
                <div class="h-full rounded" :class="core.barClass" :style="{ width: core.pct + '%' }"></div>
              </div>
              <span class="text-xs text-base-text w-9 text-right tabular-nums">{{ core.pct }}%</span>
            </div>
          </div>
        </div>

        <!-- 磁盘 -->
        <div class="lg:col-span-2 gva-card bg-container px-4 py-3">
          <div class="flex items-baseline justify-between mb-2.5">
            <span class="text-sm font-semibold tracking-tight text-base-text">磁盘</span>
            <span class="text-xs text-muted-foreground">{{ disks.length }} 个挂载点</span>
          </div>
          <template v-if="disks.length">
            <div v-for="d in disks" :key="d.mountPoint" class="mb-3 last:mb-0">
              <div class="flex justify-between text-xs mb-1">
                <span class="font-semibold text-base-text">{{ d.mountPoint }}</span>
                <span class="text-muted-foreground tabular-nums">{{ d.usedGb }} / {{ d.totalGb }} GB · {{ d.usedPercent
                  }}%</span>
              </div>
              <div class="h-2 rounded bg-control-track overflow-hidden">
                <div class="h-full rounded" :class="d.barClass" :style="{ width: d.usedPercent + '%' }"></div>
              </div>
            </div>
          </template>
          <div v-else class="text-xs text-muted-foreground py-4 text-center">
            未配置挂载点(见 config.yaml 的 disk-list)
          </div>
        </div>
      </div>

      <!-- 运行时信息 chips -->
      <div class="flex flex-wrap gap-2">
        <span v-for="chip in runtimeChips" :key="chip.label"
          class="inline-flex items-center gap-1.5 bg-muted border border-border rounded-md px-2.5 py-1 text-xs text-muted-foreground">
          {{ chip.label }} <b class="font-semibold text-base-text">{{ chip.value }}</b>
        </span>
      </div>
    </template>

    <div class="flex justify-end">
      <span class="text-xs text-muted-foreground">10 秒自动刷新 · 最近更新 {{ lastUpdated || '—' }}</span>
    </div>
  </div>
</template>

<script setup>
import { getSystemState } from '@/api/system'
import { computed, onActivated, onDeactivated, onUnmounted, ref } from 'vue'

defineOptions({
  name: 'State'
})

// 10s 轮询 × 30 个快照 ≈ 5 分钟趋势窗口; 仅前端内存态, 刷新即清空(设计文档已确认)
const HISTORY_MAX = 30

const state = ref(null)
const loaded = ref(false)
const lastUpdated = ref('')
const history = ref({ cpu: [], ram: [], disk: [], goroutine: [] })

// 阈值状态: <70 主题色 / ≥70 偏高 / ≥90 告警。
// UnoCSS 只识别字面量类名, 必须用完整映射, 禁止 `bg-${status}` 动态拼接。
const STATUS_TEXT = { primary: 'text-primary', warning: 'text-warning', error: 'text-error' }
const STATUS_BG = { primary: 'bg-primary', warning: 'bg-warning', error: 'bg-error' }
const statusOf = (pct) => (pct >= 90 ? 'error' : pct >= 70 ? 'warning' : 'primary')

const pushHistory = (list, value) => {
  list.push(value)
  if (list.length > HISTORY_MAX) list.shift()
}

const reload = async () => {
  try {
    const res = await getSystemState()
    const server = res?.data?.server
    if (!server) return // 失败静默保留上次数据, 请求层已有统一错误提示
    state.value = server

    const h = history.value
    pushHistory(h.cpu, cpuAvg.value)
    pushHistory(h.ram, server.ram?.usedPercent ?? 0)
    pushHistory(h.disk, topDisk.value ? topDisk.value.usedPercent : 0)
    pushHistory(h.goroutine, server.os?.numGoroutine ?? 0)

    lastUpdated.value = new Date().toLocaleTimeString('zh-CN', { hour12: false })
    loaded.value = true
  } catch {
    // 请求层已统一提示, 静默保留上次数据
  }
}

// 序列 -> polyline points。百分比序列用固定 0-100 标尺; 计数序列(goroutine)min-max 自适应
const toPoints = (series, { fixedScale = true } = {}) => {
  if (series.length < 2) return ''
  const W = 120
  const H = 22
  const PAD = 2
  let min = 0
  let max = 100
  if (!fixedScale) {
    min = Math.min(...series)
    max = Math.max(...series)
    if (min === max) {
      min -= 1
      max += 1 // 平线时置中显示
    }
  }
  const span = max - min
  return series
    .map((v, i) => {
      const x = (i / (series.length - 1)) * W
      const y = H - PAD - ((v - min) / span) * (H - PAD * 2)
      return `${x.toFixed(1)},${y.toFixed(1)}`
    })
    .join(' ')
}

const gb = (mb) => (mb / 1024).toFixed(1)

const cpuAvg = computed(() => {
  const cpus = state.value?.cpu?.cpus || []
  return cpus.length ? Math.round(cpus.reduce((sum, v) => sum + v, 0) / cpus.length) : 0
})

const topDisk = computed(() => {
  const list = state.value?.disk || []
  return list.reduce((max, d) => (max && max.usedPercent >= d.usedPercent ? max : d), null)
})

const kpiTiles = computed(() => {
  const s = state.value
  const h = history.value
  const ramPct = s?.ram?.usedPercent ?? 0
  const diskStatus = topDisk.value ? statusOf(topDisk.value.usedPercent) : 'primary'
  return [
    {
      key: 'cpu',
      label: 'CPU 平均',
      value: cpuAvg.value,
      suffix: '%',
      valueClass: 'text-base-text',
      lineClass: STATUS_TEXT[statusOf(cpuAvg.value)],
      points: toPoints(h.cpu)
    },
    {
      key: 'ram',
      label: '内存',
      value: ramPct,
      suffix: `% · ${gb(s?.ram?.usedMb ?? 0)} / ${gb(s?.ram?.totalMb ?? 0)} GB`,
      valueClass: 'text-base-text',
      lineClass: STATUS_TEXT[statusOf(ramPct)],
      points: toPoints(h.ram)
    },
    {
      key: 'disk',
      label: topDisk.value ? `磁盘 · ${topDisk.value.mountPoint} 最高` : '磁盘',
      value: topDisk.value ? topDisk.value.usedPercent : '—',
      suffix: topDisk.value ? '%' : '',
      // 设计稿确认: 磁盘 KPI 大数字是唯一取状态色的大数字, 其余保持正文色
      valueClass: topDisk.value ? STATUS_TEXT[diskStatus] : 'text-base-text',
      lineClass: STATUS_TEXT[diskStatus],
      points: topDisk.value ? toPoints(h.disk) : ''
    },
    {
      key: 'goroutine',
      label: 'Goroutines',
      value: s?.os?.numGoroutine ?? 0,
      suffix: '',
      valueClass: 'text-base-text',
      lineClass: 'text-primary',
      points: toPoints(h.goroutine, { fixedScale: false })
    }
  ]
})

const cores = computed(() =>
  (state.value?.cpu?.cpus || []).map((v, index) => {
    const pct = Math.round(v)
    return { index, pct, barClass: STATUS_BG[statusOf(pct)] }
  })
)

const coreSummary = computed(() => {
  const logical = state.value?.cpu?.cpus?.length ?? 0
  const physical = state.value?.cpu?.cores ?? 0
  return `${logical} 逻辑核 · ${physical} 物理核`
})

const disks = computed(() =>
  (state.value?.disk || []).map((d) => ({ ...d, barClass: STATUS_BG[statusOf(d.usedPercent)] }))
)

const runtimeChips = computed(() => {
  const os = state.value?.os || {}
  return [
    { label: 'OS', value: os.goos || '—' },
    { label: 'Go', value: os.goVersion || '—' },
    { label: '编译器', value: os.compiler || '—' },
    { label: '逻辑核', value: os.numCpu ?? '—' }
  ]
})

// keep-alive 缓存页切走时走 deactivated 而非 unmounted, 需停止轮询避免后台空转
let timer = null
const stopPolling = () => {
  if (timer) {
    clearInterval(timer)
    timer = null
  }
}
const startPolling = () => {
  stopPolling()
  timer = setInterval(reload, 10 * 1000)
}

reload()
startPolling()
onActivated(() => {
  reload()
  startPolling()
})
onDeactivated(stopPolling)
onUnmounted(stopPolling)
</script>
