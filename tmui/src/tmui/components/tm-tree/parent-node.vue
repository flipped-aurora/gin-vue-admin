<template>
	<view class="flex flex-row flex-row-center-start">
		<TmCheckbox
			:followTheme="props.followTheme"
			closeAni
			v-if="showCheckbox"
			:color="nodeData['color']"
			:indeterminate="_indeterminate"
			@change="checkboxChange"
			:disabled="nodeData['disabled']"
			v-model="_checked"
			:value="nodeData[props.fieldNames.id]"
			ref="checkboxRef"
		></TmCheckbox>
	</view>
</template>
<script lang="ts" setup>
import TmCheckbox from '../tm-checkbox/tm-checkbox.vue'

import { ref, getCurrentInstance, inject, computed, toRaw, watch, watchEffect, Ref } from 'vue'
import { baseNodeData } from './interface'
import { treeFlat, queryParentNode } from './util'
const checkboxRef = ref<InstanceType<typeof TmCheckbox> | null>(null)
const { proxy } = getCurrentInstance()
const props = defineProps({
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	/**
	 * 数据
	 * @description 生成树结构的数据。结构必须要有id字段。当然可以通过field-names来映射，如果你的唯一标识字段不是Id的话。
	 */
	data: {
		type: Object,
		default: () => {},
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
	}
})
//父级方法。
let parent = proxy.$parent
while (parent) {
	if (parent?.TreeParaentName == 'tmTree' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? undefined
	}
}
const TreeParentIds = inject('TreeParentIds', ref([]))
const TreePareantSelectedIds = inject('TreePareantSelectedIds', ref([]))
const color = parent?.$props.color ?? 'primary'
const nodeData = computed(
	() =>
		<baseNodeData>{
			icon: props.data['icon'] ?? '', //节点图标。
			color: props.data['color'] || color, //节点颜色主题
			disabled: props.data['disabled'] ?? false, //节点是否禁用
			text: props.data[props.fieldNames.text], //节点标题
			id: props.data[props.fieldNames.id] //节点标识
		}
)
const allListFlat = parent.getAllListData()

const showCheckbox = inject(
	'TreeNodeCheckable',
	computed(() => true)
)
//当前是否选中Id;
const _checked: Ref<string | number> = ref('')
//当前是否半选中状态。
const _indeterminate = ref(false)
//是否允许多选。
const _multiple = parent.$props.multiple ?? true
//本节点子节点所有节点，含 父和子
const flatlist = computed(() => treeFlat(props.data['children'] ?? [], props.fieldNames.id))
//所有数据父节点
let listParentFlat = queryParentNode(allListFlat, props.fieldNames.id)
let _listParentFlat = new Set(listParentFlat)
//本节点下所有子节点 ，不包含父节点
const childreNodelist = flatlist.value.filter((el) => !_listParentFlat.has(el))
//本节点下的呢父节点，不包含子
const parentNodelist = flatlist.value.filter((el) => _listParentFlat.has(el))

//初始时，是否有本节点。有本节点就需要 你全选下面的子节点。
let sletctedIdArray = TreeParentIds.value.filter((el: string | number) => el == nodeData.value.id)
if (flatlist.value.length > 0 && sletctedIdArray.length > 0) {
	if (_multiple) {
		parent?.onSelected(childreNodelist)
		parent?.onSelectedParent(parentNodelist)
	}
}

// function setChecked(){
//     setExpaned
// }
// watchEffect(()=>setExpaned())

function checkboxChange(e: boolean) {
	if (!_multiple && TreePareantSelectedIds.value.length >= 1) {
		parent?.checkAll(false)
	}
	if (e) {
		// parent?.onSelectedParent([nodeData.value.id,...flatlist.value])
		parent?.onSelectedParent([nodeData.value.id, ...parentNodelist])
		parent?.onSelected(childreNodelist)
		// console.log(childreNodelist)
	} else {
		// parent?.onUnSelectedParent([nodeData.value.id,...flatlist.value])
		parent?.onUnSelectedParent([nodeData.value.id, ...parentNodelist])
		parent?.onUnSelected(childreNodelist)
	}
	//父节点被点击。
	parent?.onCheck({ checked: e, data: toRaw(props.data) })
}

// console.log(TreePareantSelectedIds.value)
// console.log(TreePareantSelectedIds.value.filter(el=>el!=nodeData.value.id))
const ids = toRaw(nodeData.value.id)
function setExpaned() {
	//计算它的子集是否全部被选中了。
	//如果差集长度大于0表示没有全部选中。
	//排除本身 。
	let ar: Set<string | number> = new Set([...TreePareantSelectedIds.value, ...TreeParentIds.value])
	let _seleteds = flatlist.value.filter((el: string | number) => ar.has(el))
	// console.log(ids,_seleteds.length)
	// console.log(_seleteds.length)
	if (_seleteds.length == flatlist.value.length) {
		//全部选中后，把父级设置选中状态，否则为半选。
		_indeterminate.value = false
		_checked.value = nodeData.value.id
		// parent?.onSelectedParent([nodeData.value.id])
		return
	}
	if (_seleteds.length > 0) {
		_indeterminate.value = true
		_checked.value = ''
		//  parent?.onUnSelectedParent([nodeData.value.id])
		return
	}
	if (_seleteds.length == 0) {
		_indeterminate.value = false
		_checked.value = ''
		// parent?.onUnSelectedParent([nodeData.value.id])
	}
}
watchEffect(() => setExpaned())

function setExpanedParent() {
	//计算它的子集是否全部被选中了。
	//如果差集长度大于0表示没有全部选中。
	let ar: Set<string | number> = new Set(TreeParentIds.value)
	let _seleteds = flatlist.value.filter((el: string | number) => ar.has(el))
	if (_seleteds.length == flatlist.value.length) {
		//全部选中后，把父级设置选中状态，否则为半选。
		_indeterminate.value = false
		_checked.value = nodeData.value.id
		// parent?.onSelected([nodeData.value.id])
		return
	}
	if (_seleteds.length > 0) {
		_indeterminate.value = true
		_checked.value = ''
		//  parent?.onUnSelected([nodeData.value.id])
		return
	}
	if (_seleteds.length == 0) {
		_indeterminate.value = false
		_checked.value = ''
		// parent?.onUnSelected([nodeData.value.id])
	}
}
// watchEffect(()=>setExpanedParent())

watch(
	() => _checked.value,
	(newval, oldval) => {
		let ar: Set<string | number> = new Set([...TreePareantSelectedIds.value, ...TreeParentIds.value])
		let _seleteds = flatlist.value.filter((el: string | number) => ar.has(el))
		if (_seleteds.length == flatlist.value.length) {
			parent?.onSelectedParent([nodeData.value.id])
			return
		}
		if (_seleteds.length > 0) {
			parent?.onUnSelectedParent([nodeData.value.id])
			return
		}
		if (_seleteds.length == 0) {
			parent?.onUnSelectedParent([nodeData.value.id])
		}
	}
)

function setStatus() {
	if (nodeData.value['disabled'] === true) {
		return
	}
	checkboxRef.value?.hanlerClick()
}

defineExpose({ setStatus })
</script>
