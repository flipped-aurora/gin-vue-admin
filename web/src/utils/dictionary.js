import { useDictionaryStore } from '@/pinia/modules/dictionary'

/**
 * 生成字典缓存key
 * @param {string} type - 字典类型
 * @param {number} depth - 深度参数
 * @param {string|number|null} value - 指定节点的value
 * @returns {string} 缓存key
 */
const generateCacheKey = (type, depth, value) => {
  if (value !== null && value !== undefined) {
    return `${type}_value_${value}_depth_${depth}`
  }
  return depth === 0 ? `${type}_tree` : `${type}_depth_${depth}`
}

/**
 * 获取字典数据
 * @param {string} type - 字典类型，必填
 * @param {Object} options - 可选参数
 * @param {number} options.depth - 指定获取字典的深度，默认为0（完整树形结构）
 * @param {string|number|null} options.value - 指定节点的value，获取该节点的children，默认为null
 * @returns {Promise<Array>} 字典数据数组
 * @example
 * // 获取完整的字典树形结构
 * const dictTree = await getDict('user_status')
 *
 * // 获取指定深度的扁平化字典数据
 * const dictFlat = await getDict('user_status', {
 *  depth: 2
 * })
 *
 * // 获取指定节点的children
 * const children = await getDict('user_status', {
 *  value: 'active'
 * })
 */
export const getDict = async (
  type,
  options = {
    depth: 0,
    value: null
  }
) => {
  // 参数验证
  if (!type || typeof type !== 'string') {
    console.warn('getDict: type参数必须是非空字符串')
    return []
  }

  if (typeof options.depth !== 'number' || options.depth < 0) {
    console.warn('getDict: depth参数必须是非负数')
    options.depth = 0
  }

  try {
    const dictionaryStore = useDictionaryStore()

    // 调用store方法获取字典数据
    await dictionaryStore.getDictionary(type, options.depth, options.value)

    // 生成缓存key
    const cacheKey = generateCacheKey(type, options.depth, options.value)

    // 从缓存中获取数据
    const result = dictionaryStore.dictionaryMap[cacheKey]

    // 返回数据，确保返回数组
    return Array.isArray(result) ? result : []
  } catch (error) {
    console.error('getDict: 获取字典数据失败', { type, options, error })
    return []
  }
}

//  字典文字展示方法
export const showDictLabel = (
  dict,
  code,
  keyCode = 'value',
  valueCode = 'label'
) => {
  if (!dict) {
    return ''
  }
  const dictMap = {}
  dict.forEach((item) => {
    if (Reflect.has(item, keyCode) && Reflect.has(item, valueCode)) {
      dictMap[item[keyCode]] = item[valueCode]
    }
  })
  return Reflect.has(dictMap, code) ? dictMap[code] : ''
}
