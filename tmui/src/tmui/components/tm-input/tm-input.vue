<template>
	<tm-sheet :eventPenetrationEnabled="true" :transprent="true" :margin="props.margin" :padding="props.padding">
		<tm-sheet
			:transprent="props.transprent"
			:round="props.round"
			no-level
			:margin="[0, 0]"
			:padding="_inputPadding"
			:border="props.border"
			:text="props.text"
			:color="_color"
			:outlined="props.outlined"
			:shadow="props.shadow"
			:linear="props.linear"
			:linearDeep="props.linearDeep"
			_style="transition:border 0.24s"
		>
			<view
				class="flex flex-row relative"
				@click="inputClick($event, '')"
				:class="[propsDetail.type == 'textarea' ? propsDetail.layoutAlign : 'flex-row-center-start']"
				:style="[propsDetail.autoHeight && propsDetail.type == 'textarea' ? {} : { height: `${_height}rpx` }]"
			>
				<view v-if="propsDetail.search || propsDetail.searchLabel" class="px-9"></view>
				<slot name="left"></slot>
				<view v-if="propsDetail.prefix" class="pr-16">
					<tm-icon
						_style="transition:color 0.24s"
						:font-size="propsDetail.fontSize"
						:color="props.prefixColor"
						:name="propsDetail.prefix"
						:customicon="props.customicon"
					></tm-icon>
				</view>
				<view v-if="propsDetail.prefixLabel" class="pr-24">
					<tm-text
						_style="transition:color 0.24s"
						:font-size="propsDetail.fontSize"
						:color="props.prefixColor"
						:label="propsDetail.prefixLabel"
					></tm-text>
				</view>

				<view v-if="!isAndroid" @click="inputClick($event, 'ali')" class="flex-1 relative flex-row flex" :style="[{ width: '0px' }]">
					<!-- <view @click.stop="emits('click',$event)" class=" l-0 t-0 flex-1 " :style="{height: `${_height}rpx`,background:'red'}"></view> -->
					<input
						class="flex-1"
						:userInteractionEnabled="false"
						v-if="propsDetail.type != 'textarea'"
						:value="_value"
						:focus="propsDetail.focus"
						@focus="focus"
						@blur="blur"
						@confirm="confirm"
						@input="inputHandler"
						@keyboardheightchange="emits('keyboardheightchange', $event)"
						:password="showPasswordText"
						:maxlength="propsDetail.maxlength"
						:disabled="propsDetail.disabled"
						:cursorSpacing="propsDetail.cursorSpacing"
						:confirmType="propsDetail.confirmType"
						:confirmHold="propsDetail.confirmHold"
						:autoBlur="propsDetail.autoBlur"
						:holdKeyboard="propsDetail.holdKeyboard"
						:adjustPosition="propsDetail.adjustPosition"
						:readonly="propsDetail.readyOnly"
						:type="propsDetail.type"
						:placeholder="propsDetail.placeholder"
						:style="[
							{
								height: `${_height}rpx`,
								color: propsDetail.fontColor ? propsDetail.fontColor : tmcomputed.textColor,
								'text-align': props.align,
								fontSize: `${propsDetail.fontSize_px}px`,
								transition: 'color 0.24s'
							}
						]"
						:placeholder-style="`fontSize:${propsDetail.fontSize_px}px;${props.placeholderStyle}`"
						:ready-only="propsDetail.readyOnly"
					/>

					<textarea
						:userInteractionEnabled="false"
						v-if="propsDetail.type == 'textarea'"
						:value="_value"
						:focus="propsDetail.focus"
						@focus="focus"
						@blur="blur"
						@confirm="confirm"
						@input="inputHandler"
						@keyboardheightchange="emits('keyboardheightchange', $event)"
						:maxlength="propsDetail.maxlength"
						:disabled="propsDetail.disabled"
						:placeholder="propsDetail.placeholder"
						:cursorSpacing="propsDetail.cursorSpacing"
						:confirmHold="propsDetail.confirmHold"
						:autoBlur="propsDetail.autoBlur"
						:holdKeyboard="propsDetail.holdKeyboard"
						:cursor="propsDetail.cursor"
						:show-confirm-bar="propsDetail.showConfirmBar"
						:selectionStart="propsDetail.selectionStart"
						:selectionEnd="propsDetail.selectionEnd"
						:disable-default-padding="propsDetail.disableDefaultPadding"
						:fixed="propsDetail.fixed"
						:autoHeight="propsDetail.autoHeight"
						:readonly="propsDetail.readyOnly"
						:adjustPosition="propsDetail.adjustPosition"
						:type="propsDetail.type"
						:style="[
							propsDetail.autoHeight ? {} : { height: `${_height}rpx` },
							{
								width: 'auto',
								'word-break': 'break-word',
								color: propsDetail.fontColor ? propsDetail.fontColor : tmcomputed.textColor,
								'text-align': props.align,
								fontSize: `${propsDetail.fontSize_px}px`,
								transition: 'color 0.24s'
							}
						]"
						class="wrap flex-1"
						:placeholder-style="`fontSize:${propsDetail.fontSize_px}px;${props.placeholderStyle}`"
						:ready-only="propsDetail.readyOnly"
					></textarea>
				</view>
				<view v-if="isAndroid" class="flex-1 relative flex-row flex" :style="[{ width: '0px' }]">
					<!-- <view @click.stop="emits('click',$event)" class=" l-0 t-0 flex-1 " :style="{height: `${_height}rpx`,background:'red'}"></view> -->
					<input
						class="flex-1"
						@click.stop="emits('click', $event)"
						:userInteractionEnabled="false"
						v-if="propsDetail.type != 'textarea'"
						:value="_value"
						:focus="propsDetail.focus"
						@focus="focus"
						@blur="blur"
						@confirm="confirm"
						@input="inputHandler"
						@keyboardheightchange="emits('keyboardheightchange', $event)"
						:password="showPasswordText"
						:disabled="propsDetail.disabled"
						:cursorSpacing="propsDetail.cursorSpacing"
						:confirmType="propsDetail.confirmType"
						:confirmHold="propsDetail.confirmHold"
						:autoBlur="propsDetail.autoBlur"
						:holdKeyboard="propsDetail.holdKeyboard"
						:adjustPosition="propsDetail.adjustPosition"
						:maxlength="propsDetail.maxlength"
						:type="propsDetail.type"
						:readonly="propsDetail.readyOnly"
						:placeholder="propsDetail.placeholder"
						:style="[
							{
								height: `${_height}rpx`,
								color: propsDetail.fontColor ? propsDetail.fontColor : tmcomputed.textColor,
								'text-align': props.align,
								fontSize: `${propsDetail.fontSize_px}px`
							}
						]"
						:placeholder-style="`fontSize:${propsDetail.fontSize_px}px;${props.placeholderStyle}`"
					/>
					<textarea
						@click.stop="emits('click', $event)"
						:userInteractionEnabled="false"
						v-if="propsDetail.type == 'textarea'"
						:value="_value"
						:focus="propsDetail.focus"
						@focus="focus"
						@blur="blur"
						@confirm="confirm"
						@input="inputHandler"
						@keyboardheightchange="emits('keyboardheightchange', $event)"
						:disabled="propsDetail.disabled"
						:placeholder="propsDetail.placeholder"
						:cursorSpacing="propsDetail.cursorSpacing"
						:confirmHold="propsDetail.confirmHold"
						:autoBlur="propsDetail.autoBlur"
						:holdKeyboard="propsDetail.holdKeyboard"
						:adjustPosition="propsDetail.adjustPosition"
						:maxlength="propsDetail.maxlength"
						:autoHeight="propsDetail.autoHeight"
						:cursor="propsDetail.cursor"
						:show-confirm-bar="propsDetail.showConfirmBar"
						:selectionStart="propsDetail.selectionStart"
						:selectionEnd="propsDetail.selectionEnd"
						:disable-default-padding="propsDetail.disableDefaultPadding"
						:readonly="propsDetail.readyOnly"
						:fixed="propsDetail.fixed"
						:type="propsDetail.type"
						:style="[
							propsDetail.autoHeight ? {} : { height: `${_height}rpx` },
							{
								width: 'auto',
								'word-break': 'break-word',
								color: propsDetail.fontColor ? propsDetail.fontColor : tmcomputed.textColor,
								'text-align': props.align,
								fontSize: `${propsDetail.fontSize_px}px`
							}
						]"
						class="wrap flex-1"
						:placeholder-style="`fontSize:${propsDetail.fontSize_px}px;${props.placeholderStyle}`"
					></textarea>
				</view>
				<view @click="clearBtn" class="pl-16" v-if="propsDetail.showClear && _valueLenChar > 0">
					<tm-icon
						:customicon="props.customicon"
						_style="transition:color 0.24s"
						:userInteractionEnabled="false"
						:font-size="propsDetail.fontSize"
						:color="props.clearAndEyeColor"
						name="tmicon-times-circle-fill"
					>
					</tm-icon>
				</view>
				<view class="pl-16" v-if="_requiredError">
					<tm-icon _style="transition:color 0.24s" :font-size="propsDetail.fontSize" name="tmicon-exclamation-circle"></tm-icon>
				</view>
				<view class="pl-16" v-if="propsDetail.suffix">
					<tm-icon
						:customicon="props.customicon"
						_style="transition:color 0.24s"
						:font-size="propsDetail.fontSize"
						:color="props.suffixColor"
						:name="propsDetail.suffix"
					></tm-icon>
				</view>

				<view v-if="propsDetail.suffixLabel" class="pl-16">
					<tm-text
						_style="transition:color 0.24s"
						:font-size="propsDetail.fontSize"
						:color="props.suffixColor"
						:label="propsDetail.suffixLabel"
					></tm-text>
				</view>

				<view @click="changeSeePassword" class="pl-16" v-if="showPasswordIcon">
					<!-- tmicon-eyeslash-fill -->
					<tm-icon
						:color="props.clearAndEyeColor"
						_style="transition:color 0.24s"
						:userInteractionEnabled="false"
						:font-size="propsDetail.fontSize"
						:name="showPasswordText ? 'tmicon-eyeslash-fill' : 'tmicon-eye-fill'"
					></tm-icon>
				</view>

				<!-- #ifndef MP-ALIPAY -->
				<view v-if="propsDetail.showCharNumber && _valueLenChar > 0 && propsDetail.type != 'textarea'" class="pl-16 flex-row flex">
					<tm-text _style="transition:color 0.24s" :label="_valueLenChar"></tm-text>
					<tm-text _style="transition:color 0.24s" v-if="propsDetail.maxlength > 0" :label="'/' + propsDetail.maxlength"></tm-text>
				</view>
				<!-- 原因是支付宝小程序自带了计数器。会导致重叠。 -->
				<view
					v-if="propsDetail.showCharNumber && _valueLenChar > 0 && propsDetail.type == 'textarea'"
					class="pl-16 flex-row flex absolute r-0"
					:class="[`b-${12}`]"
				>
					<tm-text _style="transition:color 0.24s" :label="_valueLenChar"></tm-text>
					<tm-text _style="transition:color 0.24s" v-if="propsDetail.maxlength > 0" :label="'/' + propsDetail.maxlength"></tm-text>
				</view>
				<!-- #endif -->
				<slot name="right">
					<view v-if="propsDetail.search || propsDetail.searchLabel" class="pl-16">
						<TmButton
							:round="props.round"
							:width="props.searchWidth"
							:followTheme="props.followTheme"
							@click="searchClick"
							:color="props.searchBgColor"
							:font-size="24"
							:height="_height - 11"
							:padding="[16, 0]"
							:block="!props.searchWidth"
							:margin="[0, 0]"
							:fontColor="props.searchFontColor"
							:icon="propsDetail.search"
							:label="propsDetail.searchLabel"
						></TmButton>
					</view>
				</slot>
			</view>
		</tm-sheet>
	</tm-sheet>
