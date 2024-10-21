<template>
	<!-- #ifdef MP-WEIXIN -->
	<view class="flex relative" :class="[props.direction == 'horizontal' ? 'flex-row flex-row-top-center' : '']" ref="contentbody">
		<slot name="default"></slot>
	</view>
	<!-- #endif -->
	<!-- #ifndef MP-WEIXIN -->
	<view class="flex flex-row flex-row-center-center">
		<view
			class="flex relative"
			:style="{ width: widthComputed }"
			:class="[props.direction == 'horizontal' ? 'flex-row flex-row-top-center' : '']"
			ref="contentbody"
		>
			<slot name="default"></slot>
		</view>
	</view>
	<!-- #endif -->
</template>
<script lang="ts" setup>
/**
 * 步骤条
 * @description 步骤条，可以方便的预览逻辑，事务条理。需要配合tm-steps-item组件，必须放置tm-steps-item组件才能使用。
 */
import { number } from 'echarts'
import { computed, provide, ref, getCurrentInstance, watch, onUpdated, onMounted, nextTick, PropType } from 'vue'
// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
// #endif
const props = defineProps({
	/**
	 * 步骤条显示的方向。
	 * horizontal|vertical
	 */
	direction: {
		type: String as PropType<'horizontal' | 'vertical'>,
		default: 'horizontal'
	},
	//当前的步骤。可使用v-model:current
	current: {
		type: [Number],
		default: 0
	},
	//当前默认初始，显示的步骤数。
	defaultCurrent: {
		type: Number,
		default: 0
	},
	/**
	 * 当前步骤的状态。
	 * 'wait' | 'process' | 'finish' | 'error'
	 */
	status: {
		type: String as PropType<'wait' | 'process' | 'finish' | 'error'>,
		default: ''
	},
	//是否显示连线。
	showLine: {
		type: Boolean,
		default: true
	},
	//是否允许点击步骤来切换当前步骤数。
	changeable: {
		type: Boolean,
		default: false
	},
	/**
	 * 当点击步骤切换前的勾子函数，返回fase将取消切换。可以返回Promise异步。
	 */
	beforeStepChange: {
		type: [Function, Boolean],
		default: () => false
	},

	//这里提供，如果子组件tm-steps-item没有提供使用此值。如果子组件提供了，不会使用这里的参数。
	//未激活时的主题色
	color: {
		type: String,
		default: 'grey-3'
	},
	//激活时的主题色。这里提供，如果子组件tm-steps-item没有提供使用此值。如果子组件提供了，不会使用这里的参数。
	activeColor: {
		type: String,
		default: 'primary'
	},
	//样式dot,number
	type: {
		type: String,
		default: 'dot'
	}
})
/**
 * 事件说明
 * change 当前切换步骤时触发。
 * update:current 即v-model:current
 * step-click 当点击步骤时触发。
 */
const emits = defineEmits(['change', 'update:current', 'step-click'])
const proxy = getCurrentInstance()?.proxy ?? null
const _current = ref(props.defaultCurrent ?? -1)
const _countCurrent = ref(-1)
const _width = ref(0)
const widthComputed = computed(() => Math.ceil(uni.upx2px(_width.value)) + 'px')

const compoenentName = 'tmSteps'

function pushKey(width: number = 150) {
	_countCurrent.value += 1
	_width.value += width
	return _countCurrent.value
}
function delKey(width: number = 150) {
	_countCurrent.value -= 1
	_width.value -= width
	return _countCurrent.value
}
watch(
	() => props.current,
	() => {
		_current.value = props.current
		emits('change', _current.value)
	}
)

function steplick(index: number) {
	_current.value = index
	emits('step-click', index)
	emits('update:current', _current.value)
}

provide(
	'tmStepsCureent',
	computed(() => _current.value)
)
provide(
	'tmStepsCountCureent',
	computed(() => _countCurrent.value)
)
provide(
	'tmStepsCountActiveColor',
	computed(() => props.activeColor)
)
provide(
	'tmStepsCountColor',
	computed(() => props.color)
)

defineExpose({ pushKey, compoenentName, steplick, delKey })
</script>

<style></style>
