<template>
	<view>
		<!-- #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ || APP-NVUE -->

		<canvas
			@click="emits('click', $event)"
			v-if="!isPc"
			@touchstart="touchStart"
			@touchmove="touchMove"
			@touchend="touchEnd"
			:id="canvasId"
			:canvas-id="canvasId"
			class="canvas"
			:style="{ width: `${_width}rpx`, height: `${_height}rpx` }"
		></canvas>

		<canvas
			@click="emits('click', $event)"
			v-if="isPc"
			:id="canvasId"
			:canvas-id="canvasId"
			class="canvas"
			:style="{ width: `${_width}rpx`, height: `${_height}rpx` }"
		></canvas>
		<!-- #endif -->
		<!-- #ifdef MP-WEIXIN || MP-QQ -->
		<canvas
			@click="emits('click', $event)"
			@touchstart="touchStart"
			@touchmove="touchMove"
			@touchend="touchEnd"
			type="2d"
			id="canvasId"
			canvas-id="canvasId"
			class="canvas"
			:style="{ width: `${_width}rpx`, height: `${_height}rpx` }"
		></canvas>
		<!-- #endif -->
		<!-- #ifdef MP-ALIPAY -->
		<canvas
			@click="emits('click', $event)"
			@touchstart="touchStart"
			@touchmove="touchMove"
			@touchend="touchEnd"
			type="2d"
			:id="canvasId"
			:canvas-id="canvasId"
			class="canvas"
			:style="{ width: `${_width}rpx`, height: `${_height}rpx` }"
		></canvas>
		<!-- #endif -->

		<!-- #ifdef APP-NVUE -->
		<web-view
			ref="web"
			src="/hybrid/html/local.html"
			:style="{ width: `${_width}rpx`, height: `${_height}rpx` }"
			@onPostMessage="_onMessage"
		></web-view>
		<!-- #endif -->
	</view>
</template>
<script lang="ts" setup>
/**
 * Echart图表
 * @description 非nvue端：5.3.2 ,nvue端：5.4.3
 * ref:getChart:获取成功渲染的图表。
 * 事件：onInit:渲染成功后执行，并返回chart对象。
 * 安装百度图表 npm install echarts --save 后需要作下生产下的兼容，发布不影响，但开发时会报错，很烦。
 * 请找到：node_modules/echarts/lib/core/echarts.js,
 * 原文：
 *
 * if (process.env.NODE_ENV !== 'production') {
  var root =
  hasWindow ? window : global;
  defaultRenderer = root.__ECHARTS__DEFAULT__RENDERER__ || defaultRenderer;
  var devUseDirtyRect = root?.__ECHARTS__DEFAULT__USE_DIRTY_RECT__;
  defaultUseDirtyRect = devUseDirtyRect == null ? defaultUseDirtyRect : devUseDirtyRect;

  改成：
  * if (process.env.NODE_ENV !== 'production') {
   var root =
   hasWindow ? window : global;
   【修改一】defaultRenderer = root?.__ECHARTS__DEFAULT__RENDERER__ ?? defaultRenderer;
   【修改二】var devUseDirtyRect = root?.__ECHARTS__DEFAULT__USE_DIRTY_RECT__??null;
   【修改三】defaultUseDirtyRect = devUseDirtyRect == null ? defaultUseDirtyRect : devUseDirtyRect;
}
 */
import {
	getCurrentInstance,
	computed,
	ref,
	onMounted,
	nextTick,
	ComponentInternalInstance,
} from 'vue';
import WxCanvas from './canvasinit';
import mytmcharts from './nvuechart'
import * as echarts from "echarts";
import tmText from "../tm-text/tm-text.vue"

// import * as echarts from "./simple";


const proxy = getCurrentInstance()?.proxy??null;

const emits = defineEmits(['onInit', 'touchStart', 'touchMove', 'touchEnd', 'mousedown', 'mousemove', 'mouseup', 'click'])
const props = defineProps({
	width: {
		type: Number,
		default: 750
	},
	height: {
		type: Number,
		default: 450
	},
})
const canvasId = ref("canvasId")
// #ifdef H5 || APP-PLUS || APP-VUE
canvasId.value = "tm" + new Date().getTime();
// #endif
const _width = computed(() => props.width)
const _height = computed(() => props.height)
let ctx: UniApp.CanvasContext;
let ctxNode:any;
let chart: echarts.ECharts | null = null
const pixelRatio = uni.getSystemInfoSync().pixelRatio
const is2d = ref(false)
// #ifdef MP-WEIXIN || MP-ALIPAY || MP-QQ
is2d.value = true
// #endif
const isPc = ref(false)

isPc.value = uni.getSystemInfoSync().deviceType == 'pc' ? true : false;

