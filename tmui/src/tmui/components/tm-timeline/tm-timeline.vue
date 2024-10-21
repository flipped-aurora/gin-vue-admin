<template>
	<view class="flex" :style="{ 'flex-direction': reverse ? 'column-reverse' : 'column' }">
		<slot></slot>
	</view>
</template>
<script lang="ts" setup>
import { computed, provide } from 'vue'
/**
 * 时间轴
 * @description 内部只可放置 tm-timeline-item 组件，不可放置其它组件。
 * @example 示例如下：
    <tm-timeline>
        <tm-timeline-item time="20222-3-25">
            <tm-sheet :shadow="2" :margin="[0,0]">
                <tm-text label="已经到了中转站，请注意查收已经到了中转站，请注意查收已经到了中转站，请注意查收已经到了中转站，请注意查收" ></tm-text>
            </tm-sheet>
        </tm-timeline-item>
    </tm-timeline>
 */
const props = defineProps({
	//是否反转。
	reverse: {
		type: Boolean,
		default: false
	},
	//left,right,center[暂时不要使用center]
	//对齐方式，左，中，右。
	position: {
		type: String,
		default: 'left'
	}
})

provide(
	'tmTimeLinePosition',
	computed(() => props.position)
)
</script>
<style scoped></style>
