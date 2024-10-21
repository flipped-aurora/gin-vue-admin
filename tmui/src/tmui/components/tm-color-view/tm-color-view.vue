<template>
	<view @touchmove.stop="stopMove" class="flex flex-col flex-col-top-center">
		<view
			@touchmove.stop="stopMove"
			v-if="showCanvas"
			class="overflow relative"
			ref="webviewWk"
			:style="{ width: `${_width}px`, height: `${_height}px` }"
		>
			<!-- #ifdef APP-NVUE -->
			<web-view
				ref="web"
				src="/hybrid/html/local.html"
				:style="{ width: `${_width}px`, height: `${_height}px` }"
				@onPostMessage="_onMessage"
			></web-view>
			<!-- #endif -->
			<!-- #ifdef MP-WEIXIN || MP-ALIPAY || MP-QQ -->
			<canvas
				@touchstart="touchStart"
				@touchmove="touchMove"
				@touchend="touchEnd"
				class="canvas"
				type="2d"
				id="canvasId"
				:style="{ width: `${_width}px`, height: `${_height}px`, 'touch-action': 'none' }"
			></canvas>
			<!-- #endif -->
			<!-- #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ || APP-NVUE -->
			<canvas
				@touchstart="touchStart"
				@touchmove="touchMove"
				@touchend="touchEnd"
				class="canvas"
				:id="canvasId"
				:canvas-id="canvasId"
				:style="{ width: `${_width}px`, height: `${_height}px`, 'touch-action': 'none' }"
			></canvas>
			<!-- #endif -->
			<!-- #ifndef APP-NVUE -->
			<view
				id="wrapper"
				class="absolute l-0 t-0 wrapper item"
				@touchstart="colorTouch.startDrag"
				@touchmove="colorTouch.onDrag"
				@touchend="colorTouch.endDrag"
				@touchcancel="colorTouch.endDrag"
			>
				<view class="itemwk" @touchend="touchEndWk(0)" @touchmove="touchEndWk(0)"></view>
			</view>
			<view
				id="wrapper2"
				class="absolute r-0 t-0 wrapper item"
				@touchstart="colorTouch.startDrag2"
				@touchmove="colorTouch.onDrag2"
				@touchend="colorTouch.endDrag2"
				@touchcancel="colorTouch.endDrag2"
			>
				<view class="itemwk" @touchend="touchEndWk(1)" @touchmove="touchEndWk(1)"></view>
			</view>
			<!-- #endif -->

			<!-- #ifdef APP-NVUE -->
			<view id="wrapper" ref="wrapper" class="absolute l-0 t-0 wrapper item" @touchstart="nvueStartH">
				<view class="itemwk"></view>
			</view>
			<view id="wrapper2" ref="wrapper2" class="absolute r-0 t-0 wrapper item" @touchstart="nvueStartS">
				<view class="itemwk"></view>
			</view>
			<!-- #endif -->
		</view>
		<view v-if="!showCanvas" class="flex flex-row flex-row-center-center">
			<tm-icon name="tmicon-shuaxin" spin></tm-icon>
		</view>
		<view class="flex flex-row flex-row-center-center mt-32" :style="{ width: `${_width}px` }">
			<view
				class="round-0 shadow-4"
				:style="{
					width: '30px',
					height: '60rpx',
					background: `rgba(${bgcolor.r},${bgcolor.g},${bgcolor.b},${opactiy})`
				}"
			></view>
			<view style="width: 5px"></view>
			<view :style="{ width: _width - 35 + 'px' }">
				<view class="flex flex-row flex-row-center-start">
					<input v-model="bgcolor.r" class="colorInput" :style="{ width: (_width - 30) / 4 - 6 + 'px' }" />
					<input v-model="bgcolor.g" class="colorInput" :style="{ width: (_width - 30) / 4 - 6 + 'px' }" />
					<input v-model="bgcolor.b" class="colorInput" :style="{ width: (_width - 30) / 4 - 6 + 'px' }" />
					<input v-model="opactiy" class="colorInput" :style="{ width: (_width - 30) / 4 - 6 + 'px' }" />
				</view>
			</view>
		</view>
		<view class="mt-10" :style="{ width: `${_width}px` }">
			<tm-button :color="props.color" @click="confirm" block label="确认"></tm-button>
		</view>
	</view>
</template>