</template>

<script lang="ts" setup>
import { computed, PropType, ref, watch, getCurrentInstance, inject, toRaw } from 'vue'
import { inputPushItem, rulesItem } from './../tm-form-item/interface'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
import TmButton from '../tm-button/tm-button.vue'
const store = useTmpiniaStore()
const emits = defineEmits(['focus', 'blur', 'confirm', 'input', 'update:modelValue', 'clear', 'search', 'keyboardheightchange', 'click'])
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	...custom_props,
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	color: {
		type: String,
		default: 'grey-4'
	},
	searchBgColor:{
		type: String,
		default: 'primary'
	},
	searchFontColor: {
		type: String,
		default: ''
	},
	searchWidth: {
		type: Number,
		default: 0
	},
	prefixColor: {
		type: String,
		default: ''
	},
	suffixColor: {
		type: String,
		default: ''
	},
	//激活时的主题配色。
	focusColor: {
		type: String,
		default: 'primary'
	},
	/** 清除按钮，显示密码按钮的颜色 */
	clearAndEyeColor:{
		type: String,
		default: ''
	},
	//默认使用自动配色
	fontColor: {
		type: String,
		default: ''
	},
	text: {
		type: Boolean,
		default: true
	},
	outlined: {
		type: Boolean,
		default: false
	},
	border: {
		type: Number,
		default: 0
	},
	transprent: {
		type: Boolean,
		default: false
	},
	round: {
		type: Number,
		default: 3
	},
	shadow: {
		type: Number,
		default: 0
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	padding: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	height: {
		type: Number,
		default: 64
	},
	//前缀图标
	prefix: {
		type: String,
		default: ''
	},
	//前缀文字
	prefixLabel: {
		type: String,
		default: ''
	},
	//后缀图标
	suffix: {
		type: String,
		default: ''
	},
	//后缀文字
	suffixLabel: {
		type: String,
		default: ''
	},

	fontSize: {
		type: Number,
		default: 30
	},
	//tmicon-search
	search: {
		type: String,
		default: ''
	},
	//搜索
	searchLabel: {
		type: String,
		default: ''
	},
	showClear: {
		type: Boolean,
		default: false
	},
	password: {
		type: Boolean,
		default: false
	},
	//是否禁用
	disabled: {
		type: Boolean,
		default: false
	},
	placeholder: {
		type: String,
		default: '请输入内容'
	},
	//错误时，提示的文本。
	errorLabel: {
		type: String,
		default: '请输入内容'
	},
	//对齐方式。
	//left,right,center
	align: {
		type: String as PropType<'left' | 'right' | 'center'>,
		default: 'left'
	},
	modelValue: {
		type: [String, Number],
		default: ''
	},
	inputPadding: {
		type: Array as PropType<Array<number>>,
		default: () => [24, 0]
	},
	//是否显示字符统计。
	showCharNumber: {
		type: Boolean,
		default: false
	},
	maxlength: {
		type: Number,
		default: -1
	},
	type: {
		type: String as PropType<'text' | 'number' | 'idcard' | 'digit' | 'tel' | 'safe-password' | 'nickname' | 'textarea'>,
		default: 'text'
	},
	cursorSpacing: {
		type: Number,
		default: 24
	},
	confirmType: {
		type: String as PropType<'send' | 'search' | 'next' | 'go' | 'done'>,
		default: 'done'
	},
	confirmHold: {
		type: Boolean,
		default: false
	},
	autoBlur: {
		type: Boolean,
		default: true
	},
	holdKeyboard: {
		type: Boolean,
		default: false
	},
	adjustPosition: {
		type: Boolean,
		default: true
	},
	//默认的聚集状态
	focus: {
		type: Boolean,
		default: false
	},
	cursor: {
		type: Number,
		default: 0
	},
	showConfirmBar: {
		type: Boolean,
		default: true
	},
	selectionStart: {
		type: Number,
		default: -1
	},
	selectionEnd: {
		type: Number,
		default: -1
	},
	disableDefaultPadding: {
		type: Boolean,
		default: false
	},
	fixed: {
		type: Boolean,
		default: false
	},
	placeholderStyle: {
		type: String,
		default: ''
	},
	autoHeight: {
		type: Boolean,
		default: false
	},
	readyOnly: {
		type: Boolean,
		default: false
	},
	/**横向布局的对齐类,主要是用来配置文本域时,左图标需要顶对齐或者左中对齐. */
	layoutAlign: {
		type: String,
		default: 'flex-row-top-start'
	},
	customicon: {
		type: Boolean,
		default: false
	},
})

