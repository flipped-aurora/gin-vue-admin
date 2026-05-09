<template>
  <div class="wubi-practice p-4">
    <el-card class="header-card mb-4" shadow="never">
      <div class="flex items-center justify-between flex-wrap gap-2">
        <div class="text-lg font-semibold">五笔打字练习</div>
        <div class="flex items-center gap-3 flex-wrap">
          <el-select v-model="questionCount" style="width: 140px" @change="reset">
            <el-option label="10 题" :value="10" />
            <el-option label="20 题" :value="20" />
            <el-option label="40 题" :value="40" />
          </el-select>
          <el-switch v-model="showCode" active-text="显示编码" inline-prompt />
          <el-switch v-model="showPinyin" active-text="显示拼音" inline-prompt />
          <el-button type="primary" @click="reset">重新开始</el-button>
        </div>
      </div>
    </el-card>

    <el-card v-if="!finished" class="play-card mb-4" shadow="never">
      <div class="flex flex-col items-center gap-4 py-6">
        <div class="text-sm text-gray-500">
          第 {{ index + 1 }} / {{ items.length }} 题
        </div>

        <div class="text-9xl font-bold text-primary tracking-wide">
          {{ currentItem.ch }}
        </div>

        <div v-if="showCode" class="text-2xl text-gray-700 tracking-widest">
          编码: <span class="font-mono text-blue-600">{{ currentItem.code }}</span>
        </div>
        <div v-if="showPinyin" class="text-base text-gray-500">
          拼音: {{ currentItem.pinyin }}
        </div>

        <el-input
          ref="inputRef"
          v-model="inputVal"
          size="large"
          style="width: 320px"
          placeholder="请用五笔输入法键入上面的字"
          @input="onInput"
          @compositionend="onCompositionEnd"
        />

        <el-button v-if="!started" type="success" size="large" @click="start">开始</el-button>
        <el-button v-else size="small" link @click="skip">跳过 ▶</el-button>
      </div>
    </el-card>

    <el-card v-if="started || finished" class="stats-card mb-4" shadow="never">
      <el-row :gutter="16">
        <el-col :span="6">
          <el-statistic title="用时(秒)" :value="elapsed" />
        </el-col>
        <el-col :span="6">
          <el-statistic title="正确" :value="correctCount" />
        </el-col>
        <el-col :span="6">
          <el-statistic title="错误" :value="errorCount" />
        </el-col>
        <el-col :span="6">
          <el-statistic title="字/分钟" :value="wpm" :precision="1" />
        </el-col>
      </el-row>
    </el-card>

    <el-card v-if="finished" class="result-card" shadow="never">
      <template #header>
        <div class="text-base font-semibold">本轮成绩</div>
      </template>
      <el-descriptions :column="2" border>
        <el-descriptions-item label="题目数">{{ items.length }}</el-descriptions-item>
        <el-descriptions-item label="准确率">{{ accuracy.toFixed(1) }}%</el-descriptions-item>
        <el-descriptions-item label="正确">{{ correctCount }}</el-descriptions-item>
        <el-descriptions-item label="错误">{{ errorCount }}</el-descriptions-item>
        <el-descriptions-item label="用时">{{ elapsed }} 秒</el-descriptions-item>
        <el-descriptions-item label="字/分钟">{{ wpm.toFixed(1) }}</el-descriptions-item>
      </el-descriptions>
      <div class="mt-4 flex gap-2">
        <el-button type="primary" :loading="submitting" @click="submit">提交成绩</el-button>
        <el-button @click="reset">再来一轮</el-button>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed, onUnmounted, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import { pickRandom } from '@/plugin/wubi/data/wubi-dict'
import { submitScore } from '@/plugin/wubi/api/score'

const questionCount = ref(20)
const showCode = ref(true)
const showPinyin = ref(false)
const items = ref(pickRandom(questionCount.value))
const index = ref(0)
const inputVal = ref('')
const inputRef = ref(null)
const correctCount = ref(0)
const errorCount = ref(0)
const started = ref(false)
const finished = ref(false)
const startedAt = ref(0)
const finishedAt = ref(0)
const submitting = ref(false)
const tickTimer = ref(null)
const now = ref(Date.now())

const currentItem = computed(() => items.value[index.value] || { ch: '-', code: '', pinyin: '' })
const elapsed = computed(() => {
  if (!started.value) return 0
  const end = finished.value ? finishedAt.value : now.value
  return Math.max(0, Math.round((end - startedAt.value) / 1000))
})
const accuracy = computed(() => {
  const total = correctCount.value + errorCount.value
  return total === 0 ? 0 : (correctCount.value / total) * 100
})
const wpm = computed(() => {
  if (!started.value || elapsed.value === 0) return 0
  return (correctCount.value / elapsed.value) * 60
})

const start = async () => {
  started.value = true
  startedAt.value = Date.now()
  tickTimer.value = setInterval(() => { now.value = Date.now() }, 500)
  await nextTick()
  inputRef.value?.focus?.()
}

const advance = (isCorrect) => {
  if (isCorrect) correctCount.value++
  else errorCount.value++
  inputVal.value = ''
  if (index.value + 1 >= items.value.length) {
    finish()
  } else {
    index.value++
  }
}

const evaluate = (val) => {
  if (!val) return
  // 取最后一个字符与目标比较 多余的部分忽略
  const ch = val.slice(-1)
  if (ch === currentItem.value.ch) {
    advance(true)
  } else {
    advance(false)
  }
}

const onInput = (val) => {
  if (!started.value) return
  // IME 组合中 不立即判定 等 compositionend
  // val 可能已经是合成完成的中文 也可能是英文输入直接到达
  // 简单策略: 如果是非中文字符, 视为错误并跳过
  if (/^[\x00-\x7f]$/.test(val)) {
    advance(false)
    return
  }
  // 中文字符直接评估
  evaluate(val)
}

const onCompositionEnd = (e) => {
  if (!started.value) return
  evaluate(e.target.value)
}

const skip = () => advance(false)

const finish = () => {
  finishedAt.value = Date.now()
  finished.value = true
  if (tickTimer.value) {
    clearInterval(tickTimer.value)
    tickTimer.value = null
  }
}

const reset = () => {
  if (tickTimer.value) {
    clearInterval(tickTimer.value)
    tickTimer.value = null
  }
  items.value = pickRandom(questionCount.value)
  index.value = 0
  inputVal.value = ''
  correctCount.value = 0
  errorCount.value = 0
  started.value = false
  finished.value = false
  startedAt.value = 0
  finishedAt.value = 0
}

const submit = async () => {
  submitting.value = true
  try {
    const total = correctCount.value + errorCount.value
    const res = await submitScore({
      mode: 'char',
      totalCount: total,
      correctCount: correctCount.value,
      errorCount: errorCount.value,
      durationSecond: Math.max(1, elapsed.value),
      accuracy: Number(accuracy.value.toFixed(2)),
      wpm: Number(wpm.value.toFixed(2))
    })
    if (res.code === 0) {
      ElMessage.success('成绩已提交')
    } else {
      ElMessage.error(res.msg || '提交失败')
    }
  } catch (e) {
    ElMessage.error(e?.message || '提交失败')
  } finally {
    submitting.value = false
  }
}

onUnmounted(() => {
  if (tickTimer.value) clearInterval(tickTimer.value)
})
</script>

<style scoped>
.text-primary {
  color: var(--el-color-primary);
}
.font-mono {
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, 'Liberation Mono', 'Courier New', monospace;
}
</style>
