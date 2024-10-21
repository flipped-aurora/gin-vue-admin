<template>
	<view @mouseleave="touchsend" @touchcancel="touchsend" ref="tmspin" :style="{ width: `${props.width}rpx`, height: `${props.height}rpx` }">
		<!-- #ifdef APP-NVUE -->
		<gcanvas
			v-if="show"
			@touchstart="touchstart"
			@touchmove="touchsmove"
			@touchend="touchsend"
			:id="canvasId"
			:ref="canvasId"
			class="canvas"
			:style="{ width: `${props.width}rpx`, height: `${props.height}rpx` }"
		>
		</gcanvas>
		<!-- #endif -->
		<!-- #ifdef MP-WEIXIN || MP-ALIPAY || MP-QQ -->
		<canvas
			@touchstart="touchstart"
			@touchmove="touchsmove"
			@touchend="touchsend"
			@mousedown="touchstart"
			@mousemove.stop="touchsmove"
			@mouseup.stop="touchsend"
			type="2d"
			id="canvasId"
			canvas-id="canvasId"
			class="canvas"
			:style="{ width: `${props.width}rpx`, height: `${props.height}rpx` }"
		></canvas>
		<!-- #endif -->
		<!-- #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ || APP-NVUE -->
		<canvas
			@touchstart.stop="touchstart"
			@touchmove.stop="touchsmove"
			@touchend.stop="touchsend"
			@mousedown.stop="touchstart"
			@mousemove.stop="touchsmove"
			@mouseup.stop="touchsend"
			@mouseleave.stop="touchsend"
			@touchcancel.stop="touchsend"
			:id="canvasId"
			:canvas-id="canvasId"
			class="canvas"
			:style="{
				width: `${props.width}rpx`,
				height: `${props.height}rpx`,
				'touch-action': 'none'
			}"
		></canvas>
		<!-- #endif -->
	</view>
</template>

<script lang="ts" setup>
/**
 * 签名板
 * @description 方便签名业务。在当前截止最新的3.5.0版本，uni在h5端有自身的bug无法使用。等官方修复 即可使用。
 */
import { number } from 'echarts'
import {
	getCurrentInstance,
	computed,
	ref,
	ComponentInternalInstance,
	inject,
	onUpdated,
	onMounted,
	onUnmounted,
	nextTick,
	watch,
	ssrContextKey
} from 'vue'

// #ifdef APP-NVUE
import CanvasRenderingContext2D from '../../tool/gcanvas/context-2d/RenderingContext.js'
import { enable, WeexBridge } from '../../tool/gcanvas/index.js'
const dom = uni.requireNativePlugin('dom')
// #endif

const proxy = getCurrentInstance()?.proxy ?? null

const props = defineProps({
	width: {
		type: Number,
		default: 750
	},
	height: {
		type: Number,
		default: 500
	},
	lineWidth: {
		type: Number,
		default: 5
	},
	lineColor: {
		type: String,
		default: 'red'
	}
})

const canvasId = ref('canvasId')
// #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ
canvasId.value = 'tm' + uni.$tm.u.getUid(5)
// #endif

