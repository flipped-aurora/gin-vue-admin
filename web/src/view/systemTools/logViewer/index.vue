<template>
  <div
    ref="rootRef"
    class="p-2 flex flex-col overflow-hidden bg-main text-base-text"
    :style="{ height: rootHeight }"
  >
    <section class="gva-table-box !my-0 !p-0 h-full min-h-0 flex-1 overflow-hidden rounded-[var(--gva-radius)] border border-border shadow-card md:grid md:grid-cols-[240px_minmax(0,1fr)]">
      <aside class="hidden min-h-0 border-r border-border md:flex md:flex-col">
        <div class="border-b border-border p-3">
          <div class="mb-3 flex items-center justify-between">
            <el-tooltip content="上个月" placement="top">
              <button type="button" class="gva-tool-btn h-8 w-8" aria-label="上个月" @click="changeMonth(-1)">
                <svg-icon icon="lucide:chevron-left" class="h-4 w-4" />
              </button>
            </el-tooltip>
            <strong class="text-sm font-semibold tracking-wide">{{ monthLabel }}</strong>
            <el-tooltip content="下个月" placement="top">
              <button type="button" class="gva-tool-btn h-8 w-8" aria-label="下个月" @click="changeMonth(1)">
                <svg-icon icon="lucide:chevron-right" class="h-4 w-4" />
              </button>
            </el-tooltip>
          </div>

          <div class="grid grid-cols-7 gap-1 text-center text-xs text-muted-foreground">
            <span v-for="weekday in weekdays" :key="weekday" class="py-1 font-medium">{{ weekday }}</span>
          </div>
          <div v-loading="loadingDates" class="grid min-h-48 grid-cols-7 gap-1">
            <button
              v-for="cell in calendarCells"
              :key="cell.date"
              type="button"
              class="relative aspect-square min-h-7 rounded-md border border-transparent text-xs transition-colors"
              :class="calendarCellClass(cell)"
              :disabled="!cell.inMonth || !cell.available"
              :aria-label="cell.available ? `查看 ${cell.date} 日志` : `${cell.date} 无日志`"
              @click="selectDate(cell.date)"
            >
              {{ cell.day }}
              <i
                v-if="cell.inMonth && cell.available && cell.date !== selectedDate"
                class="absolute bottom-1 left-1/2 h-1 w-1 -translate-x-1/2 rounded-full bg-success"
              />
            </button>
          </div>

          <div class="mt-3 flex items-center gap-4 text-xs text-muted-foreground">
            <span class="flex items-center gap-1.5"><i class="h-1.5 w-1.5 rounded-full bg-success" />有日志</span>
            <span class="flex items-center gap-1.5"><i class="h-2.5 w-2.5 rounded-[4px] bg-primary" />已选择</span>
          </div>
        </div>

        <div class="min-h-0 flex-1 p-3">
          <div class="mb-2 flex items-center justify-between px-1 text-xs text-muted-foreground">
            <span class="font-medium">{{ selectedDate || '未选择日期' }}</span>
            <span class="rounded-full bg-muted px-2 py-0.5">{{ files.length }} 个文件</span>
          </div>
          <div v-loading="loadingFiles" class="h-full overflow-y-auto">
            <button
              v-for="file in files"
              :key="file.path"
              type="button"
              class="mb-1 flex w-full items-center gap-2 rounded-md px-2.5 py-2 text-left text-xs transition-colors hover:bg-black/5 dark:hover:bg-white/5"
              :class="file.path === activePath ? 'bg-primary-50 font-medium text-primary dark:bg-primary/15' : 'text-base-text'"
              @click="openFile(file.path)"
            >
              <svg-icon icon="lucide:file-text" class="h-4 w-4 shrink-0" :class="file.path === activePath ? 'text-primary' : 'text-muted-foreground'" />
              <span class="min-w-0 flex-1 truncate" :title="file.path">{{ file.path }}</span>
              <span class="shrink-0 tabular-nums text-muted-foreground">{{ formatFileSize(file.size) }}</span>
            </button>
            <div v-if="selectedDate && !loadingFiles && files.length === 0" class="flex h-32 flex-col items-center justify-center gap-2 text-xs text-muted-foreground">
              <svg-icon icon="lucide:folder-open" class="h-6 w-6 opacity-50" />
              该日期暂无日志文件
            </div>
          </div>
        </div>
      </aside>

      <main class="flex h-full min-h-0 min-w-0 flex-col">
        <div class="border-b border-border p-3 md:hidden">
          <div class="mb-3 flex items-center justify-between">
            <button type="button" class="gva-tool-btn h-8 w-8" aria-label="上个月" @click="changeMonth(-1)">
              <svg-icon icon="lucide:chevron-left" class="h-4 w-4" />
            </button>
            <strong class="text-sm">{{ monthLabel }}</strong>
            <button type="button" class="gva-tool-btn h-8 w-8" aria-label="下个月" @click="changeMonth(1)">
              <svg-icon icon="lucide:chevron-right" class="h-4 w-4" />
            </button>
          </div>
          <div class="grid grid-cols-1 gap-2 sm:grid-cols-2">
            <el-select
              :model-value="selectedDate"
              placeholder="选择有日志的日期"
              :loading="loadingDates"
              @change="selectDate"
            >
              <el-option v-for="item in availableDates" :key="item.date" :label="`${item.date} (${item.fileCount})`" :value="item.date" />
            </el-select>
            <el-select
              :model-value="activePath"
              placeholder="选择日志文件"
              :loading="loadingFiles"
              :disabled="files.length === 0"
              @change="openFile"
            >
              <el-option v-for="file in files" :key="file.path" :label="file.path" :value="file.path" />
            </el-select>
          </div>
        </div>

        <el-tabs
          :model-value="activePath"
          type="card"
          closable
          class="log-tabs min-w-0 border-b border-border px-3 pb-2 pt-2"
          @tab-change="openFile"
          @tab-remove="closeTab"
        >
          <el-tab-pane
            v-for="path in openedPaths"
            :key="path"
            :name="path"
          >
            <template #label>
              <span class="block max-w-52 truncate" :title="path">{{ path }}</span>
            </template>
          </el-tab-pane>
        </el-tabs>

        <div class="flex flex-wrap items-center justify-between gap-2 border-b border-border px-3 py-2">
          <div class="flex min-w-0 flex-1 items-center gap-3">
            <el-input v-model="searchText" clearable placeholder="搜索已加载内容" class="max-w-72" :disabled="!content">
              <template #prefix>
                <svg-icon icon="lucide:search" class="h-4 w-4" />
              </template>
            </el-input>
            <label class="flex shrink-0 items-center gap-2 text-xs text-muted-foreground">
              <el-switch v-model="wrapLines" aria-label="自动换行" />
              自动换行
            </label>
          </div>
          <div class="flex items-center gap-1">
            <span class="mr-2 hidden rounded-full bg-muted px-2.5 py-1 text-[11px] text-muted-foreground sm:inline">{{ contentMeta }}</span>
            <el-tooltip content="复制已加载日志" placement="top">
              <button type="button" class="gva-tool-btn h-8 w-8" aria-label="复制已加载日志" :disabled="!content" @click="copyContent">
                <svg-icon icon="lucide:copy" class="h-4 w-4" />
              </button>
            </el-tooltip>
            <el-tooltip content="刷新当前日志" placement="top">
              <button
                type="button"
                class="gva-tool-btn h-8 w-8"
                aria-label="刷新当前日志"
                :disabled="!activePath || contentBusy"
                @click="refreshContent"
              >
                <svg-icon icon="lucide:refresh-cw" class="h-4 w-4" :class="{ 'animate-spin': contentBusy }" />
              </button>
            </el-tooltip>
          </div>
        </div>

        <div v-if="activeFile" class="flex flex-wrap items-center gap-x-4 gap-y-1 border-b border-border px-3 py-2 text-xs text-muted-foreground">
          <span class="truncate font-medium text-base-text" :title="activeFile.path">{{ activeFile.path }}</span>
          <span class="tabular-nums">{{ formatFileSize(activeFile.size) }}</span>
          <span>{{ formatDate(activeFile.modifiedAt) }}</span>
          <span v-if="limitedByBytes" class="text-warning">本段已达到 2 MiB 上限</span>
        </div>

        <div class="relative min-h-0 flex-1 bg-base-text text-container">
          <div v-if="hasMore" class="absolute inset-x-0 top-0 z-10 flex justify-center py-2">
            <button
              type="button"
              class="inline-flex h-8 items-center gap-2 rounded-full bg-primary px-4 text-xs text-white shadow-lg transition-colors hover:bg-primary-600 disabled:cursor-not-allowed disabled:opacity-50"
              :disabled="contentBusy"
              @click="loadOlder"
            >
              <svg-icon icon="lucide:history" class="h-4 w-4" :class="{ 'animate-spin': loadingOlder }" />
              加载更早日志
            </button>
          </div>

          <div
            ref="contentScrollRef"
            id="log-viewer-content"
            role="tabpanel"
            v-loading="loadingContent"
            class="h-full overflow-auto font-mono text-xs leading-6"
            :class="hasMore ? 'pt-12' : 'pt-2'"
          >
            <div v-if="displayLines.length" :class="wrapLines ? 'min-w-0' : 'min-w-max'">
              <div
                v-for="line in displayLines"
                :key="line.sourceIndex"
                :data-log-line-index="line.sourceIndex"
                class="grid grid-cols-[3.5rem_minmax(0,1fr)] px-3 transition-colors hover:bg-container/5"
              >
                <span class="select-none border-r border-container/10 pr-3 text-right text-muted-foreground">{{ line.number }}</span>
                <span class="pl-3" :class="[lineTone(line.text), wrapLines ? 'whitespace-pre-wrap break-all' : 'whitespace-pre']">
                  <template v-for="(segment, index) in highlightedSegments(line.text)" :key="`${line.number}-${index}`">
                    <mark v-if="segment.match" class="rounded-sm bg-warning-300 px-0.5 text-black">{{ segment.text }}</mark>
                    <template v-else>{{ segment.text }}</template>
                  </template>
                </span>
              </div>
            </div>

            <div v-else-if="!loadingContent" class="flex h-full flex-col items-center justify-center gap-3 px-6 text-center text-sm text-muted-foreground">
              <svg-icon icon="lucide:terminal" class="h-8 w-8 opacity-40" />
              <span>{{ emptyContentText }}</span>
            </div>
          </div>
        </div>
      </main>
    </section>
  </div>
