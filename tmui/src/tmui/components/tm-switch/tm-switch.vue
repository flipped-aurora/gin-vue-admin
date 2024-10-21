<template>
	<tm-sheet
		@click="switchClick"
		:no-level="!_CheckVal"
		:followTheme="props.followTheme"
		:followDark="props.followDark"
		:dark="props.dark"
		:shadow="props.shadow"
		:outlined="props.outlined"
		:borderStyle="props.borderStyle"
		:borderDirection="props.borderDirection"
		:linearDeep="props.linearDeep"
		:linear="_CheckVal ? props.linear : ''"
		:round="viewSize.round"
		:color="_CheckVal ? props.color : props.unCheckedColor"
		:darkBgColor="_CheckVal ? '' : 'rgba(255,255,255,0.06)'"
		:height="viewSize.height"
		:width="viewSize.width"
		parenClass="switchbgani"
		:_class="['flex  relative flex-col', props.disabled ? 'opacity-4' : '']"
		:text="_CheckVal ? false : props.text"
		unit="px"
		:padding="[0, 0]"
		:margin="props.margin"
	>
		<view
			class="relative flex relative flex-col nvue"
			:style="{
				padding: '2px',
				width: `${viewSize.width}px`,
				height: `${viewSize.height}px`
			}"
		>
			<view
				:userInteractionEnabled="false"
				class="flex flex-row flex-between"
				:style="[{ width: viewSize.coenteWidth + 'px', height: viewSize.innerHeight + 'px' }]"
			>
				<view class="flex-1 flex-row flex-row-center-center"><tm-text :font-size="viewSize.fontSize" :label="props.label[0]"></tm-text></view>
				<view class="flex-1 flex-row flex-row-center-center"><tm-text :font-size="viewSize.fontSize" :label="props.label[1]"></tm-text></view>
			</view>
			<view
				:userInteractionEnabled="false"
				:class="['absolute base nvue', _CheckVal ? 'on' : 'off']"
				ref="switch"
				:style="{
					width: viewSize.innerWidth + 'px',
					height: viewSize.innerHeight + 'px'
				}"
				class="flex flex-col"
			>
				<tm-sheet
					:userInteractionEnabled="false"
					:padding="[0, 0]"
					:margin="[0, 0]"
					:height="viewSize.innerHeight"
					:width="viewSize.innerWidth"
					:color="props.barColor"
					:follow-dark="false"
					:round="viewSize.round"
					unit="px"
					_class="flex flex-center flex-row"
				>
					<tm-icon
						:followTheme="props.followTheme"
						v-if="_load"
						:font-size="viewSize.fontSize"
						:color="props.color"
						name="tmicon-shuaxin"
						spin
					></tm-icon>
					<tmTranslate name="zoom" v-if="!_load">
						<tm-icon
							v-if="props.barIcon && _CheckVal"
							:followTheme="props.followTheme"
							:font-size="viewSize.fontSize"
							:color="props.color"
							:name="props.barIcon"
						></tm-icon>
						<tm-icon
							v-if="props.offIcon && !_CheckVal"
							:followTheme="props.followTheme"
							:font-size="viewSize.fontSize"
							:name="props.offIcon"
							_class="opacity-5"
						></tm-icon>
					</tmTranslate>
				</tm-sheet>
			</view>
		</view>
	</tm-sheet>
</template>

<script lang="ts" setup>
/**
 * 开关
 * @description 便捷的选择框。可以加载中，异步开关。
 */
import { computed, ref, getCurrentInstance, onMounted, toRaw, watchEffect, watch, nextTick, inject, PropType } from 'vue'
import { inputPushItem, rulesItem } from './../tm-form-item/interface'
import { custom_props } from '../../tool/lib/minxs'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmTranslate from '../tm-translate/tm-translate.vue'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()

// #ifdef APP-PLUS-NVUE
const animation = uni.requireNativePlugin('animation')
// #endif
/**
 * 事件说明
 * update:modelValue v-model双向绑定数据。
 */
