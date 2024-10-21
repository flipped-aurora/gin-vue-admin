<template>
	<view>
		<view
			class="tm-sticky flex flex-col"
			:style="[
				props.model == 'top' ? { top: `${_offset}` } : '',
				props.model == 'left' ? { left: `${_offset}` } : '',
				{ 'z-index': props.zIndex }
			]"
		>
			<view class="flex flex-col">
				<slot name="sticky"></slot>
			</view>
		</view>
		<view>
			<slot></slot>
		</view>
	</view>
</template>
<script lang="ts" setup>
/**
 * 粘性定位
 * @description 请注意使用方法，为了兼容所有端，如果你想让粘性在某块区域，必须把其它内容套在本组件内。
 * @slot sticky 定位内容
 * @slot default 正常的布局文档内容
 * @example  <tm-sticky >
                <template v-slot:sticky>
                        <!-- 这里是粘性内容 -->
                        <text>{{ status }}{{statusw}}</text>
                </template>
                <!-- 这里放置你的内容内容 -->
            </tm-sticky>
 */
import { computed, ref } from 'vue'
const props = defineProps({
	//可能的值为:left,top
	//默认为：top
	model: {
		type: String,
		default: 'top'
	},
	offset: {
		type: [String, Number],
		default: '0px'
	},
	//层级对于nvue是无效的。
	zIndex: {
		type: Number,
		default: 50
	}
})
const _offset = computed(() => {
	if (typeof props.offset === 'number') return props.offset + 'rpx'
	return props.offset
})
</script>
<style scoped>
.tm-sticky {
	position: sticky;
}
</style>