function wrapEvent(e: Event) {
	if (!e) return;
	if (!e.preventDefault) {
		e.preventDefault = function () { };
	}
	return e;
}
onMounted(() => {
	nextTick(() => {
		if (is2d.value) {
			setTimeout(() => MpWeix_init(), 100)
		} else {
			// #ifndef APP-PLUS-NVUE
			setTimeout(() => appvueH5Other(), 50)
			// #endif
		}
	})

})
//appvue,h5,和其它平台。
function appvueH5Other(fun:Function) {
	echarts.registerPreprocessor((option: any) => {
		if (option && option.series) {
			if (option.series.length > 0) {
				option.series.forEach((series: echarts.SeriesOption) => {
					series.progressive = 0;
				});
			} else if (typeof option.series === 'object') {
				option.series.progressive = 0;
			}
		}
	});
	ctx = uni.createCanvasContext(canvasId.value, proxy);
	if (!isPc.value) {
		const canvas: any = new WxCanvas(ctx, canvasId.value, false, false)
		echarts.setPlatformAPI({ createCanvas: () => canvas });
		chart = echarts.init(canvas, "", {
			width: uni.upx2px(_width.value),
			height: uni.upx2px(_height.value),
		});
		canvas.setChart(chart);
	} else {
		const canvasNode:HTMLCanvasElement|undefined  = document.querySelector('#' + canvasId.value)?.getElementsByTagName("canvas")[0];
		document.querySelector('#' + canvasId.value)?.removeChild(document.querySelector('#' + canvasId.value)?.getElementsByTagName("div")[0])
		ctx = canvasNode?.getContext("2d")
		const canvas: any = new WxCanvas(ctx, canvasId.value, false, false)
		chart = echarts.init(canvasNode);
		chart.on("mousedown", (e) => emits('mousedown', e))
		chart.on("mousemove", (e) => emits('mousemove', e))
		chart.on("mouseup", (e) => emits('mouseup', e))
		chart.on("mouseover", (e) => emits('mouseover', e))
	}

	emits("onInit", chart)
	if(typeof fun === 'function'){
		fun(chart)
	}
}
//支付宝和微信，QQ 支持2d渲染。
function MpWeix_init(fun:Function) {
	echarts.registerPreprocessor((option: any) => {
		if (option && option.series) {
			if (option.series.length > 0) {
				option.series.forEach((series: echarts.SeriesOption) => {
					series.progressive = 0;
				});
			} else if (typeof option.series === 'object') {
				option.series.progressive = 0;
			}
		}
	});

	const query = uni.createSelectorQuery().in(proxy)
	// #ifdef MP-ALIPAY
	query.select('#canvasId').node().exec((res2) => {
		const canvasNode = res2[0].node;

		let ctxvb: UniApp.CanvasContext = canvas.getContext('2d');
		canvasNode.width = res[0].width * pixelRatio
		canvasNode.height = res[0].height * pixelRatio
		ctx = ctxvb;
		ctxNode = canvasNode
		const canvas = new WxCanvas(ctx, canvasId.value, true, canvasNode)
		echarts.setPlatformAPI({
			// Same with the old setCanvasCreator
			createCanvas() {
				return canvas;
			},
		});
		chart = echarts.init(canvas, null, {
			width: uni.upx2px(_width.value),
			height: uni.upx2px(_height.value),
			devicePixelRatio: pixelRatio
		});
		canvas.setChart(chart);
		emits("onInit", chart)
		if(typeof fun === 'function'){
			fun(chart)
		}
	})
	// #endif
	// #ifdef MP-WEIXIN || MP-QQ
	query.select('#canvasId')
		.fields({
			node: true,
			size: true,
			context: true
		})
		.exec((res) => {

			// #ifdef MP-WEIXIN
			const canvasNode = res[0].node
			const ctxvb = canvasNode.getContext('2d')
			canvasNode.width = res[0].width * pixelRatio
			canvasNode.height = res[0].height * pixelRatio
			ctxvb.scale(pixelRatio, pixelRatio)
			ctx = ctxvb;
			ctxNode = canvasNode
			const canvas = new WxCanvas(ctx, canvasId.value, true, canvasNode)
			echarts.setPlatformAPI({
				// Same with the old setCanvasCreator
				createCanvas() {
					return canvas;
				},
			});

			chart = echarts.init(canvas, null, {
				width: uni.upx2px(_width.value),
				height: uni.upx2px(_height.value),
				devicePixelRatio: pixelRatio
			});

			canvas.setChart(chart);
			emits("onInit", chart)

			// #endif
			// #ifdef MP-QQ

			const canvasNode = {}
			const ctxvb = ctx2d: res[0].context
			canvasNode.width = res[0].width * pixelRatio
			canvasNode.height = res[0].height * pixelRatio
			ctxvb.scale(pixelRatio, pixelRatio)
			ctx = ctxvb;
			const canvas = new WxCanvas(ctx, canvasId.value, true, canvasNode)
			echarts.setPlatformAPI({
				// Same with the old setCanvasCreator
				createCanvas() {
					return canvas;
				},
			});
			chart = echarts.init(canvas, null, {
				width: uni.upx2px(_width.value),
				height: uni.upx2px(_height.value),
				devicePixelRatio: pixelRatio
			});
			canvas.setChart(chart);
			emits("onInit", chart)
			// #endif

			if(typeof fun === 'function'){
				fun(chart)
			}
		})
	// #endif


}



