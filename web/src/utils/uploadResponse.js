export const getUploadErrorMessage = (response) => {
  if (response?.code === 0) {
    return ''
  }

  return response?.msg || '上传失败'
}
