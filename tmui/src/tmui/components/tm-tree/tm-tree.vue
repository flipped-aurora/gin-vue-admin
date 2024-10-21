<template>
	<view>
		<base-node-vue
			@textClick="emits('node-click', $event)"
			:followTheme="props.followTheme"
			:fieldNames="props.fieldNames"
			:data="listData"
		></base-node-vue>
	</view>
</template>
<script lang="ts" setup>
/**
 * 树状结构
 * @description 用来展示层级，文件夹结构体等.
 */
import { ref, computed, provide, reactive, watch, readonly, toRaw } from 'vue'
import { baseNodeData } from './interface'
import baseNodeVue from './base-node.vue'
import { getNodeRouter, treeFlat, queryNodeIsParent, queryParentNode } from './util'
/**
 * 事件说明
 * node-click:节点标题点击时触发的事件。
 * check:节点复选框状态变化时触发的事件。
 * expand:父节点展开/关闭时触发的事件。
 */
const emits = defineEmits(['node-click', 'check', 'expand', 'update:selectedId', 'update:expandedId'])
const props = defineProps({
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	color: {
		type: String,
		default: 'primary'
	},
	//父节点打开时的图标。
	expandedIconOpen: {
		type: String,
		default: 'tmicon-sort-down'
	},
	//父节点被关闭时显示的图标。
	expandedIconClose: {
		type: String,
		default: 'tmicon-caret-right'
	},
	//允许在节点前添加复选框。
	checkable: {
		type: Boolean,
		default: true
	},
	//是否允许多选
	multiple: {
		type: Boolean,
		default: true
	},

	/**
	 * 展开的父节点可以使用v-model:expanded-id
	 */
	expandedId: {
		type: Array,
		default: () => []
	},
	/**
	 * 默认展开的父节点
	 */
	defaultExpandedId: {
		type: Array,
		default: () => []
	},
	//选中的节点。可以使用v-model:selected-id
	selectedId: {
		type: Array,
		default: () => []
	},
	//默认选中的节点
	defaultSelectedId: {
		type: Array,
		default: () => []
	},
	/**
	 * 生成树结构的数据。结构必须要有id字段。当然可以通过field-names来映射，如果你的唯一标识字段不是Id的话。
	 */
	data: {
		type: Array,
		default: () => [],
		required: true
	},
	/**
	 * 数据结构字段的映射表。
	 */
	fieldNames: {
		type: Object,
		default: () => {
			return {
				id: 'id',
				text: 'text'
			}
		}
	},
	//显示连线。
	showLine: {
		type: [Boolean, String],
		default: true
	}
})
let listData = ref(props.data)
//所有的节点被平铺为一维。当data改变时就要取交集数据，差集需要删除。
let listDataFlat = [...treeFlat(props.data, props.fieldNames.id)]
//所有父节点
let listParentFlat = queryParentNode(props.data, props.fieldNames.id)
const selectedIds = ref([...new Set([...props.defaultSelectedId, ...props.selectedId])])
const SelectedExpandedId = ref([...new Set([...props.defaultExpandedId, ...props.expandedId])])
//保存选中的父节点。
const parentIds = ref<Array<string | number>>([])

//是否单选。
if (!props.multiple && selectedIds.value.length >= 1) {
	selectedIds.value = [selectedIds.value[0]]
}

//分离父和子节点。
let __selectedIdsSet = new Set(listParentFlat)
parentIds.value = selectedIds.value.filter((el) => __selectedIdsSet.has(el))

//过滤错误的id集。意思是如果用户提供的id不在数据props.data结构体中，就需要排除掉。以增加使用者的容错率。
// 2022年11月24日 注释下方两行,可能导致用户配音是先取得选中的值,再设置数据集,会导致过滤掉应有的数据.
// const ___a_selectedIds = new Set(selectedIds.value)
// selectedIds.value = listDataFlat.filter(el=>___a_selectedIds.has(el))
selectedIds.value = selectedIds.value.filter((el) => !__selectedIdsSet.has(el))

const ___a_ExpandedId = new Set(SelectedExpandedId.value)
SelectedExpandedId.value = listDataFlat.filter((el) => ___a_ExpandedId.has(el))

