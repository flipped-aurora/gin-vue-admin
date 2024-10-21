<template>
	<view class="col" style="box-sizing: border-box">
		<slot></slot>
	</view>
</template>

<script lang="ts" setup>
import { watchEffect, computed, ref, provide,  Ref } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
const emits = defineEmits(['change', 'update:active-key'])
const props = defineProps({
	//当前展开并激活的面板。
	activeKey: {
		type: [Array],
		default: () => []
	},
	//默认展开的面板
	defaultActiveKey: {
		type: [Array],
		default: () => []
	},
	//是否设置为单个面板展开，默认fase，允许 多个面板同时展开。
	accordion: {
		type: [Boolean, String],
		default: false
	},
	border: {
		type: [Number, String],
		default: 2
	},
	//项目展开和关闭图标的位置，可选left,right，默认是左left.
	iconPos: {
		type: String,
		default: 'left'
	},
	//展开后的图标
	openIcon: {
		type: String,
		default: 'tmicon-angle-up'
	},
	//关闭后的图标
	closeIcon: {
		type: String,
		default: 'tmicon-angle-down'
	}
})

const _activeKey:any = ref([])
watchEffect(() => {
	_activeKey.value = [...props.activeKey, ...props.defaultActiveKey]
})

if (props.accordion) {
	if (_activeKey.value.length > 0) {
		_activeKey.value = [_activeKey.value[0]]
	}
}
const cacheKey: Ref<Array<string | number>> = ref([])
const pushKey = function (key: string | number) {
	cacheKey.value = [...new Set([...cacheKey.value, key])]
}
const setKey = function (key: string | number) {
	let findkey = _activeKey.value.findIndex((el) => String(el) == String(key))
	if (props.accordion) {
		if (findkey > -1) {
			_activeKey.value = []
		} else {
			_activeKey.value = [key]
		}
	} else {
		if (findkey > -1) {
			_activeKey.value.splice(findkey, 1)
		} else {
			_activeKey.value.push(key)
		}
	}

	emits('update:active-key', _activeKey.value)
	emits('change', _activeKey.value)
}
emits('update:active-key', _activeKey.value)
defineExpose({
	tmCollapse: 'tmCollapse',
	setKey: setKey,
	pushKey: pushKey,
	border: props.border
})
provide(
	'tmCollapseKeyList',
	computed(() => _activeKey.value)
)
provide(
	'tmCollapseIconPos',
	computed(() => props.iconPos)
)
provide(
	'tmCollapseopenIcon',
	computed(() => props.openIcon)
)
provide(
	'tmCollapsecloseIcon',
	computed(() => props.closeIcon)
)
</script>

<style scoped>
.col {
	/* #ifndef APP-NVUE */
	display: flex;
	flex-direction: column;
	flex-shrink: 0;
	flex-grow: 0;
	flex-basis: auto;
	align-items: stretch;
	align-content: flex-start;
	/* #endif */
}
</style>