let ctx: CanvasRenderingContext2D
let canvasObject: HTMLCanvasElement
let ctxLeft = 0
let ctxTop = 0
let drawhd: draw
const show = ref(false) //安卓上首次要隐藏不然卡。
let isAndroid = false
// #ifdef APP-NVUE
isAndroid = uni.getSystemInfoSync().osName == 'android'
// #endif
onMounted(() => {
	// #ifdef APP-NVUE
	if (isAndroid) {
		setTimeout(() => {
			show.value = true
			setTimeout(function () {
				drawNvue_init()
			}, 100)
		}, 200)
	} else {
		show.value = true
		setTimeout(() => drawNvue_init(), 250)
	}
	// #endif

	// #ifdef MP-WEIXIN || MP-ALIPAY || MP-QQ
	setTimeout(() => MpWeix_init(), 100)
	// #endif
	// #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ || APP-NVUE
	setTimeout(() => appvueH5Other(), 50)
	// #endif
})
//appvue,h5,和其它平台。
function appvueH5Other() {
	ctx = uni.createCanvasContext(canvasId.value, proxy)
	uni.createSelectorQuery()
		.in(proxy)
		.select('#' + canvasId.value)
		.boundingClientRect((result) => {
			ctxLeft = result.left ?? 0
			ctxTop = result.top ?? 0
			drawhd = new draw(ctx, uni.upx2px(props.width), uni.upx2px(props.height))
		})
		.exec()
}
function drawNvue_init() {
	/*获取元素引用*/
	var ganvas = proxy?.$refs[canvasId.value]
	/*通过元素引用获取canvas对象*/
	var canvasObj = enable(ganvas, {
		bridge: WeexBridge
	})
	canvasObject = canvasObj
	ctx = canvasObj.getContext('2d')
	nextTick(function () {
		setTimeout(function () {
			dom?.getComponentRect(proxy?.$refs.tmspin, function (res: any) {
				if (res?.size) {
					ctxLeft = Math.floor(res.size.left)
					ctxTop = Math.floor(res.size.top)
					drawhd = new draw(ctx, uni.upx2px(props.width), uni.upx2px(props.height))
				}
			})
		}, 200)
	})
}
//支付宝和微信，QQ 支持2d渲染。
function MpWeix_init() {
	const query = uni.createSelectorQuery().in(proxy)

	// #ifdef MP-ALIPAY
	query
		.select('#canvasId')
		.node()
		.exec((res2) => {
			const canvas = res2[0].node
			let ctxvb: UniApp.CanvasContext = canvas.getContext('2d')
			ctx = ctxvb
			drawhd = new draw(ctx, uni.upx2px(props.width), uni.upx2px(props.height))
		})
	// #endif
	// #ifdef MP-WEIXIN || MP-QQ
	query
		.select('#canvasId')
		.fields({
			node: true,
			size: true,
			rect: true,
			context: true
		})
		.exec((res: Array<UniApp.NodeInfo>) => {
			ctxLeft = res[0].left
			ctxTop = res[0].top
			// #ifndef MP-QQ
			const canvas = res[0].node
			const ctxvb = canvas.getContext('2d')
			const dpr = uni.getSystemInfoSync().pixelRatio
			canvas.width = res[0].width * dpr
			canvas.height = res[0].height * dpr
			ctxvb.scale(dpr, dpr)
			ctx = ctxvb
			canvasObject = canvas
			drawhd = new draw(ctx, uni.upx2px(props.width), uni.upx2px(props.height))
			// #endif
			// #ifdef MP-QQ
			ctx = res[0].context
			drawhd = new draw(ctx, uni.upx2px(props.width), uni.upx2px(props.height))
			// #endif
		})
	// #endif
}

function touchstart(event: TouchEvent | MouseEvent) {
	if (!drawhd) return
	if (event.type.indexOf('mouse') == -1 && event.changedTouches.length == 1) {
		var touch = event.changedTouches[0]
		// #ifdef APP-NVUE
		if (isAndroid) {
			drawhd.down(touch.pageX, touch.pageY)
		} else {
			drawhd.down(touch.pageX - ctxLeft, touch.pageY - ctxTop)
		}
		// #endif
		// #ifndef APP-NVUE
		drawhd.down(touch.x, touch.y)
		// #endif
	} else {
		drawhd.down(event.pageX - event.currentTarget.offsetLeft - ctxLeft, event.pageY - event.currentTarget.offsetTop - ctxTop)
	}
}
function touchsmove(event: TouchEvent | MouseEvent) {
	if (!drawhd) return
	try{
		if (event?.preventDefault) event?.preventDefault()
		if (event?.stopPropagation) event?.stopPropagation()
	}catch(e){
		//TODO handle the exception
	}
	if (event.type.indexOf('mouse') == -1 && event.changedTouches.length == 1) {
		var touch = event.changedTouches[0]
		// #ifdef APP-NVUE
		if (isAndroid) {
			drawhd.move(touch.pageX, touch.pageY)
		} else {
			drawhd.move(touch.pageX - ctxLeft, touch.pageY - ctxTop)
		}
		// #endif
		// #ifndef APP-NVUE
		drawhd.move(touch.x, touch.y)
		// #endif
	} else {
		drawhd.move(event.pageX - event.currentTarget.offsetLeft - ctxLeft, event.pageY - event.currentTarget.offsetTop - ctxTop)
	}
}
function touchsend(event: TouchEvent | MouseEvent) {
	if (!drawhd) return
	if (event.type.indexOf('mouse') == -1 && event.changedTouches.length == 1) {
		var touch = event.changedTouches[0]

		// #ifdef APP-NVUE
		if (isAndroid) {
			drawhd.up(touch.pageX, touch.pageY)
		} else {
			drawhd.up(touch.pageX - ctxLeft, touch.pageY - ctxTop)
		}
		// #endif
		// #ifndef APP-NVUE
		drawhd.up(touch.x, touch.y)
		// #endif
	} else {
		drawhd.up(event.pageX - event.currentTarget.offsetLeft - ctxLeft, event.pageY - event.currentTarget.offsetTop - ctxTop)
	}
}
/**
 * 清除当前画板
 */
