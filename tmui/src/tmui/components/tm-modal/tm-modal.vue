<template>
	<view>
		<!-- #ifdef APP-NVUE -->
		<view @click="opens">
			<view :eventPenetrationEnabled="true"><slot name="trigger"></slot></view>
		</view>
		<!-- #endif -->
		<!-- #ifndef APP-NVUE -->
		<view @click="opens"><slot name="trigger"></slot></view>
		<!-- #endif -->
		<tm-overlay
			:zIndex="props.zIndex"
			ref="overlayAni"
			blur
			:duration="props.duration"
			@close="overclose"
			@open="OverLayOpen"
			@click="overlayClickFun"
			:align="align_rp"
			:overlayClick="false"
			v-model:show="_show"
			:teleport="props.teleport"
		>
			<tm-translate
				:reverse="reverse_rp"
				:width="anwidth"
				:height="anheight"
				ref="drawerANI"
				:auto-play="false"
				:name="aniname"
				:duration="props.duration"
			>
				<view
					@click.stop="$event.stopPropagation()"
					:style="[
						{ width: anwidth, height: anheight, boxSizing: 'border-box' },
						!props.transprent ? tmcomputed.borderCss : '',
						!props.transprent ? tmcomputed.backgroundColorCss : '',
						!props.transprent ? tmcomputed.shadowColor : '',

						customCSSStyle
					]"
					:class="[round_rp, 'flex flex-col overflow ', customClass]"
				>
					<view class="flex flex-row px-24" style="height: 44px" :class="[props.closeable ? 'flex-row-center-between' : 'flex-center']">
						<slot name="title">
							<tm-text
								:_style="props.titleStyle"
								:dark="props.dark"
								:followTheme="false"
								_class="text-overflow-1 text-weight-b text-size-m text-align-center"
								:label="props.title"
							></tm-text>
						</slot>
						<tm-icon v-if="closeable" _class="opacity-3" name="tmicon-times-circle-fill" :fontSize="32" @click="close"></tm-icon>
					</view>
					<scroll-view scroll-y :style="[props.height ? { height: contentHeight } : '']">
						<view class="px-32">
							<slot name="default">
								<tm-text
									:font-size="30"
									:dark="props.dark"
									:followTheme="false"
									:label="props.content"
									_style="line-height:46rpx"
								></tm-text>
							</slot>
						</view>
					</scroll-view>
					<view class="flex flex-row" :class="[props.splitBtn ? 'pa-32' : '']">
						<slot name="button">
							<view v-if="!props.hideCancel" class="flex-1 overflow" style="height: 90rpx">
								<tm-sheet
									:dark="props.dark"
									:followTheme="true"
									:isDisabledRoundAndriod="true"
									@click="cancel"
									:height="90"
									:linear="props.cancelLinear"
									:linearDeep="props.cancelLlinearDeep"
									text
									:color="props.cancelColor"
									:_class="['flex-center overflow flex']"
									:paren-class="props.splitBtn ? 'round-' + props.btnRound : 'round-bl-' + props.round"
									:margin="[0, 0]"
									:padding="[0, 0]"
								>
									<tm-text
										_class="text-weight-b"
										_style="line-height:90rpx"
										:dark="props.dark"
										:followTheme="false"
										:userInteractionEnabled="false"
										:label="props.cancelText"
									></tm-text>
								</tm-sheet>
							</view>
							<view v-if="props.splitBtn && !props.hideCancel" class="overflow" style="width: 24rpx"></view>
							<view class="flex-1 flex" :class="[okLoading ? 'opacity-5' : '', 'overflow']" style="height: 90rpx">
								<tm-sheet
									paretClass="flex-1"
									class="flex-1"
									:dark="props.dark"
									:followTheme="true"
									:isDisabledRoundAndriod="true"
									@click="ok"
									:height="90"
									:linear="props.okLinear"
									:linearDeep="props.okLlinearDeep"
									:color="props.okColor"
									:margin="[0, 0]"
									:_class="['flex-center overflow']"
									:paren-class="props.splitBtn ? 'round-' + props.btnRound : 'round-br-' + props.round"
									:padding="[0, 0]"
								>
									<view :userInteractionEnabled="false" class="flex flex-row">
										<view v-if="okLoading" class="pr-10"
											><tm-icon :userInteractionEnabled="false" name="tmicon-shuaxin" spin></tm-icon
										></view>
										<tm-text
											_class="text-weight-b"
											:dark="props.dark"
											:userInteractionEnabled="false"
											:label="props.okText"
										></tm-text>
									</view>
								</tm-sheet>
							</view>
						</slot>
					</view>
				</view>
			</tm-translate>
		</tm-overlay>
	</view>