const emits = defineEmits(['update:modelValue', 'change', 'click'])
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	...custom_props,
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	transprent: {
		type: Boolean,
		default: false
	},
	defaultValue: {
		type: [Boolean, String, Number],
		default: false
	},
	modelValue: {
		type: [Boolean, String, Number],
		default: false
	},
	// 未选中时的值。
	unSelected: {
		type: [Boolean, String, Number],
		default: false
	},
	// 选中时的值。
	selected: {
		type: [Boolean, String, Number],
		default: true
	},
	width: {
		type: Number,
		default: 0
	},
	height: {
		type: Number,
		default: 0
	},
	size: {
		type: String,
		default: 'normal' //mini,normal,large
	},
	//激活后的主题色
	color: {
		type: String,
		default: 'primary'
	},
	//未激活的背景色
	unCheckedColor: {
		type: String,
		default: 'grey-3'
	},
	// 小圆球冒的背景色。
	barColor: {
		type: String,
		default: 'white'
	},
	round: {
		type: Number,
		default: 10
	},
	load: {
		type: Boolean,
		default: false
	},
	beforeChecked: {
		type: [Function, Boolean, String],
		default: () => false
	},
	/**
	 * 自定义打开成功后的bar上的小图标。
	 */
	barIcon: {
		type: String,
		default: 'tmicon-check'
	},
	/**
	 * 自定义关闭的小图标
	 */
	offIcon: {
		type: String,
		default: ''
	},
	disabled: {
		type: Boolean,
		default: false
	},
	label: {
		type: Array as PropType<Array<string>>,
		default: () => ['', '']
	}
})

const viewSize = computed(() => {
	let width = 0
	let height = 0
	let fontSize = 24
	let gutter = 2

	let round = props.round
	if (props.width && props.height) {
		width = props.width
		height = props.height
		fontSize = height * 0.45
	} else {
		if (props.size == 'normal') {
			width = 100
			height = 50
			fontSize = 26
		} else if (props.size == 'mini') {
			width = 80
			height = 40
			fontSize = 22
		} else if (props.size == 'large') {
			width = 120
			height = 60
			fontSize = 32
			round = 24
		}
	}
	let gutterPx = gutter
	width = Math.ceil(uni.upx2px(width))
	height = Math.ceil(uni.upx2px(height))
	let obj = {
		width: width,
		height: height,
		innerHeight: height - gutterPx * 2,
		innerWidth: width / 2 - gutterPx * 2,
		coenteWidth: width - gutterPx * 2,
		conentWidthPx: width - gutterPx * 2,
		fontSize: fontSize,
		round: round
	}

	return obj
})

const _value = ref(props.defaultValue)
const _CheckVal = computed(() => checkVal(_value.value))
const _blackValue = _CheckVal.value
const _load = ref(false)
watchEffect(() => {
	_load.value = props.load
})
async function switchClick() {
	emits('click')
	if (_load.value || props.disabled) return
	if (typeof props.beforeChecked === 'function') {
		_load.value = true
		let p = await props.beforeChecked()
		if (typeof p === 'function') {
			p = await p()
		}
		_load.value = false
		if (!p) return
	}
	_value.value = !_CheckVal.value ? props.selected : props.unSelected
	spinNvueAni(_CheckVal.value)
	emits('change', _CheckVal.value ? props.selected : props.unSelected)
	emits('update:modelValue', _CheckVal.value ? props.selected : props.unSelected)
}
watch(
	() => props.modelValue,
	(newval: boolean | string | number) => {
		_value.value = newval
		spinNvueAni(_CheckVal.value)
	}
)
onMounted(() => {
	nextTick(() => spinNvueAni(_CheckVal.value))
})
// 检验当前是否打开还是关闭状态。
function checkVal(nowVal?: number | string | boolean) {
	let val = typeof nowVal !== 'undefined' ? nowVal : props.modelValue
	if (val === props.unSelected) {
		return false
	} else if (val === props.selected) {
		return true
	}
}
function spinNvueAni(reveser = false) {
	// #ifdef APP-NVUE
	if (!proxy?.$refs['switch']) return

	var testEl = proxy?.$refs.switch
	animation.transition(
		testEl,
		{
			styles: {
				transform: reveser ? `translateX(${viewSize.value.innerWidth + 4}px)` : 'translateX(0px)',
				transformOrigin: 'center center'
			},
			duration: 250, //ms
			timingFunction: 'ease',
			delay: 0 //ms
		},
		() => {}
	)
	// #endif
}
</script>
<style scoped>
/* #ifndef APP-NVUE */
.base {
	transform: translateX(0%);
	transition: 0.2s ease;
	left: 0;
}
.on {
	left: calc(50% + 2px);
}
.off {
	left: 2px;
}
/* #endif */
</style>
<style>
.switchbgani {
	transition-timing-function: ease;
	transition-property: background-color;
	transition-duration: 0.3s;
	transition-delay: 0ms;
}
</style>
