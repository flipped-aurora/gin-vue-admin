<template>
	<view
		@longpress=""
		@touchstart="touchStart"
		@touchmove.stop.prevent="touchMove"
		@touchend="touchEnd"
		@touchcancel="touchEnd"
		@mousedown="touchStart"
		@mousemove.stop.prevent="touchMove"
		@mouseup="touchEnd"
		@mouseleave="touchEnd"
		class="relative overflow"
		:style="[{ width: `${view_width}px`, height: `${view_height}px`, background: '#000000' }]"
	>
		<view v-if="!url" class="flex flex-center py-24"><text style="font-size: 24rpx; color: white">请先选择图片</text> </view>
		<!-- #ifndef APP-NVUE -->
		<canvas
			id="imgid"
			canvas-id="imgid"
			class="l-0 t-0"
			:style="[
				{
					width: `${rectPos.w}px`,
					height: `${rectPos.h}px`,
					transform: `translate(-1000px,-1000px)`
				}
			]"
		>
		</canvas>
		<!-- #endif -->
		<!-- transform: `translate(${imgPos.x}px,${imgPos.y}px) scale(${imgPos.scale},${imgPos.scale})` -->
		<!-- <image v-if="url" :src="url"></image> -->
		<image
			v-if="url"
			@error="imgError"
			class="absolute l-0 t-0"
			
			:userInteractionEnabled="false"
			@load="imgLoad"
			:src="url"
			:style="[
				{
					width: `${imgPos.w || 1}px`,
					height: `${imgPos.h || 1}px`,
					transform: url ? `translate(${imgPos.x}px,${imgPos.y}px) scale(${imgPos.scale},${imgPos.scale})` : `translate(-1000px,-1000px)`,
					'transform-origin': `${imgPos.center[0]}px ${imgPos.center[1]}px`
				}
			]"
		></image>
		<!-- 		   -->
		<view
			:class="[moveable ? '' : 'ani']"
			:userInteractionEnabled="false"
			class="absolute l-0 t-0 rect "
			:style="[
				{
					width: rectPos.w + 'px',
					height: rectPos.h + 'px',
					transform: `translate(${rectPos.x}px,${rectPos.y}px)`
				}
			]"
		>
		</view>

		<view class="absolute l-0 b-0 zIndex-2 black" :style="{ width: `${view_width}px`, height: `${80}px` }">
			<!-- 	<view class="flex-center mb-24">
				<tm-slider @change="scaleChange" :width="600" :max='3' :min='1'></tm-slider>
			</view> -->
			<view class="flex-row flex flex-row-center-between">
				<!-- 翻转 -->
				<!-- <view class="flex-1 flex flex-center">
					<tm-text :userInteractionEnabled="false" color="white" label="旋转"></tm-text>
				</view> -->
				<!-- 取消 -->
				<view @click="emits('cance')" class="flex-1 flex flex-center">
					<tm-text :userInteractionEnabled="false" color="white" label="取消"></tm-text>
				</view>
				<!-- 选择图片 -->
				<view @click="chooseImg" class="flex-1 flex flex-center">
					<tm-text :userInteractionEnabled="false" color="white" label="选择图片"></tm-text>
				</view>
				<!-- 确认 -->
				<view @click="pushImgToCanvas" class="flex-1 flex flex-center">
					<tm-text :userInteractionEnabled="false" color="white" label="确认"></tm-text>
				</view>
			</view>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 图片裁剪
 * @description nvue平台暂时未实现，你可以把本组件放置到vue页面来使用。
 * @description 因为uniapp平台，暂不支持在nvue页面上canvas上渲染本地路径图片，只支持远程，暂时搁置，另想plugs方法来实现。
 *
 */
