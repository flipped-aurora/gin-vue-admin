<template>
	<button
		@click="onclick"
		@touchstart="touchstart"
		@touchend="touchend"
		@longpress="emits('longpress', $event)"
		@touchcancel="touchcancel"
		@touchmove="emits('touchmove', $event)"
		@getphonenumber="emits('getphonenumber', $event)"
		@error="emits('error', $event)"
		@opensetting="emits('opensetting', $event)"
		@launchapp="emits('launchapp', $event)"
		@contact="emits('contact', $event)"
		@chooseavatar="emits('chooseavatar', $event)"
		:form-type="props.formType"
		:openType="props.openType"
		:appParameter="props.appParameter"
		:sessionFrom="props.sessionFrom"
		:sendMessageTitle="props.sendMessageTitle"
		:sendMessagePath="props.sendMessagePath"
		:sendMessageImg="props.sendMessageImg"
		:sendMessageCard="props.sendMessageCard"
		:loading="_load"
		:disabled="_disabled"
		:group-id="props.groupId"
		:guild-id="props.guildId"
		:public-id="props.publicId"
		:hover-start-time="10000000"
		hover-stop-propagation
		:class="[
			'button alipay flex flex-row flex-row-center-center',
			isclickOn && !_disabled && !_load ? props.hoverClass + ' bhover' : '',
			!_disabled && !_load ? 'webpc' : '',
			_load || _disabled ? 'opacity-8' : '',
			!isDisabledRoundAndriod ? `round-${btnSizeObj.round}` : '',
			customClass
		]"
		:style="[
			{
				marginLeft: _margin[0] + 'rpx',
				marginTop: _margin[1] + 'rpx',
				marginRight: _margin[2] + 'rpx',
				marginBottom: _margin[3] + 'rpx',
				paddingLeft: _padding[0] + 'rpx',
				paddingTop: _padding[1] + 'rpx',
				paddingRight: _padding[2] + 'rpx',
				paddingBottom: _padding[3] + 'rpx'
			},
			{
				height: btnSizeObj.h + props.unit,
				fontSize: btnSizeObj.fontSize + props.unit,
				color: textColor,
				lineHeight: btnSizeObj.h + props.unit
			},
			btnSizeObj.w && !props.block ? { width: btnSizeObj.w + props.unit } : '',
			tmcomputed.borderCss,
			_bgcolor,
			!_transprent && _shadow > 0 && !props.text ? tmcomputed.shadowColor : '',
			customCSSStyle
		]"
	>
		<slot>
			<tm-icon
				:lineHeight="btnSizeObj.fontSize * 0.9"
				v-if="_icon && !_load"
				:userInteractionEnabled="false"
				:color="_fontColor"
				:_class="_label ? 'pr-10' : ''"
				:unit="props.unit"
				:fontSize="btnSizeObj.fontSize * 0.9"
				:name="_icon"
			></tm-icon>
			<!-- #ifdef APP-NVUE -->
			<tm-text :color="props.fontColor" :font-size="btnSizeObj.fontSize" :unit="props.unit" :label="_label" _class="text-weight-b"></tm-text>
			<!-- #endif -->
			<!-- #ifndef APP-NVUE -->
			{{ _label }}
			<!-- #endif -->
		</slot>
	</button>
</template>

<script lang="ts" setup>
/**
 * 按钮
 * @description 属性与原生按钮一致，对于微信登录授权进行了便捷封装。
 * @slot default 默认插槽。
 */
import { btnSize } from './interface'
import { computed, PropType, ref, getCurrentInstance, provide } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import { cssstyle, tmVuetify } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '@/tmui/tool/lib/tmpinia'
import theme from '@/tmui/tool/theme/theme'
const store = useTmpiniaStore()

/**
 * 事件说明
 * @description 事件属性与原生 一 致
 * @links 见官网：https://uniapp.dcloud.io/component/button.html
 */