let parentFormItem: any = proxy?.$parent
while (parentFormItem) {
	if (parentFormItem?.tmFormComnameFormItem == 'tmFormComnameFormItem' || !parentFormItem) {
		break
	} else {
		parentFormItem = parentFormItem?.$parent ?? undefined
	}
}

const isAndroid = ref(false)
isAndroid.value = uni.getSystemInfoSync().osName == 'android' ? true : false
const _height = computed(() => props.height)
const _inputPadding = computed(() => {
	if (props.search !== '' || props.searchLabel !== '') {
		return [4, 0]
	}
	return props.inputPadding
})
let timerId: any = NaN
function debounce(func: Function, wait = 500, immediate = false) {
	// 清除定时器
	if (!isNaN(timerId)) clearTimeout(timerId)
	// 立即执行，此类情况一般用不到
	if (immediate) {
		var callNow = !timerId
		timerId = setTimeout(() => {
			timerId = NaN
		}, wait)

		if (callNow) typeof func === 'function' && func()
	} else {
		// 设置定时器，当最后一次操作后，timeout不会再被清除，所以在延时wait毫秒后执行func回调方法
		timerId = setTimeout(() => {
			typeof func === 'function' && func()
			timerId = NaN
		}, wait)
	}
}
const propsDetail = computed(() => {
	return {
		...props,
		fontSize_px: uni.upx2px(props.fontSize)
	}
})
const _blackValue = props.modelValue
// 设置响应式全局组件库配置表。
const tmcfg = computed(() => store.tmStore)
//自定义样式：
const customCSSStyle = computed(() => computedStyle(props))
//自定类
const customClass = computed(() => computedClass(props))
//是否暗黑模式。
const isDark = computed(() => computedDark(props, tmcfg.value))
//当前是否显示检验错误状态？
const _requiredError = ref(false)
//是否聚焦中。
const _foucsActive = ref(props.focus || false)
watch(
	() => props.focus,
	() => {
		_foucsActive.value = props.focus
	}
)
const _color = computed(() => {
	let color = props.color
	if (_foucsActive.value) {
		if (props.followTheme && store.tmStore.color) {
			color = store.tmStore.color
		} else {
			color = props.focusColor
		}
	}
	if (_requiredError.value) color = 'red'
	return color
})
//计算主题
const tmcomputed = computed(() => {
	const _props = { ...props, color: _color.value }
	return computedTheme(_props, isDark.value, tmcfg.value)
})