</template>

<script setup>
import {
  computed,
  nextTick,
  onActivated,
  onBeforeUnmount,
  onMounted,
  ref
} from 'vue'
import { ElMessage } from 'element-plus'
import { getLogContent, getLogDates, getLogFiles } from '@/api/logViewer'
import { formatDate } from '@/utils/format'

defineOptions({ name: 'LogViewer' })

const weekdays = ['一', '二', '三', '四', '五', '六', '日']
const MAX_LOADED_LINES = 5000
const MAX_LOADED_CHARS = 10 * 1024 * 1024
const now = new Date()
const todayText = formatLocalDate(now)
const visibleMonth = ref(formatMonth(now))
const availableDates = ref([])
const selectedDate = ref('')
const files = ref([])
const openedPaths = ref([])
const activePath = ref('')
const content = ref('')
const nextCursor = ref(0)
const hasMore = ref(false)
const limitedByBytes = ref(false)
const windowLimited = ref(false)
const searchText = ref('')
const wrapLines = ref(false)
const loadingDates = ref(false)
const loadingFiles = ref(false)
const loadingContent = ref(false)
const loadingOlder = ref(false)
const rootRef = ref(null)
const rootHeight = ref('auto')
const contentScrollRef = ref(null)
const contentBusy = computed(() => loadingContent.value || loadingOlder.value)