const emits = defineEmits<{
	(e: 'click', event: Event | TouchEvent): void
	(e: 'touchstart', event: Event | TouchEvent): void
	(e: 'touchmove', event: Event | TouchEvent): void
	(e: 'touchcancel', event: Event | TouchEvent): void
	(e: 'touchend', event: Event | TouchEvent): void
	(e: 'tap', event: Event | TouchEvent): void
	(e: 'longpress', event: Event | TouchEvent): void
	(e: 'getphonenumber', event: any): void
	(e: 'getUserInfo', event: any): void
	(e: 'getUserProfile', event: any): void
	(e: 'error', event: any): void
	(e: 'opensetting', event: any): void
	(e: 'launchapp', event: any): void
	(e: 'contact', event: any): void
	(e: 'chooseavatar', event: any): void
}>()
const isnvue = ref(false)
// #ifdef APP-NVUE
isnvue.value = true
// #endif
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	...custom_props,
	transprent: {
		type: Boolean,
		default: false
	},
	followTheme: {
		type: Boolean,
		default: true
	},
	hoverClass: {
		type: String,
		default: 'opacity-7'
	},
	/**
	 * mini,small,normal,middle,large
	 */
	size: {
		type: String as PropType<btnSize>,
		default: 'normal'
	},
	fontSize: {
		type: Number,
		default: 0
	},
	fontColor: {
		type: String,
		default: ''
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 16]
	},
	padding: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	//不是同层背景，默认是同层，为false
	//如果输入框表单与tmshee在同一层下。他们的黑白暗黑背景色是相同的。为了区分这个问题，需要单独指示，以便区分背景同层。
	//主意：它只在黑和白之间的色系才起作用，其它颜色下不起作用。
	noLevel: {
		type: Boolean,
		default: false
	},
	shadow: {
		type: Number,
		default: 0
	},
	width: {
		type: Number,
		default: 0
	},
	height: {
		type: Number,
		default: 0
	},
	//如果为true，采用块状flex布局，自动宽和高度。
	block: {
		type: Boolean,
		default: false
	},
	round: {
		type: Number,
		default: 0
	},
	loading: {
		type: Boolean,
		default: false
	},
	disabled: {
		type: Boolean,
		default: false
	},
	//提供时，点击后会中转到对应页面。
	url: {
		type: String,
		default: ''
	},
	label: {
		type: String,
		default: ''
	},
	icon: {
		type: String,
		default: ''
	},
	/**
	 * submit,reset 在tm-form中使用。
	 */
	formType: {
		type: String as PropType<'submit' | 'reset' | 'filterCancel' | 'filterConfirm'>,
		default: ''
	},
	//开放能力
	openType: {
		type: String,
		default: ''
	},
	//打开 APP 时，向 APP 传递的参数，open-type=launchApp时有效
	appParameter: {
		type: String,
		default: ''
	},
	sessionFrom: {
		type: String,
		default: ''
	},
	sendMessageTitle: {
		type: String,
		default: ''
	},
	sendMessagePath: {
		type: String,
		default: ''
	},
	sendMessageImg: {
		type: String,
		default: ''
	},
	sendMessageCard: {
		type: String,
		default: ''
	},
	groupId: {
		type: String,
		default: ''
	},
	guildId: {
		type: String,
		default: ''
	},
	publicId: {
		type: String,
		default: ''
	},
	unit: {
		type: String,
		default: 'rpx'
	},
	darkBgColor: {
		type: String,
		default: ''
	},
	/**禁用后的主题色 */
	disabledColor: {
		type: String,
		default: 'grey-1'
	}
})

