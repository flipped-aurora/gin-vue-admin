<template>
	<view class="flex flex-row flex-row-top-start" style="flex-flow: row wrap" :style="[_width ? { width: _width + 'rpx' } : '']">
		<slot></slot>
	</view>
</template>

<script lang="ts" setup>
/**
 * 图片组
 * @description 注意它内部只能搭配图片组件tm-image才能正常。
 * @example 相册集例子
 * <tm-image-group>
		<tm-image :padding="[5,5]" preview :width="150" :height="150" :src="'https://picsum.photos/200/300?id='+item" v-for="item in 14" :key="item"></tm-image>
	</tm-image-group
 */
import { provide, computed } from 'vue'
interface imageItem {
	width: number
	height: number
	src: string
}
const props = defineProps({
	//默认为0,宽度自动。
	width: {
		type: [Number, String],
		default: 0
	}
})
const _width = computed(() => props.width)

const _catchList: Array<imageItem> = []
function pushKey(e: imageItem) {
	let index = _catchList.findIndex((el) => {
		return el.src == e.src
	})
	if (index > -1) {
		_catchList.splice(index, 1, e)
	} else {
		_catchList.push(e)
	}
}
defineExpose({ pushKey: pushKey, tmImageGroup: 'tmImageGroup' })
provide(
	'ImagGrupList',
	computed(() => {
		return _catchList.map((el) => el.src)
	})
)
</script>
<style></style>
