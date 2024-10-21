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
 * 条形码
 * @description 这是一个条形码组件，属性多，可以生成非常个性化的组件哦~
 * 如果想知道生成的属性请查看：Baroptions类型属性。
 * 更改任意属性，都将会导致重绘
 */
import { getCurrentInstance, computed, ref, PropType, inject, ComponentInternalInstance, onMounted, nextTick, watch } from 'vue'
import CanvasRenderingContext2D from '../../tool/gcanvas/context-2d/RenderingContext.js'
import { drawBarCode, BarcodeObjType, Baroptions } from './drawing'
import { JsBarcode } from './jsbarcode/JsBarcode.js'
import JsBarcodeOptions from './jsbarcode/options/defaults.js'
// #ifdef APP-NVUE
import { enable, WeexBridge } from '../../tool/gcanvas/index.js'
// #endif
const proxy = getCurrentInstance()?.proxy ?? null

const props = defineProps({
	option: {
		type: Object as PropType<Baroptions>,
		default: () => {
			return JsBarcodeOptions
		}
	},
	width: {
		type: Number,
		default: 600
	},
	height: {
		type: Number,
		default: 240
	}
})
const vnodeCtx = proxy
const canvasId = ref('canvasId')
// #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ
canvasId.value = 'tm' + uni.$tm.u.getUid(5)
// #endif
let ctx: CanvasRenderingContext2D
let canvas2d: HTMLCanvasElement
let optsCode = computed(() => {
	let ops = { ...JsBarcodeOptions, ...props.option }
	ops.height = uni.upx2px(props.height) - 40
	return ops
})
const _width = computed(() => props.width)
const _height = computed(() => props.height)
const show = ref(false) //安卓上首次要隐藏不然卡。
onMounted(() => {
	nextTick(function () {
		// #ifdef APP-NVUE
		if (uni.getSystemInfoSync().osName == 'android') {
			setTimeout(function () {
				show.value = true
				init().then(() => draw())
			}, 200)
		} else {
			show.value = true
			init().then(() => draw())
		}
		// #endif
		// #ifndef APP-NVUE

		init().then(() => draw())
		// #endif
	})
})
function draw(opts = optsCode.value) {
	let bcode = JsBarcode(ctx, opts.text, opts)
	let BarcodeObj: BarcodeObjType = bcode._encodings[0][0]
	let is2d = canvas2d ? true : false

	drawBarCode(ctx, BarcodeObj.options, BarcodeObj, is2d, uni.upx2px(_width.value))
}
watch(
	() => props.option,
	() => {
		if (!ctx) {
			init().then(() => draw({ ...optsCode.value }))
		} else {
			draw({ ...optsCode.value })
		}
	},
	{ deep: true }
)
function init() {
	return new Promise((res, rej) => {
		// #ifdef APP-NVUE
		setTimeout(async function () {
			ctx = await drawNvue_init()
			res(true)
		}, 100)
		// #endif
		// #ifdef MP-WEIXIN || MP-ALIPAY || MP-QQ
		setTimeout(async function () {
			const { ctx2d, canvas } = await MpWeix_init()
			ctx = ctx2d
			canvas2d = canvas
			res(true)
		}, 100)
		// #endif
		// #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ || APP-NVUE
		setTimeout(async function () {
			ctx = await appvueH5Other()
			res(true)
		}, 50)
		// #endif
	})
}
//appvue,h5,和其它平台。
function appvueH5Other() {
	return Promise.resolve(uni.createCanvasContext(canvasId.value, vnodeCtx))
}
//支付宝和微信，QQ 支持2d渲染。
function MpWeix_init(): Promise<{
	ctx2d: CanvasRenderingContext2D
	canvas: HTMLCanvasElement
}> {
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
				// #ifndef MP-QQ
				const canvas = res[0].node
				const ctxvb = canvas.getContext('2d')
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
		let width = props.width
		let height = props.height
		// #ifdef APP-NVUE
		uni.showLoading({ title: '...' })
		ctx.toTempFilePath(0, 0, width, height, uni.upx2px(width), uni.upx2px(height), 'png', 1, function (res) {
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
			width: uni.upx2px(width),
			height: uni.upx2px(height),
			canvasId: canvasId.value,
			canvas: canvas2d,
			success: function (res) {
				// 在H5平台下，tempFilePath 为 base64
				su(res.tempFilePath)
			},
			fail: (res) => {
				fa(res)
			}
		})

		// #endif
	})
}

//保存条码图片。
defineExpose({ save })
</script>

<style></style>