/** -----------form专有------------ */
const formtype = computed(() => props.formType)
//父级方法。
let FormParent: any = null
let FilterParent: any = null
if (formtype.value == 'reset' || formtype.value == 'submit') {
	FormParent = proxy?.$parent
	while (FormParent) {
		if (FormParent?.tmFormComnameId == 'tmFormId' || !FormParent) {
			break
		} else {
			FormParent = FormParent?.$parent ?? undefined
		}
	}
}
//过滤器菜单 专用属性.
if (formtype.value == 'filterCancel' || formtype.value == 'filterConfirm') {
	FilterParent = proxy?.$parent
	while (FilterParent) {
		if (FilterParent?.FilterMenu == 'FilterMenu' || !FilterParent) {
			break
		} else {
			FilterParent = FilterParent?.$parent ?? undefined
		}
	}
}
/** -----------end------------ */
const _noLevel = computed(() => props.noLevel)
//是否暗黑模式。
const isDark = computed(() => computedDark(props, tmcfg.value))
// 设置响应式主题文字色。
let textColor = computed(() => {
	if (theme.isCssColor(_fontColor.value)) return _fontColor.value
	if (!props.fontColor) return tmcomputed.value.textColor
	return theme.getColor(props.fontColor).value
})
//自定义样式：
const customCSSStyle = computed(() => {
	return {
		// height: btnSizeObj.value.h + props.unit,
		...computedStyle(props)
		// fontSize:props.fontSize,
		// color:textColor,
		// border: "0px solid rgba(0, 0, 0, 0)",
		// background: "rgba(0, 0, 0, 0)",
		// borderRadius: "0px",
	}
})
//自定类
const customClass = computed(() => computedClass(props))
// 设置响应式全局组件库配置表。
const tmcfg = computed(() => store.tmStore)
//计算主题
const _shadow = computed(() => props.shadow || (store.tmuiConfig.themeConfig?.component?.button?.shadow ?? 0))
const tmcomputed = computed<cssstyle>(() => {
	let nprops = { ...props, shadow: _shadow.value }
	if (_disabled.value) {
		nprops.color = props.disabledColor
	}
	return computedTheme({ ...nprops }, isDark.value, tmcfg.value)
})

const isclickOn = ref(false)
const _load = computed(() => props.loading)
const _disabled = computed(() => props.disabled)
const _label = computed(() => props.label)
const _icon = computed(() => props.icon)
const _isBorder = computed(() => {
	if (props.outlined && props.border == 0) return 1
	if (props.border > 0) return props.border
	return 0
})
const sizeObj = computed(() => {
	let ptest = {
		block: { w: 0, h: 80, fontSize: 28, round: 3 },
		mini: { w: 88, h: 36, fontSize: 20, round: 3 },
		small: { w: 120, h: 56, fontSize: 22, round: 3 },
		normal: { w: 220, h: 80, fontSize: 28, round: 3 },
		middle: { w: 300, h: 80, fontSize: 30, round: 3 },
		large: { w: 375, h: 80, fontSize: 32, round: 3 }
	}
	if (props.unit == 'px') {
		let key: 'block' | 'mini' | 'small' | 'normal' | 'middle' | 'large' = 'block'
		let key2: 'w' | 'h' | 'fontSize' | 'round' = 'w'
		for (key in ptest) {
			for (key2 in ptest[key]) {
				ptest[key][key2] = uni.upx2px(ptest[key][key2])
			}
		}
	}
	return ptest
})
const btnSizeObj = computed(() => {
	let fontSize = props.fontSize || 0
	if (props.block) {
		return {
			w: 0,
			h: props.height || sizeObj.value.block.h,
			fontSize: fontSize || sizeObj.value.block.fontSize,
			round: props.round == -1 ? 0 : props.round || (store.tmuiConfig.themeConfig?.component?.button?.round ?? 0) || sizeObj.value.normal.round
		}
	}
	return {
		w: props.width || sizeObj.value[props.size].w,
		h: props.height || sizeObj.value[props.size].h,
		fontSize: fontSize || sizeObj.value[props.size].fontSize,
		round: props.round == -1 ? 0 : props.round || (store.tmuiConfig.themeConfig?.component?.button?.round ?? 0) || sizeObj.value[props.size].round
	}
})
const _fontColor = computed(() => props.fontColor)
const _transprent = computed(() => props.transprent)
/** 数组是左，上，右，下顺序。 */
const _margin = computed(() => {
	if (props.margin.length == 1) return [props.margin[0], props.margin[0], props.margin[0], props.margin[0]]
	if (props.margin.length == 2) return [props.margin[0], props.margin[1], props.margin[0], props.margin[1]]
	if (props.margin.length == 3) return [props.margin[0], props.margin[1], props.margin[2], 0]
	if (props.margin.length == 4) return [props.margin[0], props.margin[1], props.margin[2], props.margin[3]]
	return [0, 0, 0, 0]
})
const _padding = computed(() => {
	if (props.padding.length == 1) return [props.padding[0], props.padding[0], props.padding[0], props.padding[0]]
	if (props.padding.length == 2) return [props.padding[0], props.padding[1], props.padding[0], props.padding[1]]
	if (props.padding.length == 3) return [props.padding[0], props.padding[1], props.padding[2], 0]
	if (props.padding.length == 4) return [props.padding[0], props.padding[1], props.padding[2], props.padding[3]]
	return [0, 0, 0, 0]
})