//显示密码和关闭密码。
const showPasswordText = ref(propsDetail.value.password)
const showPasswordIcon = computed(() => props.password)
const _errorLabel = ref(props.errorLabel)
const _value = ref(props.modelValue)
const _valueLenChar = computed(() => {
	//在ios上字符的长度不能采用str.length来计算。因为一个中文=2。一个英文=1；
	let str = String(_value.value).split('')
	return str.length
})
watch(
	() => props.modelValue,
	() => (_value.value = props.modelValue)
)

function searchClick() {
	emits('search', _value.value)
}
function clearBtn() {
	_value.value = ''
	emits('update:modelValue', '')
	emits('clear')
}
function changeSeePassword() {
	showPasswordText.value = !showPasswordText.value
}
function focus(e: any) {
	_foucsActive.value = true
	emits('focus', e)
	// pushFormItem();
}
function blur(e: any) {
	_foucsActive.value = false
	// pushFormItem();
	emits('blur', e)
}
function confirm() {
	emits('confirm', _value.value)
}
function inputHandler(e: CustomEvent) {
	// #ifndef MP-WEIXIN
	_value.value = e.detail.value
	// #endif
	emits('input', e.detail.value)
	emits('update:modelValue', e.detail.value)

	return e.detail.value
}
function inputClick(e: Event, type: string) {
	// e.stopPropagation();

	if (type == 'ali') {
		debounce(
			() => {
				emits('click', e)
			},
			200,
			true
		)
		return
	} else {
		debounce(() => emits('click', e), 200, true)
	}
}
</script>

<style scoped>
/* #ifndef APP-NVUE */
input,
textarea {
	background-color: transparent;
	background: transparent;
}

/* #endif */
</style>
