import service from '@/utils/request'

export const uploadInit = (data) =>
  service({ url: '/mediaUpload/init', method: 'post', data })

export const uploadChunk = (formData, signal) =>
  service({ url: '/mediaUpload/chunk', method: 'post', data: formData, signal, headers: { 'Content-Type': 'multipart/form-data' } })

export const uploadComplete = (data) =>
  service({ url: '/mediaUpload/complete', method: 'post', data })

export const uploadCancel = (uploadId) =>
  service({ url: `/mediaUpload/${uploadId}`, method: 'delete' })
