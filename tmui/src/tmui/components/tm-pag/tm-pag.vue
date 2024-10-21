<template>
	<canvas type="webgl" :style="{ width: props.width + props.unit, height: props.height + props.unit }" :canvas-id="cid" :id="cid"></canvas>
</template>
<script lang="ts" setup>
/**
 * 使用前必须安装pag文件,
 * 微信端:npm install libpag-lite-miniprogram
 * H5,浏览器端:npm install libpag-lite
 */
import { getCanvas } from '@/tmui/tool/function/getCanvas'
import { getCurrentInstance, onMounted } from 'vue'

// #ifdef H5 || APP-VUE
import { PAGView, types } from './pag.esm.js'
// #endif
// #ifdef MP-WEIXIN
import { PAGView } from './pag'
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
	/** 如果提供了就会默认加载出来 */
	url: {
		type: String,
		default: ''
	},
	/** 当前提供了url资源后,这里为true就会自动播放 */
	autoPlay: {
		type: Boolean,
		default: true
	},
	/** 0表示无限循环. */
	loop: {
		type: Number,
		default: 0
	},
	unit: {
		type: String,
		default: 'rpx'
	}
})
const cid = 'c_' + uni.$tm.u.getUid()
const pag: any = null

onMounted(() => {
	setTimeout(function () {
		init()
	}, 100)
})
function init() {
	return new Promise((resolve, rejecj) => {
		if (!props.url) return
		getPag().then((d) => {
			uni.request({
				url: props.url,
				method: 'GET',
				responseType: 'arraybuffer',
				success(res) {
					const pagView = PAGView.init(res.data, d.ctx, {
						renderingMode: 'Webgl',
						useScale: true
					})
					pagView.setRepeatCount(props.loop)
					pagView.setScaleMode('LetterBox')
					if (props.autoPlay) {
						pagView.play()
					}
				},
				fail(error) {
					console.error('资源加载出错:', error)
					rejecj(error)
				}
			})
		})
	})
}
function getPag(): Promise<{ ctx: WebGLContextEvent; pag: PAGView }> {
	return new Promise((resolve, rejecj) => {
		getCanvas(proxy, cid).then((vx) => {
			let ctx = vx.node
			// #ifdef H5
			ctx = vx.node
			// #endif
			resolve({
				ctx: ctx,
				pag: PAGView
			})
		})
	})
}
</script>
