<template>
  <div class="gva-container">
    <div class="gva-table-box">
      <el-upload drag :auto-upload="false" :show-file-list="false" multiple :on-change="onChange">
        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
        <div class="el-upload__text">拖拽文件到此处,或 <em>点击选择</em>(支持大文件分片/断点续传)</div>
      </el-upload>

      <el-table :data="tasks" style="margin-top:16px">
        <el-table-column prop="name" label="文件名" min-width="180" show-overflow-tooltip />
        <el-table-column label="大小" width="110">
          <template #default="{ row }">{{ formatSize(row.size) }}</template>
        </el-table-column>
        <el-table-column label="进度" min-width="200">
          <template #default="{ row }">
            <el-progress :percentage="row.status === 'hashing' ? row.hashProgress : row.progress"
              :status="row.status === 'error' ? 'exception' : row.status === 'done' ? 'success' : ''" />
          </template>
        </el-table-column>
        <el-table-column label="速度" width="110">
          <template #default="{ row }">{{ row.status === 'uploading' ? formatSize(row.speed) + '/s' : '-' }}</template>
        </el-table-column>
        <el-table-column label="剩余" width="90">
          <template #default="{ row }">{{ row.status === 'uploading' && row.eta ? Math.round(row.eta) + 's' : '-' }}</template>
        </el-table-column>
        <el-table-column label="状态" width="100">
          <template #default="{ row }"><el-tag :type="tagType(row.status)">{{ statusText(row.status) }}</el-tag></template>
        </el-table-column>
        <el-table-column label="操作" width="180">
          <template #default="{ row }">
            <el-button v-if="row.status === 'uploading'" link type="primary" @click="pause(row)">暂停</el-button>
            <el-button v-if="row.status === 'paused'" link type="primary" @click="resume(row)">继续</el-button>
            <el-button v-if="row.status === 'error'" link type="warning" @click="retry(row)">重试</el-button>
            <el-button link type="danger" @click="cancel(row)">取消</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { useChunkUpload } from '@/composables/useChunkUpload'
import { UploadFilled } from '@element-plus/icons-vue'

defineOptions({ name: 'ChunkUpload' })

const { tasks, add, pause, resume, retry, cancel } = useChunkUpload()

const onChange = (uploadFile) => { if (uploadFile.raw) add([uploadFile.raw]) }

const formatSize = (b) => {
  if (!b) return '0 B'
  const u = ['B', 'KB', 'MB', 'GB']; let i = 0; let n = b
  while (n >= 1024 && i < u.length - 1) { n /= 1024; i++ }
  return `${n.toFixed(1)} ${u[i]}`
}
const statusText = (s) => ({ queued: '排队', hashing: '校验中', uploading: '上传中', paused: '已暂停', merging: '合并中', done: '完成', error: '失败', canceled: '已取消' }[s] || s)
const tagType = (s) => ({ done: 'success', error: 'danger', paused: 'warning', uploading: 'primary' }[s] || 'info')
</script>
