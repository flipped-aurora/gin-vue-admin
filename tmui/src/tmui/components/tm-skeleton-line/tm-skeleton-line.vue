<template>
	<!-- #ifndef APP-NVUE -->
	<view
		class="tmSkeletonLine flex-12 my-10"
		:class="[`round-${props.round}`]"
		:style="[
			{ backgroundColor: isDark ? '#1e1e1e' : '#ebebeb' },
			{ paddingTop: props.height / 2 + 'rpx', paddingBottom: props.height / 2 + 'rpx' }
		]"
	></view>
	<!-- #endif -->
	<!-- #ifdef APP-NVUE -->
	<view
		ref="dombg"
		class="tmSkeletonLine flex-12 my-10"
		:class="[`round-${props.round}`]"
		:style="[
			{ backgroundColor: isDark ? '#1e1e1e' : '#ebebeb' },
			{ paddingTop: props.height / 2 + 'rpx', paddingBottom: props.height / 2 + 'rpx' }
		]"
	></view>
	<!-- #endif -->
</template>

<script lang="ts" setup>
/**
 * 骨架屏-元素
 * @description 骨架加载状态屏元素，可以使用本元素组件自由的组合出你想要的任意加载造型，这对自定义骨架形状非常有用。
 */
import { getCurrentInstance, computed, onMounted, onUnmounted, nextTick, watch } from 'vue'
import { computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()
// #ifdef APP-PLUS-NVUE
const animation = uni.requireNativePlugin('animation')
// #endif
const props = defineProps({
	height: {
		type: Number,
		default: 60
	},
	dark: {
		type: [Boolean, String],
		default: false
	},
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	//是否跟随主题全局切换暗黑模式。
	followDark: {
		type: [Boolean, String],
		default: true
	},
	//圆角。
	round: {
		type: Number,
		default: 4
	}
})
const proxy = getCurrentInstance()?.proxy ?? null
// 设置响应式全局组件库配置表。
const tmcfg = computed(() => store.tmStore)

const isDark = computed(() => computedDark(props, tmcfg.value))

onMounted(() => {
	// #ifdef APP-PLUS-NVUE
	try {
		nextTick(function () {
			setTimeout(function () {
				spinNvueAni()
			}, 50)
		})
	} catch (e) {
		//TODO handle the exception
	}
	// #endif
})

// function getEl(el) {
// 	if (typeof el === "string" || typeof el === "number") return el;
// 	if (WXEnvironment) {
// 		return el.ref;
// 	} else {
// 		return el instanceof HTMLElement ? el : el.$el;
// 	}
// }

function spinNvueAni(opacity = 0) {
	let icon = proxy?.$refs?.dombg
	if (!icon) return
	animation.transition(
		icon,
		{
			styles: {
				opacity: opacity
			},
			duration: 1000, //ms
			timingFunction: 'linear',
			delay: 0 //ms
		},
		() => {
			nextTick(function () {
				spinNvueAni(opacity == 0 ? 1 : 0)
			})
		}
	)
}
</script>

<style scoped>
/* #ifndef APP-NVUE */
.tmSkeletonLine {
	animation: loading 1s linear infinite;
}

@keyframes loading {
	0% {
		opacity: 0;
	}

	50% {
		opacity: 1;
	}

	100% {
		opacity: 0;
	}
}

/* #endif */
</style>