</template>

<script lang="ts" setup>
/**
 * 对话框
 * @description 对话框
 */
import tmTranslate from '../tm-translate/tm-translate.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmOverlay from '../tm-overlay/tm-overlay.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import { getCurrentInstance, computed, ref, provide, inject, onMounted, onUnmounted, nextTick, watch, ComponentInternalInstance, toRaw } from 'vue'
import { cssstyle, tmVuetify, colorThemeType } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'

const drawerANI = ref<InstanceType<typeof tmTranslate> | null>(null)
const overlayAni = ref<InstanceType<typeof tmOverlay> | null>(null)

const store = useTmpiniaStore()
const props = defineProps({
	...custom_props,
	//是否显示遮罩
	mask: {
		type: [Boolean],
		default: false
	},

	border: {
		type: Number,
		default: 2
	},
	show: {
		type: [Boolean],
		default: false
	},
	width: {
		type: Number,
		default: 600
	},
	height: {
		type: Number,
		default: 450
	},
	round: {
		type: Number,
		default: 10
	},
	//弹出的动画时间单位ms.
	duration: {
		type: Number,
		default: 250
	},
	//是否允许点击遮罩关闭
	overlayClick: {
		type: Boolean,
		default: true
	},
	transprent: {
		type: [Boolean],
		default: false
	},
	//如果显示关闭。标题栏被替换为左标题右关闭按钮。
	closeable: {
		type: [Boolean],
		default: false
	},
	color: {
		type: String,
		default: 'white'
	},
	title: [String],
	okText: {
		type: [String],
		default: '完成'
	},
	okColor: {
		type: [String],
		default: 'primary'
	},
	okLinear: {
		type: [String],
		default: '' //left:右->左，right:左->右。top:下->上，bottom:上->下。
	},
	// 渐变的亮浅
	okLlinearDeep: {
		type: [String],
		default: 'accent' //light,dark,亮系渐变和深色渐变。
	},
	cancelColor: {
		type: [String],
		default: 'primary'
	},
	cancelText: {
		type: [String],
		default: '取消'
	},
	cancelLinear: {
		type: [String],
		default: '' //left:右->左，right:左->右。top:下->上，bottom:上->下。
	},
	// 渐变的亮浅
	cancelLlinearDeep: {
		type: [String],
		default: 'accent' //light,dark,亮系渐变和深色渐变。
	},
	//只有在分享式按钮下才有作用。
	btnRound: {
		type: Number,
		default: 24
	},
	hideCancel: {
		type: [Boolean],
		default: false
	},
	//分离式按钮。
	splitBtn: {
		type: Boolean,
		default: false
	},
	//在关闭前执行的回调函数。返回false时即取消关闭窗体。在app端返回的是true,而非app是function,需要再次执行
	beforeClose: {
		type: Function,
		default: () => {
			return () => true
		}
	},
	content: {
		type: String,
		default: ''
	},
	disabled: {
		type: Boolean,
		default: false
	},
	titleStyle: {
		type: [Array, String, Object],
		default: () => []
	},
	zIndex: {
		type: [Number, String],
		default: 999
	},
	/** 是否使用teleport */
	teleport: {
		type: Boolean,
		default: true
	}
})
const emits = defineEmits(['click', 'open', 'close', 'update:show', 'ok', 'cancel'])
const proxy = getCurrentInstance()?.proxy ?? null
// 设置响应式全局组件库配置表。
const tmcfg = computed<tmVuetify>(() => store.tmStore)
//自定义样式：
const customCSSStyle = computed(() => computedStyle(props))
//自定类
const customClass = computed(() => computedClass(props))
//是否暗黑模式。
const isDark = computed(() => computedDark(props, tmcfg.value))
//计算主题
const tmcomputed = computed<cssstyle>(() => computedTheme(props, isDark.value, tmcfg.value))
let rejCall: any = null
let resCall: any = null
let isOkClose = false
//外部调用open或者close时将会执行相关参数函数.
let nowCallFun: any = null
const reverse = ref(true)
let flag = false
let timeid = uni.$tm.u.getUid(4)
const okLoading = ref(false)
let _show = ref(props.show)
let timerId = NaN
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
		}, wait)
	}
}

