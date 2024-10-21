<template>
	<view>
		<!-- #ifdef APP-NVUE -->
		<view @click="open"
			><view :eventPenetrationEnabled="true"><slot name="trigger"></slot></view
		></view>
		<!-- #endif -->
		<!-- #ifndef APP-NVUE -->
		<view @click="open"><slot name="trigger"></slot></view>
		<!-- #endif -->
 
		<tm-overlay
			:inContent="props.inContent"
			ref="overlayAni"
			:duration="props.duration + 80"
			@open="OverLayOpen"
			@close="overclose"
			:zIndex="props.zIndex"
			:transprent="!props.mask"
			@click="overlayClickFun"
			:align="align_rp"
			:teleport="props.teleport"
			:overlayClick="false"
			v-model:show="_show"
		>
			<tm-translate
				:reverse="reverse_rp"
				:width="anwidth"
				:height="anheight"
				ref="drawerANI"
				:auto-play="false"
				:name="aniname"
				:duration="props.duration+80"
				@end="playEndEvent"
			>
				<view
					@click.stop="$event.stopPropagation()"
					:style="[
						{ width: anwidth, height: anheight },
						!props.transprent ? tmcomputed.borderCss : '',
						!props.transprent ? tmcomputed.backgroundColorCss : '',
						!props.transprent ? tmcomputed.shadowColor : '',
						customCSSStyle
					]"
					:class="[round_rp, 'flex flex-col overflow ', customClass]"
				>	
					
					<view
						v-if="!props.closeable && !props.hideHeader"
						class="flex flex-row flex-row-center-center flex-between px-24"
						style="height: 44px"
					>
						<view class="flex-4 flex-shrink">
							<tm-text v-if="!props.hideCancel && !loading && !ok_loading" @click="cancel" :label="props.cancelText"></tm-text>
						</view>
						<view class="flex-8 px-32 flex-center">
							<slot name="title"><tm-text _class="text-overflow-1 opacity-7" :label="props.title"></tm-text></slot>
						</view>
						<view class="flex-4 flex-shrink flex-row flex-row-center-end">
							<tm-text :color="okColor" @click="ok" v-if="!ok_loading && !loading" :dark="props.dark" :label="props.okText"></tm-text>
							<tm-icon
								:color="okColor"
								v-if="ok_loading || loading"
								:spin="ok_loading || loading"
								:dark="isDark"
								:_class="isDark !== true ? 'opacity-4' : ''"
								:fontSize="24"
								:name="ok_loading || loading ? 'tmicon-shuaxin' : 'tmicon-times-circle-fill'"
							></tm-icon>
						</view>
					</view>
					<view
						v-if="props.closeable && !props.hideHeader"
						class="flex flex-row flex-row-center-center flex-between px-24"
						style="height: 44px"
					>
						<view class="flex-9 pr-32">
							<slot name="title"><tm-text _class="text-overflow-1 opacity-7" :dark="props.dark" :label="props.title"></tm-text></slot>
						</view>
						<view class="flex-3 flex-shrink flex-row flex-row-center-end">
							<tm-icon
								@click="cancel"
								:dark="props.dark"
								:_class="isDark !== true ? 'opacity-3' : ''"
								:fontSize="36"
								name="tmicon-times-circle-fill"
							></tm-icon>
						</view>
					</view>
					<!-- #ifdef APP-NVUE -->
					<scroll-view :scroll-y="!props.disabbleScroll" :style="[{ height: contentHeight }]" class="overflow"
						><slot name="default"></slot
					></scroll-view>
					<!-- #endif -->
					<!-- #ifndef APP-NVUE -->
					<view
						:style="{
							overflowY: props.disabbleScroll ? 'normal' : 'auto',
							height: contentHeight
						}"
						><slot name="default"></slot
					></view>
					<!-- #endif -->
					<slot name="foot"><view v-if="props.footHeight > 0" class="flex"></view></slot>
				</view>
			</tm-translate>
		</tm-overlay>
	</view>
</template>

<script lang="ts" setup>
/**
 * 抽屉
 * @description 别名poup弹层，提供，左，右，上，下，中弹出内容。
 */
import tmTranslate from '../tm-translate/tm-translate.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmOverlay from '../tm-overlay/tm-overlay.vue'
import { getCurrentInstance, computed, ref, provide, inject, onMounted, onUnmounted, nextTick, watch } from 'vue'
import { cssstyle, tmVuetify, colorThemeType } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
import { useWindowInfo } from '../../tool/useFun/useWindowInfo'

