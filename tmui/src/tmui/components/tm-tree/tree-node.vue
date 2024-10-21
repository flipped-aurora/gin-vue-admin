<template>
	<view class="flex flex-row flex-row-center-start">
		<childNode
			@change="_checked = $event"
			ref="parentNodeRefNode"
			:followTheme="props.followTheme"
			v-if="!props.data['children']"
			:fieldNames="props.fieldNames"
			:data="props.data"
		></childNode>
		<parentNode
			ref="parentNodeRef"
			:followTheme="props.followTheme"
			v-if="props.data['children']"
			:fieldNames="props.fieldNames"
			:data="props.data"
		></parentNode>
		<TmIcon _class="pr-16" :color="nodeData['color']" :font-size="28" v-if="nodeData['icon']" :name="nodeData['icon']"></TmIcon>
		<TmText
			:color="_checked ? 'primary' : ''"
			:_class="nodeData['disabled'] && !props.data['children'] ? 'opacity-7' : ''"
			@click="textClick(nodeData)"
			:font-size="28"
			:label="nodeData.text"
		></TmText>
	</view>
</template>
<script lang="ts" setup>
import TmCheckbox from '../tm-checkbox/tm-checkbox.vue'
import TmIcon from '../tm-icon/tm-icon.vue'
import TmText from '../tm-text/tm-text.vue'
import childNode from './child-node.vue'
import parentNode from './parent-node.vue'
import { ref, getCurrentInstance, inject, computed, toRaw, watch, watchEffect, Ref } from 'vue'
import { baseNodeData } from './interface'
import { treeFlat } from './util'
const parentNodeRef = ref<InstanceType<typeof parentNode> | null>(null)
const parentNodeRefNode = ref<InstanceType<typeof childNode> | null>(null)
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['textClick'])
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
const _checked = ref(false)
//父级方法。
let parent = proxy.$parent
while (parent) {
	if (parent?.TreeParaentName == 'tmTree' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? undefined
	}
}

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

function textClick(node: baseNodeData) {
	//如果存在子级。
	if (!props.data['children']) {
		parentNodeRefNode.value?.setStatus()
	}
	emits('textClick', node)
}
</script>