const _bgcolor = computed(() => {
	if (_transprent.value === true) return `background-color:rgba(255,255,255,0);`
	if (props.darkBgColor !== '' && isDark.value === true) {
		return `background-color:${props.darkBgColor};`
	}
	if (props.linearColor.length == 2) {
		return {
			'background-image': `linear-gradient(${tmcomputed.value.linearDirectionStr},${props.linearColor[0]},${props.linearColor[1]})`
		}
	}
	if (tmcomputed.value.gradientColor?.length == 2) {
		return tmcomputed.value.backgroundColorCss
	}

	if (_noLevel.value && tmcomputed.value.isBlackAndWhite === true && isDark.value === true) {
		return `background-color: ${tmcomputed.value.inputcolor}`
	}
	return `background-color: ${tmcomputed.value.backgroundColor}`
})

function touchstart(e: Event) {
	isclickOn.value = true
	emits('touchstart', e)
}
function touchend(e: Event) {
	isclickOn.value = false
	emits('touchend', e)
}
function touchcancel(e:Event){
	isclickOn.value = false;
	emits('touchcancel', e)
}
function onclick(e: Event) {
	if (FormParent != null && typeof FormParent != 'undefined' && formtype.value && !props.loading) {
		FormParent[formtype.value]()
	}
	if (FilterParent != null && typeof FilterParent != 'undefined' && formtype.value && !props.loading) {
		FilterParent[formtype.value]()
	}
	// if button loading , not handle click
	if (props.loading) return

	emits('click', e)
	if (props.url !== '' && typeof props.url === 'string') {
		let url = props.url
		if (url[0] !== '/') url = '/' + url
		uni.navigateTo({
			url: url
		})
		return
	}
	if (props.openType == 'getUserInfo' || props.openType == 'getUserProfile') {
		// #ifdef MP-WEIXIN
		uni.getUserProfile({
			desc: '用于完善会员资料',
			success: function (userProfile) {
				if (userProfile.errMsg != 'getUserProfile:ok') {
					uni.showToast({
						title: userProfile.errMsg,
						icon: 'error',
						mask: true
					})
					return
				}
				emits('getUserInfo', userProfile)
				emits('getUserProfile', userProfile)
			},
			fail: (error) => {
				console.log(error)
				uni.showToast({
					icon: 'none',
					title: typeof error == 'object' ? error.errMsg : error
				})
			}
		})
		console.warn(
			"微信小程序已收回‘getUserProfile’以及‘getUserInfo’权限，请使用open-type='chooseAvatar'使用@chooseavatar回调，详见《微信小程序用户头像昵称获取规则调整公告》https://developers.weixin.qq.com/community/develop/doc/00022c683e8a80b29bed2142b56c01"
		)
		// #endif
	}
}

provide('appTextColor', textColor)
</script>

<style scoped>
.button {
	/* background: rgba(0, 0, 0, 0); */
	/* border: 0px solid rgba(0, 0, 0, 0); */
	padding: 0px;
	/* border-radius: 0px; */
}

.buttonHover {
	/* background: rgba(0, 0, 0, 0); */
}

.bhover {
	opacity: 0.7;
}
/* #ifdef H5 */
/* .webpc:hover {
  cursor: pointer;
  opacity: 0.7;
} */
/* #endif */
/* #ifndef APP-NVUE */
/* .bhover:hover {
  opacity: 0.7 !important;
} */
.button::after {
	background: transparent !important;
	background-color: transparent !important;
	border: none !important;
	border-radius: 0px !important;
}

/* #endif */
/* #ifdef MP */
.button.alipay {
	transition: opacity 0.7s;
}
/* #endif */
</style>