<!-- #ifndef APP-NVUE -->
<script module="colorTouch" lang="wxs" src="./colorTouch.wxs"></script>
<!-- #endif -->
<script lang="ts" setup>
import { getCanvas } from '@/tmui/tool/function/getCanvas'
import { ref, computed, watch, nextTick, getCurrentInstance, onMounted } from 'vue'
import TmText from '@/tmui/components/tm-text/tm-text.vue'
import tmIcon from '@/tmui/components/tm-icon/tm-icon.vue'
import tmButton from '@/tmui/components/tm-button/tm-button.vue'
import { colortool } from '@/tmui/tool/theme/colortool'
// #ifdef APP-NVUE
// @ts-ignore
var dom = weex.requireModule('dom')
const Binding = uni.requireNativePlugin('bindingx')
const animation = uni.requireNativePlugin('animation')
// #endif
//因为截止3.6.13+为止，nvue目前无法通过本地图片，和hybrid方案渲染canvas，来渲染本地图片，这个bug只有修复后，才可统一使用地址渲染，暂时用我的远程图片 地址。
const colorimgUrl = 'https://cdn.tmui.design/public/static/color.png'
const emits = defineEmits(['confirm', 'update:modelValue'])
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	width: {
		type: Number,
		default: 500
	},
	height: {
		type: Number,
		default: 500
	},
	color: {
		type: String,
		default: 'primary'
	},
	modelValue: {
		type: String,
		default: '#FF0000'
	}
})
let canvasId = 'c_' + uni.$tm.u.getUid(4)
// #ifdef MP-WEIXIN || MP-ALIPAY || MP-QQ
canvasId = 'canvasId'
// #endif
const _width = computed(() => {
	return 250
	// return uni.upx2px(props.width)
})
const _height = computed(() => {
	return 200
	// return uni.upx2px(props.height)
})
const _widthRpx = computed(() => uni.$tm.u.torpx(_width.value - 30))
let ctx: any = null
let ctxLeft = 0
let ctxTop = 0
let tid: any = NaN
const bindxToken = ref(null)
let dpr = uni.getSystemInfoSync().devicePixelRatio
let isAndroid = false
const showCanvas = ref(true)

let defaultColorrgba = colortool.cssToRgba(props.modelValue)
const bgcolor = ref({ ...defaultColorrgba })
const opactiy = ref(defaultColorrgba.a || 0)
onMounted(() => {
	getquerPos()
})
watch(
	() => props.modelValue,
	() => {
		let pcolor = colortool.cssToRgba(props.modelValue)
		bgcolor.value = { ...pcolor }
		opactiy.value = pcolor.a
	}
)
function confirm() {
	opactiy.value = opactiy.value <= 0 ? 0 : opactiy.value
	opactiy.value = opactiy.value >= 1 ? 1 : opactiy.value
	let cssrgba = colortool.rgbaToCss({ ...bgcolor.value, a: opactiy.value })

	emits('update:modelValue', cssrgba)
	nextTick(() => {
		emits('confirm', cssrgba)
		uni.$tm.u.setClipboardData(cssrgba)
	})
}

function getquerPos() {
	// #ifndef APP-NVUE
	uni.createSelectorQuery()
		.in(proxy)
		.select('#' + canvasId)
		.boundingClientRect((resulst) => {
			ctxLeft = resulst?.left ?? 0
			ctxTop = resulst?.top ?? 0
			getCanvas(proxy, canvasId, _width.value, _height.value).then((res) => {
				ctx = res.ctx
				renderColorHu()
				renderRectFill()
			})
		})
		.exec()
	// #endif

	// #ifdef APP-NVUE
	setTimeout(function () {
		dom.getComponentRect(proxy?.$refs.webviewWk, function (res) {
			if (res?.size) {
				ctxLeft = res?.size.left
				ctxTop = res?.size.top
			}
		})
	}, 120)
	// #endif
}

function _onMessage(e) {
	const message = e.detail.data[0]
	switch (message.action) {
		// web-view 初始化完毕
		case 'onJSBridgeReady':
			proxy.$refs.web.evalJs(`tmColorView_showdiv(${_width.value},${_height.value})`)
			// 执行渲染颜色。
			proxy.$refs.web.evalJs(`tmColorView_renderColorHu()`)
			proxy.$refs.web.evalJs(`tmColorView_renderRectFill()`)
			break

		case 'tmColorView_getColor':
			let colorstr = JSON.parse(message.tmColor)
			let activeIndex = Number(message.tmColorActive)
			bgcolor.value = colorstr
			if (activeIndex == 0) {
				proxy.$refs.web.evalJs(`tmColorView_renderRectFill(${colortool.rgbaToHsla(bgcolor.value).h})`)
			}
			break
	}
}

function touchStart(e: any) {}
function touchMove(e: any) {}
function touchEnd(e: any) {}

