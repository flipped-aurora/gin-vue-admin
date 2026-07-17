import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'
import { ref } from 'vue'
import { getUrl } from './image'

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
  } else {
    return ''
  }
}
export const formatDate = (time) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
  } else {
    return ''
  }
}

export const filterDict = (value, options) => {
  // 递归查找函数
  const findInOptions = (opts, targetValue) => {
    if (!opts || !Array.isArray(opts)) return null
    
    for (const item of opts) {
      if (item.value === targetValue) {
        return item
      }
      
      if (item.children && Array.isArray(item.children)) {
        const found = findInOptions(item.children, targetValue)
        if (found) return found
      }
    }
    
    return null
  }
  
  const rowLabel = findInOptions(options, value)
  return rowLabel && rowLabel.label
}

export const filterDataSource = (dataSource, value) => {
  // 递归查找函数
  const findInDataSource = (data, targetValue) => {
    if (!data || !Array.isArray(data)) return null
    
    for (const item of data) {
      // 检查当前项是否匹配
      if (item.value === targetValue) {
        return item
      }
      
      // 如果有children属性，递归查找
      if (item.children && Array.isArray(item.children)) {
        const found = findInDataSource(item.children, targetValue)
        if (found) return found
      }
    }
    
    return null
  }
  
  if (Array.isArray(value)) {
    return value.map((item) => {
      const rowLabel = findInDataSource(dataSource, item)
      return rowLabel?.label
    })
  }
  
  const rowLabel = findInDataSource(dataSource, value)
  return rowLabel?.label
}

export const getDictFunc = async (type) => {
  const dicts = await getDict(type)
  return dicts
}

export const ReturnArrImg = (arr) => {
  const imgArr = []
  if (arr instanceof Array) {
    // 如果是数组类型
    for (const arrKey in arr) {
        imgArr.push(getUrl(arr[arrKey]))
    }
  } else {
    imgArr.push(getUrl(arr))
  }
  return imgArr
}

export const returnArrImg = ReturnArrImg

export const onDownloadFile = (url) => {
  window.open(getUrl(url))
}

const baseUrl = ref(import.meta.env.VITE_BASE_API)

export const getBaseUrl = () => {
  return baseUrl.value === '/' ? '' : baseUrl.value
}

export const CreateUUID = () => {
  let d = new Date().getTime()
  if (window.performance && typeof window.performance.now === 'function') {
    d += performance.now()
  }
  return '00000000-0000-0000-0000-000000000000'.replace(/0/g, (c) => {
    const r = (d + Math.random() * 16) % 16 | 0 // d是随机种子
    d = Math.floor(d / 16)
    return (c === '0' ? r : (r & 0x3) | 0x8).toString(16)
  })
}
