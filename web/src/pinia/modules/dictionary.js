import { findSysDictionary } from '@/api/sysDictionary'
import { getDictionaryTreeListByType } from '@/api/sysDictionaryDetail'

import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useDictionaryStore = defineStore('dictionary', () => {
  const dictionaryMap = ref({})

  const setDictionaryMap = (dictionaryRes) => {
    dictionaryMap.value = { ...dictionaryMap.value, ...dictionaryRes }
  }

  // 过滤树形数据的深度
  const filterTreeByDepth = (items, currentDepth, targetDepth) => {
    if (targetDepth === 0) {
      // depth=0 返回全部数据
      return items
    }

    if (currentDepth >= targetDepth) {
      // 达到目标深度，移除children
      return items.map((item) => ({
        label: item.label,
        value: item.value,
        extend: item.extend
      }))
    }

    // 递归处理子项
    return items.map((item) => ({
      label: item.label,
      value: item.value,
      extend: item.extend,
      children: item.children
        ? filterTreeByDepth(item.children, currentDepth + 1, targetDepth)
        : undefined
    }))
  }

  // 将树形结构扁平化为数组（用于兼容原有的平铺格式）
  const flattenTree = (items) => {
    const result = []

    const traverse = (nodes) => {
      nodes.forEach((item) => {
        result.push({
          label: item.label,
          value: item.value,
          extend: item.extend
        })

        if (item.children && item.children.length > 0) {
          traverse(item.children)
        }
      })
    }

    traverse(items)
    return result
  }

  // 标准化树形数据，确保每个节点都包含标准的字段格式
  const normalizeTreeData = (items) => {
    return items.map((item) => ({
      label: item.label,
      value: item.value,
      extend: item.extend,
      children:
        item.children && item.children.length > 0
          ? normalizeTreeData(item.children)
          : undefined
    }))
  }

  // 根据value和depth查找指定节点并返回其children
  const findNodeByValue = (
    items,
    targetValue,
    currentDepth = 1,
    maxDepth = 0
  ) => {
    for (const item of items) {
      // 如果找到目标value的节点
      if (item.value === targetValue) {
        // 如果maxDepth为0，返回所有children
        if (maxDepth === 0) {
          return item.children ? normalizeTreeData(item.children) : []
        }
        // 否则根据depth限制返回children
        if (item.children && item.children.length > 0) {
          return filterTreeByDepth(item.children, 1, maxDepth)
        }
        return []
      }

      // 如果当前深度小于最大深度，继续在children中查找
      if (
        item.children &&
        item.children.length > 0 &&
        (maxDepth === 0 || currentDepth < maxDepth)
      ) {
        const result = findNodeByValue(
          item.children,
          targetValue,
          currentDepth + 1,
          maxDepth
        )
        if (result !== null) {
          return result
        }
      }
    }
    return null
  }

  const getDictionary = async (type, depth = 0, value = null) => {
    // 如果传入了value参数，则查找指定节点的children
    if (value !== null) {
      // 构建缓存key，包含value和depth信息
      const cacheKey = `${type}_value_${value}_depth_${depth}`

      if (
        dictionaryMap.value[cacheKey] &&
        dictionaryMap.value[cacheKey].length
      ) {
        return dictionaryMap.value[cacheKey]
      }

      try {
        // 获取完整的树形结构数据
        const treeRes = await getDictionaryTreeListByType({ type })
        if (
          treeRes.code === 0 &&
          treeRes.data &&
          treeRes.data.list &&
          treeRes.data.list.length > 0
        ) {
          // 查找指定value的节点并返回其children
          const targetNodeChildren = findNodeByValue(
            treeRes.data.list,
            value,
            1,
            depth
          )

          if (targetNodeChildren !== null) {
            let resultData
            if (depth === 0) {
              // depth=0 时返回完整的children树形结构
              resultData = targetNodeChildren
            } else {
              // 其他depth值：扁平化children数据
              resultData = flattenTree(targetNodeChildren)
            }

            const dictionaryRes = {}
            dictionaryRes[cacheKey] = resultData
            setDictionaryMap(dictionaryRes)
            return dictionaryMap.value[cacheKey]
          } else {
            // 如果没找到指定value的节点，返回空数组
            return []
          }
        }
      } catch (error) {
        console.error('根据value获取字典数据失败:', error)
        return []
      }
    }

    // 原有的逻辑：不传value参数时的处理
    // 构建缓存key，包含depth信息
    const cacheKey = depth === 0 ? `${type}_tree` : `${type}_depth_${depth}`

    if (dictionaryMap.value[cacheKey] && dictionaryMap.value[cacheKey].length) {
      return dictionaryMap.value[cacheKey]
    } else {
      try {
        // 首先尝试获取树形结构数据
        const treeRes = await getDictionaryTreeListByType({ type })
        if (
          treeRes.code === 0 &&
          treeRes.data &&
          treeRes.data.list &&
          treeRes.data.list.length > 0
        ) {
          // 使用树形结构数据
          const treeData = treeRes.data.list

          let resultData
          if (depth === 0) {
            // depth=0 时返回完整的树形结构，但要确保字段格式标准化
            resultData = normalizeTreeData(treeData)
          } else {
            // 其他depth值：根据depth参数过滤数据，然后扁平化
            const filteredData = filterTreeByDepth(treeData, 1, depth)
            resultData = flattenTree(filteredData)
          }

          const dictionaryRes = {}
          dictionaryRes[cacheKey] = resultData
          setDictionaryMap(dictionaryRes)
          return dictionaryMap.value[cacheKey]
        } else {
          // 如果没有树形数据，回退到原有的平铺方式
          const res = await findSysDictionary({ type })
          if (res.code === 0) {
            const dictionaryRes = {}
            const dict = []
            res.data.resysDictionary.sysDictionaryDetails &&
              res.data.resysDictionary.sysDictionaryDetails.forEach((item) => {
                dict.push({
                  label: item.label,
                  value: item.value,
                  extend: item.extend
                })
              })
            dictionaryRes[cacheKey] = dict
            setDictionaryMap(dictionaryRes)
            return dictionaryMap.value[cacheKey]
          }
        }
      } catch (error) {
        console.error('获取字典数据失败:', error)
        // 发生错误时回退到原有方式
        const res = await findSysDictionary({ type })
        if (res.code === 0) {
          const dictionaryRes = {}
          const dict = []
          res.data.resysDictionary.sysDictionaryDetails &&
            res.data.resysDictionary.sysDictionaryDetails.forEach((item) => {
              dict.push({
                label: item.label,
                value: item.value,
                extend: item.extend
              })
            })
          dictionaryRes[cacheKey] = dict
          setDictionaryMap(dictionaryRes)
          return dictionaryMap.value[cacheKey]
        }
      }
    }
  }

  return {
    dictionaryMap,
    setDictionaryMap,
    getDictionary
  }
})