import { ref, getCurrentInstance, onMounted, nextTick } from 'vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import tmSlider from '../tm-slider/tm-slider.vue'
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['confirm', 'cance'])
const props = defineProps({
	url: {
		type: String,
		default: ''
	},
	width: {
		type: Number,
		default: 250
	},
	height: {
		type: Number,
		default: 250
	}
})
//当前需要裁剪图片的地址。注意，如果不提供则需要选择对应的图片。
const url = ref(props.url)
const { safeArea, windowWidth, windowHeight, statusBarHeight } = uni.getSystemInfoSync()
// 视窗的宽。
const view_width = safeArea?.width || windowWidth
//视窗的高度。
let view_height = safeArea?.height ?? windowHeight - 44
// #ifdef MP
view_height = safeArea?.height ?? windowHeight
// #endif
//手指触摸的坐标点。
const points = ref({
	px: 0, //手指触摸是的1指坐标x
	py: 0, //手指触摸是的1指坐标y
	px2: 0, //手指触摸是的2指坐标x
	py2: 0, //手指触摸是的2指坐标y
	xy: 0 //两指间的距离.
})
//裁剪框的位置。和宽高。
const rectPos = ref({
	w: props.width, //当前裁剪框的宽和高
	h: props.height,
	x: 0, //当前坐标
	y: 0
})
const imgPos = ref({
	scale: 1, //当前绽放比例
	oldscale: 1,
	center: [0, 0],
	sourceImgWidth: 0, //原始图像的宽和高
	sourceImgHeight: 0,
	w: 0, //缩放后的宽和高
	h: 0,
	x: 0, //当前的坐标。
	y: 0
})
let moveable = ref(false)
let moveType = 0 //0为图片的移动。1为裁剪框的移动。
let scaleing = ref(false) //是否处于双指缩放中。
let ctx = null
onMounted(() => {
	nextTick(function () {
		setTimeout(function () {
			// #ifndef APP-NVUE
			ctx = uni.createCanvasContext('imgid', proxy)
			// #endif
		}, 50)
	})
})
setRectPos()
function touchStart(res: TouchEvent) {
	if (!url.value) return
	moveable.value = true
	let event = res.touches[0]
	let event2 = res.touches[1]
	if (res.type == 'mousedown') {
		event = res
	}
	moveType = getMoveTypeNumber(event.pageX, event.pageY)

	if (moveType == 0) {
		points.value.px = event.pageX - imgPos.value.x
		points.value.py = event.pageY - imgPos.value.y
	} else {
		points.value.px = event.pageX - rectPos.value.x
		points.value.py = event.pageY - rectPos.value.y
	}

	if (event2) {
		points.value.px2 = event2.pageX
		points.value.py2 = event2.pageY
	}

	imgPos.value.oldscale = imgPos.value.scale
}

function touchMove(res: TouchEvent) {
	if (!url.value) {
		return
	}
	let event = res.touches[0]
	let event2 = res.touches[1]
	if (res.type == 'mousedown') {
		event = res
	}
	if (!moveable.value) return
	if (moveType == 0) {
		imgPos.value.x = event.pageX - points.value.px
		imgPos.value.y = event.pageY - points.value.py
	} else {
		rectPos.value.x = event.pageX - points.value.px
		rectPos.value.y = event.pageY - points.value.py
	}

	//双向缩放暂时不实现。
	return
	// points.value.px = event.pageX - points.value.px
	// points.value.py = event.pageY -  points.value.py
	if (event2) {
		scaleing.value = true
		points.value.px2 = event2.pageX
		points.value.py2 = event2.pageY

		// 双指缩放比例计算
		var zoom =
			getDistance(
				{
					x: event.pageX,
					y: event.pageY
				},
				{
					x: event2.pageX,
					y: event2.pageY
				}
			) /
			getDistance(
				{
					x: points.value.px,
					y: points.value.py
				},
				{
					x: points.value.px2,
					y: points.value.py2
				}
			)
		var newScale = imgPos.value.oldscale * zoom

		var newscale = imgPos.value.oldscale * zoom

		imgPos.value.scale = newscale
		const center = getScaleCenter(
			{
				x: event.pageX,
				y: event.pageY
			},
			{
				x: event2.pageX,
				y: event2.pageY
			}
		)
		imgPos.value.x = center[0] - imgPos.value.center[0]
		imgPos.value.y = center[1] - imgPos.value.center[1]
		// imgPos.value.w = newscale*imgPos.value.w
		// imgPos.value.h = newscale*imgPos.value.h
	} else {
		if (scaleing.value) return
		if (moveType == 0) {
			imgPos.value.x = event.pageX - points.value.px
			imgPos.value.y = event.pageY - points.value.py
		} else {
			rectPos.value.x = event.pageX - points.value.px
			rectPos.value.y = event.pageY - points.value.py
		}
	}
}

function touchEnd(res: Event) {
	if (!url.value) {
		return
	}
	moveable.value = false
	scaleing.value = false

	if (rectPos.value.x <= 12) {
		rectPos.value.x = 12
	}
	if (rectPos.value.x >= view_width - rectPos.value.w - 12) {
		rectPos.value.x = view_width - rectPos.value.w - 12
	}

	if (rectPos.value.y <= 12) {
		rectPos.value.y = 12
	}

	if (rectPos.value.y >= view_height - rectPos.value.h - 90) {
		rectPos.value.y = view_height - rectPos.value.h - 90
	}

	if (imgPos.value.x >= rectPos.value.x) {
		imgPos.value.x = rectPos.value.x
	}
	if (imgPos.value.x + imgPos.value.w <= rectPos.value.x + rectPos.value.w) {
		imgPos.value.x = rectPos.value.x - (imgPos.value.w - rectPos.value.w)
	}
	if (imgPos.value.y >= rectPos.value.y) {
		imgPos.value.y = rectPos.value.y
	}
	if (imgPos.value.y + imgPos.value.h <= rectPos.value.y + rectPos.value.h) {
		imgPos.value.y = rectPos.value.y - (imgPos.value.h - rectPos.value.h)
	}
}

