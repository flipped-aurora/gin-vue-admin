<template>
	<!-- #ifndef APP-NVUE -->

	<view >
		<view
			v-if="!_disabled"
			@touchstart="startDrag"
			@touchmove="onDrag"
			@touchend="endDrag"
			@touchcancel="endDrag"
			:style="`width:${attr.width}px;height:${attr.height}px `"
			class="overflow relative"
			:class="[attr.disabled ? 'opacity-7' : '']"
		>
			<view class="flex flex-row flex-row-center-between">
				<view id="left" :style="{ width: `${leftWidth}px`, height: `${attr.height}px` }">
					<slot name="left"></slot>
				</view>
				<view id="right" :style="{ width: `${rightWidth}px`, height: `${attr.height}px` }">
					<slot name="right"></slot>
				</view>
			</view>
			<view id="wrapper" class="absolute l-0 t-0" :style="[`width:${attr.width}px;height:${attr.height}px;`, viewStyle]">
				<tm-sheet
					@click="emits('click')"
					:shadow="0"
					:outlined="props.outlined"
					:borderStyle="props.borderStyle"
					unit="px"
					:borderDirection="props.borderDirection"
					:linearDeep="props.linearDeep"
					:linear="props.linear"
					:round="props.round"
					:color="props.color"
					:text="_disabled"
					:transprent="props.transprent"
					:width="attr.width"
					:height="attr.height"
					:margin="[0, 0]"
					:padding="[0, 0]"
					:userInteractionEnabled="false"
				>
					<slot></slot>
				</tm-sheet>
			</view>
		</view>

		<view
			v-if="_disabled"
			:style="`width:${attr.width}px;height:${attr.height}px `"
			class="overflow relative"
			:class="[attr.disabled ? 'opacity-7' : '']"
		>
			<view class="flex flex-row flex-row-center-between">
				<view id="left" :style="{ width: `${leftWidth}px`, height: `${attr.height}px` }">
					<slot name="left"></slot>
				</view>
				<view id="right" :style="{ width: `${rightWidth}px`, height: `${attr.height}px` }">
					<slot name="right"></slot>
				</view>
			</view>

			<view id="wrapper" class="absolute l-0 t-0" :style="[`width:${attr.width}px;height:${attr.height}px `]">
				<tm-sheet
					@click="emits('click')"
					:shadow="0"
					:outlined="props.outlined"
					:borderStyle="props.borderStyle"
					unit="px"
					:borderDirection="props.borderDirection"
					:linearDeep="props.linearDeep"
					:linear="props.linear"
					:round="props.round"
					:color="props.color"
					:text="_disabled"
					:transprent="props.transprent"
					:width="attr.width"
					:height="attr.height"
					:margin="[0, 0]"
					:padding="[0, 0]"
				>
					<slot></slot>
				</tm-sheet>
			</view>
		</view>
	</view>
	<!-- #endif -->
	<!-- #ifdef APP-NVUE -->
	<view :style="`width:${attr.width}px;height:${attr.height}px `" class="overflow relative" :class="[attr.disabled ? 'opacity-7' : '']">
		<view class="flex flex-row flex-row-center-between">
			<view id="left" :style="{ width: `${leftWidth}px`, height: `${attr.height}px` }">
				<slot name="left"></slot>
			</view>
			<view id="right" :style="{ width: `${rightWidth}px`, height: `${attr.height}px` }">
				<slot name="right"></slot>
			</view>
		</view>
		<view
			@click="emits('click')"
			@touchstart.stop="touchstart"
			id="wrapper"
			ref="tabsDom"
			class="absolute l-0 t-0"
			:style="`width:${attr.width}px;height:${attr.height}px;transform:${opened ? '' : 'translate3d( 0px, 0, 0)'}`"
		>
			<tm-sheet
				:eventPenetrationEnabled="true"
				:shadow="0"
				:outlined="props.outlined"
				:borderStyle="props.borderStyle"
				unit="px"
				:borderDirection="props.borderDirection"
				:linearDeep="props.linearDeep"
				:linear="props.linear"
				:round="props.round"
				:color="props.color"
				:text="_disabled"
				:transprent="props.transprent"
				:width="attr.width"
				:height="attr.height"
				:margin="[0, 0]"
				:padding="[0, 0]"
			>
				<slot></slot>
			</tm-sheet>
		</view>
	</view>
	<!-- #endif -->
</template>

<script lang="ts" setup>
/**
 * 左滑操作栏
 * @description  向左滑动显示底部操作按钮栏。
 */