//向下传递当前默认选中的子节点。
provide('TreePareantSelectedIds', readonly(selectedIds))
//向下传递默认打开的父节点
provide('TreePareantSelectedExpandedId', readonly(SelectedExpandedId))
//父节点
provide('TreeParentIds', readonly(parentIds))

function onSelected(key: Array<string | number>) {
	selectedIds.value = [...new Set([...selectedIds.value, ...key])]
	emits('update:selectedId', selectedIds.value)
}
function onUnSelected(key: Array<string | number>) {
	let a = new Set(key)
	selectedIds.value = [...new Set([...selectedIds.value].filter((el) => !a.has(el)))]
	emits('update:selectedId', selectedIds.value)
}
function onCheck(e: baseNodeData) {
	emits('check', e)
}
function onExpand(e: any) {
	SelectedExpandedId.value = [...new Set([...SelectedExpandedId.value, e.data[props.fieldNames.id]])]
	emits('expand', { ...e.data, expanded: e.expand })
	emits('update:expandedId', SelectedExpandedId.value)
}
function onUnExpand(e: any) {
	SelectedExpandedId.value = [...new Set([...SelectedExpandedId.value].filter((el) => el != e.data[props.fieldNames.id]))]
	emits('expand', { ...e.data, expanded: e.expand })
	emits('update:expandedId', SelectedExpandedId.value)
}
function onSelectedParent(key: Array<string | number>) {
	parentIds.value = [...new Set([...parentIds.value, ...key])]
}
function onUnSelectedParent(key: Array<string | number>) {
	let a = new Set(key)
	parentIds.value = [...new Set([...parentIds.value].filter((el) => !a.has(el)))]
}
function getAllListData() {
	return toRaw(props.data)
}
watch([() => props.expandedId], () => {
	SelectedExpandedId.value = [...props.expandedId]
})
watch(
	[() => props.data],
	() => {
		listData.value = props.data
		listDataFlat = [...treeFlat(props.data)]
		//所有父节点
		listParentFlat = queryParentNode(props.data, props.fieldNames.id)
		//分离父和子节点。
		let __selectedIdsSet = new Set(listParentFlat)
		parentIds.value = selectedIds.value.filter((el) => __selectedIdsSet.has(el))

		let _a_selectedIds = new Set(selectedIds.value)
		let s = listDataFlat.filter((el) => _a_selectedIds.has(el))
		selectedIds.value = s.filter((el) => !__selectedIdsSet.has(el))
		emits('update:selectedId', selectedIds.value)
		let _a_ExpandedId = new Set(SelectedExpandedId.value)
		SelectedExpandedId.value = listDataFlat.filter((el) => _a_ExpandedId.has(el))
		emits('update:expandedId', SelectedExpandedId.value)
	},
	{ deep: true }
)
provide(
	'TreeNodeCheckable',
	computed(() => props.checkable)
)

/** 下方的方法是对象展示的ref方法 **/
/**
 * 改变所有节点状态。
 * checked:
 * true时选中所有节点。
 * false时取消所有节点
 */
function checkAll(checked = true) {
	if (checked == true) {
		onSelected(listDataFlat)
		onSelectedParent(listParentFlat)
	} else {
		selectedIds.value = []
		onUnSelectedParent(listParentFlat)
	}
	emits('update:selectedId', [...selectedIds.value, ...parentIds.value])
}
/**
 * 指定节点状态
 * @param {string|number} key 节点id
 * @param {boolean} checked 节点状态
 * description 注意，如果指定的是父节点，将会选中他的所有节节点，反之取消它所有的子节点。
 */
function checkNode(key: string | number, checked: boolean) {
	//先检查 是否存在节点，不存在返回失败。
	if (!new Set(listDataFlat).has(key)) {
		console.error('不存在该节点')
		return false
	}
	let parentData = queryNodeIsParent(toRaw(props.data), key, props.fieldNames.id)
	//选中的是父节点
	if (parentData) {
		//拆分父和子节点。

		//先平铺所有节点。
		let flat_nodataNodes = treeFlat(parentData.children, props.fieldNames.id)
		let all_parent = new Set(listParentFlat)
		//分离出父节点。
		let flat_nodataParentNodes = [...flat_nodataNodes, key].filter((el) => all_parent.has(el))
		//分离也所有子节点
		let flat_nodataChildNodes = [...flat_nodataNodes, key].filter((el) => !all_parent.has(el))

		if (checked == true) {
			onSelected(flat_nodataChildNodes)
			onSelectedParent(flat_nodataParentNodes)
		} else {
			onUnSelected(flat_nodataChildNodes)
			onUnSelectedParent(flat_nodataParentNodes)
		}
	} else {
		if (checked == true) {
			onSelected([key])
		} else {
			onUnSelected([key])
		}
	}
	emits('update:selectedId', [...selectedIds.value, ...parentIds.value])
	return true
}
/**
 * 展开或者关闭所有父节点状态
 * @param {Boolean}  checked 指定节点打开还是状态的状态。
 */