function getColor(x, y, index, iscolorget = true) {
	x = x - ctxLeft + 15
	y = y - ctxTop + 15

	if (index == 1) {
		x = x < 40 ? 40 : x
		y = y <= 2 ? 2 : y
	} else {
		x = x < 0 ? 0 : x
		y = y < 0 ? 0 : y
		y = y > 200 ? 200 : y
	}
	// #ifndef APP-NVUE
	if (ctx?.getImageData) {
		let imagedata = ctx.getImageData(x * dpr, y * dpr, 1, 1)
		let [r, g, b, a] = imagedata.data
		let bg = { r, g, b, a: 1 }
		if (iscolorget) {
			bgcolor.value = { ...bg }
		}
		if (x >= 0 && x <= 30 && y >= 0 && y <= _height.value) {
			renderRectFill(colortool.rgbaToHsla(bg).h)
		}
	} else {
		uni.canvasGetImageData(
			{
				canvasId: canvasId,
				x: x,
				y: y,
				width: 1,
				height: 1,
				success(res) {
					let [r, g, b, a] = res.data
					let bg = { r, g, b, a: 1 }
					if (iscolorget) {
						bgcolor.value = { ...bg }
					}
					if (x >= 0 && x <= 30 && y >= 0 && y <= _height.value) {
						renderRectFill(colortool.rgbaToHsla(bg).h)
					}
				},
				fail(er) {
					console.error(er)
				}
			},
			proxy
		)
	}

	// #endif
	// #ifdef APP-NVUE
	proxy.$refs.web.evalJs(`tmColorView_getColor(${x},${y},${index})`)
	// #endif
}

function renderRectFill(H: number = 0) {
	let x = 40
	let dy = 2
	let w = _width.value

	for (let i = 0; i < 100; i++) {
		let gradient = ctx.createLinearGradient(x, i, w - x, i)
		gradient.addColorStop(0, colortool.rgbaToCss(colortool.hslaToRgba({ h: H, s: 0, l: 100 - i, a: 1 })))
		gradient.addColorStop(1, colortool.rgbaToCss(colortool.hslaToRgba({ h: H, s: 100, l: 50 - i / 2, a: 1 })))
		ctx.fillStyle = gradient
		ctx.fillRect(x, (i + 1) * dy, w - x, (i + 1) * dy)
		if (ctx?.draw) {
			ctx?.draw(true)
		}
	}
}
function sleep(timeout: number) {
	return new Promise((res) => {
		setTimeout(() => {
			res()
		}, timeout)
	})
}
async function renderColorHu() {
	let dy = _height.value / 3
	let barcolorWidth = 30
	let x = 0

	let gradient = ctx.createLinearGradient(barcolorWidth / 2, 0, barcolorWidth / 2, dy)
	gradient.addColorStop(0, 'rgba(255,0,0,1)')
	gradient.addColorStop(0.5, 'rgba(255,0,255,1)')
	gradient.addColorStop(1, 'rgba(0,0,255,1)')
	ctx.fillStyle = gradient
	ctx.fillRect(x, 0, barcolorWidth, dy)
	if (ctx?.draw) {
		ctx?.draw()
	}

	let gradient2 = ctx.createLinearGradient(barcolorWidth / 2, dy, barcolorWidth / 2, dy * 2)
	gradient2.addColorStop(0, 'rgba(0,0,255,1)')
	gradient2.addColorStop(0.5, 'rgba(0,255,255,1)')
	gradient2.addColorStop(1, 'rgba(0,255,0,1)')
	ctx.fillStyle = gradient2
	ctx.fillRect(x, dy, barcolorWidth, dy)
	if (ctx?.draw) {
		ctx?.draw(true)
	}

	let gradient3 = ctx.createLinearGradient(barcolorWidth / 2, dy * 2, barcolorWidth / 2, dy * 3)
	gradient3.addColorStop(0, 'rgba(0,255,0,1)')
	gradient3.addColorStop(0.5, 'rgba(255,255,1,1)')
	gradient3.addColorStop(1, 'rgba(255,0,0,1)')
	ctx.fillStyle = gradient3
	ctx.fillRect(x, dy * 2, barcolorWidth, dy)
	// ctx.drawImage(colorimgUrl, 0, 0,_width.value,_height.value)
	if (ctx?.draw) {
		ctx?.draw(true)
	}
}