function chooseImg() {
	uni.chooseImage({
		count: 1,
		success(res) {
			let files = res.tempFilePaths
			if (files.length == 1) {
				url.value = files[0]
				imgPos.value.scale = 1
				imgPos.value.oldscale = 1
			}
		}
	})
}

function scaleChange(scale: number) {
	scale = scale > 1 ? scale - 0.5 : scale
	if (scale == imgPos.value.scale) return
	imgPos.value.scale = scale
	imgPos.value.h = rectPos.value.h * scale
	imgPos.value.w = rectPos.value.w * scale
	// setImgPosScaleXy()
	if (scale == 1) {
		setImgPosScaleXy()
	}
}

function imgLoad(res: Event) {
	const { width, height } = res.detail
	imgPos.value.sourceImgWidth = width
	imgPos.value.sourceImgHeight = height

	setImgPosScaleXy()
}

function imgError(res) {
	console.log(res)
}
//设置当前移动的类型，应该移动哪个。0移动图片，1矩形。
function getMoveTypeNumber(x: number, y: number) {
	if (x >= rectPos.value.x - 15 && x <= rectPos.value.x + 15 && y >= rectPos.value.y && y <= rectPos.value.y + rectPos.value.h) {
		return 1
	}
	if (x >= rectPos.value.x && x <= rectPos.value.w && y >= rectPos.value.y - 15 && y <= rectPos.value.y + 15) {
		return 1
	}
	if (
		x >= rectPos.value.x + rectPos.value.w - 15 &&
		x <= rectPos.value.x + rectPos.value.w + 15 &&
		y >= rectPos.value.y &&
		y <= rectPos.value.y + rectPos.value.h
	) {
		return 1
	}
	if (
		x >= rectPos.value.x &&
		x <= rectPos.value.x + rectPos.value.w &&
		y >= rectPos.value.y + rectPos.value.h - 15 &&
		y <= rectPos.value.y + rectPos.value.h + 15
	) {
		return 1
	}
	return 0
}
var getDistance = function (start, stop) {
	return Math.hypot(stop.x - start.x, stop.y - start.y)
}
var getScaleCenter = function (start, stop) {
	let minx = (stop.x + start.x) / 2
	let miny = (stop.y + start.y) / 2
	return [minx, miny]
}

///设置剪裁框的坐标。
function setRectPos() {
	let x = (view_width - rectPos.value.w) / 2
	let y = (view_height - rectPos.value.h) / 2
	rectPos.value.x = x
	rectPos.value.y = y
}
//设置当前图片的位置和宽高。
function setImgPosScaleXy() {
	imgPos.value.x = rectPos.value.x
	let now_w = rectPos.value.w
	let now_h = now_w / (imgPos.value.sourceImgWidth / imgPos.value.sourceImgHeight)
	imgPos.value.h = now_h
	imgPos.value.w = now_w

	//如果高度不够。需要放大。
	if (imgPos.value.h < rectPos.value.h) {
		imgPos.value.h = rectPos.value.h
		//计算下倍数。

		imgPos.value.w = imgPos.value.h * (now_w / now_h)
	}
	imgPos.value.y = -(imgPos.value.h - rectPos.value.h) / 2 + rectPos.value.y
	imgPos.value.center = [imgPos.value.w / 2, imgPos.value.h / 2]
}

function pushImgToCanvas() {
	if (!url.value) {
		uni.showToast({
			title: '未选择',
			icon: 'none'
		})
		return
	}
	let dx = Math.abs(imgPos.value.x - rectPos.value.x)
	let dy = Math.abs(imgPos.value.y - rectPos.value.y)
	// #ifndef APP-NVUE
	ctx.drawImage(url.value, -dx, -dy, imgPos.value.w, imgPos.value.h)
	ctx.draw(false, (res) => {
		uni.canvasToTempFilePath(
			{
				canvasId: 'imgid',
				success(ok) {
					// h5为base64
					console.log('绘制成功 ')
					emits('confirm', ok.tempFilePath)
				},
				fail(error) {
					uni.showToast({
						title:"裁切失败",
						icon:'error',
						mask:true
					})
				}
			},
			proxy
		)
	})

	// #endif
	// #ifdef APP-NVUE
	// ctx.drawImage(url.value, -dx, -dy,imgPos.value.w,imgPos.value.h)
	// 因为在nvue上,uniapp有bug，不能绘制本地路径的图片，因此曲线救国。
	// #endif
}
</script>

<style scoped>
.rect {
	border: 1px solid #ffffff;
}

.ani {
	/* #ifndef APP-NVUE */
	transition-duration: 0.3s;
	/* #endif */
	transition-property: transform;
	transition-timing-function: ease;
}
</style>
