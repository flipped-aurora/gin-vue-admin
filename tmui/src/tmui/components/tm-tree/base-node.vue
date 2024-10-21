<template>
	<view class="pa-24">
		<view v-for="(item, index) in treeNodeData" :key="index">
			<expandedNode
				:followTheme="props.followTheme"
				:fieldNames="props.fieldNames"
				:data="item"
				ref="expandedNodeRef"
				@text-click="emits('textClick', $event)"
			>
				<treeNode
					@text-click="textClick($event, item)"
					:followTheme="props.followTheme"
					:fieldNames="props.fieldNames"
					:data="item"
				></treeNode>
			</expandedNode>
		</view>
	</view>
</template>

<script lang="ts" setup>
import treeNode from './tree-node.vue'
import expandedNode from './expanded-node.vue'
import { inject, computed, getCurrentInstance, ref } from 'vue'
const expandedNodeRef = ref<InstanceType<typeof expandedNode> | null>(null)

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
	}
})
const treeNodeData = computed(() => props.data)

function textClick(event: any, item: any) {
	emits('textClick', event)
	if (item && item['children']) {
		if (Array.isArray(expandedNodeRef.value)) {
			let el = expandedNodeRef.value.find((el) => el?.filedId == item[props.fieldNames.id])
			if (el) {
				el?.setStatus()
			}
		} else {
			expandedNodeRef.value?.setStatus()
		}
	}
}
</script>
