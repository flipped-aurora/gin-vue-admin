<template>
	<view
		ref="nvueElAni"
		v-if="showValue"
		class="fixed l-0 t-0 flex flex-row flex-row-center-center on"
		:class="[showMask ? 'overflowMask' : 'overflowMaskNo']"
		:style="{ width: sysinfo.width + 'px', height: sysinfo.height + 'px', top: 0 + 'px' }"
	>
		<tm-sheet
			v-if="showValue"
			:style="{ transform: isNvue ? 'scale(0,0)' : 'scale(1,1)' }"
			ref="nvueElAniContent"
			class="scale nvueContent"
			:_style="props._style"
			:_class="props._class"
			:color="bgColor"
			:border="0"
			:shadow="10"
			:width="props.width"
			:height="props.height"
			:margin="[40, 40]"
			:round="props.round"
			:padding="props.padding"
		>
			<slot>
				<view class="flex flex-center flex-col ma-30" style="line-height: normal">
					<tm-icon
						_style="line-height: normal"
						style="line-height: normal"
						_class="pa-10"
						:spin="model_ref == 'load'"
						:color="color_ref"
						:fontSize="72"
						:name="icon_ref"
					></tm-icon>
					<tm-text :font-size="30" :_class="`pt-8 text-overflow-${lines_ref}`" :label="text_ref"></tm-text>
				</view>
			</slot>
		</tm-sheet>
	</view>
</template>

<script lang="ts" setup>
/**
 * 消息提示
 * @description 消息提示，属于全局阻断式提醒，会打断用户操作。
 */
import { language } from '../../tool/lib/language'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmTranslate from '../tm-translate/tm-translate.vue'
import tmOverlay from '../tm-overlay/tm-overlay.vue'
import { config, modelType } from './interface'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
import { getCurrentInstance, computed, ref, provide, inject, ComponentInternalInstance, onUnmounted, nextTick, watch, Ref, PropType } from 'vue'
// #ifdef APP-PLUS-NVUE
const Binding = uni.requireNativePlugin('bindingx')
const dom = uni.requireNativePlugin('dom')
const animation = uni.requireNativePlugin('animation')
// #endif
const store = useTmpiniaStore()
const tranAni = ref<InstanceType<typeof tmTranslate> | null>(null)
const Overlay = ref<InstanceType<typeof tmOverlay> | null>(null)
const emits = defineEmits(['click'])
const proxy = getCurrentInstance()?.proxy ?? null
import { useWindowInfo } from '../../tool/useFun/useWindowInfo'
const props = defineProps({
	//自定义的样式属性
	_style: {
		type: [Array, String, Object],
		default: () => {}
	},
	round: {
		type: Number,
		default: 12
	},
	padding: {
		type: Array as PropType<Array<number>>,
		default: () => [24, 0]
	},
	//自定义类名
	_class: {
		type: [Array, String],
		default: 'flex-center'
	},
	//是否显示遮罩
	mask: {
		type: [Boolean],
		default: true
	},
	//自动关闭时,需要显示多久关闭,单位ms
	duration: {
		type: Number,
		default: 1500
	},
	width: {
		type: Number,
		default: 300
	},
	height: {
		type: Number,
		default: 300
	},
	lines: {
		type: Number,
		default: 1
	}
})

const sysinfo = useWindowInfo()

const _lines = computed(() => props.lines)
let isNvue = ref(false)
// #ifdef APP-NVUE
isNvue.value = true
// #endif

const dur = ref(props.duration)
let uid: number = NaN
const bgColor = ref('white')
const model_ref: Ref<modelType> = ref('info')
const showValue = ref(false)
const icon_ref = ref('')
const text_ref = ref('')
const color_ref = ref('')
const showMask = ref(props.mask)
const dark_ref = ref(false)
const lines_ref = ref(_lines.value)

const modelIcon = computed(() => {
	return {
		load: {
			icon: 'tmicon-loading',
			color: 'primary',
			text: language('message.load.text')
		},
		error: {
			icon: 'tmicon-times-circle',
			color: 'red',
			text: language('message.error.text')
		},
		info: {
			icon: 'tmicon-info-circle',
			text: language('message.info.text'),
			color: 'black'
		},
		warn: {
			icon: 'tmicon-exclamation-circle',
			text: language('message.warn.text'),
			color: 'orange'
		},
		quest: {
			icon: 'tmicon-question-circle',
			text: language('message.quest.text'),
			color: 'pink'
		},
		success: {
			icon: 'tmicon-check-circle',
			text: language('message.success.text'),
			color: 'green'
		},
		disabled: {
			icon: 'tmicon-ban',
			text: language('message.disabled.text'),
			color: 'red'
		},
		wait: {
			icon: 'tmicon-ios-alarm',
			text: language('message.wait.text'),
			color: 'black'
		}
	}
})

