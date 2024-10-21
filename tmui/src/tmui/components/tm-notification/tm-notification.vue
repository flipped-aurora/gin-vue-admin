<template>
	<view
		v-if="showDom"
		:class="['fixed scale']"
		:style="[
			pos.left !== null ? { left: pos.left + 'px' } : '',
			pos.right !== null ? { right: pos.right + 'px' } : '',
			pos.top !== null ? { top: pos.top + 'px' } : '',
			pos.bottom !== null ? { bottom: pos.bottom + 'px' } : '',
			pos.width !== null ? { width: pos.width + 'px' } : '',
			props.shadow ? { padding: props.shadow * 4 + 'rpx' } : ''
		]"
	>
		<tm-sheet
			@click="emits('click', $event)"
			:color="color_com"
			:_class="_class"
			:_style="_style"
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
			:margin="props.margin"
			:padding="props.padding"
		>
			<slot>
				<view class="flex flex-row flex-row-center-between relative">
					<view class="flex flex-1 flex-row overflow flex-row-center-start">
						<tm-icon _class="pr-10" :fontSize="26" :name="icon_str"></tm-icon>
						<slot>
							<tm-text _class="text-overflow-1" :label="label_str"></tm-text>
						</slot>
					</view>
					<view class="pl-24 pr-12 flex flex-center" style="width: 0rpx">
						<tm-icon @click="hide" :fontSize="24" name="tmicon-times"></tm-icon>
					</view>
				</view>
			</slot>
		</tm-sheet>
	</view>
</template>

<script lang="ts" setup>
/**
 * 通知提醒
 * @description 通知提醒,总共四个角和上下，6个位置的提醒,使用时请注意内容变动即可显示。如果想一开始不想显示，不要提供内容就行。
 */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmTranslate from '../tm-translate/tm-translate.vue'
import { custom_props } from '../../tool/lib/minxs'
import { getCurrentInstance, computed, ref, provide, inject, onUpdated, onMounted, onUnmounted, nextTick, watch, PropType, watchEffect } from 'vue'
import { showOpts } from './interface'
const emits = defineEmits(['click', 'close'])
const proxy = getCurrentInstance()?.proxy ?? null
const tranmatioan = ref<InstanceType<typeof tmTranslate> | null>(null)
const props = defineProps({
	...custom_props,
	followTheme: {
		type: [Boolean],
		default: true
	},
	transprent: {
		type: [Boolean],
		default: false
	},
	border: {
		type: [Number],
		default: 0
	},
	round: {
		type: [Number],
		default: 2
	},
	shadow: {
		type: [Number],
		default: 0
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	padding: {
		type: Array as PropType<Array<number>>,
		default: () => [24, 16]
	},
	//多少秒后消失。0表示永不消失。
	duration: {
		type: Number,
		default: 2000
	},
	offset: {
		type: Array as PropType<Array<number>>,
		default: () => [32, 32] //x,y
	},
	//位置
	placement: {
		type: String,
		default: 'topLeft' //topLeft|topRight|bottomLeft|bottomRight|top|bottom
	},
	label: {
		type: String,
		default: ''
	},
	icon: {
		type: String,
		default: 'tmicon-info-circle-fill'
	}
})
const sysinfo = inject(
	'tmuiSysInfo',
	computed(() => {
		return {
			bottom: 0,
			height: 750,
			width: uni.upx2px(750),
			top: 0,
			isCustomHeader: false,
			sysinfo: null
		}
	})
)
let windowBottom = computed(() => sysinfo.value.bottom)
let windowTop = computed(() => sysinfo.value.top)
let windowWidth = computed(() => sysinfo.value.width)
let uid: number = NaN
const showDom = ref(false)
const label_str = ref(props.label)
const icon_str = ref(props.icon)
const pos = computed(() => {
	if (props.placement == 'topLeft') {
		return {
			top: windowTop.value + uni.upx2px(props.offset[1]),
			left: uni.upx2px(props.offset[0]),
			right: null,
			bottom: null,
			width: null
		}
	}
	if (props.placement == 'topRight') {
		return {
			top: windowTop.value + uni.upx2px(props.offset[1]),
			left: null,
			right: uni.upx2px(props.offset[0]),
			bottom: null,
			width: null
		}
	}
	if (props.placement == 'bottomLeft') {
		return {
			top: null,
			left: uni.upx2px(props.offset[0]),
			right: null,
			bottom: windowBottom.value + uni.upx2px(props.offset[1]),
			width: null
		}
	}
	if (props.placement == 'bottomRight') {
		return {
			top: null,
			left: null,
			right: uni.upx2px(props.offset[0]),
			bottom: windowBottom.value + uni.upx2px(props.offset[1]),
			width: null
		}
	}
	if (props.placement == 'top') {
		return {
			top: windowTop.value + uni.upx2px(props.offset[1]),
			left: uni.upx2px(props.offset[0]),
			right: null,
			bottom: null,
			width: windowWidth.value - uni.upx2px(props.offset[0]) * 2
		}
	}
	if (props.placement == 'bottom') {
		return {
			top: null,
			left: uni.upx2px(props.offset[0]),
			right: null,
			bottom: windowBottom.value + uni.upx2px(props.offset[1]),
			width: windowWidth.value - uni.upx2px(props.offset[0]) * 2
		}
	}
	return {
		left: null,
		right: null,
		bottom: null,
		width: null,
		top: null
	}
})
const color_com = ref(props.color)
watchEffect(() => {
	color_com.value = props.color
})
onMounted(() => {
	label_str.value = props.label
	icon_str.value = props.icon
})
onUnmounted(() => {
	clearTimeout(uid)
})
//手动显示
function show(arg: showOpts) {
	let { icon, label, duration } = arg || {}
	label_str.value = label || props.label || ''
	icon_str.value = icon || props.icon || ''
	color_com.value = (arg?.color ?? color_com.value) || color_com.value
	duration = typeof duration === 'undefined' ? props.duration || 0 : duration
	if (showDom.value || !isNaN(uid)) {
		showDom.value = false
		clearTimeout(uid)
		nextTick(() => {
			showDom.value = true
			if (!duration) return
			uid = setTimeout(function () {
				showDom.value = false
			}, duration)
		})
	} else {
		showDom.value = true
		if (!duration) return
		uid = setTimeout(function () {
			showDom.value = false
		}, duration)
	}
}
//手动隐藏。
function hide() {
	showDom.value = false
	clearTimeout(uid)
	emits('close')
}

defineExpose({ show: show, hide: hide })
</script>

<style scoped>
.scale {
	/* #ifndef APP-NVUE */
	animation: aniscale 0.3s;
	z-index: 500 !important;
	/* #endif */
}
/* #ifndef APP-NVUE */
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
