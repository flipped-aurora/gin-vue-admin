<template>
	<view :class="['flex flex-row', `px-${props.margin[0]}`, `pb-${props.margin[1]}`]">
		<tm-sheet
			:margin="[0, 0]"
			:padding="[24, 24]"
			paren-class="flex-1"
			class="flex-1"
			:color="_disable ? _disableBgColor : props.color"
			:_style="props._style"
			:followTheme="props.followTheme"
			:dark="props.dark"
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
			_class="flex flex-row flex-between flex-row-start-start"
		>
			<view v-if="props.showRight" class="pr-24 flex flex-col">
				<slot name="thumb">
					<tm-avatar v-if="_thumb" :size="80" :round="24" :img="_thumb"></tm-avatar>
					<view v-if="!_thumb" class="flex flex-col">
						<view class="flex-row flex flex-row-bottom-center">
							<tm-text
								:follow-dark="true"
								:userInteractionEnabled="false"
								:color="_disable ? _disableColor : props.mainColor"
								:font-size="22"
								:label="_priceDetail.prefix"
							></tm-text>
							<tm-text
								:userInteractionEnabled="false"
								:color="_disable ? _disableColor : props.mainColor"
								_class="px-10 text-weight-b"
								:font-size="42"
								:label="_priceDetail.price"
							></tm-text>
							<tm-text
								:follow-dark="true"
								:userInteractionEnabled="false"
								:color="_disable ? _disableColor : props.mainColor"
								:font-size="22"
								:label="_priceDetail.suffix"
							></tm-text>
						</view>
						<view class="flex flex-center">
							<tm-text
								:userInteractionEnabled="false"
								:color="_disable ? _disableColor : _fontColor"
								_class="pr-10 opacity-7"
								:font-size="22"
								:label="_priceDetail.subtext"
							></tm-text>
						</view>
					</view>
				</slot>
			</view>
			<view class="flex-1 flex flex-col" style="width: 0px">
				<view class="flex flex-row flex-between">
					<view class="flex flex-col flex-1">
						<slot>
							<tm-text
								:follow-dark="true"
								:color="_disable ? _disableColor : _fontColor"
								:font-size="36"
								_class="text-weight-b"
								:label="_rightDetail.title"
							></tm-text>
							<tm-text
								:follow-dark="true"
								_class="opacity-7"
								:color="_disable ? _disableColor : _fontColor"
								:font-size="24"
								:label="_rightDetail.subtitle"
							></tm-text>
						</slot>
					</view>
					<view class="flex flex-row flex-row-center-end" style="width: 150rpx">
						<slot name="btn">
							<tm-button
								@click="emits('click', $event)"
								:disabled="_disable"
								:color="_disable ? _disableColor : props.mainColor"
								:font-color="_disable ? 'grey' : ''"
								:text="props.btnTextMode"
								size="small"
								:width="120"
								:label="_btnLabel"
							></tm-button>
						</slot>
					</view>
				</view>
				<view class="flex flex-col pt-16">
					<view class="flex flex-row flex-between">
						<tm-text
							_class="opacity-7"
							:color="_disable ? _disableColor : _fontColor"
							:font-size="22"
							:label="_rightDetail.time"
						></tm-text>
						<view @click="_extraActive = !_extraActive" v-if="props.extra" class="flex flex-row flex-row-center-center opacity-7">
							<tm-text
								:userInteractionEnabled="false"
								:color="_disable ? _disableColor : _fontColor"
								_class="pr-10"
								:font-size="22"
								:label="_moreText"
							></tm-text>
							<tm-icon
								:userInteractionEnabled="false"
								:color="_disable ? _disableColor : _fontColor"
								:font-size="20"
								:name="!_extraActive ? 'tmicon-angle-down' : 'tmicon-angle-up'"
							></tm-icon>
						</view>
					</view>
					<view v-if="props.extra && _extraActive" class="flex flex-row flex-row-top-start">
						<view class="flex-1">
							<view style="height: 12px"></view>
							<slot name="extra"></slot>
						</view>
					</view>
				</view>
			</view>
		</tm-sheet>
	</view>
</template>

<script lang="ts" setup>
/**
 * 优惠卷
 */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmAvatar from '../tm-avatar/tm-avatar.vue'
import { surplice, rightSurplice } from './interface'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import tmButton from '../tm-button/tm-button.vue'
import { custom_props } from '../../tool/lib/minxs'
import { computed, PropType, ref } from 'vue'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()
const emits = defineEmits<{
	(e: 'click'): void
}>()
const props = defineProps({
	...custom_props,
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [32, 12]
	},
	transprent: {
		type: Boolean,
		default: false
	},
	text: {
		type: Boolean,
		default: false
	},
	color: {
		type: String,
		default: 'white'
	},
	fontColor: {
		type: String,
		default: 'black'
	},
	priceDetail: {
		type: Object as PropType<surplice>,
		default: () => {
			return {
				price: 100,
				suffix: '元',
				prefix: '',
				subtext: '满减券'
			}
		}
	},
	rightDetail: {
		type: Object as PropType<rightSurplice>,
		default: () => {
			return {
				title: '券的标题',
				subtitle: '券的小标题',
				time: '有效期:2022-6-3-2022-7-3'
			}
		}
	},
	shadow: {
		type: Number,
		default: 0
	},
	round: {
		type: Number,
		default: 3
	},
	border: {
		type: Number,
		default: 0
	},
	//优惠券左边是金额文本模式还是图片模式，默认为空，即金额文本模式。
	thumb: {
		type: String,
		default: ''
	},
	//是否显示左边金额或者头像。
	showRight: {
		type: Boolean,
		default: true
	},
	//是否显示额外的详情内容。
	extra: {
		type: Boolean,
		default: false
	},
	//额外内容初始打开状态。
	extraActive: {
		type: Boolean,
		default: false
	},
	moreText: {
		type: String,
		default: '规则详情'
	},
	//强调色，金额 和 按钮的主题色
	mainColor: {
		type: String,
		default: 'red'
	},
	btnTextMode: {
		type: Boolean,
		default: false
	},
	btnLabel: {
		type: String,
		default: '立即使用'
	},
	//是否禁用，等于已使用。
	disable: {
		type: Boolean,
		default: false
	},
	disableColor: {
		type: String,
		default: 'grey-1'
	},
	disableBgColor: {
		type: String,
		default: 'grey-3'
	}
})

const _priceDetail = computed(() => props.priceDetail)
const _rightDetail = computed(() => props.rightDetail)
const _thumb = computed(() => props.thumb)
const _extraActive = ref(props.extraActive)
const _moreText = computed(() => props.moreText)
const _btnLabel = computed(() => props.btnLabel)
const _disable = computed(() => props.disable)
const _disableColor = computed(() => props.disableColor)
const _disableBgColor = computed(() => props.disableBgColor)
const _isDark = computed(() => store.tmStore.dark)
const _fontColor = computed(() => {
	if (store.tmStore.dark && props.fontColor !== '' && (props.fontColor == 'black' || props.fontColor == 'white')) {
		return 'white'
	}
	return props.fontColor
})
</script>

<style></style>
