import SparkMD5 from 'spark-md5'
import { reactive } from 'vue'
import { uploadInit, uploadChunk, uploadComplete, uploadCancel } from '@/api/media'

const DEFAULT = { chunkSize: 5 * 1024 * 1024, concurrency: 3, maxRetry: 3 }

// 计算整文件与每片 md5,不整文件入内存
async function hashFile(file, chunkSize, onProgress) {
  const spark = new SparkMD5.ArrayBuffer()
  const chunkHashes = []
  const total = Math.ceil(file.size / chunkSize)
  for (let i = 0; i < total; i++) {
    const blob = file.slice(i * chunkSize, (i + 1) * chunkSize)
    const buf = await blob.arrayBuffer()
    spark.append(buf)
    const cs = new SparkMD5.ArrayBuffer()
    cs.append(buf)
    chunkHashes.push(cs.end())
    onProgress && onProgress(Math.round(((i + 1) / total) * 100))
  }
  return { fileHash: spark.end(), chunkHashes, total }
}

export function useChunkUpload(opts = {}) {
  const cfg = { ...DEFAULT, ...opts }
  const tasks = reactive([])

  function makeTask(file) {
    return reactive({
      id: `${file.name}-${file.size}-${file.lastModified}`,
      file, name: file.name, size: file.size,
      status: 'queued', // queued|hashing|uploading|paused|merging|done|error|canceled
      progress: 0, hashProgress: 0, speed: 0, eta: 0,
      uploadId: 0, error: '', media: null,
      _controllers: new Set(), _paused: false, _meta: null
    })
  }

  function add(fileList) {
    for (const f of fileList) {
      if (!tasks.find((t) => t.id === `${f.name}-${f.size}-${f.lastModified}`)) {
        const t = makeTask(f)
        tasks.push(t)
        run(t)
      }
    }
  }

  async function run(t) {
    try {
      if (!t._meta) {
        t.status = 'hashing'
        t._meta = await hashFile(t.file, cfg.chunkSize, (p) => (t.hashProgress = p))
      }
      const initRes = await uploadInit({
        fileName: t.name, fileHash: t._meta.fileHash, fileSize: t.size,
        chunkSize: cfg.chunkSize, chunkTotal: t._meta.total
      })
      const d = initRes.data
      if (d.instant) { t.progress = 100; t.status = 'done'; t.media = d.media; return }
      t.uploadId = d.uploadId
      const done = new Set(d.uploadedChunks || [])
      const pending = []
      for (let i = 0; i < t._meta.total; i++) if (!done.has(i)) pending.push(i)
      t.progress = Math.round((done.size / t._meta.total) * 100)
      await uploadPending(t, pending, done)
    } catch (e) {
      if (t.status !== 'paused' && t.status !== 'canceled') { t.status = 'error'; t.error = String(e?.message || e) }
    }
  }

  async function uploadPending(t, pending, done) {
    t.status = 'uploading'
    let cursor = 0
    let startTs = Date.now(); let startDone = done.size
    const worker = async () => {
      while (cursor < pending.length) {
        if (t._paused) return
        const idx = pending[cursor++]
        await uploadOne(t, idx)
        done.add(idx)
        t.progress = Math.round((done.size / t._meta.total) * 100)
        const dt = (Date.now() - startTs) / 1000
        if (dt > 0) {
          t.speed = ((done.size - startDone) * cfg.chunkSize) / dt
          t.eta = t.speed > 0 ? ((t._meta.total - done.size) * cfg.chunkSize) / t.speed : 0
        }
      }
    }
    await Promise.all(Array.from({ length: cfg.concurrency }, worker))
    if (t._paused || t.status === 'canceled') return
    if (done.size === t._meta.total) {
      t.status = 'merging'
      const res = await uploadComplete({ uploadId: t.uploadId })
      t.media = res.data.media; t.progress = 100; t.status = 'done'
    }
  }

  async function uploadOne(t, idx) {
    const blob = t.file.slice(idx * cfg.chunkSize, (idx + 1) * cfg.chunkSize)
    for (let attempt = 0; attempt <= cfg.maxRetry; attempt++) {
      if (t._paused) throw new Error('paused')
      const ctrl = new AbortController()
      t._controllers.add(ctrl)
      try {
        const fd = new FormData()
        fd.append('uploadId', t.uploadId)
        fd.append('chunkIndex', idx)
        fd.append('chunkHash', t._meta.chunkHashes[idx])
        fd.append('chunk', blob)
        await uploadChunk(fd, ctrl.signal)
        t._controllers.delete(ctrl)
        return
      } catch (e) {
        t._controllers.delete(ctrl)
        if (t._paused) throw new Error('paused')
        if (attempt === cfg.maxRetry) throw e
        await new Promise((r) => setTimeout(r, 300 * 2 ** attempt))
      }
    }
  }

  function pause(t) { t._paused = true; t.status = 'paused'; t._controllers.forEach((c) => c.abort()); t._controllers.clear() }
  function resume(t) { if (t.status === 'paused') { t._paused = false; run(t) } }
  function retry(t) { if (t.status === 'error') { t._paused = false; run(t) } }
  async function cancel(t) {
    t.status = 'canceled'; t._paused = true
    t._controllers.forEach((c) => c.abort()); t._controllers.clear()
    if (t.uploadId) { try { await uploadCancel(t.uploadId) } catch (e) { /* ignore */ } }
    const i = tasks.indexOf(t); if (i >= 0) tasks.splice(i, 1)
  }

  return { tasks, add, pause, resume, retry, cancel, _hashFile: hashFile }
}