function expandAll(checked = true) {
	let flatarray = queryParentNode(toRaw(props.data), props.fieldNames.id)
	if (checked) {
		SelectedExpandedId.value = [...new Set([...SelectedExpandedId.value, ...flatarray])]
	} else {
		SelectedExpandedId.value = []
	}
	emits('update:expandedId', SelectedExpandedId.value)
}
/**
 * 指定父节点展开状态
 * @param {string|number} key 节点id
 * @param {boolean} checked 节点状态
 */
function expandNode(key: string | number, checked: boolean) {
	//先检查 是否存在节点，不存在返回失败。
	if (!new Set(listDataFlat).has(key)) {
		console.error('不存在该节点')
		return false
	}
	let parentData = queryNodeIsParent(toRaw(props.data, props.fieldNames.id), key)
	if (!parentData) {
		console.error('该节点非父节点')
		return false
	}
	if (checked == true) {
		SelectedExpandedId.value = [...new Set([...SelectedExpandedId.value, key])]
	} else {
		SelectedExpandedId.value = [...new Set(SelectedExpandedId.value.filter((el) => el != key))]
	}
	emits('update:expandedId', SelectedExpandedId.value)
	return true
}
/**
 * 显示某一节点
 * @param key 需要要打开所在路径显示的节点
 * @description 将会打开它所在的所有父节点
 */
function showNode(key: string | number) {
	//先检查 是否是父节点。
	let nodepath = getNodePath(key)
	let parentData = queryNodeIsParent(toRaw(props.data), key, props.fieldNames.id)
	if (!parentData) {
		nodepath = nodepath.filter((el) => el != key)
	}

	SelectedExpandedId.value = [...new Set([...SelectedExpandedId.value, ...nodepath])]
	emits('update:expandedId', SelectedExpandedId.value)
}
/**
 * 获取选中的节点key数组
 * @param {String}  strategy all:返回所有选中的节点,parent:父子节点都选中时只返回父节点,children:只返回子节点
 */
function getCheckedNodes(strategy = 'all') {
	if (strategy == 'all') {
		return [...toRaw(selectedIds.value), ...parentIds.value]
	} else if (strategy == 'parent') {
		return toRaw(parentIds.value)
	} else if (strategy == 'children') {
		return toRaw(selectedIds.value)
	} else {
		return []
	}
}
/**
 * 获取当前展开的节点
 */
function getExpandedNodes() {
	return toRaw(SelectedExpandedId.value)
}
/**
 * 反回节点路径
 * @param key 节点id
 * @description 从父节点开始一直到本节点的路径数组。
 */
function getNodePath(key: string | number) {
	return getNodeRouter(toRaw(props.data), key, props.fieldNames.id)
}

/**
 * 搜索数据
 * @param keyword 关键词
 */
function searchData(keyword) {
	const loop = (data) => {
		const result = []
		data.forEach((item) => {
			if (item[props.fieldNames.text].toLowerCase().indexOf(keyword.toLowerCase()) > -1) {
				result.push({ ...item })
			} else if (item.children) {
				const filterData = loop(item.children)
				if (filterData.length) {
					result.push({
						...item,
						children: filterData
					})
				}
			}
		})
		return result
	}

	return loop(originTreeData)
}
defineExpose({
	TreeParaentName: 'tmTree',
	onUnSelected,
	onSelected,
	onCheck,
	onExpand,
	onUnExpand,
	onSelectedParent,
	onUnSelectedParent,
	getAllListData,
	checkAll,
	checkNode,
	expandAll,
	expandNode,
	getCheckedNodes,
	getExpandedNodes,
	getNodePath,
	showNode
})
</script>