function throttle(func: Function, wait = 500, immediate = true) {
	if (immediate) {
		if (!flag) {
			flag = true
			// 如果是立即执行，则在wait毫秒内开始时执行
			typeof func === 'function' && func()
			timeid = setTimeout(() => {
				flag = false
			}, wait)
		}
	} else {
		if (!flag) {
			flag = true
			// 如果是非立即执行，则在wait毫秒内的结束处执行
			timeid = setTimeout(() => {
				flag = false
				typeof func === 'function' && func()
			}, wait)
		}
	}
}

watch(
	() => props.show,
	(val) => {
		if (val) {
			opens()
		} else {
			close()
		}
	}
)
onMounted(() => {
	if (_show.value) {
		opens()
	}
})

const round_rp = computed(() => {
	return 'round-' + props.round
})
const reverse_rp = computed(() => {
	if (aniname.value != 'zoom') return reverse.value
	return !reverse.value
})
const aniname = computed(() => {
	return 'zoom'
})
const anwidth = computed(() => {
	return props.width + 'rpx'
})
const anheight = computed(() => {
	return props.height + 'rpx'
})
const contentHeight = computed(() => {
	let bas = 0
	if (props.splitBtn) {
		bas = uni.upx2px(64)
	}
	return uni.upx2px(props.height) - 44 - uni.upx2px(90) - bas + 'px'
})
const align_rp = computed(() => {
	return 'flex-center'
})

async function ok() {
	if (props.disabled) return
	debounce(
		async () => {
			try {
				if (typeof props.beforeClose === 'function') {
					okLoading.value = true
					let p = await props.beforeClose()
					if (typeof p === 'function') {
						p = await p()
					}
					okLoading.value = false
					if (!p) return
				}
				emits('ok', toRaw(nowCallFun))
				isOkClose = true
				close()
			} catch (e) {
				okLoading.value = false
			}
		},
		250,
		true
	)
}

function cancel() {
	if (props.disabled) return
	if (okLoading.value) return
	isOkClose = false
	emits('cancel', toRaw(nowCallFun))

	close()
}

function OverLayOpen() {
	nextTick(function () {
		drawerANI.value?.play()
	})
	emits('open')
	emits('update:show', true)
	_show.value = true
}

function overclose() {
	nextTick(() => {
		emits('close')
		emits('update:show', false)
		_show.value = false
		if (resCall && isOkClose) {
			resCall()
		}
		if (rejCall && !isOkClose) {
			rejCall()
		}
	})
}

// 外部调用
function open(arg: any = null) {
	if (okLoading.value) return
	if (_show.value == true) return
	return new Promise((res, rej) => {
		throttle(
			async () => {
				reverse.value = true
				overlayAni.value?.open(true)
				nowCallFun = arg

				rejCall = rej
				resCall = res
			},
			props.duration,
			true
		)
	})
}

// 内部调用
function opens() {
	if (props.disabled) return
	if (okLoading.value) return
	if (_show.value == true) return
	debounce(
		() => {
			reverse.value = true
			_show.value = true
			overlayAni.value?.open(true)
		},
		props.duration,
		true
	)
}

//外部手动调用关闭方法
async function close(arg: Function | null = null) {
	reverse.value = false
	if (!drawerANI.value) return
	overlayAni.value?.close()
	drawerANI.value?.play()
}
//内部通过点击遮罩关闭的方法.同时需要发送事件。
function overlayClickFun(e: Event) {
	if (_show.value == false) return
	emits('click', e)
	if (!props.overlayClick || props.disabled || okLoading.value || !overlayAni.value) return
	reverse.value = false
	throttle(
		() => {
			isOkClose = false
			overlayAni.value?.close()
			drawerANI.value?.play()
		},
		props.duration,
		true
	)
}

//外部调用的方法。
defineExpose({
	close: close,
	open: open
})
</script>

<style scoped>
.flex-left-custom {
	display: flex;
	justify-content: flex-start;
	align-items: flex-start;
}

.flex-right-custom {
	display: flex;
	justify-content: flex-start;
	align-items: flex-end;
}

.flex-top-custom {
	display: flex;
	justify-content: flex-start;
	align-items: flex-start;
}

.flex-end-custom {
	display: flex;
	justify-content: flex-end;
	align-items: flex-end;
}

.flex-center-custom {
	display: flex;
	justify-content: center;
	align-items: center;
	flex-direction: row;
}
</style>
