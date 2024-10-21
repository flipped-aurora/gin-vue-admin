<template>
	<!-- v-if="tmCascaderShowIndex==_level" -->
	<view>
		<scroll-view v-if="tmCascaderShowIndex == props.level" :style="[{ height: `${props.height}rpx` }]" scroll-y>
			<view v-for="(item, index) in _value" :key="index">
				<nodeCellVue
					:followTheme="props.followTheme"
					:level="props.level"
					:data="item"
					:keyIndex="index"
					@click="onCellClick(item)"
				></nodeCellVue>
			</view>
		</scroll-view>
		<BaseCascader
			:followTheme="props.followTheme"
			:level="_level"
			v-if="nextChildData.length > 0"
			:height="props.height"
			:color="props.color"
			:data="nextChildData"
			@cellClick="onCellClick"
		></BaseCascader>
	</view>
</template>
<script lang="ts" setup>
import { computed, ref, Ref, watch, getCurrentInstance, inject, toRaw, watchEffect, PropType } from 'vue'
import BaseCascader from './base-cascader.vue'
import BaseNode from './base-node.vue'
import nodeCellVue from './node-cell.vue'
import { childrenData } from './interface'
const emits = defineEmits(['cellClick'])
const props = defineProps({
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	/**
	 * 导入的数据
	 */
	data: {
		type: Array as PropType<Array<childrenData>>,
		default: () => [],
		required: true
	},
	height: {
		type: Number,
		default: 0,
		required: true
	},
	color: {
		type: String,
		default: 'primary'
	},
	level: {
		type: Number,
		default: 0
	},
	//点击列表中项目时再自动切换到下一项时之前执行的勾子函数，方便动态加载数据。
	beforeCellClick: {
		type: [Function, Boolean],
		default: () => {
			return false
		}
	}
})
const _value = computed(() => props.data)
const _level = props.level + 1
const tmCascaderShowIndex = inject(
	'tmCascaderShowIndex',
	computed(() => 0)
)
const ParentActivedLs = inject(
	'tmCascaderValue',
	computed(() => [])
)
const nextChildData = ref([])
watchEffect(async () => {
	if (ParentActivedLs.value[props.level]) {
		let nowobj = _value.value.filter((el) => el.id == ParentActivedLs.value[props.level])
		if (nowobj.length > 0) {
			if (
				typeof nowobj[0]?.children == 'object' &&
				Array.isArray(nowobj[0]?.children) &&
				nowobj[0]?.children?.length > 0 &&
				nowobj[0]?.children.length >= 0
			) {
				nextChildData.value = nowobj[0]?.children
			}
		}
	}
})

function onCellClick(item: any) {
	emits('cellClick', item)
}
</script>
