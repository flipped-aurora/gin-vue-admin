<template>
	<view class="relative overflow">
		<tm-sheet
			:darkBgColor="props.darkBgColor"
			@click="cellClick"
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
			:_class="props._class"
			:_style="props._style"
			:hover-class="props.hoverClass"
		>
			<view :userInteractionEnabled="true" class="flex flex-row flex-row-center-center" :class="[_computedValue.url ? 'url' : '']">
				<view
					v-if="_computedValue.showAvatar"
					:style="{
						width: `${_computedValue.avatarSize}rpx`,
						height: `${_computedValue.avatarSize}rpx`
					}"
					class="flex flex-row flex-row-center-center"
				>
					<slot name="avatar">
						<tm-image
							:round="_computedValue.avatarRound"
							:width="_computedValue.avatarSize"
							:height="_computedValue.avatarSize"
							:src="_computedValue.avatar"
						></tm-image>
					</slot>
				</view>

				<view class="flex-1 flex flex-row flex-row-center-between" style="width: 0px">
					<view>
						<view class="flex flex-5 flex-col" :class="[_computedValue.showAvatar ? 'pl-24' : '']">
							<slot name="title">
								<tm-text
									:color="_computedValue.titleColor"
									:fontSize="_computedValue.titleFontSize"
									:label="_computedValue.title"
								></tm-text>
							</slot>
							<slot name="label">
								<view v-if="_computedValue.label" class="mt-6">
									<tm-text
										:color="_computedValue.labelColor"
										:fontSize="_computedValue.labelFontSize"
										:label="_computedValue.label"
									></tm-text>
								</view>
							</slot>
						</view>
					</view>
					<view class="flex-1 flex-row flex-row-center-end" style="width: 0px">
						<slot name="rightText">
							<tm-text
								_class="nowrap pr-12"
								:color="_computedValue.rightColor"
								v-if="_computedValue.rightText"
								:fontSize="_computedValue.rightTextSize"
								:label="_computedValue.rightText"
							></tm-text>
						</slot>
						<slot name="right">
							<tm-icon v-if="_computedValue.rightIcon" _class="opacity-3" :name="_computedValue.rightIcon" :fontSize="22"> </tm-icon>
						</slot>
					</view>
				</view>
			</view>
		</tm-sheet>
		<tm-divider
			v-if="_computedValue.bottomBorder"
			:margin="[0, 0]"
			:border="2"
			color="grey-5"
			:real-color="!store.tmStore.dark"
			:style="{
				left: `${_computedValue.avatar !== '' ? _computedValue.avatarSize + _computedValue.margin[0] : 0}rpx`
			}"
		></tm-divider>
	</view>
</template>

<script lang="ts" setup>
/**
   * 单元格
   * @description 常用于列表
   * @example 示例
   * <tm-sheet :round="8" :padding="[0,0]">
      <tm-cell :bottomBorder="false" title="这是标题"></tm-cell>
      <tm-cell :bottomBorder="false"  title="这是标题"></tm-cell>
      <tm-cell  :bottomBorder="false" title="这是标题"></tm-cell>
      <tm-cell :bottomBorder="false"  title="设置" rightText="已阅读同意">
        <template v-slot:right>
          <tm-checkbox ></tm-checkbox>
        </template>
      </tm-cell>
    </tm-sheet>
   */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmImage from '../tm-image/tm-image.vue'
import tmDivider from '../tm-divider/tm-divider.vue'
import { getCurrentInstance, computed, PropType } from 'vue'
import { cssDirection } from '../../tool/lib/interface'
import { custom_props } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '@/tmui/tool/lib/tmpinia'
const emits = defineEmits(['click'])
const props = defineProps({
	...custom_props,
	shadow: {
		type: [Number],
		default: 0
	},
	round: {
		type: [Number,Array] as PropType<Array<number>|number>,
		default: 0
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [32, 0]
	},
	padding: {
		type: Array as PropType<Array<number>>,
		default: () => [24, 24]
	},
	height: {
		type: [Number],
		default: 0
	},
	width: {
		type: [Number],
		default: 0
	},
	transprent: {
		type: [Boolean],
		default: false
	},
	color: {
		type: String,
		default: 'white'
	},
	//标题
	title: {
		type: String,
		default: ''
	},
	titleFontSize: {
		type: [Number],
		default: 28
	},
	titleColor: {
		type: String,
		default: ''
	},
	//标题下方的介绍
	label: {
		type: String,
		default: ''
	},
	labelFontSize: {
		type: [Number],
		default: 22
	},
	labelColor: {
		type: String,
		default: 'grey'
	},
	//右边文字
	rightText: {
		type: String,
		default: ''
	},
	rightIcon: {
		type: String,
		default: 'tmicon-angle-right'
	},
	//右边文字
	rightColor: {
		type: String,
		default: 'grey'
	},
	//右边文字大小。
	rightTextSize: {
		type: Number,
		default: 24
	},
	showAvatar: {
		type: Boolean,
		default: false
	},
	//头像。
	//https://picsum.photos/200
	avatar: {
		type: String,
		default: ''
	},
	avatarSize: {
		type: Number,
		default: 60
	},
	avatarRound: {
		type: Number,
		default: 10
	},
	border: {
		type: [Number],
		default: 0
	},
	borderDirection: {
		type: [String],
		default: cssDirection.bottom
	},
	//显示下边线
	bottomBorder: {
		type: [Boolean],
		default: false
	},
	//当有链接地址时，将打开链接
	url: {
		type: String,
		default: ''
	},
	//暗下强制的背景色，
	//有时自动的背景，可能不是你想要暗黑背景，此时可以使用此参数，强制使用背景色，
	//只能是颜色值。
	darkBgColor: {
		type: String,
		default: ''
	},
	hoverClass: {
		type: String,
		default: 'opacity-6'
	},
})

const store = useTmpiniaStore()

function cellClick(e: any) {
	emits('click', e)
	if (props.url !== '') {
		try {
			uni.navigateTo({
				url: props.url,
				fail(error) {
					console.error('打开连接错误：', error)
				}
			})
		} catch (e) {
			//TODO handle the exception
		}
	}
}

const _computedValue = computed(() => props)
</script>

<style scoped>
.url {
	/* #ifdef H5 */
	cursor: pointer;
	/* #endif */
}
</style>
