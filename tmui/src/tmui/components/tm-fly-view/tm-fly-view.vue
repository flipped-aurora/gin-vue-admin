<template>
	<view @click="onClick" id="flyViewBody" class="">
		<!-- #ifdef APP-NVUE -->
		<view :eventPenetrationEnabled="true">
			<slot></slot>
		</view>
		<!-- #endif -->
		<!-- #ifndef APP-NVUE -->
		<slot></slot>
		<!-- #endif -->
		<view
			id="flyView"
			ref="flyView"
			class="flex flex-row flex-row-center-center"
			:eventPenetrationEnabled="true"
			:style="{
				width: props.width + 'rpx',
				height: props.height + 'rpx',
				transform: `translate(${position.x}px,${position.y}px)`,
				zIndex: props.zIndex,
				position: 'fixed'
			}"
		>
			<view v-if="showEl"><slot name="content"></slot></view>
		</view>
	</view>
</template>

<script lang="ts" setup>
import { ref, inject, computed, unref, PropType, getCurrentInstance, onUnmounted, onMounted, reactive } from 'vue'
import * as TWEEN from '@/tmui/tool/lib/tween.min.js'
import { findXOnQuadraticBezierCurve } from './parabola'
const proxy = getCurrentInstance()?.proxy ?? null
// #ifdef APP-NVUE
// @ts-ignore
var dom = weex.requireModule('dom')
const Binding = uni.requireNativePlugin('bindingx')
const animation = uni.requireNativePlugin('animation')
// #endif
const props = defineProps({
	/** 开启吸附后，吸附到边缘的动画时间，单位ms */
	duration: {
		type: Number,
		default: 600
	},
	width: {
		type: Number,
		default: 100
	},
	height: {
		type: Number,
		default: 100
	},
	/** 目标位置的x坐标，位置从页面的左上角0算起 */
	x: {
		type: Number,
		default: 0
	},
	/** 目标位置的y坐标，位置从页面的左上角0算起 */
	y: {
		type: Number,
		default: 0
	},
	/** vue层级,nvue不起作用，nvue请在页面的底部放置，即可渲染在最前面。 */
	zIndex: {
		type: Number,
		default: 100
	}
})
const sysinfo = uni.getSystemInfoSync()
let bindxToken: any = null

const position = reactive({ x: 0, y: 0 })
let requestAnimationFrameId = 0
const _x = computed(() => props.x)
const _y = computed(() => props.y)
let ParabolaObj = null
let start_x = 0
let start_y = 0
let elWidth = 0
let elTop = 0
let elLeft = 0
let elHeight = 0
let showEl = ref(false)
let tween: TWEEN | null = null
onUnmounted(() => {
	destory()
})
onMounted(() => {
	animate()
})

function getEl(el: any) {
	if (typeof el === 'string' || typeof el === 'number') return el
	if (WXEnvironment) {
		return el.ref
	} else {
		return el instanceof HTMLElement ? el : el.$el
	}
}
function spinNvueAniEnd(start: number, end: number, duration = props.duration) {
	// #ifdef APP-NVUE
	// if (!proxy?.$refs?.adsorb) return;
	// animation.transition(
	// 	proxy?.$refs.adsorb,
	// 	{
	// 		styles: {
	// 			transform: `translate(${start}px,${end}px)`,
	// 			transformOrigin: "center center",
	// 		},
	// 		duration: duration, //ms
	// 		timingFunction: "cubicBezier(0.18, 0.89, 0.32, 1)",
	// 		delay: 0, //ms
	// 	},
	// 	() => {
	// 	}
	// );
	// #endif
}

function onClick(ev: TouchEvent | MouseEvent) {
	if (!proxy) {
		console.error('错误：没有获取到元素实例。')
		return
	}
	uni.$tm.u
		.queryDom(proxy, '#flyViewBody')
		.then((r) => {
			console.log(ev)
			elTop = r.top
			elLeft = r.left
			elWidth = r.width
			elHeight = r.height
			// #ifndef APP-NVUE
			start_x = ev.touches[0].clientX
			start_y = ev.touches[0].clientY
			// #endif
			// #ifdef APP-NVUE
			start_x = r.left + r.width / 2
			start_y = r.top + r.height / 2
			// #endif
			play()
		})
		.catch((e) => {
			console.error(e)
		})
}

function play() {
	let realHeight = Math.abs(_y.value - elTop - elHeight)
	let realWidth = Math.abs(_x.value - elLeft - elWidth)

	if (tween) {
		tween.stop()
		position.x = 0
		position.y = 0
	}
	showEl.value = true
	tween = new TWEEN.Tween({ progress: 0 })
		.easing(TWEEN.Easing.Linear.None) // 缓动函数
		.to({ progress: 1 }, props?.duration ?? 0)
		.onUpdate((e: any) => {
			let y = realHeight * e.progress

			let co_x = start_x
			let co_y = start_y
			let co_x2 = start_x
			let co_y2 = start_y
			let s_x = 0
			let s_y = 0
			let e_x = _x.value - elLeft
			let e_y = realHeight

			if (start_x < _x.value) {
				co_x = start_x
				co_y = 0
				co_x2 = _x.value
				co_y2 = 0
				s_x = elWidth / 2
			} else {
				co_x = -elWidth / 2
				co_y = 0
				co_x2 = -start_x + elWidth
				co_y2 = 0
				s_x = -elWidth / 2
				// #ifdef APP-NVUE
				s_x = start_x - elWidth / 2
				// #endif
			}
			const x = findXOnQuadraticBezierCurve(s_x, s_y, co_x, co_y, co_x2, co_y2, e_x, e_y, y)
			// position.y =  start_y + e.progress*(_y.value - start_y);
			// const x_tep = getXFromParabola(ParabolaObj, _y.value*e.progress);

			position.x = x
			position.y = y
		})
		.onStart((e: any) => {})
		.onComplete((e: any) => {
			showEl.value = false
		})
		.delay(0)
		// .repeat()
		// .yoyo()
		.start()
}

function animate() {
	function animateVC() {
		requestAnimationFrameId = uni.$tm.u.requestAnimationFrame(animateVC)
		TWEEN.update()
	}
	animateVC()
}
function destory() {
	uni.$tm.u.cancelAnimationFrame(requestAnimationFrameId)
}
defineExpose({})
</script>

<style scoped>
/* #ifndef APP-NVUE */
#flyView {
}

#flyViewBody {
	position: relative;
}
/* #endif */
</style>