function show(argFs: config) {
	if (showValue.value || !isNaN(uid)) {
		showValue.value = false
		clearTimeout(uid)
		nextTick(() => {
			showAction(argFs)
		})
	} else {
		showAction(argFs)
	}
}
function showAction(argFs: config) {
	//显示所需要的参数
	let arg = argFs || {}
	let { duration, icon, text, color, dark, model, mask, lines } = arg
	model_ref.value = typeof model == 'undefined' ? model_ref.value : model
	icon_ref.value = icon = icon ?? modelIcon.value[model_ref.value].icon
	text_ref.value = text = text ?? modelIcon.value[model_ref.value].text
	color_ref.value = color = color ?? modelIcon.value[model_ref.value].color
	showMask.value = typeof mask === 'boolean' ? mask : showMask.value
	lines_ref.value = lines ?? _lines.value

	if (dark === true) {
		bgColor.value = 'black'
	}
	if (typeof dark !== 'boolean') {
		dark = store.tmStore.dark
	}
	if (color_ref.value == 'white' || color_ref.value == 'black') {
		color_ref.value = ''
	}

	dark_ref.value = dark
	if (typeof duration === 'undefined') {
		duration = props.duration
	}
	dur.value = isNaN(parseInt(String(duration))) ? 1500 : parseInt(String(duration))
	showValue.value = true
	// #ifdef APP-NVUE
	setTimeout(function () {
		showNvueAniMation()
	}, 50)
	// #endif
	if (model_ref.value != 'load') {
		uid = setTimeout(function () {
			showValue.value = false
			uid = NaN
		}, dur.value)
	}
}

function showNvueAniMation() {
	var el = proxy.$refs.nvueElAni
	var elContent = proxy.$refs.nvueElAniContent
	animation.transition(
		el,
		{
			styles: {
				backgroundColor: 'rgba(0,0,0,0)',
				transformOrigin: 'center center'
			},
			duration: 1, //ms
			timingFunction: 'ease',
			delay: 0 //ms
		},
		() => {
			animation.transition(
				el,
				{
					styles: {
						backgroundColor: 'rgba(0,0,0,0.3)',
						transformOrigin: 'center center'
					},
					duration: 220, //ms
					timingFunction: 'ease',
					delay: 0 //ms
				},
				() => {}
			)
		}
	)
	animation.transition(
		elContent,
		{
			styles: {
				opacity: 1,
				transform: 'scale(1,1)',
				transformOrigin: 'center center'
			},
			duration: 220, //ms
			timingFunction: 'ease',
			delay: 0 //ms
		},
		() => {}
	)
}

//隐藏
function hide() {
	showValue.value = false
	uid = NaN
	clearTimeout(uid)
}
defineExpose({
	show: show,
	hide: hide
})
</script>

<style scoped>
.overflowMask {
	position: fixed;
	/* #ifndef APP-NVUE */
	z-index: 1000 !important;
	animation: aniover 0.5s;
	background-color: rgba(0, 0, 0, 0.3);
	transition-duration: 0.5s;
	transition-property: background, transform;
	transition-timing-function: ease;
	transition-delay: 0.3s;
	/* #endif */
	/* #ifdef APP-NVUE */
	background-color: rgba(0, 0, 0, 0);
	/* #endif */
}
.overflowMaskNo {
	/* #ifndef APP-NVUE */
	z-index: 1000 !important;
	background-color: rgba(0, 0, 0, 0);
	/* #endif */
}
.nvueContent {
	/* #ifdef APP-NVUE */
	opacity: 0;
	/* #endif */
}
.scale {
	/* #ifndef APP-NVUE */
	animation: aniscale 0.3s;
	/* #endif */
}
/* #ifndef APP-NVUE */
@keyframes aniover {
	0% {
		background-color: rgba(0, 0, 0, 0);
	}
	100% {
		background-color: rgba(0, 0, 0, 0.3);
	}
}
@keyframes aniscale {
	0% {
		opacity: 0;
		transform: scale(0.7, 0.7);
	}
	100% {
		opacity: 1;
		transform: scale(1, 1);
	}
}
/* #endif */
</style>