let dateRequestId = 0
let fileRequestId = 0
let contentRequestId = 0

const availableDateSet = computed(() => new Set(availableDates.value.map(item => item.date)))
const monthLabel = computed(() => {
  const [year, month] = visibleMonth.value.split('-')
  return `${year} 年 ${Number(month)} 月`
})
const activeFile = computed(() => files.value.find(file => file.path === activePath.value) || null)
const contentMeta = computed(() => {
  if (!activePath.value) return ''
  const suffix = hasMore.value ? ' · 仍有更早内容' : ''
  const windowSuffix = windowLimited.value ? ' · 已限制显示窗口' : ''
  return `已加载 ${allLines.value.length} 行${suffix}${windowSuffix}`
})
const calendarCells = computed(() => {
  const [year, month] = visibleMonth.value.split('-').map(Number)
  const firstDay = new Date(year, month - 1, 1)
  const mondayOffset = (firstDay.getDay() + 6) % 7
  return Array.from({ length: 42 }, (_, index) => {
    const date = new Date(year, month - 1, 1 - mondayOffset + index)
    const dateText = formatLocalDate(date)
    return {
      date: dateText,
      day: date.getDate(),
      inMonth: date.getMonth() === month - 1,
      available: availableDateSet.value.has(dateText)
    }
  })
})
const allLines = computed(() => {
  if (!content.value) return []
  const lines = content.value.split('\n')
  if (lines[lines.length - 1] === '') lines.pop()
  return lines.map((text, index) => ({ number: index + 1, sourceIndex: index, text }))
})
const displayLines = computed(() => {
  const keyword = searchText.value.trim().toLowerCase()
  if (!keyword) return allLines.value
  return allLines.value.filter(line => line.text.toLowerCase().includes(keyword))
})
const emptyContentText = computed(() => {
  if (!activePath.value) return '请先选择日志日期和文件'
  if (searchText.value && allLines.value.length) return '当前已加载内容中没有匹配结果'
  return '日志文件为空'
})