function touchEndWk(index: number) {
	clearTimeout(tid)

	tid = setTimeout(function () {
		// #ifndef APP-NVUE
		uni.createSelectorQuery()
			.in(proxy)
			.selectAll('.wrapper')
			.boundingClientRect((resulst) => {
				if (!Array.isArray(resulst)) return
				if (resulst.length < 2) return
				let hsla = colortool.rgbaToHsla(bgcolor.value)
				if (index === 0) {
					let x0 = resulst[0].left
					let y0 = resulst[0].top
					getColor(x0, y0, 0, false)
					tid = setTimeout(() => {
						let x0 = resulst[1].left
						let y0 = resulst[1].top
						getColor(x0, y0, 1)
					}, 150)
				} else {
					let x0 = resulst[1].left
					let y0 = resulst[1].top
					getColor(x0, y0, 1)
				}
			})
			.exec()
		// #endif
		// #ifdef APP-NVUE

		// #endif
	}, 200)
}
function wxsCallPos(x, y) {
	// console.log('tmuiColor')
}
let nvue_now_left = 0
function getEl(el: HTMLElement) {
	if (typeof el === 'string' || typeof el === 'number') return el
	if (WXEnvironment) {
		return el.ref
	} else {
		return el instanceof HTMLElement ? el : el.$el
	}
}
function nvueStartH() {
	if (!proxy.$refs?.wrapper) return
	let icon = getEl(proxy.$refs.wrapper)
	let expression = `y+0`
	expression = `(y+${nvue_now_left}>=-15&&y+${nvue_now_left}<=200)?(y<-15?y+${nvue_now_left}:y+${nvue_now_left}):-y`
	let icon_bind = Binding.bind(
		{
			anchor: icon,
			eventType: 'pan',
			props: [
				{
					element: icon,
					property: 'transform.translateY',
					expression: expression
				}
			]
		},
		function (res) {
			if (res.state == 'end') {
				let ly = Math.abs(res.deltaY)
				let top = res.deltaY >= 0 ? false : true
				if (res.deltaY >= 0) {
					nvue_now_left = nvue_now_left + res.deltaY
				} else {
					nvue_now_left = nvue_now_left - ly
				}
				uni.showLoading({
					title: '处理中'
				})
				dom.getComponentRect(proxy?.$refs.wrapper, function (res) {
					if (res?.size) {
						let x0 = res?.size.left
						let y0 = res?.size.top
						getColor(x0, y0, 0)
						uni.hideLoading()
						clearTimeout(tid)
						tid = setTimeout(() => {
							dom.getComponentRect(proxy?.$refs.wrapper2, function (res) {
								if (res?.size) {
									let x0 = res?.size.left
									let y0 = res?.size.top
									getColor(x0, y0, 1)
									uni.hideLoading()
								}
							})
						}, 100)
					}
				})
			} else if (res.state == 'start') {
				// isMoveing.value = true
			}
		}
	)
	bindxToken.value = icon_bind.token
}
let nvue_color_left = 0
let nvue_color_top = 0
function nvueStartS() {
	if (!proxy.$refs?.wrapper2) return
	let icon = getEl(proxy.$refs.wrapper2)
	let expression1 = `x+0`
	let expression2 = `y+0`
	expression1 = `(x+${nvue_color_left}>=-190&&x+${nvue_color_left}<=15)?(x<15?x+${nvue_color_left}:x+${nvue_color_left}):-x`
	expression2 = `(y+${nvue_color_top}>=-12&&y+${nvue_color_top}<=187)?(y<-12?y+${nvue_color_top}:y+${nvue_color_top}):-y`
	let icon_bind = Binding.bind(
		{
			anchor: icon,
			eventType: 'pan',
			props: [
				{
					element: icon,
					property: 'transform.translateX',
					expression: expression1
				},
				{
					element: icon,
					property: 'transform.translateY',
					expression: expression2
				}
			]
		},
		function (res) {
			if (res.state == 'end') {
				let lx = Math.abs(res.deltaX)
				let ly = Math.abs(res.deltaY)
				if (res.deltaX >= 0) {
					nvue_color_left = nvue_color_left + res.deltaX
				} else {
					nvue_color_left = nvue_color_left - lx
				}
				if (res.deltaY >= 0) {
					nvue_color_top = nvue_color_top + res.deltaY
				} else {
					nvue_color_top = nvue_color_top - ly
				}

				uni.showLoading({
					title: '处理中'
				})
				dom.getComponentRect(proxy?.$refs.wrapper2, function (res) {
					if (res?.size) {
						let x0 = res?.size.left
						let y0 = res?.size.top
						getColor(x0, y0, 1)
						uni.hideLoading()
					}
				})
			} else if (res.state == 'start') {
				// isMoveing.value = true
			}
		}
	)
	bindxToken.value = icon_bind.token
}

function stopMove(e: any) {
	if (e?.stopPropagation) {
		e.stopPropagation()
	}
	if (e?.preventDefault) {
		e.preventDefault()
	}
}

defineExpose({ wxsCallPos })
</script>

<style>
.colorwk {
	border-radius: 200px;
	overflow: hidden;
}
.colorInput {
	background-color: rgba(0, 0, 0, 0.03);
	border: 3rpx solid rgba(255, 255, 255, 1);
	height: 64rpx;
	border-radius: 8rpx;
	text-align: center;
	margin-left: 5px;
	color: black;
}
.itemwk {
	width: 32px;
	height: 32px;
	border-radius: 32px;
}
.item {
	/* #ifndef APP-NVUE */
	box-sizing: border-box;
	/* #endif */
	width: 30px;
	height: 30px;
	background-color: rgba(255, 255, 255, 0.6);
	border-radius: 32px;
	border: 2px solid rgba(255, 255, 255, 0.8);
}
</style>