import { computed, nextTick, onMounted, ref, getCurrentInstance, reactive, watch } from 'vue'
import { custom_props } from '../../tool/lib/minxs'
import { defaultProps } from './props'
import tmSheet from '../tm-sheet/tm-sheet.vue'

// @ts-ignore
// #ifdef APP-NVUE
var dom = uni.requireNativePlugin('dom')
const Binding = uni.requireNativePlugin('bindingx')
const animation = uni.requireNativePlugin('animation')
// #endif
const proxy = getCurrentInstance()?.proxy ?? null
const bindxToken = ref(null)
const props = defineProps({
	...custom_props,
	...defaultProps
})
const emits = defineEmits(['click', 'open', 'close', 'update:open-status'])
const viewStyle = ref({})

const _disabled = ref(props.disabled)
const opened = ref(false)
const closed = ref(false)

const leftWidth = computed(() => props.leftWidth)
const rightWidth = computed(() => props.rightWidth)
const attr = computed(() => {
	return {
		width: Math.ceil(uni.$tm.u.topx(props.width)),
		height: Math.ceil(uni.$tm.u.topx(props.height)),
		disabled: props.disabled,
		leftWidth: Math.ceil(uni.$tm.u.topx(props.leftWidth)),
		rightWidth: Math.ceil(uni.$tm.u.topx(props.rightWidth))
	}
})

const THRESHOLD = 0.3
const MIN_DISTANCE = 10
const state = reactive({
	leftWidth: Math.ceil(uni.$tm.u.topx(props.leftWidth)),
	rightWidth: Math.ceil(uni.$tm.u.topx(props.rightWidth)),
	offset: 0,
	startOffset: 0,
	dragging: false,
	startX: 0,
	startY: 0,
	direction: '',
	deltaX: 0,
	deltaY: 0,
	offsetX: 0,
	offsetY: 0
})

// nvue bingx
let nvue_now_left = 0
function getEl(el: HTMLElement) {
	if (typeof el === 'string' || typeof el === 'number') return el
	if (WXEnvironment) {
		return el.ref
	} else {
		return el instanceof HTMLElement ? el : el.$el
	}
}
function spinNvueAniEnd(start: number, end: number, isEnd = false, duration = 300) {
	// #ifdef APP-NVUE
	if (!proxy.$refs?.tabsDom) return
	animation.transition(
		proxy.$refs.tabsDom,
		{
			styles: {
				transform: `translateX(${start + end}px)`,
				transformOrigin: 'center center'
			},
			duration: duration, //ms
			timingFunction: 'cubicBezier(0.18, 0.89, 0.32, 1)',
			delay: 0 //ms
		},
		() => {
			if (isEnd) {
				funMethod('close', true)
			} else {
				funMethod('open', true)
			}
		}
	)

	// #endif
}

function touchstart(e: TouchEvent) {
	if (_disabled.value) return
	// #ifdef APP-NVUE
	if (!proxy.$refs?.tabsDom) return

	let icon = getEl(proxy.$refs.tabsDom)
	let expression = ``
	if (nvue_now_left < 0) {
		expression = `x<120&&x>=0?x-120:(x<0?-x:0)`
	} else {
		expression = `(x<=0&&x>-120)?x+0:-x`
	}
	let icon_bind = Binding.bind(
		{
			anchor: icon,
			eventType: 'pan',
			props: [
				{
					element: icon,
					property: 'transform.translateX',
					expression: expression
				}
			]
		},
		function (res) {
			if (res.state == 'end') {
				let lx = Math.abs(res.deltaX)
				let left = res.deltaX >= 0 ? false : true

				if (nvue_now_left == -attr.value.rightWidth) {
					if (res.deltaX == 0) {
						spinNvueAniEnd(res.deltaX, 0)
					} else if (res.deltaX < 0) {
						spinNvueAniEnd(res.deltaX, lx)
					} else {
						spinNvueAniEnd(res.deltaX, -lx)
					}
					opened.value = false
					nvue_now_left = 0
				} else {
					if (lx > 30 && left) {
						spinNvueAniEnd(res.deltaX, left ? -(attr.value.rightWidth - lx) : 0, true)
						opened.value = true
						nvue_now_left = -attr.value.rightWidth
					} else {
						spinNvueAniEnd(res.deltaX, -res.deltaX)
						opened.value = false
						nvue_now_left = 0
					}
				}
			} else if (res.state == 'start') {
				// isMoveing.value = true
			}
		}
	)
	bindxToken.value = icon_bind.token
	// #endif
}

