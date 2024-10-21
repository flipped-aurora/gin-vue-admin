<template>
	<view
		ref="bodywk"
		@click="hanlder"
		:class="[customClass, 'overflow']"
		:style="[computedHeight ? { height: computedHeight } : '', computedWidth ? { width: computedWidth } : '', customCSSStyle]"
	>
		<!-- #ifdef APP-NVUE -->
		<view
			v-if="isLoadEl"
			ref="nvueElAni"
			:animation="animationData"
			:class="['flex-col flex trani', animationName + reverseAniPrefxname, customClass]"
		>
			<slot name="default"></slot>
		</view>
		<!-- #endif -->
		<!-- #ifndef APP-NVUE -->
		<view
			v-if="isLoadEl"
			ref="nvueElAni"
			:style="{
				transitionDuration: `${props.duration}ms`,
				transitionTimingFunction: 'ease'
			}"
			:class="['flex-col flex trani', animationClassName, customClass]"
		>
			<slot name="default"></slot>
		</view>
		<!-- #endif -->
	</view>
</template>

<script lang="ts" setup>
/**
 * 动画
 * @description 提供了6组动画，并且每组件动画都支持反转播放，相当于12组动画了。可用于任何元素，进入场和出场的动画。
 */
import { getCurrentInstance, computed, ref, provide, inject, onMounted, onUnmounted, nextTick, watch } from 'vue'
import { cssstyle, tmVuetify } from '../../tool/lib/interface'
import { custom_props, computedClass, computedStyle } from '../../tool/lib/minxs'
// #ifdef APP-PLUS-NVUE
const Binding = uni.requireNativePlugin('bindingx')
const dom = uni.requireNativePlugin('dom')
const animation = uni.requireNativePlugin('animation')
// #endif
// 混淆props共有参数
const props = defineProps({
	...custom_props,
	duration: {
		type: Number,
		default: 300
	},
	delay: {
		type: Number,
		default: 0
	},
	//动画名称
	name: {
		type: String,
		default: 'fade' //fade,left,right,up,down,zoom
	},
	autoPlay: {
		type: Boolean,
		default: true
	},
	disabled: {
		type: Boolean,
		default: false
	},
	height: {
		type: [Number, String],
		default: 0
	},
	width: {
		type: [Number, String],
		default: 0
	},
	//是否反向动画
	reverse: {
		type: [Boolean, String],
		default: false
	},
	//每变动一次，就重置动画一下，这个属性不对外，特殊情况使用。
	initByWechat: {
		type: Boolean,
		default: false
	}
})
const emits = defineEmits(['start', 'end', 'click'])

function hanlder(e: any) {
	emits('click', e)
}
const proxy = getCurrentInstance()?.proxy ?? null
//自定义样式：
const customCSSStyle = computed(() => computedStyle(props))
//自定类
const customClass = computed(() => computedClass(props))
//高度，
const computedHeight = computed(() => {
	if (!props.height || !Number(props.height)) {
		return 0
	}
	if (String(props.height).indexOf('px') > -1 || String(props.height).indexOf('rpx') > -1) {
		return String(props.height)
	}
	return String(props.height) + 'rpx'
})

// 宽度
const computedWidth = computed(() => {
	if (!props.width) {
		return 0
	}
	if (String(props.width).indexOf('px') > -1 || String(props.width).indexOf('rpx') > -1) {
		return props.width
	}
	return props.width + 'rpx'
})
//动画名称
const animationName = computed(() => props.name || 'fade')
//动画时长
const durationtos = computed(() => props.duration)
//是否反转。
const computedReverse = computed(() => props.reverse)
//反转动画前缀
const reverseAniPrefxname = computed(() => (computedReverse.value ? '-reverse' : ''))
//非nvue
const animationClassName = ref(animationName.value + reverseAniPrefxname.value)
//动画播放状态。0未开始,结束播放，1开始播放中,2结束。
const animationStatus = ref(0)
const tmid = ref(Number(uni.$tm.u.getUid(3)))
//是否渲染完成。
const isLoadEl = ref(false)
//css3动画数据。
const animationData = ref(null)
watch([() => props.initByWechat, () => props.name], () => {
	reset()
})
watch([() => props.name], () => {
	animationClassName.value = animationName.value + reverseAniPrefxname.value
})
function init() {
	nextTick(() => {
		isLoadEl.value = true
		if (props.autoPlay == true && !props.disabled) {
			nextTick(() => play())
		}
	})
}