function formatMonth(date) {
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}`
}

function formatLocalDate(date) {
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
}

function formatFileSize(size) {
  if (!Number.isFinite(size) || size <= 0) return '0 B'
  const units = ['B', 'KB', 'MB', 'GB']
  const unitIndex = Math.min(Math.floor(Math.log(size) / Math.log(1024)), units.length - 1)
  const value = size / (1024 ** unitIndex)
  return `${value >= 10 || unitIndex === 0 ? value.toFixed(0) : value.toFixed(1)} ${units[unitIndex]}`
}

function calendarCellClass(cell) {
  if (!cell.inMonth) return 'cursor-default text-muted-foreground opacity-30'
  if (cell.date === selectedDate.value) return 'border-primary bg-primary text-white font-semibold shadow-sm'
  if (cell.available) {
    return cell.date === todayText
      ? 'border-primary/50 text-success font-semibold hover:bg-success-50'
      : 'text-success font-medium hover:bg-success-50 hover:border-success/40'
  }
  if (cell.date === todayText) return 'border-primary/50 font-semibold text-base-text'
  return 'cursor-not-allowed text-muted-foreground opacity-50'
}

function lineTone(text) {
  if (/\b(ERROR|DPANIC|PANIC|FATAL)\b/i.test(text)) return 'text-error'
  if (/\bWARN\b/i.test(text)) return 'text-warning'
  if (/\bDEBUG\b/i.test(text)) return 'text-info'
  return 'text-container'
}

function highlightedSegments(text) {
  const keyword = searchText.value.trim()
  if (!keyword) return [{ text, match: false }]
  const lowerText = text.toLowerCase()
  const lowerKeyword = keyword.toLowerCase()
  const segments = []
  let cursor = 0
  let matchIndex = lowerText.indexOf(lowerKeyword)
  while (matchIndex !== -1) {
    if (matchIndex > cursor) segments.push({ text: text.slice(cursor, matchIndex), match: false })
    segments.push({ text: text.slice(matchIndex, matchIndex + keyword.length), match: true })
    cursor = matchIndex + keyword.length
    matchIndex = lowerText.indexOf(lowerKeyword, cursor)
  }
  if (cursor < text.length) segments.push({ text: text.slice(cursor), match: false })
  return segments.length ? segments : [{ text, match: false }]
}

function resetContentState() {
  contentRequestId++
  loadingContent.value = false
  loadingOlder.value = false
  activePath.value = ''
  content.value = ''
  nextCursor.value = 0
  hasMore.value = false
  limitedByBytes.value = false
  windowLimited.value = false
  searchText.value = ''
}

async function loadDates() {
  const requestId = ++dateRequestId
  loadingDates.value = true
  try {
    const result = await getLogDates({ month: visibleMonth.value })
    if (requestId !== dateRequestId || result.code !== 0) return
    availableDates.value = result.data.dates || []
    const latestDate = availableDates.value.at(-1)?.date || ''
    if (latestDate) {
      await selectDate(latestDate)
    } else {
      fileRequestId++
      selectedDate.value = ''
      files.value = []
      openedPaths.value = []
      resetContentState()
    }
  } finally {
    if (requestId === dateRequestId) loadingDates.value = false
  }
}

async function changeMonth(offset) {
  const [year, month] = visibleMonth.value.split('-').map(Number)
  fileRequestId++
  visibleMonth.value = formatMonth(new Date(year, month - 1 + offset, 1))
  selectedDate.value = ''
  files.value = []
  openedPaths.value = []
  resetContentState()
  await loadDates()
}

async function selectDate(date) {
  if (!date || !availableDateSet.value.has(date)) return
  const requestId = ++fileRequestId
  selectedDate.value = date
  files.value = []
  openedPaths.value = []
  resetContentState()
  loadingFiles.value = true
  try {
    const result = await getLogFiles({ date })
    if (requestId !== fileRequestId || selectedDate.value !== date || result.code !== 0) return
    files.value = result.data.files || []
    const defaultFile = files.value.find(file => file.path === 'info.log') || files.value[0]
    if (defaultFile) await openFile(defaultFile.path)
  } finally {
    if (requestId === fileRequestId) loadingFiles.value = false
  }
}

async function openFile(path) {
  if (!path || !files.value.some(file => file.path === path)) return
  if (!openedPaths.value.includes(path)) openedPaths.value.push(path)
  activePath.value = path
  searchText.value = ''
  await fetchLatestContent(path)
}

async function fetchLatestContent(path) {
  const requestId = ++contentRequestId
  const date = selectedDate.value
  loadingOlder.value = false
  loadingContent.value = true
  try {
    const result = await getLogContent({ date, path })
    if (requestId !== contentRequestId || activePath.value !== path || selectedDate.value !== date) return
    if (result.code !== 0) {
      await reloadFilesAfterContentFailure(path)
      return
    }
    applyContentResult(result.data)
    await nextTick()
    if (contentScrollRef.value) contentScrollRef.value.scrollTop = contentScrollRef.value.scrollHeight
  } finally {
    if (requestId === contentRequestId) loadingContent.value = false
  }
}

function applyContentResult(data) {
  const bounded = boundLoadedContent(data.content || '')
  content.value = bounded.content
  windowLimited.value = bounded.limited
  nextCursor.value = data.nextCursor || 0
  hasMore.value = Boolean(data.hasMore)
  limitedByBytes.value = Boolean(data.limitedByBytes)
  const file = files.value.find(item => item.path === activePath.value)
  if (file) {
    file.size = data.size
    file.modifiedAt = data.modifiedAt
  }
}

async function loadOlder() {
  if (!activePath.value || !hasMore.value || contentBusy.value) return
  const requestId = ++contentRequestId
  const date = selectedDate.value
  const path = activePath.value
  const scrollElement = contentScrollRef.value
  const previousHeight = scrollElement?.scrollHeight || 0
  const anchor = captureScrollAnchor(scrollElement)
  loadingOlder.value = true
  try {
    const result = await getLogContent({ date, path, cursor: nextCursor.value })
    if (requestId !== contentRequestId || activePath.value !== path || selectedDate.value !== date) return
    if (result.code !== 0) {
      await reloadFilesAfterContentFailure(path)
      return
    }
    const olderContent = result.data.content || ''
    const addedLineBreaks = countLineBreaks(olderContent)
    const bounded = boundLoadedContent(olderContent + content.value)
    content.value = bounded.content
    windowLimited.value = windowLimited.value || bounded.limited
    nextCursor.value = result.data.nextCursor || 0
    hasMore.value = Boolean(result.data.hasMore)
    limitedByBytes.value = Boolean(result.data.limitedByBytes)
    await nextTick()
    restoreScrollAnchor(scrollElement, anchor, addedLineBreaks, previousHeight)
  } finally {
    if (requestId === contentRequestId) loadingOlder.value = false
  }
}

async function refreshContent() {
  if (activePath.value && !contentBusy.value) await fetchLatestContent(activePath.value)
}

async function reloadFilesAfterContentFailure(failedPath) {
  const date = selectedDate.value
  if (!date) return
  const requestId = ++fileRequestId
  loadingFiles.value = true
  resetContentState()
  try {
    const result = await getLogFiles({ date })
    if (requestId !== fileRequestId || selectedDate.value !== date || result.code !== 0) return
    files.value = result.data.files || []
    const validPaths = new Set(files.value.map(file => file.path))
    if (validPaths.has(failedPath)) {
      openedPaths.value = openedPaths.value.filter(path => validPaths.has(path))
      return
    }
    openedPaths.value = openedPaths.value.filter(path => path !== failedPath && validPaths.has(path))
    const nextPath = openedPaths.value[0] || files.value[0]?.path || ''
    if (nextPath) await openFile(nextPath)
  } finally {
    if (requestId === fileRequestId) loadingFiles.value = false
  }
}

function boundLoadedContent(value) {
  let bounded = value
  let limited = false
  if (bounded.length > MAX_LOADED_CHARS) {
    bounded = bounded.slice(0, MAX_LOADED_CHARS)
    limited = true
  }

  let lineBreaks = 0
  for (let index = 0; index < bounded.length; index++) {
    if (bounded[index] !== '\n') continue
    lineBreaks++
    if (lineBreaks === MAX_LOADED_LINES && index + 1 < bounded.length) {
      bounded = bounded.slice(0, index + 1)
      limited = true
      break
    }
  }
  return { content: bounded, limited }
}

function countLineBreaks(value) {
  let count = 0
  for (let index = 0; index < value.length; index++) {
    if (value[index] === '\n') count++
  }
  return count
}

function captureScrollAnchor(scrollElement) {
  if (!scrollElement) return null
  const viewportTop = scrollElement.getBoundingClientRect().top
  const lineElements = scrollElement.querySelectorAll('[data-log-line-index]')
  const anchor = Array.from(lineElements).find(element => element.getBoundingClientRect().bottom >= viewportTop)
  if (!anchor) return null
  return {
    sourceIndex: Number(anchor.dataset.logLineIndex),
    top: anchor.getBoundingClientRect().top
  }
}

function restoreScrollAnchor(scrollElement, anchor, addedLineBreaks, previousHeight) {
  if (!scrollElement) return
  if (anchor) {
    const targetIndex = anchor.sourceIndex + addedLineBreaks
    const target = scrollElement.querySelector(`[data-log-line-index="${targetIndex}"]`)
    if (target) {
      scrollElement.scrollTop += target.getBoundingClientRect().top - anchor.top
      return
    }
  }
  scrollElement.scrollTop = Math.max(0, scrollElement.scrollTop + scrollElement.scrollHeight - previousHeight)
}

function closeTab(path) {
  const index = openedPaths.value.indexOf(path)
  if (index === -1) return
  openedPaths.value.splice(index, 1)
  if (path !== activePath.value) return
  const nextPath = openedPaths.value[index] || openedPaths.value[index - 1] || ''
  resetContentState()
  if (nextPath) openFile(nextPath)
}

async function copyContent() {
  if (!content.value) return
  try {
    if (navigator.clipboard?.writeText) {
      await navigator.clipboard.writeText(content.value)
    } else {
      const textarea = document.createElement('textarea')
      textarea.value = content.value
      textarea.setAttribute('readonly', '')
      textarea.className = 'fixed left-[-9999px] top-0'
      document.body.appendChild(textarea)
      textarea.select()
      if (!document.execCommand('copy')) throw new Error('copy failed')
      document.body.removeChild(textarea)
    }
    ElMessage.success('日志内容已复制')
  } catch {
    ElMessage.error('复制失败')
  }
}

const updateRootHeight = () => {
  const element = rootRef.value
  if (!element) return
  const top = element.getBoundingClientRect().top
  rootHeight.value = `${Math.max(320, window.innerHeight - top - 32)}px`
}

onMounted(() => {
  nextTick(updateRootHeight)
  window.addEventListener('resize', updateRootHeight)
  loadDates()
})
onActivated(updateRootHeight)
onBeforeUnmount(() => window.removeEventListener('resize', updateRootHeight))
</script>

<style scoped>
/* el-tabs 卡片样式重写为轻量 chip：去掉默认硬边框与底线，选中态主色浅底 */
.log-tabs :deep(.el-tabs--card > .el-tabs__header) {
  margin-bottom: 0;
  border-bottom: none;
}
.log-tabs :deep(.el-tabs--card > .el-tabs__header .el-tabs__nav-wrap::after) {
  content: none;
}
.log-tabs :deep(.el-tabs--card > .el-tabs__header .el-tabs__nav) {
  border: none;
  border-radius: 0;
}
.log-tabs :deep(.el-tabs--card > .el-tabs__header .el-tabs__item) {
  height: 32px;
  margin-right: 4px;
  padding: 0 12px;
  border: none;
  border-radius: 6px;
  line-height: 32px;
  font-size: 12px;
  color: rgb(var(--muted-foreground-color));
  transition: color 0.2s ease, background-color 0.2s ease;
}
.log-tabs :deep(.el-tabs--card > .el-tabs__header .el-tabs__item:hover) {
  color: rgb(var(--primary-color));
}
.log-tabs :deep(.el-tabs--card > .el-tabs__header .el-tabs__item.is-active) {
  background: rgb(var(--primary-color) / 0.1);
  color: rgb(var(--primary-color));
  font-weight: 500;
}
</style>