function clear() {
	if (!ctx) {
		uni.showToast({ title: '初始化失败', icon: 'none' })
		return
	}
	ctx.clearRect(0, 0, uni.upx2px(props.width), uni.upx2px(props.height))
	// #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ
	ctx.draw(false)
	// #endif
}
/**
 * 保存当前画板内容
 * 在h5端面返回的是base64位图片内容，其它端返回的是临时路径。
 */
interface imageData {
	width: number
	height: number
	data: ArrayBuffer
}
function save(): Promise<string> {
	return new Promise((su, fa) => {
		if (!ctx) {
			uni.showToast({ title: '初始化失败', icon: 'none' })
			fa('初始化失败')
			return
		}
		// #ifdef APP-NVUE
		uni.showLoading({ title: '...' })
		// ctx.getImageData(0,0,props.width,props.height,function(res:imageData){
		// 	console.log(ArrayBufferToBase64(res.data).length)
		// 	fa(true)
		// 	uni.hideLoading()
		// })
		ctx.toTempFilePath(0, 0, props.width, props.height, uni.upx2px(props.width), uni.upx2px(props.height), 'png', 1, function (res) {
			uni.hideLoading()
			console.log(res.errMsg)
			if (res.errMsg == 'canvasToTempFilePath:ok') {
				su(res.tempFilePath)
			} else {
				fa(res.errMsg)
			}
		})

		// #endif

		//webgl保存图片的方法
		// #ifndef APP-NVUE
		uni.canvasToTempFilePath(
			{
				x: 0,
				y: 0,
				destWidth: uni.upx2px(props.width),
				destHeight: uni.upx2px(props.height),
				width: props.width,
				height: props.height,
				// #ifndef MP-ALIPAY
				canvasId: canvasId.value,
				// #endif
				canvas: canvasObject,
				success: function (res) {
					// 在H5平台下，tempFilePath 为 base64
					su(res.tempFilePath)
				},
				fail: (res) => {
					console.error(res)
					fa(res)
				}
			},
			proxy
		)
		// #endif
	})
}
defineExpose({ save, clear })

interface points {
	x: number
	y: number
}
class draw {
	_x = 0
	_y = 0
	_lineWidth = 2
	_lineColor = 'black'
	cx: CanvasRenderingContext2D
	width = 0
	height = 0
	_isDown = false
	_points: Array<points> = []
	constructor(ctx: CanvasRenderingContext2D, w: number, h: number, lineWidth = 2, lineColor = 'black') {
		this._lineColor = lineColor
		this._lineWidth = lineWidth
		this.width = w
		this.height = h
		this.cx = ctx
	}
	down(x = 0, y = 0) {
		this._isDown = true
		this._x = x
		this._y = y
	}
	move(x = 0, y = 0) {
		if (!this._isDown) return

		// #ifdef MP-WEIXIN || MP-ALIPAY || MP-QQ
		ctx.strokeStyle = props.lineColor
		ctx.lineWidth = props.lineWidth
		ctx.lineCap = 'round'
		// #endif
		// #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ
		ctx.setStrokeStyle(props.lineColor)
		ctx.setLineWidth(props.lineWidth)
		ctx.setLineCap('round')
		// #endif

		ctx.beginPath()
		ctx.moveTo(this._x, this._y)
		ctx.lineTo(x, y)
		ctx.stroke()
		ctx.closePath()

		// #ifndef MP-WEIXIN || MP-ALIPAY
		ctx.draw(true)
		// #endif
		this._x = x
		this._y = y
	}
	up(x = 0, y = 0) {
		this._isDown = false
		this._x = x
		this._y = y
	}
}
</script>

<style scoped></style>