function play() {
	if (props.disabled == true) return
	animationStatus.value = 0
	// #ifdef APP-NVUE
	clearTimeout(tmid.value)
	nextTick(function () {
		tmid.value = setTimeout(function () {
			emits('start')
			nvueAmatons()
		}, 50)
	})
	// #endif
	// #ifndef APP-NVUE
	noNvueAmations()
	// #endif
}

function stop() {
	if (props.disabled == true) return
	clearTimeout(tmid.value)
	animationStatus.value = 0
	// noNvueAmationsReset()
}

function reset() {
	stop()

	animationStatus.value = 0
}
//对外暴露方法。
defineExpose({
	init: init,
	play: play,
	stop: stop,
	reset: reset
})
//获取当前播放状态。
function getPlayStatus() {
	return animationStatus.value
}

onMounted(() => init())
onUnmounted(() => {
	clearTimeout(tmid.value)
	animationStatus.value = 0
})

function nvueAmatons() {
	var el = proxy.$refs.nvueElAni
	let propsAni = {}

	if (animationName.value == 'fade') {
		propsAni = {
			opacity: computedReverse.value ? 0 : 1,
			transformOrigin: 'center center'
		}
	} else if (animationName.value == 'up') {
		propsAni = {
			opacity: 1,
			transform: computedReverse.value ? 'translateY(0%)' : 'translateY(-100%)',
			transformOrigin: 'center center'
		}
	} else if (animationName.value == 'down') {
		propsAni = {
			opacity: 1,
			transform: computedReverse.value ? 'translateY(0%)' : 'translateY(100%)',
			transformOrigin: 'center center'
		}
	} else if (animationName.value == 'right') {
		propsAni = {
			opacity: 1,
			transform: computedReverse.value ? 'translateX(0%)' : 'translateX(100%)',
			transformOrigin: 'center center'
		}
	} else if (animationName.value == 'left') {
		propsAni = {
			opacity: 1,
			transform: computedReverse.value ? 'translateX(0%)' : 'translateX(-100%)',
			transformOrigin: 'center center'
		}
	} else if (animationName.value == 'zoom') {
		propsAni = {
			opacity: computedReverse.value ? 0 : 1,
			transform: computedReverse.value ? 'scale(0.7,0.7)' : 'scale(1,1)',
			transformOrigin: 'center center'
		}
	}
	emits('start')
	animationStatus.value = 1
	clearTimeout(tmid.value)
	tmid.value = setTimeout(function () {
		animation.transition(
			el,
			{
				styles: propsAni,
				duration: props.duration, //ms
				timingFunction: 'ease',
				delay: 0 //ms
			},
			() => {
				emits('end')
				animationStatus.value = 2
			}
		)
	}, 20)
}
function noNvueAmations() {
	clearTimeout(tmid.value)
	tmid.value = setTimeout(() => {
		if (computedReverse.value) {
			animationClassName.value = animationName.value
		} else {
			animationClassName.value = animationName.value + '-reverse'
		}
		tmid.value = setTimeout(() => {
			emits('end')
		}, props.duration)
	}, 20)
}
</script>

<style scoped>
.fade {
	opacity: 0;
}
.fade-reverse {
	opacity: 1;
}

.up {
	transform: translateY(0%);
}

.up-reverse {
	transform: translateY(-101%);
}

.down {
	transform: translateY(0%);
}

.down-reverse {
	transform: translateY(101%);
}

.left {
	transform: translateX(0%);
}

.left-reverse {
	transform: translateX(-101%);
}

.right {
	transform: translateX(0%);
}

.right-reverse {
	transform: translateX(101%);
}

.zoom {
	transform: scale(0.7, 0.7);
	opacity: 0;
}

.zoom-reverse {
	transform: scale(1, 1);
	opacity: 1;
}
</style>
