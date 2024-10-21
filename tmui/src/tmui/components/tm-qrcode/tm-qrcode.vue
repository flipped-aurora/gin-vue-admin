<template>
	<view :style="{ width: `${_width}rpx`, height: `${_height}rpx` }">
		<!-- #ifdef APP-NVUE -->
		<gcanvas v-if="show" :id="canvasId" :ref="canvasId" class="canvas" :style="{ width: `${_width}rpx`, height: `${_height}rpx` }"> </gcanvas>
		<!-- #endif -->
		<!-- #ifdef MP-WEIXIN || MP-ALIPAY || MP-QQ -->
		<canvas type="2d" id="canvasId" canvas-id="canvasId" class="canvas" :style="{ width: `${_width}rpx`, height: `${_height}rpx` }"></canvas>
		<!-- #endif -->
		<!-- #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ || APP-NVUE -->
		<canvas :id="canvasId" :canvas-id="canvasId" class="canvas" :style="{ width: `${_width}rpx`, height: `${_height}rpx` }"></canvas>
		<!-- #endif -->
	</view>
</template>

<script lang="ts" setup>
/**
 * 二维码
 * @description 这是一个二维码组件，属性多，可以生成非常个性化的组件哦~
 * 如果想知道生成的属性请查看：qrOpts类型属性。
 * 更改任意属性，都将会导致重绘
 */
import { getCurrentInstance, computed, ref, PropType, inject, onUpdated, onMounted, onUnmounted, nextTick, watch } from 'vue'
import { qrOpts, qrOptsDefault } from './interface'
import CanvasRenderingContext2D from '../../tool/gcanvas/context-2d/RenderingContext.js'
import { qr } from './drawing'
// #ifdef APP-NVUE
import { enable, WeexBridge } from '../../tool/gcanvas/index.js'
// #endif
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	option: {
		type: Object as PropType<qrOpts>,
		default: () => {
			return qrOptsDefault
		}
	}
})
const vnodeCtx = proxy
const canvasId = ref('canvasId')
// #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ
canvasId.value = 'tm' + uni.$tm.u.getUid(5)
// #endif
let ctx: CanvasRenderingContext2D
let canvas2d: HTMLCanvasElement
let opts = computed(() => {
	return { ...qrOptsDefault, ...props.option }
})
const _width = computed(() => opts.value.size)
const _height = computed(() => opts.value.size)
const show = ref(false) //安卓上首次要隐藏不然卡。
let isAndroid = false
// #ifdef APP-NVUE
isAndroid = uni.getSystemInfoSync().osName == 'android'
// #endif
onMounted(() => {
	nextTick(async function () {
		// #ifdef APP-NVUE
		if (isAndroid) {
			setTimeout(() => {
				show.value = true
				// 不要问我为什么安卓要绘制两次。问uni，才知道原理。
				init().then(() => qr(ctx, { ...opts.value, size: uni.upx2px(_width.value) }, canvas2d))
				setTimeout(function () {
					init().then(() => qr(ctx, { ...opts.value, size: uni.upx2px(_width.value) }, canvas2d))
				}, 50)
			}, 200)
		} else {
			show.value = true
			init().then(() => qr(ctx, { ...opts.value, size: uni.upx2px(_width.value) }, canvas2d))
		}
		// #endif

		// #ifndef APP-NVUE
		init().then(() => qr(ctx, { ...opts.value, size: uni.upx2px(_width.value) }, canvas2d))
		// #endif
	})
})

watch(
	() => props.option,
	() => {
		if (!ctx) {
		} else {
			qr(ctx, { ...opts.value, size: uni.upx2px(_width.value) }, canvas2d)
		}
	},
	{ deep: true }
)
function init() {
	return new Promise((res, rej) => {
		setTimeout(async function () {
			// #ifdef APP-NVUE
			ctx = await drawNvue_init()
			// #endif
			// #ifdef MP-WEIXIN || MP-ALIPAY || MP-QQ
			const { ctx2d, canvas } = await MpWeix_init()
			ctx = ctx2d
			canvas2d = canvas
			// #endif
			// #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ || APP-NVUE
			ctx = await appvueH5Other()
			// #endif
			res(true)
		}, 150)
	})
}
//appvue,h5,和其它平台。
function appvueH5Other() {
	return Promise.resolve(uni.createCanvasContext(canvasId.value, vnodeCtx))
}
//支付宝和微信，QQ 支持2d渲染。
function MpWeix_init() {
	return new Promise((resolve, rej) => {
		const query = uni.createSelectorQuery().in(vnodeCtx)
		// #ifdef MP-ALIPAY
		query
			.select('#canvasId')
			.node()
			.exec((res2) => {
				const canvas = res2[0].node
				let ctxvb: UniApp.CanvasContext = canvas.getContext('2d')

				resolve({ ctx2d: ctxvb, canvas: canvas })
			})
		// #endif
		// #ifdef MP-WEIXIN || MP-QQ
		query
			.select('#canvasId')
			.fields({
				node: true,
				size: true,
				context: true
			})
			.exec((res) => {
				// #ifdef MP-WEIXIN
				let canvas: any = res[0]?.node
				let ctxvb: UniApp.CanvasContext = canvas.getContext('2d')
				const dpr = uni.getSystemInfoSync().pixelRatio
				canvas.width = res[0].width * dpr
				canvas.height = res[0].height * dpr
				ctxvb.scale(dpr, dpr)
				resolve({ ctx2d: ctxvb, canvas: canvas })
				// #endif

				// #ifdef MP-QQ
				resolve({ ctx2d: res[0].context, canvas: null })
				// #endif
			})
		// #endif
	})
}
function drawNvue_init() {
	/*获取元素引用*/
	var ganvas = vnodeCtx.$refs[canvasId.value]
	/*通过元素引用获取canvas对象*/
	var canvasObj = enable(ganvas, {
		bridge: WeexBridge
	})
	return canvasObj.getContext('2d')
}
function save(): Promise<string> {
	return new Promise((su, fa) => {
		if (!ctx) {
			uni.showToast({ title: '初始化失败', icon: 'none' })
			fa('初始化失败')
			return
		}
		let size = props.option.size ?? 0
		// #ifdef APP-NVUE
		uni.showLoading({ title: '...' })
		// ctx.getImageData(0,0,props.width,props.height,function(res:imageData){
		// 	console.log(ArrayBufferToBase64(res.data).length)
		// 	fa(true)
		// 	uni.hideLoading()
		// })

		ctx.toTempFilePath(0, 0, size, size, uni.upx2px(size), uni.upx2px(size), 'png', 1, function (res) {
			uni.hideLoading()
			console.log(res.errMsg)
			if (res.errMsg == 'canvasToTempFilePath:ok') {
				su(res.tempFilePath)
			} else {
				fa(res.errMsg)
			}
		})

		// #endif

		// #ifndef APP-NVUE
		uni.canvasToTempFilePath({
			x: 0,
			y: 0,
			width: uni.upx2px(size),
			height: uni.upx2px(size),
			// #ifndef MP-ALIPAY
			canvasId: canvasId.value,
			// #endif
			canvas: canvas2d,
			success: function (res) {
				// 在H5平台下，tempFilePath 为 base64
				su(res.tempFilePath)
			},
			fail: (res) => {
				console.error(res)
				fa(res)
			}
		})

		// #endif
	})
}
//保存二维码图片。
defineExpose({ save })
</script>

<style></style>
