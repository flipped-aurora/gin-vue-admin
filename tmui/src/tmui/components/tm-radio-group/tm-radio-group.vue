<template>
	<view
		class="flex"
		:class="[props.direction == 'row' ? 'flex-row' : 'flex-col', props.direction == 'customCol' ? '' : _align]"
		:style="{ flexWrap: props.direction == 'customCol' ? 'nowrap' : 'wrap' }"
	>
		<slot></slot>
	</view>
</template>
<script lang="ts" setup>
/**
 * 单选框组
 * @description 单选框组中，只能放置tm-radio组件，且必须配合tm-radio组件一起使用，不可单独使用。
 */
import { computed, nextTick, provide, ref, watch, getCurrentInstance, inject, toRaw, PropType } from 'vue'
import { inputPushItem, rulesItem } from './../tm-form-item/interface'
const emits = defineEmits(['update:modelValue', 'change'])
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	disabled: {
		type: Boolean,
		default: false
	},
	defaultValue: {
		type: [String, Number, Boolean,null],
		default: ''
	},
	modelValue: {
		type: [String, Number, Boolean,null],
		default: ''
	},
	direction: {
		type: String as PropType<'row' | 'col' | 'customCol'>,
		default: 'row' //row横排，col为竖排。
	},
	align: {
		type: String as PropType<'left' | 'center' | 'right'>,
		default: 'left'
	},
	//模式
	/**
	 * radio:正常的单选样式。
	 * button:按钮单选模式
	 */
	model: {
		type: String,
		default: 'radio' // radio,button
	}
})
let _cacheBoxList: Array<string | number | boolean> = []
const _mValue = ref(props.defaultValue || props.modelValue)

const _align = computed(() => {
	let list = {
		left: 'flex-row-center-start',
		center: 'flex-row-center-center',
		right: 'flex-row-center-end'
	}
	let listCol = {
		left: 'flex-col-center-start',
		center: 'flex-col-center-center',
		right: 'flex-col-center-end'
	}
	return props.direction == 'row' ? list[props.align] : listCol[props.align]
})

//组件唯一标识。
const checkBoxkeyId = 'tmRadioBoxGroup'
watch(
	() => props.modelValue,
	() => {
		_mValue.value = props.modelValue
	},
	{ deep: true }
)
function pushKey(key: string | number | boolean) {
	_cacheBoxList.push(key)
}
nextTick(() => {
	const _filter_key = _cacheBoxList.filter((el) => el == _mValue.value)
	if (_filter_key.length > 0) {
		_mValue.value = _filter_key[0]
	}
	emits('update:modelValue', _mValue.value)
})
//更新值、
function addKey(key: string | number | boolean) {
	_mValue.value = key
	emits('update:modelValue', _mValue.value)
	// console.log(new Date().getSeconds(),'----3')

	nextTick(() => {
		emits('change', _mValue.value)
	})
}

provide(
	'tmRadioBoxDisabled',
	computed(() => props.disabled)
)
provide('tmRadioBoxVal', _mValue)
provide(
	'tmRadioBoxModel',
	computed(() => props.model == 'radio')
)
provide('tmCheckedBoxDir', props.direction)
defineExpose({ pushKey: pushKey, addKey: addKey, checkBoxkeyId: checkBoxkeyId })
</script>