onMounted(() => {
	opened.value = props.openStatus
	// #ifdef APP-NVUE
	nvue_now_left = -attr.value.rightWidth
	nextTick(() => {
		spinNvueAniEnd(0, -attr.value.rightWidth)
	})
	// #endif
	// #ifndef APP-NVUE
	setTimeout(() => {
		initOpen()
	}, 100)
	// #endif
})

watch(
	() => props.openStatus,
	(newVal: boolean, oldVal: boolean) => {
		if (!newVal) {
			close()
		} else {
			open('right',false)
		}
	}
)

function funMethod(type: 'style' | 'closeOther' | 'open' | 'close', arg: any) {
	if (type == 'style') {
		viewStyle.value = arg
	} else if (type == 'closeOther') {
	} else if (type == 'open') {
		emits('update:open-status', true)
		if(arg){
			emits('open', true)
		}
	} else if (type == 'close') {
		emits('update:open-status', false)
		emits('close', false)
	}
}

var initOpen = function () {
	// opened为boolen类型，判断默认打开
	if (opened.value && state.rightWidth > 0) {
		swipeMove(-state.rightWidth)
	} else if (opened.value && state.leftWidth > 0) {
		swipeMove(state.leftWidth)
	}
	
}
var range = function (num: number, min: number, max: number) {
	return Math.min(Math.max(num, min), max)
}
var swipeMove = function (_offset: number) {
	if (_offset === undefined) _offset = 0

	state.offset = range(_offset, -state.rightWidth, 0)
	var transform = 'translate3d(' + state.offset + 'px, 0, 0)'
	var transition = state.dragging ? 'none' : 'transform .6s cubic-bezier(0.18, 0.89, 0.32, 1)'
	var style = {
		'-webkit-transform': transform,
		'-webkit-transition': transition,
		transform: transform,
		transition: transition
	}
	funMethod('style', style)
	
}

var close = function () {
	opened.value = false
	swipeMove(0)
	
	// #ifdef APP-NVUE
	
	nvue_now_left = -attr.value.rightWidth
	spinNvueAniEnd(-attr.value.rightWidth,attr.value.rightWidth,true)
	funMethod('close', false)
	// #endif
	// #ifndef APP-NVUE
	
	funMethod('close', false)
	// #endif
}
var open = function (position: 'left' | 'right',noEmits = true) {
	var _offset = position === 'left' ? +state.leftWidth : -state.rightWidth
	opened.value = true
	swipeMove(_offset)
	funMethod('open', noEmits)
}

var getDirection = function (x, y) {
	if (x > y && x > MIN_DISTANCE) {
		return 'horizontal'
	}
	if (y > x && y > MIN_DISTANCE) {
		return 'vertical'
	}
	return ''
}

var resetTouchStatus = function () {
	state.direction = ''
	state.deltaX = 0
	state.deltaY = 0
	state.offsetX = 0
	state.offsetY = 0
}

const startDrag = (event: TouchEvent | MouseEvent) => {
	try {
		event?.preventDefault();
		event?.stopImmediatePropagation();
	} catch (error) {
		
	}
	resetTouchStatus()
	state.startOffset = state.offset
	var touchPoint = event.touches[0]
	state.startX = touchPoint.clientX
	state.startY = touchPoint.clientY
	funMethod('closeOther', false)
}

const onDrag = (event: TouchEvent | MouseEvent) => {
	try {
		event?.preventDefault();
		event?.stopImmediatePropagation();
	} catch (error) {
		
	}
	
	var touchPoint = event.touches[0]
	state.deltaX = touchPoint.clientX - state.startX
	state.deltaY = touchPoint.clientY - state.startY
	state.offsetX = Math.abs(state.deltaX)
	state.offsetY = Math.abs(state.deltaY)
	state.direction = state.direction || getDirection(state.offsetX, state.offsetY)

	if (state.direction !== 'horizontal') {
		return
	}
	state.dragging = true
	swipeMove(state.startOffset + state.deltaX)
}
const endDrag = (event: TouchEvent | MouseEvent) => {
	state.dragging = false
	if (+state.rightWidth > 0 && -state.startOffset < +state.rightWidth && -state.offset > +state.rightWidth * THRESHOLD) {
		open('right')
		
	} else if (+state.leftWidth > 0 && state.startOffset < +state.leftWidth && state.offset > +state.leftWidth * THRESHOLD) {
		open('left')
	} else {
		// 仅在有发生侧滑的情况下自动关闭（由js控制是否异步关闭）
		if (state.startOffset !== state.offset) {
			close()
		}
	}
}

defineExpose({ close, open })
</script>

<style></style>