const drawerANI = ref<InstanceType<typeof tmTranslate> | null>(null)
const overlayAni = ref<InstanceType<typeof tmOverlay> | null>(null)

const store = useTmpiniaStore()
const props = defineProps({
	...custom_props,
	//是否显示遮罩
	mask: {
		type: [Boolean, String],
		default: true
	},
	//抽屉放置的位置
	placement: {
		type: String,
		default: 'bottom' //top|left|right|bottom|center
	},
	show: {
		type: [Boolean],
		default: false
	},
	width: {
		type: Number,
		default: 500
	},
	height: {
		type: Number,
		default: 600
	},
	round: {
		type: Number,
		default: 10
	},
	//弹出的动画时间单位ms.
	duration: {
		type: Number,
		default: 300
	},
	//是否允许点击遮罩关闭
	overlayClick: {
		type: Boolean,
		default: true
	},
	transprent: {
		type: [Boolean, String],
		default: false
	},
	//如果显示关闭。标题栏被替换为左标题右关闭按钮。
	closeable: {
		type: [Boolean, String],
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
	//true时，确认按钮将出现加载状态。
	okLoading: {
		type: [Boolean, String],
		default: false
	},
	cancelText: {
		type: [String],
		default: '取消'
	},
	hideCancel: {
		type: [Boolean, String],
		default: false
	},
	//隐藏工具栏，标题，取消，确认
	hideHeader: {
		type: [Boolean, String],
		default: false
	},
	disabled: {
		type: Boolean,
		default: false
	},
	zIndex: {
		type: [Number, String],
		default: 401
	},
	unit: {
		type: String,
		default: 'rpx'
	},
	disabbleScroll: {
		type: Boolean,
		default: false
	},
	/** 是否嵌入弹层，开启后将在它的父组件内执行弹层。 */
	inContent: {
		type: Boolean,
		default: false
	},
	footHeight: {
		type: Number,
		default: 0
	},
	/** 是否使用teleport */
	teleport: {
		type: Boolean,
		default: true
	},
	/**打开前执行 */
	beforeOpen: {
		type: Function,
		default: null
	},
	/**关点击ok前执行，如果返回是false，将阻止关闭. */
	beforeOk: {
		type: Function,
		default: null
	},
	/**点击取消前执行，如果返回fase将阻止关闭. */
	beforeCance: {
		type: Function,
		default: null
	}
})
const emits = defineEmits(['click', 'open', 'close', 'update:show', 'ok', 'cancel'])
const proxy = getCurrentInstance()?.proxy ?? null
const sysinfo = useWindowInfo();
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

const reverse = ref(true)
const timeid = ref(0)
let timerId: any = NaN
let timerIdth: any = NaN
let timerIdth_flas = false
uni.hideKeyboard()
let _show = ref(props.show)
const isPlaying = ref(false)
function debounce(func: Function, wait = 500, immediate = false) {
	// 清除定时器
	if (!isNaN(timerId)) clearTimeout(timerId)
	// 立即执行，此类情况一般用不到

	if (immediate) {
		timerId = setTimeout(() => {
			timerId = NaN
		}, wait)
		typeof func === 'function' && func()
	} else {
		// 设置定时器，当最后一次操作后，timeout不会再被清除，所以在延时wait毫秒后执行func回调方法
		timerId = setTimeout(() => {
			typeof func === 'function' && func()
		}, wait)
	}
}

function throttle(func: Function, wait = 500, immediate = true) {
	if (immediate) {
		if (!timerIdth_flas) {
			timerIdth_flas = true
			// 如果是立即执行，则在wait毫秒内开始时执行
			typeof func === 'function' && func()

			timerIdth = setTimeout(() => {
				timerIdth_flas = false
			}, wait)
		}
	} else {
		if (!timerIdth_flas) {
			timerIdth_flas = true
			// 如果是非立即执行，则在wait毫秒内的结束处执行
			timerIdth = setTimeout(() => {
				timerIdth_flas = false
				typeof func === 'function' && func()
			}, wait)
		}
	}
}

timeid.value = uni.$tm.u.getUid(4)
if (_show.value) {
	reverse.value = false
}
watch(
	() => props.show,
	(val) => {
		if (val) {
			reverse.value = true
		} else {
			reverse.value = false
			overlayAni.value?.close()
		}
		if (_show.value !== props.show) {
			nextTick(() => {
				drawerANI.value?.play()
				isPlaying.value=true;
			})
		}
		_show.value = props.show
	}
)
onMounted(() => {
	if (_show.value) {
		open()
	}
})
const ok_loading = computed(() => props.okLoading)
const loading = ref(false)
const round_rp = computed(() => {
	if (aniname.value == 'left') return 'round-r-' + props.round
	if (aniname.value == 'right') return 'round-l-' + props.round
	if (aniname.value == 'up') return 'round-b-' + props.round
	if (aniname.value == 'down') return 'round-t-' + props.round
	if (aniname.value == 'zoom') return 'round-' + props.round
})
const reverse_rp = computed(() => {
	if (aniname.value != 'zoom') return reverse.value
	return !reverse.value
})
const aniname = computed(() => {
	if (props.placement == 'center') return 'zoom'
	if (props.placement == 'top') return 'up'
	if (props.placement == 'bottom') return 'down'
	return props.placement
})
const anwidth = computed(() => {
	if (aniname.value == 'zoom') {
		return props.width + props.unit
	}
	if (props.placement == 'left' || props.placement == 'right') {
		return props.width + props.unit
	}
	return sysinfo.width + 'px'
})
const anheight = computed(() => {
	let wucha = 0
	if (props.placement == 'top' || props.placement == 'bottom' || aniname.value == 'zoom') {
		return props.height + wucha + props.unit
	}
	return sysinfo.height + 'px'
})
const contentHeight = computed(() => {
	let base_height = props.hideHeader ? 0 : 44
	let _footerHeight = uni.$tm.u.topx(props.footHeight)
	if (props.placement == 'top' || props.placement == 'bottom' || aniname.value == 'zoom') {
		let h = props.height
		if (props.unit == 'rpx') {
			h = uni.upx2px(props.height)
		}
		return h - base_height - _footerHeight + 'px'
	}
	return sysinfo.height - base_height - _footerHeight + 'px'
})
const align_rp = computed(() => {
	if (aniname.value == 'down') {
		return 'flex-col-bottom-center'
	}
	if (aniname.value == 'up') {
		return 'flex-top-custom'
	}
	if (aniname.value == 'left') {
		return 'flex-row-top-start'
	}
	if (aniname.value == 'right') {
		return 'flex-row-bottom-start'
	}
	if (aniname.value == 'zoom') {
		return 'flex-center'
	}
})

async function _beforeOpenFun() {
	if (typeof props.beforeOpen === 'function') {
		loading.value = true
		let p = await props.beforeOpen()
		if (typeof p === 'function') {
			p = await p()
		}
		loading.value = false
	}
}

async function _beforeOkFun() {
	let p = true
	if (typeof props.beforeOk === 'function') {
		loading.value = true
		p = await props.beforeOk()
		if (typeof p === 'function') {
			p = await p()
		}
		loading.value = false
		if (!p) return
	}
	return p
}

async function _beforeCancelFun() {
	let p = true
	if (typeof props.beforeCance === 'function') {
		loading.value = true
		p = await props.beforeCance()
		if (typeof p === 'function') {
			p = await p()
		}
		loading.value = false
		if (!p) return
	}
	return p
}

function OverLayOpen() {
	// nextTick(function() {
	// 	drawerANI.value?.play();
	// })
	_beforeOpenFun()
	_show.value = true
	emits('open')
	emits('update:show', true)
}
function overclose() {
	nextTick(() => {
		_show.value = false
		emits('close')
		emits('update:show', false)
	})
}
async function overlayClickFun(e: Event) {
	
	if (!props.overlayClick || props.disabled || !overlayAni.value || loading.value||isPlaying.value) return
	emits('click', e)
	if (!(await _beforeCancelFun())) return
	reverse.value = false
	throttle(
		async () => {
			emits('cancel')
			overlayAni.value?.close()
			drawerANI.value?.play()
			
		},
		props.duration + 80,
		true
	)
}
function playEndEvent(){
	isPlaying.value = false
}


async function ok() {
	if (props.disabled || loading.value) return
	if (!(await _beforeOkFun())) return
	reverse.value = false
	debounce(
		() => {
			emits('ok')
			overlayAni.value?.close()
			drawerANI.value?.play()
		},
		500,
		true
	)
}

async function cancel() {
	if (props.disabled || loading.value) return
	if (!(await _beforeCancelFun())) return
	reverse.value = false
	debounce(
		() => {
			emits('cancel')
			overlayAni.value?.close()
			drawerANI.value?.play()
		},
		500,
		true
	)
}

//外部调用。
function open() {
	reverse.value = true
	_show.value = true
	nextTick(() => {
		drawerANI.value?.play()
	})
}

//外部手动调用关闭方法
function close() {
	reverse.value = false
	overlayAni.value?.close()
	drawerANI.value?.play()
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