function getChart() {
	return new Promise((res,rej)=>{
		if (is2d.value) {
			setTimeout(() => MpWeix_init(chart=>res(chart)), 100)
		} else {
			// #ifdef APP-PLUS-NVUE
			res(chart)
			// #endif
			// #ifndef APP-PLUS-NVUE
			setTimeout(() => appvueH5Other(chart=>res(chart)), 50)
			// #endif
		}
	})
}

function getImg(){
	return new Promise((res,rej)=>{
		if(!ctx){
			console.error('tmChart没有初始化')
			rej('没有初始化')
			return
		}
		uni.canvasToTempFilePath({
		  x: 0,
		  y: 0,
		  width: _width.value,
		  height: _height.value,
		  canvasId: canvasId.value,
		  canvas: ctxNode,
		  success: function (respone) {
		    // 在H5平台下，tempFilePath 为 base64
		    res(respone.tempFilePath);
		  },
		  fail: (respone) => {
			console.error(respone)
		    rej(res);
		  },
		});
	})
}

function _onMessage(e){
	const message = e.detail.data[0]
	switch (message.action) {
	  // web-view 初始化完毕
	  case 'onJSBridgeReady':
		// 初始化图表库。
		let w = uni.$tm.u.topx(_width.value);
		let h = uni.$tm.u.topx(_height.value);
	    proxy?.$refs.web.evalJs(`echart_createDom(${w},${h})`);
	    proxy?.$refs.web.evalJs(`echart_createChart()`);
		chart = new mytmcharts(proxy?.$refs.web,w,h)
		emits("onInit", chart)
	    break

	   case 'tmColorView_getColor':

	   break
	}
}




defineExpose({getChart,getImg})

function compareVersion(v11: string, v22: string) {
	let v1 = v11.split('.')
	let v2 = v22.split('.')
	const len = Math.max(v1.length, v2.length)

	while (v1.length < len) {
		v1.push('0')
	}
	while (v2.length < len) {
		v2.push('0')
	}

	for (let i = 0; i < len; i++) {
		const num1 = parseInt(v1[i])
		const num2 = parseInt(v2[i])

		if (num1 > num2) {
			return 1
		} else if (num1 < num2) {
			return -1
		}
	}
	return 0
}

function touchStart(e: TouchEvent) {
	if (chart && e.touches.length > 0) {
		var touch = e.touches[0];
		var handler = chart.getZr().handler;
		handler.dispatch('mousedown', {
			zrX: touch.x,
			zrY: touch.y,
			preventDefault: () => { },
			stopPropagation: () => { }
		});
		handler.dispatch('mousemove', {
			zrX: touch.x,
			zrY: touch.y,
			preventDefault: () => { },
			stopPropagation: () => { }
		});
		handler.processGesture(wrapTouch(e), 'start');
		emits('touchStart', e)
	}
}

function touchMove(e: TouchEvent) {
	if (chart && e.touches.length > 0) {
		var touch = e.touches[0];
		var handler = chart.getZr().handler;
		handler.dispatch('mousemove', {
			zrX: touch.x,
			zrY: touch.y,
			preventDefault: () => { },
			stopPropagation: () => { }
		});
		handler.processGesture(wrapTouch(e), 'change');
		emits('touchMove', e)
	}
}

function touchEnd(e: TouchEvent) {
	if (chart) {
		const touch = e.changedTouches ? e.changedTouches[0] : {};
		var handler = chart.getZr().handler;
		handler.dispatch('mouseup', {
			zrX: touch.x,
			zrY: touch.y,
			preventDefault: () => { },
			stopPropagation: () => { }
		});
		handler.dispatch('click', {
			zrX: touch.x,
			zrY: touch.y
		});
		handler.processGesture(wrapTouch(e), 'end');
		emits('touchEnd', e)
	}
}

function wrapTouch(event: TouchEvent) {
	for (let i = 0; i < event.touches.length; ++i) {
		const touch = event.touches[i];
		touch.offsetX = touch.x;
		touch.offsetY = touch.y;
	}
	return event;
}
</script>
