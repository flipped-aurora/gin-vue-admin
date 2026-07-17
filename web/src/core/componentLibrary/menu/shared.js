// menu 的有限枚举与纯工具单一事实源（不含 Vue / 主题依赖，供组件与组合式共用，避免循环依赖）。
export const MENU_THEMES = ['design', 'light', 'group']

// 顶级横向菜单项间距（px）：同时驱动模板 gap 与溢出计算，避免二者漂移。
export const MENU_GAP = 8

// 菜单缩进：基础左内边距 + 每级步长（design 风格步长更大，配合左侧主色竖条）。
const MENU_INDENT_BASE = 12
const MENU_INDENT_STEP = 12
const MENU_INDENT_STEP_DESIGN = 16
export const menuIndent = (depth, theme) =>
  MENU_INDENT_BASE +
  depth * (theme === 'design' ? MENU_INDENT_STEP_DESIGN : MENU_INDENT_STEP) +
  'px'

// 过滤隐藏项：菜单各处渲染前的统一入口。
export const visibleItems = (items) => {
  const list = items || []
  return list.filter((i) => !i.hidden)
}

// key 是否落在某节点子树内（含自身）。
export const isKeyInSubtree = (node, key) => {
  const children = node.children || []
  return node.name === key || children.some((c) => isKeyInSubtree(c, key))
}

// 在根级菜单里求 key 所属的一级菜单节点（跳过隐藏一级），找不到返回 null。
export const resolveActiveTop = (rootMenus, key) => {
  const topMenus = rootMenus || []
  return topMenus.find((top) => !top.hidden && isKeyInSubtree(top, key)) || null
}

/**
 * 在菜单树中按 name 求从根到该节点的祖先路径（含自身），找不到返回 []。
 * @param {Array<{name:string,children?:Array}>} menus
 * @param {string} targetKey
 * @returns {string[]}
 */
export function findMenuPath(menus, targetKey) {
  const path = []
  const dfs = (item) => {
    path.push(item.name)
    if (item.name === targetKey) return true
    for (const child of item.children || []) {
      if (dfs(child)) return true
    }
    path.pop()
    return false
  }
  for (const m of menus || []) {
    if (dfs(m)) return [...path]
  }
  return []
}
