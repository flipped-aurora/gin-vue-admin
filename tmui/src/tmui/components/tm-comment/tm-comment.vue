<template>
	<tm-sheet
		:color="props.color"
		:followTheme="props.followTheme"
		:dark="props.dark"
		:followDark="props.followDark"
		:round="props.round"
		:shadow="props.shadow"
		:outlined="props.outlined"
		:border="props.border"
		:borderStyle="props.borderStyle"
		:borderDirection="props.borderDirection"
		:text="props.text"
		:transprent="props.transprent"
		:linear="props.linear"
		:linearDeep="props.linearDeep"
		:width="props.width"
		:height="props.height"
		:margin="props.margin"
		:padding="props.padding"
	>
		<view class="flex flex-row flex-row-top-start">
			<view v-if="ctl.avatar" class="mr-24" style="min-width: 60rpx; align-items: stretch">
				<tm-avatar
					:round="12"
					@click="emits('avatar-click', $event)"
					:followTheme="props.followTheme"
					:dark="props.dark"
					:followDark="props.followDark"
					:size="60"
					:img="ctl.avatar"
				></tm-avatar>
			</view>
			<view class="flex flex-9 flex-col wh5">
				<view class="flex flex-row" :class="[_align.time]">
					<slot name="author">
						<tm-text
							@click="emits('author-click', $event)"
							:followTheme="props.followTheme"
							:dark="props.dark"
							:followDark="props.followDark"
							:color="ctl.authorColor"
							:fontSize="ctl.authorFontSize"
							:label="ctl.author"
						></tm-text>
					</slot>
					<slot name="time">
						<tm-text
							@click="emits('time-click', $event)"
							v-if="_time"
							:followTheme="props.followTheme"
							:dark="props.dark"
							:followDark="props.followDark"
							_class="pl-24 opacity-5"
							:fontSize="22"
							:label="_time"
						></tm-text>
					</slot>
				</view>
				<view class="flex flex-col py-12">
					<slot name="content">
						<tm-text
							@click="emits('content-click', $event)"
							:followTheme="props.followTheme"
							:dark="props.dark"
							:followDark="props.followDark"
							:fontSize="28"
							:label="ctl.content"
						>
						</tm-text>
					</slot>
				</view>
				<view class="flex flex-row flex-1 flex-nowrap" :class="[_align.action]">
					<slot name="actions">
						<!-- <view class="flex flex-row flex-center pl-16">
							<tm-icon :followTheme="props.followTheme" :dark="props.dark" :followDark="props.followDark" :font-size="24" name="tmicon-comment-dots"></tm-icon>
							<tm-text :followTheme="props.followTheme" :dark="props.dark" :followDark="props.followDark" _class="pl-10"  :font-size="24" label="125"></tm-text>
						</view>
						<view class="flex flex-row flex-center pl-16">
							<tm-icon :followTheme="props.followTheme" :dark="props.dark" :followDark="props.followDark" :font-size="24" name="tmicon-md-heart"></tm-icon>
							<tm-text :followTheme="props.followTheme" :dark="props.dark" :followDark="props.followDark" _class="pl-10" :font-size="24" label="125"></tm-text>
						</view> -->
					</slot>
				</view>
				<tm-divider v-if="ctl.borderBottom"></tm-divider>
				<slot></slot>
			</view>
		</view>
	</tm-sheet>
</template>

<script lang="ts" setup>
/**
 * 评论
 * @description 评论可以嵌套使用。
 * @template default,author,time,content,actions
 * @emits ['avatar-click','author-click','content-click','time-click']
 */
import { getCurrentInstance, computed, ref, provide, inject, onUpdated, onMounted, onUnmounted, nextTick, watch, PropType } from 'vue'
import { cssstyle, tmVuetify, colorThemeType } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmDivider from '../tm-divider/tm-divider.vue'
import tmAvatar from '../tm-avatar/tm-avatar.vue'
const emits = defineEmits(['avatar-click', 'author-click', 'content-click', 'time-click'])
const props = defineProps({
	...custom_props,
	shadow: {
		type: [Number],
		default: 0
	},
	round: {
		type: [Number],
		default: 4
	},
	border: {
		type: [Number],
		default: 0
	},
	width: {
		type: [Number],
		default: 0
	},
	height: {
		type: [Number],
		default: 0
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [32, 8]
	},
	padding: {
		type: Array as PropType<Array<number>>,
		default: () => [24, 24]
	},
	transprent: {
		type: [Boolean],
		default: false
	},
	color: {
		type: String,
		default: 'white'
	},
	//是否显示底部的边线。
	borderBottom: {
		type: [Boolean],
		default: false
	},
	author: {
		type: String,
		default: ''
	},
	authorColor: {
		type: String,
		default: 'primary'
	},
	authorFontSize: {
		type: Number,
		default: 26
	},
	avatar: {
		type: String,
		default: ''
	},
	content: {
		type: String,
		default: ''
	},
	time: {
		type: String,
		default: ''
	},
	//时间和下面的操作按钮是靠左还是靠右。
	align: {
		type: String,
		default: 'right' //right,left
	},
	// 自动格式化时间
	autoFormatTime: {
		type: Boolean,
		default: true
	}
})

// 距离当前时间段。
const _time = computed(() => {
	if (!props.autoFormatTime) return props.time
	return uni.$tm.u.getDateToNewData(props.time)
})

const _align = computed(() => {
	let agn = {
		time: 'flex-between',
		action: 'flex-row-center-end flex-1'
	}
	if (props.align == 'left') {
		agn = {
			time: 'flex-start',
			action: 'flex-row-center-start'
		}
	}
	return agn
})
const ctl = computed(() => {
	return {
		time: props.time,
		content: props.content,
		author: props.author,
		authorColor: props.authorColor,
		avatar: props.avatar,
		authorFontSize: props.authorFontSize,
		borderBottom: props.borderBottom
	}
})
</script>

<style scoped>
.wh5 {
	width: 0;
}
</style>
